# Titlebar

Gonomo can hide the native Windows titlebar while preserving the native caption buttons over your web content.

```ts
export default defineConfig({
  window: {
    titleBarStyle: {
      hidden: true,
      overlay: true,
      darkMode: true,
      captionColor: '#202427',
      textColor: '#cdd6f4',
    },
  },
})
```

## CSS Variables

Call `gonomoInit()` once in your frontend. It injects:

| Variable | Description |
|---|---|
| `--gonomo-caption-button-height` | Height of the native caption buttons |
| `--gonomo-caption-buttons-width` | Total width reserved for native caption buttons |
| `--gonomo-resize-border-top` | Top resize border area |

Example:

```css
.titlebar {
  height: var(--gonomo-caption-button-height);
  padding-right: var(--gonomo-caption-buttons-width);
  padding-top: var(--gonomo-resize-border-top);
  -webkit-app-region: drag;
}

.titlebar button {
  -webkit-app-region: no-drag;
}
```

## Window Controls

Use `gonomo` helpers for custom controls:

```ts
import { gonomoClose, gonomoMaximize, gonomoMinimize, gonomoRestore } from 'gonomo'
```
