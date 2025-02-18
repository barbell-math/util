package translators

import (
	"net"
	"net/netip"
)

type (
	// Represents a cmd line argument that will be translated to a [netip.Addr]
	// type.
	Addr struct{}

	// Represents a cmd line argument that should be considered a port value.
	// Nothing more than an alias for BuiltinUint16
	Port = BuiltinUint16

	// Represents a cmd line argument that will be translated to a
	// [netip.AddrPort] type.
	AddrPort struct{}

	// Represents a cmd line argument that will be translated to a
	// [netip.Prefix] type.
	Prefix struct{}

	// Represents a cmd line argument that will be translated to a
	// [net.MACAddr] type.
	MACAddr struct{}

	CIDRVals struct {
		net.IP
		net.IPNet
	}
	// Represents a cmd line argument that will be considered a network mask
	// following the CIDR format. The string will be translated into a [net.IP]
	// and [net.IPNet] values.
	CIDRTranslator struct{}
)

func (_ Addr) Translate(arg string) (netip.Addr, error) {
	return netip.ParseAddr(arg)
}
func (_ Addr) Reset() {
	// intentional noop - Addr has no state that needs to be reset
}

func (_ AddrPort) Translate(arg string) (netip.AddrPort, error) {
	return netip.ParseAddrPort(arg)
}
func (_ AddrPort) Reset() {
	// intentional noop - AddrPort has no state that needs to be reset
}

func (_ Prefix) Translate(arg string) (netip.Prefix, error) {
	return netip.ParsePrefix(arg)
}
func (_ Prefix) Reset() {
	// intentional noop - Prefix has no state that needs to be reset
}

func (_ MACAddr) Translate(arg string) (net.HardwareAddr, error) {
	return net.ParseMAC(arg)
}
func (_ MACAddr) Reset() {
	// intentional noop - MACAddr has no state that needs to be reset
}

func (_ CIDRTranslator) Translate(arg string) (CIDRVals, error) {
	ip, netip, err := net.ParseCIDR(arg)
	rv := CIDRVals{IP: ip}
	if netip != nil {
		rv.IPNet = *netip
	}
	return rv, err
}
func (_ CIDRTranslator) Reset() {
	// intentional noop - CIDRTranslator has no state that needs to be reset
}
