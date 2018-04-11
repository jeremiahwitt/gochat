#COMP 445 - Lab Assignment 3
Jeremiah Witt - 40017534

## Info
This assignment was written with Go, using the following libraries:

- dep (https://github.com/golang/dep) to help with dependency management
- testify (https://github.com/stretchr/testify) to aid with the development of unit tests

Please follow the instructions above to install dep in your system. Once it has been installed, '
you can run the command `dep ensure`. This will pull all necessary dependencies into the `vendor` directory.

You do not have to compile if you do not wish - I have included a docker image prebuilt that you can use.

## Execution
It is intended that Docker will be used to execute the application. If you wish to build the image yourself,
please follow the instructions below. Otherwise, you can load the provided docker image as follows

```
docker load -i gochat.tar
```

This will allow you to use the `gochat` docker image on your system.

### Starting Docker Image
Regardless of how you get the image, you can start it with the following command:

``` 
docker run -it gochat
```

## Compilation
To build, you have two options:

### Build With Go
Go can be used directly to build the project. You _must_ unzip the source code into your Gopath directory. You can then 
compile by running the following in the project directory:

```
go build
```

This will create an executable called `comp445la3` (or whatever is the name of the parent directory).

### Build with Docker
Docker can make the compilation process even easier. It requires all necessary dependencies in the
`vendor` directory - however, I have included them with my submission. You can compile as follows from
within the project directory:

```
docker build -t gochat .
```

The Dockerfile uses a multistage build. The first stage uses a 'go' container to build the project within
a container. The second stage then takes the compiled application and places it within a smaller alpine container.

