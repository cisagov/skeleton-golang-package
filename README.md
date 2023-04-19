# skeleton-golang-package #

[![GitHub Build Status](https://github.com/cisagov/skeleton-golang-package/workflows/build/badge.svg)](https://github.com/cisagov/skeleton-golang-package/actions)

This is a generic skeleton project that can be used to quickly get a
new [cisagov](https://github.com/cisagov) Go project started.
This skeleton project contains [licensing information](LICENSE), as
well as [pre-commit hooks](https://pre-commit.com) and
[GitHub Actions](https://github.com/features/actions) configurations
appropriate for the major languages that we use.

In many cases you will instead want to use one of the more specific
skeleton projects derived from this one.

## Installation ##

To install the example package, use `go build`:

```console
cd ./skeleton-golang-package/src/example
go build
```

This will produce an executable `example` within the current directory.

Execute the `./example` binary to run the program.

You should see the contents of the main package as definied in `main.go` as expected.

### Nix ###

If you have [Nix](https://nixos.org/download.html) installed you can use
the available [flake.nix](flake.nix) within the root of the project to build
the `example` package:

```console
cd ./skeleton-golang-package
nix build
```

After the build has completed you will see a `result` directory within the
root of the project. In that directory contains a `bin` directory where the
executable `example` can be found: `result/bin/example`

You can execute the binary from the root of the `skeleton-golang-package`
directory to run the program:

```console
./result/bin/example
```

Alternatively, if you want to execute an interactive build environment
with all available dependencies for this project you can run the below
 command in the project root:

```console
cd ./skeleton-golang-package
nix develop
```

## Testing ##

You can execute tests for the  `example` package by running the below
command from within the `./skeleton-golang-package/src/example` directory:

```console
cd ./skeleton-golang-package/src/example
go test -v test/example_test.go
```

This will run tests for the `example` package as defined in
`./skeleton-golang-package/src/example/test/example_test.go`

## New Repositories from a Skeleton ##

Please see our [Project Setup guide](https://github.com/cisagov/development-guide/tree/develop/project_setup)
for step-by-step instructions on how to start a new repository from
a skeleton. This will save you time and effort when configuring a
new repository!

## Contributing ##

We welcome contributions!  Please see [`CONTRIBUTING.md`](CONTRIBUTING.md) for
details.

## License ##

This project is in the worldwide [public domain](LICENSE).

This project is in the public domain within the United States, and
copyright and related rights in the work worldwide are waived through
the [CC0 1.0 Universal public domain
dedication](https://creativecommons.org/publicdomain/zero/1.0/).

All contributions to this project will be released under the CC0
dedication. By submitting a pull request, you are agreeing to comply
with this waiver of copyright interest.
