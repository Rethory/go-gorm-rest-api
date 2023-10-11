# Go REST API CRUD con PostgreSQL
Este proyecto es una REST API desarrollada en el lenguaje Go (Golang), que utiliza PostgreSQL como base de datos, Gorilla Mux, Gorm. La aplicación se ejecuta en un contenedor Docker para una fácil implementación y escalabilidad.
---
### Autor
#### Javier de Jesús Nájera Moreno
---
# Tabla de Contenidos
- Configuración
- Instalación
- Ejemplo de uso

# Configuración
Requiere tener previamente instalado [Go(Golang)](https://go.dev) y [Docker](https://www.docker.com/products/docker-desktop/)
---

Modulo Go Mux para el servidor
```
go get -u github.com/gorilla/mux
```
Modulo GORM
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```
Modulo Air (opcional)
```
go install github.com/cosmtrek/air@latest
```

# Instalación
Requiere tener previamente instalado [Go(Golang)](https://go.dev) y [Docker](https://www.docker.com/products/docker-desktop/)
---

## Clona el repositorio
```
git clone https://github.com/Rethory/go-gorm-rest-api.git

# Accede al directorio del proyecto
cd go-gorm-rest-api

# Ejecuta la aplicación ó usa 'air' si decidiste instalar el modulo air
go run main.go
```
# Ejemplo de uso de solicitudes HTTP de la API

Admite uso de:
- [x] [GET] http://localhost:3000/propietarios ó mascotas Para recibir la información de esa tabla.
- [x] [POST] http://localhost:3000/propietarios ó mascotas Para añadir un nuevo Propietario o Mascota en la base de datos.
- [x] [GET] http://localhost:3000/propietarios ó mascotas/{id} Para recibir la información en especifico de un propietario o mascota por su ID.
- [x] [DELETE] http://localhost:3000/propietarios ó mascotas/{id} Para eliminar un Propietario o Mascota por su ID.
---
### Tabla Mascotas

| Nombre      | Especie  | Raza              | Genero  | Peso | Temperatura | Observación        | MascotaID | PropietarioID |
|-------------|----------|-------------------|---------|------|-------------|--------------------|-----------|---------------|
|Firulais     | Perro    | Golden retriever  | Macho   | 34   | 37          | No tiene apetito   | 1         | 1             |
|Sam          | Gato     | Angora            | Hembra  | 12   | 38          | Infección          | 2         | 1             |
|Albert       | Perro    | Pastor aleman     | Macho   | 38   | 38          | Infección          | 3         | 2             | 

### Tabla Propietarios

| Nombre     | Apellido       | Telefono      |
|------------|----------------|---------------|
| Kevin      | Gonzales       | 123456789     |
| Alexa      | Ramirez        | 987654321     |

Código json
```
{
    "Nombre": "Firulais",
    "Especie": "Perro",
    "Raza": "Golden retriever",
    "Genero": "Macho",
    "Peso": 34,
    "Temperatura": 37,
    "Observacion": "No tiene apetitor",
    "MascotaID": 1,
    "PropietarioID": 1
}
```
```
{
    "Nombre": "Sam",
    "Especie": "Gato",
    "Raza": "Angora",
    "Genero": "Hembra",
    "Peso": 12,
    "Temperatura": 38,
    "Observacion": "Infeccion",
    "MascotaID": 2,
    "PropietarioID": 1
}
```
```
{
    "Nombre": "Albert",
    "Especie": "Perro",
    "Raza": "Pastor Aleman",
    "Genero": "Macho",
    "Peso": 38,
    "Temperatura": 38,
    "Observacion": "Infeccion",
    "MascotaID": 3,
    "PropietarioID": 2
}
```
```
{
    "nombre": "Kevin",
    "apellido": "Gonzales",
    "telefono": 123456789
}
```
```
{
    "nombre": "Alexa",
    "apellido": "Ramirez",
    "telefono": 987654321
}
```


