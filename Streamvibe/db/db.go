package db

import (
	"database/sql" // Importa el paquete de SQL para la interacción con bases de datos.
	"log"          // Importa el paquete log para registrar mensajes de error.

	_ "github.com/denisenkom/go-mssqldb" // Importa el driver MSSQL para Go.
)

var DB *sql.DB // DB es una variable global que mantendrá la conexión a la base de datos.

// Init inicializa la conexión a la base de datos.
func Init() {
	var err error
	// sql.Open establece una conexión a la base de datos.
	// "sqlserver" es el nombre del driver, y la cadena de conexión sigue.
	DB, err = sql.Open("sqlserver", "server=DESKTOP-05LFD9C;port=1433;database=Servicio")
	if err != nil {
		// Si hay un error al abrir la conexión, se registra y se termina el programa.
		log.Fatal("Error al conectar a la base de datos:", err)
	}
}

// Close cierra la conexión a la base de datos.
func Close() {
	// DB.Close() cierra la conexión a la base de datos.
	if err := DB.Close(); err != nil {
		// Si hay un error al cerrar la conexión, se registra y se termina el programa.
		log.Fatal("Error al cerrar la conexión de la base de datos:", err)
	}
}
