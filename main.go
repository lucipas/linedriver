package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WebDriverClient struct {
	BaseURL   string
	SessionID string
}

// CreateSession initiates a new WebDriver session
func (c *WebDriverClient) CreateSession(browserName string) error {
	payload := map[string]interface{}{
		"capabilities": map[string]interface{}{
			"alwaysMatch": map[string]interface{}{
				"browserName": browserName,
				"moz:firefoxOptions": map[string]interface{}{
					"prefs": map[string]interface{}{
						"dom.webdriver.enabled": false,
						"enable_automation":     false,
					},
					"args": []string{
						// TODO: Put this into a config
						"--profile", "/home/aerialtramway/.config/mozilla/firefox/m50jazby.default-release/",
					},
				},
			},
		},
	}
	data, _ := json.Marshal(payload)

	resp, err := http.Post(c.BaseURL+"/session", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	val, ok := result["value"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to parse session response: %v", result)
	}
	c.SessionID = val["sessionId"].(string)
	return nil
}

// NavigateTo sends a URL navigation command
func (c *WebDriverClient) NavigateTo(url string) {
	payload := map[string]string{"url": url}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", c.BaseURL+"/session/"+c.SessionID+"/url", bytes.NewBuffer(data))
	http.DefaultClient.Do(req)
}

// DeleteSession closes the browser
func (c *WebDriverClient) DeleteSession() {
	req, _ := http.NewRequest("DELETE", c.BaseURL+"/session/"+c.SessionID, nil)
	http.DefaultClient.Do(req)
}

func (c *WebDriverClient) GetSource() string {
	req, _ := http.NewRequest("GET", c.BaseURL+"/session/"+c.SessionID+"/source", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
	// Read the body content
	body, _ := io.ReadAll(resp.Body)
	b := string(body)
	return b
}

func (c *WebDriverClient) GetTitle() string {
	req, _ := http.NewRequest("GET", c.BaseURL+"/session/"+c.SessionID+"/title", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
	// Read the body content
	body, _ := io.ReadAll(resp.Body)
	b := string(body)
	return b
}

func (c *WebDriverClient) GetUrl() string {
	req, _ := http.NewRequest("GET", c.BaseURL+"/session/"+c.SessionID+"/url", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
	// Read the body content
	body, _ := io.ReadAll(resp.Body)
	b := string(body)
	return b
}

func (c *WebDriverClient) NewWin() {
	req, _ := http.NewRequest("POST", c.BaseURL+"/session/"+c.SessionID+"/window/new", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
}

func (c *WebDriverClient) FindElement(sType, selector string) string {
	payload := map[string]string{"using": sType, "value": selector}
	data, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", c.BaseURL+"/session/"+c.SessionID+"/element", bytes.NewBuffer(data))
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	// Read the body content
	body, _ := io.ReadAll(resp.Body)
	b := string(body)
	return b
}

func (c *WebDriverClient) Screenshot() string {
	req, _ := http.NewRequest("GET", c.BaseURL+"/session/"+c.SessionID+"/screenshot", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
	// Read the body content
	body, _ := io.ReadAll(resp.Body)
	b := string(body)
	return b
}

func (c *WebDriverClient) Refresh() {
	req, _ := http.NewRequest("POST", c.BaseURL+"/session/"+c.SessionID+"/refresh", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
}


func (c *WebDriverClient) Back() {
	req, _ := http.NewRequest("POST", c.BaseURL+"/session/"+c.SessionID+"/back", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
}
func (c *WebDriverClient) Forward() {
	req, _ := http.NewRequest("POST", c.BaseURL+"/session/"+c.SessionID+"/forward", nil)

	resp, _ := http.DefaultClient.Do(req)
	// Ensure the body is closed after reading
	defer resp.Body.Close()
}






func main() {
	// Define CLI Flags
	port := 0
	browserPtr := flag.String("b", "firefox", "Browser name (firefox/chrome)")
	ctxPtr := flag.String("ctx", "", "Browser Context to automate")
	flag.Parse()

	if(*browserPtr == "firefox") {
		port = 4444
	} else if(*browserPtr == "chrome") {
		port = 37679
	}

	client := WebDriverClient{BaseURL: fmt.Sprintf("http://localhost:%d", port)}
	if(*ctxPtr == "") {
		client.CreateSession(*browserPtr)
		fmt.Println(client.SessionID)
		os.Exit(0)
	} else {
		client.SessionID = *ctxPtr
	}
	// Initialize

	switch flag.Arg(0) {
	case "navigate": {
		client.NavigateTo(flag.Arg(1))
	}
	case "close": {
		client.DeleteSession()
	}
	case "src": {
		out := client.GetSource()
		fmt.Println(out)
	}
	case "url": {
		out := client.GetUrl()
		fmt.Println(out)
	}
	case "title": {
		out := client.GetTitle()
		fmt.Println(out)
	}
	case "refresh": {
		client.Refresh()
	}
	case "back": {
		client.Back()
	}
	case "forward": {
		client.Forward()
	}
}
