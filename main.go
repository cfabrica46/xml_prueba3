package main

import (
	"encoding/xml"
	"fmt"
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

	u1 := User{Username: "xtron", Password: "01234"}
	u2 := User{Username: "cfrabrica46", Password: "12345"}
	slice := []User{u1, u2}

	users := Users{Users: slice}

	fmt.Println(users)

	users2 := Users{Users: []User{User{Username: "xtron", Password: "01234"}, User{Username: "cfabrica46", Password: "12345"}}}

	data, err := xml.MarshalIndent(users2, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	archivo.Write(data)
}
