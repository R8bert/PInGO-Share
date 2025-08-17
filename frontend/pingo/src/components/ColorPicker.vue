<template>
  <div class="space-y-4">
    <!-- Color Input -->
    <div class="relative">
      <label :for="inputId" class="block text-sm font-medium mb-2 transition-colors duration-300"
             :style="{ color: isDark ? '#f9fafb' : '#111827' }">
        {{ label }}
      </label>
      <div class="relative">
        <input
          :id="inputId"
          type="color"
          :value="modelValue"
          @input="handleColorInput"
          class="w-full h-12 rounded-xl border-2 cursor-pointer transition-all duration-300 hover:scale-105"
          :style="{
            borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
            backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)'
          }"
        />
        <div class="absolute inset-y-0 right-3 flex items-center pointer-events-none">
          <div class="w-8 h-8 rounded-lg border-2 border-white shadow-lg"
               :style="{ backgroundColor: modelValue }"></div>
        </div>
      </div>
    </div>

    <!-- Hex Input -->
    <div class="relative">
      <input
        type="text"
        :value="modelValue"
        @input="handleHexInput"
        placeholder="#3b82f6"
        maxlength="7"
        class="w-full px-4 py-3 rounded-xl border transition-all duration-300 font-mono text-center"
        :style="{
          backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
          borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
          color: isDark ? '#f9fafb' : '#111827'
        }"
      />
    </div>

    <!-- Preset Colors -->
    <div v-if="showPresets" class="space-y-2">
      <p class="text-sm font-medium transition-colors duration-300"
         :style="{ color: isDark ? '#f9fafb' : '#111827' }">
        Quick Colors
      </p>
      <div class="grid grid-cols-8 gap-2">
        <button
          v-for="color in presetColors"
          :key="color"
          @click="$emit('update:modelValue', color)"
          class="w-8 h-8 rounded-lg border-2 border-white shadow-lg hover:scale-110 transition-transform duration-200"
          :style="{ backgroundColor: color }"
          :title="color"
        ></button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTheme } from '../composables/useTheme'

const { isDark } = useTheme()

interface Props {
  modelValue: string
  label: string
  showPresets?: boolean
  inputId?: string
}

interface Emits {
  (e: 'update:modelValue', value: string): void
}

withDefaults(defineProps<Props>(), {
  showPresets: true,
  inputId: () => `color-input-${Math.random().toString(36).substr(2, 9)}`
})

const emit = defineEmits<Emits>()

const presetColors = [
  '#3b82f6', '#8b5cf6', '#ec4899', '#ef4444',
  '#f59e0b', '#10b981', '#06b6d4', '#6366f1',
  '#84cc16', '#f97316', '#14b8a6', '#a855f7',
  '#22c55e', '#eab308', '#64748b', '#374151'
]

const handleColorInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target?.value) {
    emit('update:modelValue', target.value)
  }
}

const handleHexInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  let value = target.value
  
  // Ensure it starts with #
  if (!value.startsWith('#')) {
    value = '#' + value
  }
  
  // Validate hex format
  const hexRegex = /^#[0-9A-Fa-f]{6}$/
  if (hexRegex.test(value)) {
    emit('update:modelValue', value)
    target.value = value
  }
}
</script>
