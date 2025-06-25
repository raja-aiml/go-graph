# Step 2: Design the Target Go Architecture

Following the analysis in [STEP1_ANALYSIS.md](STEP1_ANALYSIS.md), we outline a high-level architecture for migrating the repository to Go.

## Repository Layout

- Continue using a **monorepo** so Go modules live alongside existing Python and JS packages under `libs/`.
- Each Go library resides in its own folder with a `go.mod` file. The proposed layout mirrors the current libraries:
  - `libs/langgraph-go` – core graph framework.
  - `libs/checkpoint-go` – checkpoint interfaces and shared types.
  - `libs/checkpoint-postgres-go` – Postgres implementation of the checkpoint saver.
  - `libs/checkpoint-sqlite-go` – SQLite implementation of the checkpoint saver.
  - `libs/cli-go` – command‑line utilities.
  - `libs/prebuilt-go` – higher‑level agent APIs.
  - `libs/sdk-go` – Go SDK for interacting with the platform API (already started).

## Modules and Packages

- Each module should be self-contained but importable by others, similar to the Python packages.
- Use Go modules for dependency management. Downstream packages declare requirements on the core libraries, following the dependency map from the analysis:

```
langgraph-go
├── checkpoint-go
│   ├── checkpoint-postgres-go
│   └── checkpoint-sqlite-go
├── cli-go
└── prebuilt-go

sdk-go
└── langgraph-go
```

## Goals Inherited from Step 1

The design must provide feature parity with the functionality identified in Step 1:

1. **Graph Framework** – `langgraph-go` exposes types for building and running stateful graphs.
2. **Checkpointing** – the checkpoint modules persist state using Go database drivers.
3. **CLI Utilities** – `cli-go` wraps common project commands and development tools.
4. **High-level Agent APIs** – `prebuilt-go` mirrors agent patterns like ReAct.
5. **SDKs** – `sdk-go` handles authentication and API interaction using the standard library.
6. **Testing and Examples** – each module includes Go tests and example programs.

Keeping this structure lets us gradually rewrite components while sharing conventions across languages.
