<script setup lang="ts">
const activeSection = inject<Ref<string>>("activeSection")!;
const { onClose, onMinimize, onMaximize, onRestore, isMaximized } = useGonomo();
</script>

<template>
    <!-- Overview -->
    <section v-if="activeSection === 'overview'">
        <h1 class="section-title">Gonomo</h1>
        <p class="section-desc">Native desktop apps with web technologies. Gonomo wraps your Nuxt/Vue app into a
            standalone Windows executable using WebView2.</p>

        <div class="card">
            <h2 class="card-title">What is Gonomo?</h2>
            <p style="font-size:12px;opacity:0.7;line-height:1.6;margin:0;">
                Gonomo is a CLI tool and Go runtime that takes your Nuxt (or any) frontend and packages it into a native
                Windows executable.
                It provides a bridge between JavaScript and the Windows OS — giving you control over the window,
                titlebar, taskbar, and splash screen.
            </p>
        </div>

        <div style="display:grid;grid-template-columns:1fr 1fr;gap:12px;margin-top:16px;">
            <div class="card">
                <h2 class="card-title">Frontend Agnostic</h2>
                <p style="font-size:12px;opacity:0.6;margin:0;line-height:1.5;">
                    Works with Nuxt, Vite, or any framework. Your frontend builds to static files or a Node server
                    bundled inside the .exe.
                </p>
            </div>
            <div class="card">
                <h2 class="card-title">Native Bridge</h2>
                <p style="font-size:12px;opacity:0.6;margin:0;line-height:1.5;">
                    Call window.gonomo.* from JavaScript to control the window, titlebar, colors, and splash — no IPC
                    boilerplate.
                </p>
            </div>
            <div class="card">
                <h2 class="card-title">Custom Titlebar</h2>
                <p style="font-size:12px;opacity:0.6;margin:0;line-height:1.5;">
                    Hide the native titlebar and render your own with CSS. Native caption buttons float over your
                    content via the overlay system.
                </p>
            </div>
            <div class="card">
                <h2 class="card-title">Splash Screen</h2>
                <p style="font-size:12px;opacity:0.6;margin:0;line-height:1.5;">
                    Native splash window with multiple layouts, images, and progress indication while the backend boots.
                </p>
            </div>
        </div>
    </section>

    <!-- Configuration -->
    <section v-if="activeSection === 'config'">
        <h1 class="section-title">Configuration</h1>
        <p class="section-desc">The <code>gonomo.config.ts</code> file controls every aspect of the build and runtime.
        </p>

        <div class="card">
            <h2 class="card-title">gonomo.config.ts Reference</h2>
            <pre style="margin-top:8px;">{
  "name": "my-app",
  "title": "My App",
  "icon": "icon.ico",
  "window": {
    "width": 1400,
    "height": 900,
    "maximized": false,
    "titleBarStyle": {
      "hidden": true,
      "overlay": true,
      "darkMode": true,
      "captionColor": "#202427",
      "textColor": "#cdd6f4"
    }
  },
  "splash": {
    "enabled": true,
    "layout": "centered",
    "image": "icon.png",
    "backgroundColor": "#1e1e2e",
    "foregroundColor": "#cdd6f4",
    "text": "Loading...",
    "minDuration": 2000
  },
  "build": {
    "command": "pnpm run build",
    "outputDir": ".output",
    "entry": "server/index.mjs",
    "runtime": "node",
    "embed": "full"
  },
  "dev": {
    "command": "pnpm run dev",
    "url": "http://localhost:3000"
  },
  "output": {
    "dir": "./dist",
    "name": "app.exe"
  }
}</pre>
        </div>

        <div style="display:grid;grid-template-columns:1fr 1fr;gap:12px;margin-top:16px;">
            <div class="card">
                <h2 class="card-title">Window Options</h2>
                <div class="prop-grid" style="margin-top:8px;">
                    <span class="prop-key">width</span><span class="prop-val">1400 (default)</span>
                    <span class="prop-key">height</span><span class="prop-val">900 (default)</span>
                    <span class="prop-key">maximized</span><span class="prop-val">false</span>
                    <span class="prop-key">titleBarStyle</span><span class="prop-val">"normal" | "hidden" |
                        object</span>
                </div>
            </div>
            <div class="card">
                <h2 class="card-title">titleBarStyle Object</h2>
                <div class="prop-grid" style="margin-top:8px;">
                    <span class="prop-key">hidden</span><span class="prop-val">boolean</span>
                    <span class="prop-key">overlay</span><span class="prop-val">boolean</span>
                    <span class="prop-key">darkMode</span><span class="prop-val">boolean</span>
                    <span class="prop-key">captionColor</span><span class="prop-val">hex string</span>
                    <span class="prop-key">textColor</span><span class="prop-val">hex string</span>
                </div>
            </div>
            <div class="card">
                <h2 class="card-title">Splash Options</h2>
                <div class="prop-grid" style="margin-top:8px;">
                    <span class="prop-key">enabled</span><span class="prop-val">false</span>
                    <span class="prop-key">layout</span><span class="prop-val">centered | minimal | custom …</span>
                    <span class="prop-key">image</span><span class="prop-val">path to PNG</span>
                    <span class="prop-key">minDuration</span><span class="prop-val">ms (default 400)</span>
                </div>
            </div>
            <div class="card">
                <h2 class="card-title">Build Options</h2>
                <div class="prop-grid" style="margin-top:8px;">
                    <span class="prop-key">runtime</span><span class="prop-val">"node" | "static"</span>
                    <span class="prop-key">embed</span><span class="prop-val">"full" | "none"</span>
                    <span class="prop-key">entry</span><span class="prop-val">server entry path</span>
                    <span class="prop-key">command</span><span class="prop-val">build script</span>
                </div>
            </div>
        </div>
    </section>

    <!-- Window API -->
    <section v-if="activeSection === 'window-api'">
        <h1 class="section-title">Window API</h1>
        <p class="section-desc">Control the native window from JavaScript via <code>window.gonomo</code>.</p>

        <div class="card">
            <h2 class="card-title">Window Controls</h2>
            <div style="display:flex;gap:8px;margin-top:12px;flex-wrap:wrap;">
                <button class="btn btn-primary" @click="onMinimize">Minimize</button>
                <button class="btn btn-primary" @click="onMaximize">Maximize</button>
                <button class="btn btn-primary" @click="onRestore">Restore</button>
                <button class="btn" @click="onClose">Close</button>
            </div>
            <p style="font-size:12px;opacity:0.5;margin:8px 0 0;">IsMaximized: <code>{{ isMaximized }}</code></p>
        </div>

        <div class="card" style="margin-top:12px;">
            <h2 class="card-title">Method Reference</h2>
            <pre style="margin-top:8px;">window.gonomo.Close()          // Close the window
window.gonomo.Minimize()       // Minimize
window.gonomo.Maximize()       // Maximize
window.gonomo.Restore()        // Restore from max/min
window.gonomo.IsMaximized()    // Returns Promise&lt;boolean&gt;

window.gonomo.setDarkMode(bool)
window.gonomo.setCaptionColor(hex)
window.gonomo.setTextColor(hex)
window.gonomo.setTitlebarVisible(bool)
window.gonomo.setTitleBarOverlay(bool)</pre>
        </div>

        <div class="card" style="margin-top:12px;">
            <h2 class="card-title">Composable Usage (Vue)</h2>
            <pre style="margin-top:8px;">import { useGonomo } from "~/composables/useGonomo";

const {
  gonomo,              // window.gonomo ref
  isMaximized,         // reactive boolean
  appTitle,            // window title
  onClose,
  onMinimize,
  onMaximize,
  onRestore,
  onToggleMaximize,
} = useGonomo();</pre>
        </div>
    </section>

    <!-- Titlebar Styling -->
    <section v-if="activeSection === 'titlebar'">
        <h1 class="section-title">Titlebar Styling</h1>
        <p class="section-desc">Hide the OS titlebar and build your own with CSS + native caption buttons.</p>

        <div style="display:grid;grid-template-columns:1fr 1fr;gap:12px;">
            <div class="card">
                <h2 class="card-title">Overlay Mode</h2>
                <p style="font-size:12px;opacity:0.6;line-height:1.5;">
                    When <code>hidden: true</code> + <code>overlay: true</code>, the native titlebar
                    is hidden but the minimize, maximize, and close buttons are rendered by the OS
                    floating over your web content — including Windows 11 Snap Layouts.
                </p>
            </div>
            <div class="card">
                <h2 class="card-title">Draggable Region</h2>
                <p style="font-size:12px;opacity:0.6;line-height:1.5;">
                    Use <code>-webkit-app-region: drag</code> on your custom titlebar to enable
                    window dragging. Set <code>no-drag</code> on interactive children (menus, buttons).
                    Double-clicking a drag region toggles maximize natively.
                </p>
            </div>
        </div>

        <div class="card" style="margin-top:12px;">
            <h2 class="card-title">Configuration Example</h2>
            <pre style="margin-top:8px;">{
  "titleBarStyle": {
    "hidden": true,
    "overlay": true,
    "darkMode": true,
    "captionColor": "#202427",
    "textColor": "#cdd6f4"
  }
}</pre>
        </div>
    </section>

    <!-- CSS Variables -->
    <section v-if="activeSection === 'css-vars'">
        <h1 class="section-title">CSS Variables</h1>
        <p class="section-desc">Gonomo injects CSS custom properties on <code>:root</code> for precise titlebar sizing.
        </p>

        <div class="card">
            <h2 class="card-title">Injected by Bridge</h2>
            <div class="prop-grid" style="margin-top:8px;">
                <span class="prop-key">--gonomo-caption-button-height</span>
                <span class="prop-val">Native caption button height (DPI-scaled, e.g. 32px)</span>
                <span class="prop-key">--gonomo-caption-buttons-width</span>
                <span class="prop-val">Total width of min+max+close buttons (DPI-scaled, e.g. 138px)</span>
            </div>
        </div>

        <div class="card" style="margin-top:12px;">
            <h2 class="card-title">Usage</h2>
            <pre style="margin-top:8px;">.titlebar {
  height: var(--gonomo-caption-button-height, 32px);
  padding-right: var(--gonomo-caption-buttons-width, 138px);
}</pre>
            <p style="font-size:12px;opacity:0.5;margin-top:8px;">
                The fallback values match 96 DPI. On high-DPI displays the actual values scale automatically.
            </p>
        </div>

        <div class="card" style="margin-top:12px;">
            <h2 class="card-title">Theme Variables</h2>
            <p style="font-size:12px;opacity:0.6;line-height:1.5;">
                The example app defines these theme variables in <code>assets/css/theme.css</code>.
                Set your <code>captionColor</code> in gonomo.config.ts to match <code>--color-surface</code> for a
                seamless look.
            </p>
            <pre style="margin-top:8px;">:root {
  --color-background: oklch(25.75% 0.008 240.18);
  --color-surface:   oklch(30.75% 0.008 240.18);
  --color-accent:    oklch(45.75% 0.18 240.18);
  --color-text:      oklch(92.75% 0.02 240.18);
}</pre>
        </div>
    </section>

    <!-- Splash Screen -->
    <section v-if="activeSection === 'splash'">
        <h1 class="section-title">Splash Screen</h1>
        <p class="section-desc">A native splash window shown while the backend server starts.</p>

        <div style="display:grid;grid-template-columns:1fr 1fr;gap:12px;">
            <div class="card">
                <h2 class="card-title">Layouts</h2>
                <div class="prop-grid" style="margin-top:8px;">
                    <span class="prop-key">centered</span><span class="prop-val">Image + text centered</span>
                    <span class="prop-key">minimal</span><span class="prop-val">Minimal text only</span>
                    <span class="prop-key">top-banner</span><span class="prop-val">Image banner at top</span>
                    <span class="prop-key">bottom-banner</span><span class="prop-val">Image banner at bottom</span>
                    <span class="prop-key">split</span><span class="prop-val">Side-by-side image + text</span>
                    <span class="prop-key">full-image</span><span class="prop-val">Full window image</span>
                    <span class="prop-key">custom</span><span class="prop-val">Transparent with color-key</span>
                </div>
            </div>
            <div class="card">
                <h2 class="card-title">Configuration</h2>
                <pre style="margin-top:8px;">{
  "splash": {
    "enabled": true,
    "layout": "centered",
    "backgroundColor": "#1e1e2e",
    "foregroundColor": "#cdd6f4",
    "image": "icon.png",
    "text": "My App",
    "minDuration": 2000,
    "width": 480,
    "height": 320
  }
}</pre>
            </div>
        </div>
    </section>

    <!-- Build & Dev -->
    <section v-if="activeSection === 'build'">
        <h1 class="section-title">Build & Dev</h1>
        <p class="section-desc">Commands and workflows for development and production builds.</p>

        <div style="display:grid;grid-template-columns:1fr 1fr;gap:12px;">
            <div class="card">
                <h2 class="card-title">Dev Mode</h2>
                <pre style="margin-top:8px;">npm run gonomo:dev</pre>
                <p style="font-size:12px;opacity:0.6;line-height:1.5;margin-top:8px;">
                    Starts the frontend dev server, compiles a dev executable that connects to it,
                    and launches the app. The <code>dev.url</code> and <code>dev.command</code> in
                    gonomo.config.ts controls the dev server.
                </p>
            </div>
            <div class="card">
                <h2 class="card-title">Production Build</h2>
                <pre style="margin-top:8px;">npm run gonomo:build</pre>
                <p style="font-size:12px;opacity:0.6;line-height:1.5;margin-top:8px;">
                    Runs the frontend build command, embeds all assets into the Go binary,
                    and produces a single <code>.exe</code> file. The <code>output</code> config
                    controls the destination.
                </p>
            </div>
        </div>

        <div class="card" style="margin-top:12px;">
            <h2 class="card-title">Runtime Options</h2>
            <table style="width:100%;font-size:12px;border-collapse:collapse;margin-top:8px;">
                <thead>
                    <tr style="border-bottom:1px solid var(--color-border);">
                        <th style="text-align:left;padding:6px 8px;font-weight:600;">Option</th>
                        <th style="text-align:left;padding:6px 8px;font-weight:600;">Description</th>
                    </tr>
                </thead>
                <tbody>
                    <tr style="border-bottom:1px solid var(--color-border);">
                        <td style="padding:6px 8px;"><code>runtime: "node"</code></td>
                        <td style="padding:6px 8px;opacity:0.7;">Bundles Node.js for SSR apps (Nuxt, Nitro)</td>
                    </tr>
                    <tr style="border-bottom:1px solid var(--color-border);">
                        <td style="padding:6px 8px;"><code>runtime: "static"</code></td>
                        <td style="padding:6px 8px;opacity:0.7;">Serves static files, no Node runtime needed</td>
                    </tr>
                    <tr style="border-bottom:1px solid var(--color-border);">
                        <td style="padding:6px 8px;"><code>embed: "full"</code></td>
                        <td style="padding:6px 8px;opacity:0.7;">All assets inside the .exe (single file)</td>
                    </tr>
                    <tr>
                        <td style="padding:6px 8px;"><code>embed: "extract"</code></td>
                        <td style="padding:6px 8px;opacity:0.7;">Assets extracted to sidecar directory</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </section>
</template>

<style scoped>
section {
    max-width: 800px;
}
</style>
