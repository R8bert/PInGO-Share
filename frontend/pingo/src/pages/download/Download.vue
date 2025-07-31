<template>
  <div class="min-h-screen flex flex-col transition-colors duration-300" 
       :style="{ backgroundColor: isDark ? '#000000' : '#f8fafc' }">
    <!-- Main Content -->
    <main class="relative">
      <!-- Hero Section with background effect -->
      <div class="relative overflow-hidden transition-colors duration-300 pb-16">
        <!-- Floating circles animation (simpler version) -->
        <div class="absolute inset-0 pointer-events-none">
          <div class="absolute top-1/4 left-1/4 w-72 h-72 rounded-full blur-3xl animate-float-slow"
               :style="{ backgroundColor: isDark ? '#1e3a8a20' : '#dbeafe20' }"></div>
          <div class="absolute bottom-1/4 right-1/4 w-64 h-64 rounded-full blur-3xl animate-float-slower"
               :style="{ backgroundColor: isDark ? '#7c3aed20' : '#f3e8ff20' }"></div>
        </div>
        
        <div class="relative max-w-screen-xl mx-auto px-6 lg:px-8 pt-20 pb-16">
          <div class="text-center">
            <!-- Main headline with stagger animation -->
            <h1 class="text-5xl font-bold mb-6 animate-fade-in transition-colors duration-300" 
                :style="{ color: isDark ? '#ffffff' : '#000000' }">
              Download <span class="text-transparent bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 bg-clip-text animate-rainbow">files</span>
            </h1>
            
            <!-- Subtitle with delay -->
            <p class="text-xl max-w-2xl mx-auto animate-fade-in-delay transition-colors duration-300"
               :style="{ color: isDark ? '#cccccc' : '#666666' }">
              {{ loading ? 'Loading your files...' : `${files.length} file${files.length !== 1 ? 's' : ''} ready for download` }}
            </p>
            
            <!-- Back to home link -->
            <router-link 
              to="/" 
              class="inline-flex items-center mt-6 px-4 py-2 rounded-lg font-medium transition-all duration-200 hover:scale-105"
              :style="{ 
                color: isDark ? '#60a5fa' : '#2563eb',
                backgroundColor: isDark ? '#1e3a8a20' : '#dbeafe40'
              }"
            >
              ← Share new files
            </router-link>
          </div>
        </div>
      </div>

      <!-- Content Section -->
      <div class="relative max-w-4xl mx-auto px-6 lg:px-8 pb-32">
        <!-- Loading State -->
        <div v-if="loading" 
             class="rounded-2xl shadow-xl overflow-hidden animate-pulse transition-colors duration-300 p-12 text-center"
             :style="{ 
               backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
               borderColor: isDark ? '#374151cc' : '#e5e7ebcc',
               borderWidth: '1px'
             }">
          <div class="animate-spin w-12 h-12 border-4 border-blue-500 border-t-transparent rounded-full mx-auto mb-4"></div>
          <p class="text-lg font-medium transition-colors duration-300"
             :style="{ color: isDark ? '#f9fafb' : '#111827' }">Loading your files...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="error" 
             class="rounded-2xl shadow-xl overflow-hidden animate-shake transition-colors duration-300 p-12 text-center"
             :style="{ 
               backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
               borderColor: isDark ? '#dc2626cc' : '#ef4444cc',
               borderWidth: '2px'
             }">
          <XMarkIcon class="w-16 h-16 text-red-500 mx-auto mb-4" />
          <h2 class="text-2xl font-bold mb-4 text-red-500">Error</h2>
          <p class="text-lg mb-6 transition-colors duration-300"
             :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ error }}</p>
          <router-link 
            to="/" 
            class="inline-flex items-center px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-xl font-semibold transition-all duration-200 hover:scale-105"
          >
            ← Share new files
          </router-link>
        </div>

        <!-- Success State -->
        <div v-else class="space-y-8">
          <!-- Uploader Info -->
          <div v-if="uploader" 
               class="rounded-2xl shadow-xl overflow-hidden animate-slide-up transition-colors duration-300 relative"
               :style="{ 
                 backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
                 borderColor: isDark ? '#374151cc' : '#e5e7ebcc',
                 borderWidth: '1px'
               }">
            <!-- Gradient header background -->
            <div class="absolute top-0 left-0 right-0 h-24 bg-gradient-to-br from-blue-500 via-purple-500 to-pink-500 opacity-10"></div>
            
            <div class="relative p-8">
              <!-- "Shared by" label -->
              <div class="flex items-center mb-6">
                <div class="w-1 h-6 bg-gradient-to-b from-blue-500 to-purple-500 rounded-full mr-3"></div>
                <h2 class="text-sm font-semibold uppercase tracking-wider transition-colors duration-300"
                    :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Shared by</h2>
              </div>
              
              <!-- User info with enhanced styling -->
              <div class="flex flex-col sm:flex-row items-start sm:items-start space-y-4 sm:space-y-0 sm:space-x-6">
                <!-- Avatar with enhanced styling -->
                <div class="relative group flex-shrink-0">
                  <div class="w-20 h-20 rounded-full p-1 bg-gradient-to-br from-blue-500 via-purple-500 to-pink-500 animate-pulse-slow">
                    <div class="w-full h-full rounded-full flex items-center justify-center transition-colors duration-300"
                         :style="{ backgroundColor: isDark ? '#1a1a1a' : '#ffffff' }">
                      <img v-if="uploader.avatar" 
                           :src="`http://localhost:8080${uploader.avatar}`" 
                           alt="Avatar" 
                           class="w-16 h-16 rounded-full object-cover shadow-lg" 
                           @error="handleAvatarError" />
                      <UserIcon v-else 
                               class="w-10 h-10 transition-colors duration-300"
                               :style="{ color: isDark ? '#60a5fa' : '#2563eb' }" />
                    </div>
                  </div>
                  <!-- Online status indicator -->
                  <div class="absolute bottom-1 right-1 w-4 h-4 bg-green-500 rounded-full border-2 transition-colors duration-300"
                       :style="{ borderColor: isDark ? '#1a1a1a' : '#ffffff' }"></div>
                </div>
                
                <!-- User details with enhanced typography -->
                <div class="flex-1 min-w-0">
                  <h3 class="text-2xl font-bold mb-1 transition-colors duration-300"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ uploader.username }}</h3>
                  <p class="text-lg transition-colors duration-300 mb-3"
                     :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ uploader.email }}</p>
                  
                  <!-- Sharing stats -->
                  <div class="flex flex-wrap items-center gap-4 text-sm mb-4">
                    <div class="flex items-center space-x-1">
                      <div class="w-2 h-2 rounded-full bg-blue-500"></div>
                      <span class="transition-colors duration-300"
                            :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ files.length }} file{{ files.length > 1 ? 's' : '' }}</span>
                    </div>
                    <div class="flex items-center space-x-1">
                      <div class="w-2 h-2 rounded-full bg-purple-500"></div>
                      <span class="transition-colors duration-300"
                            :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ formatTotalSize() }}</span>
                    </div>
                    <div v-if="uploader.expirationDate" class="flex items-center space-x-1">
                      <div class="w-2 h-2 rounded-full" 
                           :class="formatExpirationDate(uploader.expirationDate).isExpired ? 'bg-red-500' : 'bg-green-500'"></div>
                      <span class="transition-colors duration-300 font-medium"
                            :class="formatExpirationDate(uploader.expirationDate).color">
                        {{ formatExpirationDate(uploader.expirationDate).text }}
                      </span>
                    </div>
                  </div>

                  <!-- Expiration Notice (if applicable) -->
                  <div v-if="uploader.expirationDate" class="p-3 rounded-lg border transition-colors duration-300"
                       :style="{ 
                         backgroundColor: formatExpirationDate(uploader.expirationDate).isExpired 
                           ? (isDark ? '#7f1d1d' : '#fee2e2') 
                           : (isDark ? '#064e3b' : '#d1fae5'),
                         borderColor: formatExpirationDate(uploader.expirationDate).isExpired 
                           ? (isDark ? '#dc2626' : '#ef4444') 
                           : (isDark ? '#059669' : '#10b981')
                       }">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4 flex-shrink-0" 
                           :class="formatExpirationDate(uploader.expirationDate).isExpired ? 'text-red-500' : 'text-green-600'"
                           fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd"/>
                      </svg>
                      <div class="flex-1">
                        <p class="text-sm font-medium transition-colors duration-300"
                           :class="formatExpirationDate(uploader.expirationDate).isExpired ? 'text-red-700' : 'text-green-700'">
                          {{ formatExpirationDate(uploader.expirationDate).isExpired ? 'Files have expired' : 'Files expire in:' }}
                        </p>
                        <p class="text-xs transition-colors duration-300"
                           :class="formatExpirationDate(uploader.expirationDate).isExpired ? 'text-red-600' : 'text-green-600'">
                          {{ formatExpirationDate(uploader.expirationDate).text }}
                          <span class="ml-2 opacity-75">
                            ({{ new Date(uploader.expirationDate).toLocaleDateString() }} {{ new Date(uploader.expirationDate).toLocaleTimeString() }})
                          </span>
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
                
                <!-- Share badge -->
                <div class="flex-shrink-0">
                  <div class="inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold transition-colors duration-300"
                       :style="{ 
                         backgroundColor: isDark ? '#1e3a8a' : '#dbeafe',
                         color: isDark ? '#60a5fa' : '#2563eb'
                       }">
                    <svg class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20">
                      <path d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z"/>
                    </svg>
                    Shared
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Files Header with Rainbow Animation -->
          <div class="rounded-2xl shadow-xl overflow-hidden animate-slide-up transition-colors duration-300"
               :style="{ 
                 backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
                 borderColor: isDark ? '#374151cc' : '#e5e7ebcc',
                 borderWidth: '1px'
               }">
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
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 714 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
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
              class="rounded-2xl shadow-xl overflow-hidden animate-slide-up transition-colors duration-300"
              :style="{ 
                animationDelay: `${index * 0.1}s`,
                backgroundColor: isDark ? '#1a1a1a' : '#ffffff',
                borderColor: isDark ? '#374151cc' : '#e5e7ebcc',
                borderWidth: '1px'
              }"
            >
              <div class="p-6">
                <div class="flex items-center justify-between mb-4">
                  <div class="flex items-center space-x-4">
                    <div class="w-16 h-16 rounded-xl flex items-center justify-center transition-colors duration-300"
                         :style="{ backgroundColor: isDark ? '#1e3a8a' : '#dbeafe' }">
                      <img :src="getFileIcon(file)" alt="File type" class="w-10 h-10" />
                    </div>
                    <div>
                      <h3 class="text-xl font-bold transition-colors duration-300"
                          :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ file.name }}</h3>
                      <p class="transition-colors duration-300"
                         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ formatFileSize(file.size) }}</p>
                      <p class="text-sm transition-colors duration-300"
                         :style="{ color: isDark ? '#6b7280' : '#9ca3af' }">{{ getFileExtension(file).toUpperCase() }} file</p>
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
                          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 714 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
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
                  <div class="border-t pt-6 transition-colors duration-300"
                       :style="{ borderColor: isDark ? '#374151' : '#e5e7eb' }">
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
                        class="w-full h-96 rounded-lg shadow-md border transition-colors duration-300"
                        :style="{ borderColor: isDark ? '#374151' : '#e5e7eb' }"
                      >
                        <div class="flex items-center justify-center h-96 rounded-lg transition-colors duration-300"
                             :style="{ backgroundColor: isDark ? '#111827' : '#f9fafb' }">
                          <div class="text-center">
                            <DocumentIcon class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                                         :style="{ color: isDark ? '#6b7280' : '#9ca3af' }" />
                            <p class="mb-4 transition-colors duration-300"
                               :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">PDF preview not supported in your browser</p>
                            <button 
                              @click="downloadFile(index)"
                              class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-all duration-200"
                            >
                              Download to view
                            </button>
                          </div>
                        </div>
                      </object>
                    </div>

                    <!-- Text Preview -->
                    <div v-else-if="isText(file)" class="rounded-lg p-4 transition-colors duration-300"
                         :style="{ backgroundColor: isDark ? '#111827' : '#f9fafb' }">
                      <pre class="text-sm whitespace-pre-wrap overflow-auto max-h-64 transition-colors duration-300"
                           :style="{ color: isDark ? '#e5e7eb' : '#374151' }">{{ textContent.get(index) || 'Loading...' }}</pre>
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
                      <DocumentIcon class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                                   :style="{ color: isDark ? '#6b7280' : '#9ca3af' }" />
                      <p class="mb-4 transition-colors duration-300"
                         :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Preview not available for this file type</p>
                      <button 
                        @click="downloadFile(index)"
                        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-all duration-200"
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
  expirationDate?: string
}

// State
const route = useRoute()
const files = ref<FileInfo[]>([])
const uploader = ref<UploaderInfo | null>(null)
const loading = ref(true)
const error = ref('')
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

const formatExpirationDate = (dateString: string) => {
  const expirationDate = new Date(dateString)
  const now = new Date()
  const diffMs = expirationDate.getTime() - now.getTime()
  
  if (diffMs <= 0) {
    return { text: 'Expired', isExpired: true, color: 'text-red-500' }
  }
  
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24))
  const diffHours = Math.floor((diffMs % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  
  if (diffDays > 0) {
    const color = diffDays <= 1 ? 'text-orange-500' : 'text-green-500'
    return { 
      text: `${diffDays} day${diffDays > 1 ? 's' : ''} remaining`, 
      isExpired: false, 
      color 
    }
  } else if (diffHours > 0) {
    return { 
      text: `${diffHours} hour${diffHours > 1 ? 's' : ''} remaining`, 
      isExpired: false, 
      color: 'text-orange-500' 
    }
  } else {
    const diffMinutes = Math.floor((diffMs % (1000 * 60 * 60)) / (1000 * 60))
    return { 
      text: `${diffMinutes} minute${diffMinutes > 1 ? 's' : ''} remaining`, 
      isExpired: false, 
      color: 'text-red-500' 
    }
  }
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

const loadFiles = async () => {
  try {
    const id = route.params.id as string
    if (!id) {
      error.value = 'No file ID provided'
      return
    }

    // Request file metadata from backend
    const response = await axios.get(`http://localhost:8080/files/${id}`)
    files.value = response.data.files || []
    uploader.value = response.data.uploader || null
    
    // For testing: Add expiration date if not present in backend response
    if (uploader.value && !uploader.value.expirationDate) {
      // Add a test expiration date (24 hours from now)
      const expDate = new Date()
      expDate.setHours(expDate.getHours() + 24)
      uploader.value.expirationDate = expDate.toISOString()
    }
    
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
  await loadFiles()
})
</script>

<style scoped>
/* Animations matching Home.vue */
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

@keyframes pulse-slow {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.8; }
}

/* Animation Classes */
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

.animate-pulse-slow {
  animation: pulse-slow 3s ease-in-out infinite;
}

/* Enhanced hover effects */
.hover\:scale-105:hover {
  transform: scale(1.05);
}

.hover\:scale-110:hover {
  transform: scale(1.10);
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Custom focus states */
input:focus, button:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

/* Mobile Optimizations */
@media (max-width: 768px) {
  .text-5xl {
    font-size: 2.5rem;
  }
  
  .text-2xl {
    font-size: 1.5rem;
  }
  
  .space-x-6 > :not([hidden]) ~ :not([hidden]) {
    margin-left: 1rem;
  }
  
  .space-x-4 > :not([hidden]) ~ :not([hidden]) {
    margin-left: 0.5rem;
  }
  
  .flex-wrap {
    flex-wrap: wrap;
  }
  
  .gap-4 {
    gap: 0.5rem;
  }
  
  .hidden {
    display: none;
  }
  
  .sm\:inline {
    display: inline;
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