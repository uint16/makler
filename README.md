GOAL:
The goal of this project is to have a one stop portal for all things parliament related.

TODO:
See github issues



SETUP
Install and setup GOPATH and GOROOT - https://golang.org/doc/install

HOW TO RUN
Makler is beego application. To run it you will need to install beego

```go get github.com/astaxie/beego```

Clone the repository in any directory or subdirectory in GOPATH/src

Install dependencies for the project

```go get ./...```

Note: If the above command does not work install each/failed dependency separately

To run application in debug mode use

```bee run```

To run in production, build the project and run the executable

```go build main.go```

then

```./main```
