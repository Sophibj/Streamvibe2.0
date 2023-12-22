package main

import (
	"Streamvibe/auth"
	"Streamvibe/db"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Estructuras para deserializar la respuesta JSON
type Genre struct {
	ID   int    `json:"id"`   // Identificador único del género en la respuesta JSON.
	Name string `json:"name"` // Nombre del género en la respuesta JSON.
}

type GenreList struct {
	Genres []Genre `json:"genres"` // Lista de géneros en la respuesta JSON.
}

func main() {
	// Inicialización de la base de datos
	db.Init()
	defer db.Close() //se usa para cerrar la conexión a la base de datos

	var userID int
	for {
		fmt.Println("Selecciona la Opción deseada")
		fmt.Println("1. Crear Nuevo Usuario")
		fmt.Println("2. Iniciar Sesión")
		fmt.Println("3. Salir")
		fmt.Print("Ingrese su elección: ")

		// Crear un scanner para leer la entrada del usuario
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("Opción no válida:", err)
			continue // Volver al inicio del bucle si la opción no es válida
		}

		switch choice {
		case 1:
			// Crear un nuevo usuario llamando a la función NewUser() del paquete "auth"
			auth.NewUser()
		case 2:
			// Iniciar sesión y mostrar el menú del usuario si el inicio de sesión es exitoso
			userID = auth.Login()
			if userID != 0 {
				auth.UserMenu()
			}
		case 3:
			// Salir de la aplicación
			fmt.Println("Gracias por usar StreamVibe. ¡Hasta luego!")
			return // Salir del programa
		default:
			// Opción no válida
			fmt.Println("Opción no válida")
		}
	}
}
