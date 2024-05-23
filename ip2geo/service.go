package ip2geo

import (
	"encoding/json"
	"net/http"
	"strings"

	ipresolver "torq.io/ip2geo/ip2geo/ipResolver"
)

type IpResolverService struct {
	resolver ipresolver.IpToGeoResolver
}

func CreateIpResolverService() *IpResolverService {
	config := NewConfig()
	switch ipReolverType := strings.ToLower(config.GetDatastoreType()); ipReolverType {
	case "csv":
		filepath := config.GetDatastoreUrl()
		return &IpResolverService{ipresolver.CsvIpResolver(filepath)}
	default:
		panic("could not find any configured data store")
	}
}

func (service *IpResolverService) FindCountry(w http.ResponseWriter, req *http.Request) {

	// TODO: input validation
	ip := req.URL.Query().Get("ip")

	println("got a request for %s", ip)

	countryAndCity, err := service.resolver.Resolve(ip)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		e, _ := json.Marshal(ErrorResponseBody{err.Error()})
		w.Write(e)
		return
	}

	resp, err := json.Marshal(&countryAndCity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e, _ := json.Marshal(ErrorResponseBody{err.Error()})
		w.Write(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
