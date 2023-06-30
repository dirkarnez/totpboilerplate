package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/pquerna/otp/totp"
)

var (
	secret string
)

func main() {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _ := key.Image(200, 200)
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// Now Validate that the user's successfully added the passcode.
	passcode := promptForPasscode()
	valid := totp.Validate(passcode, key.Secret())

	fmt.Scanln()

	if valid {
		// User successfully used their TOTP, save it to your backend!
		storeSecret("alice@example.com", key.Secret())
	}

	passcode = promptForPasscode()
	//secret := getSecret("alice@example.com")

	if totp.Validate(passcode, secret) {
		fmt.Println("Success! continue login process.")
	}
}

func storeSecret(email, secretIn string) {
	log.Println("alice@example.com", secretIn)
	secret = secretIn
}

func promptForPasscode() string {
	fmt.Println("promptForPasscode")
	var input string
	fmt.Scanln(&input)
	return input
}
