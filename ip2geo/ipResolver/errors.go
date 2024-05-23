package ipresolver

import "fmt"

type IpMissingError struct {
	Ip string
}

func (e *IpMissingError) Error() string {
	return fmt.Sprintf("IP %s is missing from the database", e.Ip)
}
