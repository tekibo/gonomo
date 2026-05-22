<script setup lang="ts">
import type { GonomoSplashConfig } from 'gonomo';

interface SplashLayout {
  id: string
  label: string
  description: string
}

const layouts: SplashLayout[] = [
  { id: 'centered', label: 'Centered', description: 'Image centered in upper portion, text below' },
  { id: 'minimal', label: 'Minimal', description: 'Large centered text, no image' },
  { id: 'top-banner', label: 'Top Banner', description: 'Image banner at top, text below' },
  { id: 'bottom-banner', label: 'Bottom Banner', description: 'Text at top, image banner at bottom' },
  { id: 'split', label: 'Split', description: 'Image left, text right side by side' },
  { id: 'full-image', label: 'Full Image', description: 'Image fills entire background, text overlaid' },
  { id: 'custom', label: 'Custom', description: 'Transparent background, centered image only' },
];

const props = defineProps<{ config: GonomoSplashConfig | null }>();

function layoutSvg(id: string): string {
  const bg = props.config?.backgroundColor || '#fafafa';
  const fg = props.config?.foregroundColor || '#202427';
  const imgColor = '#888';
  switch (id) {
    case 'centered':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${bg}"/>
        <rect x="35" y="8" width="50" height="38" rx="4" fill="${imgColor}" opacity="0.4"/>
        <rect x="25" y="52" width="70" height="8" rx="4" fill="${fg}" opacity="0.7"/>
        <rect x="40" y="64" width="40" height="4" rx="2" fill="${fg}" opacity="0.4"/>
      </svg>`;
    case 'minimal':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${bg}"/>
        <rect x="20" y="28" width="80" height="12" rx="4" fill="${fg}" opacity="0.8"/>
        <rect x="35" y="44" width="50" height="6" rx="3" fill="${fg}" opacity="0.4"/>
      </svg>`;
    case 'top-banner':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${bg}"/>
        <rect x="0" y="0" width="120" height="36" fill="${imgColor}" opacity="0.35"/>
        <rect x="25" y="44" width="70" height="8" rx="4" fill="${fg}" opacity="0.7"/>
        <rect x="40" y="56" width="40" height="4" rx="2" fill="${fg}" opacity="0.4"/>
      </svg>`;
    case 'bottom-banner':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${bg}"/>
        <rect x="25" y="8" width="70" height="8" rx="4" fill="${fg}" opacity="0.7"/>
        <rect x="40" y="20" width="40" height="4" rx="2" fill="${fg}" opacity="0.4"/>
        <rect x="0" y="44" width="120" height="36" fill="${imgColor}" opacity="0.35"/>
      </svg>`;
    case 'split':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${bg}"/>
        <rect x="8" y="15" width="44" height="50" rx="4" fill="${imgColor}" opacity="0.35"/>
        <rect x="62" y="28" width="50" height="8" rx="4" fill="${fg}" opacity="0.7"/>
        <rect x="68" y="40" width="38" height="4" rx="2" fill="${fg}" opacity="0.4"/>
      </svg>`;
    case 'full-image':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${imgColor}" opacity="0.5"/>
        <rect x="0" y="0" width="120" height="80" fill="${bg}" opacity="0.3"/>
        <rect x="25" y="28" width="70" height="10" rx="4" fill="#fff" opacity="0.9"/>
        <rect x="38" y="42" width="44" height="5" rx="2.5" fill="#fff" opacity="0.6"/>
      </svg>`;
    case 'custom':
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80">
        <rect width="120" height="80" fill="${bg}" opacity="0.15"/>
        <rect x="0" y="0" width="120" height="80" fill="url(#checkers)"/>
        <rect x="30" y="10" width="60" height="60" rx="8" fill="${imgColor}" opacity="0.6"/>
        <defs><pattern id="checkers" width="8" height="8" patternUnits="userSpaceOnUse">
          <rect width="4" height="4" fill="#ccc" opacity="0.3"/>
          <rect x="4" y="4" width="4" height="4" fill="#ccc" opacity="0.3"/>
        </pattern></defs>
      </svg>`;
    default:
      return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 80"><rect width="120" height="80" fill="${bg}"/></svg>`;
  }
}
</script>

<template>
  <div class="splash-preview">
    <h2>Splash Layout Presets</h2>
    <p class="hint">Configure via <code>splash.layout</code> in <code>gonomo.config.ts</code></p>
    <div class="layout-grid">
      <div v-for="l in layouts" :key="l.id" class="layout-card" :class="{ active: props.config?.layout === l.id }"
        @click="$emit('select', l.id)">
        <div class="layout-svg" v-html="layoutSvg(l.id)" />
        <div class="layout-info">
          <strong>{{ l.label }}</strong>
          <span>{{ l.description }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.splash-preview {
  margin-bottom: 28px;
}

.hint {
  font-size: 13px;
  opacity: 0.6;
  margin: 0 0 16px;
}

.hint code {
  font-size: 12px;
  background: rgba(255, 255, 255, 0.08);
  padding: 2px 6px;
  border-radius: 4px;
}

.layout-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 12px;
}

.layout-card {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 12px;
  border-radius: 8px;
  border: 2px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.03);
  cursor: pointer;
  transition: border-color 0.15s, background 0.15s;
}

.layout-card:hover {
  border-color: rgba(255, 255, 255, 0.2);
  background: rgba(255, 255, 255, 0.06);
}

.layout-card.active {
  border-color: var(--color-accent);
  background: var(--color-accent-soft);
}

.layout-svg {
  width: 100%;
  aspect-ratio: 3/2;
  border-radius: 4px;
  overflow: hidden;
}

.layout-svg :deep(svg) {
  display: block;
  width: 100%;
  height: 100%;
}

.layout-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.layout-info strong {
  font-size: 13px;
}

.layout-info span {
  font-size: 11px;
  opacity: 0.6;
}
</style>
