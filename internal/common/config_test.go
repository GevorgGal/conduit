package common

import "testing"

func TestParseConfig(t *testing.T) {
	cfg := map[string]string{
		"url":    "http://localhost:8086",
		"token":  "random-token",
		"org":    "random-org",
		"bucket": "random-bucket",
	}
	_, err := ParseConfig(cfg)
	if err != nil {
		t.Errorf("ParseConfig() error = %v, wantErr %v", err, nil)
	}
}
