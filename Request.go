package stormiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 注册函数
func register(username, password string) error {
	url := serveraddr + "/register"
	_, _, err := Request(url, username, password)
	return err
}

var serveraddr = "http://118.25.196.166:8379"

// 登录函数
func login(username, password string) (string, string, error) {
	url := serveraddr + "/login"
	redis_addr, password, err := Request(url, username, password)
	return redis_addr, password, err
}

func Request(url, username, password string) (string, string, error) {
	data := map[string]string{
		"username": username,
		"password": password,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", "", fmt.Errorf("error marshaling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("error decoding response: %v", err)
	}

	// 获取redis-addr
	if redisAddr, ok := result["redis-addr"].(string); ok {
		// 获取password
		if password, ok := result["redis-password"].(string); ok {
			return redisAddr, password, nil
		}
	}

	return "", "", nil
}
