# entropy-chat-api

Api Rest en donde los usuarios pueden realizar una serie de acciones como comunicarse con sus amigos y familiares de una forma sencilla utilizando el lenguaje de programación Go

## Instalación

- Instalar [Go Programming Language](https://golang.org/doc/install)

- Instalar [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

- Clonar el repositorio `git clone git@github.com:jgersain/entropy-chat-api.git`

- Moverse al directorio del proyecto `cd entropy-chat-api`

- Instalar los siguientes paquetes

```
go get github.com/joho/godotenv
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/postgres
```