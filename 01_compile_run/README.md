## compile and run

within the current repository, you cannot use `color.go` as is because it's not a module, 
bye following the error message, you can correct this : `go mod init`

it will create a go.mod without any dependencies : 

```
module github.com/VincentLemaireAgriconomie/formation-go-poker/01_compile_run

go 1.13
```

but we need `"github.com/logrusorgru/aurora"`

we can get it with `go get`

you will see a new file `go.sum` containing the checksum for the project and his dependencies.

Now we can format our file with `go fmt` this ensure that everyone in go is following the same rules.

to run the programme : `go run .` 

to compile the programme : `go build .`