package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	// 32 byte key for AES-256
	// 24 byte key for AES-192
	// 16 byte key for AES-128
	keys1 := [3][]byte{
		[]byte("asuperstrong32bitpasswordgohere!"),
		[]byte("asuperstrong24bitpasswor"),
		[]byte("asuperstrong16bi")}

	message := os.Args[1]

	for i := 0; i < len(keys1); i++ {
		cipherKey := keys1[i]
		//Encrypt the text:
		encrypted, err := encrypt(cipherKey, message)
		checkError(err)
		//Print the key and cipher text:
		fmt.Printf("\n\tCIPHER KEY: %s\n", string(cipherKey))
		fmt.Printf("\tENCRYPTED: %s\n", encrypted)
		//Decrypt the text:
		decrypted, err := decrypt(cipherKey, encrypted)
		checkError(err)
		//Print re-decrypted text:
		fmt.Printf("\tDECRYPTED: %s\n\n", decrypted)
	}

}

/*
 *	FUNCTION		: encrypt
 *	DESCRIPTION		:
 *		This function takes a string and a cipher key and uses AES to encrypt the message
 *
 *	PARAMETERS		:
 *		byte[] key	: Byte array containing the cipher key
 *		string message	: String containing the message to encrypt
 *
 *	RETURNS			:
 *		string encoded	: String containing the encoded user input
 *		error err	: Error message
 */
func encrypt(key []byte, message string) (encoded string, err error) {
	//Create byte array from the input string
	plainText := []byte(message)

	//Create a new AES cipher using the key
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv is the ciphertext up to the blocksize (16) - initialization vector
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	//Encrypt the data
	// CFB (short for cipher feedback) is an AES block cipher mode similar to
	// the CBC mode in the sense that for the encryption of a block
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//Return string encoded in base64
	return base64.RawStdEncoding.EncodeToString(cipherText), err
}

/*
 *	FUNCTION		: decrypt
 *	DESCRIPTION		:
 *		This function takes a string and a key and uses AES to decrypt the string into plain text
 *
 *	PARAMETERS		:
 *		byte[] key	: Byte array containing the cipher key
 *		string secure	: String containing an encrypted message
 *
 *	RETURNS			:
 *		string decoded	: String containing the decrypted equivalent of secure
 *		error err	: Error message
 */
func decrypt(key []byte, secure string) (decoded string, err error) {
	//Remove base64 encoding:
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

	//IF DecodeString failed, exit:
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
