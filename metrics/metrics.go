/*
Copyright © 2025 Mad Scientist Research LLC.
This file is part of Fastbound Downloader.
*/

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// MetricsRegistry creates our own custom metrics registry with no defaults
var MetricsRegistry = prometheus.NewRegistry()

var (
	// DownloadedBooksTotal counts the total number of successful bound book downloads
	DownloadedBooksTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "fastbound_downloader_downloaded_books_total",
		Help: "The total number of successful bound book downloads",
	})

	// SkippedBookDownloadsTotal counts the total number of bound book downloads that were skipped because there was nothing new to fetch
	SkippedBookDownloadsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "fastbound_downloader_skipped_book_downloads_total",
		Help: "The total number of times the book was already downloaded or Fastbound had not generated one yet",
	})

	// FailedBookDownloadsTotal counts the total number of failed bound book downloads
	FailedBookDownloadsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "fastbound_downloader_failed_book_downloads_total",
		Help: "The total number of failed attempts at downloading a bound book",
	})
)

// A function to initialize our registry with our counters
func init() {
	MetricsRegistry.MustRegister(DownloadedBooksTotal)
	MetricsRegistry.MustRegister(SkippedBookDownloadsTotal)
	MetricsRegistry.MustRegister(FailedBookDownloadsTotal)
}
