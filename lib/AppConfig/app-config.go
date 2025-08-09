package appconfig

import (
	"bushuray-tui/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type AppConfig struct {
	SocksPort     int       `json:"socks-port"`
	HttpPort      int       `json:"http-port"`
	CoreTCPPort   int       `json:"core-tcp-port"`
	TestPortRange PortRange `json:"test-port-range"`
	NoBackground  bool      `json:"no-background,omitzero"`
}

type PortRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

var is_config_loaded bool = false

var application_configuration AppConfig = defaultConfig()

func defaultConfig() AppConfig {
	return AppConfig{
		SocksPort:   3090,
		HttpPort:    3091,
		CoreTCPPort: 4897,
		TestPortRange: PortRange{
			Start: 3095,
			End:   30120,
		},
		NoBackground: false,
	}
}

func GetConfig() AppConfig {
	if !is_config_loaded {
		panic("attempted to access config before loading")
	}
	return application_configuration
}

func LoadConfig() {
	config, err := readConfig()
	if err != nil {
		log.Println("failed to read config file:", err, "using default config")
	}
	application_configuration = config
	is_config_loaded = true
}

func readConfig() (AppConfig, error) {
	var default_config AppConfig = defaultConfig()
	home_dir, err := utils.GetHomeDir()
	if err != nil {
		return default_config, err
	}
	var dir_path = filepath.Join(home_dir, ".config", "bushuray")
	var config_path = filepath.Join(dir_path, "config.json")
	file_bytes, err := os.ReadFile(config_path)
	if err == nil {
		if err := json.Unmarshal(file_bytes, &default_config); err != nil {
			return default_config, fmt.Errorf("failed to parse config file, invalid json")
		}
		return default_config, nil
	}

	if !os.IsNotExist(err) {
		return default_config, fmt.Errorf("failed to read config file")
	}

	if err := os.MkdirAll(dir_path, 0777); err != nil {
		return default_config, fmt.Errorf("config file was not found and failed to create config directory " + dir_path + ": " + err.Error())
	}

	default_config_json, err := json.MarshalIndent(defaultConfig(), "", " ")

	if err != nil {
		return default_config, fmt.Errorf("config file was not found and failed to parse default config: " + err.Error())
	}

	if err := os.WriteFile(config_path, default_config_json, 0666); err != nil {
		return default_config, fmt.Errorf("config file was not found and failed to write default config: " + err.Error())
	}

	return default_config, nil
}
