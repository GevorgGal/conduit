package source

import (
	"context"
	"testing"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

func TestInfluxDBSource_ConfigureAndOpen(t *testing.T) {
	src := NewInfluxDBSource()
	cfg := map[string]string{
		"url":    "http://fakeurl:9999",
		"token":  "fake-token",
		"org":    "fake-org",
		"bucket": "fake-bucket",
	}
	err := src.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Configure failed: %v", err)
	}

	position := sdk.Position("")

	if err := src.Open(context.Background(), position); err != nil {
		t.Errorf("Open() error = %v, wantErr %v", err, nil)
	}
}
