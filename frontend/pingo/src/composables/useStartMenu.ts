import { ref } from 'vue'

// Shared state for Start Menu
const startMenuOpen = ref(false)

export function useStartMenu() {
  const toggleStartMenu = () => {
    startMenuOpen.value = !startMenuOpen.value
  }

  const closeStartMenu = () => {
    startMenuOpen.value = false
  }

  const openStartMenu = () => {
    startMenuOpen.value = true
  }

  return {
    startMenuOpen,
    toggleStartMenu,
    closeStartMenu,
    openStartMenu
  }
}
