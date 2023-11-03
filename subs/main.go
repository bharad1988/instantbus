package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"time"

	"github.com/bharad1988/instantbus/sbclient"
)

func main() {
	// var wait chan bool
	serviceBusAddress := "localhost:5001"
	ss := sbclient.SubscriptionService{}
	ss2 := sbclient.SubscriptionService{}
	ss2.StartPublisher(serviceBusAddress)
	ss.StartSubscriptionService(serviceBusAddress)
	ss.SubscribeTopic("world")
	ss.SubscribeTopic("earth")
	for i := 0; i < 100; i++ {
		m := fmt.Sprintf("test %d", i)
		testMsg := []byte(m)
		ss2.SendMessage("world", testMsg)
		ss2.SendMessage("earth", testMsg)
	}

	plainText := "Hello, World!"
	fmt.Println("This is an original:", plainText)
	encrypted, err := GetAESEncrypted(plainText)
	if err != nil {
		fmt.Println("Error during encryption", err)
	}
	ss2.SendMessage("earth", encrypted)
	fmt.Println("This is an encrypted:", encrypted)
	decrypted, err := GetAESDecrypted(encrypted)

	if err != nil {
		fmt.Println("Error during decryption", err)
	}
	fmt.Println("This is a decrypted:", string(decrypted))
	time.Sleep(5 * time.Second)
	fmt.Println("len of messages ", len(ss.GetAllMessages("earth")))
	/*
		for i := 0; i < len(ss.GetAllMessages("earth")); i++ {
			fmt.Printf("All message - %s", ss.GetAllMessages("earth")[i])
		}
	*/
	// for {
	newMessages := ss.GetAllUnreadMessages("earth")
	for _, m := range newMessages {
		fmt.Printf("topic : earth : message %s \n", m)
	}
	time.Sleep(time.Second * 1)
	// }
	ss2.SendMessage("earth", encrypted)
	ss.GetAllUnreadMessages("earth")
	for _, m := range newMessages {
		d, _ := GetAESDecrypted(m)
		fmt.Printf("topic : earth : message %s \n", d)
	}
	time.Sleep(time.Second * 1)

	// <-wait
}

// GetAESDecrypted decrypts given text in AES 256 CBC
func GetAESDecrypted(ciphertext []byte) ([]byte, error) {
	key := "my32digitkey12345678901234567890"
	iv := "my16digitIvKey12"
	/* 	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	   	if err != nil {
	   		return nil, err
	   	} */
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

// GetAESEncrypted encrypts given text in AES 256 CBC
func GetAESEncrypted(plaintext string) ([]byte, error) {
	key := "my32digitkey12345678901234567890"
	iv := "my16digitIvKey12"

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)
	// str := base64.StdEncoding.EncodeToString(ciphertext)

	return ciphertext, nil
}

func testIt() {
	plainText := "Hello, World!"
	fmt.Println("This is an original:", plainText)

	encrypted, err := GetAESEncrypted(plainText)

	if err != nil {
		fmt.Println("Error during encryption", err)
	}

	fmt.Println("This is an encrypted:", encrypted)

	decrypted, err := GetAESDecrypted(encrypted)

	if err != nil {
		fmt.Println("Error during decryption", err)
	}
	fmt.Println("This is a decrypted:", string(decrypted))
}
