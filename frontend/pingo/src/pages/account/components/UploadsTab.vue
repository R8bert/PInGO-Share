<template>
  <div>
    <!-- Filter Controls -->
    <div v-if="!isLoading && uploads.length > 0" class="mb-6 p-4 rounded-xl border transition-colors duration-300"
         :style="{ 
           backgroundColor: isDark ? '#1f2937' : '#f8fafc',
           borderColor: isDark ? '#374151' : '#e2e8f0'
         }">
      <div class="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
        <!-- Filter by Status -->
        <div class="flex items-center space-x-3">
          <label class="text-sm font-medium transition-colors duration-300"
                 :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
            Filter:
          </label>
          <select 
            v-model="deletionFilter"
            @change="onFilterChange"
            class="px-3 py-2 rounded-lg border text-sm transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
            :style="{ 
              backgroundColor: isDark ? '#374151' : '#ffffff',
              borderColor: isDark ? '#4b5563' : '#d1d5db',
              color: isDark ? '#f3f4f6' : '#374151'
            }">
            <option value="all">Show All</option>
            <option value="active">Not Deleted Only</option>
            <option value="deleted">Deleted Only</option>
          </select>
        </div>

        <!-- Items per page -->
        <div class="flex items-center space-x-3">
          <label class="text-sm font-medium transition-colors duration-300"
                 :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
            Items per page:
          </label>
          <select 
            v-model="itemsPerPage"
            @change="onItemsPerPageChange"
            class="px-3 py-2 rounded-lg border text-sm transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
            :style="{ 
              backgroundColor: isDark ? '#374151' : '#ffffff',
              borderColor: isDark ? '#4b5563' : '#d1d5db',
              color: isDark ? '#f3f4f6' : '#374151'
            }">
            <option value="10">10</option>
            <option value="25">25</option>
            <option value="50">50</option>
            <option value="100">100</option>
          </select>
        </div>
      </div>

      <!-- Results Info -->
      <div class="mt-3 text-sm transition-colors duration-300"
           :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
        Showing {{ displayedUploads.length }} of {{ filteredUploads.length }} uploads
        <span v-if="deletionFilter !== 'all'">
          ({{ deletionFilter === 'deleted' ? 'deleted' : 'active' }})
        </span>
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
        class="rounded-xl p-6 border hover:shadow-md transition-all duration-300 upload-container"
        :class="[
          upload.is_deleted ? 'opacity-75' : '',
          newlyDeletedUploads.has(upload.upload_id) ? 'animate-delete-transition' : ''
        ]"
        :style="{ 
          backgroundColor: upload.is_deleted 
            ? (isDark ? '#7f1d1d' : '#fecaca') 
            : (isDark ? '#1f2937' : '#f9fafb'),
          borderColor: upload.is_deleted 
            ? (isDark ? '#dc2626' : '#ef4444') 
            : (isDark ? '#374151' : '#e5e7eb')
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
                <h3 class="font-medium transition-colors duration-300"
                    :style="{ color: isDark ? '#f9fafb' : '#111827' }">
                  {{ JSON.parse(upload.files).length }} file{{ JSON.parse(upload.files).length > 1 ? 's' : '' }}
                </h3>
                <p class="text-sm transition-colors duration-300"
                   :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
                  {{ formatFileSize(upload.total_size) }} • {{ formatDate(upload.created_at) }}
                </p>
              </div>
            </div>
            
            <div class="flex flex-wrap gap-2 mb-3">
              <span
                v-for="filename in JSON.parse(upload.files)"
                :key="filename"
                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium transition-colors duration-300"
                :style="{ 
                  backgroundColor: isDark ? '#1e3a8a' : '#dbeafe',
                  color: isDark ? '#93c5fd' : '#1e40af'
                }"
              >
                {{ filename }}
              </span>
            </div>

            <div class="flex items-center space-x-4 text-sm">
              <span v-if="upload.expires_at" :class="[
                'flex items-center',
                isExpiringSoon(upload.expires_at) ? 'text-red-600' : ''
              ]" :style="!isExpiringSoon(upload.expires_at) ? { color: isDark ? '#9ca3af' : '#6b7280' } : {}">
                <ClockIcon class="w-4 h-4 mr-1" />
                {{ upload.expires_at ? `Expires ${formatDate(upload.expires_at)}` : 'Never expires' }}
              </span>
              <span v-else class="flex items-center text-green-600">
                <InfinityIcon class="w-4 h-4 mr-1" />
                Never expires
              </span>
              <span v-if="upload.is_reverse" class="flex items-center text-purple-600">
                <ShareIcon class="w-4 h-4 mr-1" />
                Reverse Upload
              </span>
              <span v-if="upload.is_deleted" class="flex items-center text-red-600 font-medium">
                <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                </svg>
                Deleted
              </span>
              <span v-else :class="[
                'flex items-center',
                upload.is_available ? 'text-green-600' : 'text-red-600'
              ]">
                <div class="w-2 h-2 rounded-full mr-1" :class="[
                  upload.is_available ? 'bg-green-500' : 'bg-red-500'
                ]"></div>
                {{ upload.is_available ? 'Available' : 'Unavailable' }}
              </span>
            </div>
          </div>

          <div class="flex items-center space-x-2">
            <!-- Action buttons (disabled for deleted uploads) -->
            <template v-if="!upload.is_deleted">
              <button
                @click="copyToClipboard(`${getBaseUrl()}${upload.download_url}`)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="isDark ? 'hover:bg-blue-500/20 bg-blue-500/10' : 'hover:bg-blue-50 bg-blue-25'"
                title="Copy link"
              >
                <ClipboardDocumentIcon class="w-5 h-5 transition-colors duration-300"
                                     :class="isDark ? 'text-blue-400 group-hover:text-blue-300' : 'text-blue-600 group-hover:text-blue-700'" />
              </button>
              <button
                @click="goToFiles(upload.download_url)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="isDark ? 'hover:bg-emerald-500/20 bg-emerald-500/10' : 'hover:bg-emerald-50 bg-emerald-25'"
                title="Go to files"
              >
                <FolderIcon class="w-5 h-5 transition-colors duration-300"
                           :class="isDark ? 'text-emerald-400 group-hover:text-emerald-300' : 'text-emerald-600 group-hover:text-emerald-700'" />
              </button>
              <button
                @click="toggleAvailability(upload.upload_id, upload.is_available)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="upload.is_available 
                  ? (isDark ? 'hover:bg-amber-500/20 bg-amber-500/10' : 'hover:bg-amber-50 bg-amber-25')
                  : (isDark ? 'hover:bg-green-500/20 bg-green-500/10' : 'hover:bg-green-50 bg-green-25')"
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
                  :class="isDark ? 'hover:bg-purple-500/20 bg-purple-500/10' : 'hover:bg-purple-50 bg-purple-25'"
                  title="Change expiration"
                >
                  <ClockIcon class="w-5 h-5 transition-colors duration-300"
                            :class="isDark ? 'text-purple-400 group-hover:text-purple-300' : 'text-purple-600 group-hover:text-purple-700'" />
                </button>
                <div v-if="showExpirationDropdown[upload.upload_id]" 
                     class="absolute right-0 mt-2 w-48 rounded-md shadow-lg border z-10 transition-colors duration-300"
                     :style="{ 
                       backgroundColor: isDark ? '#1f2937' : '#ffffff',
                       borderColor: isDark ? '#374151' : '#e5e7eb'
                     }">
                  <div class="py-1">
                    <button
                      @click="changeExpiration(upload.upload_id, '7days')"
                      class="block w-full text-left px-4 py-2 text-sm transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#e5e7eb' : '#374151' }"
                    >
                      7 days
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, '1month')"
                      class="block w-full text-left px-4 py-2 text-sm transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#e5e7eb' : '#374151' }"
                    >
                      1 month
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, '6months')"
                      class="block w-full text-left px-4 py-2 text-sm transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#e5e7eb' : '#374151' }"
                    >
                      6 months
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, '1year')"
                      class="block w-full text-left px-4 py-2 text-sm transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#e5e7eb' : '#374151' }"
                    >
                      1 year
                    </button>
                    <button
                      @click="changeExpiration(upload.upload_id, 'never')"
                      class="block w-full text-left px-4 py-2 text-sm transition-colors duration-300 hover:bg-gray-100 dark:hover:bg-gray-800"
                      :style="{ color: isDark ? '#e5e7eb' : '#374151' }"
                    >
                      Never expires
                    </button>
                  </div>
                </div>
              </div>
              
              <!-- Delete Button -->
              <button
                @click="deleteUpload(upload.upload_id)"
                class="p-2 rounded-lg transition-all duration-300 hover:scale-105 hover:shadow-md group"
                :class="isDark ? 'hover:bg-red-500/20 bg-red-500/10' : 'hover:bg-red-50 bg-red-25'"
                title="Delete upload"
              >
                <TrashIcon class="w-5 h-5 transition-colors duration-300"
                          :class="isDark ? 'text-red-400 group-hover:text-red-300' : 'text-red-600 group-hover:text-red-700'" />
              </button>
            </template>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination (if there are more items than displayed) -->
    <div v-if="filteredUploads.length > itemsPerPage" class="mt-6 text-center">
      <div class="inline-flex items-center space-x-2 p-2 rounded-lg border transition-colors duration-300"
           :style="{ 
             backgroundColor: isDark ? '#1f2937' : '#f8fafc',
             borderColor: isDark ? '#374151' : '#e2e8f0'
           }">
        <span class="text-sm transition-colors duration-300"
              :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">
          Showing {{ displayedUploads.length }} of {{ filteredUploads.length }} items
        </span>
        <button 
          v-if="currentlyShowing < filteredUploads.length"
          @click="loadMore"
          class="px-3 py-1 text-sm bg-blue-600 hover:bg-blue-700 text-white rounded-md transition-colors duration-200"
        >
          Load More ({{ Math.min(itemsPerPage, filteredUploads.length - currentlyShowing) }})
        </button>
      </div>
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
