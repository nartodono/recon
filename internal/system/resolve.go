package system

import "net"

func ValidateResolvable(target string) error {
	if net.ParseIP(target) != nil {
		return nil
	}
	if _, err := net.LookupHost(target); err != nil {
		return ResolveError{Target: target}
	}
	return nil
}
