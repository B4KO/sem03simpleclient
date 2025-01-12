package main

import (
	"github.com/B4KO/is105sem03/mycrypt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "173.17.0.5:8080")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	encryptedMsg := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)

	_, err = conn.Write([]byte(string(encryptedMsg)))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := string(buf[:n])

	unencryptedResponse := mycrypt.Krypter([]rune(response), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
	log.Printf("reply from proxy: %s", string(unencryptedResponse))
}
