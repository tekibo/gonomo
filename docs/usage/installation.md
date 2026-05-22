# Installation Guide

To protect against npm package poisoning and maintain a tighter supply chain, Gonomo intentionally avoids publishing to the NPM registry. Instead, you consume the official TypeScript packages directly from the GitHub Releases.

This gives you a secure, per-project installation (similar to Expo CLI) without relying on global installs or the NPM registry.

## 1. Install Dependencies from GitHub

In your project directory, use your package manager to install the packed tarballs directly from the latest GitHub Release URL.

**Using npm:**
```bash
npm install -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo-cli.tgz https://github.com/tekibo/gonomo/releases/latest/download/gonomo-core.tgz
```

**Using pnpm:**
```bash
pnpm add -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo-cli.tgz https://github.com/tekibo/gonomo/releases/latest/download/gonomo-core.tgz
```

**Using bun:**
```bash
bun add -d https://github.com/tekibo/gonomo/releases/latest/download/gonomo-cli.tgz https://github.com/tekibo/gonomo/releases/latest/download/gonomo-core.tgz
```

## 2. Initialize Gonomo

Once the CLI and Core packages are installed as `devDependencies`, you can run the CLI locally within your project using `npx`, `pnpm dlx`, or `bunx`.

```bash
npx gonomo init
```

This command will:
1. Generate your `gonomo.config.ts` configuration file.
2. Download the compiled Go engine binary (`gonomo.exe`, etc.) specifically tailored for your OS into the `.gonomo/bin/` folder.
3. Automatically inject `gonomo:dev` and `gonomo:build` scripts into your `package.json`.

## 3. Run Commands

Now you can use Gonomo safely using the standard package manager scripts, knowing everything is isolated to your project:

```bash
npm run gonomo:dev
```
```bash
npm run gonomo:build
```
