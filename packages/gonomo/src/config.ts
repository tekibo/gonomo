import { pathToFileURL } from 'node:url'
import { join } from 'node:path'
import { existsSync } from 'node:fs'
import { writeFile } from 'node:fs/promises'
import type { GonomoUserConfig } from './types.js'

/**
 * Load the user's gonomo.config.ts and return the config object.
 * Uses tsx internally to handle TypeScript at runtime.
 */
export async function loadUserConfig(cwd: string): Promise<GonomoUserConfig> {
  const configPath = join(cwd, 'gonomo.config.ts')
  if (!existsSync(configPath)) {
    throw new Error(
      'No gonomo.config.ts found. Run `npx gonomo init` to create one.'
    )
  }

  // tsx provides the loader needed to import .ts files at runtime
  const { register } = await import('tsx/esm/api')
  register()

  const imported = await import(pathToFileURL(configPath).href)
  return imported.default as GonomoUserConfig
}

/**
 * Serialize a GonomoUserConfig to the JSON format the Go binary expects
 * and write it as gonomo.json.
 */
export async function writeConfigJson(config: GonomoUserConfig, dest: string): Promise<void> {
  const json = {
    name: config.name,
    title: config.title ?? config.name,
    icon: config.icon ?? '',
    window: {
      width: config.window?.width ?? 1400,
      height: config.window?.height ?? 900,
      maximized: config.window?.maximized ?? false,
      titleBarStyle: config.window?.titleBarStyle ?? {
        hidden: true,
        overlay: true,
        darkMode: true,
        captionColor: '#202427',
        textColor: '#cdd6f4',
      },
    },
    splash: {
      enabled: config.splash?.enabled ?? false,
      layout: config.splash?.layout ?? 'centered',
      backgroundColor: config.splash?.backgroundColor ?? '#fafafa',
      foregroundColor: config.splash?.foregroundColor ?? '#202427',
      image: config.splash?.image ?? '',
      text: config.splash?.text ?? config.name,
      tagline: config.splash?.tagline ?? '',
      minDuration: config.splash?.minDuration ?? 2000,
      width: config.splash?.width ?? 480,
      height: config.splash?.height ?? 320,
    },
    build: {
      command: config.build.command,
      cwd: config.build.cwd ?? '.',
      outputDir: config.build.outputDir ?? '.output',
      entry: config.build.entry ?? 'server/index.mjs',
      runtime: config.build.runtime ?? 'node',
      embed: config.build.embed ?? 'full',
      nodeVersion: config.build.nodeVersion ?? '',
    },
    dev: {
      command: config.dev.command,
      cwd: config.dev.cwd ?? '.',
      url: config.dev.url ?? 'http://localhost:3000',
    },
    output: {
      dir: config.output?.dir ?? './dist',
      name: config.output?.name ?? `${config.name}.exe`,
    },
  }

  await writeFile(dest, JSON.stringify(json, null, 2) + '\n')
}
