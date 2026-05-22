import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const tomlPath = path.resolve(__dirname, '../../../global.toml');
const outPath = path.resolve(__dirname, '../src/constants.ts');

const toml = fs.readFileSync(tomlPath, 'utf-8');
const match = toml.match(/repo\s*=\s*"([^"]+)"/);

if (!match) {
  console.error("Failed to find 'repo' in global.toml");
  process.exit(1);
}

const content = `// Generated from global.toml. DO NOT EDIT.\nexport const GONOMO_REPO = "${match[1]}";\n`;
fs.writeFileSync(outPath, content);
console.log("Generated src/constants.ts from global.toml");
