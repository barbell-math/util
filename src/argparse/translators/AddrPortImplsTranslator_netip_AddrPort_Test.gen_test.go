package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"net/netip"
	"testing"
)

func TestAddrPortValueImplementsTranslator_netip_AddrPort_(t *testing.T) {
	var typeThing AddrPort
	var iFaceThing Translator[netip.AddrPort] = typeThing
	_ = iFaceThing
}

func TestAddrPortPntrImplementsTranslator_netip_AddrPort_(t *testing.T) {
	var typeThing AddrPort
	var iFaceThing Translator[netip.AddrPort] = &typeThing
	_ = iFaceThing
}
