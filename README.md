# skeleton-golang-package #

[![GitHub Build Status](https://github.com/cisagov/skeleton-golang-package/workflows/build/badge.svg)](https://github.com/cisagov/skeleton-golang-package/actions)

This is a generic skeleton project that can be used to quickly get a
new [cisagov](https://github.com/cisagov) Go project started.
This skeleton project contains [licensing information](LICENSE), as
well as [pre-commit hooks](https://pre-commit.com) and
[GitHub Actions](https://github.com/features/actions) configurations
appropriate for a Go package.

## Installation ##

To install the example package, use `go build`:

```console
go build --ldflags "-X 'main.version=$(git describe --tag --abbrev=0)'"  -o skeleton-golang-package-$(git describe --tag --abbrev=0)
```

This will produce an executable `skeleton-golang-package-v*` with an appended
version number within the current directory.

Execute the `./skeleton-golang-package-v*` binary to run the program.

You should see the contents of the main package as defined in [main.go](main.go).

### Nix ###

If you have [Nix](https://nixos.org/download.html) installed you can use
the [flake.nix](flake.nix) configuration file located at the root of the
project to build and run the `skeleton-golang-package` package:

```console
nix run
```

After the build has completed the program will automatically execute. The build
will also have created a `result` directory within the root of the project.
This directory contains a `bin` subdirectory where the executable
`skeleton-golang-package` can be found:
`result/bin/skeleton-golang-package`

To run the program simply execute the binary from the project root:

```console
result/bin/skeleton-golang-package
```

Alternatively, if you want to execute an interactive build environment
with all available dependencies for this project you can run this
command in the project root:

```console
nix develop
```

If you make changes to the project files you will need to update `vendorSha256`
in `flake.nix`  You can do this by running `nix build` and you should be presented
with an error similar to:

```console
error: hash mismatch in fixed-output derivation '/nix/store/11a71sp1wynjgxinx6yb0yhli4q659zi-skeleton-golang-package-20230424-go-modules.drv':
         specified: sha256-+aBatR0sHWkykrQr8AMD0P5xoc4VEXtA+egp6PxsKzY=
            got:    sha256-gqfvjULp2VApWQl7yFVj45meYpcS4XefUtEUy+TtAH4=
error: 1 dependencies of derivation '/nix/store/nm6xhgfawf67sy89rl9azdzvnjc5r7cr-skeleton-golang-package-20230424.drv' failed to build
```

You would then copy the `got` hash and use it to replace the previous
 `vendorSha256` value within `flake.nix`.

If `flake.nix` changes or is modified you will also need to update the
`flake.lock` file. You can do this by running: `nix flake update`. It is
also recommended to then run `nix build` after this operation to ensure
that the build is successful.

## SBOM ##

A Software Bill of Materials (SBOM) is like a recipe card for software,
detailing all the ingredients (components, libraries, dependencies)
used to cook up a digital product. It's a crucial tool for understanding
what goes into software, helping with things like spotting vulnerabilities,
ensuring license compliance, and keeping an eye on supply chain security.
However, SBOM isn't a magic fix for security issues itself—it's more of a
roadmap that helps teams navigate potential risks. It's not a substitute
for good coding practices or regular updates, but a helpful companion in
the journey towards stronger software security.

The SBOM found in this project was generated using [cyclonedx-gomod](https://github.com/CycloneDX/cyclonedx-gomod).
The exact commands used can be found below.

```bash
cyclonedx-gomod mod -std -json -output bom.json  ./
```

Generates an SBOM in json format, includes the aggregate of modules
required by all packages in the target module. This optionally
includes modules required by tests and test packages. Build
constraints are NOT evaluated, allowing for a "whole picture"
view on the target module's dependencies.

```bash
cyclonedx-gomod app -std -json -output sbom.json  ./
```

Generates an SBOM in json format for only the
modules the target application depends on. Modules required by tests
or packages that are not imported by the application are not included.
Build constraints are evaluated, which enables a very detailed view of
what's really compiled into an application's binary.

This project does not inlude executables but if that changes in the future
or if the application is modified to include additional dependencies or updates
the SBOM should be updated to reflect that.

## Testing ##

You can execute tests for the `example` package by running the following
command from within the project root directory:

```console
go test -v test/example_test.go
```

This will run tests for the `example` package as defined in
[`example_test.go`](test/example_test.go).

## Documentation ##

If you would like to view this project's documentation
using `godoc`, install `godoc` by running:

```console
go install golang.org/x/tools/cmd/godoc@latest
```

Then, run the command below and point your browser to `http://localhost:6060`:

```console
godoc -http=:6060
```

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
