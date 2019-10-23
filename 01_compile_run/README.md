## compile and run

within the current repository, you cannot use `color.go` as is because it's not a module, 
bye following the error message, you can correct this : `go mod init`

it will create a go.mod without any dependencies : 

```
module github.com/VincentLemaireAgriconomie/formation-go-poker/01_compile_run

go 1.13
```

but we need `"github.com/logrusorgru/aurora"`

`go get` to get it

you will see a new file `go.sum` containing the checksum for the project and his dependencies.

`go fmt` allow you to format your code this ensure that everyone in go is following the same rules. ( please do it before every commit )

`go run .` to run the programme 

`go build .` to compile the programme ( it will be in ./bin/ )

`go vet ./...` to check potential error

the [playground](https://play.golang.org/) basically just do a go run in a container.