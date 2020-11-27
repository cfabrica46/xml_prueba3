package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type user struct {
	Username string
	Password string
}

func main() {

	databases := []user{
		user{
			Username: "xtron",
			Password: "01234",
		},
		user{
			Username: "cfabrica46",
			Password: "12345",
		},
	}

	fmt.Println(databases)

	data, err := xml.MarshalIndent(databases, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	archivo, err := os.OpenFile("databases.xml", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	_, err = archivo.Write(data)

	if err != nil {
		log.Fatal(err)
	}

}
