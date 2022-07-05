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
