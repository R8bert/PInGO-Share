<template>
  <div class="min-h-screen transition-colors duration-300 relative overflow-hidden" 
       :style="{ backgroundColor: isDark ? '#000000' : '#fbfbfd' }">
    
    <!-- Starfield Background -->
    <div v-if="isDark" class="fixed inset-0 z-0">
      <!-- Twinkling Stars -->
      <div v-for="star in stars" :key="`star-${star.id}`"
           class="absolute rounded-full bg-white animate-twinkle"
           :style="{
             left: star.x + '%',
             top: star.y + '%',
             width: star.size + 'px',
             height: star.size + 'px',
             animationDelay: star.delay + 's',
             animationDuration: star.duration + 's'
           }">
      </div>
      
      <!-- Constellations -->
      <div v-for="constellation in constellations" :key="`constellation-${constellation.id}`"
           class="absolute opacity-30"
           :style="{
             left: constellation.x + '%',
             top: constellation.y + '%'
           }">
        <svg width="60" height="60" viewBox="0 0 60 60">
          <circle cx="10" cy="10" r="1" fill="white" />
          <circle cx="30" cy="20" r="1.5" fill="white" />
          <circle cx="45" cy="15" r="1" fill="white" />
          <circle cx="20" cy="40" r="1" fill="white" />
          <circle cx="50" cy="45" r="1.5" fill="white" />
          <line x1="10" y1="10" x2="30" y2="20" stroke="white" stroke-width="0.3" opacity="0.5" />
          <line x1="30" y1="20" x2="45" y2="15" stroke="white" stroke-width="0.3" opacity="0.5" />
          <line x1="30" y1="20" x2="20" y2="40" stroke="white" stroke-width="0.3" opacity="0.5" />
          <line x1="20" y1="40" x2="50" y2="45" stroke="white" stroke-width="0.3" opacity="0.5" />
        </svg>
      </div>
      
      <!-- Nebula Effect -->
      <div class="absolute top-20 left-1/4 w-96 h-96 rounded-full blur-3xl opacity-10 animate-float-slow bg-blue-500"></div>
      <div class="absolute bottom-20 right-1/4 w-80 h-80 rounded-full blur-3xl opacity-10 animate-float-slower bg-purple-500"></div>
    </div>

    <!-- Main Content -->
    <main class="relative z-10 pt-8">
      <!-- Hero Section with Apple-style Typography -->
      <section class="relative overflow-hidden">
        <!-- Subtle background gradient -->
        <div class="absolute inset-0">
          <div class="absolute top-20 left-1/4 w-96 h-96 rounded-full blur-3xl opacity-20 animate-float-slow"
               :style="{ backgroundColor: isDark ? '#3b82f6' : '#60a5fa' }"></div>
          <div class="absolute bottom-20 right-1/4 w-80 h-80 rounded-full blur-3xl opacity-20 animate-float-slower"
               :style="{ backgroundColor: isDark ? '#8b5cf6' : '#a78bfa' }"></div>
        </div>
        
        <div class="relative max-w-4xl mx-auto px-6 lg:px-8 pt-20 pb-16 text-center">
          
          <!-- Apple-style large headline -->
          <h1 class="text-6xl lg:text-7xl font-bold leading-tight mb-6 animate-fade-in"
              :style="{ 
                color: isDark ? '#ffffff' : '#1d1d1f',
                fontFamily: 'system-ui, -apple-system, sans-serif'
              }">
            Your <span class="text-transparent bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 bg-clip-text animate-rainbow">files</span> are ready
          </h1>
          
          <!-- Apple-style subtitle -->
          <p class="text-xl lg:text-2xl font-light max-w-2xl mx-auto leading-relaxed animate-fade-in-delay transition-colors duration-300"
             :style="{ 
               color: isDark ? '#a1a1aa' : '#6e6e73',
               fontFamily: 'system-ui, -apple-system, sans-serif'
             }">
            {{ loading ? 'Preparing your download experience...' : `${files.length} file${files.length !== 1 ? 's' : ''} shared with you` }}
          </p>
        </div>
      </section>

      <!-- Content Section -->
      <section class="relative max-w-4xl mx-auto px-6 lg:px-8 pb-20">
        
        <!-- Loading State - Apple Style -->
        <div v-if="loading" 
             class="mx-auto max-w-md animate-pulse">
          <div class="rounded-3xl p-12 text-center backdrop-blur-xl border transition-all duration-300"
               :style="{ 
                 backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(255, 255, 255, 0.7)',
                 borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
               }">
            <div class="w-16 h-16 mx-auto mb-6 rounded-full animate-spin border-4 border-blue-500 border-t-transparent"></div>
            <p class="text-lg font-medium transition-colors duration-300"
               :style="{ color: isDark ? '#ffffff' : '#1d1d1f' }">Loading files...</p>
          </div>
        </div>

        <!-- Error State - Apple Style -->
        <div v-else-if="error" 
             class="mx-auto max-w-md animate-scale-in">
          <div class="rounded-3xl p-12 text-center backdrop-blur-xl border transition-all duration-300"
               :style="{ 
                 backgroundColor: isDark ? 'rgba(239, 68, 68, 0.1)' : 'rgba(254, 226, 226, 0.8)',
                 borderColor: isDark ? 'rgba(239, 68, 68, 0.3)' : 'rgba(239, 68, 68, 0.2)'
               }">
            <div class="w-16 h-16 mx-auto mb-6 rounded-full flex items-center justify-center"
                 :style="{ backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(254, 226, 226, 1)' }">
              <XMarkIcon class="w-8 h-8 text-red-500" />
            </div>
            <h2 class="text-2xl font-semibold mb-4 text-red-500">Oops!</h2>
            <p class="text-lg mb-8 transition-colors duration-300"
               :style="{ color: isDark ? '#d1d5db' : '#374151' }">{{ error }}</p>
            <router-link 
              to="/" 
              class="inline-flex items-center px-8 py-4 bg-blue-500 hover:bg-blue-600 text-white rounded-2xl font-semibold transition-all duration-200 hover:scale-105"
            >
              ← Go Home
            </router-link>
          </div>
        </div>

        <!-- Success State - Apple Style -->
        <div v-else class="space-y-8">
          
          <!-- Uploader Card - Apple Style -->
          <div v-if="uploader" 
               class="rounded-3xl p-8 backdrop-blur-xl border transition-all duration-300 animate-slide-up"
               :style="{ 
                 backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(255, 255, 255, 0.7)',
                 borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
               }">
            
            <div class="flex items-center space-x-6">
              <!-- Apple-style Avatar -->
              <div class="relative flex-shrink-0">
                <div class="w-20 h-20 rounded-2xl overflow-hidden shadow-xl"
                     :style="{ backgroundColor: isDark ? '#1e1e1e' : '#f5f5f7' }">
                  <img v-if="uploader.avatar" 
                       :src="`http://localhost:8080${uploader.avatar}`" 
                       alt="Avatar" 
                       class="w-full h-full object-cover" 
                       @error="handleAvatarError" />
                  <div v-else class="w-full h-full flex items-center justify-center">
                    <UserIcon class="w-10 h-10 transition-colors duration-300"
                             :style="{ color: isDark ? '#60a5fa' : '#2563eb' }" />
                  </div>
                </div>
                <!-- Online indicator -->
                <div class="absolute -bottom-1 -right-1 w-6 h-6 bg-green-500 rounded-full border-4 transition-colors duration-300"
                     :style="{ borderColor: isDark ? '#000000' : '#fbfbfd' }"></div>
              </div>
              
              <!-- User Info -->
              <div class="flex-1 min-w-0">
                <h3 class="text-2xl font-semibold mb-1 transition-colors duration-300"
                    :style="{ 
                      color: isDark ? '#ffffff' : '#1d1d1f',
                      fontFamily: 'system-ui, -apple-system, sans-serif'
                    }">{{ uploader.username }}</h3>
                <p class="text-lg transition-colors duration-300 mb-3"
                   :style="{ color: isDark ? '#a1a1aa' : '#6e6e73' }">{{ uploader.email }}</p>
                
                <!-- Stats with Apple-style pills -->
                <div class="flex flex-wrap gap-3">
                  <div class="px-4 py-2 rounded-full transition-colors duration-300"
                       :style="{ 
                         backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)',
                         color: isDark ? '#60a5fa' : '#2563eb'
                       }">
                    {{ files.length }} file{{ files.length > 1 ? 's' : '' }}
                  </div>
                  <div class="px-4 py-2 rounded-full transition-colors duration-300"
                       :style="{ 
                         backgroundColor: isDark ? 'rgba(139, 92, 246, 0.2)' : 'rgba(139, 92, 246, 0.1)',
                         color: isDark ? '#a78bfa' : '#7c3aed'
                       }">
                    {{ formatTotalSize() }}
                  </div>
                  <div v-if="uploader.expirationDate" 
                       class="px-4 py-2 rounded-full transition-colors duration-300"
                       :style="{ 
                         backgroundColor: formatExpirationDate(uploader.expirationDate).isExpired 
                           ? (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)')
                           : (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)'),
                         color: formatExpirationDate(uploader.expirationDate).isExpired 
                           ? '#ef4444' 
                           : '#22c55e'
                       }">
                    {{ formatExpirationDate(uploader.expirationDate).text }}
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Files Header with Apple-style Action Button -->
          <div class="flex items-center justify-between mb-8">
            <div>
              <h2 class="text-3xl font-semibold transition-colors duration-300"
                  :style="{ 
                    color: isDark ? '#ffffff' : '#1d1d1f',
                    fontFamily: 'system-ui, -apple-system, sans-serif'
                  }">Download</h2>
              <p class="text-lg transition-colors duration-300 mt-1"
                 :style="{ color: isDark ? '#a1a1aa' : '#6e6e73' }">{{ formatTotalSize() }} total</p>
            </div>
            
            <!-- Apple-style Download All Button -->
            <button 
              @click="downloadAll"
              :disabled="downloadingAll"
              class="group relative px-8 py-4 bg-blue-500 hover:bg-blue-600 text-white rounded-2xl font-semibold transition-all duration-200 hover:scale-105 disabled:opacity-50 shadow-lg"
            >
              <span v-if="downloadingAll" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-3 h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
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

          <!-- Files Grid - Apple Style -->
          <div class="grid grid-cols-1 gap-6">
            <div 
              v-for="(file, index) in files" 
              :key="index" 
              class="group rounded-3xl p-6 backdrop-blur-xl border transition-all duration-300 hover:scale-[1.02] animate-slide-up"
              :style="{ 
                animationDelay: `${index * 0.1}s`,
                backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(255, 255, 255, 0.7)',
                borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
              }"
            >
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-4 flex-1 min-w-0">
                  <!-- Apple-style File Icon -->
                  <div class="w-16 h-16 rounded-2xl flex items-center justify-center flex-shrink-0 transition-colors duration-300"
                       :style="{ backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)' }">
                    <img :src="getFileIcon(file)" alt="File type" class="w-10 h-10" />
                  </div>
                  
                  <!-- File Details -->
                  <div class="flex-1 min-w-0">
                    <h3 class="text-xl font-semibold mb-1 truncate transition-colors duration-300"
                        :style="{ 
                          color: isDark ? '#ffffff' : '#1d1d1f',
                          fontFamily: 'system-ui, -apple-system, sans-serif'
                        }">{{ file.name }}</h3>
                    <div class="flex items-center space-x-4 text-sm">
                      <span class="transition-colors duration-300"
                            :style="{ color: isDark ? '#a1a1aa' : '#6e6e73' }">{{ formatFileSize(file.size) }}</span>
                      <span class="px-3 py-1 rounded-full text-xs font-medium transition-colors duration-300"
                            :style="{ 
                              backgroundColor: isDark ? 'rgba(139, 92, 246, 0.2)' : 'rgba(139, 92, 246, 0.1)',
                              color: isDark ? '#a78bfa' : '#7c3aed'
                            }">
                        {{ getFileExtension(file).toUpperCase() }}
                      </span>
                    </div>
                  </div>
                </div>
                
                <!-- Action Buttons -->
                <div class="flex items-center space-x-3 flex-shrink-0">
                  <!-- Preview Button -->
                  <button 
                    v-if="canPreview(file)" 
                    @click="togglePreview(index)"
                    class="p-3 rounded-xl transition-all duration-200 hover:scale-110"
                    :style="{
                      backgroundColor: previewingFiles.has(index) 
                        ? (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)')
                        : (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)'),
                      color: previewingFiles.has(index) ? '#ef4444' : '#22c55e'
                    }"
                    :title="previewingFiles.has(index) ? 'Hide Preview' : 'Show Preview'"
                  >
                    <EyeIcon v-if="!previewingFiles.has(index)" class="w-5 h-5" />
                    <EyeSlashIcon v-else class="w-5 h-5" />
                  </button>
                  
                  <!-- Download Button -->
                  <button 
                    @click="downloadFile(index)"
                    :disabled="downloadingFiles.has(index)"
                    class="px-6 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-xl font-medium transition-all duration-200 hover:scale-105 disabled:opacity-50"
                  >
                    <span v-if="downloadingFiles.has(index)" class="flex items-center">
                      <svg class="animate-spin -ml-1 mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 714 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                      </svg>
                      Loading
                    </span>
                    <span v-else class="flex items-center">
                      <ArrowDownTrayIcon class="w-4 h-4 mr-2" />
                      Download
                    </span>
                  </button>
                </div>
              </div>

              <!-- Preview Section - Apple Style -->
              <div v-if="previewingFiles.has(index)" class="mt-8 animate-slide-down">
                <div class="border-t pt-6 transition-colors duration-300"
                     :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }">
                  
                  <!-- Image Preview -->
                  <div v-if="isImage(file)" class="text-center">
                    <img 
                      :src="getFileUrl(file)" 
                      :alt="file.name"
                      class="max-w-full max-h-96 mx-auto rounded-2xl shadow-2xl"
                      @error="handlePreviewError"
                    />
                  </div>

                  <!-- Video Preview -->
                  <div v-else-if="isVideo(file)" class="text-center">
                    <video 
                      controls 
                      class="max-w-full max-h-96 mx-auto rounded-2xl shadow-2xl"
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
                      class="w-full h-96 rounded-2xl shadow-2xl border transition-colors duration-300"
                      :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }"
                    >
                      <div class="flex items-center justify-center h-96 rounded-2xl transition-colors duration-300"
                           :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)' }">
                        <div class="text-center">
                          <DocumentIcon class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                                       :style="{ color: isDark ? '#6b7280' : '#9ca3af' }" />
                          <p class="mb-4 transition-colors duration-300"
                             :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">PDF preview not supported</p>
                          <button 
                            @click="downloadFile(index)"
                            class="px-6 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-xl font-medium transition-all duration-200"
                          >
                            Download to view
                          </button>
                        </div>
                      </div>
                    </object>
                  </div>

                  <!-- Text Preview -->
                  <div v-else-if="isText(file)" class="rounded-2xl p-6 transition-colors duration-300"
                       :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)' }">
                    <pre class="text-sm whitespace-pre-wrap overflow-auto max-h-64 transition-colors duration-300 font-mono"
                         :style="{ color: isDark ? '#e5e7eb' : '#374151' }">{{ textContent.get(index) || 'Loading...' }}</pre>
                  </div>

                  <!-- Audio Preview -->
                  <div v-else-if="isAudio(file)" class="text-center">
                    <div class="inline-block p-6 rounded-2xl transition-colors duration-300"
                         :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)' }">
                      <audio 
                        controls 
                        class="max-w-md"
                        :src="getFileUrl(file)"
                      >
                        Your browser does not support the audio tag.
                      </audio>
                    </div>
                  </div>

                  <!-- Unsupported Preview -->
                  <div v-else class="text-center py-12">
                    <DocumentIcon class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                                 :style="{ color: isDark ? '#6b7280' : '#9ca3af' }" />
                    <p class="mb-6 text-lg transition-colors duration-300"
                       :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Preview not available</p>
                    <button 
                      @click="downloadFile(index)"
                      class="px-8 py-4 bg-blue-500 hover:bg-blue-600 text-white rounded-2xl font-semibold transition-all duration-200 hover:scale-105"
                    >
                      Download to view
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useTheme } from '../../composables/useTheme'
import { XMarkIcon, ArrowDownTrayIcon, DocumentIcon, EyeIcon, EyeSlashIcon, UserIcon } from '@heroicons/vue/24/outline'
import axios from 'axios'

// Types
interface FileInfo {
  name: string
  size: number
  url: string
}

interface UploaderInfo {
  username: string
  email: string
  avatar?: string
  expirationDate?: string
}

// Dark mode
const { isDark } = useTheme()

// Router
const route = useRoute()

// State
const loading = ref(true)
const error = ref('')
const files = ref<FileInfo[]>([])
const uploader = ref<UploaderInfo | null>(null)
const previewingFiles = ref(new Set<number>())
const downloadingFiles = ref(new Set<number>())
const downloadingAll = ref(false)
const textContent = ref(new Map<number, string>())

// Star animation data
const stars = ref<Array<{
  id: number
  x: number
  y: number
  size: number
  delay: number
  duration: number
}>>([])

const constellations = ref<Array<{
  id: number
  x: number
  y: number
}>>([])

// Initialize star field
const initializeStars = () => {
  if (!isDark.value) return
  
  // Generate twinkling stars
  for (let i = 0; i < 200; i++) {
    stars.value.push({
      id: i,
      x: Math.random() * 100,
      y: Math.random() * 100,
      size: Math.random() * 2 + 1,
      delay: Math.random() * 5,
      duration: Math.random() * 3 + 2
    })
  }
  
  // Generate constellations
  for (let i = 0; i < 8; i++) {
    constellations.value.push({
      id: i,
      x: Math.random() * 90,
      y: Math.random() * 90
    })
  }
}

// Computed properties
const shareId = computed(() => route.params.id as string)

// Utility functions
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
  const extension = getFileExtension(file).toLowerCase()
  const iconMap: Record<string, string> = {
    pdf: '/src/assets/images/train/icons/file-pdf.png',
    jpg: '/src/assets/images/train/icons/file-jpg.png',
    jpeg: '/src/assets/images/train/icons/file-jpg.png',
    png: '/src/assets/images/train/icons/file-png.png',
    mp4: '/src/assets/images/train/icons/file-mp4.png',
    avi: '/src/assets/images/train/icons/file-mp4.png',
    mov: '/src/assets/images/train/icons/file-mp4.png',
    docx: '/src/images/icons/word_docx_doc.svg',
    doc: '/src/images/icons/word_docx_doc.svg',
    ppt: '/src/images/icons/ppt_pptx_file.svg',
    pptx: '/src/images/icons/ppt_pptx_file.svg',
    folder: '/src/assets/images/train/icons/file-folder.png'
  }
  return iconMap[extension] || '/src/assets/images/train/icons/file-folder.png'
}

const getFileExtension = (file: FileInfo) => {
  return file.name.split('.').pop() || ''
}

const canPreview = (file: FileInfo) => {
  const extension = getFileExtension(file).toLowerCase()
  const previewableTypes = [
    'jpg', 'jpeg', 'png', 'gif', 'webp', 'svg',
    'mp4', 'webm', 'ogg', 'avi', 'mov',
    'pdf',
    'txt', 'md', 'js', 'ts', 'html', 'css', 'json', 'xml', 'csv',
    'mp3', 'wav', 'flac', 'aac', 'ogg'
  ]
  return previewableTypes.includes(extension)
}

const isImage = (file: FileInfo) => {
  const extension = getFileExtension(file).toLowerCase()
  return ['jpg', 'jpeg', 'png', 'gif', 'webp', 'svg'].includes(extension)
}

const isVideo = (file: FileInfo) => {
  const extension = getFileExtension(file).toLowerCase()
  return ['mp4', 'webm', 'ogg', 'avi', 'mov'].includes(extension)
}

const isPdf = (file: FileInfo) => {
  return getFileExtension(file).toLowerCase() === 'pdf'
}

const isText = (file: FileInfo) => {
  const extension = getFileExtension(file).toLowerCase()
  return ['txt', 'md', 'js', 'ts', 'html', 'css', 'json', 'xml', 'csv'].includes(extension)
}

const isAudio = (file: FileInfo) => {
  const extension = getFileExtension(file).toLowerCase()
  return ['mp3', 'wav', 'flac', 'aac', 'ogg'].includes(extension)
}

const getFileUrl = (file: FileInfo) => {
  return `http://localhost:8080${file.url}`
}

const formatExpirationDate = (dateString: string) => {
  const now = new Date()
  const expiration = new Date(dateString)
  const timeDiff = expiration.getTime() - now.getTime()
  
  if (timeDiff <= 0) {
    return {
      text: 'Expired',
      isExpired: true,
      color: 'text-red-500'
    }
  }
  
  const days = Math.floor(timeDiff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((timeDiff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  const minutes = Math.floor((timeDiff % (1000 * 60 * 60)) / (1000 * 60))
  
  let text = ''
  if (days > 0) {
    text = `${days} day${days > 1 ? 's' : ''}`
  } else if (hours > 0) {
    text = `${hours} hour${hours > 1 ? 's' : ''}`
  } else {
    text = `${minutes} minute${minutes > 1 ? 's' : ''}`
  }
  
  return {
    text,
    isExpired: false,
    color: 'text-green-500'
  }
}

// API functions
const fetchShareData = async () => {
  try {
    loading.value = true
    const response = await axios.get(`http://localhost:8080/files/${shareId.value}`)
    
    // Backend returns { files: [...], uploader: {...} } directly
    files.value = response.data.files || []
    uploader.value = response.data.uploader || null
  } catch (err: any) {
    console.error('Error fetching share data:', err)
    if (err.response?.status === 404) {
      error.value = 'Files not found or have expired'
    } else {
      error.value = err.response?.data?.error || 'Failed to load files'
    }
  } finally {
    loading.value = false
  }
}

const togglePreview = async (index: number) => {
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
    return
  }
  
  previewingFiles.value.add(index)
  
  const file = files.value[index]
  if (isText(file) && !textContent.value.has(index)) {
    try {
      const response = await axios.get(getFileUrl(file), { responseType: 'text' })
      textContent.value.set(index, response.data)
    } catch (err) {
      console.error('Error loading text content:', err)
      textContent.value.set(index, 'Error loading file content')
    }
  }
}

const downloadFile = async (index: number) => {
  const file = files.value[index]
  downloadingFiles.value.add(index)
  
  try {
    const response = await axios.get(getFileUrl(file), {
      responseType: 'blob'
    })
    
    const blob = new Blob([response.data])
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = file.name
    link.click()
    window.URL.revokeObjectURL(url)
  } catch (err) {
    console.error('Error downloading file:', err)
  } finally {
    downloadingFiles.value.delete(index)
  }
}

const downloadAll = async () => {
  downloadingAll.value = true
  
  try {
    const response = await axios.get(`http://localhost:8080/download/${shareId.value}`, {
      responseType: 'blob'
    })
    
    const blob = new Blob([response.data])
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    
    // If multiple files, backend returns ZIP; if single file, returns the file directly
    if (files.value.length > 1) {
      link.download = `shared-files-${shareId.value}.zip`
    } else if (files.value.length === 1) {
      link.download = files.value[0].name
    } else {
      link.download = `files-${shareId.value}.zip`
    }
    
    link.click()
    window.URL.revokeObjectURL(url)
  } catch (err) {
    console.error('Error downloading all files:', err)
    
    // Fallback: download files individually
    for (let i = 0; i < files.value.length; i++) {
      await downloadFile(i)
      // Add small delay between downloads
      await new Promise(resolve => setTimeout(resolve, 500))
    }
  } finally {
    downloadingAll.value = false
  }
}

const handleAvatarError = () => {
  if (uploader.value) {
    uploader.value.avatar = undefined
  }
}

const handlePreviewError = () => {
  console.log('Preview error occurred')
}

// Lifecycle
onMounted(() => {
  initializeStars()
  fetchShareData()
})
</script>

<style scoped>
/* Star animations */
@keyframes twinkle {
  0%, 100% { 
    opacity: 0.3; 
    transform: scale(1); 
  }
  50% { 
    opacity: 1; 
    transform: scale(1.2); 
  }
}

/* Apple-style animations */
@keyframes fade-in {
  from { 
    opacity: 0; 
    transform: translateY(20px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

@keyframes fade-in-delay {
  0%, 30% { 
    opacity: 0; 
    transform: translateY(20px); 
  }
  100% { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

@keyframes slide-up {
  from { 
    opacity: 0; 
    transform: translateY(30px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

@keyframes slide-down {
  from { 
    opacity: 0; 
    transform: translateY(-20px); 
  }
  to { 
    opacity: 1; 
    transform: translateY(0); 
  }
}

@keyframes scale-in {
  from { 
    opacity: 0; 
    transform: scale(0.95); 
  }
  to { 
    opacity: 1; 
    transform: scale(1); 
  }
}

@keyframes float-slow {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(2deg); }
}

@keyframes float-slower {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-30px) rotate(-2deg); }
}

@keyframes rainbow {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.animate-twinkle {
  animation: twinkle var(--duration, 3s) ease-in-out infinite;
}

.animate-fade-in {
  animation: fade-in 0.8s ease-out forwards;
}

.animate-fade-in-delay {
  animation: fade-in-delay 1.2s ease-out forwards;
}

.animate-slide-up {
  animation: slide-up 0.6s ease-out forwards;
}

.animate-slide-down {
  animation: slide-down 0.4s ease-out forwards;
}

.animate-scale-in {
  animation: scale-in 0.5s ease-out forwards;
}

.animate-float-slow {
  animation: float-slow 6s ease-in-out infinite;
}

.animate-float-slower {
  animation: float-slower 8s ease-in-out infinite;
}

.animate-rainbow {
  background-size: 400% 400%;
  animation: rainbow 3s ease infinite;
}

/* Apple-style backdrop blur */
.backdrop-blur-xl {
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
}

/* Apple-style shadows */
.shadow-xl {
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

.shadow-2xl {
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
}

/* Apple-style smooth scrolling */
* {
  scroll-behavior: smooth;
}

/* Apple-style reduced motion support */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
</style>
