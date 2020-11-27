package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type user struct {
	Username string
	Password string
}

var usuario, password string

func main() {

	log.SetFlags(log.Llongfile)

	var eleccion1, eleccion2 string

	databases := []user{}

	load(&databases)

	for {
		fmt.Println("¿Desea Registrase o Iniciar Secion? [R/I]")
		fmt.Scan(&eleccion1)

		e1 := strings.ToLower(eleccion1)

		if e1 == "r" {

			registrar(&databases)

			fmt.Println("¿Desea Iniciar Secion? [S/N]")
			fmt.Scan(&eleccion2)

			e2 := strings.ToLower(eleccion2)

			if e2 == "s" {
				ingresar(databases)
				break
			}
			if e2 == "n" {
				break
			}
		}
		if e1 == "i" {

			ingresar(databases)
			break
		}

	}

}

//en caso de un proyecto real el archivo databases.json seria ignorado por git
func load(databases *[]user) {

	archivoxml, err := os.OpenFile("databases.xml", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivoxml.Close()

	*databases = []user{}

	buf := make([]byte, 4)

	contenido := []byte{}

	for {

		_, err := archivoxml.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		contenido = append(contenido, buf...)

	}
	fmt.Println(string(contenido))
	err = xml.Unmarshal(contenido, &*databases)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*databases)
}

func registrar(databases *[]user) {

	fmt.Println("Nombre de Usuario")
	fmt.Scan(&usuario)
	fmt.Println("Contraseña")
	fmt.Scan(&password)

	nuevousuario := user{Username: usuario, Password: password}
	fmt.Println(*databases)
	*databases = append(*databases, nuevousuario)

	data, err := xml.MarshalIndent(databases, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	nuevadata := bytes.NewBuffer(data)

	buf := make([]byte, 4)

	archivoxml, err := os.OpenFile("databases.xml", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivoxml.Close()

	err = archivoxml.Truncate(0)

	if err != nil {
		log.Fatal(err)
	}

	for {

		n, err := nuevadata.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		_, err = archivoxml.Write(buf[:n])

		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Print("Nuevo usuario\t")
	fmt.Printf("Username: %v || Password: %v\n", nuevousuario.Username, nuevousuario.Password)

}

func ingresar(databases []user) {

	var ingreso bool

	fmt.Println("Ingrese su nombre de usuario")
	fmt.Scan(&usuario)
	fmt.Println("Ingrese su nombre de contraseña")
	fmt.Scan(&password)

	for i := range databases {

		if usuario == databases[i].Username && password == databases[i].Password {

			ingreso = true

			fmt.Println("Bienvenido")
			fmt.Println(databases[i].Username)

		}
	}

	if ingreso == false {
		fmt.Println("Nombre y/o Contraseña equivocadas")
	}
}
