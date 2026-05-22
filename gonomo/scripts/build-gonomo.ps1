param(
    [string]$Output = ""
)

$ErrorActionPreference = "Stop"

$goRoot = Resolve-Path (Join-Path $PSScriptRoot "..")
$repoRoot = Resolve-Path (Join-Path $goRoot "..")
$pngPath = Join-Path $repoRoot "resources\gonomo.png"
$iconPath = Join-Path $repoRoot "resources\gonomo.ico"
$sysoPath = Join-Path $goRoot "cmd\gonomo\rsrc.syso"

if ($Output -eq "") {
    $Output = Join-Path $repoRoot "gonomo.exe"
}

Push-Location $goRoot
try {
    if (-not (Test-Path -LiteralPath $pngPath)) {
        throw "resources/gonomo.png is required for the gonomo.exe icon"
    }

    if ((-not (Test-Path -LiteralPath $iconPath)) -or ((Get-Item -LiteralPath $pngPath).LastWriteTime -gt (Get-Item -LiteralPath $iconPath).LastWriteTime)) {
        & go run .\tools\generateicon\main.go -src $pngPath -out $iconPath
        if ($LASTEXITCODE -ne 0) { throw "failed to generate resources/gonomo.ico" }
    }

    $rsrc = Get-Command rsrc -ErrorAction SilentlyContinue
    if ($null -eq $rsrc) {
        throw "rsrc is required to embed the gonomo.exe icon. Install it with: go install github.com/akavel/rsrc@latest"
    }

    & $rsrc.Source -ico $iconPath -o $sysoPath
    if ($LASTEXITCODE -ne 0) { throw "rsrc failed to generate gonomo/cmd/gonomo/rsrc.syso" }

    & go build -o $Output .\cmd\gonomo
    if ($LASTEXITCODE -ne 0) { throw "go build failed" }
}
finally {
    Pop-Location
}
