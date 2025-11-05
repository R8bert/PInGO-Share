import { ref, watch } from 'vue'

// Global state - this will work
const isDarkMode = ref(false)

// Load theme from localStorage on initialization
const loadTheme = () => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme === 'dark') {
    isDarkMode.value = true
  } else if (savedTheme === 'light') {
    isDarkMode.value = false
  } else {
    // Default to system preference if no saved theme
    isDarkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches
  }
  applyTheme()
}

// Apply theme to DOM
const applyTheme = () => {
  const html = document.documentElement

  if (isDarkMode.value) {
    html.classList.add('dark')
    html.style.backgroundColor = '#000000'
    document.body.style.backgroundColor = '#000000'
    document.body.style.color = '#ffffff'
  } else {
    html.classList.remove('dark')
    html.style.backgroundColor = '#ffffff'
    document.body.style.backgroundColor = '#ffffff'
    document.body.style.color = '#000000'
  }
}

// Listen for system theme changes
const setupSystemThemeListener = () => {
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')

  const handleChange = (e: MediaQueryListEvent) => {
    // Only update if user hasn't manually set a theme
    const savedTheme = localStorage.getItem('theme')
    if (!savedTheme) {
      isDarkMode.value = e.matches
      applyTheme()
    }
  }

  // Modern browsers
  if (mediaQuery.addEventListener) {
    mediaQuery.addEventListener('change', handleChange)
  } else {
    // Fallback for older browsers
    mediaQuery.addListener(handleChange)
  }

  return () => {
    if (mediaQuery.removeEventListener) {
      mediaQuery.removeEventListener('change', handleChange)
    } else {
      mediaQuery.removeListener(handleChange)
    }
  }
}

// Initialize theme on first load
loadTheme()
const cleanupSystemThemeListener = setupSystemThemeListener()

export function useTheme() {
  const theme = ref<'light' | 'dark'>(isDarkMode.value ? 'dark' : 'light')
  const isDark = isDarkMode

    // Simple toggle that DEFINITELY works
  const toggleTheme = () => {
    isDarkMode.value = !isDarkMode.value
    theme.value = isDarkMode.value ? 'dark' : 'light'
    
    // Save to localStorage
    localStorage.setItem('theme', theme.value)
    
    applyTheme()
    
    console.log('THEME TOGGLED:', isDarkMode.value ? 'DARK' : 'LIGHT')
  }

  // Reset theme to browser preference (removes saved preference)
  const resetThemeToBrowser = () => {
    // Remove saved theme preference
    localStorage.removeItem('theme')
    
    // Set to current browser preference
    isDarkMode.value = window.matchMedia('(prefers-color-scheme: dark)').matches
    theme.value = isDarkMode.value ? 'dark' : 'light'
    
    applyTheme()
    
    console.log('THEME RESET TO BROWSER:', isDarkMode.value ? 'DARK' : 'LIGHT')
  }

  // Watch for changes and keep theme in sync
  watch(isDarkMode, (newValue) => {
    theme.value = newValue ? 'dark' : 'light'
  })

  return {
    theme,
    isDark,
    toggleTheme,
    resetThemeToBrowser,
    cleanupSystemThemeListener
  }
}