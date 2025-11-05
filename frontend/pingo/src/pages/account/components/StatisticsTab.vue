<template>
  <div>
    <!-- Loading skeleton -->
    <div v-if="isLoading" class="space-y-6">
      <div class="animate-pulse rounded-xl p-6 border transition-colors duration-300"
           :style="{ 
             backgroundColor: isDark ? '#374151' : '#e5e7eb',
             borderColor: isDark ? '#4b5563' : '#d1d5db'
           }">
        <div class="h-6 rounded w-1/3 mb-4 transition-colors duration-300"
             :style="{ backgroundColor: isDark ? '#4b5563' : '#d1d5db' }"></div>
        <div class="space-y-3">
          <div class="h-4 rounded w-full transition-colors duration-300"
               :style="{ backgroundColor: isDark ? '#4b5563' : '#d1d5db' }"></div>
          <div class="h-4 rounded w-2/3 transition-colors duration-300"
               :style="{ backgroundColor: isDark ? '#4b5563' : '#d1d5db' }"></div>
        </div>
      </div>
    </div>
    
    <!-- Statistics Content -->
    <div v-else class="space-y-8">
      <div>
        <h3 class="text-xl font-bold mb-6 transition-colors duration-300"
            :style="{ color: isDark ? '#f9fafb' : '#111827' }">Platform Statistics</h3>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="rounded-xl p-6 border transition-colors duration-300"
               :style="{ 
                 backgroundColor: isDark ? '#1e3a8a' : '#eff6ff',
                 borderColor: isDark ? '#3b82f6' : '#3b82f6'
               }">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold transition-colors duration-300"
                    :style="{ color: isDark ? '#e0e7ff' : '#111827' }">Total Users</h3>
                <p class="text-3xl font-bold transition-colors duration-300"
                   :style="{ color: isDark ? '#93c5fd' : '#2563eb' }">{{ adminStats.totalUsers }}</p>
              </div>
              <IconUserGroup class="w-12 h-12 opacity-20 transition-colors duration-300"
                             :style="{ color: isDark ? '#60a5fa' : '#3b82f6' }" />
            </div>
          </div>
          
          <div class="rounded-xl p-6 border transition-colors duration-300"
               :style="{ 
                 backgroundColor: isDark ? '#059669' : '#f0fdf4',
                 borderColor: isDark ? '#10b981' : '#10b981'
               }">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold transition-colors duration-300"
                    :style="{ color: isDark ? '#d1fae5' : '#111827' }">Total Uploads</h3>
                <p class="text-3xl font-bold transition-colors duration-300"
                   :style="{ color: isDark ? '#86efac' : '#16a34a' }">{{ adminStats.totalUploads }}</p>
              </div>
              <IconCloudArrowUp class="w-12 h-12 opacity-20 transition-colors duration-300"
                                :style="{ color: isDark ? '#34d399' : '#10b981' }" />
            </div>
          </div>
          
          <div class="rounded-xl p-6 border transition-colors duration-300"
               :style="{ 
                 backgroundColor: isDark ? '#7c3aed' : '#fdf4ff',
                 borderColor: isDark ? '#8b5cf6' : '#8b5cf6'
               }">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-semibold transition-colors duration-300"
                    :style="{ color: isDark ? '#ede9fe' : '#111827' }">Storage Used</h3>
                <p class="text-3xl font-bold transition-colors duration-300"
                   :style="{ color: isDark ? '#c4b5fd' : '#9333ea' }">{{ formatBytes(adminStats.totalStorage) }}</p>
              </div>
              <IconArchiveBox class="w-12 h-12 opacity-20 transition-colors duration-300"
                              :style="{ color: isDark ? '#a78bfa' : '#8b5cf6' }" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTheme } from '../../../composables/useTheme'
import IconUserGroup from '~icons/solar/users-group-two-rounded-bold'
import IconCloudArrowUp from '~icons/solar/cloud-upload-bold'
import IconArchiveBox from '~icons/solar/archive-bold'

const { isDark } = useTheme()

// Props
interface Props {
  adminStats: {
    totalUsers: number
    totalUploads: number
    totalStorage: number
  }
  isLoading: boolean
}

defineProps<Props>()

// Helper functions
const formatBytes = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
