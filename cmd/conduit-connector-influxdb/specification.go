package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func GetSpecification() sdk.Specification {
	return sdk.Specification{
		Name:        "conduit-connector-influxdb",
		Summary:     "InfluxDB connector for Conduit that supports both source and destination operations.",
		Description: "This connector allows Conduit to interact with InfluxDB, enabling data extraction from and data writing to InfluxDB. It supports full lifecycle operations including create, read, update, and delete data points.",
		Version:     "v1.0.0",
		Author:      "Gevorg Galstyan",
	}
}
