# Build TypeScript Package

The TypeScript workspace contains one package:

| Package | Path | Package Name |
|---|---|---|
| Gonomo | `packages/gonomo` | `gonomo` |

## Install

From the repository root:

```bash
pnpm install
```

## Build

```bash
pnpm run --filter gonomo build
```

## Package Responsibilities

`gonomo` is the main package. It serves two purposes:
1. It is the CLI that handles `init`, loads `gonomo.config.ts`, writes the temporary `gonomo.json` expected by the Go engine, and delegates `dev`/`build` to the downloaded Go binary.
2. It acts as the frontend library. It exports `defineConfig()`, runtime API helpers (like `gonomoInit()`), and TypeScript types. It has no framework dependency.
