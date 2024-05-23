package ipresolver

import (
	"os"
	"testing"
)

func TestCreateCsvDatastoreWhenMissingFileShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	CsvIpResolver("path/to/nothing")
}

func TestCreateCsvDatastoreWithEmptyFileShouldReturnDefault(t *testing.T) {
	f, err := os.CreateTemp(".", "empty.csv")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer os.Remove(f.Name())

	resolver := CsvIpResolver(f.Name())
	countryAndCity, err := resolver.Resolve("localhost")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if countryAndCity.Country != "unknown" || countryAndCity.City != "unknown" {
		t.Fatalf("unexpected results when resolving from default resolver")
	}
}

func TestCreateCsvDatastoreWithDataShouldReturnThatData(t *testing.T) {
	f, err := os.CreateTemp(".", "empty.csv")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer os.Remove(f.Name())
	f.Write([]byte("1.2.3.4,Country,City"))
	resolver := CsvIpResolver(f.Name())
	countryAndCity, err := resolver.Resolve("1.2.3.4")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if countryAndCity.Country != "Country" || countryAndCity.City != "City" {
		t.Fatalf("unexpected results when resolving from csv resolver")
	}
}
