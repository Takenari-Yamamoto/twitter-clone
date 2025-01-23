package config

import "time"

const (
	SecretKey     = "your-secret-key" // 本番環境では環境変数から読み込むべき
	TokenDuration = 24 * time.Hour
) 