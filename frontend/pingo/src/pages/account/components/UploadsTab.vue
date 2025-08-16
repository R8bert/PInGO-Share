<template>
  <div class="space-y-6">
    <!-- Filter Controls -->
    <div v-if="!isLoading && uploads.length > 0">
      <div class="relative">
        <div class="absolute inset-0 bg-gradient-to-r from-blue-600/10 to-purple-600/10 rounded-2xl blur-xl"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-6"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.03)' : 'rgba(255,255,255,0.7)',
               borderColor: isDark ? 'rgba(255,255,255,0.08)' : 'rgba(0,0,0,0.08)'
             }">
          <div class="flex flex-wrap items-center justify-between gap-4">
            <div class="flex flex-wrap items-center gap-4">
              <!-- Filter by Status -->
              <div class="flex items-center gap-3">
                <label class="text-sm font-medium opacity-80"
                       :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
                  Filter:
                </label>
                <select 
                  v-model="deletionFilter"
                  @change="onFilterChange"
                  class="px-4 py-2 rounded-xl border backdrop-blur-sm text-sm font-medium transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-blue-500/50"
                  :style="{ 
                    backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                    borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                    color: isDark ? '#f9fafb' : '#111827'
                  }">
                  <option value="all">All Files</option>
                  <option value="active">Active Files</option>
                  <option value="deleted">Deleted Files</option>
                </select>
              </div>

              <!-- Items per page -->
              <div class="flex items-center gap-3">
                <label class="text-sm font-medium opacity-80"
                       :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
                  Show:
                </label>
                <select 
                  v-model="itemsPerPage"
                  @change="onItemsPerPageChange"
                  class="px-4 py-2 rounded-xl border backdrop-blur-sm text-sm font-medium transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-blue-500/50"
                  :style="{ 
                    backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                    borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
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
            <div class="px-4 py-2 rounded-xl backdrop-blur-sm border"
                 :style="{
                   backgroundColor: isDark ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.05)',
                   borderColor: 'rgba(59, 130, 246, 0.2)',
                   color: '#3b82f6'
                 }">
              <div class="flex items-center space-x-2">
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                </svg>
                <span class="text-sm font-medium">{{ displayedUploads.length }} of {{ filteredUploads.length }} uploads</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex justify-center py-20">
      <div class="relative">
        <div class="w-16 h-16 rounded-2xl bg-gradient-to-r from-blue-500 to-purple-500 flex items-center justify-center animate-pulse">
          <svg class="w-8 h-8 text-white animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
          </svg>
        </div>
      </div>
    </div>
    
    <!-- Empty State -->
    <div v-else-if="uploads.length === 0" class="text-center py-20">
      <div class="relative inline-block mb-6">
        <div class="w-20 h-20 rounded-2xl bg-gradient-to-r from-blue-500 to-purple-500 flex items-center justify-center">
          <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
          </svg>
        </div>
      </div>
      <h3 class="text-2xl font-bold mb-4"
          :style="{ color: isDark ? '#ffffff' : '#000000' }">
        No uploads yet
      </h3>
      <p class="text-lg mb-8 opacity-80"
         :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
        Start by uploading your first file to get started.
      </p>
      <router-link
        to="/"
        class="inline-flex items-center space-x-2 px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold rounded-2xl transition-all duration-200 transform hover:scale-105"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
        </svg>
        <span>Upload Files</span>
      </router-link>
    </div>

    <!-- No Results -->
    <div v-else-if="displayedUploads.length === 0" class="text-center py-20">
      <div class="relative inline-block mb-6">
        <div class="w-20 h-20 rounded-2xl bg-gradient-to-r from-gray-500 to-gray-600 flex items-center justify-center">
          <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </div>
      </div>
      <h3 class="text-2xl font-bold mb-4"
          :style="{ color: isDark ? '#ffffff' : '#000000' }">
        No uploads found
      </h3>
      <p class="text-lg opacity-80"
         :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
        Try changing your filter settings to see more results.
      </p>
    </div>

    <!-- Upload Cards -->
    <div v-else class="space-y-6">
      <div
        v-for="upload in displayedUploads"
        :key="upload.id"
        class="group relative">
        
        <!-- Card Background Glow -->
        <div class="absolute inset-0 bg-gradient-to-r opacity-0 group-hover:opacity-100 transition-opacity duration-300 rounded-2xl blur-xl"
             :class="upload.is_deleted ? 'from-red-600/20 to-orange-600/20' : 'from-green-600/20 to-blue-600/20'"></div>
        
        <!-- Main Card -->
        <div class="relative backdrop-blur-xl border rounded-2xl p-6 transition-all duration-300"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
               borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
               opacity: upload.is_deleted ? '0.7' : '1'
             }">
          
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <div class="flex items-center space-x-4 mb-4">
                <!-- File Icon -->
                <div class="w-12 h-12 rounded-xl bg-gradient-to-r from-blue-500 to-purple-500 flex items-center justify-center">
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
                  <p class="text-sm opacity-70"
                     :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                    {{ formatFileSize(upload.total_size) }} • {{ formatDate(upload.created_at) }}
                  </p>
                </div>
              </div>
              
              <!-- File List -->
              <div class="flex flex-wrap gap-2 mb-4">
                <span
                  v-for="filename in JSON.parse(upload.files)"
                  :key="filename"
                  class="inline-flex items-center px-3 py-1 rounded-lg text-xs font-medium backdrop-blur-sm border"
                  :style="{ 
                    backgroundColor: isDark ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.05)',
                    borderColor: 'rgba(59, 130, 246, 0.2)',
                    color: '#3b82f6'
                  }"
                >
                  {{ filename }}
                </span>
              </div>

              <!-- Status and Actions -->
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <!-- Status Indicator -->
                  <div v-if="upload.is_deleted" class="px-3 py-1 rounded-lg backdrop-blur-sm border"
                       :style="{
                         backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)',
                         borderColor: 'rgba(239, 68, 68, 0.3)',
                         color: '#ef4444'
                       }">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                      </svg>
                      <span class="text-sm font-medium">Deleted</span>
                    </div>
                  </div>
                  
                  <div v-else class="px-3 py-1 rounded-lg backdrop-blur-sm border"
                       :style="{
                         backgroundColor: upload.is_available 
                           ? (isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)')
                           : (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)'),
                         borderColor: upload.is_available ? 'rgba(34, 197, 94, 0.3)' : 'rgba(239, 68, 68, 0.3)',
                         color: upload.is_available ? '#22c55e' : '#ef4444'
                       }">
                    <div class="flex items-center space-x-2">
                      <div class="w-2 h-2 rounded-full"
                           :style="{ backgroundColor: upload.is_available ? '#22c55e' : '#ef4444' }"></div>
                      <span class="text-sm font-medium">{{ upload.is_available ? 'Available' : 'Unavailable' }}</span>
                    </div>
                  </div>

                  <!-- Expiration Info -->
                  <div v-if="!upload.is_deleted && upload.expires_at" class="px-3 py-1 rounded-lg backdrop-blur-sm border"
                       :style="{
                         backgroundColor: isExpiringSoon(upload.expires_at) 
                           ? (isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)')
                           : (isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)'),
                         borderColor: isExpiringSoon(upload.expires_at) ? 'rgba(239, 68, 68, 0.3)' : 'rgba(59, 130, 246, 0.3)',
                         color: isExpiringSoon(upload.expires_at) ? '#ef4444' : '#3b82f6'
                       }">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                      </svg>
                      <span class="text-sm font-medium">Expires {{ formatDate(upload.expires_at) }}</span>
                    </div>
                  </div>

                  <!-- Never Expires -->
                  <div v-else-if="!upload.is_deleted && !upload.expires_at" class="px-3 py-1 rounded-lg backdrop-blur-sm border"
                       :style="{
                         backgroundColor: isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)',
                         borderColor: 'rgba(34, 197, 94, 0.3)',
                         color: '#22c55e'
                       }">
                    <div class="flex items-center space-x-2">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                      </svg>
                      <span class="text-sm font-medium">Never expires</span>
                    </div>
                  </div>
                </div>

                <!-- Actions -->
                <div class="flex items-center space-x-2">
                  <!-- Copy Link -->
                  <button v-if="!upload.is_deleted"
                          @click="$emit('copyToClipboard', `${getBaseUrl()}/download/${upload.upload_id}`)"
                          class="p-2 rounded-lg backdrop-blur-sm border transition-all duration-200 hover:scale-105"
                          :style="{
                            backgroundColor: isDark ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.05)',
                            borderColor: 'rgba(59, 130, 246, 0.2)',
                            color: '#3b82f6'
                          }"
                          title="Copy share link">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                    </svg>
                  </button>

                  <!-- Toggle Availability -->
                  <button v-if="!upload.is_deleted"
                          @click="$emit('toggleAvailability', upload.upload_id, upload.is_available)"
                          class="p-2 rounded-lg backdrop-blur-sm border transition-all duration-200 hover:scale-105"
                          :style="{
                            backgroundColor: upload.is_available 
                              ? (isDark ? 'rgba(239, 68, 68, 0.1)' : 'rgba(239, 68, 68, 0.05)')
                              : (isDark ? 'rgba(34, 197, 94, 0.1)' : 'rgba(34, 197, 94, 0.05)'),
                            borderColor: upload.is_available ? 'rgba(239, 68, 68, 0.2)' : 'rgba(34, 197, 94, 0.2)',
                            color: upload.is_available ? '#ef4444' : '#22c55e'
                          }"
                          :title="upload.is_available ? 'Disable share' : 'Enable share'">
                    <svg v-if="upload.is_available" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                    </svg>
                  </button>

                  <!-- Delete Button -->
                  <button @click="$emit('deleteUpload', upload.upload_id)"
                          class="p-2 rounded-lg backdrop-blur-sm border transition-all duration-200 hover:scale-105"
                          :style="{
                            backgroundColor: isDark ? 'rgba(239, 68, 68, 0.1)' : 'rgba(239, 68, 68, 0.05)',
                            borderColor: 'rgba(239, 68, 68, 0.2)',
                            color: '#ef4444'
                          }"
                          title="Delete upload">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
    <div v-if="filteredUploads.length > itemsPerPage">
      <div class="relative">
        <div class="absolute inset-0 bg-gradient-to-r from-blue-600/10 to-purple-600/10 rounded-2xl blur-xl"></div>
        <div class="relative backdrop-blur-xl border rounded-2xl p-4"
             :style="{
               backgroundColor: isDark ? 'rgba(255,255,255,0.03)' : 'rgba(255,255,255,0.7)',
               borderColor: isDark ? 'rgba(255,255,255,0.08)' : 'rgba(0,0,0,0.08)'
             }">
          <div class="flex items-center justify-between">
            <p class="text-sm opacity-70"
               :style="{ color: isDark ? '#f3f4f6' : '#374151' }">
              Showing {{ currentPage * itemsPerPage + 1 }} to {{ Math.min((currentPage + 1) * itemsPerPage, filteredUploads.length) }} of {{ filteredUploads.length }} results
            </p>
            
            <div class="flex items-center space-x-2">
              <button @click="currentPage = 0"
                      :disabled="currentPage === 0"
                      class="px-4 py-2 rounded-lg backdrop-blur-sm border transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                      :style="{
                        backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                        borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                        color: isDark ? '#f9fafb' : '#111827'
                      }">
                First
              </button>
              
              <button @click="currentPage--"
                      :disabled="currentPage === 0"
                      class="px-4 py-2 rounded-lg backdrop-blur-sm border transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                      :style="{
                        backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                        borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                        color: isDark ? '#f9fafb' : '#111827'
                      }">
                Previous
              </button>
              
              <span class="px-4 py-2 rounded-lg backdrop-blur-sm border"
                    :style="{
                      backgroundColor: isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.1)',
                      borderColor: 'rgba(59, 130, 246, 0.3)',
                      color: '#3b82f6'
                    }">
                {{ currentPage + 1 }} of {{ totalPages }}
              </span>
              
              <button @click="currentPage++"
                      :disabled="currentPage >= totalPages - 1"
                      class="px-4 py-2 rounded-lg backdrop-blur-sm border transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                      :style="{
                        backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                        borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                        color: isDark ? '#f9fafb' : '#111827'
                      }">
                Next
              </button>
              
              <button @click="currentPage = totalPages - 1"
                      :disabled="currentPage >= totalPages - 1"
                      class="px-4 py-2 rounded-lg backdrop-blur-sm border transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                      :style="{
                        backgroundColor: isDark ? 'rgba(255,255,255,0.05)' : 'rgba(255,255,255,0.8)',
                        borderColor: isDark ? 'rgba(255,255,255,0.1)' : 'rgba(0,0,0,0.1)',
                        color: isDark ? '#f9fafb' : '#111827'
                      }">
                Last
              </button>
            </div>
          </div>
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
/* Upload container animations */
.group {
  transform-origin: center;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}
</style>
