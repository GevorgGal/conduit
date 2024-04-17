package destination

import (
	"context"
	"testing"
)

func TestInfluxDBDestination_ConfigureAndOpen(t *testing.T) {
	dest := NewInfluxDBDestination()
	cfg := map[string]string{
		"url":    "http://fakeurl:9999",
		"token":  "fake-token",
		"org":    "fake-org",
		"bucket": "fake-bucket",
	}
	err := dest.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Configure failed: %v", err)
	}

	if err := dest.Open(context.Background()); err != nil {
		t.Errorf("Open() error = %v, wantErr %v", err, nil)
	}
}
