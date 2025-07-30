<template>
  <div class="min-h-screen bg-white">
    <!-- Main Content -->
    <main class="relative">
      <!-- Hero Section with background effect -->
      <div class="relative bg-gradient-to-b from-blue-50/30 to-white overflow-hidden">
        <!-- Floating circles animation -->
        <div class="absolute inset-0 pointer-events-none">
          <div class="absolute top-1/4 left-1/4 w-96 h-96 bg-blue-100/20 rounded-full blur-3xl animate-float-slow"></div>
          <div class="absolute bottom-1/4 right-1/4 w-80 h-80 bg-purple-100/20 rounded-full blur-3xl animate-float-slower"></div>
        </div>
        
        <div class="relative max-w-screen-xl mx-auto px-6 lg:px-8 pt-20 pb-32">
          <div class="text-center">
            <!-- Main headline with stagger animation -->
            <h1 class="text-5xl font-bold text-gray-900 mb-6 animate-fade-in">
              Share files <span class="text-transparent bg-gradient-to-r from-red-500 via-yellow-500 via-green-500 via-blue-500 via-indigo-500 to-purple-500 bg-clip-text animate-rainbow">simply</span>
            </h1>
            
            <!-- Subtitle with delay -->
            <p class="text-xl text-gray-600 max-w-2xl mx-auto animate-fade-in-delay">
              Transfer files of any size securely and quickly. No registration required.
            </p>
          </div>
        </div>
      </div>

      <!-- Upload Section -->
      <div class="relative -mt-16 max-w-4xl mx-auto px-6 lg:px-8 pb-32">
        <!-- Upload Area -->
        <div class="bg-white rounded-2xl shadow-xl border border-gray-200/50 overflow-hidden animate-slide-up">
          <!-- Upload Zone -->
          <div 
            @dragover.prevent 
            @drop.prevent="onDrop" 
            @dragenter="isDragging = true" 
            @dragleave="isDragging = false"
            @click="triggerFileInput"
            :class="[
              'relative border-2 border-dashed transition-all duration-300 cursor-pointer',
              isDragging ? 'border-blue-400 bg-blue-50 scale-105' : 'border-gray-300 hover:border-blue-400 hover:bg-gray-50'
            ]"
            class="p-12"
          >
            <input type="file" ref="fileInput" @change="onFileChange" class="hidden" multiple />
            
            <!-- Animated layer with icons at 45° -->
            <div v-if="selectedFiles.length > 0" class="absolute inset-0 -z-10 opacity-10 pointer-events-none">
              <div class="absolute -inset-[100%] origin-center rotate-45">
                <div class="w-full h-full animate-diagonal-scroll bg-repeat" :style="{ backgroundImage: `url(${getFileIcon(selectedFiles[0])})`, backgroundSize: '80px 80px' }" />
              </div>
            </div>
            
            <div class="text-center">
              <div class="mb-6">
                <CloudArrowUpIcon 
                  :class="[
                    'w-16 h-16 mx-auto transition-all duration-300',
                    isDragging ? 'text-blue-500 scale-110' : 'text-gray-400 animate-float'
                  ]" 
                />
              </div>
              
              <h3 class="text-2xl font-semibold text-gray-900 mb-2">
                {{ selectedFiles.length > 0 ? `${selectedFiles.length} file${selectedFiles.length > 1 ? 's' : ''} selected` : 'Drop your files here' }}
              </h3>
              
              <p class="text-gray-600 mb-6">
                {{ selectedFiles.length > 0 ? selectedFiles.map(f => f.name).join(', ') : 'or click to browse' }}
              </p>
              
              <button 
                v-if="selectedFiles.length === 0"
                class="inline-flex items-center px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-all duration-200 hover:scale-105"
              >
                Select files
              </button>
            </div>

            <!-- Drag overlay -->
            <div 
              v-if="isDragging" 
              class="absolute inset-0 bg-blue-100/50 rounded-lg flex items-center justify-center animate-pulse"
            >
              <div class="text-blue-600 font-semibold text-lg">Drop files here</div>
            </div>
          </div>

          <!-- Files List with Individual Previews -->
          <div v-if="selectedFiles.length > 0" class="space-y-4">
            <div v-for="(file, index) in selectedFiles" :key="index" class="border border-gray-200 rounded-lg p-4 bg-gray-50">
              <div class="flex items-center justify-between mb-3">
                <div class="flex items-center space-x-3">
                  <div class="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center">
                    <img :src="getFileIcon(file)" alt="File type" class="w-6 h-6" />
                  </div>
                  <div>
                    <h4 class="font-medium text-gray-900">{{ file.name }}</h4>
                    <p class="text-sm text-gray-600">{{ formatFileSize(file.size) }}</p>
                  </div>
                </div>
                <div class="flex items-center space-x-2">
                  <button 
                    v-if="['mp4', 'pdf', 'jpg', 'jpeg', 'png'].includes(getFileExtension(file))" 
                    @click="togglePreview(index)"
                    class="text-white px-3 py-1 rounded text-sm font-medium hover:opacity-90 transition-all duration-200"
                    :class="previewingFiles.has(index) ? 'bg-red-500 hover:bg-red-600' : 'bg-green-500 hover:bg-green-600'"
                    :disabled="isUploading">
                    {{ previewingFiles.has(index) ? 'Hide' : 'Preview' }}
                  </button>
                  <button 
                    @click="removeFile(index)"
                    class="p-1 text-gray-400 hover:text-red-600 transition-all duration-200"
                  >
                    <XMarkIcon class="w-4 h-4" />
                  </button>
                </div>
              </div>

              <!-- Individual Preview -->
              <div v-if="previewingFiles.has(index)" class="mt-3">
                <!-- Image preview for JPG/JPEG/PNG -->
                <div v-if="['jpg', 'jpeg', 'png'].includes(getFileExtension(file))" class="w-full">
                  <img 
                    :src="previewUrls.get(index)" 
                    :alt="file.name"
                    class="w-full max-w-md mx-auto rounded-lg shadow-md" 
                    @error="() => console.error('Image preview failed')">
                </div>

                <!-- Video preview for .mp4 -->
                <div v-else-if="getFileExtension(file) === 'mp4'" class="w-full">
                  <video 
                    :ref="(el) => setVideoRef(index, el)"
                    controls 
                    class="w-full max-w-md mx-auto rounded-lg shadow-md" 
                    :src="previewUrls.get(index)" 
                    @loadedmetadata="updateVideoMetadata">
                    Your browser does not support the video tag.
                  </video>
                </div>

                <!-- PDF preview for .pdf -->
                <div v-else-if="getFileExtension(file) === 'pdf'" class="w-full">
                  <object 
                    :data="previewUrls.get(index)" 
                    type="application/pdf" 
                    class="w-full max-w-md mx-auto h-64 rounded-lg shadow-md">
                    <p>Your browser does not support PDF preview. <a :href="previewUrls.get(index)" download>Download the PDF</a> to view it.</p>
                  </object>
                </div>
              </div>
            </div>
          </div>

                    <!-- File Details & Controls -->
          <div v-if="selectedFiles.length > 0" class="border-t border-gray-200 bg-gray-50/50 animate-slide-down">
            <div class="p-6">
              <!-- Total Files Summary -->
              <div class="flex items-center justify-between mb-6">
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center animate-scale-in">
                    <img :src="getFileIcon(selectedFiles[0])" alt="File type" class="w-8 h-8" />
                  </div>
                  <div>
                    <h4 class="font-medium text-gray-900 animate-fade-in">{{ selectedFiles.length }} file{{ selectedFiles.length > 1 ? 's' : '' }} selected</h4>
                    <p class="text-sm text-gray-600 animate-fade-in-delay">{{ formatTotalFileSize() }}</p>
                  </div>
                </div>
                <button 
                  @click="clearFile"
                  class="p-2 text-gray-400 hover:text-gray-600 transition-all duration-200 hover:scale-110"
                >
                  <XMarkIcon class="w-5 h-5" />
                </button>
              </div>

              <!-- Size Warning -->
              <div v-if="getTotalFileSize() > maxUploadSize" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg animate-pulse">
                <p class="text-red-800 text-sm font-medium">
                  Total file size exceeds the maximum limit of {{ formatFileSize(maxUploadSize) }}
                </p>
              </div>

              <!-- Email Input -->
              <div class="mb-6">
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Send to (optional)
                </label>
                <div class="relative">
                  <EnvelopeIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                  <input 
                    v-model="email"
                    type="email" 
                    placeholder="recipient@example.com"
                    class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200"
                  />
                </div>
              </div>

              <!-- Validity Selection -->
              <div class="mb-6" v-if="validityOptions.length > 0">
                <label class="block text-sm font-medium text-gray-700 mb-3">
                  File expiration
                </label>
                <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
                  <label
                    v-for="option in validityOptions"
                    :key="option.value"
                    :class="[
                      'relative flex items-center justify-center px-3 py-2 rounded-lg border-2 cursor-pointer transition-all text-sm',
                      selectedValidity === option.value
                        ? 'border-blue-500 bg-blue-50 text-blue-700'
                        : 'border-gray-200 bg-white text-gray-700 hover:border-blue-300'
                    ]"
                  >
                    <input
                      type="radio"
                      v-model="selectedValidity"
                      :value="option.value"
                      class="sr-only"
                    />
                    <div class="text-center">
                      <div class="font-medium">{{ option.label }}</div>
                      <div class="text-xs opacity-75">{{ option.description }}</div>
                    </div>
                  </label>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="flex space-x-3">
                <button 
                  @click="handleUpload"
                  :disabled="isUploading || selectedFiles.length === 0 || getTotalFileSize() > maxUploadSize"
                  :class="[
                    'flex-1 py-3 px-6 rounded-lg font-medium transition-all duration-200',
                    isUploading || selectedFiles.length === 0 || getTotalFileSize() > maxUploadSize
                      ? 'bg-gray-300 text-gray-500 cursor-not-allowed'
                      : 'bg-blue-600 hover:bg-blue-700 text-white hover:scale-105'
                  ]"
                >
                  <span v-if="isUploading" class="flex items-center justify-center">
                    <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Uploading...
                  </span>
                  <span v-else>Transfer</span>
                </button>
                
                <button 
                  @click="clearFile"
                  class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-all duration-200 hover:scale-105"
                >
                  Clear
                </button>
              </div>

              <!-- Progress Bar -->
              <div v-if="progress > 0 && progress < 100" class="mt-6">
                <div class="flex justify-between text-sm text-gray-600 mb-2">
                  <span>Uploading...</span>
                  <span>{{ progress }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2 overflow-hidden">
                  <div 
                    class="bg-blue-600 h-2 rounded-full transition-all duration-300 animate-pulse"
                    :style="{ width: progress + '%' }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Success Message - WeTransfer Style -->
        <div 
          v-if="message && message.type === 'success'" 
          class="mt-8 bg-white rounded-2xl shadow-xl border border-green-200 overflow-hidden animate-scale-in"
        >
          <div class="bg-gradient-to-r from-green-500 to-emerald-600 p-6">
            <div class="flex items-center space-x-4">
              <div class="w-12 h-12 bg-white/20 rounded-full flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
              </div>
              <div>
                <h3 class="text-xl font-bold text-white">Transfer complete! 🎉</h3>
                <p class="text-green-100">Your file has been uploaded successfully</p>
              </div>
            </div>
          </div>
          
          <div class="p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-600 mb-2">Share this link with your recipient:</p>
                <div class="bg-gray-50 rounded-lg p-3 font-mono text-sm text-gray-800 break-all">
                  {{ downloadUrl }}
                </div>
              </div>
              <div class="flex space-x-3 ml-6">
                <button 
                  @click="copyToClipboard(downloadUrl)"
                  class="group relative bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 text-white px-6 py-3 rounded-xl font-semibold transition-all duration-300 hover:scale-105 hover:shadow-lg transform active:scale-95"
                >
                  <span class="flex items-center">
                    <svg class="w-5 h-5 mr-2 transition-transform group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                    </svg>
                    Copy Link
                  </span>
                  <div class="absolute inset-0 rounded-xl bg-white opacity-0 group-hover:opacity-20 transition-opacity duration-300"></div>
                </button>
                <router-link 
                  v-if="downloadPath"
                  :to="downloadPath"
                  class="group relative bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 text-white px-6 py-3 rounded-xl font-semibold transition-all duration-300 hover:scale-105 hover:shadow-lg transform active:scale-95"
                >
                  <span class="flex items-center">
                    <svg class="w-5 h-5 mr-2 transition-transform group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                    </svg>
                    View Files
                  </span>
                  <div class="absolute inset-0 rounded-xl bg-white opacity-0 group-hover:opacity-20 transition-opacity duration-300"></div>
                </router-link>
              </div>
            </div>
          </div>
        </div>

        <!-- Error Message -->
        <div 
          v-if="message && message.type === 'error'" 
          class="mt-8 bg-red-50 border border-red-200 rounded-2xl p-6 animate-shake"
        >
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-red-100 rounded-full flex items-center justify-center">
              <XMarkIcon class="w-5 h-5 text-red-600" />
            </div>
            <div>
              <h3 class="font-bold text-red-900">Upload failed</h3>
              <p class="text-red-700">{{ message.text }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Features Section - WeTransfer Style -->
      <section class="bg-gray-50 py-24">
        <div class="max-w-screen-xl mx-auto px-6 lg:px-8">
          <div class="text-center mb-16">
            <h2 class="text-4xl font-bold text-gray-900 mb-4 animate-slide-up">
              Simple, secure file sharing
            </h2>
            <p class="text-xl text-gray-600 max-w-2xl mx-auto animate-slide-up-delay">
              Trusted by millions of users worldwide
            </p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
            <div class="text-center animate-slide-up">
              <div class="w-16 h-16 bg-blue-100 rounded-2xl flex items-center justify-center mx-auto mb-6 transform hover:scale-110 transition-transform duration-300">
                <svg class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900 mb-3">Secure transfers</h3>
              <p class="text-gray-600 leading-relaxed">End-to-end encrypted file transfers with automatic deletion after download.</p>
            </div>
            
            <div class="text-center animate-slide-up-delay">
              <div class="w-16 h-16 bg-purple-100 rounded-2xl flex items-center justify-center mx-auto mb-6 transform hover:scale-110 transition-transform duration-300">
                <svg class="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900 mb-3">Lightning fast</h3>
              <p class="text-gray-600 leading-relaxed">Upload files up to 2GB with optimized performance and global CDN delivery.</p>
            </div>
            
            <div class="text-center animate-slide-up-delay-2">
              <div class="w-16 h-16 bg-green-100 rounded-2xl flex items-center justify-center mx-auto mb-6 transform hover:scale-110 transition-transform duration-300">
                <svg class="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900 mb-3">Any file type</h3>
              <p class="text-gray-600 leading-relaxed">Send any file format - documents, images, videos, and more.</p>
            </div>
          </div>
        </div>
      </section>
    </main>
    <!-- WeTransfer-style Footer -->
    <footer class="bg-white border-t border-gray-100">
      <div class="max-w-screen-xl mx-auto px-6 lg:px-8 py-12">
        <div class="text-center">
          <div class="flex items-center justify-center mb-6">
            <img v-if="logoPath" :src="`http://localhost:8080${logoPath}`" class="h-8 w-auto mr-3" alt="Logo" @error="handleImageError" />
            <div class="text-2xl font-bold text-gray-900">PinGO</div>
          </div>
          <p class="text-gray-600 max-w-md mx-auto mb-8">
            Simple, secure file sharing for everyone. Made with ❤️ for seamless file transfers.
          </p>
          <div class="flex items-center justify-center space-x-6 text-sm text-gray-500">
            <a href="#" class="hover:text-gray-700 transition-colors">Privacy</a>
            <a href="#" class="hover:text-gray-700 transition-colors">Terms</a>
            <a href="#" class="hover:text-gray-700 transition-colors">Support</a>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuth } from '../../composables/useAuth'
import axios from 'axios'
import { CloudArrowUpIcon, EnvelopeIcon, XMarkIcon } from '@heroicons/vue/24/outline'

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

// Use auth composable
const { user, getSettings } = useAuth()

// State
const selectedFiles = ref<File[]>([])
const fileInput = ref<HTMLInputElement | null>(null)
const email = ref('')
const downloadUrl = ref('')
const downloadPath = ref('')
const isUploading = ref(false)
const progress = ref(0)
const isDragging = ref(false)
const message = ref<Message | null>(null)
const logoPath = ref<string | null>(null)
const maxUploadSize = ref<number>(104857600) // Default 100 MB
const previewingFiles = ref<Set<number>>(new Set())
const previewUrls = ref<Map<number, string>>(new Map())
const videoRefs = ref<Map<number, HTMLVideoElement>>(new Map())
const videoMetadata = ref<VideoMetadata>({ duration: 0 })

// Validity options and selected validity
const validityOptions = ref([
  { value: '7days', label: '7 Days', description: 'One week' },
  { value: '1month', label: '1 Month', description: '30 days' },
  { value: '6months', label: '6 Months', description: 'Half year' },
  { value: '1year', label: '1 Year', description: '12 months' },
  { value: 'never', label: 'Never', description: 'Permanent' }
])
const selectedValidity = ref('7days')
const maxAllowedValidity = ref('7days')

// Icon mapping
const iconMap: Record<string, string> = {
  pdf: filePdfIcon,
  png: filePngIcon,
  jpg: fileJpgIcon,
  jpeg: fileJpgIcon,
  docx: fileDocxIcon,
  mp4: fileMp4Icon,
}

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

const getFileIcon = (file: File) => {
  if (!file) return fileFolderIcon
  const ext = file.name.split('.').pop()?.toLowerCase() || ''
  return iconMap[ext] || fileFolderIcon
}

const getFileExtension = (file: File) => {
  return file?.name.split('.').pop()?.toLowerCase() || ''
}

const createPreviewUrl = (file: File, index: number) => {
  const url = URL.createObjectURL(file)
  previewUrls.value.set(index, url)
  return url
}

const togglePreview = (index: number) => {
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
    // Stop video if playing
    const videoEl = videoRefs.value.get(index)
    if (videoEl) {
      videoEl.pause()
      videoEl.currentTime = 0
    }
    // Clean up preview URL
    const url = previewUrls.value.get(index)
    if (url) {
      URL.revokeObjectURL(url)
      previewUrls.value.delete(index)
    }
  } else {
    previewingFiles.value.add(index)
    createPreviewUrl(selectedFiles.value[index], index)
  }
}

const setVideoRef = (index: number, el: any) => {
  if (el && el instanceof HTMLVideoElement) {
    videoRefs.value.set(index, el)
  }
}

const removeFile = (index: number) => {
  // Clean up preview if active
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
    const url = previewUrls.value.get(index)
    if (url) {
      URL.revokeObjectURL(url)
      previewUrls.value.delete(index)
    }
  }
  
  // Remove the file
  selectedFiles.value.splice(index, 1)
  
  // Reset all preview indices since array shifted
  const newPreviewingFiles = new Set<number>()
  const newPreviewUrls = new Map<number, string>()
  const newVideoRefs = new Map<number, HTMLVideoElement>()
  
  previewingFiles.value.forEach(oldIndex => {
    if (oldIndex > index) {
      newPreviewingFiles.add(oldIndex - 1)
      const url = previewUrls.value.get(oldIndex)
      if (url) newPreviewUrls.set(oldIndex - 1, url)
      const videoEl = videoRefs.value.get(oldIndex)
      if (videoEl) newVideoRefs.set(oldIndex - 1, videoEl)
    } else if (oldIndex < index) {
      newPreviewingFiles.add(oldIndex)
      const url = previewUrls.value.get(oldIndex)
      if (url) newPreviewUrls.set(oldIndex, url)
      const videoEl = videoRefs.value.get(oldIndex)
      if (videoEl) newVideoRefs.set(oldIndex, videoEl)
    }
  })
  
  previewingFiles.value = newPreviewingFiles
  previewUrls.value = newPreviewUrls
  videoRefs.value = newVideoRefs
}

const updateVideoMetadata = () => {
  // This function can be expanded if needed for multiple video metadata tracking
  console.log('Video metadata updated')
}

const triggerFileInput = () => {
  fileInput.value?.click()
}

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (target.files?.length) {
    selectedFiles.value = Array.from(target.files)
    isDragging.value = false
    message.value = null
    videoMetadata.value.duration = 0
  }
}

const onDrop = (e: DragEvent) => {
  if (e.dataTransfer?.files.length) {
    selectedFiles.value = Array.from(e.dataTransfer.files)
    isDragging.value = false
    message.value = null
    videoMetadata.value.duration = 0
  }
}

const handleUpload = async () => {
  if (selectedFiles.value.length === 0) return

  isUploading.value = true
  progress.value = 0
  message.value = null

  const formData = new FormData()
  selectedFiles.value.forEach((file) => {
    formData.append(`files`, file)
  })
  formData.append('email', email.value)
  formData.append('validity', selectedValidity.value)

  try {
    const res = await axios.post('/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (event) => {
        if (event.total) progress.value = Math.round((event.loaded * 100) / event.total)
      },
    })

    const downloadId = res.data.download_url.split('/').pop()
    downloadUrl.value = `${window.location.origin}/download/${downloadId}`
    downloadPath.value = `/download/${downloadId}`
    message.value = { text: `${selectedFiles.value.length} file${selectedFiles.value.length > 1 ? 's' : ''} uploaded successfully!`, type: 'success' }
  } catch (err: any) {
    console.error('Upload error:', err)
    const errorMessage = err.response?.data?.error || 'Upload failed. Please try again.'
    message.value = { text: errorMessage, type: 'error' }
  } finally {
    isUploading.value = false
  }
}

const clearFile = () => {
  // Clean up all preview URLs
  previewUrls.value.forEach(url => URL.revokeObjectURL(url))
  
  // Stop all videos
  videoRefs.value.forEach(video => {
    video.pause()
    video.currentTime = 0
  })
  
  selectedFiles.value = []
  isDragging.value = false
  message.value = null
  downloadUrl.value = ''
  downloadPath.value = ''
  progress.value = 0
  previewingFiles.value.clear()
  previewUrls.value.clear()
  videoRefs.value.clear()
  videoMetadata.value.duration = 0
  
  if (fileInput.value) {
    fileInput.value.value = ''
  }
}

const handleImageError = () => {
  console.error('Logo image failed to load:', logoPath.value)
  logoPath.value = null
}

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    message.value = { text: 'Link copied to clipboard!', type: 'success' }
    setTimeout(() => {
      message.value = null
    }, 3000)
  } catch (err) {
    console.error('Failed to copy text: ', err)
    message.value = { text: 'Failed to copy link', type: 'error' }
    setTimeout(() => {
      message.value = null
    }, 3000)
  }
}

const loadSettings = async () => {
  try {
    const settings = await getSettings()
    logoPath.value = settings.logo || null
    maxUploadSize.value = settings.maxUploadSize || 104857600
    maxAllowedValidity.value = settings.maxValidity || '7days'
    
    // Filter validity options based on max allowed
    const validityOrder = ['7days', '1month', '6months', '1year', 'never']
    const maxIndex = validityOrder.indexOf(maxAllowedValidity.value)
    
    if (maxIndex !== -1) {
      const allowedOptions = validityOrder.slice(0, maxIndex + 1)
      validityOptions.value = validityOptions.value.filter(option => 
        allowedOptions.includes(option.value)
      )
    }
    
    // Set default validity to first available option
    if (validityOptions.value.length > 0) {
      selectedValidity.value = validityOptions.value[0].value
    }
  } catch (error) {
    console.error('Error loading settings:', error)
  }
}

onMounted(() => {
  loadSettings()
})
</script>

<style scoped>
/* Optimized Animations */
@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes slide-up {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes slide-down {
  from { opacity: 0; transform: translateY(-20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes scale-in {
  from { opacity: 0; transform: scale(0.8); }
  to { opacity: 1; transform: scale(1); }
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

@keyframes float-delayed {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-15px); }
}

@keyframes float-slow {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(2deg); }
}

@keyframes float-slower {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-15px) rotate(-1deg); }
}

@keyframes rainbow {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* Animation Classes */
.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
  animation: fade-in 0.8s ease-out 0.2s both;
}

.animate-slide-up-delay {
  animation: slide-up 0.6s ease-out 0.3s both;
}

.animate-slide-up-delay-2 {
  animation: slide-up 0.6s ease-out 0.6s both;
}

.animate-slide-down {
  animation: slide-down 0.4s ease-out;
}

.animate-scale-in {
  animation: scale-in 0.3s ease-out;
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

.animate-float-delayed {
  animation: float-delayed 8s ease-in-out infinite;
}

.animate-float-slow {
  animation: float-slow 8s ease-in-out infinite;
}

.animate-float-slower {
  animation: float-slower 10s ease-in-out infinite;
}

.animate-rainbow {
  background-size: 400% 400%;
  animation: rainbow 3s ease-in-out infinite;
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Enhanced hover effects */
.hover\:scale-105:hover {
  transform: scale(1.05);
}

.hover\:scale-110:hover {
  transform: scale(1.10);
}

/* Continuous diagonal scroll effect */
@keyframes diagonal-scroll { 
  0% { background-position: 0 0; } 
  100% { background-position: -240px 240px; } 
}
.animate-diagonal-scroll { 
  animation: diagonal-scroll 10s linear infinite; 
}

/* Progress bar pulse */
.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

/* Custom focus states */
input:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

/* Mobile Optimizations */
@media (max-width: 768px) {
  .text-5xl {
    font-size: 2.5rem;
  }
}

/* Reduced Motion Support */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
</style>
