# Configuration

Gonomo uses `gonomo.config.ts` with `defineConfig()` from `gonomo`.

```ts
import { defineConfig } from 'gonomo'

export default defineConfig({
  name: 'MyApp',
  title: 'My App',
  icon: './icon.ico',
  window: {
    width: 1200,
    height: 800,
    maximized: false,
    titleBarStyle: {
      hidden: true,
      overlay: true,
      darkMode: true,
      captionColor: '#202427',
      textColor: '#cdd6f4',
      captionButtonWidth: 46,
      captionButtonHeight: 32,
    },
  },
  splash: {
    enabled: true,
    layout: 'centered',
    backgroundColor: '#fafafa',
    foregroundColor: '#202427',
    text: 'My App',
    minDuration: 1200,
    width: 460,
    height: 300,
  },
  build: {
    command: 'pnpm run build',
    cwd: '.',
    outputDir: '.output',
    entry: 'server/index.mjs',
    runtime: 'node',
    embed: 'full',
  },
  dev: {
    command: 'pnpm run dev',
    cwd: '.',
    url: 'http://localhost:3000',
  },
  output: {
    dir: './dist',
    name: 'MyApp.exe',
  },
})
```

## Top-Level Fields

| Field | Type | Description |
|---|---|---|
| `name` | `string` | Internal app name and default title/output base |
| `title` | `string` | Window title; defaults to `name` |
| `icon` | `string` | Path to the user app `.ico` file |
| `window` | `object` | Main window options |
| `splash` | `object` | Native splash screen options |
| `build` | `object` | Frontend build and runtime packaging options |
| `dev` | `object` | Dev server command and URL |
| `output` | `object` | Final `.exe` output path |

## Build Runtime

Use `runtime: 'node'` for SSR frameworks like Nuxt and Next standalone output. Use `runtime: 'static'` for static Vite-style apps.
