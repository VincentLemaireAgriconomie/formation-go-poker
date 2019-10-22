## installation : 

first step is to [download](https://golang.org/dl/) the language for your system, (this repository is written with go1.13 as target) 
then follow the [installation steps](https://golang.org/doc/install).

For the rest of the project, we will assume that we are on linux or macOs. we will alsao assume that we dont change the GOPATH and use ~/go

since go1.11 golang provides a package manager like composer or npm, to be sure it's always in use, we add

in `~/.bashrc`:

`alias go='GO111MODULE=on go'`

otherwise it wont be activated and you could use the wrong version on the local.

now if you try to build with `go build` like the [installation steps](https://golang.org/doc/install) says, it wont work anymore.

time to move on [compile and run](../01_compile_run/README.md)
