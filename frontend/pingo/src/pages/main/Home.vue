<template>
  <div class="min-h-screen relative overflow-hidden" :style="{ '--button-from': buttonFromColor, '--button-to': buttonToColor }">
    <!-- Dynamic background with overlay -->
    <div class="absolute inset-0 z-0">
      <img 
        ref="backgroundImage"
        src="https://4kwallpapers.com/images/wallpapers/poppy-flower-black-3840x2160-17227.jpg" 
        class="w-full h-full object-cover"
        alt="Background"
      />
      <div class="absolute inset-0 bg-gradient-to-br from-blue-900/70 to-purple-900/70"></div>
    </div>

    <!-- Main content -->
    <div class="relative z-10 min-h-screen px-4 py-8">
      <div class="max-w-7xl mx-auto">
        <!-- Upload box (left side) -->
        <div class="w-full sm:w-1/2 md:w-1/3 bg-white/90 backdrop-blur-lg shadow-xl rounded-xl p-6 space-y-6 animate-fade-in relative overflow-hidden">
          
          <!-- Animated layer with icons at 45° -->
          <div
            v-if="selectedFile"
            class="absolute inset-0 -z-10 opacity-10 pointer-events-none"
          >
            <div class="absolute -inset-[100%] origin-center rotate-45">
              <div
                class="w-full h-full animate-diagonal-scroll bg-repeat"
                :style="{ backgroundImage: `url(${getFileIcon()})`, backgroundSize: '80px 80px' }"
              />
            </div>
          </div>

          <h1 class="text-2xl font-extrabold text-gray-900 tracking-tight relative z-10">
            Share Files Effortlessly
          </h1>
          <p class="text-gray-500 text-sm relative z-10">
            Upload files securely with a single click or drag-and-drop
          </p>

          <!-- Drag-and-drop zone -->
          <div
            @dragover.prevent
            @drop.prevent="onDrop"
            @dragenter="isDragging = true"
            @dragleave="isDragging = false"
            class="relative border-2 border-dashed border-gray-300 p-6 rounded-xl text-center cursor-pointer transition-all duration-300 hover:border-blue-400 hover:bg-blue-50/50"
            :class="{ 'border-blue-400 bg-blue-50/50 scale-105': isDragging }"
            @click="triggerFileInput"
          >
            <input
              type="file"
              ref="fileInput"
              @change="onFileChange"
              class="hidden"
              aria-label="Select file to upload"
            />
            <div class="flex flex-col items-center space-y-2">
              <CloudArrowUpIcon
                class="w-10 h-10 text-gray-400 transition-colors duration-300"
                :class="{ 'text-blue-500': isDragging }"
              />
              <p class="text-sm font-medium text-gray-600">
                <span class="text-blue-500 underline hover:text-blue-600">Choose a file</span>
                or drag it here
              </p>
              <p
                v-if="selectedFile"
                class="mt-1 text-sm font-semibold text-gray-800 truncate max-w-xs"
              >
                {{ selectedFile.name }}
              </p>
            </div>
            <div
              v-if="isDragging"
              class="absolute inset-0 rounded-xl bg-blue-100 opacity-20 animate-pulse"
            />
          </div>

          <!-- Preview section -->
          <div v-if="selectedFile" class="space-y-4">
            <button
              v-if="['mp4', 'pdf'].includes(getFileExtension())"
              @click="togglePreview"
              :class="`w-full text-white py-2 rounded-lg text-sm font-semibold hover:opacity-90 transition-all duration-200 ${isPreviewing ? 'bg-red-500 hover:bg-red-600' : 'bg-green-500 hover:bg-green-600'}`"
              :disabled="isUploading"
              aria-label="Toggle preview"
            >
              {{ isPreviewing ? 'Hide Preview' : 'Show Preview' }}
            </button>

            <!-- Video preview for .mp4 -->
            <div v-if="isPreviewing && getFileExtension() === 'mp4'" class="w-full">
              <video
                ref="videoPreview"
                controls
                class="w-full rounded-lg shadow-md"
                :src="previewUrl"
                @loadedmetadata="updateVideoMetadata"
              >
                Your browser does not support the video tag.
              </video>
            </div>

            <!-- PDF preview for .pdf -->
            <div v-if="isPreviewing && getFileExtension() === 'pdf'" class="w-full">
              <object
                :data="previewUrl"
                type="application/pdf"
                class="w-full h-96 rounded-lg shadow-md"
              >
                <p>Your browser does not support PDF preview. <a :href="previewUrl" download>Download the PDF</a> to view it.</p>
              </object>
            </div>
          </div>

          <!-- Email field -->
          <div class="relative">
            <input
              v-model="email"
              type="email"
              placeholder="Recipient's email (optional)"
              class="w-full border border-gray-300 px-3 py-2 rounded-lg text-sm focus:ring-2 focus:ring-blue-400 focus:border-blue-400 outline-none transition-all duration-200"
              aria-label="Recipient's email"
            />
            <EnvelopeIcon
              class="absolute right-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
            />
          </div>

          <!-- Upload button -->
          <button
            :disabled="!selectedFile || isUploading"
            @click="handleUpload"
            class="w-full bg-gradient-to-r from-[var(--button-from)] to-[var(--button-to)] text-white py-2.5 rounded-lg text-sm font-semibold hover:opacity-90 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300 transform hover:-translate-y-0.5"
            aria-label="Upload file"
          >
            {{ isUploading ? 'Uploading...' : 'Share Now' }}
          </button>

          <!-- Progress bar -->
          <div
            v-if="progress > 0 && progress < 100"
            class="w-full bg-gray-100 rounded-full h-2 overflow-hidden"
          >
            <div
              class="bg-gradient-to-r from-[var(--button-from)] to-[var(--button-to)] h-full transition-all duration-300"
              :style="{ width: progress + '%' }"
            />
          </div>
        </div>

        <!-- File info section (right side) -->
        <div v-if="selectedFile" class="mt-6 sm:mt-0 sm:absolute sm:top-8 sm:right-8 w-full sm:w-1/2 md:w-2/5 bg-white/90 backdrop-blur-lg shadow-xl rounded-xl p-6 space-y-6">
          <h2 class="text-xl font-semibold text-gray-900">Share Details</h2>
          <div class="space-y-4">
            <div>
              <p class="text-gray-500 text-sm">File Name:</p>
              <p class="text-gray-800 font-medium truncate">{{ selectedFile.name }}</p>
            </div>
            <div>
              <p class="text-gray-500 text-sm">Size:</p>
              <p class="text-gray-800 font-medium">{{ formatFileSize(selectedFile.size) }}</p>
            </div>
            <div>
              <p class="text-gray-500 text-sm">Type:</p>
              <p class="text-gray-800 font-medium">{{ selectedFile.type || 'Unknown' }}</p>
            </div>
            <div>
              <p class="text-gray-500 text-sm">Extension:</p>
              <p class="text-gray-800 font-medium">{{ getFileExtension() }}</p>
            </div>
            <div>
              <p class="text-gray-500 text-sm">Last Modified:</p>
              <p class="text-gray-800 font-medium">{{ new Date(selectedFile.lastModified).toLocaleString() }}</p>
            </div>
            <div>
              <p class="text-gray-500 text-sm">Creation Date:</p>
              <p class="text-gray-800 font-medium">{{ new Date(selectedFile.lastModified).toLocaleString() }}</p> <!-- Fallback to lastModified -->
            </div>
            <div v-if="videoMetadata.duration">
              <p class="text-gray-500 text-sm">Duration:</p>
              <p class="text-gray-800 font-medium">{{ formatDuration(videoMetadata.duration) }}</p>
            </div>
            <div v-if="email">
              <p class="text-gray-500 text-sm">Recipient:</p>
              <p class="text-gray-800 font-medium">{{ email }}</p>
            </div>
          </div>
          <div v-if="message" class="text-center text-sm animate-slide-up">
            <p :class="message.type === 'success' ? 'text-green-600' : 'text-red-600'">
              {{ message.text }}
            </p>
            <a
              v-if="message.type === 'success' && downloadUrl"
              :href="downloadUrl"
              target="_blank"
              class="text-blue-500 underline hover:text-blue-600 text-xs break-all"
            >
              {{ downloadUrl }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import axios from 'axios'
import { CloudArrowUpIcon, EnvelopeIcon } from '@heroicons/vue/24/outline'
import ColorThief from 'colorthief'

// Import icons
import fileMp4Icon from './images/train/icons/file-mp4.png'
import filePdfIcon from './images/train/icons/file-pdf.png'
import fileFolderIcon from './images/train/icons/file-folder.png'
import filePngIcon from './images/train/icons/file-png.png'
import fileJpgIcon from './images/train/icons/file-jpg.png'
import fileDocxIcon from './images/train/icons/file-docx.png'

interface Message {
  text: string
  type: 'success' | 'error'
}

interface VideoMetadata {
  duration: number
}

const selectedFile = ref<File | null>(null)
const fileInput = ref<HTMLInputElement | null>(null)
const email = ref('')
const downloadUrl = ref('')
const isUploading = ref(false)
const progress = ref(0)
const isDragging = ref(false)
const message = ref<Message | null>(null)
const useFolderIcon = ref(false)
const isPreviewing = ref(false)
const previewUrl = ref<string | undefined>(undefined)
const videoPreview = ref<HTMLVideoElement | null>(null)
const backgroundImage = ref<HTMLImageElement | null>(null)
const buttonFromColor = ref<string>('#3b82f6') // Default blue-500
const buttonToColor = ref<string>('#6366f1')   // Default indigo-600
const videoMetadata = ref<VideoMetadata>({ duration: 0 })

const iconMap: Record<string, string> = {
  pdf: filePdfIcon,
  png: filePngIcon,
  jpg: fileJpgIcon,
  jpeg: fileJpgIcon,
  docx: fileDocxIcon,
  mp4: fileMp4Icon,
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const formatDuration = (seconds: number) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = Math.floor(seconds % 60)
  return `${minutes}:${remainingSeconds < 10 ? '0' : ''}${remainingSeconds}`
}

const triggerFileInput = () => {
  fileInput.value?.click()
}

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (target.files?.length) {
    selectedFile.value = target.files[0]
    useFolderIcon.value = false
    isDragging.value = false
    message.value = null
    updatePreviewUrl()
    videoMetadata.value.duration = 0 // Reset duration
  }
}

const onDrop = (e: DragEvent) => {
  if (e.dataTransfer?.files.length) {
    selectedFile.value = e.dataTransfer.files[0]
    useFolderIcon.value = false
    isDragging.value = false
    message.value = null
    updatePreviewUrl()
    videoMetadata.value.duration = 0 // Reset duration
  }
}

const handleUpload = async () => {
  if (!selectedFile.value) return

  isUploading.value = true
  progress.value = 0
  message.value = null

  const formData = new FormData()
  formData.append('file', selectedFile.value)
  formData.append('email', email.value)

  try {
    const res = await axios.post('http://localhost:8080/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
      onUploadProgress: (event) => {
        if (event.total) {
          progress.value = Math.round((event.loaded * 100) / event.total)
        }
      },
    })

    downloadUrl.value = `http://localhost:8080${res.data.download_url}`
    message.value = {
      text: 'File uploaded successfully!',
      type: 'success',
    }
  } catch (err) {
    console.error('Upload error:', err)
    message.value = {
      text: 'Upload failed. Please try again.',
      type: 'error',
    }
  } finally {
    isUploading.value = false
  }
}

const getFileIcon = () => {
  if (!selectedFile.value) return ''
  const ext = selectedFile.value.name.split('.').pop()?.toLowerCase() || ''
  return useFolderIcon.value ? fileFolderIcon : (iconMap[ext] || fileFolderIcon)
}

const getFileExtension = () => {
  return selectedFile.value?.name.split('.').pop()?.toLowerCase() || ''
}

const updatePreviewUrl = () => {
  if (selectedFile.value) {
    previewUrl.value = URL.createObjectURL(selectedFile.value)
  } else {
    previewUrl.value = undefined
  }
}

const togglePreview = () => {
  isPreviewing.value = !isPreviewing.value
  if (isPreviewing.value && videoPreview.value) {
    videoPreview.value.play().catch(err => console.error('Autoplay blocked:', err))
  } else if (!isPreviewing.value && videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
}

const updateVideoMetadata = () => {
  if (videoPreview.value && getFileExtension() === 'mp4') {
    videoMetadata.value.duration = videoPreview.value.duration
  }
}

// Cleanup object URL on unmount or file change
watch(selectedFile, () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    updatePreviewUrl()
  }
  if (!isPreviewing.value && videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
  videoMetadata.value.duration = 0 // Reset duration on file change
})

onUnmounted(() => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
  }
  if (videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
})

// Dynamic color adaptation based on background
onMounted(async () => {
  if (backgroundImage.value) {
    await backgroundImage.value.decode()
    const colorThief = new ColorThief()
    const [r, g, b] = colorThief.getColor(backgroundImage.value)
    const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255
    const isDark = luminance < 0.5

    const hsl = rgbToHsl(r, g, b)
    let complementaryHue = (hsl.h + 180) % 360
    const saturation = hsl.s
    const lightness = isDark ? Math.min(0.7, hsl.l + 0.2) : Math.max(0.3, hsl.l - 0.2)

    const fromColor = hslToRgb(complementaryHue, saturation, lightness)
    const toColor = hslToRgb((complementaryHue + 30) % 360, saturation, lightness)
    buttonFromColor.value = `rgb(${fromColor.r}, ${fromColor.g}, ${fromColor.b})`
    buttonToColor.value = `rgb(${toColor.r}, ${toColor.g}, ${toColor.b})`
  }
})

// Helper functions for color conversion
function rgbToHsl(r: number, g: number, b: number) {
  r /= 255; g /= 255; b /= 255
  const max = Math.max(r, g, b), min = Math.min(r, g, b)
  let h, s, l = (max + min) / 2

  if (max === min) {
    h = s = 0
  } else {
    const d = max - min
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min)
    switch (max) {
      case r: h = (g - b) / d + (g < b ? 6 : 0); break
      case g: h = (b - r) / d + 2; break
      case b: h = (r - g) / d + 4; break
    }
    h /= 6
  }

  return { h: h * 360, s, l }
}

function hslToRgb(h: number, s: number, l: number) {
  h /= 360; s /= 100; l /= 100
  let r, g, b
  if (s === 0) {
    r = g = b = l
  } else {
    const hue2rgb = (p: number, q: number, t: number) => {
      if (t < 0) t += 1
      if (t > 1) t -= 1
      if (t < 1/6) return p + (q - p) * 6 * t
      if (t < 1/2) return q
      if (t < 2/3) return p + (q - p) * (2/3 - t) * 6
      return p
    }
    const q = l < 0.5 ? l * (1 + s) : l + s - l * s
    const p = 2 * l - q
    r = hue2rgb(p, q, h + 1/3)
    g = hue2rgb(p, q, h)
    b = hue2rgb(p, q, h - 1/3)
  }
  return { r: Math.round(r * 255), g: Math.round(g * 255), b: Math.round(b * 255) }
}
</script>

<style scoped>
/* Fade in and slide up animations */
.animate-fade-in {
  animation: fadeIn 0.6s ease-in-out;
}
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(15px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-slide-up {
  animation: slideUp 0.4s ease-in-out;
}
@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(15px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Continuous diagonal scroll effect */
@keyframes diagonal-scroll {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: -240px 240px;
  }
}
.animate-diagonal-scroll {
  animation: diagonal-scroll 10s linear infinite;
}

/* Ensure overlay covers the entire dialog */
.absolute.inset-0 {
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

/* Styles for icons and overlay */
.bg-repeat {
  background-repeat: repeat;
}

/* Extended container to cover corners when rotated */
.absolute.inset-100p {
  top: -100%;
  left: -100%;
  right: -100%;
  bottom: -100%;
}

/* Responsive adjustments */
@media (max-width: 640px) {
  .sm\:absolute {
    position: relative;
    top: auto;
    right: auto;
    width: 100%;
    margin-top: 1.5rem;
  }
}

/* Dynamic button styles */
button {
  background: var(--button-from);
  background: linear-gradient(to right, var(--button-from), var(--button-to));
}
button:hover {
  opacity: 0.9;
}
</style>