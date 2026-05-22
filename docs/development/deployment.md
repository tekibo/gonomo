# Deployment Flow

Releases for Gonomo are published entirely to one place for better supply chain security:

1. **GitHub Releases:** The compiled `gonomo.exe` binaries for all platforms, as well as the NPM packages packaged as `.tgz` tarballs.

Automated GitHub Actions workflows are configured in `.github/workflows` to fully automate the process.

## The Intended Flow

Once CI/CD is configured, the intended deployment flow is:

1. **Tagging:** A developer pushes a new git tag (e.g., `v1.0.0`).
2. **Binaries Build:** The GitHub Actions intercept the tag and:
   - Compile the Go engine for Windows, macOS (Intel & ARM), and Linux.
   - Build the TypeScript packages (`packages/cli` and `packages/core`) and pack them into `.tgz` tarballs.
   - Create a new GitHub Release with the tag.
   - Upload the compiled binaries and the package tarballs to the release assets.

## Global Constants

Constants that impact deployment across both Go and TypeScript environments are kept in a single source of truth at the root of the repository:

`global.toml`

When building the TypeScript `gonomo` package, a prebuild script automatically extracts values (like the GitHub repo name) from `global.toml` into a JS constant, ensuring that the CLI downloads binaries from the correct repository location.
