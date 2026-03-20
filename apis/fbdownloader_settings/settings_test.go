/*
Copyright © 2025 Mad Scientist Research LLC.
This file is part of Fastbound Downloader.
*/

package fbdownloader_settings

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

// TestCheckForSettingsFile validate the CheckForSettingsFile function
func TestCheckForSettingsFile(t *testing.T) {
	// Create a temporary file to store settings in
	tempFile, err := os.CreateTemp("", "settings.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Warning: failed to remove temp file %s: %v", tempFile.Name(), err)
		}
	}()

	err = tempFile.Chmod(0400)
	if err != nil {
		log.Fatal(err)
	}

	// Test that the file exists and the mode is correct
	t.Run("Test with correct mode", func(t *testing.T) {
		CheckForSettingsFile(tempFile.Name())
	})
}

// TestValidateSettingsFile validate the validateSettingsFile function
func TestValidateSettingsFile(t *testing.T) {
	// Create a list of test configs to validate as pass/fail
	tests := []struct {
		name     string
		settings FBDConfig
		wantErr  bool
	}{
		{
			name: "Valid settings.json",
			settings: FBDConfig{
				Fastbound: struct {
					AccountNumber string `json:"account-number"`
					ApiKey        string `json:"api-key"`
					AuditUser     string `json:"audit-user"`
				}{
					AccountNumber: "123456",
					ApiKey:        "kkJ4K3dHoHqZzNvoDJ",
					AuditUser:     "pgibbons@initech.com",
				},
				Paths: struct {
					BoundBooks string `json:"bound-books"`
				}{
					BoundBooks: "/books/",
				},
				IsCron:                    false,
				DisableMetrics:            false,
				MetricsPort:               ":9090",
				ScanningIntervalInMinutes: 1440,
			},
			wantErr: false,
		},
		{
			name: "Invalid settings.json",
			settings: FBDConfig{
				Fastbound: struct {
					AccountNumber string `json:"account-number"`
					ApiKey        string `json:"api-key"`
					AuditUser     string `json:"audit-user"`
				}{
					AccountNumber: "123456",
					ApiKey:        "",
					AuditUser:     "pgibbons@initech.com",
				},
				Paths: struct {
					BoundBooks string `json:"bound-books"`
				}{
					BoundBooks: "/books/",
				},
				IsCron:                    false,
				DisableMetrics:            false,
				MetricsPort:               ":9090",
				ScanningIntervalInMinutes: 1440,
			},
			wantErr: true,
		},
	}
	// Loop through the tests and see if they behave as expected
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateSettingsFile(test.settings)
			if (err != nil) != test.wantErr {
				t.Errorf("validateSettingsFile() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

// TestReadSettingsFile validate the ReadSettingsFile function
func TestReadSettingsFile(t *testing.T) {
	// Create a temporary file to store settings in
	tempFile, err := os.CreateTemp("", "settings.json")
	if err != nil {
		t.Errorf("Could not create temp file: %v", err)
	}
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Logf("Warning: failed to remove temp file %s: %v", tempFile.Name(), err)
		}
	}()

	// Create test settings
	testConfig := FBDConfig{
		Fastbound: struct {
			AccountNumber string `json:"account-number"`
			ApiKey        string `json:"api-key"`
			AuditUser     string `json:"audit-user"`
		}{
			AccountNumber: "123456",
			ApiKey:        "kkJ4K3dHoHqZzNvoDJ",
			AuditUser:     "pgibbons@initech.com",
		},
		Paths: struct {
			BoundBooks string `json:"bound-books"`
		}{
			BoundBooks: "/books/",
		},
		IsCron:                    false,
		DisableMetrics:            false,
		MetricsPort:               ":9090",
		ScanningIntervalInMinutes: 1440,
	}
	// Write the test settings to the file
	jsonData, _ := json.MarshalIndent(testConfig, "", " ")
	_ = os.WriteFile(tempFile.Name(), jsonData, 0644)

	// Validate by reading the test settings file
	t.Run("Test valid settings", func(t *testing.T) {
		_, err := ReadSettingsFile(tempFile.Name())
		if err != nil {
			t.Errorf("Expected no error but got: %v", err)
		}
	})
}
