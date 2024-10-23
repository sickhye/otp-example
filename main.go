package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"

	"github.com/pquerna/otp/hotp"
)

func isValidGameUID(gameUID string) bool {
	if len(gameUID) < 12 {
		return false
	}
	for _, char := range gameUID {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func serializeToJSON(gameUID, otp string) (string, error) {
	data := map[string]string{
		"guid": gameUID,
		"otp":  otp,
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func deserializeFromJSON(jsonData string) (string, string, error) {
	var data map[string]string
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", "", err
	}
	return data["guid"], data["otp"], nil
}

func main() {
	fmt.Println("Server (Go)")

	if len(os.Args) < 2 {
		log.Fatalf("Error: gameUID must be provided as a command-line argument.")
	}

	gameUID := os.Args[1]

	if !isValidGameUID(gameUID) {
		log.Fatalf("Error: Invalid gameUID. It must be at least 12 characters long and contain only letters and digits.")
	}

	// Base32 encoding (padding processed according to RFC 4648)
	base32Encoded := base32.StdEncoding.EncodeToString([]byte(gameUID))

	// Remove padding if necessary
	secret := strings.TrimRight(base32Encoded, "=")

	fmt.Printf("Base32 encoded secret key: %s\n", secret)

	// Set the counter value (increase counter every time the user authenticates)
	counter := uint64(time.Now().Unix())

	// Generate HOTP code
	otp, err := hotp.GenerateCode(secret, counter)
	if err != nil {
		log.Fatalf("Error generating HOTP: %v", err)
	}

	fmt.Printf("Generated HOTP code: %s\n", otp)

	// Serialize to JSON format
	jsonData, err := serializeToJSON(gameUID, otp)
	if err != nil {
		log.Fatalf("JSON serialization error: %v", err)
	}
	fmt.Printf("Serialized JSON data: %s\n", jsonData)

	// Encode as Base64
	base64Data := base64.StdEncoding.EncodeToString([]byte(jsonData))
	fmt.Printf("Base64 encoded JSON data: %s\n", base64Data)

	// Decode Base64 data
	decodedJsonData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		log.Fatalf("Base64 decoding error: %v", err)
	}

	fmt.Printf("Decoded JSON data: %s\n", decodedJsonData)

	// Deserialize the decoded JSON data
	deserializedGuid, deserializedOtp, err := deserializeFromJSON(string(decodedJsonData))
	if err != nil {
		log.Fatalf("JSON deserialization error: %v", err)
	}
	fmt.Printf("Deserialized data - GUID: %s, OTP: %s\n", deserializedGuid, deserializedOtp)

	// Validate HOTP code
	valid := hotp.Validate(deserializedOtp, counter, secret)
	if valid {
		fmt.Println("HOTP code is valid.")
	} else {
		fmt.Println("HOTP code is invalid.")
	}
	fmt.Print("\n")

	// Call the JavaScript code with the Base64 encoded data
	cmd := exec.Command("node", "./client/decode.js", base64Data)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing JavaScript code: %v", err)
	}

	// Output from JavaScript code
	fmt.Println(string(output))
}
