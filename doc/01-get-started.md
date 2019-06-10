# Get started

## Install go
https://golang.org/dl/

```bash
# mac brew
brew install go
```

## Workspace

### Folder structure
Recommended workspace structure [(link)](https://golang.org/doc/code.html#Workspaces):
```
/Users/<user.name>/Documents/Workspace/go/
    bin/
    pkg/
    src/ #libraries will be installed here
        github.com/
            <user_name>/
                <project_name>/
```

### env
Set **GOPATH** environment variable
[(github)](https://github.com/golang/go/wiki/SettingGOPATH#zsh)
[(golang)](https://golang.org/doc/code.html#GOPATH)

```
export GOPATH=$HOME/Documents/Workspace/go
export GOBIN=$GOPATH/bin
export PATH=$GOBIN:$PATH
```

`$GOPATH` configuration is not needed if you use [Go Modules](15-go-modules.md)

## Hello World
To import a package (for example "fmt"):

```go
// single import
import "fmt"

// grouped imports
import (
	"os"
	"strings"
)
```
Dot imports, underscore imports:
```go
package main

import (
    "fmt"
    . "fmt" // 1 
    _ "log" // 2
    format "fmt" // 3
)

func main() {
    fmt.Println("hello world")
    Println("hello world")    
    format.Println("hello world")
}
```
- **1**: With dot before the package doesn't force to put fmt. before Println
- **2**: The package could not be used (for development purposes)
- **3**: Alias import

## Go Commands

https://golang.org/cmd/go/

```bash
# Version
go version

# Run compiles and runs the main package comprising
# the named Go source files.
go run file.go

# Compiles and build an executable
# if file name is not written it will compile a bin 
# with the folder (package) name 
go build hello.go

# Compiles, build and generate a binary inside the 
# workspace/bin folder
go install 

# Clean all the executables in the current folder
go clean

# Get a Package. It will be downloaded in $GOPATH/src directory
go get github.com/satori/go.uuid
go get github.com/slaparra/go-training
go get -u github.com/slaparra/go-training //to get the updates

# Show go environment variables
go env

# Format the code
go fmt

# With -d prints the diff
gofmt -d -w file.go

# Golint is a linter for Go source code.
# Gofmt reformats Go source code, whereas golint prints out style mistakes.
golint ./...

# Report suspicious constructs
go vet

# Show race conditions info
go run --race gofile.go
```

- [Oficial documentation: Golang wiki](https://github.com/golang/go/wiki)
- [Gofmt](https://golang.org/cmd/gofmt/)
- [Golint](https://github.com/golang/lint)

## Cross-compile
*One of Go's most powerful features is the ability to cross-build executables for any Go-supported foreign platform. This makes testing and package distribution much easier, because you don't need to have access to a specific platform in order to distribute your package for it.* (digitalocean.com)

```
env GOOS=target-OS GOARCH=target-architecture go build package-import-path
```
<small>([Possible combinations of GOOS and GOARCH you can use])</small>

[Possible combinations of GOOS and GOARCH you can use]: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures

- [How To Build Go Executables for Multiple Platforms](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)
- [Cross compilation just got a whole lot better in Go 1.5](https://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5)
- [Cross Compiling, Golang Cookbook](https://golangcookbook.com/chapters/running/cross-compiling/)