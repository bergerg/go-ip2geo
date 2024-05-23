package main

import (
	"fmt"
	"net/http"

	"torq.io/ip2geo/ip2geo"
)

func main() {
	config := ip2geo.NewConfig()
	port := config.GetPort()
	service := ip2geo.CreateIpResolverService()

	findCountryRouter := http.NewServeMux()
	findCountryRouter.HandleFunc("/v1/find-country", service.FindCountry)

	rps := config.GetRateLimitRps()
	println(rps)
	rateLimitedService := ip2geo.WithRateLimiter(rps, findCountryRouter)

	fmt.Printf("server listening on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), rateLimitedService)
}
