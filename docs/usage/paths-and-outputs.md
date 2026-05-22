# Paths and Outputs

## Files Created By Init

| Path | Purpose |
|---|---|
| `gonomo.config.ts` | Type-safe user config |
| `.gonomo/bin/` | Downloaded Go engine binary |
| `package.json` scripts | `gonomo:dev` and `gonomo:build` |

## Files Created By Build

| Path | Purpose |
|---|---|
| `gonomo.json` | Temporary JSON serialized from `gonomo.config.ts` for the Go engine |
| `.gonomo/build/` | Generated Go runtime project |
| `.gonomo/bin/node-*.exe` | Downloaded Node runtime when `build.runtime` is `node` |
| `.gonomo/build/internal/resources/` | Embedded frontend, Node runtime, and user app icon |
| `output.dir/output.name` | Final Windows `.exe` |

## Icons

The user app icon comes from `icon` in `gonomo.config.ts`. It is copied into `.gonomo/build/internal/resources/icon.ico` and compiled into the final `.exe`.

The framework icon for `gonomo.exe` is not used by user apps.
