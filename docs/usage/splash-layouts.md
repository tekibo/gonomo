# Splash Layouts

The splash screen is a native Win32 popup displayed while the runtime starts.

```ts
export default defineConfig({
  splash: {
    enabled: true,
    layout: 'centered',
    backgroundColor: '#fafafa',
    foregroundColor: '#202427',
    image: './splash.png',
    text: 'My App',
    tagline: 'Starting...',
    minDuration: 1200,
    width: 460,
    height: 300,
  },
})
```

## Layouts

| Layout | Description |
|---|---|
| `centered` | Image above centered text |
| `minimal` | Text-only splash |
| `top-banner` | Image banner on top, text below |
| `bottom-banner` | Text on top, image banner below |
| `split` | Image left, text right |
| `full-image` | Image fills splash with overlaid text |
| `custom` | Transparent splash showing only the image |

## Notes

Use PNG for splash images when you need transparency. `minDuration` is in milliseconds and prevents the splash from closing too quickly on fast starts.
