# Comandos Utilizados

    docker-compose up -d
    docker exec -it appproduct bash

Como nós vamos trabalhar usando 'Go Mod' que é uma forma do Go gerenciar as dependências externas automaticamente, precisamos criar o nome do nosso módulo:

    go mod init github.com/taakahashi/hexagon-with-go