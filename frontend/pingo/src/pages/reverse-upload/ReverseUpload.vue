<template>
  <div class="min-h-screen" :class="isDark ? 'bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900' : 'bg-gradient-to-br from-blue-50 via-white to-indigo-50'">
    <!-- Subtle Background Pattern -->
    <div class="absolute inset-0 bg-grid-pattern opacity-5"></div>
    
    <div class="relative py-12 px-4">
      <div class="max-w-2xl mx-auto">
        
        <!-- Modern Header -->
        <div class="text-center mb-10 animate-fade-in">
          <div class="relative inline-flex items-center justify-center w-16 h-16 mb-6">
            <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-2xl blur-lg opacity-60"></div>
            <div class="relative w-16 h-16 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-2xl flex items-center justify-center shadow-xl">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"/>
              </svg>
            </div>
          </div>
          <h1 class="text-3xl font-bold mb-3 bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">
            Reverse Share Upload
          </h1>
          <p class="text-lg" :class="isDark ? 'text-gray-300' : 'text-gray-600'">
            Upload files using your reverse share token
          </p>
        </div>

        <!-- Modern Upload Form -->
        <div class="glass-morphism rounded-2xl border p-8 shadow-2xl animate-slide-up"
             :class="isDark ? 'bg-gray-800/80 border-gray-700/50 backdrop-blur-xl' : 'bg-white/80 border-white/50 backdrop-blur-xl'"
             style="animation-delay: 0.2s;">
          
          <form @submit.prevent="handleUpload" class="space-y-8">
            
            <!-- Token Input -->
            <div class="animate-slide-up" style="animation-delay: 0.3s;">
              <label for="token" class="block text-sm font-semibold mb-3" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
                🔐 Reverse Share Token
              </label>
              <div class="relative group">
                <input 
                  v-model="token"
                  type="text" 
                  id="token" 
                  placeholder="Enter your reverse share token"
                  class="w-full px-6 py-4 rounded-xl border-2 transition-all duration-300 focus:outline-none focus:ring-4 focus:ring-blue-500/20 group-hover:border-blue-400"
                  :class="isDark 
                    ? 'bg-gray-700/50 border-gray-600 text-white placeholder-gray-400 focus:border-blue-500 focus:bg-gray-700' 
                    : 'bg-gray-50/50 border-gray-300 text-gray-900 placeholder-gray-500 focus:border-blue-500 focus:bg-white'"
                  required
                />
                <div class="absolute inset-0 rounded-xl bg-gradient-to-r from-blue-500/10 to-indigo-500/10 opacity-0 group-hover:opacity-100 transition-opacity duration-300 pointer-events-none"></div>
              </div>
            </div>

            <!-- Email Input -->
            <div class="animate-slide-up" style="animation-delay: 0.4s;">
              <label for="email" class="block text-sm font-semibold mb-3" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
                📧 Email Address <span class="text-xs font-normal text-blue-500">(optional)</span>
              </label>
              <div class="relative group">
                <input 
                  v-model="email"
                  type="email" 
                  id="email" 
                  placeholder="your.email@example.com"
                  class="w-full px-6 py-4 rounded-xl border-2 transition-all duration-300 focus:outline-none focus:ring-4 focus:ring-blue-500/20 group-hover:border-blue-400"
                  :class="isDark 
                    ? 'bg-gray-700/50 border-gray-600 text-white placeholder-gray-400 focus:border-blue-500 focus:bg-gray-700' 
                    : 'bg-gray-50/50 border-gray-300 text-gray-900 placeholder-gray-500 focus:border-blue-500 focus:bg-white'"
                />
                <div class="absolute inset-0 rounded-xl bg-gradient-to-r from-blue-500/10 to-indigo-500/10 opacity-0 group-hover:opacity-100 transition-opacity duration-300 pointer-events-none"></div>
              </div>
            </div>

            <!-- Expiration Selector -->
            <div class="animate-slide-up" style="animation-delay: 0.5s;">
              <label class="block text-sm font-semibold mb-4" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
                ⏰ File Expiration
              </label>
              <div class="grid grid-cols-3 gap-3">
                <button
                  v-for="(option, index) in expirationOptions"
                  :key="option.value"
                  type="button"
                  @click="validity = option.value"
                  class="relative px-4 py-3 rounded-xl text-sm font-medium transition-all duration-300 transform hover:scale-105 focus:outline-none"
                  :class="validity === option.value
                    ? 'bg-gradient-to-r from-blue-500 to-indigo-600 text-white shadow-lg shadow-blue-500/25'
                    : (isDark 
                        ? 'bg-gray-700/50 text-gray-300 border border-gray-600 hover:bg-gray-600/50 hover:border-blue-500/50' 
                        : 'bg-gray-50/50 text-gray-700 border border-gray-200 hover:bg-white hover:border-blue-300 shadow-sm')"
                  :style="`animation-delay: ${0.6 + index * 0.1}s`"
                >
                  {{ option.label }}
                  <div v-if="validity === option.value" class="absolute inset-0 bg-gradient-to-r from-blue-400/20 to-indigo-400/20 rounded-xl animate-pulse"></div>
                </button>
              </div>
            </div>

            <!-- Modern File Upload -->
            <div class="animate-slide-up" style="animation-delay: 0.8s;">
              <label class="block text-sm font-semibold mb-4" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
                📁 Select Files
              </label>
              <div class="relative">
                <input
                  type="file"
                  id="files"
                  multiple
                  @change="handleFileSelect"
                  class="hidden"
                />
                <label 
                  for="files"
                  class="group relative flex flex-col items-center justify-center w-full h-40 border-2 border-dashed rounded-2xl cursor-pointer transition-all duration-300 hover:scale-[1.02]"
                  :class="isDark 
                    ? 'border-gray-600 bg-gray-700/30 hover:bg-gray-600/30 hover:border-blue-500/50' 
                    : 'border-gray-300 bg-gray-50/30 hover:bg-blue-50/50 hover:border-blue-400'"
                >
                  <!-- Upload Icon with Animation -->
                  <div class="relative mb-4">
                    <div class="absolute inset-0 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-2xl blur-xl opacity-20 group-hover:opacity-40 transition-opacity duration-300"></div>
                    <div class="relative w-12 h-12 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-2xl flex items-center justify-center shadow-lg group-hover:shadow-xl transition-all duration-300 transform group-hover:scale-110">
                      <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
                      </svg>
                    </div>
                  </div>
                  
                  <div class="text-center">
                    <p class="text-lg font-semibold mb-2" :class="isDark ? 'text-white' : 'text-gray-900'">
                      {{ selectedFiles.length > 0 ? `${selectedFiles.length} file(s) selected` : 'Drop files here or click to browse' }}
                    </p>
                    <p class="text-sm" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
                      Support for any file type
                    </p>
                  </div>
                </label>
              </div>
              
              <!-- Selected Files -->
              <div v-if="selectedFiles.length > 0" class="mt-6 space-y-3">
                <p class="text-sm font-semibold" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
                  Selected Files ({{ selectedFiles.length }})
                </p>
                <div class="space-y-2">
                  <div v-for="(file, index) in selectedFiles" :key="file.name"
                       class="flex items-center gap-4 p-4 rounded-xl border transition-all duration-300 hover:shadow-md animate-fade-in"
                       :class="isDark ? 'bg-gray-700/40 border-gray-600/50 hover:bg-gray-600/40' : 'bg-white/60 border-gray-200/60 hover:bg-white/80'"
                       :style="`animation-delay: ${0.9 + index * 0.1}s`">
                    <div class="w-10 h-10 rounded-xl bg-gradient-to-r from-blue-500 to-indigo-600 flex items-center justify-center shadow-md">
                      <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                      </svg>
                    </div>
                    <div class="flex-1 min-w-0">
                      <p class="text-sm font-semibold truncate" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
                        {{ file.name }}
                      </p>
                      <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-500'">
                        {{ formatFileSize(file.size) }}
                      </p>
                    </div>
                    <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Modern Submit Button -->
            <button
              type="submit"
              :disabled="!token || selectedFiles.length === 0 || isUploading"
              class="w-full py-4 px-6 rounded-xl font-semibold text-white transition-all duration-300 transform hover:scale-[1.02] focus:outline-none focus:ring-4 focus:ring-blue-500/20 disabled:scale-100 disabled:opacity-50 disabled:cursor-not-allowed relative overflow-hidden group animate-slide-up"
              :class="!token || selectedFiles.length === 0 || isUploading
                ? 'bg-gray-400'
                : 'bg-gradient-to-r from-blue-600 to-indigo-600 hover:from-blue-700 hover:to-indigo-700 shadow-lg hover:shadow-xl shadow-blue-500/25'"
              style="animation-delay: 1.2s;"
            >
              <!-- Shimmer Effect -->
              <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
              
              <span v-if="isUploading" class="flex items-center justify-center gap-3">
                <div class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                Uploading files...
              </span>
              <span v-else class="flex items-center justify-center gap-3">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"/>
                </svg>
                Upload Files Securely
              </span>
            </button>
          </form>
        </div>

        <!-- Modern Upload Progress -->
        <div v-if="isUploading" class="mt-8 p-6 rounded-2xl border backdrop-blur-xl animate-fade-in"
             :class="isDark ? 'bg-gray-800/80 border-gray-700/50' : 'bg-blue-50/80 border-blue-200/50'">
          <div class="flex items-center gap-4">
            <div class="relative">
              <div class="w-12 h-12 border-4 border-blue-500/30 border-t-blue-500 rounded-full animate-spin"></div>
              <div class="absolute inset-0 w-12 h-12 border-4 border-transparent border-t-indigo-500 rounded-full animate-spin" style="animation-direction: reverse; animation-duration: 1.5s;"></div>
            </div>
            <div>
              <p class="font-semibold text-lg" :class="isDark ? 'text-blue-400' : 'text-blue-800'">
                Uploading files...
              </p>
              <p class="text-sm" :class="isDark ? 'text-blue-300' : 'text-blue-600'">
                Please wait while we securely process your files
              </p>
            </div>
          </div>
        </div>

        <!-- Modern Success Result -->
        <div v-if="uploadResult && uploadResult.success" class="mt-8 p-8 rounded-2xl border backdrop-blur-xl animate-fade-in"
             :class="isDark ? 'bg-gray-800/80 border-gray-700/50' : 'bg-green-50/80 border-green-200/50'">
          <div class="flex items-start gap-6">
            <div class="w-14 h-14 rounded-2xl bg-gradient-to-r from-green-500 to-emerald-600 flex items-center justify-center shadow-xl">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="font-bold text-xl mb-4" :class="isDark ? 'text-green-400' : 'text-green-800'">
                🎉 Upload Successful!
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                <div class="p-4 rounded-xl backdrop-blur-sm" :class="isDark ? 'bg-gray-700/50' : 'bg-white/60'">
                  <p class="font-semibold text-sm mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">Files Uploaded</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-600'">{{ uploadResult.data.files.join(', ') }}</p>
                </div>
                <div class="p-4 rounded-xl backdrop-blur-sm" :class="isDark ? 'bg-gray-700/50' : 'bg-white/60'">
                  <p class="font-semibold text-sm mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">Total Files</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-600'">{{ uploadResult.data.count }}</p>
                </div>
                <div class="p-4 rounded-xl backdrop-blur-sm" :class="isDark ? 'bg-gray-700/50' : 'bg-white/60'">
                  <p class="font-semibold text-sm mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">Expires</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-600'">{{ uploadResult.data.expires_at ? new Date(uploadResult.data.expires_at).toLocaleDateString() : 'Never' }}</p>
                </div>
              </div>
              <div class="p-6 rounded-2xl border-2 border-dashed backdrop-blur-sm" :class="isDark ? 'bg-gray-700/30 border-green-500/30' : 'bg-white/80 border-green-300/50'">
                <p class="text-sm font-semibold mb-3" :class="isDark ? 'text-gray-300' : 'text-gray-700'">📎 Download URL</p>
                <code class="text-sm font-mono break-all block p-3 rounded-lg" 
                      :class="isDark ? 'bg-gray-800/50 text-green-400' : 'bg-gray-100 text-green-600'">
                  {{ uploadResult.data.download_url }}
                </code>
              </div>
            </div>
          </div>
        </div>

        <!-- Modern Error Result -->
        <div v-if="uploadResult && !uploadResult.success" class="mt-8 p-8 rounded-2xl border backdrop-blur-xl animate-fade-in"
             :class="isDark ? 'bg-gray-800/80 border-gray-700/50' : 'bg-red-50/80 border-red-200/50'">
          <div class="flex items-start gap-6">
            <div class="w-14 h-14 rounded-2xl bg-gradient-to-r from-red-500 to-pink-600 flex items-center justify-center shadow-xl">
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <div>
              <h3 class="font-bold text-xl mb-3" :class="isDark ? 'text-red-400' : 'text-red-800'">
                ❌ Upload Failed
              </h3>
              <p class="text-sm mb-4" :class="isDark ? 'text-red-300' : 'text-red-700'">
                {{ uploadResult.error }}
              </p>
              <div class="p-4 rounded-xl backdrop-blur-sm" :class="isDark ? 'bg-gray-700/30' : 'bg-white/60'">
                <p class="text-xs font-medium" :class="isDark ? 'text-red-400' : 'text-red-600'">
                  💡 Please check your reverse token and try again
                </p>
              </div>
            </div>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useTheme } from '../../composables/useTheme'

const { isDark } = useTheme()
const route = useRoute()

const token = ref('')
const email = ref('')
const validity = ref('7days')
const selectedFiles = ref<File[]>([])
const isUploading = ref(false)
const uploadResult = ref<{ success: boolean; data?: any; error?: string } | null>(null)

const expirationOptions = [
  { value: '7days', label: '7 Days' },
  { value: '1month', label: '1 Month' },
  { value: '6months', label: '6 Months' },
  { value: '1year', label: '1 Year' },
  { value: 'never', label: 'Never' }
]

onMounted(() => {
  // Pre-fill token from URL parameter if provided
  if (route.params.token) {
    token.value = route.params.token as string
  }
})

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    selectedFiles.value = Array.from(target.files)
  }
}

const handleUpload = async () => {
  if (!token.value || !selectedFiles.value.length) {
    alert('Please enter a token and select files')
    return
  }

  isUploading.value = true
  uploadResult.value = null

  const formData = new FormData()
  
  // Add files
  for (const file of selectedFiles.value) {
    formData.append('files', file)
  }
  
  // Add metadata
  if (email.value) formData.append('email', email.value)
  formData.append('validity', validity.value)

  try {
    const response = await fetch(`http://localhost:8080/reverse-upload/${token.value}`, {
      method: 'POST',
      body: formData
    })

    const data = await response.json()

    if (response.ok) {
      uploadResult.value = {
        success: true,
        data
      }
      // Reset form
      email.value = ''
      validity.value = '7days'
      selectedFiles.value = []
      // Reset file input
      const fileInput = document.getElementById('files') as HTMLInputElement
      if (fileInput) fileInput.value = ''
    } else {
      throw new Error(data.error || 'Upload failed')
    }
  } catch (error) {
    uploadResult.value = {
      success: false,
      error: error instanceof Error ? error.message : 'Upload failed'
    }
  } finally {
    isUploading.value = false
  }
}
</script>

<style scoped>
/* Background Grid Pattern */
.bg-grid-pattern {
  background-image: 
    linear-gradient(rgba(59, 130, 246, 0.1) 1px, transparent 1px),
    linear-gradient(90deg, rgba(59, 130, 246, 0.1) 1px, transparent 1px);
  background-size: 20px 20px;
}

/* Glass Morphism Effect */
.glass-morphism {
  backdrop-filter: blur(16px) saturate(180%);
  -webkit-backdrop-filter: blur(16px) saturate(180%);
}

/* Modern Animations */
.animate-fade-in {
  animation: fadeIn 0.8s ease-out forwards;
  opacity: 0;
}

@keyframes fadeIn {
  to {
    opacity: 1;
  }
}

.animate-slide-up {
  animation: slideUp 0.8s ease-out forwards;
  opacity: 0;
  transform: translateY(30px);
}

@keyframes slideUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Hover Effects */
.group:hover .w-12.h-12 {
  transform: scale(1.1);
}

/* Input Focus Effects */
input:focus {
  transform: translateY(-2px);
}

/* Button Hover Effects */
button:hover:not(:disabled) {
  transform: translateY(-2px);
}

button:active:not(:disabled) {
  transform: translateY(0);
}

/* File Item Animation */
.animate-fade-in {
  animation: fadeInScale 0.6s ease-out forwards;
  opacity: 0;
  transform: scale(0.95);
}

@keyframes fadeInScale {
  to {
    opacity: 1;
    transform: scale(1);
  }
}

/* Upload Zone Hover */
.group:hover {
  transform: translateY(-2px);
}

/* Smooth Transitions */
* {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Responsive Design */
@media (max-width: 640px) {
  .grid-cols-3 {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

/* Animation Delays */
[style*="animation-delay"] {
  animation-fill-mode: both;
}

/* Gradient Text */
.bg-gradient-to-r.bg-clip-text {
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

/* Pulse Animation for Dots */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

/* Loading Spinner Enhancement */
.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
