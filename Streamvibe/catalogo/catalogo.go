package catalogo

import (
	// Importación de paquetes necesarios para el funcionamiento del código.
	"Streamvibe/db"     // Importa el paquete para la conexión a la base de datos.
	"Streamvibe/search" // Importa un paquete personalizado para la búsqueda.
	"bufio"             // Para leer la entrada del usuario.
	"database/sql"      // Para interactuar con bases de datos SQL
	"encoding/json"     // Para decodificar datos JSON.
	"fmt"               // Para imprimir y formatear textos.
	"log"               // Para registrar errores y mensajes.
	"net/http"          // Para realizar solicitudes HTTP.
	"os"                // Para interactuar con el sistema operativo.
	"strconv"           // Para convertir strings a otros tipos de datos.
)

// Movie representa la estructura de una película.
type Movie struct {
	APImovieID int    `json:"id"`          // Identificador de la película en la API.
	Title      string `json:"title"`       // Título de la película.
	Overview   string `json:"overview"`    // Descripción general de la película
	PosterPath string `json:"poster_path"` // Ruta del póster de la película.
}

// ApiResponse representa la respuesta de la API.
type ApiResponse struct {
	Page         int     `json:"page"`          // Número de página actual de la respuesta.
	Results      []Movie `json:"results"`       // Lista de películas devueltas en la respuesta.
	TotalPages   int     `json:"total_pages"`   // Número total de páginas disponibles.
	TotalResults int     `json:"total_results"` // Número total de resultados disponibles.
}

// ShowCatalog muestra el catálogo de películas populares.
func ShowCatalog() {
	// URL de la API para obtener películas populares.
	url := "https://api.themoviedb.org/3/movie/popular?api_key=52abf4732494d35e674f7b2345c0486f&language=es-ES"

	// Realiza una solicitud HTTP GET a la URL.
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al realizar la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close() // Asegura que la respuesta se cierre al final de la función.

	var apiResponse ApiResponse
	// Decodifica la respuesta JSON en la estructura ApiResponse.
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return
	}

	// Imprime el título de cada película en el catálogo.
	for i, movie := range apiResponse.Results {
		fmt.Printf("%d: %s\n", i+1, movie.Title)
		fmt.Println("URL de la imagen:", "https://image.tmdb.org/t/p/w500"+movie.PosterPath)
	}

	// Lee la selección del usuario desde la entrada estándar.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese el número de la película que desea seleccionar: ")
	scanner.Scan()
	input := scanner.Text()
	selection, err := strconv.Atoi(input)
	// Verifica si la selección es válida.
	if err != nil || selection < 1 || selection > len(apiResponse.Results) {
		fmt.Println("Selección no válida")
		return
	}

	// Obtiene la película seleccionada de la lista de resultados.
	selectedMovie := apiResponse.Results[selection-1]

	// Verifica si la película ya está en la base de datos.
	var existingMovieID int
	err = db.DB.QueryRow("SELECT MovieID FROM Movies WHERE APImovieID = @p1", selectedMovie.APImovieID).Scan(&existingMovieID)

	// Maneja los errores de la consulta a la base de datos.
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error al verificar la película en la base de datos: %v\n", err)
		return
	}

	// Si la película no existe en la base de datos, la inserta.
	if existingMovieID == 0 {
		_, err = db.DB.Exec("INSERT INTO Movies (Title, Overview, PosterPath, APImovieID) VALUES (@p1, @p2, @p3, @p4)",
			selectedMovie.Title, selectedMovie.Overview, "https://image.tmdb.org/t/p/w500"+selectedMovie.PosterPath, selectedMovie.APImovieID)

		if err != nil {
			log.Printf("Error al insertar la película seleccionada en la base de datos: %v\n", err)
			return
		}
		fmt.Println("Película seleccionada guardada en la base de datos.")
	} else {
		// Si la película ya existe, informa al usuario.
		fmt.Println("La película seleccionada ya existe en la base de datos.")
	}

	// Busca la película en YouTube.
	search.SearchMovieYoutube(selectedMovie.Title + " película")
}
