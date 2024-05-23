package ip2geo

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestConfigDatastoreTypeMissingKeyShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewConfig().GetDatastoreType()
}

func TestConfigDatastoreTypeDefinedShouldReturnValue(t *testing.T) {
	expected := "ds_type"
	os.Setenv(datastoreType, expected)
	dsType := NewConfig().GetDatastoreType()
	if dsType != expected {
		t.Fatalf(fmt.Sprintf("expecting config %s to be %s, but %s was found", datastoreType, expected, dsType))
	}
}

func TestConfigDatastoreUrlMissingKeyShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewConfig().GetDatastoreUrl()
}

func TestConfigDatastoreUrlDefinedShouldReturnValue(t *testing.T) {
	expected := "url"
	os.Setenv(datastoreUrl, expected)
	url := NewConfig().GetDatastoreUrl()
	if url != expected {
		t.Fatalf(fmt.Sprintf("expecting config %s to be %s, but %s was found", datastoreUrl, expected, url))
	}
}

func TestConfigRateLimitRpsMissingKeyShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	NewConfig().GetRateLimitRps()
}

func TestConfigRateLimitRpsNonIntegerShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	expected := "hi"
	os.Setenv(rateLimitRps, expected)
	NewConfig().GetRateLimitRps()
}

func TestConfigRateLimitRpsDefinedShouldReturnValue(t *testing.T) {
	expected := 666
	os.Setenv(rateLimitRps, strconv.Itoa(expected))
	rps := NewConfig().GetRateLimitRps()
	if rps != expected {
		t.Fatalf(fmt.Sprintf("expecting config %s to be %v, but %v was found", datastoreUrl, expected, rps))
	}
}

func TestConfigPortMissingKeyShouldFallback(t *testing.T) {
	configuredPort := NewConfig().GetPort()
	if configuredPort != 8080 {
		t.Fatalf(fmt.Sprintf("expecting config %s to be %v, but %v was found", port, 8080, configuredPort))
	}
}

func TestConfigPortNonIntegerShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	expected := "hi"
	os.Setenv(port, expected)
	NewConfig().GetPort()
}

func TestConfigPortDefinedShouldReturnValue(t *testing.T) {
	expected := 8888
	os.Setenv(port, strconv.Itoa(expected))
	configuredPort := NewConfig().GetPort()
	if configuredPort != expected {
		t.Fatalf(fmt.Sprintf("expecting config %s to be %v, but %v was found", datastoreUrl, expected, configuredPort))
	}
}
