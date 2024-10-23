# otp-example

```sh
$ go run main.go otpsampleuser1
Server (Go)
Base32 encoded secret key: N52HA43BNVYGYZLVONSXEMI
Generated HOTP code: 985461
Serialized JSON data: {"guid":"otpsampleuser1","otp":"985461"}
Base64 encoded JSON data: eyJndWlkIjoib3Rwc2FtcGxldXNlcjEiLCJvdHAiOiI5ODU0NjEifQ==
Decoded JSON data: {"guid":"otpsampleuser1","otp":"985461"}
Deserialized data - GUID: otpsampleuser1, OTP: 985461
HOTP code is valid.

Clinet (Javascript)
GUID: otpsampleuser1
OTP: 985461
```
