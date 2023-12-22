package auth

// Importaciones de paquetes necesarios para este módulo
import (
	"Streamvibe/catalogo" // Importa el paquete de catálogo que hemos creado, para manejar el catálogo de películas.
	"Streamvibe/db"       // Importa el paquete de base de datos para las operaciones de la base de datos.
	"Streamvibe/generos"
	"Streamvibe/search"
	"bufio"        // Importamos el paquete 'bufio'. 'bufio' se utiliza para la lectura y escritura en búfer. Un 'búfer' es como un 'almacén temporal' - bufio nos da funciones para la entrada y salida de búfer.
	"database/sql" //para trabajar con funciones que nos dejen interactuar con bases de datos SQL.
	"fmt"          // Proporciona funciones para formatear y imprimir texto.
	"log"          // Funciones de registro para errores y mensajes importantes
	"os"           //interfaz para interactuar con el sistema operativo.
	"strconv"      //funciones para convertir strings a otros tipos de datos.
)

// Definimos la estructura de una película:
type Pelicula struct {
	PeliculaID      int
	Titulo          string
	Director        string
	Duracion        int
	AnioLanzamiento int
	Descripcion     string
	Disponible      bool
	Genero          string
}

// declaramos la funcion NewUser para la creación de un nuevo usuario.
func NewUser() {
	var nombre, email, contraseña string // Declaramos variables locales para almacenar nombre, email y contraseña
	// Creamos  un nuevo Scanner para leer las entradas del usuario desde la consola
	scanner := bufio.NewScanner(os.Stdin) // 'bufio.NewScanner' crea un objeto que puede leer datos de 'os.Stdin', 'bufio' hace la lectura  más eficiente almacenando los datos leídos en un búfer
	fmt.Print("Ingrese nombre: ")
	scanner.Scan()               // 'Scan' lee el nombre ingresado por el usuario
	nombre = scanner.Text()      // 'Text' extrae la línea leída por 'Scan' como un string
	fmt.Print("Ingrese email: ") // Repetimos el proceso para email y contraseña.
	scanner.Scan()
	email = scanner.Text()
	fmt.Print("Ingrese contraseña: ")
	scanner.Scan()
	contraseña = scanner.Text()
	// Insertamos los datos del nuevo usuario en la base de datos SQL:
	// Aquí estamos diciendo a la bd que se quiere agregar un nuevo usuario. Utilizamos 'db.DB.Exec' para enviar un query a la base de datos:
	_, err := db.DB.Exec("INSERT INTO Clientes (Nombre, Email, Contraseña) VALUES (@p1, @p2, @p3)", nombre, email, contraseña)
	// Manejo de errores: Si algo sale mal o hay un problema con la base de datos:
	if err != nil { //si err es diferente de nil  entonces ejecutamos el código dentro de las llaves e informamos de dicho error.
		log.Fatal("Error al crear el usuario:", err)
	}
	fmt.Println("Usuario creado exitosamente") // Sino hay error informamos que se ha creado el usuario.
}

// Usamos la función Login para iniciar sesión:
func Login() int {
	var email, contraseña string          // declaramos variables para guardar el email y contraseña que el usuario va a escribir
	scanner := bufio.NewScanner(os.Stdin) // Pedimos al usuario su email y contraseña. Usamos  scanner para leer lo que escriba
	fmt.Print("Ingrese email: ")
	scanner.Scan()
	email = scanner.Text() // Guardamos el email en la variable email
	fmt.Print("Ingrese contraseña: ")
	scanner.Scan()
	contraseña = scanner.Text() // Guardamos la contraseña en la variable contraseña

	// buscamos en la base de datos si existe un usuario con ese email y contraseña:
	var clienteID int //// Aquí declaramos una variable 'clienteID' donde almacenaremos el ID del cliente si encontramos una coincidencia
	// 'db.DB.QueryRow' es un método que ejecuta una consulta SQL que espera devolver solo una fila (un solo resultado).
	//Consulta: estamos diciendo al sql que queremos que el ClienteID de la tabla Clientes donde el Email sea igual a @p1 y la Contraseña sea igual a @p2 se registre y genere un ID de Cliente
	err := db.DB.QueryRow("SELECT ClienteID FROM Clientes WHERE Email = @p1 AND Contraseña = @p2", email, contraseña).Scan(&clienteID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario y/o contraseña incorrecto") // Si no encontramos un usuario, mostramos un mensaje de error.
		} else {
			log.Fatal("Error al iniciar sesión:", err) // Si hay otro tipo de error al realizar la consulta, lo mostramos y terminamos la función.
		}
		return 0
	}

	// Si las credenciales coinciden, mostramos un mensaje de bienvenida y devolvemos el ID del cliente.
	fmt.Printf("Inicio de sesión exitoso. ¡BIENVENIDO A STREAMVIBE!! ClienteID: %d\n", clienteID)
	return clienteID
}

func UserMenu() {
	// Este es el menú que ve el usuario una vez iniciada la sesión.
	for {
		// Mostramos las opciones disponibles para el usuario.
		fmt.Println("Selecciona la Opción deseada")
		fmt.Println("1. Catalogo de Películas")
		fmt.Println("2. Ver Géneros de Películas")
		fmt.Println("3. Buscar Películas en TMDB")
		fmt.Println("4. Buscar Videos en YouTube")
		fmt.Println("5. Cerrar Sesión")
		fmt.Print("Ingrese su elección: ")

		// Creamos un nuevo scanner para leer la entrada del usuario desde la consola.
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()                              // Leemos la entrada del usuario.
		choice, err := strconv.Atoi(scanner.Text()) // Convertimos la entrada (texto) a un número entero.

		if err != nil {
			// Si hay un error en la conversión (por ejemplo, el usuario no ingresó un número), informamos y reiniciamos el bucle.
			log.Println("Opción no válida:", err)
			continue
		}

		// Usamos un switch para manejar las diferentes opciones basadas en la elección del usuario.
		switch choice {
		case 1:
			catalogo.ShowCatalog() // Si elige la opción 1, mostramos el catálogo de películas.
		case 2:
			generos.ShowGenres() // Si elige la opción 2, mostramos los géneros de películas.
		case 3:
			search.Searchtmdb() // Si elige la opción 3, realizamos una búsqueda en TMDB.
		case 4:
			search.SearchYoutube() // Si elige la opción 4, realizamos una búsqueda en YouTube.
		case 5:
			// Si elige la opción 5, imprimimos un mensaje de despedida y salimos del bucle con 'return'.
			fmt.Println("Sesión cerrada. ¡Gracias por visitar StreamVibe!")
			return
		default:
			// Si el usuario elige una opción que no está en el menú, le informamos que no es válida.
			fmt.Println("Opción no válida")
		}
	}
}
