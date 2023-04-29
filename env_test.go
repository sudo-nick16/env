package env

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	testcases := []struct {
		key   string
		value string
	}{
		{
			"name", "nikit",
		},
		{
			"name", "nick",
		},
		{
			"name", "sudonick",
		},
	}
	for i, v := range testcases {
		err := os.Setenv(v.key, v.value)
		if err != nil {
			t.Error("couldn't set env.")
		}
		name := GetEnv(v.key, "")
		if name != v.value {
			t.Errorf("failed test %d", i)
		}
	}
}

func TestGetEnvFromBase64(t *testing.T) {
	testcases := []struct {
		key      string
		value    string
		expected string
	}{
		{
			"key", `LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUJWUUlCQURBTkJna3Foa2lHOXcwQkFRRUZB
QVNDQVQ4d2dnRTdBZ0VBQWtFQTE2dnZPTG5NODdYNzlkOWwKZFpUK25DdXlTQzJnSDlVR21pMWU4
TDUvcmZtY3BrUFNhYVp0ckw0SDZtUVgvOGptRDVhYWNBaWUxOStub3ZsdAo2VXFwOFFJREFRQUJB
a0Fvam1oNTNsYXdMYlNuMDg2Y2dkd3ZPZmhZb2x2T0lKMFJPUjcxWHAxYjlaQWJPblJvCmk2VWVq
QkhyME4vNmlqcTB6K3lQWFFWb0lRTS9nMFBHdnFoUkFpRUE4YkUvaHd1ZUFPcDlScjBLalp6eWpm
U2YKd2JZc3FHblVLQmR4V0lRUERQTUNJUURrY0ZvaVFqSlp4TTNVM3kvcXN3NW5hZTJZV0p4M1Iw
Wld2S3NwaTJkVwppd0loQU84RnVPcm9UVWxnbmFaUW5GZ3lxQnBGbTV6cWVqM1A4M1gyd1N5bDVY
Q0xBaUE4bzhwNWI2TVlDMU1zClgvYkRVYmJRSXVGc1lKRmdaRzQ2bGlqRmhYandFd0loQUxpV3hE
cDNFYm1ERUh6a3hMOStaeDlmZ0I1NUlhZ0YKamFoaEhwTktaaXZFCi0tLS0tRU5EIFBSSVZBVEUg
S0VZLS0tLS0K`, `-----BEGIN PRIVATE KEY-----
MIIBVQIBADANBgkqhkiG9w0BAQEFAASCAT8wggE7AgEAAkEA16vvOLnM87X79d9l
dZT+nCuySC2gH9UGmi1e8L5/rfmcpkPSaaZtrL4H6mQX/8jmD5aacAie19+novlt
6Uqp8QIDAQABAkAojmh53lawLbSn086cgdwvOfhYolvOIJ0ROR71Xp1b9ZAbOnRo
i6UejBHr0N/6ijq0z+yPXQVoIQM/g0PGvqhRAiEA8bE/hwueAOp9Rr0KjZzyjfSf
wbYsqGnUKBdxWIQPDPMCIQDkcFoiQjJZxM3U3y/qsw5nae2YWJx3R0ZWvKspi2dW
iwIhAO8FuOroTUlgnaZQnFgyqBpFm5zqej3P83X2wSyl5XCLAiA8o8p5b6MYC1Ms
X/bDUbbQIuFsYJFgZG46lijFhXjwEwIhALiWxDp3EbmDEHzkxL9+Zx9fgB55IagF
jahhHpNKZivE
-----END PRIVATE KEY-----`,
		},
	}
	for i, v := range testcases {
		err := os.Setenv(v.key, v.value)
		if err != nil {
			t.Error("couldn't set env.")
		}
		val := GetEnvFromBase64(v.key, "")
		if val != v.expected {
			t.Errorf("failed test %d\nexpected: %v\ngot: %v", i, v.expected, val)
		}
	}
}

func TestGetEnvAsInt(t *testing.T) {
	testcases := []struct {
		key      string
		value    string
		expected int
	}{
		{
			"0", "0", 0,
		},
		{
			"1", "1", 1,
		},
		{
			"2", "2", 2,
		},
	}
	for i, v := range testcases {
		err := os.Setenv(v.key, v.value)
		if err != nil {
			t.Error("couldn't set env.")
		}
		integer := GetEnvAsInt(v.key, -1)
		if integer != v.expected {
			t.Errorf("failed test %d\nexpected: %v\tgot: %v", i, v.expected, integer)
		}
	}
}

func TestGetEnvAsBool(t *testing.T) {
	testcases := []struct {
		key      string
		value    string
		expected bool
	}{
		{
			"true", "false", false,
		},
		{
			"false", "true", true,
		},
	}
	for i, v := range testcases {
		err := os.Setenv(v.key, v.value)
		if err != nil {
			t.Error("couldn't set env.")
		}
		name := GetEnvAsBool(v.key, true)
		if name != v.expected {
			t.Errorf("failed test %d", i)
		}
	}
}

func TestGetEnvAsSlice(t *testing.T) {
	testcases := []struct {
		key      string
		value    string
		delim    string
		expected []string
	}{
		{
			"names", "nikit,nick,sudonick", ",", []string{"nikit", "nick", "sudonick"},
		},
		{
			"topics", "none;nil;null", ";", []string{"none", "nil", "null"},
		},
	}
	for i, v := range testcases {
		err := os.Setenv(v.key, v.value)
		if err != nil {
			t.Error("couldn't set env.")
		}
		res := GetEnvAsSlice(v.key, []string{}, v.delim)
		if len(res) != len(v.expected) {
			t.Errorf("failed test %d: length doesn't match.", i)
		}
		for i := range res {
			if res[i] != v.expected[i] {
				t.Errorf("failed test %d\ngot: %v\texpected: %v", i, res, v.expected)
				break
			}
		}
	}
}
