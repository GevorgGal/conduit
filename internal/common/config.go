package common

import (
	"fmt"
)

type Config struct {
	URL    string
	Token  string
	Org    string
	Bucket string
}

func ParseConfig(cfg map[string]string) (Config, error) {
	url, ok := cfg["url"]
	if !ok {
		return Config{}, fmt.Errorf("url configuration is required")
	}
	token, ok := cfg["token"]
	if !ok {
		return Config{}, fmt.Errorf("token configuration is required")
	}
	org, ok := cfg["org"]
	if !ok {
		return Config{}, fmt.Errorf("org configuration is required")
	}
	bucket, ok := cfg["bucket"]
	if !ok {
		return Config{}, fmt.Errorf("bucket configuration is required")
	}

	return Config{
		URL:    url,
		Token:  token,
		Org:    org,
		Bucket: bucket,
	}, nil
}
