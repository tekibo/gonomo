<template>
    <div class="layout">
        <TitleBar />
        <div class="body">
            <AppSidebar v-model:section="activeSection" />
            <main class="content">
                <slot />
            </main>
        </div>
    </div>
</template>

<script setup lang="ts">
const activeSection = ref("overview");
provide("activeSection", activeSection);
</script>

<style scoped>
.layout {
    display: flex;
    flex-direction: column;
    height: 100%;
}

.body {
    flex: 1;
    display: flex;
    min-height: 0;
}

.content {
    flex: 1;
    overflow-y: auto;
    padding: 24px 32px;
    min-height: 0;
}
</style>
