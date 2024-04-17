package source

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GevorgGal/conduit-connector-influxdb/internal/common"
	sdk "github.com/conduitio/conduit-connector-sdk"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBSource struct {
	sdk.UnimplementedSource
	client   influxdb2.Client
	queryAPI api.QueryAPI
	config   common.Config
}

func NewInfluxDBSource() sdk.Source {
	return sdk.SourceWithMiddleware(&InfluxDBSource{}, sdk.DefaultSourceMiddleware()...)
}

func (s *InfluxDBSource) Configure(ctx context.Context, cfg map[string]string) error {
	config, err := common.ParseConfig(cfg)
	if err != nil {
		return err
	}
	s.config = config
	return nil
}

func (s *InfluxDBSource) Open(ctx context.Context, position sdk.Position) error {
	client := influxdb2.NewClient(s.config.URL, s.config.Token)
	s.client = client
	s.queryAPI = client.QueryAPI(s.config.Org)
	return nil
}

func (s *InfluxDBSource) Read(ctx context.Context) (sdk.Record, error) {
	query := fmt.Sprintf(`from(bucket:"%s") |> range(start: -1h)`, s.config.Bucket)
	result, err := s.queryAPI.Query(ctx, query)
	if err != nil {
		return sdk.Record{}, err
	}
	defer result.Close()

	if result.Next() {
		if result.Err() != nil {
			return sdk.Record{}, result.Err()
		}

		values := result.Record().Values()
		keyData, err := json.Marshal(map[string]interface{}{"id": values["id"]})
		if err != nil {
			return sdk.Record{}, err
		}
		payloadData, err := json.Marshal(values)
		if err != nil {
			return sdk.Record{}, err
		}

		record := sdk.Util.Source.NewRecordCreate(
			sdk.Position(result.Record().Time().String()),
			sdk.Metadata{"source": "influxdb"},
			sdk.RawData(keyData),
			sdk.RawData(payloadData),
		)
		return record, nil
	}

	return sdk.Record{}, ctx.Err()
}

func (s *InfluxDBSource) Teardown(ctx context.Context) error {
	if s.client != nil {
		s.client.Close()
	}
	return nil
}
