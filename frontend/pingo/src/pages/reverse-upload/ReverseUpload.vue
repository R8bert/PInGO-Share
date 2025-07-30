<template>
  <div class="min-h-screen bg-gray-100 py-8 px-4">
    <div class="max-w-2xl mx-auto">
      <div class="bg-white rounded-lg shadow-lg p-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-6 text-center">Upload Files to PInGO Share</h1>
        <p class="text-gray-600 mb-8 text-center">Upload files using a reverse share token</p>
        
        <form @submit.prevent="handleUpload" class="space-y-6">
          <div>
            <label for="token" class="block text-sm font-medium text-gray-700 mb-2">Reverse Share Token</label>
            <input 
              v-model="token"
              type="text" 
              id="token" 
              placeholder="Enter your reverse share token"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              required
              :disabled="isUploading"
            >
          </div>
          
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-2">Your Email (optional)</label>
            <input 
              v-model="email"
              type="email" 
              id="email" 
              placeholder="your.email@example.com"
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              :disabled="isUploading"
            >
          </div>
          
          <div>
            <label for="validity" class="block text-sm font-medium text-gray-700 mb-2">File Expiration</label>
            <select 
              v-model="validity"
              id="validity" 
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              :disabled="isUploading"
            >
              <option value="7days">7 Days</option>
              <option value="1month">1 Month</option>
              <option value="6months">6 Months</option>
              <option value="1year">1 Year</option>
              <option value="never">Never</option>
            </select>
          </div>
          
          <div>
            <label for="files" class="block text-sm font-medium text-gray-700 mb-2">Select Files</label>
            <input 
              @change="handleFileSelect"
              type="file" 
              id="files" 
              multiple
              class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              required
              :disabled="isUploading"
            >
          </div>
          
          <button 
            type="submit" 
            class="w-full bg-blue-600 hover:bg-blue-700 disabled:bg-blue-400 text-white font-medium py-3 px-4 rounded-lg transition-colors"
            :disabled="isUploading || !selectedFiles.length || !token"
          >
            {{ isUploading ? 'Uploading...' : 'Upload Files' }}
          </button>
        </form>
        
        <!-- Upload Progress -->
        <div v-if="isUploading" class="mt-6 bg-blue-50 border border-blue-200 rounded-lg p-4">
          <div class="flex items-center">
            <div class="animate-spin rounded-full h-5 w-5 border-b-2 border-blue-600"></div>
            <span class="ml-3 text-blue-800">Uploading files...</span>
          </div>
        </div>
        
        <!-- Success Result -->
        <div v-if="uploadResult && uploadResult.success" class="mt-6 p-4 bg-green-100 border border-green-400 rounded-lg">
          <h3 class="text-green-800 font-medium mb-2">Upload Successful!</h3>
          <p class="text-green-700">Files uploaded successfully.</p>
          <p class="text-sm text-green-600 mt-2">Download URL: <code class="bg-green-200 px-1 rounded">{{ uploadResult.data.download_url }}</code></p>
          <p class="text-sm text-green-600">Files: {{ uploadResult.data.files.join(', ') }}</p>
          <p class="text-sm text-green-600">Total files: {{ uploadResult.data.count }}</p>
          <p class="text-sm text-green-600">
            {{ uploadResult.data.expires_at ? `Expires: ${new Date(uploadResult.data.expires_at).toLocaleDateString()}` : 'Never expires' }}
          </p>
        </div>
        
        <!-- Error Result -->
        <div v-if="uploadResult && !uploadResult.success" class="mt-6 p-4 bg-red-100 border border-red-400 rounded-lg">
          <h3 class="text-red-800 font-medium mb-2">Upload Failed</h3>
          <p class="text-red-700">{{ uploadResult.error }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const token = ref('')
const email = ref('')
const validity = ref('7days')
const selectedFiles = ref<File[]>([])
const isUploading = ref(false)
const uploadResult = ref<{ success: boolean; data?: any; error?: string } | null>(null)

onMounted(() => {
  // Pre-fill token from URL parameter if provided
  if (route.params.token) {
    token.value = route.params.token as string
  }
})

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
