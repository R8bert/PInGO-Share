<template>
  <div class="space-y-6">
    <!-- Filter Controls -->
    <div v-if="!isLoading && uploads.length > 0" class="bg-white dark:bg-gray-800 rounded-lg border p-4 shadow-sm"
         :style="{
           backgroundColor: isDark ? '#1f2937' : '#ffffff',
           borderColor: isDark ? '#374151' : '#e5e7eb'
         }">
      <div class="flex flex-wrap items-center justify-between gap-4">
        <div class="flex flex-wrap items-center gap-4">
          <!-- Filter by Status -->
          <div class="flex items-center gap-3">
            <label class="text-sm font-medium"
                   :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
              Filter:
            </label>
            <select 
              v-model="deletionFilter"
              @change="onFilterChange"
              class="px-3 py-2 rounded-lg border text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
              :style="{ 
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#4b5563' : '#d1d5db',
                color: isDark ? '#f9fafb' : '#111827'
              }">
              <option value="all">All Files</option>
              <option value="active">Active Files</option>
              <option value="deleted">Deleted Files</option>
            </select>
          </div>

          <!-- Items per page -->
          <div class="flex items-center gap-3">
            <label class="text-sm font-medium"
                   :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
              Show:
            </label>
            <select 
              v-model="itemsPerPage"
              @change="onItemsPerPageChange"
              class="px-3 py-2 rounded-lg border text-sm font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
              :style="{ 
                backgroundColor: isDark ? '#374151' : '#ffffff',
                borderColor: isDark ? '#4b5563' : '#d1d5db',
                color: isDark ? '#f9fafb' : '#111827'
              }">
              <option value="10">10</option>
              <option value="25">25</option>
              <option value="50">50</option>
              <option value="100">100</option>
            </select>
          </div>
        </div>

        <!-- Results Info -->
        <div class="px-3 py-2 bg-blue-50 dark:bg-blue-900/20 text-blue-700 dark:text-blue-300 rounded-lg border border-blue-200 dark:border-blue-700">
          <div class="flex items-center space-x-2">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <span class="text-sm font-medium">{{ displayedUploads.length }} of {{ filteredUploads.length }} uploads</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex justify-center py-20">
      <div class="text-center">
        <div class="w-12 h-12 border-4 border-blue-200 border-t-blue-600 rounded-full animate-spin mx-auto mb-4"></div>
        <p class="text-lg font-medium" :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Loading your uploads...</p>
      </div>
    </div>
    
    <!-- Empty State -->
    <div v-else-if="uploads.length === 0" class="text-center py-20">
      <div class="w-24 h-24 bg-gradient-to-br from-blue-500 to-purple-600 rounded-2xl flex items-center justify-center mx-auto mb-6 shadow-lg">
        <svg class="w-12 h-12 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
        </svg>
      </div>
      <h3 class="text-2xl font-bold mb-3"
          :style="{ color: isDark ? '#ffffff' : '#000000' }">
        No uploads yet
      </h3>
      <p class="text-lg mb-8"
         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
        Your uploaded files will appear here once you start sharing.
      </p>
      <router-link
        to="/"
        class="inline-flex items-center space-x-2 px-6 py-3 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold rounded-lg transition-all duration-200 shadow-lg hover:shadow-xl"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
        </svg>
        <span>Upload Your First Files</span>
      </router-link>
    </div>

    <!-- No Results -->
    <div v-else-if="displayedUploads.length === 0" class="text-center py-16">
      <div class="w-20 h-20 bg-gray-100 dark:bg-gray-700 rounded-2xl flex items-center justify-center mx-auto mb-6">
        <svg class="w-10 h-10 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </div>
      <h3 class="text-xl font-semibold mb-2"
          :style="{ color: isDark ? '#ffffff' : '#000000' }">
        No uploads found
      </h3>
      <p class="text-lg"
         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
        Try adjusting your filter settings to see more results.
      </p>
    </div>

    <!-- Upload Cards -->
    <div v-else class="space-y-4">
      <div
        v-for="upload in displayedUploads"
        :key="upload.id"
        class="bg-white dark:bg-gray-800 rounded-lg border shadow-sm hover:shadow-md transition-shadow duration-200"
        :style="{
          backgroundColor: isDark ? '#1f2937' : '#ffffff',
          borderColor: isDark ? '#374151' : '#e5e7eb',
          opacity: upload.is_deleted ? '0.6' : '1'
        }">
        
        <div class="p-6">
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center space-x-4 mb-4">
                <!-- File Type Icon -->
                <div class="w-12 h-12 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
                  <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
                  </svg>
                </div>
                
                <!-- Upload Info -->
                <div class="flex-1">
                  <h3 class="text-lg font-semibold mb-1"
                      :style="{ color: isDark ? '#ffffff' : '#000000' }">
                    {{ JSON.parse(upload.files).length }} file{{ JSON.parse(upload.files).length > 1 ? 's' : '' }}
                  </h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    {{ formatFileSize(upload.total_size) }} • Uploaded {{ formatDate(upload.created_at) }}
                  </p>
                </div>
              </div>
              
              <!-- File List -->
              <div class="flex flex-wrap gap-2 mb-4">
                <span
                  v-for="filename in JSON.parse(upload.files)"
                  :key="filename"
                  class="inline-flex items-center px-3 py-1 bg-blue-50 dark:bg-blue-900/20 text-blue-700 dark:text-blue-300 text-sm font-medium rounded-full border border-blue-200 dark:border-blue-700"
                >
                  {{ filename }}
                </span>
              </div>

              <!-- Status and Actions -->
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <!-- Status Badge -->
                  <div v-if="upload.is_deleted" 
                       class="inline-flex items-center px-3 py-1 bg-red-50 dark:bg-red-900/20 text-red-700 dark:text-red-300 text-sm font-medium rounded-full border border-red-200 dark:border-red-700">
                    <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                    Deleted
                  </div>
                  
                  <div v-else 
                       :class="[
                         'inline-flex items-center px-3 py-1 text-sm font-medium rounded-full border',
                         upload.is_available 
                           ? 'bg-green-50 dark:bg-green-900/20 text-green-700 dark:text-green-300 border-green-200 dark:border-green-700'
                           : 'bg-gray-50 dark:bg-gray-900/20 text-gray-700 dark:text-gray-300 border-gray-200 dark:border-gray-700'
                       ]">
                    <div class="w-2 h-2 rounded-full mr-2"
                         :class="upload.is_available ? 'bg-green-500' : 'bg-gray-400'"></div>
                    {{ upload.is_available ? 'Available' : 'Unavailable' }}
                  </div>

                  <!-- Expiration Info -->
                  <div v-if="!upload.is_deleted && upload.expires_at" 
                       :class="[
                         'inline-flex items-center px-3 py-1 text-sm font-medium rounded-full border',
                         isExpiringSoon(upload.expires_at)
                           ? 'bg-orange-50 dark:bg-orange-900/20 text-orange-700 dark:text-orange-300 border-orange-200 dark:border-orange-700'
                           : 'bg-blue-50 dark:bg-blue-900/20 text-blue-700 dark:text-blue-300 border-blue-200 dark:border-blue-700'
                       ]">
                    <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                    </svg>
                    Expires {{ formatDate(upload.expires_at) }}
                  </div>

                  <div v-else-if="!upload.is_deleted && !upload.expires_at" 
                       class="inline-flex items-center px-3 py-1 bg-green-50 dark:bg-green-900/20 text-green-700 dark:text-green-300 text-sm font-medium rounded-full border border-green-200 dark:border-green-700">
                    <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                    </svg>
                    Never expires
                  </div>
                </div>

                <!-- Action Buttons -->
                <div class="flex items-center space-x-2">
                  <!-- Copy Link -->
                  <button v-if="!upload.is_deleted"
                          @click="$emit('copyToClipboard', `${getBaseUrl()}/download/${upload.upload_id}`)"
                          class="p-2 text-gray-400 hover:text-blue-600 hover:bg-blue-50 dark:hover:bg-blue-900/20 rounded-lg transition-colors"
                          title="Copy share link">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                    </svg>
                  </button>

                  <!-- Toggle Availability -->
                  <button v-if="!upload.is_deleted"
                          @click="$emit('toggleAvailability', upload.upload_id, upload.is_available)"
                          :class="[
                            'p-2 rounded-lg transition-colors',
                            upload.is_available
                              ? 'text-gray-400 hover:text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20'
                              : 'text-gray-400 hover:text-green-600 hover:bg-green-50 dark:hover:bg-green-900/20'
                          ]"
                          :title="upload.is_available ? 'Make unavailable' : 'Make available'">
                    <svg v-if="upload.is_available" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                    </svg>
                    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>

                  <!-- Delete Button -->
                  <button @click="$emit('deleteUpload', upload.upload_id)"
                          class="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
                          title="Delete upload">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="filteredUploads.length > itemsPerPage" class="bg-white dark:bg-gray-800 rounded-lg border p-4 shadow-sm"
         :style="{
           backgroundColor: isDark ? '#1f2937' : '#ffffff',
           borderColor: isDark ? '#374151' : '#e5e7eb'
         }">
      <div class="flex items-center justify-between">
        <p class="text-sm text-gray-700 dark:text-gray-300">
          Showing {{ currentPage * itemsPerPage + 1 }} to {{ Math.min((currentPage + 1) * itemsPerPage, filteredUploads.length) }} of {{ filteredUploads.length }} results
        </p>
        
        <div class="flex items-center space-x-2">
          <button @click="currentPage = 0"
                  :disabled="currentPage === 0"
                  class="px-4 py-2 text-sm font-medium border rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  :style="{
                    backgroundColor: isDark ? '#374151' : '#ffffff',
                    borderColor: isDark ? '#4b5563' : '#d1d5db',
                    color: isDark ? '#f9fafb' : '#111827'
                  }">
            First
          </button>
          
          <button @click="currentPage--"
                  :disabled="currentPage === 0"
                  class="px-4 py-2 text-sm font-medium border rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  :style="{
                    backgroundColor: isDark ? '#374151' : '#ffffff',
                    borderColor: isDark ? '#4b5563' : '#d1d5db',
                    color: isDark ? '#f9fafb' : '#111827'
                  }">
            Previous
          </button>
          
          <span class="px-4 py-2 text-sm font-medium bg-blue-50 dark:bg-blue-900/20 text-blue-700 dark:text-blue-300 rounded-lg border border-blue-200 dark:border-blue-700">
            {{ currentPage + 1 }} of {{ totalPages }}
          </span>
          
          <button @click="currentPage++"
                  :disabled="currentPage >= totalPages - 1"
                  class="px-4 py-2 text-sm font-medium border rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  :style="{
                    backgroundColor: isDark ? '#374151' : '#ffffff',
                    borderColor: isDark ? '#4b5563' : '#d1d5db',
                    color: isDark ? '#f9fafb' : '#111827'
                  }">
            Next
          </button>
          
          <button @click="currentPage = totalPages - 1"
                  :disabled="currentPage >= totalPages - 1"
                  class="px-4 py-2 text-sm font-medium border rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                  :style="{
                    backgroundColor: isDark ? '#374151' : '#ffffff',
                    borderColor: isDark ? '#4b5563' : '#d1d5db',
                    color: isDark ? '#f9fafb' : '#111827'
                  }">
            Last
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useTheme } from '../../../composables/useTheme'

const { isDark } = useTheme()

// Props
interface Props {
  uploads: any[]
  isLoading: boolean
}

const props = defineProps<Props>()

// Emits
defineEmits<{
  copyToClipboard: [text: string]
  toggleAvailability: [uploadId: string, currentAvailability: boolean]
  deleteUpload: [uploadId: string]
}>()

// State
const deletionFilter = ref('all')
const itemsPerPage = ref(10)
const currentPage = ref(0)

// Computed
const filteredUploads = computed(() => {
  if (deletionFilter.value === 'deleted') {
    return props.uploads.filter(upload => upload.is_deleted)
  } else if (deletionFilter.value === 'active') {
    return props.uploads.filter(upload => !upload.is_deleted)
  }
  return props.uploads
})

const displayedUploads = computed(() => {
  const start = currentPage.value * itemsPerPage.value
  const end = start + itemsPerPage.value
  return filteredUploads.value.slice(start, end)
})

const totalPages = computed(() => {
  return Math.ceil(filteredUploads.value.length / itemsPerPage.value)
})

// Methods
const onFilterChange = () => {
  currentPage.value = 0
}

const onItemsPerPageChange = () => {
  currentPage.value = 0
}

const formatDate = (dateString: string | undefined) => {
  if (!dateString) return 'Unknown'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i]
}

const isExpiringSoon = (expiresAt: string) => {
  const expirationDate = new Date(expiresAt)
  const now = new Date()
  const diffTime = expirationDate.getTime() - now.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  return diffDays <= 7 && diffDays > 0
}

const getBaseUrl = () => {
  return window.location.origin
}

// Watch for uploads changes to reset pagination
watch(() => props.uploads, () => {
  currentPage.value = 0
})
</script>

<style scoped>
/* Clean, simple styling */
.transition-shadow {
  transition: box-shadow 0.2s ease;
}
</style>
