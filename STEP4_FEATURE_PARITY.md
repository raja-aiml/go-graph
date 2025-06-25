# Step 4: Plan Feature Parity

Building on the repository survey in [STEP1_ANALYSIS.md](STEP1_ANALYSIS.md), the architecture defined in [STEP2_DESIGN.md](STEP2_DESIGN.md), and the toolchain preparation from [STEP3_TOOLCHAIN.md](STEP3_TOOLCHAIN.md), this step maps each existing library to its Go counterpart.

## Library Mapping

- **checkpoint** → `checkpoint-go`
  - Go interfaces for checkpoint savers.
  - Implement Postgres and SQLite modules (`checkpoint-postgres-go`, `checkpoint-sqlite-go`) that satisfy these interfaces.
- **cli** → `cli-go`
  - Recreate command-line utilities using the `cobra` framework.
  - Provide project setup and development commands.
- **langgraph** → `langgraph-go`
  - Port graph building, execution, and state management types.
- **prebuilt** → `prebuilt-go`
  - Offer high-level agent patterns such as ReAct in Go.
- **sdk-py** / **sdk-js** → `sdk-go`
  - Implement authentication, API calls, and streaming responses.

## Replacing Third-Party Dependencies

Python dependencies identified in Step&nbsp;1 lack direct Go equivalents. The migration will replace them with native packages or popular Go libraries:

| Python Dependency | Planned Go Replacement |
| ----------------- | --------------------- |
| `httpx`           | Go `net/http` |
| `orjson`          | Go `encoding/json` |
| `psycopg`         | `database/sql` + `github.com/lib/pq` |
| `aiosqlite`/`sqlite-vec` | `github.com/mattn/go-sqlite3` |
| `click`           | `github.com/spf13/cobra` |
| `xxhash`          | `github.com/cespare/xxhash/v2` |
| `pydantic`        | Struct tags + validation libraries |

Each Go module will maintain its own `go.mod` file and participate in the repository-level `go.work` defined in Step&nbsp;3. This ensures cross-module imports work seamlessly during the migration.
