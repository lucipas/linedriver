package main

import (
	"bytes"
	"encoding/json"
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
						"dom.webdriver.enabled": true,
						"enable_automation":     true,
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

func (c *WebDriverClient) GetSource() (string) {
    req, _ := http.NewRequest("GET", c.BaseURL+"/session/"+c.SessionID+"/source", nil)

    resp, _ := http.DefaultClient.Do(req)
    // Ensure the body is closed after reading
    defer resp.Body.Close()
    // Read the body content
    body, _ := io.ReadAll(resp.Body)
    b := string(body)
    return b
}
func main() {
	// TODO: make exe dependent on the flags and other things 

	// TODO:  Configuration: change to 9515 for Chrome or 4444 for Firefox
	client := WebDriverClient{BaseURL: "http://localhost:4444"}

//	fmt.Println("Starting session...")
	if err := client.CreateSession("firefox"); err != nil {
		fmt.Printf("Could not connect: %v\n", err)
		os.Exit(1)
	}
//	fmt.Printf("Session %s started.\n", client.SessionID)

//	fmt.Println("Navigating to https://archlinux.org...")
	client.NavigateTo("")

	f := client.GetSource()
	fmt.Println(f)

//	fmt.Println("Press Enter to close the session...")
//	fmt.Scanln()

	client.DeleteSession()
//	fmt.Println("Session closed.")
}
