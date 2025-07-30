<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900">Account</h1>
            <p class="text-gray-600 mt-1">Manage your profile and uploads</p>
          </div>
          <router-link
            to="/"
            class="inline-flex items-center px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
          >
            <ArrowLeftIcon class="w-5 h-5 mr-2" />
            Back to Home
          </router-link>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- User Profile -->
        <div class="lg:col-span-1">
          <div class="bg-white rounded-xl shadow-sm p-6">
            <div class="text-center">
              <div class="w-20 h-20 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
                <UserIcon class="w-10 h-10 text-blue-600" />
              </div>
              <h2 class="text-xl font-bold text-gray-900">{{ user?.username }}</h2>
              <p class="text-gray-600">{{ user?.email }}</p>
              <p class="text-sm text-gray-500 mt-2">
                Member since {{ formatDate(user?.created_at) }}
              </p>
            </div>
            
            <div class="mt-6 pt-6 border-t border-gray-200">
              <div class="space-y-4">
                <div class="flex items-center justify-between">
                  <span class="text-gray-600">Total uploads</span>
                  <span class="font-semibold text-gray-900">{{ uploads.length }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span class="text-gray-600">Total size</span>
                  <span class="font-semibold text-gray-900">{{ formatTotalSize() }}</span>
                </div>
              </div>
            </div>

            <button
              @click="handleLogout"
              class="w-full mt-6 bg-red-600 hover:bg-red-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
            >
              Sign out
            </button>
          </div>
        </div>

        <!-- Upload History -->
        <div class="lg:col-span-2">
          <div class="bg-white rounded-xl shadow-sm">
            <div class="p-6 border-b border-gray-200">
              <div class="flex items-center justify-between">
                <h3 class="text-lg font-semibold text-gray-900">Upload History</h3>
                <button
                  @click="refreshUploads"
                  :disabled="isRefreshing"
                  class="inline-flex items-center px-3 py-2 text-sm font-medium text-gray-600 hover:text-gray-900 disabled:opacity-50"
                >
                  <ArrowPathIcon 
                    class="w-4 h-4 mr-1"
                    :class="{ 'animate-spin': isRefreshing }"
                  />
                  Refresh
                </button>
              </div>
            </div>

            <div v-if="uploads.length === 0" class="p-8 text-center">
              <CloudArrowUpIcon class="w-12 h-12 mx-auto text-gray-400 mb-4" />
              <h4 class="text-lg font-medium text-gray-900 mb-2">No uploads yet</h4>
              <p class="text-gray-600 mb-4">Start by uploading your first file!</p>
              <router-link
                to="/"
                class="inline-flex items-center px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
              >
                Upload Files
              </router-link>
            </div>

            <div v-else class="divide-y divide-gray-200">
              <div
                v-for="upload in uploads"
                :key="upload.id"
                class="p-6 hover:bg-gray-50 transition-colors"
              >
                <div class="flex items-start justify-between">
                  <div class="flex-1">
                    <div class="flex items-center space-x-3 mb-2">
                      <DocumentIcon class="w-5 h-5 text-gray-400" />
                      <span class="font-medium text-gray-900">
                        {{ getFileNames(upload.files).join(', ') }}
                      </span>
                      <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium"
                            :class="isExpired(upload.expires_at) ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800'">
                        {{ isExpired(upload.expires_at) ? 'Expired' : 'Active' }}
                      </span>
                    </div>
                    
                    <div class="text-sm text-gray-600 space-y-1">
                      <p>Size: {{ formatFileSize(upload.total_size) }}</p>
                      <p>Uploaded: {{ formatDate(upload.created_at) }}</p>
                      <p>Expires: {{ formatDate(upload.expires_at) }}</p>
                      <p v-if="upload.email">Sent to: {{ upload.email }}</p>
                    </div>
                  </div>

                  <div class="flex items-center space-x-2 ml-4">
                    <button
                      v-if="!isExpired(upload.expires_at)"
                      @click="copyShareLink(upload.download_url)"
                      class="inline-flex items-center px-3 py-2 text-sm font-medium text-blue-600 hover:text-blue-700 bg-blue-50 hover:bg-blue-100 rounded-lg transition-colors"
                    >
                      <LinkIcon class="w-4 h-4 mr-1" />
                      Copy Link
                    </button>
                    
                    <router-link
                      v-if="!isExpired(upload.expires_at)"
                      :to="upload.download_url"
                      class="inline-flex items-center px-3 py-2 text-sm font-medium text-green-600 hover:text-green-700 bg-green-50 hover:bg-green-100 rounded-lg transition-colors"
                    >
                      <EyeIcon class="w-4 h-4 mr-1" />
                      View
                    </router-link>

                    <button
                      @click="deleteUpload(upload.upload_id)"
                      class="inline-flex items-center px-3 py-2 text-sm font-medium text-red-600 hover:text-red-700 bg-red-50 hover:bg-red-100 rounded-lg transition-colors"
                    >
                      <TrashIcon class="w-4 h-4 mr-1" />
                      Delete
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Success/Error Messages -->
    <div
      v-if="message"
      class="fixed bottom-4 right-4 max-w-sm w-full bg-white rounded-lg shadow-lg border-l-4 p-4 animate-slide-up"
      :class="message.type === 'success' ? 'border-green-500' : 'border-red-500'"
    >
      <div class="flex items-center">
        <div
          class="flex-shrink-0"
          :class="message.type === 'success' ? 'text-green-500' : 'text-red-500'"
        >
          <CheckCircleIcon v-if="message.type === 'success'" class="w-5 h-5" />
          <XCircleIcon v-else class="w-5 h-5" />
        </div>
        <div class="ml-3">
          <p class="text-sm text-gray-900">{{ message.text }}</p>
        </div>
        <button
          @click="message = null"
          class="ml-auto pl-3"
        >
          <XMarkIcon class="w-4 h-4 text-gray-400 hover:text-gray-600" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useRouter } from 'vue-router'
import {
  UserIcon,
  CloudArrowUpIcon,
  DocumentIcon,
  LinkIcon,
  EyeIcon,
  TrashIcon,
  ArrowLeftIcon,
  ArrowPathIcon,
  CheckCircleIcon,
  XCircleIcon,
  XMarkIcon
} from '@heroicons/vue/24/outline'

interface Message {
  text: string
  type: 'success' | 'error'
}

const { user, uploads, logout, fetchUploads, deleteUpload: removeUpload } = useAuth()
const router = useRouter()

const isRefreshing = ref(false)
const message = ref<Message | null>(null)

const formatDate = (dateString?: string) => {
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

const formatTotalSize = () => {
  const total = uploads.value.reduce((sum, upload) => sum + upload.total_size, 0)
  return formatFileSize(total)
}

const getFileNames = (filesJson: string) => {
  try {
    return JSON.parse(filesJson)
  } catch {
    return ['Unknown files']
  }
}

const isExpired = (expiresAt: string) => {
  return new Date(expiresAt) < new Date()
}

const copyShareLink = async (downloadUrl: string) => {
  try {
    const fullUrl = `${window.location.origin}${downloadUrl}`
    await navigator.clipboard.writeText(fullUrl)
    showMessage('Link copied to clipboard!', 'success')
  } catch (error) {
    showMessage('Failed to copy link', 'error')
  }
}

const deleteUpload = async (uploadId: string) => {
  if (!confirm('Are you sure you want to delete this upload? This action cannot be undone.')) {
    return
  }

  const result = await removeUpload(uploadId)
  if (result.success) {
    showMessage(result.message, 'success')
  } else {
    showMessage(result.message, 'error')
  }
}

const refreshUploads = async () => {
  isRefreshing.value = true
  const result = await fetchUploads()
  if (!result.success) {
    showMessage(result.message, 'error')
  }
  isRefreshing.value = false
}

const handleLogout = async () => {
  await logout()
  router.push('/auth')
}

const showMessage = (text: string, type: 'success' | 'error') => {
  message.value = { text, type }
  setTimeout(() => {
    message.value = null
  }, 5000)
}

onMounted(() => {
  fetchUploads()
})
</script>

<style scoped>
@keyframes slide-up {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-slide-up {
  animation: slide-up 0.3s ease-out;
}
</style>
