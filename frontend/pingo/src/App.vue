<template>
  <div class="min-h-screen transition-colors duration-300">
    <Navbar />
    
    <main class="relative">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted, onBeforeUnmount } from 'vue'
import { useAuth } from './composables/useAuth'
import { useTheme } from './composables/useTheme'
import Navbar from './components/shared/Navbar.vue'

const { fetchCurrentUser } = useAuth()
const { cleanupSystemThemeListener } = useTheme() // Initialize theme system

let initTimeout: number | null = null

onMounted(async () => {
  // Debounce the initialization to prevent rapid refresh issues
  if (initTimeout) {
    clearTimeout(initTimeout)
  }
  
  initTimeout = setTimeout(async () => {
    // Try to authenticate user on app start
    await fetchCurrentUser()
  }, 50) // Small delay to debounce rapid page refreshes
})

onBeforeUnmount(() => {
  if (initTimeout) {
    clearTimeout(initTimeout)
  }
  // Cleanup theme listener
  if (cleanupSystemThemeListener) {
    cleanupSystemThemeListener()
  }
})
</script>