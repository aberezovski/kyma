package testkit

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	internalAPIUrlEnvName = "INTERNAL_API_URL"
	externalAPIUrlEnvName = "EXTERNAL_API_URL"
	gatewayUrlEnvName     = "GATEWAY_URL"
	skipVerifyEnvName     = "SKIP_SSL_VERIFY"
	centralEnvName        = "CENTRAL"
)

type TestConfig struct {
	InternalAPIUrl string
	ExternalAPIUrl string
	GatewayUrl     string
	SkipSslVerify  bool
	Central        bool
}

func ReadConfig() (TestConfig, error) {
	internalAPIUrl, found := os.LookupEnv(internalAPIUrlEnvName)
	if !found {
		return TestConfig{}, errors.New(fmt.Sprintf("failed to read %s environment variable", internalAPIUrlEnvName))
	}

	externalAPIUrl, found := os.LookupEnv(externalAPIUrlEnvName)
	if !found {
		return TestConfig{}, errors.New(fmt.Sprintf("failed to read %s environment variable", externalAPIUrlEnvName))
	}

	gatewayUrl, found := os.LookupEnv(gatewayUrlEnvName)
	if !found {
		return TestConfig{}, errors.New(fmt.Sprintf("failed to read %s environment variable", gatewayUrlEnvName))
	}

	skipVerify := false
	sv, found := os.LookupEnv(skipVerifyEnvName)
	if found {
		skipVerify, _ = strconv.ParseBool(sv)
	}

	central := false
	c, found := os.LookupEnv(centralEnvName)
	if found {
		central, _ = strconv.ParseBool(c)
	}

	config := TestConfig{
		InternalAPIUrl: internalAPIUrl,
		ExternalAPIUrl: externalAPIUrl,
		GatewayUrl:     gatewayUrl,
		SkipSslVerify:  skipVerify,
		Central:        central,
	}

	log.Printf("Read configuration: %+v", config)

	return config, nil
}
