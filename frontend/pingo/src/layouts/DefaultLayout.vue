<template>
  <div class="flex flex-col min-h-screen">
    <LiquidGlassNavbar />
    <main class="flex-grow pt-20">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import LiquidGlassNavbar from '../components/shared/LiquidGlassNavbar.vue'
</script>