package search

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Movie define la estructura para almacenar información de películas.
type Movie struct {
	Title       string `json:"title"`        // Campo para el título de la película, mapeado del JSON 'title'.
	Overview    string `json:"overview"`     // Campo para la descripción general, mapeado de 'overview'.
	ReleaseDate string `json:"release_date"` // Campo para la fecha de lanzamiento, mapeado de 'release_date'.
}

// SearchMovieYoutube realiza una búsqueda en YouTube basada en un término de búsqueda proporcionado.
func SearchMovieYoutube(searchTerm string) {
	// Crear un contexto "vacío" que se usará en la solicitud a la API de YouTube.
	ctx := context.Background()
	// Definir la clave API para autenticar las solicitudes a la API de YouTube.
	apiKey := "[TU_API_KEY]"

	// Intentar crear un servicio de YouTube usando el contexto y la clave API.
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	// Si hay un error al crear el servicio, se detiene el programa y se muestra el error.
	if err != nil {
		log.Fatalf("Error al crear el servicio de YouTube: %v", err)
	}

	// Configurar la solicitud de búsqueda en YouTube con el término de búsqueda y un límite de 5 resultados.
	call := service.Search.List([]string{"snippet"}).Q(searchTerm).MaxResults(5)
	// Ejecutar la solicitud de búsqueda y obtener la respuesta.
	response, err := call.Do()
	// Si hay un error en la solicitud, se detiene el programa y se muestra el error.
	if err != nil {
		log.Fatalf("Error al realizar la solicitud de búsqueda: %v", err)
	}

	// Iterar a través de los resultados de la búsqueda.
	for _, item := range response.Items {
		// Imprimir el título y la descripción de cada resultado.
		fmt.Printf("Título: %s\n", item.Snippet.Title)
		fmt.Printf("Descripción: %s\n", item.Snippet.Description)
		// Si el resultado es un video de YouTube, se formate y muestra el URL.
		if item.Id.Kind == "youtube#video" {
			videoUrl := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId)
			fmt.Printf("Ver en YouTube: %s\n", videoUrl)
		}
		// Imprimir una línea en blanco para separar los resultados.
		fmt.Println()
	}
}

// MovieSearchResponse representa la respuesta de la búsqueda de películas de TMDB.
type MovieSearchResponse struct {
	Results []Movie `json:"results"` // Define una estructura para almacenar los resultados de películas, mapeando el JSON de la respuesta.
}

func Searchtmdb() {
	// Solicitamos al usuario que ingrese un término de búsqueda para películas.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingresa el nombre de la Película que quieres buscar: ")
	scanner.Scan()               // Leemos la entrada del usuario.
	searchTerm := scanner.Text() // Almacenamos el término de búsqueda ingresado en 'searchTerm'.

	apiKey := "52abf4732494d35e674f7b2345c0486f"     // Aquí está nuestra clave API para TMDB.
	escapedSearchTerm := url.QueryEscape(searchTerm) // 'QueryEscape' asegura que el término de búsqueda sea seguro para usar en una URL.
	// Formatear la URL para la solicitud de búsqueda de películas a TMDB.
	// 'fmt.Sprintf' se usa para construir una cadena con formato, insertando 'apiKey' y 'escapedSearchTerm' en la URL.
	// 'apiKey' es la clave API para autenticar la solicitud a TMDB.
	// 'escapedSearchTerm' es el término de búsqueda proporcionado por el usuario, escapado para ser seguro en una URL.
	// La URL incluye parámetros para especificar la clave API, el término de búsqueda, evitar contenido para adultos, el idioma de respuesta y la página de resultados.
	searchURL := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s&include_adult=false&language=es-ES&page=1", apiKey, escapedSearchTerm)

	req, err := http.NewRequest("GET", searchURL, nil) // Creamos una nueva solicitud HTTP GET con la URL de búsqueda.
	if err != nil {
		log.Fatal(err) // Si hay un error al crear la solicitud, lo mostramos y detenemos la ejecución.
	}

	req.Header.Add("accept", "application/json") // Añadimos una cabecera para aceptar respuestas en formato JSON.

	res, err := http.DefaultClient.Do(req) // Enviamos la solicitud y obtenemos la respuesta.
	if err != nil {
		log.Fatal(err) // Si hay un error al realizar la solicitud, lo mostramos y detenemos la ejecución.
	}
	defer res.Body.Close() // Cerramos el cuerpo de la respuesta al final.

	body, err := io.ReadAll(res.Body) // Leemos todo el cuerpo de la respuesta.
	if err != nil {
		log.Fatal(err) // Si hay un error al leer la respuesta, lo mostramos y detenemos la ejecución.
	}

	var response MovieSearchResponse      // Creamos una variable para almacenar los resultados deserializados.
	err = json.Unmarshal(body, &response) // Deserializamos el cuerpo de la respuesta (en formato JSON) a nuestra estructura.
	if err != nil {
		log.Fatal("Error al deserializar la respuesta: ", err) // Manejamos cualquier error en la deserialización.
	}

	// Iteramos sobre los resultados de películas y mostramos información relevante de cada una.
	for _, movie := range response.Results {
		fmt.Printf("Título: %s\n", movie.Title)
		fmt.Printf("Fecha de lanzamiento: %s\n", movie.ReleaseDate)
		fmt.Printf("Descripción: %s\n\n", movie.Overview)
	}
}

// BUSQUEDA EN YOUTUBE
// Estructura para almacenar la información relevante de cada video.
type Video struct {
	Title       string `json:"title"`       // 'Title' almacena el título del video, mapeado del campo 'title' en JSON.
	Description string `json:"description"` // 'Description' almacena la descripción del video, mapeado de 'description'.
}

func SearchYoutube() {
	ctx := context.Background()                         // Crear un contexto vacío para usar en la API de YouTube.
	apiKey := "AIzaSyC8LiUPsvPqJQow-VfgzueO9DchLHwr8Hk" // Definir la clave API para YouTube.

	// Crear un servicio de YouTube usando el contexto y la clave API.
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error al crear el servicio de YouTube: %v", err) // Manejar errores en la creación del servicio.
	}

	// Solicitar al usuario el término de búsqueda.
	fmt.Print("Ingrese el término de búsqueda: ")
	scanner := bufio.NewScanner(os.Stdin) // Crear un scanner para leer la entrada del usuario.
	scanner.Scan()                        // Leer la entrada del usuario.
	searchTerm := scanner.Text()          // Almacenar el término de búsqueda.

	// Configurar y realizar una solicitud de búsqueda en YouTube.
	call := service.Search.List([]string{"snippet"}).Q(searchTerm).MaxResults(5)
	response, err := call.Do() // Ejecutar la solicitud y obtener la respuesta.
	if err != nil {
		log.Fatalf("Error al realizar la solicitud de búsqueda: %v", err) // Manejar errores en la solicitud.
	}

	// Iterar sobre los resultados y mostrar título y descripción.
	for _, item := range response.Items {
		fmt.Printf("Título: %s\n", item.Snippet.Title)            // Imprimir el título del video.
		fmt.Printf("Descripción: %s\n", item.Snippet.Description) // Imprimir la descripción del video.

		// Si el resultado es un video de YouTube, crear y mostrar su URL.
		if item.Id.Kind == "youtube#video" {
			videoUrl := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId) // Formatear la URL del video.
			fmt.Printf("Ver en YouTube: %s\n", videoUrl)                                   // Imprimir la URL del video.
		}

		fmt.Println() // Imprimir una línea en blanco para separar los resultados.
	}
}

// ShowMoviesAndSearchRelatedContent muestra las películas y busca contenido relacionado en YouTube.
func ShowMoviesAndSearchRelatedContent() {
	// Solicitar al usuario la selección de una película.
	fmt.Println("Seleccione una película (ingrese el número):")
	apiKey := "52abf4732494d35e674f7b2345c0486f"                          // Clave API para TMDB.
	url := "https://api.themoviedb.org/3/movie/popular?api_key=" + apiKey // URL para obtener películas populares de TMDB.

	resp, err := http.Get(url) // Realizar una solicitud GET a la URL.
	if err != nil {
		log.Fatal("Error al obtener películas:", err) // Manejar errores en la obtención de películas.
	}
	defer resp.Body.Close() // Asegurarse de cerrar el cuerpo de la respuesta al final.

	var response MovieSearchResponse // Crear una variable para almacenar la respuesta deserializada.
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal("Error al decodificar respuesta:", err) // Manejar errores en la decodificación de la respuesta.
	}

	// Iterar sobre los resultados de películas y mostrarlos.
	for i, movie := range response.Results {
		fmt.Printf("%d: %s\n", i+1, movie.Title) // Imprimir el índice y el título de cada película.
	}

	scanner := bufio.NewScanner(os.Stdin)       // Crear un scanner para leer la entrada del usuario.
	scanner.Scan()                              // Leer la entrada del usuario.
	choice, err := strconv.Atoi(scanner.Text()) // Convertir la elección del usuario a un número entero.
	if err != nil || choice < 1 || choice > len(response.Results) {
		fmt.Println("Selección no válida") // Manejar selecciones inválidas.
		return
	}

	selectedMovie := response.Results[choice-1].Title // Obtener el título de la película seleccionada.
	fmt.Println("Has seleccionado:", selectedMovie)   // Mostrar la película seleccionada.

	// Buscar contenido relacionado en YouTube.
	ctx := context.Background()                                          // Crear un nuevo contexto vacío.
	ytApiKey := "AIzaSyC8LiUPsvPqJQow-VfgzueO9DchLHwr8Hk"                // Clave API para YouTube.
	service, err := youtube.NewService(ctx, option.WithAPIKey(ytApiKey)) // Crear un servicio de YouTube con la clave API.
	if err != nil {
		log.Fatalf("Error al crear el servicio de YouTube: %v", err) // Manejar errores en la creación del servicio de YouTube.
	}

	searchTerm := selectedMovie + " película"                                    // Formular el término de búsqueda agregando 'película' al título seleccionado.
	call := service.Search.List([]string{"snippet"}).Q(searchTerm).MaxResults(5) // Configurar la solicitud de búsqueda en YouTube.
	ytResponse, err := call.Do()                                                 // Ejecutar la solicitud y obtener la respuesta.
	if err != nil {
		log.Fatalf("Error al realizar la solicitud de búsqueda en YouTube: %v", err) // Manejar errores en la solicitud de búsqueda.
	}

	fmt.Println("Contenido relacionado en YouTube:") // Anunciar el inicio de los resultados de YouTube.
	for _, item := range ytResponse.Items {
		fmt.Printf("Título: %s\n", item.Snippet.Title)            // Imprimir el título de cada video relacionado.
		fmt.Printf("Descripción: %s\n", item.Snippet.Description) // Imprimir la descripción de cada video.
		if item.Id.Kind == "youtube#video" {
			videoUrl := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id.VideoId) // Formatear y generar la URL del video.
			fmt.Printf("Ver en YouTube: %s\n\n", videoUrl)                                 // Imprimir la URL del video.
		}
	}
}
