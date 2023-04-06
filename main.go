package main

import (
	"log"
	"os"

	qrgenerator "github.com/alexanderosadc/qr-generator/pkg/qr-generator"
)

func main() {
	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatalf("file has not been generated: %s", err)
	}
	defer file.Close()

	err = qrgenerator.CreateQRCode(file, "45-35-26")
	if err != nil {
		log.Fatal(err)
	}
	// qrCode, err = qrgenerator.CreateQRCode("github.com")
}
