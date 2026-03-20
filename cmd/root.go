/*
Copyright © 2025 Mad Scientist Research LLC.
This file is part of Fastbound Downloader.
*/

package cmd

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ScienceGuns/fastbound-downloader/apis/fbdownloader_settings"
	"github.com/ScienceGuns/fastbound-downloader/metrics"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fbdownloader",
	Short: functionHelpShort,
	Long: fmt.Sprintf(`%s

Version: %s`, functionHelpLong, shortVersion),
	// Run root's command normally
	Run: func(cmd *cobra.Command, args []string) {
		settings := pullSettings()

		// Start the Prometheus metrics server only if not disabled by one or more flags that prevent the functionality
		if !settings.IsCron && !settings.DisableMetrics {
			log.Printf("Metrics server starting on %s\n", settings.MetricsPort)
			go func() {
				http.Handle("/metrics", promhttp.HandlerFor(metrics.MetricsRegistry, promhttp.HandlerOpts{}))
				log.Fatal(http.ListenAndServe(settings.MetricsPort, nil))
			}()
			log.Printf("Waiting 5 minutes before scanning\n")
			time.Sleep(5 * time.Minute)
		}

		if settings.IsCron {
			// Run a cycle and exit
			rotationCycle(settings)
			os.Exit(0)
		} else {
			// Run an initial cycle immediately
			now := time.Now().Format("2006-01-02 15:04 MST")
			log.Printf("Running a cycle at %s\n", now)
			rotationCycle(settings)
			// Run future cycles only AFTER the interval has occurred
			ticker := time.NewTicker(time.Duration(settings.ScanningIntervalInMinutes) * time.Minute)
			for range ticker.C {
				now := time.Now().Format("2006-01-02 15:04 MST")
				log.Printf("Running a cycle at %s\n", now)
				rotationCycle(settings)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&SettingsFilePath, "settings-path", "/config/settings.json", "OPTIONAL: Specify an alternate settings file path.")
}

// pullSettings Loads the settings from file and outputs them
func pullSettings() fbdownloader_settings.FBDConfig {
	// Check if the settings file exists and has the correct mode
	fbdownloader_settings.CheckForSettingsFile(SettingsFilePath)
	Settings, err := fbdownloader_settings.ReadSettingsFile(SettingsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return *Settings
}
