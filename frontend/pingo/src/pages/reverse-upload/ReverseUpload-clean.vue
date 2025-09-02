<template>
  <div class="min-h-screen" :class="isDark ? 'bg-gray-900' : 'bg-gray-50'">
    <div class="py-8 px-4">
      <div class="max-w-2xl mx-auto">
        
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="w-12 h-12 mx-auto mb-4 rounded-xl bg-blue-500 flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"/>
            </svg>
          </div>
          <h1 class="text-2xl font-semibold mb-2" :class="isDark ? 'text-white' : 'text-gray-900'">
            Reverse Share Upload
          </h1>
          <p class="text-sm" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
            Upload files using your reverse share token
          </p>
        </div>

        <!-- Upload Form -->
        <div class="rounded-lg border p-6"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'">
          
          <form @submit.prevent="handleUpload" class="space-y-6">
            
            <!-- Token Input -->
            <div>
              <label for="token" class="block text-sm font-medium mb-2" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                Reverse Share Token
              </label>
              <input 
                v-model="token"
                type="text" 
                id="token" 
                placeholder="Enter your reverse share token"
                class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                :class="isDark 
                  ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400' 
                  : 'bg-gray-50 border-gray-300 text-gray-900 placeholder-gray-500'"
                required
              />
            </div>

            <!-- Email Input -->
            <div>
              <label for="email" class="block text-sm font-medium mb-2" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                Email Address <span class="text-xs text-gray-500">(optional)</span>
              </label>
              <input 
                v-model="email"
                type="email" 
                id="email" 
                placeholder="your.email@example.com"
                class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                :class="isDark 
                  ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400' 
                  : 'bg-gray-50 border-gray-300 text-gray-900 placeholder-gray-500'"
              />
            </div>

            <!-- Expiration Selector -->
            <div>
              <label class="block text-sm font-medium mb-3" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                File Expiration
              </label>
              <div class="grid grid-cols-3 gap-3">
                <button
                  v-for="option in expirationOptions"
                  :key="option.value"
                  type="button"
                  @click="validity = option.value"
                  class="px-4 py-2 rounded-lg text-sm font-medium border transition-colors"
                  :class="validity === option.value
                    ? 'bg-blue-500 text-white border-blue-500'
                    : (isDark 
                        ? 'bg-gray-700 text-gray-300 border-gray-600 hover:bg-gray-600' 
                        : 'bg-gray-50 text-gray-700 border-gray-300 hover:bg-gray-100')"
                >
                  {{ option.label }}
                </button>
              </div>
            </div>

            <!-- File Upload -->
            <div>
              <label class="block text-sm font-medium mb-3" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                Select Files
              </label>
              <div class="border-2 border-dashed rounded-lg p-8 text-center"
                   :class="isDark 
                     ? 'border-gray-600 bg-gray-700/50' 
                     : 'border-gray-300 bg-gray-50'">
                <input
                  type="file"
                  id="files"
                  multiple
                  @change="handleFileSelect"
                  class="hidden"
                />
                <label for="files" class="cursor-pointer">
                  <div class="w-12 h-12 mx-auto mb-4 rounded-lg bg-blue-100 flex items-center justify-center"
                       :class="isDark ? 'bg-blue-900/50' : 'bg-blue-100'">
                    <svg class="w-6 h-6 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
                    </svg>
                  </div>
                  <p class="text-sm font-medium mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                    Click to upload files
                  </p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-500'">
                    or drag and drop files here
                  </p>
                </label>
              </div>
              
              <!-- Selected Files -->
              <div v-if="selectedFiles.length > 0" class="mt-4">
                <p class="text-sm font-medium mb-2" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                  Selected Files ({{ selectedFiles.length }})
                </p>
                <div class="space-y-2">
                  <div v-for="file in selectedFiles" :key="file.name"
                       class="flex items-center justify-between p-3 rounded-lg border"
                       :class="isDark ? 'bg-gray-700 border-gray-600' : 'bg-gray-50 border-gray-200'">
                    <div class="flex items-center gap-3">
                      <div class="w-8 h-8 rounded bg-blue-100 flex items-center justify-center"
                           :class="isDark ? 'bg-blue-900/50' : 'bg-blue-100'">
                        <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                        </svg>
                      </div>
                      <div>
                        <p class="text-sm font-medium" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                          {{ file.name }}
                        </p>
                        <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-500'">
                          {{ formatFileSize(file.size) }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="!token || selectedFiles.length === 0 || isUploading"
              class="w-full py-3 px-4 rounded-lg font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              :class="isDark 
                ? 'bg-blue-600 hover:bg-blue-700 text-white' 
                : 'bg-blue-500 hover:bg-blue-600 text-white'"
            >
              <span v-if="isUploading">Uploading...</span>
              <span v-else>Upload Files</span>
            </button>
          </form>
        </div>

        <!-- Upload Progress -->
        <div v-if="isUploading" class="mt-6 p-4 rounded-lg border"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-blue-50 border-blue-200'">
          <div class="flex items-center gap-3">
            <div class="w-5 h-5 border-2 border-blue-500/30 border-t-blue-500 rounded-full animate-spin"></div>
            <span class="text-sm" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
              Uploading files...
            </span>
          </div>
        </div>

        <!-- Success Result -->
        <div v-if="uploadResult && uploadResult.success" class="mt-6 p-6 rounded-lg border"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-green-50 border-green-200'">
          <div class="flex items-start gap-4">
            <div class="w-10 h-10 rounded-lg bg-green-500 flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="font-medium mb-3" :class="isDark ? 'text-green-400' : 'text-green-800'">
                Upload Successful
              </h3>
              <div class="space-y-3 text-sm" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <p class="font-medium">Files Uploaded</p>
                    <p class="text-xs">{{ uploadResult.data.files.join(', ') }}</p>
                  </div>
                  <div>
                    <p class="font-medium">Total Files</p>
                    <p class="text-xs">{{ uploadResult.data.count }}</p>
                  </div>
                  <div>
                    <p class="font-medium">Expires</p>
                    <p class="text-xs">{{ uploadResult.data.expires_at ? new Date(uploadResult.data.expires_at).toLocaleDateString() : 'Never' }}</p>
                  </div>
                </div>
                <div class="p-3 rounded border" :class="isDark ? 'bg-gray-900 border-gray-600' : 'bg-white border-gray-300'">
                  <p class="font-medium mb-1">Download URL</p>
                  <code class="text-xs break-all" :class="isDark ? 'text-green-400' : 'text-green-600'">
                    {{ uploadResult.data.download_url }}
                  </code>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Error Result -->
        <div v-if="uploadResult && !uploadResult.success" class="mt-6 p-6 rounded-lg border"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-red-50 border-red-200'">
          <div class="flex items-start gap-4">
            <div class="w-10 h-10 rounded-lg bg-red-500 flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </div>
            <div>
              <h3 class="font-medium mb-2" :class="isDark ? 'text-red-400' : 'text-red-800'">
                Upload Failed
              </h3>
              <p class="text-sm" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                {{ uploadResult.error }}
              </p>
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
/* Simple transitions */
* {
  transition: all 0.2s ease;
}
</style>
