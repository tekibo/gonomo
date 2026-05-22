# Architecture

Gonomo has two layers:

| Layer | Implementation | Responsibility |
|---|---|---|
| User-facing CLI | TypeScript package `gonomo` | `init`, config loading, dependency/script injection, Go binary download |
| Build engine | Go project in `gonomo/` | Frontend build orchestration, runtime project generation, final `.exe` compilation |

## Runtime Generation

The Go engine reads `gonomo.json`, which is generated from `gonomo.config.ts` by the TypeScript CLI. It then creates `.gonomo/build`, copies the embedded runtime template into that directory, copies the configured user app icon and frontend output, embeds Node when required, and runs `go build` to produce the final app executable.

The embedded runtime template lives in `gonomo/internal/runtime` and uses module path `gonomo/runtime`.

## User App Icons

The user app icon is configured with `icon` in `gonomo.config.ts`. During build, that file is copied to `.gonomo/build/internal/resources/icon.ico` and also used to generate `.gonomo/build/rsrc.syso`, which gives the final user app `.exe` its Windows icon.

## Framework Icon

The `gonomo.exe` framework binary icon is separate. Its source asset is `resources/gonomo.png`; the build script converts it to ICO for `rsrc`.
