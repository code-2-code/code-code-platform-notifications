package main

import (
	"log/slog"
	"os"
	"strings"

	"code-code.internal/platform-k8s/internal/wecomcallback"
)

func main() {
	publisher, err := wecomcallback.NewNATSPublisher(
		envOrDefault("WECOM_CALLBACK_NATS_URL", "nats://nats.code-code-infra.svc.cluster.local:4222"),
		envOrDefault("WECOM_CALLBACK_NATS_SUBJECT", "platform.wecom.messages.received"),
	)
	if err != nil {
		slog.Error("fatal error", "error", err)
		os.Exit(1)
	}
	defer publisher.Close()

	server, err := wecomcallback.NewServer(wecomcallback.Config{
		Addr:           envOrDefault("WECOM_CALLBACK_HTTP_ADDR", ":8080"),
		Path:           envOrDefault("WECOM_CALLBACK_PATH", "/wecom/callback"),
		Token:          requiredEnv("WECOM_CALLBACK_TOKEN"),
		EncodingAESKey: requiredEnv("WECOM_CALLBACK_ENCODING_AES_KEY"),
	}, publisher, slog.Default())
	if err != nil {
		slog.Error("fatal error", "error", err)
		os.Exit(1)
	}
	slog.Info("wecom-callback-adapter listening", "addr", envOrDefault("WECOM_CALLBACK_HTTP_ADDR", ":8080"), "path", envOrDefault("WECOM_CALLBACK_PATH", "/wecom/callback"))
	if err := server.ListenAndServe(); err != nil {
		slog.Error("fatal error", "error", err)
		os.Exit(1)
	}
}

func envOrDefault(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func requiredEnv(key string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		slog.Error("required environment variable is missing", "key", key)
		os.Exit(1)
	}
	return value
}
