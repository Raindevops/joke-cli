# ✨ Joke CLI ✨

Joke CLI is a command line interface used to print jokes in the terminal.

This project use external API to get jokes : 
- [Chuck norris API](https://api.chucknorris.io/jokes/)
- [icanhasdadjoke](https://icanhazdadjoke.com/)
- [Official joke api](https://official-joke-api.appspot.com/)

Developped with [cobra](https://github.com/spf13/cobra).

## 🔨 How to build the project 🔨

Joke CLI is publicly available on my [dockerhub repository](https://hub.docker.com/repositories/raindevops).

If you want to build an run this project locally : 

    go get . 
    go build -o joke-cli main.go
    ./joke-cli

## 🏃 Run the project🏃

Joke CLI can print 3 types of Jokes :

- Chuck Norris phrase
- Dad's jokes
- General jokes

💻 Get one of each by running : 

    podman run --rm raindevops/joke-cli:1.0.0 chuckNorris
    
    podman run --rm raindevops/joke-cli:1.0.0 dadjoke

    podamn run --rm raindevops/joke-cli:1.0.0 jokes

Each command has their own flags to customize the joke you want to have

for more informations : 

    podman run --rm raindevops/joke-cli:1.0.0 jokes --help