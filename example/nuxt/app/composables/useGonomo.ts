import { gonomoInit } from 'gonomo';
import type { GonomoAPI } from 'gonomo'

export function useGonomo() {
  const gonomo = ref<GonomoAPI | null>(null);
  const splashDismissed = ref(false);
  const isMaximized = ref(false);
  const appTitle = ref("");

  onMounted(async () => {
    const api = gonomoInit();
    if (!api) return;

    gonomo.value = api;
    splashDismissed.value = true;

    appTitle.value = api.Config?.title ?? "";
    isMaximized.value = api.IsMaximized();
  });

  const onClose = () => gonomo.value?.Close();
  const onMinimize = () => gonomo.value?.Minimize();
  const onMaximize = () => gonomo.value?.Maximize();
  const onRestore = () => gonomo.value?.Restore();
  const onToggleMaximize = () => {
    if (isMaximized.value) {
      onRestore();
    } else {
      onMaximize();
    }
    isMaximized.value = !isMaximized.value;
  };
  const onToggleTitlebar = (visible: boolean) =>
    gonomo.value?.setTitlebarVisible(visible);
  const onToggleTitleBarOverlay = (enabled: boolean) =>
    gonomo.value?.setTitleBarOverlay(enabled);

  return {
    gonomo,
    splashDismissed,
    isMaximized,
    appTitle,
    onClose,
    onMinimize,
    onMaximize,
    onRestore,
    onToggleMaximize,
    onToggleTitlebar,
    onToggleTitleBarOverlay,
  };
}
