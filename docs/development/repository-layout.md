# Repository Layout

The repo is organized as a small monorepo:

| Path | Purpose |
|---|---|
| `gonomo/` | Go build engine and generated-app runtime template |
| `packages/core/` | Framework-agnostic frontend runtime API and config types |
| `packages/cli/` | npm CLI published as `gonomo` |
| `resources/gonomo.png` | Source icon for the `gonomo.exe` framework binary |
| `example/` | Example apps used while developing Gonomo |
| `docs/development/` | Contributor documentation |
| `docs/usage/` | End-user documentation |

The Go engine is intentionally contained in `gonomo/` so the root stays focused on docs, packages, examples, and repo-level assets.
