# Step 3: Set Up the Go Toolchain

Building on the analysis from [STEP1_ANALYSIS.md](STEP1_ANALYSIS.md) and the repository layout proposed in [STEP2_DESIGN.md](STEP2_DESIGN.md), this step prepares the repository to build Go code.

## Install Go

Install Go 1.23 or higher. The official installer sets the `GOROOT` variable automatically. Create a workspace directory for `GOPATH` if you need to override the default (`~/go`):

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## Initialize Go Modules

Each Go package under `libs/` uses its own `go.mod` file. A `go.work` file at the repository root ties these modules together for local development:

```bash
go work init ./libs/sdk-go
```

Additional modules will be added to `go.work` as they are created. This allows imports between modules without requiring `replace` directives.

## Next Steps

With the toolchain configured, future steps will implement Go equivalents of the existing libraries and ensure tests run through `make test` in each Go module.
