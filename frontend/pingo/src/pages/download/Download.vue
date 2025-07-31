<template>
  <div class="min-h-screen transition-colors duration-300"
       :style="{ backgroundColor: isDark ? '#0a0a0a' : '#ffffff' }">
    <!-- Main Content -->
    <main class="relative">
      <!-- Hero Section -->
      <div class="relative overflow-hidden transition-colors duration-300"
           :style="{ backgroundColor: isDark ? '#111827' : '#f8fafc' }">
        <!-- Floating circles animation -->
        <div class="absolute inset-0 pointer-events-none">
          <div class="absolute top-1/4 left-1/4 w-96 h-96 rounded-full blur-3xl animate-float-slow"
               :style="{ backgroundColor: isDark ? '#1e3a8a20' : '#dbeafe20' }"></div>
          <div class="absolute bottom-1/4 right-1/4 w-80 h-80 rounded-full blur-3xl animate-float-slower"
               :style="{ backgroundColor: isDark ? '#7c3aed20' : '#f3e8ff20' }"></div>
        </div>
        
        <div class="relative max-w-screen-xl mx-auto px-6 lg:px-8 pt-20 pb-16">
          <div class="text-center">
            <h1 class="text-4xl font-bold mb-4 animate-fade-in transition-colors duration-300"
                :style="{ color: isDark ? '#f9fafb' : '#111827' }">
              Download <span class="text-transparent bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 bg-clip-text animate-rainbow">files</span>
            </h1>
            <p class="text-xl max-w-2xl mx-auto animate-fade-in-delay transition-colors duration-300"
               :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">
              Preview and download your shared files
            </p>
            <router-link 
              to="/" 
              class="inline-flex items-center mt-6 font-medium transition-colors hover:opacity-80"
              :style="{ color: isDark ? '#60a5fa' : '#2563eb' }"
            >
              ← Share new files
            </router-link>
          </div>
        </div>
      </div>

      <!-- Download Section -->
      <div class="relative -mt-8 max-w-6xl mx-auto px-6 lg:px-8 pb-32">
        <!-- Loading State -->
        <div v-if="loading" 
             class="rounded-2xl shadow-xl border p-12 text-center animate-pulse transition-colors duration-300"
             :style="{ 
               backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
               borderColor: isDark ? '#374151' : '#e5e7eb'
             }">
          <div class="w-16 h-16 rounded-full mx-auto mb-4 animate-spin transition-colors duration-300"
               :style="{ backgroundColor: isDark ? '#374151' : '#e5e7eb' }"></div>
          <p class="transition-colors duration-300"
             :style="{ color: isDark ? '#9ca3af' : '#4b5563' }">Loading files...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="error" 
             class="rounded-2xl shadow-xl border p-12 text-center animate-shake transition-colors duration-300"
             :style="{ 
               backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
               borderColor: isDark ? '#dc2626' : '#fca5a5'
             }">
          <div class="w-16 h-16 bg-red-100 dark:bg-red-900 rounded-full flex items-center justify-center mx-auto mb-4">
            <XMarkIcon class="w-8 h-8 text-red-600" />
          </div>
          <h3 class="text-2xl font-bold mb-2 transition-colors duration-300"
              :style="{ color: isDark ? '#fca5a5' : '#dc2626' }">Files not found</h3>
          <p class="mb-6 transition-colors duration-300"
             :style="{ color: isDark ? '#f87171' : '#dc2626' }">{{ error }}</p>
          <router-link 
            to="/" 
            class="inline-flex items-center px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-all duration-200 hover:scale-105"
          >
            Share new files
          </router-link>
        </div>

        <!-- Files Content -->
        <div v-else-if="files.length > 0" class="space-y-6">
          <!-- Uploader Information -->
          <div v-if="uploader" class="bg-white rounded-2xl shadow-xl border border-gray-200/50 p-6 animate-slide-up">
            <h3 class="text-lg font-semibold text-gray-900 mb-4">Shared by</h3>
            <div class="flex items-center space-x-4">
              <div class="w-12 h-12 rounded-full overflow-hidden bg-gradient-to-r from-blue-500 to-indigo-600 flex items-center justify-center transition-transform duration-300 hover:scale-150">
                <img 
                  v-if="uploader.avatar" 
                  :src="`http://localhost:8080${uploader.avatar}`" 
                  :alt="uploader.username"
                  class="w-full h-full object-cover"
                  @error="handleAvatarError"
                />
                <UserIcon v-else class="w-6 h-6 text-white" />
              </div>
              <div>
                <p class="font-medium text-gray-900">{{ uploader.username }}</p>
                <p class="text-sm text-gray-600">{{ uploader.email }}</p>
              </div>
            </div>
          </div>

          <!-- Files Header with Rainbow Animation -->
          <div class="bg-white rounded-2xl shadow-xl border border-gray-200/50 overflow-hidden animate-slide-up">
            <div class="bg-gradient-to-r from-red-500 via-yellow-500 via-green-500 via-blue-500 via-indigo-500 to-purple-500 bg-size-400 animate-rainbow p-6">
              <div class="flex items-center justify-between">
                <div>
                  <h2 class="text-2xl font-bold text-white">{{ files.length }} file{{ files.length > 1 ? 's' : '' }} shared</h2>
                  <p class="text-white/90">{{ formatTotalSize() }} total</p>
                </div>
                <button 
                  @click="downloadAll"
                  :disabled="downloadingAll"
                  class="bg-white/20 hover:bg-white/30 backdrop-blur-sm text-white px-6 py-3 rounded-xl font-semibold transition-all duration-300 hover:scale-105 disabled:opacity-50 hover:shadow-lg"
                >
                  <span v-if="downloadingAll" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Downloading...
                  </span>
                  <span v-else class="flex items-center">
                    <ArrowDownTrayIcon class="w-5 h-5 mr-2" />
                    Download All
                  </span>
                </button>
              </div>
            </div>
          </div>

          <!-- Files List -->
          <div class="grid grid-cols-1 gap-6">
            <div 
              v-for="(file, index) in files" 
              :key="index" 
              class="bg-white rounded-2xl shadow-xl border border-gray-200/50 overflow-hidden animate-slide-up"
              :style="{ animationDelay: `${index * 0.1}s` }"
            >
              <div class="p-6">
                <div class="flex items-center justify-between mb-4">
                  <div class="flex items-center space-x-4">
                    <div class="w-16 h-16 bg-blue-100 rounded-xl flex items-center justify-center">
                      <img :src="getFileIcon(file)" alt="File type" class="w-10 h-10" />
                    </div>
                    <div>
                      <h3 class="text-xl font-bold text-gray-900">{{ file.name }}</h3>
                      <p class="text-gray-600">{{ formatFileSize(file.size) }}</p>
                      <p class="text-sm text-gray-500">{{ getFileExtension(file).toUpperCase() }} file</p>
                    </div>
                  </div>
                  <div class="flex items-center space-x-3">
                    <button 
                      v-if="canPreview(file)" 
                      @click="togglePreview(index)"
                      class="group relative p-2 rounded-lg text-sm font-medium hover:opacity-90 transition-all duration-200 hover:scale-110"
                      :class="previewingFiles.has(index) ? 'bg-red-500 hover:bg-red-600 text-white' : 'bg-green-500 hover:bg-green-600 text-white'"
                      :title="previewingFiles.has(index) ? 'Hide Preview' : 'Show Preview'"
                    >
                      <EyeIcon v-if="!previewingFiles.has(index)" class="w-5 h-5" />
                      <EyeSlashIcon v-else class="w-5 h-5" />
                    </button>
                    <button 
                      @click="downloadFile(index)"
                      :disabled="downloadingFiles.has(index)"
                      class="group relative bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-all duration-200 hover:scale-105 disabled:opacity-50"
                    >
                      <span v-if="downloadingFiles.has(index)" class="flex items-center">
                        <svg class="animate-spin -ml-1 mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        Downloading...
                      </span>
                      <span v-else class="flex items-center">
                        <ArrowDownTrayIcon class="w-4 h-4 mr-2" />
                        Download
                      </span>
                    </button>
                  </div>
                </div>

                <!-- Preview Section -->
                <div v-if="previewingFiles.has(index)" class="mt-6 animate-slide-down">
                  <div class="border-t border-gray-200 pt-6">
                    <!-- Image Preview -->
                    <div v-if="isImage(file)" class="text-center">
                      <img 
                        :src="getFileUrl(file)" 
                        :alt="file.name"
                        class="max-w-full max-h-96 mx-auto rounded-lg shadow-md"
                        @error="handlePreviewError"
                      />
                    </div>

                    <!-- Video Preview -->
                    <div v-else-if="isVideo(file)" class="text-center">
                      <video 
                        controls 
                        class="max-w-full max-h-96 mx-auto rounded-lg shadow-md"
                        :src="getFileUrl(file)"
                        @error="handlePreviewError"
                      >
                        Your browser does not support the video tag.
                      </video>
                    </div>

                    <!-- PDF Preview -->
                    <div v-else-if="isPdf(file)" class="text-center">
                      <object 
                        :data="getFileUrl(file)" 
                        type="application/pdf" 
                        class="w-full h-96 rounded-lg shadow-md border"
                      >
                        <div class="flex items-center justify-center h-96 bg-gray-100 rounded-lg">
                          <div class="text-center">
                            <DocumentIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
                            <p class="text-gray-600">PDF preview not supported in your browser</p>
                            <button 
                              @click="downloadFile(index)"
                              class="mt-4 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-all duration-200"
                            >
                              Download to view
                            </button>
                          </div>
                        </div>
                      </object>
                    </div>

                    <!-- Text Preview -->
                    <div v-else-if="isText(file)" class="bg-gray-50 rounded-lg p-4">
                      <pre class="text-sm text-gray-800 whitespace-pre-wrap overflow-auto max-h-64">{{ textContent.get(index) || 'Loading...' }}</pre>
                    </div>

                    <!-- Audio Preview -->
                    <div v-else-if="isAudio(file)" class="text-center">
                      <audio 
                        controls 
                        class="w-full max-w-md mx-auto"
                        :src="getFileUrl(file)"
                      >
                        Your browser does not support the audio tag.
                      </audio>
                    </div>

                    <!-- Unsupported Preview -->
                    <div v-else class="text-center py-8">
                      <DocumentIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
                      <p class="text-gray-600">Preview not available for this file type</p>
                      <button 
                        @click="downloadFile(index)"
                        class="mt-4 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-all duration-200"
                      >
                        Download to view
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useTheme } from '../../composables/useTheme'
import axios from 'axios'
import { XMarkIcon, ArrowDownTrayIcon, DocumentIcon, EyeIcon, EyeSlashIcon, UserIcon } from '@heroicons/vue/24/outline'

const { isDark } = useTheme()

// Import icons
import fileMp4Icon from '../main/images/train/icons/file-mp4.png'
import filePdfIcon from '../main/images/train/icons/file-pdf.png'
import fileFolderIcon from '../main/images/train/icons/file-folder.png'
import filePngIcon from '../main/images/train/icons/file-png.png'
import fileJpgIcon from '../main/images/train/icons/file-jpg.png'
import fileDocxIcon from '../main/images/train/icons/file-docx.png'

interface FileInfo {
  name: string
  size: number
  url: string
}

interface UploaderInfo {
  username: string
  avatar: string
  email: string
}

// State
const route = useRoute()
const files = ref<FileInfo[]>([])
const uploader = ref<UploaderInfo | null>(null)
const loading = ref(true)
const error = ref('')
const logoPath = ref<string | null>(null)
const previewingFiles = ref<Set<number>>(new Set())
const downloadingFiles = ref<Set<number>>(new Set())
const downloadingAll = ref(false)
const textContent = ref<Map<number, string>>(new Map())

// Icon mapping
const iconMap: Record<string, string> = {
  pdf: filePdfIcon,
  png: filePngIcon,
  jpg: fileJpgIcon,
  jpeg: fileJpgIcon,
  docx: fileDocxIcon,
  mp4: fileMp4Icon,
  txt: fileDocxIcon,
  doc: fileDocxIcon,
}

// Helper functions
const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const formatTotalSize = () => {
  const total = files.value.reduce((sum, file) => sum + file.size, 0)
  return formatFileSize(total)
}

const getFileIcon = (file: FileInfo) => {
  const ext = getFileExtension(file)
  return iconMap[ext] || fileFolderIcon
}

const getFileExtension = (file: FileInfo) => {
  return file.name.split('.').pop()?.toLowerCase() || ''
}

const getFileUrl = (file: FileInfo) => {
  return `http://localhost:8080${file.url}`
}

const canPreview = (file: FileInfo) => {
  const ext = getFileExtension(file)
  return ['jpg', 'jpeg', 'png', 'gif', 'webp', 'mp4', 'webm', 'ogg', 'pdf', 'txt', 'mp3', 'wav', 'ogg'].includes(ext)
}

const isImage = (file: FileInfo) => {
  const ext = getFileExtension(file)
  return ['jpg', 'jpeg', 'png', 'gif', 'webp', 'svg', 'bmp'].includes(ext)
}

const isVideo = (file: FileInfo) => {
  const ext = getFileExtension(file)
  return ['mp4', 'webm', 'ogg', 'avi', 'mov'].includes(ext)
}

const isPdf = (file: FileInfo) => {
  return getFileExtension(file) === 'pdf'
}

const isText = (file: FileInfo) => {
  const ext = getFileExtension(file)
  return ['txt', 'md', 'json', 'xml', 'csv'].includes(ext)
}

const isAudio = (file: FileInfo) => {
  const ext = getFileExtension(file)
  return ['mp3', 'wav', 'ogg', 'aac', 'm4a'].includes(ext)
}

const togglePreview = async (index: number) => {
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
  } else {
    previewingFiles.value.add(index)
    
    // Load text content if it's a text file
    const file = files.value[index]
    if (isText(file) && !textContent.value.has(index)) {
      try {
        const response = await axios.get(getFileUrl(file), { responseType: 'text' })
        textContent.value.set(index, response.data)
      } catch (err) {
        textContent.value.set(index, 'Error loading file content')
      }
    }
  }
}

const downloadFile = async (index: number) => {
  downloadingFiles.value.add(index)
  
  try {
    const file = files.value[index]
    const response = await axios.get(getFileUrl(file), {
      responseType: 'blob'
    })
    
    const blob = new Blob([response.data])
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = file.name
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
  } catch (err) {
    console.error('Download failed:', err)
  } finally {
    downloadingFiles.value.delete(index)
  }
}

const downloadAll = async () => {
  if (files.value.length === 1) {
    // If only one file, download it directly
    await downloadFile(0)
    return
  }

  downloadingAll.value = true
  
  try {
    const id = route.params.id as string
    const response = await axios.get(`http://localhost:8080/download/${id}`, {
      responseType: 'blob'
    })
    
    const blob = new Blob([response.data])
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = 'files.zip'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
  } catch (err) {
    console.error('Download all failed:', err)
  } finally {
    downloadingAll.value = false
  }
}

const handlePreviewError = () => {
  console.error('Failed to load preview')
}

const handleAvatarError = () => {
  console.error('Failed to load avatar')
}

const loadSettings = async () => {
  try {
    const response = await axios.get('http://localhost:8080/settings')
    logoPath.value = response.data.logo || null
  } catch (error) {
    console.error('Error loading settings:', error)
  }
}

const loadFiles = async () => {
  try {
    const id = route.params.id as string
    if (!id) {
      error.value = 'No file ID provided'
      return
    }

    // For now, we'll make a request to get file metadata
    // This will need to be implemented in the backend
    const response = await axios.get(`http://localhost:8080/files/${id}`)
    files.value = response.data.files || []
    uploader.value = response.data.uploader || null
    
    if (files.value.length === 0) {
      error.value = 'No files found or files have expired'
    }
  } catch (err: any) {
    console.error('Error loading files:', err)
    error.value = err.response?.data?.error || 'Files not found or have expired'
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  await loadSettings()
  await loadFiles()
})
</script>

<style scoped>
/* Reuse animations from Home.vue */
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

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  10%, 30%, 50%, 70%, 90% { transform: translateX(-2px); }
  20%, 40%, 60%, 80% { transform: translateX(2px); }
}

/* Animation Classes */
/* .animate-hover-zoom removed: replaced with Tailwind classes for hover zoom effect */


.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
  animation: fade-in 0.8s ease-out 0.2s both;
}

.animate-slide-up {
  animation: slide-up 0.6s ease-out both;
}

.animate-slide-down {
  animation: slide-down 0.4s ease-out;
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

.bg-size-400 {
  background-size: 400% 400%;
}

.animate-shake {
  animation: shake 0.5s ease-in-out;
}

/* Enhanced hover effects */
.hover\:scale-105:hover {
  transform: scale(1.05);
}

/* Custom focus states */
input:focus, button:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

/* Mobile Optimizations */
@media (max-width: 768px) {
  .text-4xl {
    font-size: 2rem;
  }
  
  .grid-cols-1 {
    grid-template-columns: 1fr;
  }
  
  .flex-col {
    flex-direction: column;
  }
  
  .space-x-3 > :not([hidden]) ~ :not([hidden]) {
    margin-left: 0;
    margin-top: 0.75rem;
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
