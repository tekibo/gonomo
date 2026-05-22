# Gonomo Nuxt Example

This example shows a Nuxt app configured for Gonomo with `gonomo.config.ts` and `gonomo`.

From the repository root:

```bash
pnpm install
pnpm run --filter web build
pnpm run --filter web gonomo:build
```

For dev mode:

```bash
pnpm run --filter web gonomo:dev
```

The example uses workspace dependencies for local development:

- `gonomo` from `packages/cli`
- `gonomo` from `packages/core`
