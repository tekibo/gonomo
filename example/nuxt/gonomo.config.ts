import { defineConfig } from 'gonomo'

export default defineConfig({
  name: 'web',
  title: 'web',
  icon: 'icon.png',
  window: {
    width: 1400,
    height: 900,
    titleBarStyle: {
      hidden: true,
      overlay: true,
      darkMode: true,
      captionColor: '#202427',
      textColor: '#cdd6f4',
    },
  },
  splash: {
    enabled: true,
    layout: "custom",
    image: "icon.png",
    minDuration: 2000,
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
    name: 'web.exe',
  },
})
