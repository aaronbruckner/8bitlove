package config

import (
	"encoding/json"
	"errors"
	"os"
)

type AuthConfig struct {
	Aws_access_key string
	Aws_secret_key string
}

var ErrLocalAuthConfigNotPresent = errors.New("ErrLocalAuthConfigNotPresent")

const AUTH_CONFIG_FILE_NAME = ".localAuthConfig.json"

func LoadLocalAuthConfigIfPresent() (AuthConfig, error) {
	var config AuthConfig

	authBytes, err := os.ReadFile(AUTH_CONFIG_FILE_NAME)
	if err != nil {
		return config, ErrLocalAuthConfigNotPresent
	}

	err = json.Unmarshal(authBytes, &config)

	if err != nil {
		return config, ErrLocalAuthConfigNotPresent
	}

	return config, nil
}

func SaveLocalAuthConfig(config AuthConfig) {
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(AUTH_CONFIG_FILE_NAME, jsonBytes, 0666)
	if err != nil {
		panic(err)
	}
}
