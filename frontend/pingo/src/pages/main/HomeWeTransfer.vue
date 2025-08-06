<template>
  <div class="min-h-screen bg-white dark:bg-gray-900">
    <!-- Main Container -->
    <div class="max-w-4xl mx-auto px-4 py-12">
      <!-- Header Section -->
      <div class="text-center mb-12">
        <h1 class="text-5xl md:text-6xl font-light text-gray-900 dark:text-white mb-6 leading-tight">
          Transfer up to 2GB free
        </h1>
        <p class="text-xl text-gray-600 dark:text-gray-300 font-light">
          No registration required
        </p>
      </div>

      <!-- Main Transfer Card -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 overflow-hidden">
        
        <!-- Upload Section -->
        <div class="p-8 md:p-12">
          <!-- Add Files Button/Zone -->
          <div 
            v-if="selectedFiles.length === 0"
            @dragover.prevent 
            @drop.prevent="onDrop" 
            @dragenter="isDragging = true" 
            @dragleave="isDragging = false"
            @click="triggerFileInput"
            :class="[
              'border-2 border-dashed rounded-lg p-16 text-center cursor-pointer transition-all duration-300',
              isDragging 
                ? 'border-blue-400 bg-blue-50 dark:bg-blue-900/20' 
                : 'border-gray-300 dark:border-gray-600 hover:border-gray-400 dark:hover:border-gray-500'
            ]"
          >
            <input type="file" ref="fileInput" @change="onFileChange" class="hidden" multiple />
            
            <!-- Plus Icon -->
            <div class="w-16 h-16 mx-auto mb-6 border-2 border-gray-300 dark:border-gray-600 rounded-full flex items-center justify-center">
              <svg class="w-8 h-8 text-gray-400 dark:text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
            </div>
            
            <h3 class="text-xl font-light text-gray-900 dark:text-white mb-2">Add your files</h3>
            <p class="text-gray-500 dark:text-gray-400">or drop them here</p>
          </div>

          <!-- Files List -->
          <div v-if="selectedFiles.length > 0" class="space-y-0 border border-gray-200 dark:border-gray-700 rounded-lg overflow-hidden">
            <div 
              v-for="(file, index) in selectedFiles" 
              :key="index" 
              class="flex items-center p-4 border-b border-gray-100 dark:border-gray-700 last:border-b-0 hover:bg-gray-50 dark:hover:bg-gray-700/50"
            >
              <!-- File Icon -->
              <div class="w-12 h-12 flex items-center justify-center bg-gray-100 dark:bg-gray-700 rounded mr-4">
                <svg class="w-6 h-6 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
              </div>
              
              <!-- File Info -->
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium text-gray-900 dark:text-white truncate">{{ file.name }}</p>
                <p class="text-xs text-gray-500 dark:text-gray-400">{{ formatFileSize(file.size) }}</p>
              </div>
              
              <!-- Remove Button -->
              <button 
                @click="removeFile(index)"
                class="ml-4 p-1 text-gray-400 hover:text-red-500 dark:hover:text-red-400"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
            
            <!-- Add More Files Button -->
            <div class="p-4 bg-gray-50 dark:bg-gray-800">
              <button 
                @click="triggerFileInput"
                class="flex items-center text-sm text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300"
              >
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                </svg>
                Add more files
              </button>
            </div>
          </div>
        </div>

        <!-- Transfer Settings -->
        <div v-if="selectedFiles.length > 0" class="border-t border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-800/50">
          <div class="p-8 md:p-12 space-y-8">
            
            <!-- Email to -->
            <div>
              <label class="block text-sm font-medium text-gray-900 dark:text-white mb-3">
                Email to
              </label>
              <div class="relative">
                <input 
                  v-model="emailTo"
                  type="email" 
                  placeholder="Enter email address"
                  class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
                <div v-if="emailTo" class="absolute right-3 top-3">
                  <button @click="emailTo = ''" class="text-gray-400 hover:text-gray-600 dark:hover:text-gray-300">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <!-- Email from -->
            <div>
              <label class="block text-sm font-medium text-gray-900 dark:text-white mb-3">
                Your email
              </label>
              <input 
                v-model="emailFrom"
                type="email" 
                placeholder="Enter your email address"
                class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400 focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <!-- Message -->
            <div>
              <label class="block text-sm font-medium text-gray-900 dark:text-white mb-3">
                Message
              </label>
              <textarea 
                v-model="transferMessage"
                placeholder="Add a message (optional)"
                rows="4"
                class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400 focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- Transfer Button -->
        <div v-if="selectedFiles.length > 0" class="border-t border-gray-200 dark:border-gray-700 p-8 md:p-12">
          <div class="flex items-center justify-between">
            <!-- File Summary -->
            <div class="text-sm text-gray-600 dark:text-gray-400">
              {{ selectedFiles.length }} file{{ selectedFiles.length > 1 ? 's' : '' }}, {{ formatTotalFileSize() }}
            </div>
            
            <!-- Transfer Button -->
            <button 
              @click="handleTransfer"
              :disabled="isUploading || getTotalFileSize() > maxUploadSize"
              :class="[
                'px-8 py-3 rounded-md font-medium transition-all duration-200',
                isUploading || getTotalFileSize() > maxUploadSize
                  ? 'bg-gray-300 dark:bg-gray-600 text-gray-500 dark:text-gray-400 cursor-not-allowed'
                  : 'bg-blue-600 hover:bg-blue-700 text-white shadow-sm'
              ]"
            >
              {{ isUploading ? 'Transferring...' : 'Transfer' }}
            </button>
          </div>

          <!-- Progress Bar -->
          <div v-if="progress > 0 && progress < 100" class="mt-6">
            <div class="flex justify-between text-xs text-gray-600 dark:text-gray-400 mb-2">
              <span>Uploading files...</span>
              <span>{{ progress }}%</span>
            </div>
            <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-1.5">
              <div 
                class="bg-blue-600 h-1.5 rounded-full transition-all duration-300 ease-out"
                :style="{ width: progress + '%' }"
              ></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Success State -->
      <div v-if="transferComplete" class="mt-8 bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-8 md:p-12">
        <div class="text-center mb-8">
          <!-- Success Icon -->
          <div class="w-16 h-16 bg-green-100 dark:bg-green-900/30 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
            </svg>
          </div>
          
          <h2 class="text-2xl font-light text-gray-900 dark:text-white mb-2">Transfer complete!</h2>
          <p class="text-gray-600 dark:text-gray-400">Your files have been uploaded successfully</p>
        </div>

        <!-- Download Link -->
        <div class="bg-gray-50 dark:bg-gray-800/50 rounded-lg p-6">
          <label class="block text-sm font-medium text-gray-900 dark:text-white mb-3">
            Download link
          </label>
          <div class="flex">
            <input 
              :value="downloadUrl"
              readonly
              class="flex-1 px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-l-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white text-sm"
            />
            <button 
              @click="copyToClipboard(downloadUrl)"
              class="px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-r-md border border-blue-600 font-medium text-sm"
            >
              Copy
            </button>
          </div>
        </div>

        <!-- Send Another -->
        <div class="mt-8 text-center">
          <button 
            @click="resetTransfer"
            class="text-blue-600 dark:text-blue-400 hover:text-blue-700 dark:hover:text-blue-300 font-medium"
          >
            Send another transfer
          </button>
        </div>
      </div>

      <!-- Error State -->
      <div v-if="errorMessage" class="mt-8 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg p-6">
        <div class="flex items-start">
          <div class="w-6 h-6 text-red-600 dark:text-red-400 mr-3 mt-0.5">
            <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-sm font-medium text-red-800 dark:text-red-200">Transfer failed</h3>
            <p class="text-sm text-red-700 dark:text-red-300 mt-1">{{ errorMessage }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- GitHub Icon - Bottom Left -->
    <a 
      href="https://github.com/R8bert/PInGO-Share" 
      target="_blank" 
      rel="noopener noreferrer" 
      class="fixed bottom-4 left-4 z-40 text-gray-400 hover:text-gray-600 dark:hover:text-gray-300"
      title="View GitHub Repository"
    >
      <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
      </svg>
    </a>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'

// State
const selectedFiles = ref<File[]>([])
const fileInput = ref<HTMLInputElement | null>(null)
const emailTo = ref('')
const emailFrom = ref('')
const transferMessage = ref('')
const downloadUrl = ref('')
const isUploading = ref(false)
const progress = ref(0)
const isDragging = ref(false)
const transferComplete = ref(false)
const errorMessage = ref('')
const maxUploadSize = ref<number>(2147483648) // 2GB like WeTransfer

// Helper functions
const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const getTotalFileSize = () => {
  return selectedFiles.value.reduce((total, file) => total + file.size, 0)
}

const formatTotalFileSize = () => {
  return formatFileSize(getTotalFileSize())
}

const triggerFileInput = () => {
  fileInput.value?.click()
}

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (target.files?.length) {
    selectedFiles.value = Array.from(target.files)
    isDragging.value = false
    errorMessage.value = ''
    transferComplete.value = false
  }
}

const onDrop = (e: DragEvent) => {
  if (e.dataTransfer?.files.length) {
    selectedFiles.value = Array.from(e.dataTransfer.files)
    isDragging.value = false
    errorMessage.value = ''
    transferComplete.value = false
  }
}

const removeFile = (index: number) => {
  selectedFiles.value.splice(index, 1)
}

const handleTransfer = async () => {
  if (selectedFiles.value.length === 0) return

  isUploading.value = true
  progress.value = 0
  errorMessage.value = ''
  transferComplete.value = false

  const formData = new FormData()
  selectedFiles.value.forEach((file) => {
    formData.append(`files`, file)
  })
  formData.append('email', emailTo.value)
  formData.append('senderEmail', emailFrom.value)
  formData.append('message', transferMessage.value)
  formData.append('validity', '7days') // Default validity

  try {
    const res = await axios.post('/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (event) => {
        if (event.total) progress.value = Math.round((event.loaded * 100) / event.total)
      },
    })

    const downloadId = res.data.download_url.split('/').pop()
    downloadUrl.value = `${window.location.origin}/download/${downloadId}`
    transferComplete.value = true
    progress.value = 100
  } catch (err: any) {
    console.error('Upload error:', err)
    const error = err.response?.data?.error || 'Transfer failed. Please try again.'
    errorMessage.value = error
  } finally {
    isUploading.value = false
  }
}

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    // Could add temporary success feedback here
  } catch (err) {
    console.error('Failed to copy text: ', err)
    errorMessage.value = 'Failed to copy link'
  }
}

const resetTransfer = () => {
  selectedFiles.value = []
  emailTo.value = ''
  emailFrom.value = ''
  transferMessage.value = ''
  downloadUrl.value = ''
  transferComplete.value = false
  errorMessage.value = ''
  progress.value = 0
  isUploading.value = false
  
  if (fileInput.value) {
    fileInput.value.value = ''
  }
}
</script>

<style scoped>
/* Smooth transitions for all interactive elements */
* {
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, opacity, box-shadow, transform, filter;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Custom focus states for better accessibility */
input:focus,
textarea:focus {
  outline: none;
}

/* Loading animation */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.animate-spin {
  animation: spin 1s linear infinite;
}

/* Hover effects */
button:disabled {
  cursor: not-allowed;
}

/* Mobile optimizations */
@media (max-width: 640px) {
  .text-6xl {
    font-size: 3rem;
    line-height: 1;
  }
  
  .px-16 {
    padding-left: 2rem;
    padding-right: 2rem;
  }
  
  .py-16 {
    padding-top: 3rem;
    padding-bottom: 3rem;
  }
}

/* Dark mode scrollbar */
.dark ::-webkit-scrollbar {
  width: 8px;
}

.dark ::-webkit-scrollbar-track {
  background: #1f2937;
}

.dark ::-webkit-scrollbar-thumb {
  background: #4b5563;
  border-radius: 4px;
}

.dark ::-webkit-scrollbar-thumb:hover {
  background: #6b7280;
}
</style>
