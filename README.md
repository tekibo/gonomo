<div align="center">
  <pre>
┌──────────────────────────────────────────────────────────────────────────────┐
│ ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄        ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄       ▄▄  ▄▄▄▄▄▄▄▄▄▄▄ │
│▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌      ▐░▌▐░░░░░░░░░░░▌▐░░▌     ▐░░▌▐░░░░░░░░░░░▌│
│▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░▌░▌     ▐░▌▐░█▀▀▀▀▀▀▀█░▌▐░▌░▌   ▐░▐░▌▐░█▀▀▀▀▀▀▀█░▌│
│▐░▌          ▐░▌       ▐░▌▐░▌▐░▌    ▐░▌▐░▌       ▐░▌▐░▌▐░▌ ▐░▌▐░▌▐░▌       ▐░▌│
│▐░▌ ▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░▌ ▐░▌   ▐░▌▐░▌       ▐░▌▐░▌ ▐░▐░▌ ▐░▌▐░▌       ▐░▌│
│▐░▌▐░░░░░░░░▌▐░▌       ▐░▌▐░▌  ▐░▌  ▐░▌▐░▌       ▐░▌▐░▌  ▐░▌  ▐░▌▐░▌       ▐░▌│
│▐░▌ ▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░▌   ▐░▌ ▐░▌▐░▌       ▐░▌▐░▌   ▀   ▐░▌▐░▌       ▐░▌│
│▐░▌       ▐░▌▐░▌       ▐░▌▐░▌    ▐░▌▐░▌▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌│
│▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌     ▐░▐░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄█░▌│
│▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌      ▐░░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌│
│ ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀        ▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀ │
└──────────────────────────────────────────────────────────────────────────────┘
  </pre>

  <h3>The lightweight desktop framework.</h3>
  <p>Build native desktop apps with web technologies using Go and WebView2.</p>
</div>

---

# 🚀 Quick Start

To maintain a tighter supply chain, Gonomo avoids publishing to the NPM registry.

You consume the official TypeScript packages directly from GitHub Releases.

<details open>
<summary><strong>npm</strong></summary>

<br />

### 1. Install from GitHub Releases

```bash
npm install -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo.tgz
```

### 2. Initialize your project

```bash
npx gonomo init
```

### 3. Run and build

```bash
npm run gonomo:dev
npm run gonomo:build
```

</details>

<details>
<summary><strong>pnpm</strong></summary>

<br />

### 1. Install from GitHub Releases

```bash
pnpm add -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo.tgz
```

### 2. Initialize your project

```bash
pnpm dlx gonomo init
```

### 3. Run and build

```bash
pnpm run gonomo:dev
pnpm run gonomo:build
```

</details>

<details>
<summary><strong>yarn</strong></summary>

<br />

### 1. Install from GitHub Releases

```bash
yarn add -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo.tgz
```

### 2. Initialize your project

```bash
yarn gonomo init
```

### 3. Run and build

```bash
yarn gonomo:dev
yarn gonomo:build
```

</details>

<details>
<summary><strong>bun</strong></summary>

<br />

### 1. Install from GitHub Releases

```bash
bun add -D https://github.com/tekibo/gonomo/releases/latest/download/gonomo.tgz
```

### 2. Initialize your project

```bash
bunx gonomo init
```

### 3. Run and build

```bash
bun run gonomo:dev
bun run gonomo:build
```

</details>

The final app does not require Node.js on the end user's machine.

For Node-backed apps, Gonomo downloads and embeds a standalone `node.exe` during build.

## ⚙️ Configuration

Gonomo uses configurations from `gonomo.config.ts`:

```ts
import { defineConfig } from 'gonomo'

export default defineConfig({
  name: 'MyApp',
  title: 'My App',
  icon: './icon.ico',
  build: {
    command: 'pnpm run build',
    outputDir: '.output',
    entry: 'server/index.mjs',
    runtime: 'node',
    embed: 'full',
  },
  dev: {
    command: 'pnpm run dev',
    url: 'http://localhost:3000',
  },
  output: {
    dir: './dist',
    name: 'MyApp.exe',
  },
  window: {
    width: 1200,
    height: 800,
    titleBarStyle: {
      hidden: true,
      overlay: true,
      darkMode: true,
    },
  },
  splash: {
    enabled: true,
    layout: 'centered',
    text: 'My App',
    backgroundColor: '#fafafa',
    foregroundColor: '#202427',
  },
})
```

## 🌐 Frontend Integration

Use `gonomo` in your frontend to initialize the bridge and access native window helpers:

```ts
import { gonomoInit, gonomoClose, gonomoMinimize } from 'gonomo'

// Must be called on mount to hide the native splash screen
gonomoInit()

// Native window controls
gonomoMinimize()
gonomoClose()
```

## 📚 Documentation

### Usage Guides

- [Installation Guide](./docs/usage/installation.md)
- [Configuration Reference](./docs/usage/configuration.md)
- [CLI Commands](./docs/usage/commands.md)
- [Paths and Outputs](./docs/usage/paths-and-outputs.md)
- [Frontend API](./docs/usage/frontend-api.md)
- [Titlebar Customization](./docs/usage/titlebar.md)
- [Splash Screen Layouts](./docs/usage/splash-layouts.md)

### Contributor Documentation

- [Development Home](./docs/development/README.md)
- [Testing Locally](./docs/development/testing-locally.md)
- [Deployment Flow](./docs/development/deployment.md)
- [Building the Go Engine](./docs/development/build-gonomo.md)
- [Building the TS Packages](./docs/development/typescript-packages.md)
