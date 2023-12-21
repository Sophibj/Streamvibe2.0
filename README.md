# Streamvibe2.0

<p align="center">
  <img src="https://github.com/Sophibj/Streamvibe2.0/blob/main/StreamVibe.jpg" alt="StreamVibe">
</p>

## Introducción
Streamvibe2.0 es un proyecto innovador para amantes del cine, utilizando APIs de YouTube y TMDb para sugerir y buscar películas, ofreciendo una plataforma dinámica y fácil de usar.

## Características
- **Sugerencias de Películas:** Utiliza APIs de YouTube y TMDb para recomendaciones personalizadas.
- **Funcionalidad de Búsqueda:** Capacidad eficiente para encontrar películas.
- **Cuentas de Usuario:** Permite a los usuarios crear cuentas personales.
- **Catálogos:** Navegación y visualización de extensos catálogos de películas.

## Lenguaje
Este código se encuentra programado en lenguage GO (Golang).

## Instalación
Para usar todas los módulos se deben pegar en la ruta de documentos y abrir con el IDE de preferencia, para este proyecto nosotros hemos usado VisualSutudio con el paquete de Lenguage GO.

## Uso

A continuación se describen los módulos principales de Streamvibe2.0 y cómo utilizarlos:

### [auth.go](https://github.com/Sophibj/Streamvibe2.0/blob/main/Streamvibe/auth/auth.go)

Este módulo gestiona la autenticación y administración de usuarios. Incluye funcionalidades para crear nuevos usuarios (`NewUser`) y para que los usuarios existentes inicien sesión (`Login`). Tras la autenticación, el usuario accede al menú principal (`UserMenu`), donde puede navegar por el catálogo de películas, visualizar géneros, buscar películas en TMDb y buscar videos en YouTube.

### [catalogo.go](https://github.com/Sophibj/Streamvibe2.0/blob/main/catalogo.go)

El módulo `catalogo` maneja la visualización y gestión del catálogo de películas. Permite a los usuarios ver un listado de películas populares y seleccionar una para más detalles. La función `ShowCatalog` conecta con la API de TMDb para recuperar y mostrar películas populares.

### [db.go](https://github.com/Sophibj/Streamvibe2.0/blob/main/db.go)

Este módulo configura y maneja la conexión con la base de datos SQL del proyecto (`Init` y `Close`). Aquí se definen operaciones básicas de base de datos que otros módulos utilizan para almacenar y recuperar datos.

### [generos.go](https://github.com/Sophibj/Streamvibe2.0/blob/main/generos.go)

`generos.go` se encarga de la gestión de géneros de películas. Permite mostrar una lista de géneros disponibles (`ShowGenres`) y añadirlos a la base de datos. Este módulo se conecta a la API de TMDb para obtener los géneros.

### [search.go](https://github.com/Sophibj/Streamvibe2.0/blob/main/search.go)

El módulo `search` proporciona funcionalidades para buscar películas en la base de datos de TMDb (`Searchtmdb`) y para buscar videos relacionados en YouTube (`SearchYoutube`). Permite a los usuarios buscar contenido específico basado en sus intereses.

### [main.go](https://github.com/Sophibj/Streamvibe2.0/blob/main/main.go)

`main.go` es el punto de entrada del programa. Este módulo inicia la conexión a la base de datos y presenta al usuario un menú inicial donde puede elegir entre crear un nuevo usuario, iniciar sesión o salir de la aplicación.



## Agradecimientos
Agradecimientos a los creadores del proyecto Sophia Ibarra (Sophibj) y Sebastián Becerra (Jugger_SB) por el repositorio.
Un agradecimiento a Juan David Moromenacho por el acompañamiento
