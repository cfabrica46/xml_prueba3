package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type user struct {
	Username string
	Password string
}

func main() {

	//el contenido de databases.xml esta creado por un marshal y de parametro un slice de user

	archivo, err := os.OpenFile("databases.xml", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	contenido, err := ioutil.ReadAll(archivo)

	if err != nil {
		log.Fatal(err)
	}

	users := []user{}

	err = xml.Unmarshal(contenido, &users)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
