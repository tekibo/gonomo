import { readFile, writeFile, mkdir } from 'node:fs/promises'
import { existsSync } from 'node:fs'
import { join } from 'node:path'
import { execSync } from 'node:child_process'
import { downloadGonomo } from './download.js'

const configTsTemplate = `import { defineConfig } from 'gonomo'

export default defineConfig({
  name: '{{name}}',
  title: '{{name}}',
  icon: 'icon.ico',
  window: {
    width: 1400,
    height: 900,
    maximized: true,
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
    layout: 'centered',
    backgroundColor: '#fafafa',
    foregroundColor: '#202427',
    text: '{{name}}',
    minDuration: 2000,
    width: 460,
    height: 300,
  },
  build: {
    command: '{{buildCmd}}',
    cwd: '.',
    outputDir: '{{outputDir}}',
    entry: '{{entry}}',
    runtime: '{{runtime}}',
    embed: 'full',
  },
  dev: {
    command: '{{devCmd}}',
    cwd: '.',
    url: 'http://localhost:3000',
  },
  output: {
    dir: './dist',
    name: '{{name}}.exe',
  },
})
`

function detectPackageManager(): string {
  if (existsSync('pnpm-lock.yaml')) return 'pnpm'
  if (existsSync('yarn.lock')) return 'yarn'
  if (existsSync('bun.lockb')) return 'bun'
  return 'npm'
}

function addPackage(pm: string, name: string, dev = false): void {
  const devFlag = pm === 'bun' ? '-d' : '-D'
  execSync(`${pm} add ${dev ? `${devFlag} ` : ''}${name}`, { stdio: 'inherit' })
}

export async function runInit(args: string[]) {
  console.log('Initializing gonomo project...')

  const pm = detectPackageManager()
  const cmdPrefix = `${pm} run `

  let name = 'my-app'
  let runtime = 'node'
  let outputDir = '.output'
  let entry = 'server/index.mjs'

  // Read package.json for name and framework detection
  const pkgPath = join(process.cwd(), 'package.json')
  if (existsSync(pkgPath)) {
    const pkg = JSON.parse(await readFile(pkgPath, 'utf-8'))
    if (pkg.name) name = pkg.name

    const deps: string[] = []
    if (pkg.dependencies) deps.push(...Object.keys(pkg.dependencies))
    if (pkg.devDependencies) deps.push(...Object.keys(pkg.devDependencies))
    const depSet = new Set(deps)

    if (depSet.has('nuxt') || depSet.has('nuxt3')) {
      console.log('Detected Nuxt')
      outputDir = '.output'
      entry = 'server/index.mjs'
    } else if (depSet.has('next')) {
      console.log('Detected Next.js')
      outputDir = '.next/standalone'
      entry = 'server.js'
    } else if (depSet.has('vite') || depSet.has('@sveltejs/kit')) {
      console.log('Detected static frontend')
      runtime = 'static'
      outputDir = 'dist'
      entry = ''
    }
  }

  // Create .gonomo directory
  const gonomoDir = join(process.cwd(), '.gonomo')
  await mkdir(gonomoDir, { recursive: true })

  // Download gonomo binary
  const binPath = await downloadGonomo(gonomoDir)
  console.log(`Gonomo binary ready at ${binPath}`)

  // Write gonomo.config.ts
  const configContent = configTsTemplate
    .replace(/{{name}}/g, name)
    .replace('{{buildCmd}}', `${cmdPrefix}build`)
    .replace('{{devCmd}}', `${cmdPrefix}dev`)
    .replace('{{outputDir}}', outputDir)
    .replace('{{entry}}', entry)
    .replace('{{runtime}}', runtime)

  await writeFile('gonomo.config.ts', configContent)
  console.log('Created gonomo.config.ts')

  // Add runtime and local CLI dependencies
  console.log('Adding gonomo packages...')
  addPackage(pm, 'gonomo', true)

  // Update package.json scripts
  if (existsSync(pkgPath)) {
    const pkg = JSON.parse(await readFile(pkgPath, 'utf-8'))
    pkg.scripts ??= {}
    if (!pkg.scripts['gonomo:dev']) pkg.scripts['gonomo:dev'] = 'gonomo dev'
    if (!pkg.scripts['gonomo:build']) pkg.scripts['gonomo:build'] = 'gonomo build'
    await writeFile(pkgPath, JSON.stringify(pkg, null, 2) + '\n')
    console.log('Updated package.json with gonomo scripts')
  }

  console.log(`\nDone! Run \`${pm} run gonomo:dev\` to start developing.`)
}
