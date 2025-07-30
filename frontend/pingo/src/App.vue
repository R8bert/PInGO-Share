<template>
  <div class="min-h-screen bg-gray-50">
    <Navbar />
    <main class="relative">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useAuth } from './composables/useAuth'
import Navbar from './components/Navbar.vue'

const { fetchCurrentUser } = useAuth()

onMounted(async () => {
  // Try to authenticate user on app start
  await fetchCurrentUser()
})
</script>