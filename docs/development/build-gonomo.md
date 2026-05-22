# Build gonomo.exe From Source

The Go project lives in `gonomo/`. The built binary is the backend engine that the npm CLI downloads from GitHub releases.

## Prerequisites

- Go
- `rsrc` for Windows icon resources
- PowerShell on Windows

Install `rsrc`:

```powershell
go install github.com/akavel/rsrc@latest
```

## Build

From the repository root:

```powershell
powershell -ExecutionPolicy Bypass -File ./gonomo/scripts/build-gonomo.ps1
```

This builds `gonomo.exe` at the repo root by default.

To choose an output path:

```powershell
powershell -ExecutionPolicy Bypass -File ./gonomo/scripts/build-gonomo.ps1 -Output ./dist/gonomo.exe
```

## Icon

The source icon is `resources/gonomo.png`. The build script converts it to `resources/gonomo.ico`, then uses `rsrc` to produce `gonomo/cmd/gonomo/rsrc.syso`.

Generated files are ignored:

- `resources/gonomo.ico`
- `gonomo/cmd/gonomo/rsrc.syso`
- `gonomo.exe`

## Manual Go Commands

If you do not need the icon resource, you can build the Go engine directly:

```powershell
go build -C ./gonomo ./...
```

To run `go generate` for the icon resources:

```powershell
go generate -C ./gonomo ./cmd/gonomo
```
