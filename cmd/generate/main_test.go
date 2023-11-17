package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToCamelWithInitialisms(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{arg: "destination_id", want: "DestinationID"},
		{arg: "someDestinationId", want: "SomeDestinationID"},
		{arg: "fooApi", want: "FooAPI"},
		{arg: "fooAscii", want: "FooASCII"},
		{arg: "fooCpu", want: "FooCPU"},
		{arg: "fooCss", want: "FooCSS"},
		{arg: "fooDns", want: "FooDNS"},
		{arg: "fooEof", want: "FooEOF"},
		{arg: "fooGuid", want: "FooGUID"},
		{arg: "fooHtml", want: "FooHTML"},
		{arg: "fooHttp", want: "FooHTTP"},
		{arg: "fooHttps", want: "FooHTTPS"},
		{arg: "fooId", want: "FooID"},
		{arg: "fooIp", want: "FooIP"},
		{arg: "fooJson", want: "FooJSON"},
		{arg: "fooQps", want: "FooQPS"},
		{arg: "fooRam", want: "FooRAM"},
		{arg: "fooRpc", want: "FooRPC"},
		{arg: "fooSsh", want: "FooSSH"},
		{arg: "fooTcp", want: "FooTCP"},
		{arg: "fooTls", want: "FooTLS"},
		{arg: "fooTtl", want: "FooTTL"},
		{arg: "fooUdp", want: "FooUDP"},
		{arg: "fooUi", want: "FooUI"},
		{arg: "fooUid", want: "FooUID"},
		{arg: "fooUuid", want: "FooUUID"},
		{arg: "fooUri", want: "FooURI"},
		{arg: "fooUrl", want: "FooURL"},
		{arg: "fooXml", want: "FooXML"},
		{arg: "fooXss", want: "FooXSS"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			require.Equal(t, tt.want, toCamelWithInitialisms(tt.arg))
		})
	}
}
