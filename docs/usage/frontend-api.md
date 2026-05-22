# Frontend API

Install the frontend library:

```bash
npm install gonomo
```

Initialize once when your app mounts:

```ts
import { gonomoInit } from 'gonomo'

gonomoInit()
```

React:

```tsx
useEffect(() => {
  gonomoInit()
}, [])
```

Vue:

```ts
onMounted(() => {
  gonomoInit()
})
```

## Helpers

| Helper | Description |
|---|---|
| `gonomoClose()` | Close the window |
| `gonomoMinimize()` | Minimize the window |
| `gonomoMaximize()` | Maximize the window |
| `gonomoRestore()` | Restore the window |
| `gonomoIsMaximized()` | Returns whether the window is maximized |
| `gonomoSetDarkMode(enabled)` | Toggle dark titlebar rendering |
| `gonomoSetCaptionColor(hex)` | Set native caption background color |
| `gonomoSetTextColor(hex)` | Set native caption text/icon color |
| `gonomoSetTitlebarVisible(visible)` | Show/hide the titlebar |
| `gonomoSetTitleBarOverlay(enabled)` | Toggle titlebar overlay |

`gonomoInit()` also writes CSS variables used for titlebar layout:

| Variable | Description |
|---|---|
| `--gonomo-caption-button-height` | Native caption button height |
| `--gonomo-caption-buttons-width` | Total width of minimize/maximize/close buttons |
| `--gonomo-resize-border-top` | Top resize border inset |
