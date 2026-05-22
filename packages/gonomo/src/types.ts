// --- Runtime API types (injected by Go binary into window.gonomo) ---

export interface GonomoConfig {
  title: string
  icon: string
  width: number
  height: number
  titleBarStyle?: 'hidden'
  titleBarOverlay?: boolean
  hideTitlebar?: boolean
  titlebar?: GonomoTitlebarConfig
  splash?: GonomoSplashConfig
}

export interface GonomoTitlebarConfig {
  mode?: 'normal' | 'hidden' | 'onlyNativeButtons'
  darkMode: boolean
  captionColor: string
  textColor: string
  titleBarOverlay?: boolean
  captionButtonWidth?: number
  captionButtonHeight?: number
}

export interface GonomoSplashConfig {
  enabled: boolean
  layout: 'centered' | 'minimal' | 'top-banner' | 'bottom-banner' | 'split' | 'full-image' | 'custom'
  backgroundColor: string
  foregroundColor: string
  image: string
  text: string
  tagline?: string
  minDuration: number
  width: number
  height: number
}

export interface GonomoAPI {
  Config: GonomoConfig
  Titlebar: GonomoTitlebarConfig
  captionButtonHeight: number
  captionButtonsWidth: number
  resizeBorderTop: number
  Close: () => void
  Minimize: () => void
  Maximize: () => void
  Restore: () => void
  IsMaximized: () => boolean
  setDarkMode: (enabled: boolean) => void
  setCaptionColor: (hex: string) => void
  setTextColor: (hex: string) => void
  setTitlebarVisible: (visible: boolean) => void
  setTitleBarOverlay: (enabled: boolean) => void
  dismissSplash: () => void
}

declare global {
  interface Window {
    gonomo?: GonomoAPI
  }
}

// --- User-facing config types (for gonomo.config.ts) ---

export interface GonomoWindowConfig {
  width?: number
  height?: number
  maximized?: boolean
  titleBarStyle?: {
    hidden?: boolean
    overlay?: boolean
    darkMode?: boolean
    captionColor?: string
    textColor?: string
    captionButtonWidth?: number
    captionButtonHeight?: number
  }
}

export interface GonomoBuildConfig {
  command: string
  cwd?: string
  outputDir?: string
  entry?: string
  runtime?: 'node' | 'static'
  embed?: 'full' | 'none'
  nodeVersion?: string
}

export interface GonomoDevConfig {
  command: string
  cwd?: string
  url?: string
}

export interface GonomoOutputConfig {
  dir?: string
  name?: string
}

export interface GonomoSplashUserConfig {
  enabled?: boolean
  layout?: 'centered' | 'minimal' | 'top-banner' | 'bottom-banner' | 'split' | 'full-image' | 'custom'
  backgroundColor?: string
  foregroundColor?: string
  image?: string
  text?: string
  tagline?: string
  minDuration?: number
  width?: number
  height?: number
}

export interface GonomoUserConfig {
  name: string
  title?: string
  icon?: string
  window?: GonomoWindowConfig
  build: GonomoBuildConfig
  dev: GonomoDevConfig
  output?: GonomoOutputConfig
  splash?: GonomoSplashUserConfig
}
