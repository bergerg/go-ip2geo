package ipresolver

import (
	"bufio"
	"os"
	"strings"
)

type CsvFileDatastore struct {
	Ip2Geo map[string]CountryCity
}

func (csfdb *CsvFileDatastore) Resolve(ip string) (CountryCity, error) {
	val, ok := csfdb.Ip2Geo[ip]
	if ok {
		return val, nil
	} else {
		return val, &IpMissingError{ip}
	}
}

func CsvIpResolver(filePath string) IpToGeoResolver {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	data := make(map[string]CountryCity)
	data["localhost"] = CountryCity{"unknown", "unknown"} // add at least one key with default values

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), ",")
		// TODO: data validations
		data[splitLine[0]] = CountryCity{splitLine[1], splitLine[2]}
	}

	return &CsvFileDatastore{data}
}
