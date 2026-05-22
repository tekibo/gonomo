# Commands

## Initialize

First, install the framework packages directly from the latest GitHub Release:

```bash
npm install -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo-cli.tgz https://github.com/tekibo/gonomo/releases/latest/download/gonomo-core.tgz
```

Then, initialize your project:

```bash
npx gonomo init
```

Creates `gonomo.config.ts`, downloads the Go engine into `.gonomo/bin`, configures your framework, and injects package scripts.

## Dev Mode

```bash
npm run gonomo:dev
```

Runs your configured dev server and opens it in a native WebView2 window.

## Build

```bash
npm run gonomo:build
```

Runs your frontend build, downloads `node.exe` when `build.runtime` is `node`, generates the temporary Go project in `.gonomo/build`, and compiles the final Windows `.exe`.

## Build Flags

```bash
gonomo build --skip-frontend
gonomo build --clean
gonomo build --verbose
```

| Flag | Description |
|---|---|
| `--skip-frontend` | Skip the configured frontend build command |
| `--clean` | Remove `.gonomo/build` before generating the runtime project |
| `--verbose` | Reserved for detailed output |
