# bazelinit

`bazelinit` is a CLI tool for initializing code repositories with Bazel. It simplifies the process of setting up a new project by creating the necessary Bazel configuration files for different languages.

## Installation

To install `bazelinit`, you can use `go install`:

```sh
go install github.com/yashsoodini/bazelinit@latest
```

Make sure your `GOPATH` is set up correctly and that `$GOPATH/bin` is in your `PATH`.

## Usage

To use `bazelinit`, run the command in the root directory of your project.

### Go

For Go projects, use the `go` subcommand. You need to provide the module path for your project.

```sh
bazelinit go --module_path=<your-go-module-path>
```

For example:

```sh
bazelinit go --module_path=github.com/my-user/my-project
```

This will create the following files:
- `MODULE.bazel`: Contains the necessary Go dependencies for Bazel.
- `BUILD`: A basic BUILD file that contains a :gazelle target to generate BUILD files for your Go packages.
- `go.mod`: A go.mod file for your project.

### C++

For C++ projects, use the `c++` subcommand.

```sh
bazelinit c++
```

This will create the following files:
- `MODULE.bazel`: Contains the necessary C++ dependencies for Bazel.
- `BUILD`: A basic BUILD file that contains a :gazelle target to generate BUILD files for your C++ packages.