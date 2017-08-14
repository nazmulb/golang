# golang
Go is a general-purpose language designed with systems programming in mind. It was initially developed at Google in the year 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is strongly and statically typed, provides inbuilt support for garbage collection, and supports concurrent programming.

Programs are constructed using packages, for efficient management of dependencies. Go programming implementations use a traditional compile and link model to generate executable binaries. The Go programming language was announced in November 2009 and is used in some of the Google's production systems.

Sample code:

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```

## Features of Go Programming
The most important features of Go programming are listed below −

- Support for environment adopting patterns similar to dynamic languages. For example, type inference (x := 0 is valid declaration of a variable x of type int)

- Compilation time is fast.

- Inbuilt concurrency support: lightweight processes (via go routines), channels, select statement.

- Go programs are simple, concise, and safe.

- Support for Interfaces and Type embedding.

- Production of statically linked native binaries without external dependencies.

## Features Excluded Intentionally
To keep the language simple and concise, the following features commonly available in other similar languages are omitted in Go −

- Support for type inheritance

- Support for method or operator overloading

- Support for circular dependencies among packages

- Support for pointer arithmetic

- Support for assertions

- Support for generic programming

## Installing Go

<a href="https://golang.org/dl/">Download the package file</a>, open it, and follow the prompts to install the Go tools.

## Code organization

### Overview
- Go programmers typically keep all their Go code in a single workspace.
- A workspace contains many version control repositories (managed by Git, for example).
- Each repository contains one or more packages.
- Each package consists of one or more Go source files in a single directory.
- The path to a package's directory determines its import path.

Note that this differs from other programming environments in which every project has a separate workspace and workspaces are closely tied to version control repositories.

### Workspaces
A workspace is a directory hierarchy with three directories at its root:

- `src` contains Go source files,
- `pkg` contains package objects, and
- `bin` contains executable commands.

The go tool builds source packages and installs the resulting binaries to the `pkg` and `bin` directories.

The src subdirectory typically contains multiple version control repositories (such as for Git or Mercurial) that track the development of one or more source packages.

To give you an idea of how a workspace looks in practice, here's an example:

```
bin/
    hello                          # command executable
pkg/
    darwin_amd64/
        github.com/nazmulb/golang/learning/
            nazmulpkg.a            # package object
src/
    github.com/nazmulb/golang/
        .git/                      # Git repository metadata
        learning/
            hello/
                hello.go           # command source
            nazmulpkg/
                show.go            # package source
                show_test.go       # test source
        README.md
    golang.org/x/image/
        .git/                      # Git repository metadata
        bmp/
            reader.go              # package source
            writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

The tree above shows a workspace containing two repositories (example and image). The example repository contains two commands (hello and outyet) and one library (stringutil). The image repository contains the bmp package and several others.

A typical workspace contains many source repositories containing many packages and commands. Most Go programmers keep all their Go source code and dependencies in a single workspace.

### The GOPATH environment variable
The `GOPATH` environment variable specifies the location of your workspace. It defaults to a directory named go inside your home directory.

The command `go env GOPATH` prints the effective current `GOPATH`:
```js
go env GOPATH
/Users/nazmulbasher/go
```

Edit your `~/.zshrc` file:
```js
vim ~/.zshrc
```

Add the following line to change the `GOPATH` to your workspace:
```js
export GOPATH=/Volumes/MyComputer/projects/golang
```

For convenience, add the workspace's `bin` subdirectory to your `PATH` as well:
```js
export PATH=$PATH:$(go env GOPATH)/bin
```

Now save and quit `~/.zshrc` file using vim and source the change from it:

```js
source ~/.zshrc
```

## Your first program

To compile and run a simple program, first choose a package path (we'll use github.com/nazmulb/golang/learning/hello) and create a corresponding package directory inside your workspace:

```js
mkdir $GOPATH/src/github.com/nazmulb/golang/learning/hello
```

Next, create a file named `hello.go` inside that directory, containing the following Go code.

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```

Now you can build and install that program with the go tool:

```js
go install github.com/nazmulb/golang/learning/hello
```

Note that you can run this command from anywhere on your system. The go tool finds the source code by looking for the `github.com/nazmulb/golang/learning/hello` package inside the workspace specified by `GOPATH`.

You can also omit the package path if you run `go install` from the package directory:

```js
cd $GOPATH/src/github.com/nazmulb/golang/learning/hello
go install
```

This command builds the hello command, producing an executable binary. It then installs that binary to the workspace's `bin` directory as `hello` (or, under Windows, hello.exe). In our example, that will be `$GOPATH/bin/hello`, which is `/Volumes/MyComputer/projects/golang/bin/hello`.

The go tool will only print output when an error occurs, so if these commands produce no output they have executed successfully.

You can now run the program by typing its full path at the command line:

```js
$GOPATH/bin/hello
Hello, world.
```

Or, as you have added `$GOPATH/bin` to your PATH, just type the binary name:

```js
hello
Hello, world.
```