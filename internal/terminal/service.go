package terminal

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mobile-agy/internal/auth"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Service struct {
	cmd       *exec.Cmd
	stdin     io.WriteCloser
	stdout    io.ReadCloser
	isRunning bool
	mutex     sync.Mutex
	clients   map[chan []byte]bool
	clientMux sync.Mutex
	history   []byte
}

type OpenAISettings struct {
	APIBase          string   `json:"apiBase"`
	APIKey           string   `json:"apiKey,omitempty"`
	APIKeySet        bool     `json:"apiKeySet"`
	APIKeyMasked     string   `json:"apiKeyMasked,omitempty"`
	ConfiguredModels string   `json:"configuredModels"`
	AvailableModels  []string `json:"availableModels,omitempty"`
}

func NewService() *Service {
	return &Service{
		clients: make(map[chan []byte]bool),
	}
}

func (s *Service) StartSession(workspaceDir string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.isRunning {
		return nil
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		if _, err := auth.SafeLookPath("bash"); err == nil {
			cmd = auth.SafeCommand("bash", "-i")
		} else if _, err := auth.SafeLookPath("powershell"); err == nil {
			cmd = auth.SafeCommand("powershell")
		} else {
			cmd = auth.SafeCommand("cmd")
		}
	} else {
		cmd = auth.SafeCommand("bash", "-i")
	}

	cmd.Dir = workspaceDir
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "TERM=xterm-256color")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		stdin.Close()
		return err
	}
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		stdin.Close()
		stdout.Close()
		return err
	}

	s.cmd = cmd
	s.stdin = stdin
	s.stdout = stdout
	s.isRunning = true

	// Read loop
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				data := make([]byte, n)
				copy(data, buf[:n])
				s.Broadcast(data)
			}
			if err != nil {
				break
			}
		}
		s.mutex.Lock()
		s.isRunning = false
		if s.stdin != nil {
			_ = s.stdin.Close()
		}
		if s.stdout != nil {
			_ = s.stdout.Close()
		}
		s.mutex.Unlock()
	}()

	return nil
}

func (s *Service) WriteInput(data string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.isRunning || s.stdin == nil {
		return fmt.Errorf("terminal session not running")
	}

	_, err := s.stdin.Write([]byte(data))
	return err
}

func (s *Service) RegisterClient(ch chan []byte) {
	s.clientMux.Lock()
	if s.clients == nil {
		s.clients = make(map[chan []byte]bool)
	}
	s.clients[ch] = true
	
	// Send history to new client
	if len(s.history) > 0 {
		histCopy := make([]byte, len(s.history))
		copy(histCopy, s.history)
		select {
		case ch <- histCopy:
		default:
		}
	}
	s.clientMux.Unlock()
}

func (s *Service) UnregisterClient(ch chan []byte) {
	s.clientMux.Lock()
	if s.clients != nil {
		delete(s.clients, ch)
	}
	s.clientMux.Unlock()
}

func (s *Service) Broadcast(data []byte) {
	s.clientMux.Lock()
	defer s.clientMux.Unlock()
	
	s.history = append(s.history, data...)
	if len(s.history) > 20000 {
		s.history = s.history[len(s.history)-20000:]
	}
	
	for ch := range s.clients {
		select {
		case ch <- data:
		default:
		}
	}
}

func (s *Service) KillSession() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.isRunning && s.cmd != nil && s.cmd.Process != nil {
		_ = s.cmd.Process.Kill()
	}
}

// StartCommand executes a bash command and returns its stdout/stderr reader
func (s *Service) StartCommand(ctx context.Context, command string, activeWorkspaceDir string) (*exec.Cmd, io.ReadCloser, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		if _, err := auth.SafeLookPath("bash"); err == nil {
			cmd = auth.SafeCommand("bash", "-c", command)
		} else if _, err := auth.SafeLookPath("powershell"); err == nil {
			cmd = auth.SafeCommand("powershell", "-Command", command)
		} else {
			cmd = auth.SafeCommand("cmd", "/c", command)
		}
	} else {
		cmd = auth.SafeCommand("bash", "-c", command)
	}

	cmd.Dir = activeWorkspaceDir
	cmd.Env = os.Environ()

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	return cmd, stdoutPipe, nil
}

func defaultOpenAIBase() string {
	if apiBase := os.Getenv("OPENAI_API_BASE"); apiBase != "" {
		return apiBase
	}
	return "https://api.openai.com/v1"
}

func maskAPIKey(key string) string {
	if key == "" {
		return ""
	}
	if len(key) <= 8 {
		return "********"
	}
	return key[:4] + strings.Repeat("*", len(key)-8) + key[len(key)-4:]
}

func parseConfiguredOpenAIModels(modelsEnv string) []string {
	var models []string
	for _, p := range strings.Split(modelsEnv, ",") {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			models = append(models, "openai/"+trimmed)
		}
	}
	return models
}

func (s *Service) GetOpenAISettings(fetchModels bool) OpenAISettings {
	settings := OpenAISettings{
		APIBase:          defaultOpenAIBase(),
		APIKeySet:        os.Getenv("OPENAI_API_KEY") != "",
		APIKeyMasked:     maskAPIKey(os.Getenv("OPENAI_API_KEY")),
		ConfiguredModels: os.Getenv("OPENAI_MODELS"),
	}
	if fetchModels && settings.APIKeySet {
		models, err := s.FetchOpenAIModels("", "")
		if err == nil {
			settings.AvailableModels = models
		}
	}
	return settings
}

func (s *Service) SaveOpenAISettings(apiKey, apiBase, models string, clearAPIKey bool) error {
	apiKey = strings.TrimSpace(apiKey)
	apiBase = strings.TrimSpace(apiBase)
	models = strings.TrimSpace(models)

	if apiBase == "" {
		apiBase = "https://api.openai.com/v1"
	}

	updates := map[string]*string{
		"OPENAI_API_BASE": &apiBase,
		"OPENAI_MODELS":   &models,
	}

	if clearAPIKey {
		os.Unsetenv("OPENAI_API_KEY")
		updates["OPENAI_API_KEY"] = nil
	} else if apiKey != "" {
		os.Setenv("OPENAI_API_KEY", apiKey)
		updates["OPENAI_API_KEY"] = &apiKey
	}

	os.Setenv("OPENAI_API_BASE", apiBase)
	if models == "" {
		os.Unsetenv("OPENAI_MODELS")
	} else {
		os.Setenv("OPENAI_MODELS", models)
	}

	return updateEnvFile(".env", updates)
}

func updateEnvFile(path string, updates map[string]*string) error {
	seen := map[string]bool{}
	var lines []string

	file, err := os.Open(path)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, "#") || !strings.Contains(line, "=") {
				lines = append(lines, line)
				continue
			}
			key := strings.TrimSpace(strings.SplitN(line, "=", 2)[0])
			if val, ok := updates[key]; ok {
				seen[key] = true
				if val == nil {
					continue
				}
				lines = append(lines, key+"="+strconv.Quote(*val))
				continue
			}
			lines = append(lines, line)
		}
		if scanErr := scanner.Err(); scanErr != nil {
			_ = file.Close()
			return scanErr
		}
		_ = file.Close()
	} else if !os.IsNotExist(err) {
		return err
	}

	for key, val := range updates {
		if seen[key] || val == nil {
			continue
		}
		lines = append(lines, key+"="+strconv.Quote(*val))
	}

	content := strings.Join(lines, "\n")
	if content != "" {
		content += "\n"
	}
	return os.WriteFile(path, []byte(content), 0600)
}

type OpenAIKeyPoolEntry struct {
	ID           string   `json:"id"`
	Label        string   `json:"label"`
	APIKey       string   `json:"apiKey,omitempty"`
	APIKeyMasked string   `json:"apiKeyMasked,omitempty"`
	APIBase      string   `json:"apiBase"`
	Models       string   `json:"models"`
	IsActive     bool     `json:"isActive"`
	CreatedAt    int64    `json:"createdAt"`
}

func getOpenAIKeysPoolFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(homeDir, ".gemini", "antigravity-cli")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "openai_keys_pool.json"), nil
}

func (s *Service) LoadOpenAIKeysPool() ([]OpenAIKeyPoolEntry, error) {
	path, err := getOpenAIKeysPoolFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []OpenAIKeyPoolEntry{}, nil
		}
		return nil, err
	}
	var pool []OpenAIKeyPoolEntry
	if err := json.Unmarshal(data, &pool); err != nil {
		return []OpenAIKeyPoolEntry{}, nil
	}
	return pool, nil
}

func (s *Service) SaveOpenAIKeysPool(pool []OpenAIKeyPoolEntry) error {
	path, err := getOpenAIKeysPoolFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(pool, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func (s *Service) AddOrUpdateOpenAIKey(label, apiKey, apiBase, models string, setActive bool) error {
	label = strings.TrimSpace(label)
	apiKey = strings.TrimSpace(apiKey)
	apiBase = strings.TrimSpace(apiBase)
	models = strings.TrimSpace(models)

	if apiBase == "" {
		apiBase = "https://api.openai.com/v1"
	}
	if label == "" {
		if strings.Contains(apiBase, "deepseek") {
			label = "DeepSeek AI"
		} else if strings.Contains(apiBase, "groq") {
			label = "Groq High-Speed"
		} else if strings.Contains(apiBase, "openrouter") {
			label = "OpenRouter AI"
		} else if strings.Contains(apiBase, "localhost") || strings.Contains(apiBase, "127.0.0.1") {
			label = "Ollama Local"
		} else {
			label = "OpenAI Provider"
		}
	}

	pool, _ := s.LoadOpenAIKeysPool()

	existingIdx := -1
	for i, entry := range pool {
		if (apiKey != "" && entry.APIKey == apiKey && entry.APIBase == apiBase) || (label != "" && entry.Label == label && entry.APIBase == apiBase) {
			existingIdx = i
			break
		}
	}

	entryID := fmt.Sprintf("key-%d", time.Now().UnixNano())
	if existingIdx >= 0 {
		entryID = pool[existingIdx].ID
		if apiKey == "" {
			apiKey = pool[existingIdx].APIKey
		}
	}

	if setActive {
		for i := range pool {
			pool[i].IsActive = false
		}
	}

	newEntry := OpenAIKeyPoolEntry{
		ID:        entryID,
		Label:     label,
		APIKey:    apiKey,
		APIBase:   apiBase,
		Models:    models,
		IsActive:  setActive,
		CreatedAt: time.Now().Unix(),
	}

	if existingIdx >= 0 {
		pool[existingIdx] = newEntry
	} else {
		pool = append(pool, newEntry)
	}

	if err := s.SaveOpenAIKeysPool(pool); err != nil {
		return err
	}

	if setActive {
		return s.SaveOpenAISettings(apiKey, apiBase, models, false)
	}
	return nil
}

func (s *Service) SetActiveOpenAIKey(id string) error {
	pool, err := s.LoadOpenAIKeysPool()
	if err != nil {
		return err
	}

	var activeEntry *OpenAIKeyPoolEntry
	for i := range pool {
		if pool[i].ID == id {
			pool[i].IsActive = true
			activeEntry = &pool[i]
		} else {
			pool[i].IsActive = false
		}
	}

	if activeEntry == nil {
		return fmt.Errorf("API key dengan ID '%s' tidak ditemukan", id)
	}

	if err := s.SaveOpenAIKeysPool(pool); err != nil {
		return err
	}

	return s.SaveOpenAISettings(activeEntry.APIKey, activeEntry.APIBase, activeEntry.Models, false)
}

func (s *Service) DeleteOpenAIKey(id string) error {
	pool, err := s.LoadOpenAIKeysPool()
	if err != nil {
		return err
	}

	newPool := []OpenAIKeyPoolEntry{}
	wasActive := false
	for _, entry := range pool {
		if entry.ID == id {
			if entry.IsActive {
				wasActive = true
			}
			continue
		}
		newPool = append(newPool, entry)
	}

	if wasActive && len(newPool) > 0 {
		newPool[0].IsActive = true
		_ = s.SaveOpenAISettings(newPool[0].APIKey, newPool[0].APIBase, newPool[0].Models, false)
	} else if wasActive && len(newPool) == 0 {
		_ = s.SaveOpenAISettings("", "https://api.openai.com/v1", "", true)
	}

	return s.SaveOpenAIKeysPool(newPool)
}

func (s *Service) GetOpenAIKeysPoolForClient() []OpenAIKeyPoolEntry {
	pool, _ := s.LoadOpenAIKeysPool()
	activeKey := os.Getenv("OPENAI_API_KEY")
	activeBase := os.Getenv("OPENAI_API_BASE")

	clientPool := make([]OpenAIKeyPoolEntry, len(pool))
	for i, entry := range pool {
		isActive := entry.IsActive
		if activeKey != "" && entry.APIKey == activeKey && entry.APIBase == activeBase {
			isActive = true
		}
		clientPool[i] = OpenAIKeyPoolEntry{
			ID:           entry.ID,
			Label:        entry.Label,
			APIKeyMasked: maskAPIKey(entry.APIKey),
			APIBase:      entry.APIBase,
			Models:       entry.Models,
			IsActive:     isActive,
			CreatedAt:    entry.CreatedAt,
		}
	}
	return clientPool
}

func (s *Service) FetchOpenAIModels(apiKey, apiBase string) ([]string, error) {
	apiKey = strings.TrimSpace(apiKey)
	apiBase = strings.TrimSpace(apiBase)
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}
	if apiBase == "" {
		apiBase = defaultOpenAIBase()
	}
	if apiKey == "" && !strings.Contains(apiBase, "localhost") && !strings.Contains(apiBase, "127.0.0.1") {
		return nil, fmt.Errorf("API key belum diisi")
	}

	url := strings.TrimSuffix(apiBase, "/") + "/models"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; AGY-Mobile-IDE/1.1.2)")

	dialer := &net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := dialer.DialContext(ctx, network, addr)
			if err == nil {
				return conn, nil
			}

			host, port, splitErr := net.SplitHostPort(addr)
			if splitErr != nil {
				return nil, err
			}

			resolver := &net.Resolver{
				PreferGo: true,
				Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
					d := net.Dialer{Timeout: 3 * time.Second}
					c, e := d.DialContext(ctx, "udp", "1.1.1.1:53")
					if e == nil {
						return c, nil
					}
					c2, e2 := d.DialContext(ctx, "udp", "8.8.8.8:53")
					if e2 == nil {
						return c2, nil
					}
					return d.DialContext(ctx, "udp", "8.8.4.4:53")
				},
			}

			ips, lookupErr := resolver.LookupHost(ctx, host)
			if lookupErr != nil || len(ips) == 0 {
				return nil, err
			}

			for _, ip := range ips {
				c, e := dialer.DialContext(ctx, "tcp4", net.JoinHostPort(ip, port))
				if e == nil {
					return c, nil
				}
				c2, e2 := dialer.DialContext(ctx, "tcp", net.JoinHostPort(ip, port))
				if e2 == nil {
					return c2, nil
				}
			}
			return nil, err
		},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke endpoint %s: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("endpoint %s mengembalikan status %d: %s", url, resp.StatusCode, string(body))
	}

	var parsed struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf("gagal parse balasan JSON: %v", err)
	}

	var models []string
	for _, item := range parsed.Data {
		id := strings.TrimSpace(item.ID)
		if id != "" {
			models = append(models, "openai/"+id)
		}
	}
	return models, nil
}

// GetModelsList fetches available models from agy CLI or falls back to defaults
func (s *Service) GetModelsList() ([]string, error) {
	var models []string

	hasToken := false
	homeDir, errToken := os.UserHomeDir()
	if errToken == nil {
		tokenPath := filepath.Join(homeDir, ".gemini", "antigravity-cli", "antigravity-oauth-token")
		if _, errStat := os.Stat(tokenPath); errStat == nil {
			hasToken = true
		}
	}

	if hasToken {
		agyPath := auth.FindAgyPath()
		var outputBytes []byte
		var err error

		useDirect := false
		if _, lookErr := auth.SafeLookPath("script"); lookErr != nil {
			useDirect = true
		}

		if useDirect {
			cmdDirect := auth.SafeCommand(agyPath, "models")
			cmdDirect.Env = os.Environ()
			outputBytes, err = cmdDirect.Output()
		} else {
			cmdStr := fmt.Sprintf("%s models", agyPath)
			cmd := auth.SafeCommand("script", "-q", "-f", "-c", cmdStr, "/dev/null")
			cmd.Env = os.Environ()
			outputBytes, err = cmd.Output()

			if err != nil {
				cmdDirect := auth.SafeCommand(agyPath, "models")
				cmdDirect.Env = os.Environ()
				outputBytes, err = cmdDirect.Output()
			}
		}

		if err == nil {
			lines := strings.Split(string(outputBytes), "\n")
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if trimmed == "" {
					continue
				}
				if strings.Contains(trimmed, "Fetching") || strings.Contains(trimmed, "⠋") || strings.Contains(trimmed, "⠙") || strings.Contains(trimmed, "⠹") || strings.Contains(trimmed, "⠸") || strings.Contains(trimmed, "⠼") || strings.Contains(trimmed, "⠴") || strings.Contains(trimmed, "⠦") || strings.Contains(trimmed, "⠧") || strings.Contains(trimmed, "⠇") || strings.Contains(trimmed, "⠏") {
					continue
				}
				fields := strings.Fields(trimmed)
				if len(fields) > 0 {
					models = append(models, fields[0])
				}
			}
		}
	}

	if len(models) == 0 {
		models = []string{
			"gemini-3.5-flash-high",
			"gemini-3.5-flash-medium",
			"gemini-3.5-flash-low",
			"gemini-3.1-pro-high",
			"gemini-3.1-pro-low",
			"claude-sonnet-4-6",
			"claude-opus-4-6-thinking",
			"gpt-oss-120b-medium",
		}
	}

	// Append OpenAI-compatible models if configured. Prefer the live /models
	// endpoint so the dropdown reflects models available for the configured key.
	if os.Getenv("OPENAI_API_KEY") != "" {
		openAIModels, fetchErr := s.FetchOpenAIModels("", "")
		if fetchErr == nil && len(openAIModels) > 0 {
			models = append(models, openAIModels...)
		} else if configured := parseConfiguredOpenAIModels(os.Getenv("OPENAI_MODELS")); len(configured) > 0 {
			models = append(models, configured...)
		} else {
			models = append(models,
				"openai/gpt-4o",
				"openai/gpt-4o-mini",
				"openai/deepseek-chat",
				"openai/deepseek-reasoner",
			)
		}
	}

	return models, nil
}
