package main

import (
	"github.com/GevorgGal/conduit-connector-influxdb/internal/destination"
	"github.com/GevorgGal/conduit-connector-influxdb/internal/source"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(sdk.Connector{
		NewSpecification: GetSpecification,
		NewSource:      source.NewInfluxDBSource,
		NewDestination: destination.NewInfluxDBDestination,
	})
}
