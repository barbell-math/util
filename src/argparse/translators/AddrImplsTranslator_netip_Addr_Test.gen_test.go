package translators

// Code generated by ../../../bin/ifaceImplCheck - DO NOT EDIT.
import (
	"net/netip"
	"testing"
)

func TestAddrValueImplementsTranslator_netip_Addr_(t *testing.T) {
	var typeThing Addr
	var iFaceThing Translator[netip.Addr] = typeThing
	_ = iFaceThing
}

func TestAddrPntrImplementsTranslator_netip_Addr_(t *testing.T) {
	var typeThing Addr
	var iFaceThing Translator[netip.Addr] = &typeThing
	_ = iFaceThing
}
