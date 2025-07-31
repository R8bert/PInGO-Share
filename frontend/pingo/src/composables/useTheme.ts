import { ref } from 'vue'

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

// Initialize theme on first load
loadTheme()

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

  return {
    theme,
    isDark,
    toggleTheme
  }
}