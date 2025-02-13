package bootstrap

import (
	"os"
	"strconv"
)

const (
	_nameAppDefault = "reto-amaris-beer"

	_highDemandThresholdDefault = 100

	_streamNameDefault = "beer:high_demand"
)

func getApplicationName() string {
	appName := os.Getenv("SERVICE")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}

func getHighDemandThreshold() int {
	threshold := os.Getenv("HIGH_DEMAND_THRESHOLD")
	if threshold == "" {
		return _highDemandThresholdDefault
	}

	demand, err := strconv.Atoi(threshold)
	if err != nil {
		return _highDemandThresholdDefault
	}

	return demand
}

func getStreamName() string {
	streamName := os.Getenv("STREAM_NAME")
	if streamName == "" {
		return _streamNameDefault
	}

	return streamName
}
