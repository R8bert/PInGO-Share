<template>
  <div class="min-h-screen p-4 sm:p-6" :class="isDark ? 'bg-gray-900' : 'bg-gray-50'">
    <!-- Header Section -->
    <div class="max-w-7xl mx-auto">
      <!-- Page Title and Stats -->
      <div class="mb-6 sm:mb-8">
        <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between mb-6 gap-6">
          <div class="min-w-0">
            <h1 class="text-2xl sm:text-3xl font-bold truncate" :class="isDark ? 'text-white' : 'text-gray-900'">My Uploads</h1>
            <p class="mt-1 text-sm sm:text-base" :class="isDark ? 'text-gray-400' : 'text-gray-600'">Manage and track your file uploads</p>
          </div>
          
          <!-- Quick Stats Cards -->
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-3 sm:gap-4 w-full lg:w-auto">
            <div class="rounded-xl p-3 sm:p-4 shadow-sm border" 
                 :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'">
              <div class="flex items-center gap-2 sm:gap-3">
                <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-lg flex items-center justify-center flex-shrink-0"
                     :class="isDark ? 'bg-blue-900/30' : 'bg-blue-100'">
                  <img :src="getSystemIcon('filesUploaded')" alt="Total uploads" 
                       class="w-5 h-5 sm:w-6 sm:h-6 object-contain" />
                </div>
                <div class="min-w-0">
                  <div class="text-lg sm:text-2xl font-bold" :class="isDark ? 'text-white' : 'text-gray-900'">{{ uploads.length }}</div>
                  <div class="text-xs sm:text-sm truncate" :class="isDark ? 'text-gray-400' : 'text-gray-500'">Total Uploads</div>
                </div>
              </div>
            </div>
            
            <div class="rounded-xl p-3 sm:p-4 shadow-sm border" 
                 :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'">
              <div class="flex items-center gap-2 sm:gap-3">
                <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-lg flex items-center justify-center flex-shrink-0"
                     :class="isDark ? 'bg-green-900/30' : 'bg-green-100'">
                  <IconCheckCircle class="w-5 h-5 sm:w-6 sm:h-6" :class="isDark ? 'text-green-400' : 'text-green-600'" />
                </div>
                <div class="min-w-0">
                  <div class="text-lg sm:text-2xl font-bold" :class="isDark ? 'text-white' : 'text-gray-900'">{{ activeUploads.length }}</div>
                  <div class="text-xs sm:text-sm truncate" :class="isDark ? 'text-gray-400' : 'text-gray-500'">Active Files</div>
                </div>
              </div>
            </div>
            
            <div class="rounded-xl p-3 sm:p-4 shadow-sm border" 
                 :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'">
              <div class="flex items-center gap-2 sm:gap-3">
                <div class="w-8 h-8 sm:w-10 sm:h-10 rounded-lg flex items-center justify-center flex-shrink-0"
                     :class="isDark ? 'bg-purple-900/30' : 'bg-purple-100'">
                  <img :src="getSystemIcon('totalFiles')" alt="Total files" 
                       class="w-5 h-5 sm:w-6 sm:h-6 object-contain" />
                </div>
                <div class="min-w-0">
                  <div class="text-lg sm:text-2xl font-bold" :class="isDark ? 'text-white' : 'text-gray-900'">{{ totalFilesCount }}</div>
                  <div class="text-xs sm:text-sm truncate" :class="isDark ? 'text-gray-400' : 'text-gray-500'">Total Files</div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Filter Tabs -->
        <div v-if="!isLoading && uploads.length > 0" class="flex flex-wrap gap-2">
          <button 
            v-for="filter in filterOptions" 
            :key="filter.value"
            @click="deletionFilter = filter.value; onFilterChange()"
            :class="[
              'px-3 sm:px-4 py-2 rounded-lg font-medium transition-all duration-200 flex items-center gap-2 text-sm sm:text-base',
              deletionFilter === filter.value 
                ? 'bg-blue-600 text-white shadow-md' 
                : isDark 
                  ? 'bg-gray-800 text-gray-300 hover:bg-gray-700 border border-gray-700'
                  : 'bg-white text-gray-700 hover:bg-gray-50 border border-gray-200'
            ]"
          >
            <span class="truncate">{{ filter.label }}</span>
            <span v-if="filter.count !== undefined" 
                  :class="[
                    'px-2 py-0.5 text-xs rounded-full flex-shrink-0',
                    deletionFilter === filter.value 
                      ? 'bg-white/20 text-white' 
                      : isDark
                        ? 'bg-gray-700 text-gray-400'
                        : 'bg-gray-100 text-gray-600'
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
          <p class="" :class="isDark ? 'text-gray-400' : 'text-gray-600'">Loading your uploads...</p>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else-if="filteredUploads.length === 0" class="text-center py-20">
        <div class="w-24 h-24 rounded-full flex items-center justify-center mx-auto mb-6"
             :class="isDark ? 'bg-gray-800' : 'bg-gray-100'">
          <img :src="getSystemIcon('filesUploaded2')" alt="No uploads" 
               class="w-12 h-12 opacity-50 object-contain" 
               style="width: 48px !important; height: 48px !important;" />
        </div>
        <h3 class="text-xl font-semibold mb-2" :class="isDark ? 'text-white' : 'text-gray-900'">No uploads found</h3>
        <p class="max-w-md mx-auto" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
          {{ deletionFilter === 'all' ? 'You haven\'t uploaded any files yet.' : `No ${deletionFilter} uploads found.` }}
        </p>
      </div>

      <!-- Uploads List -->
      <div v-else class="space-y-4">
        <transition-group name="list" tag="div" class="space-y-4">
          <div
            v-for="upload in displayedUploads"
            :key="upload.upload_id"
            class="group rounded-xl shadow-sm border hover:shadow-md transition-all duration-200 overflow-hidden"
            :class="[
              selectedUpload?.id === upload.id ? 'ring-2 ring-blue-500' : '',
              isDark ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'
            ]"
          >
            <!-- Main Upload Row -->
            <div class="p-4 sm:p-6">
              <div class="flex flex-col sm:flex-row sm:items-center gap-4 overflow-hidden">
                <!-- Left side: Icon, Title, and Meta -->
                <div class="flex items-center gap-3 sm:gap-4 flex-1 min-w-0 overflow-hidden">
                  <!-- File Type Icon -->
                  <div class="flex items-center justify-center flex-shrink-0">
                    <img 
                      :src="getUploadIconPath(upload)" 
                      :alt="getUploadIconAltText(upload)" 
                      class="w-8 h-8 sm:w-10 sm:h-10 object-contain"
                    />
                  </div>
                  
                  <!-- Upload Info -->
                  <div class="flex-1 min-w-0 overflow-hidden">
                    <div class="flex flex-col sm:flex-row sm:items-center gap-2 sm:gap-3 mb-1 min-w-0">
                      <h3 class="font-semibold text-sm sm:text-base truncate min-w-0" :class="isDark ? 'text-white' : 'text-gray-900'">Upload {{ upload.upload_id }}</h3>
                      
                      <!-- Reverse Share Tab -->
                      <span 
                        v-if="upload.is_reverse"
                        class="px-2 py-1 text-xs font-medium rounded-full inline-flex items-center gap-1 w-fit flex-shrink-0"
                        :class="isDark 
                          ? 'bg-orange-500/20 text-orange-400 border border-orange-500/30' 
                          : 'bg-orange-100 text-orange-700 border border-orange-200'"
                      >
                        <IconArrowUp class="w-3 h-3" />
                        Reverse Share
                      </span>
                      
                      <!-- Status Badge -->
                      <span 
                        :class="[
                          'px-2 py-1 text-xs font-medium rounded-full inline-block w-fit flex-shrink-0',
                          upload.is_deleted 
                            ? isDark 
                              ? 'bg-red-900/30 text-red-400'
                              : 'bg-red-100 text-red-800'
                            : !upload.is_available
                              ? isDark
                                ? 'bg-gray-700 text-gray-400'
                                : 'bg-gray-100 text-gray-800'
                              : upload.expires_at && isExpiringSoon(upload.expires_at)
                                ? isDark
                                  ? 'bg-yellow-900/30 text-yellow-400'
                                  : 'bg-yellow-100 text-yellow-800'
                                : isDark
                                  ? 'bg-green-900/30 text-green-400'
                                  : 'bg-green-100 text-green-800'
                        ]"
                      >
                        {{ getStatusText(upload) }}
                      </span>
                    </div>
                    
                    <!-- Upload Metadata -->
                    <div class="flex flex-wrap items-center gap-2 sm:gap-4 lg:gap-6 text-xs sm:text-sm overflow-hidden" :class="isDark ? 'text-gray-400' : 'text-gray-500'">
                      <span class="whitespace-nowrap flex-shrink-0">{{ JSON.parse(upload.files).length }} files</span>
                      <span class="whitespace-nowrap flex-shrink-0">{{ getUploadSize(upload) }}</span>
                      <span class="whitespace-nowrap flex-shrink-0">Created {{ formatDate(upload.created_at) }}</span>
                      <span v-if="upload.expires_at" class="whitespace-nowrap flex-shrink-0">Expires {{ formatDate(upload.expires_at) }}</span>
                    </div>
                  </div>
                </div>

                <!-- Right side: Action Buttons -->
                <div class="flex items-center gap-2 flex-shrink-0 w-full sm:w-auto">
                  <button
                    @click="openFile(upload)"
                    class="bg-blue-600 hover:bg-blue-700 text-white px-3 sm:px-4 py-2 rounded-lg text-xs sm:text-sm font-medium transition-colors duration-200 flex items-center gap-2 flex-1 sm:flex-none justify-center"
                  >
                    <IconArrowUpRight class="w-3 h-3 sm:w-4 sm:h-4" />
                    <span class="hidden sm:inline">Open</span>
                  </button>
                  
                  <button
                    @click="$emit('copyToClipboard', `${getBaseUrl()}/download/${upload.upload_id}`)"
                    class="px-2 sm:px-3 py-2 rounded-lg text-xs sm:text-sm font-medium transition-colors duration-200 flex items-center justify-center"
                    :class="isDark ? 'bg-gray-700 hover:bg-gray-600 text-gray-300' : 'bg-gray-100 hover:bg-gray-200 text-gray-700'"
                    title="Copy Link"
                  >
                    <IconCopy class="w-3 h-3 sm:w-4 sm:h-4" />
                  </button>
                  
                  <button
                    @click="selectUpload(upload)"
                    class="px-3 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center justify-center"
                    :class="isDark ? 'bg-gray-700 hover:bg-gray-600 text-gray-300' : 'bg-gray-100 hover:bg-gray-200 text-gray-700'"
                    title="More Options"
                  >
                    <IconChevronDown 
                      class="w-4 h-4 transition-transform duration-200" 
                      :class="selectedUpload?.id === upload.id ? 'rotate-180' : ''"
                    />
                  </button>
                </div>
              </div>
            </div>

            <!-- Expanded Details (when selected) -->
            <div v-if="selectedUpload?.id === upload.id" class="border-t p-6"
                 :class="[
                   isDark ? 'border-gray-700 bg-gray-900/50' : 'border-gray-200 bg-gray-50'
                 ]">
              <div class="grid md:grid-cols-2 gap-6">
                <!-- Files List -->
                <div>
                  <h4 class="font-medium mb-3" :class="isDark ? 'text-white' : 'text-gray-900'">Files in this upload:</h4>
                  <div class="space-y-2 max-h-60 overflow-y-auto">
                    <div v-for="(file, fileIndex) in JSON.parse(upload.files)" :key="fileIndex" 
                         class="flex items-center gap-3 p-3 rounded-lg min-w-0 overflow-hidden"
                         :class="isDark ? 'bg-gray-800' : 'bg-white' "
                         :title="file">
                      <div class="flex items-center justify-center flex-shrink-0">
                        <img 
                          :src="getFileIconPath(file)" 
                          :alt="getFileIconAltText(file)" 
                          class="w-6 h-6 object-contain"
                          style="width: 24px !important; height: 24px !important;"
                        />
                      </div>
                      <div class="flex-1 min-w-0 overflow-hidden">
                        <span class="text-sm block truncate break-all" :class="isDark ? 'text-white' : 'text-gray-900'" style="word-break: break-all; overflow-wrap: break-word;">{{ file }}</span>
                      </div>
                    </div>
                  </div>
                </div>
                
                <!-- Additional Actions -->
                <div>
                  <h4 class="font-medium mb-3" :class="isDark ? 'text-white' : 'text-gray-900'">Actions:</h4>
                  <div class="space-y-3">
                    <button
                      @click="downloadSingleFile(upload)"
                      class="w-full bg-green-600 hover:bg-green-700 text-white px-4 py-3 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                    >
                      <IconDownload class="w-4 h-4" />
                      Download Files
                    </button>
                    
                    <!-- Change Expiration -->
                    <div v-if="!upload.is_deleted" class="space-y-2">
                      <div class="flex items-center justify-between">
                        <IconClock class="w-4 h-4" />
                        <button
                          @click="toggleExpirationEdit(upload.upload_id)"
                          class="text-xs px-2 py-1 rounded transition-colors duration-200"
                          :class="editingExpiration[upload.upload_id] 
                            ? 'bg-amber-100 text-amber-700 hover:bg-amber-200' 
                            : (isDark ? 'bg-gray-700 text-gray-300 hover:bg-gray-600' : 'bg-gray-100 text-gray-600 hover:bg-gray-200')"
                        >
                          {{ editingExpiration[upload.upload_id] ? 'Cancel' : 'Change' }}
                        </button>
                      </div>
                      
                      <div v-if="!editingExpiration[upload.upload_id]" 
                           class="text-sm p-2 rounded-lg"
                           :class="isDark ? 'bg-gray-800 text-gray-300' : 'bg-gray-100 text-gray-600'">
                        {{ upload.expires_at ? formatDate(upload.expires_at) : 'Never expires' }}
                      </div>
                      
                      <div v-else class="space-y-2">
                        <select
                          v-model="newExpirationValue[upload.upload_id]"
                          class="w-full px-3 py-2 text-sm rounded-lg border transition-colors duration-200"
                          :class="isDark 
                            ? 'bg-gray-800 border-gray-600 text-gray-200 focus:border-amber-500' 
                            : 'bg-white border-gray-300 text-gray-900 focus:border-amber-500'"
                        >
                          <option value="7days">7 Days</option>
                          <option value="1month">1 Month</option>
                          <option value="6months">6 Months</option>
                          <option value="1year">1 Year</option>
                          <option value="never">Never</option>
                        </select>
                        
                        <div class="flex gap-2">
                          <button
                            @click="changeExpiration(upload.upload_id)"
                            class="flex-1 px-3 py-2 bg-amber-600 hover:bg-amber-700 text-white text-sm font-medium rounded-lg transition-colors duration-200"
                          >
                            Save
                          </button>
                          <button
                            @click="toggleExpirationEdit(upload.upload_id)"
                            class="px-3 py-2 text-sm font-medium rounded-lg transition-colors duration-200"
                            :class="isDark ? 'bg-gray-700 text-gray-300 hover:bg-gray-600' : 'bg-gray-200 text-gray-700 hover:bg-gray-300'"
                          >
                            Cancel
                          </button>
                        </div>
                      </div>
                    </div>
                    
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
                      <IconEye v-if="upload.is_available" class="w-4 h-4" />
                      <IconEyeSlash v-else class="w-4 h-4" />
                      {{ upload.is_available ? 'Hide Upload' : 'Show Upload' }}
                    </button>
                    
                    <button
                      @click="$emit('deleteUpload', upload.upload_id)"
                      class="w-full bg-red-600 hover:bg-red-700 text-white px-4 py-3 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                    >
                      <IconTrash class="w-4 h-4" />
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
                : isDark 
                  ? 'bg-gray-800 text-gray-300 hover:bg-gray-700 border border-gray-700'
                  : 'bg-white text-gray-700 hover:bg-gray-50 border border-gray-200'
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
import { useTheme } from '../../../composables/useTheme'
import { useIcons } from '../../../composables/useIcons'
import { getApiUrl } from '../../../utils/apiUtils'
import IconCheckCircle from '~icons/solar/check-circle-bold'
import IconArrowUpRight from '~icons/tabler/folder-symlink'
import IconCopy from '~icons/solar/copy-bold'
import IconChevronDown from '~icons/solar/alt-arrow-down-line-duotone'
import IconDownload from '~icons/solar/download-bold'
import IconEye from '~icons/solar/eye-bold'
import IconEyeSlash from '~icons/solar/eye-closed-bold'
import IconTrash from '~icons/solar/trash-bin-minimalistic-bold'
import IconClock from '~icons/solar/clock-circle-bold'
import IconArrowUp from '~icons/solar/arrow-up-bold'

// Use composables
const { isDark } = useTheme()
const { 
  getUploadIcon, 
  getUploadIconAlt, 
  getFileIcon, 
  getFileIconAlt, 
  getSystemIcon 
} = useIcons()

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
const editingExpiration = ref<Record<string, boolean>>({})
const newExpirationValue = ref<Record<string, string>>({})

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
    value: 'reverse', 
    label: 'Reverse Share', 
    count: props.uploads.filter(u => u.is_reverse && !u.is_deleted).length 
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
  } else if (deletionFilter.value === 'reverse') {
    return props.uploads.filter(upload => upload.is_reverse && !upload.is_deleted)
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
  return getUploadIcon(files)
}

const getUploadIconAltText = (upload: any) => {
  const files = JSON.parse(upload.files)
  return getUploadIconAlt(files)
}

const getFileIconPath = (filename: string) => {
  return getFileIcon(filename)
}

const getFileIconAltText = (filename: string) => {
  return getFileIconAlt(filename)
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
    const response = await fetch(getApiUrl(`/download/${upload.upload_id}`), {
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

const getUploadSize = (upload: any) => {
  // Try to use the total_size field (this is what the backend actually stores)
  if (upload.total_size && upload.total_size > 0) {
    return formatFileSize(upload.total_size)
  }
  
  // Fallback to size field if total_size is not available
  if (upload.size && upload.size > 0) {
    return formatFileSize(upload.size)
  }
  
  // Try to use file_size if available  
  if (upload.file_size && upload.file_size > 0) {
    return formatFileSize(upload.file_size)
  }
  
  // Check if there's a files_size field
  if (upload.files_size && upload.files_size > 0) {
    return formatFileSize(upload.files_size)
  }
  
  // As a last resort, show that size is unavailable
  return 'Size unavailable'
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

// Expiration change functionality
const toggleExpirationEdit = (uploadId: string) => {
  editingExpiration.value[uploadId] = !editingExpiration.value[uploadId]
  if (editingExpiration.value[uploadId]) {
    // Initialize with current value or default
    newExpirationValue.value[uploadId] = '7days'
  }
}

const changeExpiration = async (uploadId: string) => {
  try {
    const response = await fetch(getApiUrl(`/uploads/${uploadId}/expiration`), {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      body: JSON.stringify({
        validity: newExpirationValue.value[uploadId]
      })
    })

    if (response.ok) {
      // Close edit mode
      editingExpiration.value[uploadId] = false
      // Emit an event to refresh uploads from parent
      // We could also update the local state here, but emitting refresh is cleaner
      window.location.reload() // Simple reload for now, could be improved
    } else {
      console.error('Failed to update expiration')
      alert('Failed to update expiration date')
    }
  } catch (error) {
    console.error('Error updating expiration:', error)
    alert('Failed to update expiration date')
  }
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
