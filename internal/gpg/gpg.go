package gpg

import (
	"fmt"
	"os"

	"github.com/hibare/GoCommon/v2/pkg/crypto/gpg"
	"golang.org/x/term"
)

func readPasswordFromPrompt() (string, error) {
	fmt.Print("Enter passphrase: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // Print a new line after the password input
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	return password, nil
}

func GPGDecryptor(privateKeyPath, publicKeyPath, inputGPGFile string) {
	gpg := gpg.GPG{}

	gpg.SetPrivateKey(privateKeyPath)

	gpg.SetPublicKey(publicKeyPath)

	pass, err := readPasswordFromPrompt()
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Println("Decrypting file...")
	dout, err := gpg.DecryptFile(inputGPGFile, pass)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	fmt.Printf("File decrypted at %s\n", dout)
}
