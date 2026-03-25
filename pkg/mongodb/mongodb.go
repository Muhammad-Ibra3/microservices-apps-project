package mongodb

import (
	"context"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

type Config struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

// NewMongoDBConn Create new MongoDB client
func NewMongoDBConn(ctx context.Context, cfg *Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.URI)

	// Respect credentials embedded in URI first; only fall back to config user/password.
	if cfg.User != "" && cfg.Password != "" {
		if parsedURI, err := url.Parse(cfg.URI); err == nil && parsedURI.User == nil {
			clientOptions.SetAuth(options.Credential{
				Username: cfg.User,
				Password: cfg.Password,
			})
		}
	}

	client, err := mongo.NewClient(
		clientOptions.
			SetConnectTimeout(connectTimeout).
			SetMaxConnIdleTime(maxConnIdleTime).
			SetMinPoolSize(minPoolSize).
			SetMaxPoolSize(maxPoolSize))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
