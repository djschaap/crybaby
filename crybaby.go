package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Maximum number of iterations to run; zero means no limit.
	maxIter := 2

	interval := time.Second

	// Generate random UUID to identify this invocation.
	id := uuid.New()

	fmt.Printf("id=%s startup\n", id)

	pfx := "logprefix"
	slogOpts := &slog.HandlerOptions{AddSource: true}

	// sample:
	// 2023/12/12 16:59:38 INFO slog info to stderr default id=162dafad-e228-4768-a21a-8cde950d9a5b i=0 logger=slog-default dst=stderr l=info
	logStdout := log.New(os.Stdout, pfx, 0)
	logStderr := log.New(os.Stderr, pfx, 0)

	// sample:
	// time=2023-12-12T16:59:39.214Z level=INFO msg="slog info to stderr text" id=162dafad-e228-4768-a21a-8cde950d9a5b i=0 logger=slog-text dst=stderr l=info
	slogTextStdout := slog.New(slog.NewTextHandler(os.Stdout, slogOpts))
	slogTextStderr := slog.New(slog.NewTextHandler(os.Stderr, slogOpts))

	// sample:
	// {"time":"2023-12-12T16:59:40.215605213Z","level":"INFO","msg":"slog info to stderr json","id":"162dafad-e228-4768-a21a-8cde950d9a5b","i":0,"logger":"slog-json","dst":"stderr","l":"info"}
	slogJSONStdout := slog.New(slog.NewJSONHandler(os.Stdout, slogOpts))
	slogJSONStderr := slog.New(slog.NewJSONHandler(os.Stderr, slogOpts))

	for i := 0; maxIter == 0 || i < maxIter; i++ {
		fmt.Printf("id=%s i=%d\n", id, i)

		// log
		time.Sleep(interval)
		logStdout.Printf("log to stdout; id=%s i=%d logger=log dst=stdout l=info", id, i)
		logStderr.Printf("log to stderr; id=%s i=%d logger=log dst=stderr l=info", id, i)

		// slog - default (stderr only)
		time.Sleep(interval)
		slog.Info(
			"slog info to stderr default",
			"id", id,
			"i", i,
			"logger", "slog-default",
			"dst", "stderr",
			"l", "info",
		)
		slog.Warn(
			"slog warn to stderr default",
			"id", id,
			"i", i,
			"logger", "slog-default",
			"dst", "stderr",
			"l", "warn",
		)
		slog.Error(
			"slog error to stderr default",
			"id", id,
			"i", i,
			"logger", "slog-default",
			"dst", "stderr",
			"l", "error",
		)

		// slog - text
		time.Sleep(interval)
		slogTextStdout.Info(
			"slog info to stdout text",
			"id", id,
			"i", i,
			"logger", "slog-text",
			"dst", "stdout",
			"l", "info",
		)
		slogTextStderr.Info(
			"slog info to stderr text",
			"id", id,
			"i", i,
			"logger", "slog-text",
			"dst", "stderr",
			"l", "info",
		)
		slogTextStdout.Warn(
			"slog warn to stdout text",
			"id", id,
			"i", i,
			"logger", "slog-text",
			"dst", "stdout",
			"l", "warn",
		)
		slogTextStderr.Warn(
			"slog warn to stderr text",
			"id", id,
			"i", i,
			"logger", "slog-text",
			"dst", "stderr",
			"l", "warn",
		)
		slogTextStdout.Error(
			"slog error to stdout text",
			"id", id,
			"i", i,
			"logger", "slog-text",
			"dst", "stdout",
			"l", "error",
		)
		slogTextStderr.Error(
			"slog error to stderr text",
			"id", id,
			"i", i,
			"logger", "slog-text",
			"dst", "stderr",
			"l", "error",
		)

		// slog - json
		time.Sleep(interval)
		slogJSONStdout.Info(
			"slog info to stdout json",
			"id", id,
			"i", i,
			"logger", "slog-json",
			"dst", "stdout",
			"l", "info",
		)
		slogJSONStderr.Info(
			"slog info to stderr json",
			"id", id,
			"i", i,
			"logger", "slog-json",
			"dst", "stderr",
			"l", "info",
		)
		slogJSONStdout.Warn(
			"slog warn to stdout json",
			"id", id,
			"i", i,
			"logger", "slog-json",
			"dst", "stdout",
			"l", "warn",
		)
		slogJSONStderr.Warn(
			"slog warn to stderr json",
			"id", id,
			"i", i,
			"logger", "slog-json",
			"dst", "stderr",
			"l", "warn",
		)
		slogJSONStdout.Error(
			"slog error to stdout json",
			"id", id,
			"i", i,
			"logger", "slog-json",
			"dst", "stdout",
			"l", "error",
		)
		slogJSONStderr.Error(
			"slog error to stderr json",
			"id", id,
			"i", i,
			"logger", "slog-json",
			"dst", "stderr",
			"l", "error",
		)
	}

	fmt.Printf("id=%s finished\n", id)
}
