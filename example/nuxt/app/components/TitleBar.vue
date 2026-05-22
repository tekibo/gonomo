<script setup lang="ts"></script>

<template>
    <header class="titlebar drag">
        <div class="titlebar-start">
            <slot name="start" />
        </div>

        <nav class="menubar">
            <slot name="menubar" />
        </nav>
    </header>
    <div class="border"></div>
</template>

<style scoped>
.border {
    border-bottom: 1px solid var(--color-border);
    width: 100%;
}

.titlebar {
    display: flex;
    align-items: center;
    height: var(--gonomo-caption-button-height);
    padding-right: var(--gonomo-caption-buttons-width);
    background: var(--color-background);
    user-select: none;
    flex-shrink: 0;
    z-index: 999;
}

.titlebar-start {
    display: flex;
    align-items: center;
    padding: 0 12px;
}

.menubar {
    display: flex;
    padding-bottom: var(--gonomo-resize-border-top);
    align-items: stretch;
    -webkit-app-region: no-drag;
}

.drag {
    -webkit-app-region: drag;
}
</style>
