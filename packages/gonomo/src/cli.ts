#!/usr/bin/env node

async function main() {
  const command = process.argv[2]
  const args = process.argv.slice(3)

  switch (command) {
    case 'init': {
      const { runInit } = await import('./init.js')
      await runInit(args)
      break
    }
    case 'dev': {
      const { runDev } = await import('./dev.js')
      await runDev(args)
      break
    }
    case 'build': {
      const { runBuild } = await import('./build.js')
      await runBuild(args)
      break
    }
    case 'version':
      console.log('gonomo v0.0.0')
      break
    default:
      console.log(`
gonomo — Desktop App Builder

Usage:
  gonomo <command> [arguments]

Commands:
  init       Bootstrap a new project
  dev        Start dev mode
  build      Build the final executable
  version    Print version
`)
      process.exit(command ? 1 : 0)
  }
}

main().catch((err) => {
  console.error(err)
  process.exit(1)
})
