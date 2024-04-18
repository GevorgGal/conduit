# Conduit Connector InfluxDB

## Overview

Conduit Connector InfluxDB is a connector for Conduit that enables interaction with InfluxDB, allowing data extraction from and data writing to InfluxDB. It supports full lifecycle operations including create, read, update, and delete data points.

## Installation

To use the Conduit Connector InfluxDB, you can either build it from source or download pre-built binaries from the releases page.

### Build from Source

1. Clone the repository:

    ```bash
    git clone https://github.com/GevorgGal/conduit-connector-influxdb.git
    ```

2. Navigate to the project directory:

    ```bash
    cd conduit-connector-influxdb
    ```

3. Build the connector:

    ```bash
    make build
    ```

### Download Pre-built Binaries

Pre-built binaries can be downloaded from the releases page on GitHub.

## Usage

Once you have built or downloaded the Conduit Connector InfluxDB, you can use it by running the compiled binary and providing the necessary configuration parameters.

Example usage:

```bash
./conduit-connector-influxdb --config=config.yml
