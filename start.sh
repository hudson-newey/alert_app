#!/bin/bash
if [ ! -f ./src/server ]; then
    go build -o ./src/server ./src/server.go;
fi

./src/server ./src/app.html
