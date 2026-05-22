import { mkdir, chmod } from "node:fs/promises";
import { join } from "node:path";
import { platform, arch } from "node:os";
import { createWriteStream, existsSync } from "node:fs";
import { GONOMO_REPO } from "./constants.js";

function gonomoAssetName(): string {
  const p = platform();
  const a = arch();
  if (p === "win32") return `gonomo-x64.exe`;
  if (p === "darwin") return `gonomo-${a === "arm64" ? "arm64" : "x64"}`;
  return `gonomo-linux-${a === "arm64" ? "arm64" : "x64"}`;
}

async function downloadFile(url: string, dest: string): Promise<void> {
  const res = await fetch(url);
  if (!res.ok)
    throw new Error(`Download failed: ${res.status} ${res.statusText}`);
  const reader = res.body?.getReader();
  if (!reader) throw new Error("No response body");
  const writer = createWriteStream(dest);
  const pump = async () => {
    while (true) {
      const { done, value } = await reader.read();
      if (done) break;
      writer.write(value);
    }
    writer.close();
  };
  await pump();
}

export async function downloadGonomo(
  gonomoDir: string,
  version = "latest",
): Promise<string> {
  const binDir = join(gonomoDir, "bin");
  await mkdir(binDir, { recursive: true });

  const binPath = join(binDir, gonomoAssetName());
  if (existsSync(binPath)) {
    console.log("Using existing local gonomo binary at ${binPath}");
    return binPath;
  }

  const tag = version === "latest" ? "latest" : "tags/${version}";
  const releaseUrl = `https://api.github.com/repos/${GONOMO_REPO}/releases/${tag}`;

  console.log(`Downloading gonomo binary (${version})...`);
  const res = await fetch(releaseUrl, {
    headers: {
      Accept: "application/vnd.github.v3+json",
      "User-Agent": "gonomo-cli",
    },
  });
  if (!res.ok) throw new Error(`GitHub API error: ${res.status}`);
  const release: any = await res.json();
  const asset = release.assets?.find((a: any) => a.name === gonomoAssetName());
  if (!asset) throw new Error(`No binary found for ${platform()}/${arch()}`);

  console.log(`Downloading ${asset.name}...`);
  await downloadFile(asset.browser_download_url, binPath);
  if (platform() !== "win32") {
    await chmod(binPath, 0o755);
  }
  return binPath;
}

