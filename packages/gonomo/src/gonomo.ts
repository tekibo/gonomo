import type { GonomoAPI, GonomoUserConfig } from './types.js'

/** Initialize the gonomo bridge: set CSS vars and dismiss splash. Call once on mount. */
export function gonomoInit(): GonomoAPI | null {
  const api = (globalThis as any).gonomo as GonomoAPI | undefined
  if (!api) return null

  const r = document.documentElement
  r.style.setProperty('--gonomo-caption-button-height', api.captionButtonHeight + 'px')
  r.style.setProperty('--gonomo-caption-buttons-width', api.captionButtonsWidth + 'px')
  r.style.setProperty('--gonomo-resize-border-top', api.resizeBorderTop + 'px')

  if (api.dismissSplash) api.dismissSplash()

  return api
}

// --- Named API helpers ---

export function gonomoClose(): void {
  ;(globalThis as any).gonomo?.Close()
}

export function gonomoMinimize(): void {
  ;(globalThis as any).gonomo?.Minimize()
}

export function gonomoMaximize(): void {
  ;(globalThis as any).gonomo?.Maximize()
}

export function gonomoRestore(): void {
  ;(globalThis as any).gonomo?.Restore()
}

export function gonomoIsMaximized(): boolean {
  return !!(globalThis as any).gonomo?.IsMaximized()
}

export function gonomoSetDarkMode(enabled: boolean): void {
  ;(globalThis as any).gonomo?.setDarkMode(enabled)
}

export function gonomoSetCaptionColor(hex: string): void {
  ;(globalThis as any).gonomo?.setCaptionColor(hex)
}

export function gonomoSetTextColor(hex: string): void {
  ;(globalThis as any).gonomo?.setTextColor(hex)
}

export function gonomoSetTitlebarVisible(visible: boolean): void {
  ;(globalThis as any).gonomo?.setTitlebarVisible(visible)
}

export function gonomoSetTitleBarOverlay(enabled: boolean): void {
  ;(globalThis as any).gonomo?.setTitleBarOverlay(enabled)
}

// --- Config helper ---

/** Create a fully typed gonomo config with autocompletion. Use as the default export of gonomo.config.ts */
export function defineConfig(config: GonomoUserConfig): GonomoUserConfig {
  return config
}
