package main

import "testing"

func TestToCamelWithInitialisms(t *testing.T) {
	tests := []struct {
		test string
		arg  string
		want string
	}{
		{test: "1", arg: "destination_id", want: "DestinationID"},
		{test: "2", arg: "someDestinationID", want: "SomeDestinationID"},
		{test: "3", arg: "fooApi", want: "FooAPI"},
		{test: "4", arg: "fooAscii", want: "FooASCII"},
		{test: "5", arg: "fooCpu", want: "FooCPU"},
		{test: "6", arg: "fooCss", want: "FooCSS"},
		{test: "7", arg: "fooDns", want: "FooDNS"},
		{test: "8", arg: "fooEof", want: "FooEOF"},
		{test: "9", arg: "fooGuid", want: "FooGUID"},
		{test: "10", arg: "fooHtml", want: "FooHTML"},
		{test: "11", arg: "fooHttp", want: "FooHTTP"},
		{test: "12", arg: "fooHttps", want: "FooHTTPS"},
		{test: "13", arg: "fooId", want: "FooID"},
		{test: "14", arg: "fooIp", want: "FooIP"},
		{test: "15", arg: "fooJson", want: "FooJSON"},
		{test: "16", arg: "fooQps", want: "FooQPS"},
		{test: "17", arg: "fooRam", want: "FooRAM"},
		{test: "18", arg: "fooRpc", want: "FooRPC"},
		{test: "19", arg: "fooSsh", want: "FooSSH"},
		{test: "20", arg: "fooTcp", want: "FooTCP"},
		{test: "21", arg: "fooTls", want: "FooTLS"},
		{test: "22", arg: "fooTtl", want: "FooTTL"},
		{test: "23", arg: "fooUdp", want: "FooUDP"},
		{test: "24", arg: "fooUi", want: "FooUI"},
		{test: "25", arg: "fooUid", want: "FooUID"},
		{test: "26", arg: "fooUuid", want: "FooUUID"},
		{test: "27", arg: "fooUri", want: "FooURI"},
		{test: "28", arg: "fooUrl", want: "FooURL"},
		{test: "29", arg: "fooUtf8", want: "FooUTF8"},
		{test: "30", arg: "fooXml", want: "FooXML"},
		{test: "31", arg: "fooXss", want: "FooXSS"},
	}
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			if got := toCamelWithInitialisms(tt.arg); got != tt.want {
				t.Errorf("toCamelWithInitialisms() = %v, want %v", got, tt.want)
			}
		})
	}
}
