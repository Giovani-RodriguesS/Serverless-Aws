#!/bin/bash

echo "Itens a ser formatados: "

cd ../src/

modulos=($(ls -d */))
echo "${modulos[@]}"

for item in "${modulos[@]}"; do
    echo "Formatando: $item"
    cd $item
    go fmt ./...
    go mod tidy
    cd ..
done

echo "Formatação concluída!"