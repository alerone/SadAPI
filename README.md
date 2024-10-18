# ToDo List API 📒

![GitHub top language](https://img.shields.io/github/languages/top/alerone/SadAPI?color=%2377CDFF)


## _Tarea para la asignatura SAD sobre los seminarios 1 y 2_

A continuación se detallará la forma de usar este programa y todas las funcionalidades que trae el mismo.

## Instalación 💻⬇️

Para desplegar el API en el dispositivo utilizamos docker-compose (se requiere de docker para desplegar el proyecto). Para ello debes situarte en el directorio raíz del proyecto y ejecutar en el terminal:

```bash
docker-compose up --build
```

Al ejecutar esa línea por terminal ya estará todo el proyecto desplegado. Ahora falta realizar las consultas pertinentes al API en la dirección **localhost:8080** para gestionar tus ToDos. 😁

## Uso 🧠

A continuación se detallan las posibles consultas que se pueden realizar al API.

- `Crear` una ToDo: `POST` a http://localhost:8080/
- Obtener `todas` las tareas: `GET` a http://localhost:8080/
- `Buscar una` ToDo por ID: `GET` a http://localhost:8080/:id
- Buscar varias ToDo por `título`: `GET` a http://localhost:8080/title/:title
- `Actualizar` una ToDo por ID: `PUT` a http://localhost:8080/:id
- `Borrar` una ToDo por ID: `DELETE` a http://localhost:8080/:id
- Marcar una ToDo como `Completada` por ID: `Get` a http://localhost:8080/complete/:id

Importante destacar que cuando pone `:id` o `:title` debes cambiar eso por el id o el título de la tarea a la que quieres realizar la acción.

## Tecnología utilizada 🤖🖥️

La tecnología que se ha utilizado para el desarrollo del proyecto es:

- [Golang](https://go.dev/) - Lenguaje de programación utilizado

- [Gin](github.com/gin-gonic/gin) - Librería de Go para crear un API REST

- [pq](github.com/lib/pq) - Librería de Go para manipular una base de datos en PostgreSQL

- [docker](https://www.docker.com/) - Tecnología de para automatizar la implementación de aplicaciones en contenedores
