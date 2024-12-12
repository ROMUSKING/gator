package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	CurrentUserName string `json:"current_user_name"`
	DBURL           string `json:"db_url"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	var cfg Config
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("can't get home dir %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("can't get config file: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("can't read config file: %v", err)
	}
	return cfg, nil
}

func (cfg *Config) SetUser(user_name string) error {
	cfg.CurrentUserName = user_name

	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("can't write config file: %v", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {

	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = fmt.Sprintf("%v/%v", path, configFileName)
	return path, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, dat, 0666)

	if err != nil {
		return err
	}
	return nil
}
