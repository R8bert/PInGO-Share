<template>
  <div class="flex flex-col min-h-screen">
    <Navbar />
    <main class="flex-grow pt-20">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/shared/Navbar.vue'
</script>