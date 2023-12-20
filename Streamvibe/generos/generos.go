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

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GenreList struct {
	Genres []Genre `json:"genres"`
}

func ShowGenres() {
	apiKey := "52abf4732494d35e674f7b2345c0486f"
	url := fmt.Sprintf("https://api.themoviedb.org/3/genre/movie/list?api_key=%s&language=es", apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var genreList GenreList
	err = json.Unmarshal(body, &genreList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lista de Géneros:")
	for i, genre := range genreList.Genres {
		fmt.Printf("%d: %s\n", i+1, genre.Name)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Ingrese el número del género que desea seleccionar: ")
	scanner.Scan()
	input := scanner.Text()
	selection, err := strconv.Atoi(input)
	if err != nil || selection < 1 || selection > len(genreList.Genres) {
		fmt.Println("Selección no válida")
		return
	}

	selectedGenre := genreList.Genres[selection-1]

	// Insertar el género seleccionado en la base de datos
	_, err = db.DB.Exec("INSERT INTO Genres (GenreID, Name) VALUES (@p1, @p2)",
		selectedGenre.ID, selectedGenre.Name)

	if err != nil {
		log.Printf("Error al insertar el género seleccionado en la base de datos: %v\n", err)

	}
}
