package ipresolver

type CountryCity struct {
	Country, City string
}

type IpToGeoResolver interface {
	Resolve(ip string) (CountryCity, error)
}
