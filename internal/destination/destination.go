package destination

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GevorgGal/conduit-connector-influxdb/internal/common"
	sdk "github.com/conduitio/conduit-connector-sdk"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2api "github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBDestination struct {
	sdk.UnimplementedDestination
	client   influxdb2.Client
	writeAPI influxdb2api.WriteAPIBlocking
	config   common.Config
}

func NewInfluxDBDestination() sdk.Destination {
	return sdk.DestinationWithMiddleware(&InfluxDBDestination{}, sdk.DefaultDestinationMiddleware()...)
}

func (d *InfluxDBDestination) Configure(ctx context.Context, cfg map[string]string) error {
	config, err := common.ParseConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to parse configuration: %w", err)
	}
	d.config = config
	return nil
}

func (d *InfluxDBDestination) Open(ctx context.Context) error {
	client := influxdb2.NewClientWithOptions(d.config.URL, d.config.Token, influxdb2.DefaultOptions().SetBatchSize(20))
	d.client = client
	d.writeAPI = client.WriteAPIBlocking(d.config.Org, d.config.Bucket)
	return nil
}

func (d *InfluxDBDestination) Write(ctx context.Context, records []sdk.Record) (int, error) {
	for _, record := range records {
		var data map[string]interface{}
		if record.Operation == sdk.OperationCreate || record.Operation == sdk.OperationUpdate {
			if err := json.Unmarshal(record.Payload.After.Bytes(), &data); err != nil {
				return 0, fmt.Errorf("failed to unmarshal record payload: %w", err)
			}
		}

		point := influxdb2.NewPoint(
			"measurement",
			map[string]string{"id": string(record.Key.Bytes())},
			data,
			time.Now(),
		)

		if err := d.writeAPI.WritePoint(ctx, point); err != nil {
			return 0, fmt.Errorf("failed to write point: %w", err)
		}
	}

	return len(records), nil
}

func (d *InfluxDBDestination) Teardown(ctx context.Context) error {
	if d.client != nil {
		d.client.Close()
	}
	return nil
}
