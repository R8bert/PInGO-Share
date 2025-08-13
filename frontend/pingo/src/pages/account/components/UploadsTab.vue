<template>
  <div>
    <!-- Filter Controls -->
    <div v-if="!isLoading && uploads.length > 0" class="mb-4 flex flex-wrap items-center justify-between gap-3 p-4 rounded-lg border transition-colors duration-300"
         :style="{ 
           backgroundColor: isDark ? '#111827' : '#ffffff',
           borderColor: isDark ? '#374151' : '#e5e7eb'
         }">
      <div class="flex flex-wrap items-center gap-3">
        <!-- Filter by Status -->
        <div class="flex items-center gap-2">
          <label class="text-sm font-medium transition-colors duration-300"
                 :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
            Filter:
          </label>
          <select 
            v-model="deletionFilter"
            @change="onFilterChange"
            class="px-3 py-1.5 rounded-md border text-sm transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
            :style="{ 
              backgroundColor: isDark ? '#1f2937' : '#f9fafb',
              borderColor: isDark ? '#4b5563' : '#d1d5db',
              color: isDark ? '#f9fafb' : '#111827'
            }">
            <option value="all">All</option>
            <option value="active">Active</option>
            <option value="deleted">Deleted</option>
          </select>
        </div>

        <!-- Items per page -->
        <div class="flex items-center gap-2">
          <label class="text-sm font-medium transition-colors duration-300"
                 :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
            Show:
          </label>
          <select 
            v-model="itemsPerPage"
            @change="onItemsPerPageChange"
            class="px-3 py-1.5 rounded-md border text-sm transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
            :style="{ 
              backgroundColor: isDark ? '#1f2937' : '#f9fafb',
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
      <div class="text-sm font-medium transition-colors duration-300"
           :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
        {{ displayedUploads.length }} of {{ filteredUploads.length }} uploads
      </div>
    </div>

    <div v-if="isLoading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
    </div>
    
    <div v-else-if="uploads.length === 0" class="text-center py-12">
      <CloudArrowUpIcon class="w-16 h-16 mx-auto mb-4 transition-colors duration-300"
                        :style="{ color: isDark ? '#4b5563' : '#d1d5db' }" />
      <h3 class="text-lg font-medium mb-2 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">No uploads yet</h3>
      <p class="mb-6 transition-colors duration-300"
         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Start by uploading your first file.</p>
      <router-link
        to="/"
        class="inline-flex items-center px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
      >
        Upload Files
      </router-link>
    </div>

    <div v-else-if="displayedUploads.length === 0" class="text-center py-12">
      <svg class="w-16 h-16 mx-auto mb-4 transition-colors duration-300" 
           :style="{ color: isDark ? '#4b5563' : '#d1d5db' }"
           fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
      </svg>
      <h3 class="text-lg font-medium mb-2 transition-colors duration-300"
          :style="{ color: isDark ? '#f9fafb' : '#111827' }">No uploads found</h3>
      <p class="mb-6 transition-colors duration-300"
         :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
        Try changing your filter settings.
      </p>
    </div>

    <div v-else class="space-y-4">
      <div
        v-for="upload in displayedUploads"
        :key="upload.id"
        class="rounded-xl p-6 border hover:shadow-lg transition-all duration-300 upload-container"
        :class="[
          upload.is_deleted ? 'opacity-90' : '',
          newlyDeletedUploads.has(upload.upload_id) ? 'animate-delete-transition' : ''
        ]"
        :style="{ 
          backgroundColor: upload.is_deleted 
            ? (isDark ? '#991b1b' : '#fee2e2') 
            : (isDark ? '#111827' : '#ffffff'),
          borderColor: upload.is_deleted 
            ? (isDark ? '#dc2626' : '#f87171') 
            : (isDark ? '#374151' : '#e5e7eb'),
          boxShadow: upload.is_deleted 
            ? (isDark ? '0 0 0 1px rgba(220, 38, 38, 0.1)' : '0 0 0 1px rgba(239, 68, 68, 0.1)')
            : 'none'
        }"
      >
        <div class="flex items-start justify-between">
          <div class="flex-1">
            <div class="flex items-center space-x-3 mb-3">
              <img src="../../../images/icons/com.belmoussaoui.ReadItLater.svg" 
                   alt="Files uploaded" 
                   class="w-12 h-12 transition-all duration-300"
                   :class="isDark ? '' : 'drop-shadow-sm'"
                   :style="{ opacity: isDark ? '0.9' : '1' }" />
              <div>
                <h3 class="font-semibold text-base transition-colors duration-300"
                    :style="{ color: isDark ? '#f9fafb' : '#111827' }">
                  {{ JSON.parse(upload.files).length }} file{{ JSON.parse(upload.files).length > 1 ? 's' : '' }}
                </h3>
                <p class="text-sm font-medium transition-colors duration-300"
                   :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                  {{ formatFileSize(upload.total_size) }} • {{ formatDate(upload.created_at) }}
                </p>
              </div>
            </div>
            
            <div class="flex flex-wrap gap-2 mb-4">
              <span
                v-for="filename in JSON.parse(upload.files)"
                :key="filename"
                class="inline-flex items-center px-3 py-1 rounded-full text-xs font-semibold transition-colors duration-300"
                :style="{ 
                  backgroundColor: isDark ? '#1e40af' : '#dbeafe',
                  color: isDark ? '#bfdbfe' : '#1e40af'
                }"
              >
                {{ filename }}
              </span>
            </div>

            <div class="flex items-center space-x-4 text-sm">
              <!-- Only show expiration for non-deleted uploads -->
              <template v-if="!upload.is_deleted">
                <span v-if="upload.expires_at" :class="[
                  'flex items-center font-medium',
                  isExpiringSoon(upload.expires_at) ? 'text-red-500' : ''
                ]" :style="!isExpiringSoon(upload.expires_at) ? { color: isDark ? '#d1d5db' : '#4b5563' } : {}">
                  <ClockIcon class="w-4 h-4 mr-1" />
                  {{ upload.expires_at ? `Expires ${formatDate(upload.expires_at)}` : 'Never expires' }}
                </span>
                <span v-else class="flex items-center text-emerald-500 font-medium">
                  <InfinityIcon class="w-4 h-4 mr-1" />
                  Never expires
                </span>
              </template>
              <span v-if="upload.is_reverse" class="flex items-center text-purple-500 font-medium">
                <ShareIcon class="w-4 h-4 mr-1" />
                Reverse Upload
              </span>
              <span v-if="!upload.is_deleted" :class="[
                'flex items-center font-medium',
                upload.is_available ? 'text-emerald-500' : 'text-red-500'
              ]">
                <div class="w-2 h-2 rounded-full mr-1" :class="[
                  upload.is_available ? 'bg-emerald-500' : 'bg-red-500'
                ]"></div>
                {{ upload.is_available ? 'Available' : 'Unavailable' }}
              </span>
            </div>
          </div>

          <div class="flex items-center space-x-2">
            <!-- Deleted status badge (for deleted uploads) -->
            <div v-if="upload.is_deleted" class="flex flex-col font-bold text-xs px-3 py-2 rounded-md"
                 :style="{ 
                   backgroundColor: isDark ? '#7f1d1d' : '#fecaca',
                   color: isDark ? '#fecaca' : '#7f1d1d'
                 }">
              <div class="flex items-center">
                <svg class="w-4 h-4 mr-1 font-bold" fill="currentColor" stroke="currentColor" viewBox="0 0 24 24" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
                DELETED
              </div>
              <div class="text-xs font-medium mt-1 opacity-90">
                {{ formatDate(upload.deleted_at) }}
              </div>
              <div v-if="upload.deletion_reason && getDeletionReason(upload.deletion_reason)" class="text-xs font-medium opacity-80">
                {{ getDeletionReason(upload.deletion_reason) }}
              </div>
            </div>
            
            <!-- Action buttons (disabled for deleted uploads) -->
            <template v-if="!upload.is_deleted">
              <button
                @click="copyToClipboard(`${getBaseUrl()}${upload.download_url}`)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="isDark ? 'hover:bg-blue-600/20 bg-blue-600/10' : 'hover:bg-blue-100 bg-blue-50'"
                title="Copy link"
              >
                <ClipboardDocumentIcon class="w-5 h-5 transition-colors duration-300"
                                     :class="isDark ? 'text-blue-400 group-hover:text-blue-300' : 'text-blue-600 group-hover:text-blue-700'" />
              </button>
              <button
                @click="goToFiles(upload.download_url)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="isDark ? 'hover:bg-emerald-600/20 bg-emerald-600/10' : 'hover:bg-emerald-100 bg-emerald-50'"
                title="Go to files"
              >
                <FolderIcon class="w-5 h-5 transition-colors duration-300"
                           :class="isDark ? 'text-emerald-400 group-hover:text-emerald-300' : 'text-emerald-600 group-hover:text-emerald-700'" />
              </button>
              <button
                @click="toggleAvailability(upload.upload_id, upload.is_available)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="upload.is_available 
                  ? (isDark ? 'hover:bg-amber-600/20 bg-amber-600/10' : 'hover:bg-amber-100 bg-amber-50')
                  : (isDark ? 'hover:bg-green-600/20 bg-green-600/10' : 'hover:bg-green-100 bg-green-50')"
                :title="upload.is_available ? 'Make unavailable' : 'Make available'"
              >
                <EyeSlashIcon v-if="upload.is_available" class="w-5 h-5 transition-colors duration-300"
                             :class="isDark ? 'text-amber-400 group-hover:text-amber-300' : 'text-amber-600 group-hover:text-amber-700'" />
                <EyeIcon v-else class="w-5 h-5 transition-colors duration-300"
                        :class="isDark ? 'text-green-400 group-hover:text-green-300' : 'text-green-600 group-hover:text-green-700'" />
              </button>
            
              <!-- Expiration Dropdown -->
              <div class="relative">
                <button
                  @click="toggleExpirationDropdown(upload.upload_id)"
                  class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                  :class="isDark ? 'hover:bg-purple-600/20 bg-purple-600/10' : 'hover:bg-purple-100 bg-purple-50'"
                  title="Change expiration"
                >
                  <ClockIcon class="w-5 h-5 transition-colors duration-300"
                            :class="isDark ? 'text-purple-400 group-hover:text-purple-300' : 'text-purple-600 group-hover:text-purple-700'" />
                </button>
                <div v-if="showExpirationDropdown[upload.upload_id]" 
                     class="absolute right-0 mt-2 w-48 rounded-lg shadow-lg border z-10 transition-colors duration-300"
                     :style="{ 
                       backgroundColor: isDark ? '#111827' : '#ffffff',
                       borderColor: isDark ? '#374151' : '#e5e7eb'
                     }">
                  <div class="py-1">
                    <button
                      @click="changeExpiration(upload.upload_id, '7days')"
                      class="block w-full text-left px-4 py-2 text-sm font-medium transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }"
                    >
                      7 days
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, '1month')"
                      class="block w-full text-left px-4 py-2 text-sm font-medium transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }"
                    >
                      1 month
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, '6months')"
                      class="block w-full text-left px-4 py-2 text-sm font-medium transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }"
                    >
                      6 months
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, '1year')"
                      class="block w-full text-left px-4 py-2 text-sm font-medium transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }"
                    >
                      1 year
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, 'never')"
                      class="block w-full text-left px-4 py-2 text-sm font-medium transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#f9fafb' : '#111827' }"
                    >
                      Never expires
                    </button>
                  </div>
                </div>
              </div>
              
              <!-- Delete Button -->
              <button
                @click="deleteUpload(upload.upload_id)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group border"
                :class="isDark ? 'hover:bg-red-500/30 bg-red-500/20 border-red-400/50' : 'hover:bg-red-100 bg-red-50 border-red-200'"
                title="Delete upload"
              >
                <TrashIcon class="w-5 h-5 transition-colors duration-300 font-bold"
                          :class="isDark ? 'text-red-300 group-hover:text-red-200' : 'text-red-700 group-hover:text-red-800'" />
              </button>
            </template>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination (if there are more items than displayed) -->
        <!-- Load More/Show Less Controls -->
    <div v-if="filteredUploads.length > itemsPerPage" 
         class="flex justify-center items-center space-x-4 p-6 rounded-lg border transition-colors duration-300"
         :style="{ 
           backgroundColor: isDark ? '#111827' : '#ffffff',
           borderColor: isDark ? '#374151' : '#e5e7eb'
         }">
      <button
        v-if="currentlyShowing < filteredUploads.length"
        @click="loadMore"
        class="px-6 py-3 rounded-lg font-semibold transition-all duration-300 bg-blue-600 text-white hover:bg-blue-700 hover:shadow-lg transform hover:scale-105"
      >
        Load More ({{ Math.min(itemsPerPage, filteredUploads.length - currentlyShowing) }} more)
      </button>
      
      <span class="font-bold text-lg transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">
        Showing {{ currentlyShowing }} of {{ filteredUploads.length }} uploads
      </span>
      
      <button
        v-if="currentlyShowing > itemsPerPage"
        @click="currentlyShowing = itemsPerPage"
        class="px-6 py-3 rounded-lg font-semibold transition-all duration-300 bg-gray-600 text-white hover:bg-gray-700 hover:shadow-lg transform hover:scale-105"
      >
        Show Less
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from '../../../composables/useTheme'
import {
  CloudArrowUpIcon,
  ClockIcon,
  ShareIcon,
  ClipboardDocumentIcon,
  FolderIcon,
  EyeIcon,
  EyeSlashIcon,
  TrashIcon
} from '@heroicons/vue/24/outline'

const { isDark } = useTheme()

// Props
interface Props {
  uploads: any[]
  isLoading: boolean
  showExpirationDropdown: Record<string, boolean>
}

const props = defineProps<Props>()

// Emits
const emit = defineEmits<{
  copyToClipboard: [text: string]
  toggleAvailability: [uploadId: string, currentAvailability: boolean]
  changeExpiration: [uploadId: string, validity: string]
  toggleExpirationDropdown: [uploadId: string]
  deleteUpload: [uploadId: string]
}>()

// Router
const router = useRouter()

// Cookie utility functions
const setCookie = (name: string, value: string, days: number = 30) => {
  const expires = new Date()
  expires.setTime(expires.getTime() + (days * 24 * 60 * 60 * 1000))
  document.cookie = `${name}=${value};expires=${expires.toUTCString()};path=/`
}

const getCookie = (name: string): string | null => {
  const nameEQ = name + "="
  const ca = document.cookie.split(';')
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i]
    while (c.charAt(0) === ' ') c = c.substring(1, c.length)
    if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length)
  }
  return null
}

// Filter and pagination state with cookie persistence
const deletionFilter = ref<'all' | 'active' | 'deleted'>(
  (() => {
    const saved = getCookie('uploads_deletion_filter') as 'all' | 'active' | 'deleted'
    return ['all', 'active', 'deleted'].includes(saved) ? saved : 'all'
  })()
)
const itemsPerPage = ref<number>(
  (() => {
    const saved = parseInt(getCookie('uploads_items_per_page') || '25')
    return [10, 25, 50, 100].includes(saved) ? saved : 25
  })()
)
const currentlyShowing = ref<number>(itemsPerPage.value)

// Animation state for newly deleted uploads
const newlyDeletedUploads = ref<Set<string>>(new Set())

// Function to handle animation after successful delete
const handleSuccessfulDelete = (uploadId: string) => {
  newlyDeletedUploads.value.add(uploadId)
  // Remove the animation class after animation completes
  setTimeout(() => {
    newlyDeletedUploads.value.delete(uploadId)
  }, 1000) // Duration matches CSS animation
}

// Watch for changes in uploads to detect newly deleted items
watch(() => props.uploads, (newUploads, oldUploads) => {
  if (oldUploads) {
    // Find uploads that changed from not deleted to deleted
    newUploads.forEach(newUpload => {
      const oldUpload = oldUploads.find(old => old.upload_id === newUpload.upload_id)
      if (oldUpload && !oldUpload.is_deleted && newUpload.is_deleted) {
        handleSuccessfulDelete(newUpload.upload_id)
      }
    })
  }
}, { deep: true })

// For infinity icon
const InfinityIcon = { template: '<span class="text-lg">∞</span>' }

// Computed properties for filtering and pagination
const filteredUploads = computed(() => {
  if (deletionFilter.value === 'all') {
    return props.uploads
  } else if (deletionFilter.value === 'deleted') {
    return props.uploads.filter(upload => upload.is_deleted)
  } else {
    return props.uploads.filter(upload => !upload.is_deleted)
  }
})

const displayedUploads = computed(() => {
  return filteredUploads.value.slice(0, Math.min(currentlyShowing.value, filteredUploads.value.length))
})

// Filter change handlers
const onFilterChange = () => {
  // Reset the currently showing count when filter changes
  currentlyShowing.value = itemsPerPage.value
  // Save to cookie
  setCookie('uploads_deletion_filter', deletionFilter.value)
}

const onItemsPerPageChange = () => {
  // Update the currently showing count when items per page changes
  currentlyShowing.value = itemsPerPage.value
  // Save to cookie
  setCookie('uploads_items_per_page', itemsPerPage.value.toString())
}

// Watch for filter changes to save to cookies
watch(deletionFilter, (newValue) => {
  setCookie('uploads_deletion_filter', newValue)
})

watch(itemsPerPage, (newValue) => {
  setCookie('uploads_items_per_page', newValue.toString())
})

const loadMore = () => {
  // Increase the number of items being shown
  const newCount = currentlyShowing.value + itemsPerPage.value
  currentlyShowing.value = Math.min(newCount, filteredUploads.value.length)
}

// Helper functions
const formatDate = (dateString: string | undefined) => {
  if (!dateString) return 'Unknown'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const isExpiringSoon = (expiresAt: string) => {
  if (!expiresAt) return false
  const expiry = new Date(expiresAt)
  const now = new Date()
  const timeDiff = expiry.getTime() - now.getTime()
  const daysDiff = timeDiff / (1000 * 3600 * 24)
  return daysDiff <= 1 && daysDiff > 0
}

const getBaseUrl = () => {
  return window.location.origin
}

const getDeletionReason = (deletionReason: any) => {
  // Handle SQL nullable string format: { "String": "value", "Valid": true }
  if (typeof deletionReason === 'string') {
    return deletionReason
  }
  if (typeof deletionReason === 'object' && deletionReason?.Valid && deletionReason?.String) {
    return deletionReason.String
  }
  return null
}

const goToFiles = (downloadUrl: string) => {
  router.push(downloadUrl)
}

// Event handlers
const copyToClipboard = (text: string) => {
  emit('copyToClipboard', text)
}

const toggleAvailability = (uploadId: string, currentAvailability: boolean) => {
  emit('toggleAvailability', uploadId, currentAvailability)
}

const changeExpiration = (uploadId: string, validity: string) => {
  emit('changeExpiration', uploadId, validity)
}

const toggleExpirationDropdown = (uploadId: string) => {
  emit('toggleExpirationDropdown', uploadId)
}

const deleteUpload = (uploadId: string) => {
  emit('deleteUpload', uploadId)
}

// Initialize component when mounted
onMounted(() => {
  // Ensure currentlyShowing matches itemsPerPage on load
  currentlyShowing.value = itemsPerPage.value
})
</script>

<style scoped>
/* Upload container animations */
.upload-container {
  transform-origin: center;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Delete transition animation - transforms to deleted state */
.animate-delete-transition {
  animation: deleteTransition 1s cubic-bezier(0.4, 0, 0.2, 1) forwards;
}

@keyframes deleteTransition {
  0% {
    transform: scale(1);
    filter: brightness(1) saturate(1);
  }
  20% {
    transform: scale(1.02);
    filter: brightness(1.1) saturate(1.2);
  }
  40% {
    transform: scale(0.98);
    filter: brightness(0.9) saturate(1.5);
  }
  60% {
    transform: scale(1.01);
    filter: brightness(0.8) saturate(2);
  }
  80% {
    transform: scale(0.99);
    filter: brightness(0.7) saturate(1.8) hue-rotate(10deg);
  }
  100% {
    transform: scale(1);
    filter: brightness(0.75) saturate(1.5) hue-rotate(15deg);
  }
}

/* Deleted state styling enhancement */
.upload-container.animate-delete-transition {
  background-color: var(--deleted-bg) !important;
  border-color: var(--deleted-border) !important;
}
</style>
