package generos

import (
	"Streamvibe/db"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Genre define la estructura para almacenar información de géneros de películas.
type Genre struct {
	ID   int    `json:"id"`   // Campo 'ID' para almacenar el identificador del género, mapeado desde 'id' en JSON.
	Name string `json:"name"` // Campo 'Name' para almacenar el nombre del género, mapeado desde 'name'.
}

// GenreList define la estructura para una lista de géneros.
type GenreList struct {
	Genres []Genre `json:"genres"` // Campo 'Genres' es un slice que almacena múltiples géneros.
}

// ShowGenres muestra los géneros disponibles y permite al usuario seleccionar uno.
func ShowGenres() {
	apiKey := "52abf4732494d35e674f7b2345c0486f"                                                       // Clave API para TMDB.
	url := fmt.Sprintf("https://api.themoviedb.org/3/genre/movie/list?api_key=%s&language=es", apiKey) // URL para obtener la lista de géneros de películas.

	req, err := http.NewRequest("GET", url, nil) // Crear una nueva solicitud HTTP GET.
	if err != nil {
		log.Fatal(err) // Manejar errores al crear la solicitud.
	}

	req.Header.Add("accept", "application/json") // Añadir cabecera para aceptar JSON.

	res, err := http.DefaultClient.Do(req) // Enviar la solicitud y obtener la respuesta.
	if err != nil {
		log.Fatal(err) // Manejar errores al realizar la solicitud.
	}
	defer res.Body.Close() // Asegurarse de cerrar el cuerpo de la respuesta.

	body, err := io.ReadAll(res.Body) // Leer todo el cuerpo de la respuesta.
	if err != nil {
		log.Fatal(err) // Manejar errores al leer la respuesta.
	}

	var genreList GenreList                // Crear una variable para almacenar la lista de géneros.
	err = json.Unmarshal(body, &genreList) // Deserializar el cuerpo de la respuesta a la estructura GenreList.
	if err != nil {
		log.Fatal(err) // Manejar errores en la deserialización.
	}

	fmt.Println("Lista de Géneros:") // Imprimir el encabezado de la lista de géneros.
	for i, genre := range genreList.Genres {
		fmt.Printf("%d: %s\n", i+1, genre.Name) // Imprimir cada género con su índice.
	}

	scanner := bufio.NewScanner(os.Stdin) // Crear un scanner para leer la entrada del usuario.
	fmt.Print("Ingrese el número del género que desea seleccionar: ")
	scanner.Scan()                        // Leer la entrada del usuario.
	input := scanner.Text()               // Almacenar la entrada del usuario.
	selection, err := strconv.Atoi(input) // Convertir la entrada a un número entero.
	if err != nil || selection < 1 || selection > len(genreList.Genres) {
		fmt.Println("Selección no válida") // Manejar selecciones inválidas.
		return
	}

	selectedGenre := genreList.Genres[selection-1] // Obtener el género seleccionado.

	// Comprobar si el género ya existe en la base de datos.
	var exists int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM Genres WHERE GenreID = @p1", selectedGenre.ID).Scan(&exists)
	if err != nil {
		log.Fatal(err) // Manejar errores en la consulta a la base de datos.
	}

	if exists == 0 {
		// Insertar el género seleccionado en la base de datos si no existe.
		_, err = db.DB.Exec("INSERT INTO Genres (GenreID, Name) VALUES (@p1, @p2)", selectedGenre.ID, selectedGenre.Name)
		if err != nil {
			log.Printf("Error al insertar el género seleccionado en la base de datos: %v\n", err) // Manejar errores en la inserción.
		}
	} else {
		fmt.Println("El género ya existe en la base de datos.") // Informar al usuario si el género ya existe.
	}
}
