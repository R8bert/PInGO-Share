<template>
  <div class="min-h-screen transition-colors duration-300" 
       :style="{ backgroundColor: isDark ? '#000000' : '#ffffff' }">
    
    <!-- WebGL Background -->
    <div class="fixed inset-0 pointer-events-none overflow-hidden z-0">
      <WebGLBackground 
        :hue-shift="isDark ? 240 : 120"
        :noise-intensity="isDark ? 0.03 : 0.01"
        :scanline-intensity="isDark ? 0.1 : 0.05"
        :speed="0.3"
        :scanline-frequency="isDark ? 0.5 : 0.2"
        :warp-amount="0.1"
        :resolution-scale="0.8"
        class="opacity-40"
      />
    </div>

    <!-- Main Content -->
    <main class="relative z-10">
      
      <!-- Hero Section -->
      <section class="relative min-h-screen flex items-center justify-center px-4 sm:px-6 lg:px-8">
        <div class="max-w-7xl mx-auto w-full">
          <div class="text-center space-y-8">
            
            <!-- Badge -->
            <div class="inline-flex items-center space-x-2 px-4 py-2 rounded-full border backdrop-blur-sm animate-fade-in"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.05)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                   color: isDark ? '#a1a1aa' : '#71717a'
                 }">
              <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
              <span class="text-sm font-medium">Secure & Fast File Transfer</span>
            </div>

            <!-- Main Headline -->
            <div class="space-y-6 animate-fade-in-delay">
              <h1 class="text-5xl sm:text-6xl lg:text-7xl font-bold tracking-tight">
                <span :style="{ color: isDark ? '#ffffff' : '#000000' }">Share files</span>
                <br />
                <span class="text-transparent bg-gradient-to-r from-blue-600 via-purple-600 to-pink-600 bg-clip-text">
                  effortlessly
                </span>
              </h1>
              
              <p class="text-xl sm:text-2xl max-w-3xl mx-auto leading-relaxed"
                 :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                Transfer files of any size securely and instantly. No registration required, 
                no limits, just pure simplicity.
              </p>
            </div>

            <!-- CTA Buttons -->
            <div class="flex flex-col sm:flex-row items-center justify-center gap-4 animate-fade-in-delay-2">
              <button 
                @click="scrollToUpload"
                class="group relative px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-2xl transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-blue-500/25"
              >
                <span class="relative z-10">Start Uploading</span>
                <div class="absolute inset-0 bg-gradient-to-r from-blue-700 to-purple-700 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              </button>
              
              <button 
                @click="scrollToFeatures"
                class="px-8 py-4 border-2 font-semibold rounded-2xl transition-all duration-300 hover:scale-105 backdrop-blur-sm"
                :style="{
                  borderColor: isDark ? 'rgba(255,255,255,0.2)' : 'rgba(0,0,0,0.2)',
                  color: isDark ? '#ffffff' : '#000000'
                }"
              >
                Learn More
              </button>
            </div>

          </div>
        </div>

        <!-- Scroll Indicator -->
        <div class="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce">
          <div class="w-6 h-10 border-2 rounded-full flex justify-center"
               :style="{ borderColor: isDark ? 'rgba(255,255,255,0.3)' : 'rgba(0,0,0,0.3)' }">
            <div class="w-1 h-3 bg-gradient-to-b from-blue-500 to-purple-500 rounded-full mt-2 animate-pulse"></div>
          </div>
        </div>
      </section>

      <!-- Upload Section -->
      <section ref="uploadSection" class="relative py-20 px-4 sm:px-6 lg:px-8">
        <div class="max-w-4xl mx-auto">
          
          <!-- Upload Card -->
          <div class="relative">
            <!-- Background Glow -->
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-3xl blur-3xl"></div>
            
            <!-- Main Card -->
            <div class="relative backdrop-blur-xl border rounded-3xl p-8 sm:p-12"
                 :style="{
                   backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                   borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                 }">
              
              <!-- Upload Zone -->
              <div 
                @dragover.prevent 
                @drop.prevent="onDrop" 
                @dragenter="isDragging = true" 
                @dragleave="isDragging = false"
                @click="triggerFileInput"
                :class="[
                  'relative border-2 border-dashed transition-all duration-300 cursor-pointer rounded-2xl p-12',
                  isDragging 
                    ? 'border-blue-500 bg-blue-500/10 scale-[1.02]' 
                    : 'border-gray-300 dark:border-gray-600 hover:border-blue-400 hover:bg-blue-500/5'
                ]"
              >
                <input type="file" ref="fileInput" @change="onFileChange" class="hidden" multiple />
                
                <div class="text-center space-y-6">
                  <!-- Upload Icon -->
                  <div class="relative mx-auto w-24 h-24">
                    <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl blur-xl opacity-30"></div>
                    <div class="relative bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl p-6 flex items-center justify-center">
                      <CloudArrowUpIcon class="w-12 h-12 text-white" />
                    </div>
                  </div>

                  <!-- Upload Text -->
                  <div class="space-y-2">
                    <h3 class="text-2xl font-bold"
                        :style="{ color: isDark ? '#ffffff' : '#000000' }">
                      {{ selectedFiles.length > 0 ? `${selectedFiles.length} file${selectedFiles.length > 1 ? 's' : ''} selected` : 'Drop files here or click to browse' }}
                    </h3>
                    <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                      Support for any file type • No size limits • Encrypted transfer
                    </p>
                  </div>

                  <!-- Features -->
                  <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 pt-6">
                    <div class="flex items-center justify-center space-x-2 p-3 rounded-xl"
                         :style="{ backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.05)' }">
                      <div class="w-2 h-2 bg-green-500 rounded-full"></div>
                      <span class="text-sm font-medium" :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                        Secure
                      </span>
                    </div>
                    <div class="flex items-center justify-center space-x-2 p-3 rounded-xl"
                         :style="{ backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.05)' }">
                      <div class="w-2 h-2 bg-blue-500 rounded-full"></div>
                      <span class="text-sm font-medium" :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                        Fast
                      </span>
                    </div>
                    <div class="flex items-center justify-center space-x-2 p-3 rounded-xl"
                         :style="{ backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.05)' }">
                      <div class="w-2 h-2 bg-purple-500 rounded-full"></div>
                      <span class="text-sm font-medium" :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                        Simple
                      </span>
                    </div>
                  </div>
                </div>

                <!-- Drag overlay -->
                <div v-if="isDragging" 
                     class="absolute inset-0 rounded-2xl flex items-center justify-center z-20 bg-blue-500/20 backdrop-blur-sm">
                  <div class="text-2xl font-bold text-blue-600">Drop files here</div>
                </div>
              </div>

              <!-- Selected Files -->
              <div v-if="selectedFiles.length > 0" class="mt-8 space-y-4">
                <h4 class="text-lg font-semibold" :style="{ color: isDark ? '#ffffff' : '#000000' }">
                  Selected Files
                </h4>
                <div class="grid gap-4">
                  <div v-for="(file, index) in selectedFiles" :key="index" 
                       class="flex items-center justify-between p-4 rounded-xl border"
                       :style="{
                         backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(0,0,0,0.05)',
                         borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                       }">
                    <div class="flex items-center space-x-3">
                      <div class="w-10 h-10 bg-gradient-to-r from-blue-500 to-purple-500 rounded-lg flex items-center justify-center">
                        <span class="text-white text-xs font-bold">{{ getFileExtension(file).toUpperCase() }}</span>
                      </div>
                      <div>
                        <p class="font-medium" :style="{ color: isDark ? '#ffffff' : '#000000' }">{{ file.name }}</p>
                        <p class="text-sm" :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">{{ formatFileSize(file.size) }}</p>
                      </div>
                    </div>
                    <button @click="removeFile(index)" 
                            class="p-2 rounded-full hover:bg-red-500/20 text-red-500 transition-colors">
                      <XMarkIcon class="w-5 h-5" />
                    </button>
                  </div>
                </div>

                <!-- Upload Button -->
                <div class="pt-6">
                  <button @click="uploadFiles" 
                          :disabled="isUploading"
                          class="w-full py-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-2xl transition-all duration-300 hover:scale-[1.02] hover:shadow-2xl disabled:opacity-50 disabled:hover:scale-100">
                    {{ isUploading ? 'Uploading...' : 'Upload Files' }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Features Section -->
      <section ref="featuresSection" class="relative py-20 px-4 sm:px-6 lg:px-8">
        <div class="max-w-7xl mx-auto">
          <div class="text-center mb-16">
            <h2 class="text-4xl sm:text-5xl font-bold mb-6"
                :style="{ color: isDark ? '#ffffff' : '#000000' }">
              Why choose PInGO Share?
            </h2>
            <p class="text-xl max-w-3xl mx-auto"
               :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
              Built with modern technology and user experience in mind
            </p>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            
            <!-- Feature 1 -->
            <div class="relative group">
              <div class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              <div class="relative backdrop-blur-xl border rounded-2xl p-8 transition-all duration-300 group-hover:scale-105"
                   :style="{
                     backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                     borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                   }">
                <div class="w-16 h-16 bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl flex items-center justify-center mb-6">
                  <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                  </svg>
                </div>
                <h3 class="text-xl font-bold mb-4" :style="{ color: isDark ? '#ffffff' : '#000000' }">
                  End-to-End Encryption
                </h3>
                <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                  Your files are encrypted before they leave your device, ensuring maximum security.
                </p>
              </div>
            </div>

            <!-- Feature 2 -->
            <div class="relative group">
              <div class="absolute inset-0 bg-gradient-to-r from-green-600/20 to-blue-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              <div class="relative backdrop-blur-xl border rounded-2xl p-8 transition-all duration-300 group-hover:scale-105"
                   :style="{
                     backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                     borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                   }">
                <div class="w-16 h-16 bg-gradient-to-r from-green-500 to-blue-500 rounded-2xl flex items-center justify-center mb-6">
                  <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                  </svg>
                </div>
                <h3 class="text-xl font-bold mb-4" :style="{ color: isDark ? '#ffffff' : '#000000' }">
                  Lightning Fast
                </h3>
                <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                  Optimized for speed with parallel uploads and smart compression.
                </p>
              </div>
            </div>

            <!-- Feature 3 -->
            <div class="relative group">
              <div class="absolute inset-0 bg-gradient-to-r from-purple-600/20 to-pink-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              <div class="relative backdrop-blur-xl border rounded-2xl p-8 transition-all duration-300 group-hover:scale-105"
                   :style="{
                     backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                     borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)'
                   }">
                <div class="w-16 h-16 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl flex items-center justify-center mb-6">
                  <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                  </svg>
                </div>
                <h3 class="text-xl font-bold mb-4" :style="{ color: isDark ? '#ffffff' : '#000000' }">
                  No Registration
                </h3>
                <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                  Start sharing immediately without creating accounts or providing personal info.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Footer -->
      <footer class="relative py-12 px-4 sm:px-6 lg:px-8 border-t"
              :style="{ borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)' }">
        <div class="max-w-7xl mx-auto text-center">
          <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
            © 2025 PInGO Share. Built with ❤️ for the community.
          </p>
        </div>
      </footer>
    </main>

    <!-- Messages -->
    <div v-if="message.text" 
         class="fixed bottom-6 right-6 p-4 rounded-2xl shadow-2xl z-50 backdrop-blur-xl border"
         :style="{
           backgroundColor: message.type === 'success' 
             ? (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)')
             : (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)'),
           borderColor: message.type === 'success' ? '#22c55e' : '#ef4444',
           color: message.type === 'success' ? '#22c55e' : '#ef4444'
         }">
      {{ message.text }}
    </div>

    <!-- GitHub Icon - Bottom Left -->
    <a 
      href="https://github.com/R8bert/PInGO-Share" 
      target="_blank" 
      rel="noopener noreferrer" 
      class="fixed bottom-4 left-4 z-40"
      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }"
      title="View GitHub Repository"
    >
      <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
      </svg>
    </a>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
import axios from 'axios'
import { CloudArrowUpIcon, XMarkIcon } from '@heroicons/vue/24/outline'
import WebGLBackground from '../../components/WebGLBackground.vue'

const { isAuthenticated } = useAuth()
const { isDark } = useTheme()

// Refs
const fileInput = ref<HTMLInputElement | null>(null)
const uploadSection = ref<HTMLElement | null>(null)
const featuresSection = ref<HTMLElement | null>(null)

// State
const selectedFiles = ref<File[]>([])
const isDragging = ref(false)
const isUploading = ref(false)
const message = ref({ text: '', type: 'success' as 'success' | 'error' })

// Methods
const triggerFileInput = () => {
  fileInput.value?.click()
}

const onFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    selectedFiles.value = Array.from(target.files)
  }
}

const onDrop = (event: DragEvent) => {
  isDragging.value = false
  if (event.dataTransfer?.files) {
    selectedFiles.value = Array.from(event.dataTransfer.files)
  }
}

const removeFile = (index: number) => {
  selectedFiles.value.splice(index, 1)
}

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const getFileExtension = (file: File): string => {
  return file.name.split('.').pop()?.toLowerCase() || ''
}

const uploadFiles = async () => {
  if (selectedFiles.value.length === 0) return
  
  isUploading.value = true
  
  try {
    const formData = new FormData()
    selectedFiles.value.forEach(file => {
      formData.append('files', file)
    })

    const response = await axios.post('http://localhost:8080/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        ...(isAuthenticated.value && localStorage.getItem('auth_token') 
          ? { 'Authorization': `Bearer ${localStorage.getItem('auth_token')}` }
          : {})
      },
      withCredentials: true
    })

    if (response.data.uploadId) {
      message.value = { text: 'Files uploaded successfully!', type: 'success' }
      selectedFiles.value = []
      
      // Redirect to download page
      setTimeout(() => {
        window.location.href = `/download/${response.data.uploadId}`
      }, 1500)
    }
  } catch (error) {
    console.error('Upload error:', error)
    message.value = { text: 'Upload failed. Please try again.', type: 'error' }
  } finally {
    isUploading.value = false
    
    // Clear message after 3 seconds
    setTimeout(() => {
      message.value = { text: '', type: 'success' }
    }, 3000)
  }
}

const scrollToUpload = () => {
  uploadSection.value?.scrollIntoView({ behavior: 'smooth' })
}

const scrollToFeatures = () => {
  featuresSection.value?.scrollIntoView({ behavior: 'smooth' })
}

onMounted(() => {
  // Any initialization code
})
</script>

<style scoped>
@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fade-in-delay {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes fade-in-delay-2 {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
  animation: fade-in-delay 0.8s ease-out 0.2s both;
}

.animate-fade-in-delay-2 {
  animation: fade-in-delay-2 0.8s ease-out 0.4s both;
}
</style>
