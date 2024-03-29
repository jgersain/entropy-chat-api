# REST API Entropy chat

Api Rest en donde los usuarios pueden realizar una serie de acciones como comunicarse con sus amigos y familiares de una forma sencilla utilizando el lenguaje de programación Go

## Instalación

- Instalar [Go Programming Language](https://golang.org/doc/install)

- Instalar [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

- Clonar el repositorio `git clone git@github.com:jgersain/entropy-chat-api.git`

- Moverse al directorio del proyecto `cd entropy-chat-api`

- Instalar los siguientes paquetes

```
$ go get github.com/joho/godotenv
$ go get github.com/jinzhu/gorm
$ go get golang.org/x/crypto/bcrypt
$ go get github.com/dgrijalva/jwt-go
$ go get github.com/gorilla/mux
$ go get github.com/jinzhu/gorm/dialects/postgres
$ go get github.com/badoux/checkmail
$ go get github.com/stretchr/testify
```

- Crear una instancia de postgres utilizando docker. En el ejemplo el usuario es `pgdev` y la contraseña es `developer`

```
$ docker pull postgres:latest
$ docker run --name postgres-entropy -p 127.0.0.1:5432:5432 -d postgres:latest
$ docker exec -it postgres-entropy psql -U postgres -c "create role pgdev with login superuser password 'developer'"
```

- Crear el archivo `.env` con el contenido del template `.env.example` y establecer su configuración 

- Ejecutar el servidor `go run main.go`, por default se ejecuta en [http://127.0.0.1:8080](http://127.0.0.1:8080)

## Documentación

### Endpoints

La raíz del API en desarrollo se encuentra en http://127.0.0.1/api

### Registrarse

```
POST /users
Accept: application/json
```

##### Esquema del body

```
{
    "name":"John Connor",
    "age":"26",
    "email":"john@connor.com",
    "password":"startx",
    "profile_photo":""
}
```

##### Respuesta

| Status        | Significado   | Descripción                            |
| ------------- |:-------------:| :-------------------------------------:|
| 201           | Created       | Información creada satisfactoriamente  |
| 400           | Bad Request   | Los parámetros enviados no son válidos |

Respuesta con estatus 201

```
{
    "id": 1,
    "name": "John Connor",
    "age": "26",
    "email": "john@connor.com",
    "profile_photo": "",
    "password": "$2a$10$d2rI9YzFvSGNxwu2z.GDEO9MUrxyj08KLiK14.u4G67nu086F2Xni",
    "created_at": "2019-12-13T20:24:43.2412471-06:00",
    "updated_at": "2019-12-13T20:24:43.2412471-06:00",
}
```

Respuesta con estatus 400

```
{
    "error": "El correo electrónico es necesario"
}
```

### Autenticarse

```
POST /login
Accept: application/json
```

##### Esquema del body

```
{
    "email":"john@connor.com",
    "password":"startx"
}
```

##### Respuesta

| Status        | Significado   | Descripción                                                  |
| ------------- |:-------------:| :-----------------------------------------------------------:|
| 200           | OK            | La solicitud ha tenido exito                                 |
| 401           | Unauthorized  | El usuario no tiene permiso de realizar la acción solicitada |

Respuesta con estatus 200

```
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1NzYzMjE4NDEsInVzZXJfaWQiOjF9.FlY4DR9PVM4Xfcy5In8rcnk2sTLUlpapv3ovWrUReNM",
    "token_type": "Bearer"
}
```

Respuesta con estatus 401

```
{
    "error": "valores no válidos"
}
```

---
**NOTA**

Para interactuar con los demás endpoints es necesario generar un token

---

### Editar información del perfil

```
PUT /users/{id}
Accept: application/json
```

##### Esquema del body

```
{
    "name":"Juan Contreras",
    "age":"45",
    "profile_photo":"",
    "email":"juan@contreras.com",
    "password":"startx"
}
```

##### Respuesta

| Status        | Significado   | Descripción                                                  |
| ------------- |:-------------:| :-----------------------------------------------------------:|
| 200           | OK            | La solicitud ha tenido exito                                 |
| 401           | Unauthorized  | El usuario no tiene permiso de realizar la acción solicitada |

Respuesta con estatus 200

```
{
    "age": "45",
    "created_at": "2019-12-14T10:29:35.014696Z",
    "email": "johny@connor.com",
    "id": 1,
    "name": "John Connor Díaz",
    "password": "$2a$10$40aOB6LuKb5kXHb8FKA/CuASHsuPBZ7KHUnqL85JRT5jb8w5w5CJK"
    "updated_at": "2019-12-14T10:29:35.014696Z",
}
```

Respuesta con el estatus 401

```
{
    "error": "No autorizado"
}
```

### Anadir contacto

```
POST /contacts
Accept: application/json
```

##### Esquema del body

```
{
    "name":"Cristiano Ronaldo",
    "nickname":"CR/",
    "email":"cr7@mail.com",
    "phone":"5546768765",
    "address":"Villa de las manzanas 208",
    "user_id":1
}
```

##### Respuesta

| Status        | Significado   | Descripción                                                  |
| ------------- |:-------------:| :-----------------------------------------------------------:|
| 201           | OK            | Información creada satisfactoriamente                        |
| 401           | Unauthorized  | El usuario no tiene permiso de realizar la acción solicitada |

Respuesta con estatus 201

```
{
    "id": 1,
    "name": "Cristiano Ronaldo",
    "nickname": "CR7",
    "email": "cr7@mail.com",
    "phone": "5546768765",
    "address": "Villa de las manzanas 208",
    "user_id": 1,
    "created_at": "2019-12-14T17:15:18.7959126-06:00",
    "updated_at": "2019-12-14T17:15:18.7959126-06:00"
}
```

Respuesta con el estatus 401

```
{
    "error": "No autorizado"
}
```

### Consultar lista de contactos

```
GET /contacts?user_id={id}
Accept: application/json
```

##### Esquema del body

##### Respuesta

| Status        | Significado   | Descripción                                                  |
| ------------- |:-------------:| :-----------------------------------------------------------:|
| 200           | OK            | La solicitud ha tenido exito                                 |
| 401           | Unauthorized  | El usuario no tiene permiso de realizar la acción solicitada |

Respuesta con estatus 200 ```/contacts?user_id=1```

```
[
    {
        "id": 1,
        "name": "Cristiano Ronaldo",
        "nickname": "CR/",
        "email": "cr7@gmail.com",
        "phone": "5546768765",
        "address": "Villa de las manzanas 208",
        "user_id": 1,
        "created_at": "2019-12-15T02:09:02.771054Z",
        "updated_at": "2019-12-15T02:09:02.771054Z"
    },
    {
        "id": 2,
        "name": "Leonel Messi",
        "nickname": "Lio",
        "email": "lio@mail.com",
        "phone": "5546768799",
        "address": "Villa de las manzanas 209",
        "user_id": 1,
        "created_at": "2019-12-15T02:23:22.693701Z",
        "updated_at": "2019-12-15T02:23:22.693701Z"
    }
]
```

Respuesta con el estatus 401

```
{
    "error": "No autorizado"
}
```

### Ver un contacto en especifico

```
GET /contacts/{id}?user_id={id}
Accept: application/json
```

##### Esquema del body

##### Respuesta

| Status        | Significado   | Descripción                                                  |
| ------------- |:-------------:| :-----------------------------------------------------------:|
| 200           | OK            | La solicitud ha tenido exito                                 |
| 401           | Unauthorized  | El usuario no tiene permiso de realizar la acción solicitada |

Respuesta con estatus 200 ```/contacts/1?user_id=1```

```
{
    "id": 1,
    "name": "Cristiano Ronaldo",
    "nickname": "CR/",
    "email": "cr7@gmail.com",
    "phone": "5546768799",
    "address": "Villa de las manzanas 209",
    "user_id": 1,
    "created_at": "2019-12-15T05:39:37.499016Z",
    "updated_at": "2019-12-15T05:39:37.499016Z"
}
```

Respuesta con el estatus 401

```
{
    "error": "No autorizado"
}
```






