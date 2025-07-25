<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-100 via-indigo-50 to-purple-100 px-4 py-8">
    <div class="max-w-7xl mx-auto flex items-start justify-between gap-8">
      <!-- Panou de upload (stânga) -->
      <div class="w-1/3 bg-white/90 backdrop-blur-lg shadow-xl rounded-xl p-6 space-y-6 animate-fade-in relative overflow-hidden">
        
        <!-- Layer animat cu iconițe la 45° -->
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

        <!-- Zona de drag-and-drop -->
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

        <!-- Câmp email -->
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

        <!-- Buton de upload -->
        <button
          :disabled="!selectedFile || isUploading"
          @click="handleUpload"
          class="w-full bg-gradient-to-r from-blue-500 to-indigo-600 text-white py-2.5 rounded-lg text-sm font-semibold hover:from-blue-600 hover:to-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300 transform hover:-translate-y-0.5"
          aria-label="Upload file"
        >
          {{ isUploading ? 'Uploading...' : 'Share Now' }}
        </button>

        <!-- Bară de progres -->
        <div
          v-if="progress > 0 && progress < 100"
          class="w-full bg-gray-100 rounded-full h-2 overflow-hidden"
        >
          <div
            class="bg-gradient-to-r from-blue-500 to-indigo-600 h-full transition-all duration-300"
            :style="{ width: progress + '%' }"
          />
        </div>
      </div>

      <!-- Panou dreapta (rezervat pentru extensii) -->
      <div class="w-2/3 bg-white/90 backdrop-blur-lg shadow-xl rounded-xl p-6 space-y-6">
        <h2 class="text-xl font-semibold text-gray-900">Share Details</h2>
        <p class="text-gray-500 text-sm" v-if="selectedFile">
          Selected file: <strong>{{ selectedFile.name }}</strong><br>
          Size: {{ (selectedFile.size / 1024 / 1024).toFixed(2) }} MB
        </p>
        <p class="text-gray-500 text-sm" v-else>
          No file selected. Drag or click to upload a file.
        </p>
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
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { CloudArrowUpIcon, EnvelopeIcon } from '@heroicons/vue/24/outline'

// ✅ Importuri corecte din folderul local
import filePdfIcon from './images/train/icons/file-pdf.png'
import fileFolderIcon from './images/train/icons/file-folder.png'
import filePngIcon from './images/train/icons/file-png.png'
import fileJpgIcon from './images/train/icons/file-jpg.png'
import fileDocxIcon from './images/train/icons/file-docx.png'

interface Message {
  text: string
  type: 'success' | 'error'
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

const iconMap: Record<string, string> = {
  pdf: filePdfIcon,
  png: filePngIcon,
  jpg: fileJpgIcon,
  jpeg: fileJpgIcon,
  docx: fileDocxIcon,
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
  }
}

const onDrop = (e: DragEvent) => {
  if (e.dataTransfer?.files.length) {
    selectedFile.value = e.dataTransfer.files[0]
    useFolderIcon.value = false
    isDragging.value = false
    message.value = null
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
</script>

<style scoped>
/* Fade in și slide up */
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

/* Efect de scroll diagonal continuu */
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

/* Asigurăm că overlay-ul acoperă întregul dialog */
.absolute.inset-0 {
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

/* Stiluri pentru iconițe și overlay */
.bg-repeat {
  background-repeat: repeat;
}

/* Container extins pentru a acoperi colțurile la rotire */
.absolute.inset-100p {
  top: -100%;
  left: -100%;
  right: -100%;
  bottom: -100%;
}
</style>