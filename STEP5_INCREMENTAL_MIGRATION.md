# Step 5: Incremental Migration

Building upon the analysis in [STEP1_ANALYSIS.md](STEP1_ANALYSIS.md), the
design described in [STEP2_DESIGN.md](STEP2_DESIGN.md), the toolchain setup in
[STEP3_TOOLCHAIN.md](STEP3_TOOLCHAIN.md), and the feature mapping in
[STEP4_FEATURE_PARITY.md](STEP4_FEATURE_PARITY.md), this step outlines how to
rewrite the repository one module at a time.

## Approach

1. **Start with Core Libraries**
   - Migrate `langgraph` to `langgraph-go` first so other packages have a solid
     foundation.
   - Port the checkpoint interfaces next (`checkpoint-go`), followed by the
     Postgres and SQLite implementations.
2. **Integrate Gradually**
   - Introduce Go modules alongside existing Python code using interface layers
     or small wrappers.
   - Keep tests running for both languages during the transition.
3. **Update Downstream Packages**
   - Once a core library is ready, update dependents (e.g., `cli-go`,
     `sdk-go`, `prebuilt-go`) to use the new Go code.
   - Maintain compatibility with existing tooling until all modules are
     migrated.
4. **Remove Legacy Code**
   - After verifying each module, deprecate and delete the corresponding
     non-Go implementation.
   - Clean up documentation and examples to reference the Go versions.

This incremental strategy lets the Go code evolve in place while existing
packages continue to function, reducing the risk of breaking changes during the
migration.
