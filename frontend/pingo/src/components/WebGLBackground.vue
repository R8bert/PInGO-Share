<template>
  <canvas
    ref="canvasRef"
    class="absolute inset-0 w-full h-full"
    :class="className"
  />
</template>

<script setup lang="ts">
import { useWebGLBackground } from '../composables/useWebGLBackground'

interface Props {
  hueShift?: number
  noiseIntensity?: number
  scanlineIntensity?: number
  speed?: number
  scanlineFrequency?: number
  warpAmount?: number
  resolutionScale?: number
  className?: string
}

const props = withDefaults(defineProps<Props>(), {
  hueShift: 0,
  noiseIntensity: 0.02,
  scanlineIntensity: 0.1,
  speed: 0.5,
  scanlineFrequency: 0.5,
  warpAmount: 0.1,
  resolutionScale: 1,
  className: ''
})

const { canvasRef } = useWebGLBackground({
  hueShift: props.hueShift,
  noiseIntensity: props.noiseIntensity,
  scanlineIntensity: props.scanlineIntensity,
  speed: props.speed,
  scanlineFrequency: props.scanlineFrequency,
  warpAmount: props.warpAmount,
  resolutionScale: props.resolutionScale
})
</script>
