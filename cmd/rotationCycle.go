/*
Copyright © 2025 Mad Scientist Research LLC.
This file is part of Fastbound Downloader.
*/

package cmd

import (
	"github.com/ScienceGuns/fastbound-downloader/apis/fastbound"
	"github.com/ScienceGuns/fastbound-downloader/apis/fbdownloader_settings"
	"github.com/ScienceGuns/fastbound-downloader/metrics"
	"log"
)

// rotationCycle This function runs the core logic of the Fastbound Downloader
func rotationCycle(settings fbdownloader_settings.FBDConfig) {
	log.Printf("Downloading the latest bound book for account %s\n", settings.Fastbound.AccountNumber)
	// Download the daily Bound Book
	downloadedBook, err := fastbound.DownloadBoundBook(fastboundAPIBaseURL, settings)
	if err != nil {
		metrics.FailedBookDownloadsTotal.Inc()
		return
	}
	if downloadedBook != "" {
		metrics.DownloadedBooksTotal.Inc()
		log.Printf("Downloaded the bound book %s\n", downloadedBook)
	} else {
		metrics.SkippedBookDownloadsTotal.Inc()
	}
}
