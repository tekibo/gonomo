# Testing Locally

When developing the Gonomo engine or the NPM CLI, you will want to test them together locally without publishing.

## Using a Local `gonomo.exe` Engine

The npm CLI (`gonomo`) delegates all `dev` and `build` commands to the downloaded Go binary. By default, `gonomo init` downloads the latest binary from GitHub into a `.gonomo/bin` directory inside your app workspace.

To test a locally built engine:

1. Build `gonomo.exe` from source (see [Build gonomo.exe From Source](./build-gonomo.md)).
2. Create an example project or use the `example/` directory.
3. Copy your locally built `gonomo.exe` into the example project's `.gonomo/bin/` directory, overwriting the downloaded one. Make sure it's named according to your platform (e.g. `gonomo-x64.exe` on Windows).
4. Run your application. The CLI will now use your local binary.

## Testing the Local NPM Packages

The TypeScript package (`gonomo`) can be built by running:

```powershell
pnpm run --filter gonomo build
```

To test the `init` command locally inside the `example` folder, or any other directory, you can run the CLI script directly using Node:

```powershell
cd example
node ../../packages/gonomo/dist/cli.js init
```

This will run the initialization process using the local, just-built `init` logic. Note that `init` will add the package to `package.json`, and since the example folder is inside a PNPM workspace, it will correctly link `gonomo` via the workspace protocol.

After initialization, you can run:

```powershell
pnpm run gonomo:dev
```

And it will execute your local copy of the CLI (and whichever binary is inside `.gonomo/bin/`).
