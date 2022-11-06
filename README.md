# SWE
Backend is written in Golang. It performs authentification and user registration for admin.
There is a Dockerfile given. Comands to build and run Docker:
```
docker build -t {name} .
docker run -p 8080:8080 {name}
```
We run our backend and frontend in the same host machine. 
You can also run backend by going to its repository and running:
```
go run main.go
```
Frontend is written in Nuxt(Vue.js). To run Frontend you should write these commands:
```
yarn install
yarn dev
```
This is the prototype of our registration system. We will perfect it before the presentation time.
The main struggle was working with different people and adjusting to each other. 
