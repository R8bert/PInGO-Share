<template>
  <button
    @click="handleToggle"
    class="relative inline-flex items-center w-14 h-7 rounded-full transition-all duration-300 ease-in-out focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
    :style="{
      backgroundColor: isDark ? '#4f46e5' : '#e5e7eb'
    }"
    :title="`Switch to ${isDark ? 'light' : 'dark'} mode`"
    aria-label="Toggle theme"
  >
    <!-- Toggle circle -->
    <div
      class="w-6 h-6 rounded-full transition-all duration-300 ease-in-out shadow-lg flex items-center justify-center absolute top-0.5"
      :style="{
        backgroundColor: isDark ? '#ffffff' : '#ffffff',
        transform: isDark ? 'translateX(28px)' : 'translateX(2px)'
      }"
    >
      <!-- Sun Icon (Light Mode) -->
      <transition
        enter-active-class="transition-all duration-200 ease-out"
        enter-from-class="scale-0 rotate-90 opacity-0"
        enter-to-class="scale-100 rotate-0 opacity-100"
        leave-active-class="transition-all duration-200 ease-in"
        leave-from-class="scale-100 rotate-0 opacity-100"
        leave-to-class="scale-0 -rotate-90 opacity-0"
      >
        <svg
          v-if="!isDark"
          class="w-4 h-4 text-amber-500"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path
            fill-rule="evenodd"
            d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z"
            clip-rule="evenodd"
          />
        </svg>
      </transition>

      <!-- Moon Icon (Dark Mode) -->
      <transition
        enter-active-class="transition-all duration-200 ease-out"
        enter-from-class="scale-0 -rotate-90 opacity-0"
        enter-to-class="scale-100 rotate-0 opacity-100"
        leave-active-class="transition-all duration-200 ease-in"
        leave-from-class="scale-100 rotate-0 opacity-100"
        leave-to-class="scale-0 rotate-90 opacity-0"
      >
        <svg
          v-if="isDark"
          class="w-4 h-4 text-indigo-600"
          fill="currentColor"
          viewBox="0 0 20 20"
        >
          <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z" />
        </svg>
      </transition>
    </div>
  </button>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTheme } from '../../composables/useTheme'

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
/* Clean toggle with subtle hover effects */
button:hover {
  transform: scale(1.02);
}

button:active {
  transform: scale(0.98);
}

/* Smooth transitions for all elements */
button, div, svg {
  transition: all 0.2s ease;
}

/* Focus styles for accessibility */
button:focus {
  outline: none;
}
</style>
