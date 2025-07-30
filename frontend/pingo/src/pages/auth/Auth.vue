<template>
  <div>
    <Login v-if="mode === 'login'" @switch-to-register="mode = 'register'" />
    <Register v-else @switch-to-login="mode = 'login'" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Login from './Login.vue'
import Register from './Register.vue'

const mode = ref<'login' | 'register'>('login')
</script>
