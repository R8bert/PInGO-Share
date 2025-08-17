<template>
  <canvas
    ref="canvasRef"
    class="absolute inset-0 w-full h-full"
    :class="className"
  />
</template>

<script setup lang="ts">
import { computed, watch, ref } from 'vue'
import { useWebGLBackground } from '../composables/useWebGLBackground'
import { useSettings } from '../composables/useSettings'

interface Props {
  hueShift?: number
  noiseIntensity?: number
  scanlineIntensity?: number
  speed?: number
  scanlineFrequency?: number
  warpAmount?: number
  resolutionScale?: number
  className?: string
  useSettingsColor?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  hueShift: 0,
  noiseIntensity: 0.02,
  scanlineIntensity: 0.1,
  speed: 0.5,
  scanlineFrequency: 0.5,
  warpAmount: 0.1,
  resolutionScale: 1,
  className: '',
  useSettingsColor: true
})

const { settings } = useSettings()
const backgroundKey = ref(0) // Force re-initialization when needed

// Convert hex color to hue shift
const hexToHueShift = (hex: string): number => {
  if (!hex || hex.length !== 7) return 0
  
  const r = parseInt(hex.slice(1, 3), 16) / 255
  const g = parseInt(hex.slice(3, 5), 16) / 255
  const b = parseInt(hex.slice(5, 7), 16) / 255
  
  const max = Math.max(r, g, b)
  const min = Math.min(r, g, b)
  let h = 0
  
  if (max !== min) {
    const d = max - min
    switch (max) {
      case r: h = (g - b) / d + (g < b ? 6 : 0); break
      case g: h = (b - r) / d + 2; break
      case b: h = (r - g) / d + 4; break
    }
    h /= 6
  }
  
  return h * 360
}

const computedHueShift = computed(() => {
  if (!props.useSettingsColor || !settings.value?.websiteColor) {
    return props.hueShift
  }
  return hexToHueShift(settings.value.websiteColor)
})

// Watch for settings changes and force re-initialization
watch(() => settings.value?.websiteColor, () => {
  if (props.useSettingsColor) {
    backgroundKey.value++
  }
})

const { canvasRef } = useWebGLBackground({
  hueShift: computedHueShift.value,
  noiseIntensity: props.noiseIntensity,
  scanlineIntensity: props.scanlineIntensity,
  speed: props.speed,
  scanlineFrequency: props.scanlineFrequency,
  warpAmount: props.warpAmount,
  resolutionScale: props.resolutionScale
})
</script>
