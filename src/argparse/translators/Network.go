package translators

import (
	"net"
	"net/netip"
)

//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Addr
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Port
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=AddrPort
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=Prefix
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=MACAddr
//go:generate ../../../bin/ifaceImplCheck -typeToCheck=CIDRTranslator

type (
	// Represents a cmd line argument that will be translated to a [netip.Addr]
	// type.
	//gen:ifaceImplCheck ifaceName Translator[netip.Addr]
	//gen:ifaceImplCheck imports net/netip
	//gen:ifaceImplCheck valOrPntr both
	Addr struct{}

	// Represents a cmd line argument that should be considered a port value.
	// Nothing more than an alias for BuiltinUint16
	//gen:ifaceImplCheck ifaceName Translator[uint16]
	//gen:ifaceImplCheck valOrPntr both
	Port = BuiltinUint16

	// Represents a cmd line argument that will be translated to a
	// [netip.AddrPort] type.
	//gen:ifaceImplCheck ifaceName Translator[netip.AddrPort]
	//gen:ifaceImplCheck imports net/netip
	//gen:ifaceImplCheck valOrPntr both
	AddrPort struct{}

	// Represents a cmd line argument that will be translated to a
	// [netip.Prefix] type.
	//gen:ifaceImplCheck ifaceName Translator[netip.Prefix]
	//gen:ifaceImplCheck imports net/netip
	//gen:ifaceImplCheck valOrPntr both
	Prefix struct{}

	// Represents a cmd line argument that will be translated to a
	// [net.MACAddr] type.
	//gen:ifaceImplCheck ifaceName Translator[net.HardwareAddr]
	//gen:ifaceImplCheck imports net
	//gen:ifaceImplCheck valOrPntr both
	MACAddr struct{}

	CIDRVals struct {
		net.IP
		net.IPNet
	}
	// Represents a cmd line argument that will be considered a network mask
	// following the CIDR format. The string will be translated into a [net.IP]
	// and [net.IPNet] values.
	//gen:ifaceImplCheck ifaceName Translator[CIDRVals]
	//gen:ifaceImplCheck valOrPntr both
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
