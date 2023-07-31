package gpg

import (
	"fmt"
	"os"

	"github.com/hibare/GoCommon/pkg/crypto/gpg"
	"golang.org/x/term"
)

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	// Read the file content into a buffer
	data := make([]byte, fileSize)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

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

	data, err := readFile(privateKeyPath)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	gpg.PrivateKey = string(data)

	data, err = readFile(publicKeyPath)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	gpg.PublicKey = string(data)

	pass, err := readPasswordFromPrompt()
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	gpg.Passphrase = pass

	fmt.Println("Decrypting file...")
	dout, err := gpg.DecryptFile(inputGPGFile)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	fmt.Printf("File decrypted at %s\n", dout)
}
