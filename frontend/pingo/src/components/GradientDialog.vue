<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
    <!-- Backdrop -->
    <div class="absolute inset-0 bg-black/50 backdrop-blur-sm" @click="closeDialog"></div>
    
    <!-- Dialog -->
    <div class="relative w-full max-w-2xl max-h-[90vh] overflow-y-auto rounded-3xl border shadow-2xl"
         :style="{
           backgroundColor: isDark ? 'rgba(0,0,0,0.9)' : 'rgba(255,255,255,0.95)',
           borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
         }">
      
      <!-- Header -->
      <div class="p-6 border-b"
           :style="{ borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)' }">
        <div class="flex items-center justify-between">
          <h3 class="text-2xl font-bold transition-colors duration-300"
              :style="{ color: isDark ? '#f9fafb' : '#111827' }">
            {{ title }}
          </h3>
          <button @click="closeDialog"
                  class="p-2 rounded-xl transition-all duration-200 hover:scale-110"
                  :style="{
                    backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                    color: isDark ? '#f9fafb' : '#111827'
                  }">
            <IconXMark class="w-6 h-6" />
          </button>
        </div>
      </div>

      <!-- Content -->
      <div class="p-6 space-y-6">
        <!-- Mode Selector -->
        <div class="flex rounded-xl p-1 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.05)'
             }">
          <button
            @click="mode = 'single'"
            :class="[
              'flex-1 px-4 py-2 rounded-lg transition-all duration-200 font-medium',
              mode === 'single' 
                ? 'bg-gradient-to-r from-blue-600 to-purple-600 text-white shadow-lg' 
                : (isDark ? 'text-gray-300 hover:text-white' : 'text-gray-600 hover:text-gray-900')
            ]"
          >
            Single Color
          </button>
          <button
            @click="mode = 'manual'"
            :class="[
              'flex-1 px-4 py-2 rounded-lg transition-all duration-200 font-medium',
              mode === 'manual' 
                ? 'bg-gradient-to-r from-blue-600 to-purple-600 text-white shadow-lg' 
                : (isDark ? 'text-gray-300 hover:text-white' : 'text-gray-600 hover:text-gray-900')
            ]"
          >
            Manual Colors
          </button>
        </div>

        <!-- Single Color Mode -->
        <div v-if="mode === 'single'" class="space-y-4">
          <ColorPicker
            v-model="singleColor"
            label="Base Color"
            input-id="base-color"
          />
          
          <div class="space-y-3">
            <p class="text-sm font-medium transition-colors duration-300"
               :style="{ color: isDark ? '#f9fafb' : '#111827' }">
              Generated Gradient Preview
            </p>
            <div class="h-20 rounded-xl border-2 border-white shadow-lg relative overflow-hidden"
                 :style="{ background: generatedGradient }">
              <div class="absolute inset-0 flex items-center justify-center">
                <span class="text-white font-semibold text-lg drop-shadow-lg">
                  Generated Gradient
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Manual Color Mode -->
        <div v-if="mode === 'manual'" class="space-y-4">
          <ColorPicker
            v-model="colors.color1"
            label="First Color"
            input-id="color1"
          />
          <ColorPicker
            v-model="colors.color2"
            label="Second Color"
            input-id="color2"
          />
          <ColorPicker
            v-model="colors.color3"
            label="Third Color"
            input-id="color3"
          />
          
          <div class="space-y-3">
            <p class="text-sm font-medium transition-colors duration-300"
               :style="{ color: isDark ? '#f9fafb' : '#111827' }">
              Manual Gradient Preview
            </p>
            <div class="h-20 rounded-xl border-2 border-white shadow-lg relative overflow-hidden"
                 :style="{ background: manualGradient }">
              <div class="absolute inset-0 flex items-center justify-center">
                <span class="text-white font-semibold text-lg drop-shadow-lg">
                  Manual Gradient
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <div class="p-6 border-t flex justify-end space-x-3"
           :style="{ borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)' }">
        <button @click="closeDialog"
                class="px-6 py-3 rounded-xl transition-all duration-200 font-medium"
                :style="{
                  backgroundColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                  color: isDark ? '#f9fafb' : '#111827'
                }">
          Cancel
        </button>
        <button @click="applyColors"
                class="px-6 py-3 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-medium rounded-xl transition-all duration-200 hover:scale-105">
          Apply Colors
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useTheme } from '../composables/useTheme'
import ColorPicker from './ColorPicker.vue'
import IconXMark from '~icons/solar/close-circle-bold-duotone'

const { isDark } = useTheme()

interface Props {
  isOpen: boolean
  title: string
  initialColors?: {
    color1: string
    color2: string
    color3: string
  }
}

interface Emits {
  (e: 'close'): void
  (e: 'apply', colors: { color1: string; color2: string; color3: string }): void
}

const props = withDefaults(defineProps<Props>(), {
  initialColors: () => ({
    color1: '#3b82f6',
    color2: '#8b5cf6',
    color3: '#ec4899'
  })
})

const emit = defineEmits<Emits>()

const mode = ref<'single' | 'manual'>('single')
const singleColor = ref('#3b82f6')
const colors = ref({
  color1: props.initialColors.color1,
  color2: props.initialColors.color2,
  color3: props.initialColors.color3
})

// Generate gradient from single color
const generatedGradient = computed(() => {
  const baseColor = singleColor.value
  const hsl = hexToHsl(baseColor)
  
  // Generate complementary colors
  const color1 = baseColor
  const color2 = hslToHex((hsl.h + 60) % 360, Math.max(hsl.s - 10, 40), Math.min(hsl.l + 10, 80))
  const color3 = hslToHex((hsl.h + 120) % 360, Math.max(hsl.s - 5, 50), Math.max(hsl.l - 10, 30))
  
  return `linear-gradient(135deg, ${color1} 0%, ${color2} 50%, ${color3} 100%)`
})

const manualGradient = computed(() => {
  return `linear-gradient(135deg, ${colors.value.color1} 0%, ${colors.value.color2} 50%, ${colors.value.color3} 100%)`
})

// Color conversion utilities
function hexToHsl(hex: string) {
  const r = parseInt(hex.slice(1, 3), 16) / 255
  const g = parseInt(hex.slice(3, 5), 16) / 255
  const b = parseInt(hex.slice(5, 7), 16) / 255
  
  const max = Math.max(r, g, b)
  const min = Math.min(r, g, b)
  let h: number, s: number, l = (max + min) / 2
  
  if (max === min) {
    h = s = 0 // achromatic
  } else {
    const d = max - min
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min)
    
    switch (max) {
      case r: h = (g - b) / d + (g < b ? 6 : 0); break
      case g: h = (b - r) / d + 2; break
      case b: h = (r - g) / d + 4; break
      default: h = 0
    }
    h /= 6
  }
  
  return { h: h * 360, s: s * 100, l: l * 100 }
}

function hslToHex(h: number, s: number, l: number) {
  h = h % 360
  s = Math.max(0, Math.min(100, s)) / 100
  l = Math.max(0, Math.min(100, l)) / 100
  
  const c = (1 - Math.abs(2 * l - 1)) * s
  const x = c * (1 - Math.abs((h / 60) % 2 - 1))
  const m = l - c / 2
  let r = 0, g = 0, b = 0
  
  if (0 <= h && h < 60) {
    r = c; g = x; b = 0
  } else if (60 <= h && h < 120) {
    r = x; g = c; b = 0
  } else if (120 <= h && h < 180) {
    r = 0; g = c; b = x
  } else if (180 <= h && h < 240) {
    r = 0; g = x; b = c
  } else if (240 <= h && h < 300) {
    r = x; g = 0; b = c
  } else if (300 <= h && h < 360) {
    r = c; g = 0; b = x
  }
  
  r = Math.round((r + m) * 255)
  g = Math.round((g + m) * 255)
  b = Math.round((b + m) * 255)
  
  return '#' + [r, g, b].map(x => x.toString(16).padStart(2, '0')).join('')
}

// Update colors when single color changes
watch(singleColor, (newColor) => {
  if (mode.value === 'single') {
    const hsl = hexToHsl(newColor)
    colors.value = {
      color1: newColor,
      color2: hslToHex((hsl.h + 60) % 360, Math.max(hsl.s - 10, 40), Math.min(hsl.l + 10, 80)),
      color3: hslToHex((hsl.h + 120) % 360, Math.max(hsl.s - 5, 50), Math.max(hsl.l - 10, 30))
    }
  }
})

const closeDialog = () => {
  emit('close')
}

const applyColors = () => {
  emit('apply', colors.value)
  closeDialog()
}

// Initialize with props
watch(() => props.initialColors, (newColors) => {
  colors.value = { ...newColors }
}, { immediate: true })
</script>
