<script setup lang="ts">
const sections = [
    { id: "overview", label: "Overview" },
    { id: "config", label: "Configuration" },
    { id: "window-api", label: "Window API" },
    { id: "titlebar", label: "Titlebar Styling" },
    { id: "css-vars", label: "CSS Variables" },
    { id: "splash", label: "Splash Screen" },
    { id: "build", label: "Build & Dev" },
];

const active = defineModel<string>("section", { default: "overview" });
</script>

<template>
    <aside class="sidebar">
        <div class="sidebar-section-label">Docs</div>
        <nav class="sidebar-nav">
            <button
                v-for="s in sections"
                :key="s.id"
                class="sidebar-item"
                :class="{ active: active === s.id }"
                @click="active = s.id"
            >
                {{ s.label }}
            </button>
        </nav>
    </aside>
</template>

<style scoped>
.sidebar {
    width: 200px;
    display: flex;
    flex-direction: column;
    background: var(--color-background);
    border-right: 1px solid var(--color-border);
    flex-shrink: 0;
    overflow-y: auto;
    padding: 12px 0;
}

.sidebar-section-label {
    font-size: 10px;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 1px;
    opacity: 0.4;
    padding: 8px;
}

.sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 1px;
}

.sidebar-item {
    display: block;
    width: 100%;
    text-align: left;
    padding: 7px 16px;
    font-size: 12px;
    background: none;
    border: none;
    color: var(--color-text);
    cursor: pointer;
    opacity: 0.7;
}
.sidebar-item:hover {
    background: var(--color-surface-high);
    opacity: 1;
}
.sidebar-item.active {
    background: var(--color-accent-soft);
    opacity: 1;
    font-weight: 600;
}
</style>
