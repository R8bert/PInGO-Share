<template>
  <div class="min-h-screen transition-colors duration-300">
    <Navbar />
    
    <main class="relative">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuth } from './composables/useAuth'
import { useTheme } from './composables/useTheme'
import Navbar from './components/Navbar.vue'

const { fetchCurrentUser } = useAuth()
const { } = useTheme() // Initialize theme system

onMounted(async () => {
  // Try to authenticate user on app start
  await fetchCurrentUser()
})
</script>