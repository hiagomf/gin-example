package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	HTTPPort int `json:"http_port"`
}

var singleton *configuration

func Load() error {
	bytesJson, err := os.ReadFile("../config/config.json")
	if err != nil {
		return fmt.Errorf("erro ao ler json: %s", err.Error())
	}

	singleton = &configuration{}
	if err := json.Unmarshal(bytesJson, singleton); err != nil {
		return fmt.Errorf("erro ao decodificar json: %s", err.Error())
	}

	return nil
}

func Get() *configuration {
	return singleton
}
