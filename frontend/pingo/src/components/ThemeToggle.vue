<template>
  <button
    @click="handleToggle"
    :class="[
      'relative inline-flex items-center justify-center p-2 rounded-lg transition-all duration-300',
      'hover:scale-110 active:scale-95',
      'bg-white dark:bg-gray-800 shadow-lg hover:shadow-xl',
      'border border-gray-200 dark:border-gray-600',
      'text-gray-600 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400'
    ]"
    :title="`Switch to ${isDark ? 'light' : 'dark'} mode`"
    aria-label="Toggle theme"
  >
    <!-- Sun Icon (Light Mode) -->
    <transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="scale-0 rotate-90 opacity-0"
      enter-to-class="scale-100 rotate-0 opacity-100"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="scale-100 rotate-0 opacity-100"
      leave-to-class="scale-0 -rotate-90 opacity-0"
    >
      <svg
        v-if="!isDark"
        class="w-5 h-5"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
        />
      </svg>
    </transition>

    <!-- Moon Icon (Dark Mode) -->
    <transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="scale-0 -rotate-90 opacity-0"
      enter-to-class="scale-100 rotate-0 opacity-100"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="scale-100 rotate-0 opacity-100"
      leave-to-class="scale-0 rotate-90 opacity-0"
    >
      <svg
        v-if="isDark"
        class="w-5 h-5"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
        />
      </svg>
    </transition>

    <!-- Ripple effect on click -->
    <div
      v-if="isClicked"
      class="absolute inset-0 rounded-lg bg-blue-400 opacity-30 animate-ping"
    ></div>
  </button>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTheme } from '../composables/useTheme'

const { isDark, toggleTheme } = useTheme()
const isClicked = ref(false)

const handleToggle = () => {
  isClicked.value = true
  toggleTheme()
  
  // Reset click animation
  setTimeout(() => {
    isClicked.value = false
  }, 300)
}
</script>

<style scoped>
/* Additional hover effects */
button:hover {
  transform: scale(1.1);
}

button:active {
  transform: scale(0.95);
}

/* Smooth icon transitions */
svg {
  transition: all 0.3s ease;
}
</style>
