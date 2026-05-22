import { join } from 'node:path'
import { spawn } from 'node:child_process'
import { existsSync } from 'node:fs'
import { loadUserConfig, writeConfigJson } from './config.js'

async function findGonomoBin(): Promise<string> {
  const binDir = join(process.cwd(), '.gonomo', 'bin')
  if (!existsSync(binDir)) {
    throw new Error('No .gonomo directory found. Run `npx gonomo init` first.')
  }
  const { readdir } = await import('node:fs/promises')
  const files = await readdir(binDir)
  const bin = files.find(f => f.startsWith('gonomo'))
  if (!bin) throw new Error('No gonomo binary found in .gonomo/bin')
  return join(binDir, bin)
}

export async function runDev(args: string[]) {
  // Load typed config and write temp gonomo.json for the Go binary
  const config = await loadUserConfig(process.cwd())
  await writeConfigJson(config, join(process.cwd(), '.gonomo', 'gonomo.json'))

  const bin = await findGonomoBin()
  console.log(`Starting dev mode...`)

  const child = spawn(bin, ['dev', ...args], {
    stdio: 'inherit',
    cwd: process.cwd(),
  })

  child.on('exit', (code) => process.exit(code ?? 1))
}
