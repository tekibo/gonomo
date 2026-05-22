export {
  gonomoInit,
  gonomoClose,
  gonomoMinimize,
  gonomoMaximize,
  gonomoRestore,
  gonomoIsMaximized,
  gonomoSetDarkMode,
  gonomoSetCaptionColor,
  gonomoSetTextColor,
  gonomoSetTitlebarVisible,
  gonomoSetTitleBarOverlay,
  defineConfig,
} from './gonomo.js'

export type {
  GonomoAPI,
  GonomoConfig,
  GonomoTitlebarConfig,
  GonomoSplashConfig,
  GonomoUserConfig,
  GonomoWindowConfig,
  GonomoBuildConfig,
  GonomoDevConfig,
  GonomoOutputConfig,
  GonomoSplashUserConfig,
} from './types.js'
