# Decks API

This project provides a REST API for working with card decks.

It provides the following endpoints:

    GET /decks/{uuid} - Returns a deck with the given uuid.
    POST /decks - Creates a new deck.
    PUT /decks/{uuid}/{count} - Draws cards from the deck.

The application use a postgres database to store the decks and cards.

You also need to export some environment variables related to DATABASE:

```
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_NAME=decksapi
```

# Possible improvements

I would like to make the code more readable, modular and with a better architecture, but due to my small knowledge in Golang 
I had to study a lot to learn some patterns, learn the language and understand some good practices. I think that with a little
time I can become much more proficient in Golang and I can make better code.
I would like for example to better separate some concerns on the application, put persistence code in repositories, create use cases
for the entrypoints, etc, but really don't know yet the better way to do this in Golang.

# Run Tests

To run tests, run `go test -v ./...`

# Developing

### To have the enable live reload while developing, you can use nodemon.

1) Install nodemon:
```
npm i -g nodemon
```

2) Execute the application with nodemon:
```
nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
```

# TODO

- [ ] Criar Testes
- [ ] Atualizar README