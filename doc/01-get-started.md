# Get started

## Install go
https://golang.org/dl/

```
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


## Go Commands

https://golang.org/cmd/go/

```
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

# Show go environment variables
go env

# Format the code
go fmt

# Golint is a linter for Go source code.
# Gofmt reformats Go source code, whereas golint prints out style mistakes.
golint ./...

# Show race conditions info
go run --race gofile.go
```

- [Oficial documentation: Golang wiki](https://github.com/golang/go/wiki)
- [Golint](https://github.com/golang/lint)