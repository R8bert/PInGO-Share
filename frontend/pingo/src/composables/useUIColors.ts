import { computed } from 'vue'
import { useSettings } from './useSettings'

export function useUIColors() {
  const { settings } = useSettings()

  // Gradient colors for buttons and UI elements
  const gradientColors = computed(() => {
    if (!settings.value) {
      return {
        color1: '#3b82f6',
        color2: '#8b5cf6', 
        color3: '#ec4899'
      }
    }
    
    return {
      color1: settings.value.gradientColor1 || '#3b82f6',
      color2: settings.value.gradientColor2 || '#8b5cf6',
      color3: settings.value.gradientColor3 || '#ec4899'
    }
  })

  // CSS gradient string for use in styles
  const primaryGradient = computed(() => {
    const colors = gradientColors.value
    return `linear-gradient(135deg, ${colors.color1} 0%, ${colors.color2} 50%, ${colors.color3} 100%)`
  })

  // Two-color gradient for simpler buttons
  const buttonGradient = computed(() => {
    const colors = gradientColors.value
    return `linear-gradient(to right, ${colors.color1}, ${colors.color2})`
  })

  // Hover gradient with darker shades
  const hoverGradient = computed(() => {
    const colors = gradientColors.value
    return `linear-gradient(to right, ${darkenColor(colors.color1, 10)}, ${darkenColor(colors.color2, 10)})`
  })

  // Primary color (first gradient color)
  const primaryColor = computed(() => gradientColors.value.color1)

  // Secondary color (second gradient color)  
  const secondaryColor = computed(() => gradientColors.value.color2)

  // Accent color (third gradient color)
  const accentColor = computed(() => gradientColors.value.color3)

  // Website background color from settings
  const websiteColor = computed(() => settings.value?.websiteColor || '#3b82f6')

  return {
    gradientColors,
    primaryGradient,
    buttonGradient, 
    hoverGradient,
    primaryColor,
    secondaryColor,
    accentColor,
    websiteColor
  }
}

// Utility function to darken a hex color by a percentage
function darkenColor(hex: string, percent: number): string {
  if (!hex || hex.length !== 7) return hex
  
  const num = parseInt(hex.slice(1), 16)
  const amt = Math.round(2.55 * percent)
  const R = (num >> 16) - amt
  const G = (num >> 8 & 0x00FF) - amt  
  const B = (num & 0x0000FF) - amt
  
  return '#' + (0x1000000 + (R < 255 ? (R < 1 ? 0 : R) : 255) * 0x10000 +
    (G < 255 ? (G < 1 ? 0 : G) : 255) * 0x100 +
    (B < 255 ? (B < 1 ? 0 : B) : 255)).toString(16).slice(1)
}
