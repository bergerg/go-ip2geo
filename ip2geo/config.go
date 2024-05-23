package ip2geo

import (
	"fmt"
	"os"
	"strconv"
)

const (
	port          = "PORT"
	rateLimitRps  = "RATE_LIMIT_RPS"
	datastoreType = "DATASTORE_TYPE"
	datastoreUrl  = "DATASTORE_URL"
)

type Ip2GeoConfig interface {
	GetPort() int
	GetRateLimitRps() int
	GetDatastoreType() string
	GetDatastoreUrl() string
}

func NewConfig() Ip2GeoConfig {
	return &EnvironmentConfig{}
}

type EnvironmentConfig struct{}

func (ec *EnvironmentConfig) GetRateLimitRps() int {
	rpsAsStr := os.Getenv(rateLimitRps)
	if len(rpsAsStr) == 0 {
		panic(fmt.Sprintf("could not find required configuration %s", rateLimitRps))
	}
	rps, err := strconv.Atoi(rpsAsStr)
	if err != nil {
		panic(err)
	}
	return rps
}

func (ec *EnvironmentConfig) GetPort() int {
	portAsStr := os.Getenv(port)
	if len(portAsStr) == 0 {
		portAsStr = "8080"
	}
	rps, err := strconv.Atoi(portAsStr)
	if err != nil {
		panic(err)
	}
	return rps
}

func (ec *EnvironmentConfig) GetDatastoreType() string {
	dsType := os.Getenv(datastoreType)
	if len(dsType) == 0 {
		panic(fmt.Sprintf("could not find required configuration %s", datastoreType))
	}
	return dsType
}

func (ec *EnvironmentConfig) GetDatastoreUrl() string {
	dsUrl := os.Getenv(datastoreUrl)
	if len(dsUrl) == 0 {
		panic(fmt.Sprintf("could not find required configuration %s", datastoreUrl))
	}
	return dsUrl
}
