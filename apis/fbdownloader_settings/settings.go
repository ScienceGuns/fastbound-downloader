/*
Copyright © 2025 Mad Scientist Research LLC.
This file is part of Fastbound Downloader.
*/

package fbdownloader_settings

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// FBDConfig A struct to keep track of known values in settings.json
type FBDConfig struct {
	Fastbound struct {
		AccountNumber string `json:"account-number"`
		ApiKey        string `json:"api-key"`
		AuditUser     string `json:"audit-user"`
	} `json:"fastbound"`
	Paths struct {
		BoundBooks string `json:"bound-books"`
	} `json:"paths"`
	IsCron                    bool   `json:"is-cron,omitempty"`
	DisableMetrics            bool   `json:"disable-metrics,omitempty"`
	MetricsPort               string `json:"metrics-port,omitempty"`
	ScanningIntervalInMinutes uint   `json:"scanning-interval,omitempty"`
}

// CheckForSettingsFile Check if the settings file exists and has the correct mode
func CheckForSettingsFile(settingsFilePath string) {
	settingsFile, err := os.Stat(settingsFilePath)
	if err != nil {
		// If the error is not nil, check if this is because the file doesn't exist
		if os.IsNotExist(err) {
			log.Fatalf("Unable to find settings file: %v", err)
		} else {
			log.Fatalf("Unable to read settings file: %v", err)
		}
	}
	// If the mode is incorrect then we exit with warning to rotate credentials
	if settingsFile.Mode() != 0400 {
		log.Fatalf("Settings file detected but mode is not 0400!\n" +
			"You SHOULD rotate any credentials in the file after fixing mode.")
	}
}

// validateSettingsFile Validate that the contents of the settings file are sane
func validateSettingsFile(settings FBDConfig) error {
	if len(settings.Fastbound.AccountNumber) < 6 {
		return fmt.Errorf("fastbound account number appears to be in the wrong format")
	}
	if len(settings.Fastbound.ApiKey) == 0 {
		return fmt.Errorf("fastbound API key appears to be blank")
	}
	if settings.Paths.BoundBooks == "" {
		return fmt.Errorf("bound book path seems to be invalid")
	}
	return nil
}

// ReadSettingsFile Read the settings.json file and store the data
func ReadSettingsFile(settingsFilePath string) (*FBDConfig, error) {
	fileData, err := os.ReadFile(settingsFilePath)
	if err != nil {
		return nil, fmt.Errorf("failure reading discovered config file: %v", err)
	}

	var outputConfig FBDConfig
	err = json.Unmarshal(fileData, &outputConfig)
	if err != nil {
		return nil, fmt.Errorf("failure reading discovered config file: %v", err)
	}

	// Set default metrics port if left unconfigured
	if outputConfig.MetricsPort == "" {
		outputConfig.MetricsPort = ":9090"
	} else {
		// Ensure the port string starts with a colon as we will assume this elsewhere
		if outputConfig.MetricsPort[0] != ':' {
			outputConfig.MetricsPort = ":" + outputConfig.MetricsPort
		}
	}

	// Set default scanning interval to 1440 minutes (1 day) if left unconfigured
	if outputConfig.ScanningIntervalInMinutes == 0 {
		outputConfig.ScanningIntervalInMinutes = 1440
	}

	// Validate settings config
	err = validateSettingsFile(outputConfig)
	if err != nil {
		return nil, err
	}
	return &outputConfig, nil
}
