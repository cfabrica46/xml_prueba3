package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Users es exportado...
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

//User es exportado...
type User struct {
	XMLName  xml.Name `xml:"user"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

func main() {

	archivo, err := os.OpenFile("databases.xml", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("archivo se abrio con exito")

	defer archivo.Close()

	data, err := ioutil.ReadAll(archivo)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	var users Users

	err = xml.Unmarshal(data, &users)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users.Users[0].Username)

}
