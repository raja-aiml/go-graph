# Step 1: Analyze the Existing Codebase

This document catalogs the current repository layout and major dependencies. It also highlights the key features that would need to be re-implemented in Go.

## Repository Structure

The repository is organized as a monorepo under `libs/` where each subdirectory is a standalone package.

- `checkpoint` – base interfaces for checkpoint savers.
- `checkpoint-postgres` – Postgres implementation of checkpoint persistence.
- `checkpoint-sqlite` – SQLite implementation of checkpoint persistence.
- `cli` – command-line interface for managing LangGraph projects.
- `langgraph` – core Python framework for building stateful, multi-actor agents.
- `prebuilt` – high-level agent APIs built on top of `langgraph`.
- `sdk-js` – JavaScript/TypeScript SDK for interacting with the LangGraph REST API.
- `sdk-py` – Python SDK for interacting with the platform API.
- `sdk-go` – early Go SDK containing a simple `Hello` function and basic auth types.

Additional folders include documentation (`docs/`) and example notebooks (`examples/`).

## Key Dependencies

Across the Python libraries, common dependencies include:

- `langchain-core` for shared interfaces
- `httpx` and `orjson` in the SDK
- `psycopg` for Postgres checkpointers
- `aiosqlite` and `sqlite-vec` for SQLite support
- `click` for the CLI
- `xxhash` and `pydantic` within the core framework

The JavaScript SDK relies on TypeScript, React, and various testing/formatting tools. The Go SDK currently has no external dependencies beyond the standard library.

## Features to Re‑implement in Go

To fully migrate the codebase to Go, the following functionality would require Go implementations:

1. **Graph Framework** – the `langgraph` library provides graph building, execution, and state management utilities. A Go equivalent would need to support similar abstractions (channels, state objects, runner loops, etc.).
2. **Checkpointing** – persistence layers (currently Postgres and SQLite implementations) must be recreated in Go, including the `CheckpointSaver` interfaces and their database integrations.
3. **CLI Utilities** – commands for creating and running projects, plus development helpers such as hot reloading.
4. **High‑level Agent APIs** – the `prebuilt` package bundles reusable agent patterns like ReAct agents.
5. **SDKs** – client libraries for interacting with the LangGraph API. The Go SDK would need parity with the Python and JS SDKs, handling authentication and streaming responses.
6. **Testing and Examples** – a comprehensive test suite and example applications mirroring the current Python notebooks and scripts.

Third‑party modules such as `aiosqlite`, `psycopg`, and `httpx` do not have direct Go equivalents, so native Go libraries or community packages will be required.

