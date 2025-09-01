<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900 p-6">
    <!-- Header Section -->
    <div class="max-w-7xl mx-auto">
      <!-- Page Title and Stats -->
      <div class="mb-8">
        <div class="flex items-center justify-between mb-6">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">My Uploads</h1>
            <p class="text-gray-600 dark:text-gray-400 mt-1">Manage and track your file uploads</p>
          </div>
          
          <!-- Quick Stats Cards -->
          <div class="flex gap-4">
            <div class="bg-white dark:bg-gray-800 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-700">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-blue-100 dark:bg-blue-900/30 rounded-lg flex items-center justify-center">
                  <img src="/src/assets/svg/icons/files_uploaded.svg" alt="Total uploads" 
                       class="w-6 h-6 object-contain" 
                       style="width: 24px !important; height: 24px !important;" />
                </div>
                <div>
                  <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ uploads.length }}</div>
                  <div class="text-sm text-gray-500 dark:text-gray-400">Total Uploads</div>
                </div>
              </div>
            </div>
            
            <div class="bg-white dark:bg-gray-800 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-700">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-green-100 dark:bg-green-900/30 rounded-lg flex items-center justify-center">
                  <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div>
                  <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ activeUploads.length }}</div>
                  <div class="text-sm text-gray-500 dark:text-gray-400">Active Files</div>
                </div>
              </div>
            </div>
            
            <div class="bg-white dark:bg-gray-800 rounded-xl p-4 shadow-sm border border-gray-200 dark:border-gray-700">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-purple-100 dark:bg-purple-900/30 rounded-lg flex items-center justify-center">
                  <img src="/src/assets/svg/icons/total_files_icon.svg" alt="Total files" 
                       class="w-6 h-6 object-contain" 
                       style="width: 24px !important; height: 24px !important;" />
                </div>
                <div>
                  <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ totalFilesCount }}</div>
                  <div class="text-sm text-gray-500 dark:text-gray-400">Total Files</div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Filter Tabs -->
        <div v-if="!isLoading && uploads.length > 0" class="flex gap-2">
          <button 
            v-for="filter in filterOptions" 
            :key="filter.value"
            @click="deletionFilter = filter.value; onFilterChange()"
            :class="[
              'px-4 py-2 rounded-lg font-medium transition-all duration-200 flex items-center gap-2',
              deletionFilter === filter.value 
                ? 'bg-blue-600 text-white shadow-md' 
                : 'bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 border border-gray-200 dark:border-gray-700'
            ]"
          >
            {{ filter.label }}
            <span v-if="filter.count !== undefined" 
                  :class="[
                    'px-2 py-0.5 text-xs rounded-full',
                    deletionFilter === filter.value 
                      ? 'bg-white/20 text-white' 
                      : 'bg-gray-100 dark:bg-gray-700 text-gray-600 dark:text-gray-400'
                  ]">
              {{ filter.count }}
            </span>
          </button>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="flex items-center justify-center py-20">
        <div class="text-center">
          <div class="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p class="text-gray-600 dark:text-gray-400">Loading your uploads...</p>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredUploads.length === 0" class="text-center py-20">
        <div class="w-24 h-24 bg-gray-100 dark:bg-gray-800 rounded-full flex items-center justify-center mx-auto mb-6">
          <img src="/src/assets/svg/icons/files_uploaded2.svg" alt="No uploads" 
               class="w-12 h-12 opacity-50 object-contain" 
               style="width: 48px !important; height: 48px !important;" />
        </div>
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">No uploads found</h3>
        <p class="text-gray-600 dark:text-gray-400 max-w-md mx-auto">
          {{ deletionFilter === 'all' ? 'You haven\'t uploaded any files yet.' : `No ${deletionFilter} uploads found.` }}
        </p>
      </div>

      <!-- Uploads List -->
      <div v-else class="space-y-4">
        <transition-group name="list" tag="div" class="space-y-4">
          <div
            v-for="upload in displayedUploads"
            :key="upload.upload_id"
            class="group bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 hover:shadow-md transition-all duration-200 overflow-hidden"
            :class="selectedUpload?.id === upload.id ? 'ring-2 ring-blue-500' : ''"
          >
            <!-- Main Upload Row -->
            <div class="p-6">
              <div class="flex items-center justify-between">
                <!-- Left side: Icon, Title, and Meta -->
                <div class="flex items-center gap-4 flex-1 min-w-0">
                  <!-- File Type Icon -->
                  <div class="flex items-center justify-center flex-shrink-0">
                    <img 
                      :src="getUploadIconPath(upload)" 
                      :alt="getUploadIconAlt(upload)" 
                      class="w-10 h-10 object-contain"
                      style="width: 45px !important; height: 45px !important;"
                    />
                  </div>
                  
                  <!-- Upload Info -->
                  <div class="flex-1 min-w-0">
                    <div class="flex items-center gap-3 mb-1">
                      <h3 class="font-semibold text-gray-900 dark:text-white">Upload {{ upload.upload_id }}</h3>
                      <!-- Status Badge -->
                      <span 
                        :class="[
                          'px-2 py-1 text-xs font-medium rounded-full',
                          upload.is_deleted 
                            ? 'bg-red-100 dark:bg-red-900/30 text-red-800 dark:text-red-400'
                            : !upload.is_available
                              ? 'bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-400'
                              : upload.expires_at && isExpiringSoon(upload.expires_at)
                                ? 'bg-yellow-100 dark:bg-yellow-900/30 text-yellow-800 dark:text-yellow-400'
                                : 'bg-green-100 dark:bg-green-900/30 text-green-800 dark:text-green-400'
                        ]"
                      >
                        {{ getStatusText(upload) }}
                      </span>
                    </div>
                    
                    <!-- Upload Metadata -->
                    <div class="flex items-center gap-6 text-sm text-gray-500 dark:text-gray-400">
                      <span>{{ JSON.parse(upload.files).length }} files</span>
                      <span>{{ formatFileSize(upload.size || 0) }}</span>
                      <span>Created {{ formatDate(upload.created_at) }}</span>
                      <span v-if="upload.expires_at">Expires {{ formatDate(upload.expires_at) }}</span>
                    </div>
                  </div>
                </div>

                <!-- Right side: Action Buttons -->
                <div class="flex items-center gap-2 flex-shrink-0">
                  <button
                    @click="openFile(upload)"
                    class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                    Open
                  </button>
                  
                  <button
                    @click="$emit('copyToClipboard', `${getBaseUrl()}/download/${upload.upload_id}`)"
                    class="bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 px-3 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center justify-center"
                    title="Copy Link"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                  </button>
                  
                  <button
                    @click="selectUpload(upload)"
                    class="bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 px-3 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center justify-center"
                    title="More Options"
                  >
                    <svg 
                      class="w-4 h-4 transition-transform duration-200" 
                      :class="selectedUpload?.id === upload.id ? 'rotate-180' : ''"
                      fill="none" 
                      stroke="currentColor" 
                      viewBox="0 0 24 24"
                    >
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>

            <!-- Expanded Details (when selected) -->
            <div v-if="selectedUpload?.id === upload.id" class="border-t border-gray-200 dark:border-gray-700 p-6 bg-gray-50 dark:bg-gray-900/50">
              <div class="grid md:grid-cols-2 gap-6">
                <!-- Files List -->
                <div>
                  <h4 class="font-medium text-gray-900 dark:text-white mb-3">Files in this upload:</h4>
                  <div class="space-y-2 max-h-60 overflow-y-auto">
                    <div v-for="(file, fileIndex) in JSON.parse(upload.files)" :key="fileIndex" 
                         class="flex items-center gap-3 p-3 bg-white dark:bg-gray-800 rounded-lg">
                      <div class="flex items-center justify-center flex-shrink-0">
                        <img 
                          :src="getFileIconPath(file)" 
                          :alt="getFileIconAlt(file)" 
                          class="w-6 h-6 object-contain"
                          style="width: 24px !important; height: 24px !important;"
                        />
                      </div>
                      <span class="text-sm text-gray-900 dark:text-white truncate">{{ file }}</span>
                    </div>
                  </div>
                </div>
                
                <!-- Additional Actions -->
                <div>
                  <h4 class="font-medium text-gray-900 dark:text-white mb-3">Actions:</h4>
                  <div class="space-y-3">
                    <button
                      @click="downloadSingleFile(upload)"
                      class="w-full bg-green-600 hover:bg-green-700 text-white px-4 py-3 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                      </svg>
                      Download Files
                    </button>
                    
                    <button
                      v-if="!upload.is_deleted"
                      @click="$emit('toggleAvailability', upload.upload_id, upload.is_available)"
                      :class="[
                        'w-full px-4 py-3 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2',
                        upload.is_available 
                          ? 'bg-orange-600 hover:bg-orange-700 text-white' 
                          : 'bg-blue-600 hover:bg-blue-700 text-white'
                      ]"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path v-if="upload.is_available" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L8.464 8.464M9.878 9.878l-3-3m5.5 5.5l3 3m-3-3l3-3m-3 3l-3 3" />
                        <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        <path v-if="!upload.is_available" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                      </svg>
                      {{ upload.is_available ? 'Hide Upload' : 'Show Upload' }}
                    </button>
                    
                    <button
                      @click="$emit('deleteUpload', upload.upload_id)"
                      class="w-full bg-red-600 hover:bg-red-700 text-white px-4 py-3 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                    >
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                      Delete Upload
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </transition-group>
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="flex justify-center mt-8">
        <div class="flex gap-2">
          <button
            v-for="page in totalPages"
            :key="page"
            @click="currentPage = page - 1"
            :class="[
              'px-4 py-2 rounded-lg font-medium transition-colors duration-200',
              currentPage === page - 1
                ? 'bg-blue-600 text-white'
                : 'bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 border border-gray-200 dark:border-gray-700'
            ]"
          >
            {{ page }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'

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
const itemsPerPage = ref(12)
const currentPage = ref(0)
const selectedUpload = ref<any>(null)

// Dynamic filter options with counts
const filterOptions = computed(() => [
  { 
    value: 'all', 
    label: 'All Files', 
    count: props.uploads.length 
  },
  { 
    value: 'active', 
    label: 'Active', 
    count: props.uploads.filter(u => !u.is_deleted).length 
  },
  { 
    value: 'deleted', 
    label: 'Deleted', 
    count: props.uploads.filter(u => u.is_deleted).length 
  }
])

// Computed stats
const activeUploads = computed(() => props.uploads.filter(upload => !upload.is_deleted))
const totalFilesCount = computed(() => {
  return props.uploads.reduce((total, upload) => {
    return total + JSON.parse(upload.files).length
  }, 0)
})

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

// Helper methods
const getUploadIconPath = (upload: any) => {
  const files = JSON.parse(upload.files)
  
  // If multiple files, use default multiple files icon
  if (files.length > 1) {
    return '/src/assets/svg/icons/default_multiple_files.svg'
  }
  
  // For single file, determine the file type
  const filename = files[0]
  const ext = getFileExtension(filename).toLowerCase()
  
  // Map file extensions to specific SVG icons
  if (['pdf'].includes(ext)) {
    return '/src/assets/svg/icons/files_pdf.svg'
  }
  if (['doc', 'docx'].includes(ext)) {
    return '/src/assets/svg/icons/word_docx_doc.svg'
  }
  if (['ppt', 'pptx'].includes(ext)) {
    return '/src/assets/svg/icons/ppt_pptx_file.svg'
  }
//   video files
  if (['mp4', 'mov', 'avi', 'mkv'].includes(ext)) {
    return '/src/assets/svg/icons/files_video.svg'
  }

  // For file types without specific icons, use the no icon fallback
  return '/src/assets/svg/icons/files_no_icon.svg'
}

const getUploadIconAlt = (upload: any) => {
  const files = JSON.parse(upload.files)
  
  if (files.length > 1) {
    return 'Multiple files'
  }
  
  const filename = files[0]
  const ext = getFileExtension(filename).toLowerCase()
  
  if (['pdf'].includes(ext)) return 'PDF file'
  if (['doc', 'docx'].includes(ext)) return 'Word document'
  if (['ppt', 'pptx'].includes(ext)) return 'PowerPoint presentation'
  
  return 'File'
}

const getFileIconPath = (filename: string) => {
  const ext = getFileExtension(filename).toLowerCase()
  
  // Map file extensions to specific SVG icons
  if (['pdf'].includes(ext)) {
    return '/src/assets/svg/icons/files_pdf.svg'
  }
  if (['doc', 'docx'].includes(ext)) {
    return '/src/assets/svg/icons/word_docx_doc.svg'
  }
  if (['ppt', 'pptx'].includes(ext)) {
    return '/src/assets/svg/icons/ppt_pptx_file.svg'
  }
  
  // For file types without specific icons, use the no icon fallback
  return '/src/assets/svg/icons/files_no_icon.svg'
}

const getFileIconAlt = (filename: string) => {
  const ext = getFileExtension(filename).toLowerCase()
  
  if (['pdf'].includes(ext)) return 'PDF file'
  if (['doc', 'docx'].includes(ext)) return 'Word document'
  if (['ppt', 'pptx'].includes(ext)) return 'PowerPoint presentation'
  
  return 'File'
}

const getFileExtension = (filename: string) => {
  return filename.split('.').pop() || ''
}

const getStatusText = (upload: any) => {
  if (upload.is_deleted) return 'Deleted'
  if (!upload.is_available) return 'Hidden'
  if (upload.expires_at && isExpiringSoon(upload.expires_at)) return 'Expiring Soon'
  return 'Active'
}

const selectUpload = (upload: any) => {
  selectedUpload.value = selectedUpload.value?.id === upload.id ? null : upload
}

const openFile = (upload: any) => {
  const downloadUrl = `${getBaseUrl()}/download/${upload.upload_id}`
  window.open(downloadUrl, '_blank')
}

const downloadSingleFile = async (upload: any) => {
  try {
    const response = await fetch(`/api/download/${upload.upload_id}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const contentDisposition = response.headers.get('Content-Disposition')
    let filename = 'download'
    if (contentDisposition) {
      const matches = contentDisposition.match(/filename="([^"]*)"/)
      if (matches) {
        filename = matches[1]
      }
    } else {
      const files = JSON.parse(upload.files)
      filename = files.length === 1 ? files[0] : `upload_${upload.upload_id}.zip`
    }

    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.style.display = 'none'
    a.href = url
    a.download = filename
    document.body.appendChild(a)
    a.click()
    window.URL.revokeObjectURL(url)
    document.body.removeChild(a)
  } catch (error) {
    console.error('Download failed:', error)
    alert('Download failed. Please try again.')
  }
}

const onFilterChange = () => {
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
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: scale(0.95);
}
</style>
