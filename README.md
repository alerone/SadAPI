# ToDo List API 📒 ![GitHub top language](https://img.shields.io/github/languages/top/alerone/SadAPI?color=%2377CDFF) ![GitHub last commit](https://img.shields.io/github/last-commit/alerone/SadAPI?color=%23bc0bbf) ![GitHub Created At](https://img.shields.io/github/created-at/alerone/SadAPI?color=%230dba69) ![GitHub repo size](https://img.shields.io/github/repo-size/alerone/SadAPI?color=%23390385)




## _Tarea para la asignatura SAD sobre los seminarios 1 y 2_

A continuación se detallará la forma de usar este programa y todas las funcionalidades que trae el mismo.

## Instalación 💻⬇️

Para desplegar el API en el dispositivo utilizamos docker-compose (se requiere de docker para desplegar el proyecto). Para ello debes situarte en el directorio raíz del proyecto y ejecutar en el terminal:

```bash
docker-compose up --build
```

Al ejecutar esa línea por terminal ya estará todo el proyecto desplegado. Ahora falta realizar las consultas pertinentes al API en la dirección **http://localhost:8080** para gestionar tus ToDos. 😁

### Explicación en detalle
Para desplegar todo el sistema se utilizan dos contenedores: `SadAPI` y una base de datos `postgres`. Con [`docker compose`](./docker-compose.yaml) conseguimos automatizar el despliegue de ambos contenedores de forma que estos contenedores se pueden conectar para que el API obtenga los credenciales de la base de datos y su dirección.

Para guardar las credenciales fuera del archivo docker-compose he utilizado un archivo [`.env`](./src/.env) donde se almacena la información de usuario, contraseña y nombre de la base de datos.

La imagen de la base de datos es una imagen `postgres:latest` que es extraída del registro público de Docker. La imagen SadApi, en cambio, es una imagen creada por mí ([SadApi Dockerfile](./src/Dockerfile)) que construye el código de la aplicación en la primera fase y luego extrae lo unicamente necesario al contenedor final para tener una imagen y un contenedor lo más liviano posible.

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

### Tips de uso 📓
Puedes utilizar la herramienta `curl` para probar estas funcionalidades. Por ejemplo:

**Obtener todas las tareas**
```bash
curl -X GET http://localhost:8080/
```
**Crear una ToDo**
```bash
curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d '{
    "title": "Título de la tarea",
    "description": "Descripción de la tarea",
    "completed": false
}'
```

También puedes utilizar herramientas gráficas para enviar peticiones REST al API como `postman` o alguna extensión de `Visual Studio Code` como `ThunderClient`


## Tecnología utilizada 🤖🖥️

La tecnología que se ha utilizado para el desarrollo del proyecto es:

- [Golang](https://go.dev/) - Lenguaje de programación utilizado

- [Gin](github.com/gin-gonic/gin) - Librería de Go para crear un API REST

- [pq](github.com/lib/pq) - Librería de Go para manipular una base de datos en PostgreSQL

- [docker](https://www.docker.com/) - Tecnología de para automatizar la implementación de aplicaciones en contenedores
