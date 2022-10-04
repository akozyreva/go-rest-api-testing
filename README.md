# Go-rest-api-testing
## Pre-requirements
- Go must be installed
- Specify GOPATH, like adding in `.zprofile`:

```
export GOPATH=$HOME/projects/go-workspace
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```
Next projects must be located under in `GOPATH` folder
## How to run tests
```
go test
```
## Docker
To build image:
```
docker build  -t go-rest-api .
```
To run image:
```
docker run -it --rm -v ${PWD}:/rest-api-basics/ go-rest-api 
```
