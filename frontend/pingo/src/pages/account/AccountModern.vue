<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50">
    <!-- Top Navigation -->
    <nav class="bg-white/80 backdrop-blur-sm border-b border-white/20 sticky top-0 z-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center space-x-3">
            <router-link to="/" class="flex items-center space-x-2 group">
              <ArrowLeftIcon class="w-5 h-5 text-gray-600 group-hover:text-blue-600 transition-colors" />
              <span class="text-gray-700 group-hover:text-blue-600 font-medium transition-colors">Back to Home</span>
            </router-link>
          </div>
          <div class="flex items-center space-x-4">
            <div class="flex items-center space-x-2">
              <div class="w-8 h-8 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-full flex items-center justify-center">
                <UserIcon class="w-4 h-4 text-white" />
              </div>
              <span class="text-gray-700 font-medium">{{ user?.username }}</span>
              <span v-if="isAdmin" class="px-2 py-1 bg-gradient-to-r from-purple-500 to-pink-500 text-white text-xs font-medium rounded-full">
                ADMIN
              </span>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Header Section -->
      <div class="text-center mb-12">
        <div class="inline-flex items-center justify-center w-20 h-20 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-2xl mb-6 shadow-lg">
          <UserIcon class="w-10 h-10 text-white" />
        </div>
        <h1 class="text-4xl font-bold bg-gradient-to-r from-gray-900 to-gray-600 bg-clip-text text-transparent mb-2">
          {{ user?.username }}'s Dashboard
        </h1>
        <p class="text-xl text-gray-600 mb-4">{{ user?.email }}</p>
        <div class="inline-flex items-center space-x-4 text-sm text-gray-500">
          <span class="flex items-center">
            <CalendarIcon class="w-4 h-4 mr-1" />
            Member since {{ formatDate(user?.created_at) }}
          </span>
          <span v-if="isAdmin" class="flex items-center text-purple-600">
            <ShieldCheckIcon class="w-4 h-4 mr-1" />
            Administrator
          </span>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div class="bg-white/70 backdrop-blur-sm rounded-2xl p-6 border border-white/20 shadow-sm hover:shadow-md transition-all">
          <div class="flex items-center">
            <div class="p-3 bg-blue-100 rounded-xl">
              <CloudArrowUpIcon class="w-6 h-6 text-blue-600" />
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Total Uploads</p>
              <p class="text-2xl font-bold text-gray-900">{{ uploads.length }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white/70 backdrop-blur-sm rounded-2xl p-6 border border-white/20 shadow-sm hover:shadow-md transition-all">
          <div class="flex items-center">
            <div class="p-3 bg-green-100 rounded-xl">
              <ArchiveBoxIcon class="w-6 h-6 text-green-600" />
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Total Size</p>
              <p class="text-2xl font-bold text-gray-900">{{ formatTotalSize() }}</p>
            </div>
          </div>
        </div>

        <div class="bg-white/70 backdrop-blur-sm rounded-2xl p-6 border border-white/20 shadow-sm hover:shadow-md transition-all">
          <div class="flex items-center">
            <div class="p-3 bg-purple-100 rounded-xl">
              <ClockIcon class="w-6 h-6 text-purple-600" />
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-600">Active Files</p>
              <p class="text-2xl font-bold text-gray-900">{{ activeUploads.length }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Content Tabs -->
      <div class="bg-white/70 backdrop-blur-sm rounded-2xl border border-white/20 shadow-sm">
        <!-- Tab Navigation -->
        <div class="border-b border-gray-200/50">
          <nav class="flex space-x-8 px-6">
            <button
              @click="activeTab = 'uploads'"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === 'uploads'
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              <div class="flex items-center space-x-2">
                <FolderIcon class="w-4 h-4" />
                <span>My Uploads</span>
              </div>
            </button>
            <button
              @click="activeTab = 'reverse'"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === 'reverse'
                  ? 'border-green-500 text-green-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              <div class="flex items-center space-x-2">
                <ShareIcon class="w-4 h-4" />
                <span>Reverse Share</span>
              </div>
            </button>
            <button
              v-if="isAdmin"
              @click="activeTab = 'admin'"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === 'admin'
                  ? 'border-purple-500 text-purple-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              <div class="flex items-center space-x-2">
                <CogIcon class="w-4 h-4" />
                <span>Admin Settings</span>
              </div>
            </button>
          </nav>
        </div>

        <!-- Tab Content -->
        <div class="p-6">
          <!-- Uploads Tab -->
          <div v-if="activeTab === 'uploads'">
            <div v-if="isLoading" class="flex justify-center py-12">
              <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            </div>
            
            <div v-else-if="uploads.length === 0" class="text-center py-12">
              <CloudArrowUpIcon class="w-16 h-16 text-gray-300 mx-auto mb-4" />
              <h3 class="text-lg font-medium text-gray-900 mb-2">No uploads yet</h3>
              <p class="text-gray-500 mb-6">Start by uploading your first file.</p>
              <router-link
                to="/"
                class="inline-flex items-center px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
              >
                Upload Files
              </router-link>
            </div>

            <div v-else class="space-y-4">
              <div
                v-for="upload in uploads"
                :key="upload.id"
                class="bg-gray-50/50 rounded-xl p-6 border border-gray-200/50 hover:bg-gray-50 transition-colors"
              >
                <div class="flex items-start justify-between">
                  <div class="flex-1">
                    <div class="flex items-center space-x-3 mb-3">
                      <div class="p-2 bg-blue-100 rounded-lg">
                        <DocumentIcon class="w-5 h-5 text-blue-600" />
                      </div>
                      <div>
                        <h3 class="font-medium text-gray-900">
                          {{ JSON.parse(upload.files).length }} file{{ JSON.parse(upload.files).length > 1 ? 's' : '' }}
                        </h3>
                        <p class="text-sm text-gray-500">
                          {{ formatFileSize(upload.total_size) }} • {{ formatDate(upload.created_at) }}
                        </p>
                      </div>
                    </div>
                    
                    <div class="flex flex-wrap gap-2 mb-3">
                      <span
                        v-for="filename in JSON.parse(upload.files)"
                        :key="filename"
                        class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
                      >
                        {{ filename }}
                      </span>
                    </div>

                    <div class="flex items-center space-x-4 text-sm">
                      <span class="flex items-center text-gray-500">
                        <LinkIcon class="w-4 h-4 mr-1" />
                        {{ upload.download_url }}
                      </span>
                      <span v-if="upload.expires_at" :class="[
                        'flex items-center',
                        isExpiringSoon(upload.expires_at) ? 'text-red-600' : 'text-gray-500'
                      ]">
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
                      <span :class="[
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
                    <button
                      @click="copyToClipboard(`${getBaseUrl()}${upload.download_url}`)"
                      class="p-2 text-gray-400 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
                      title="Copy link"
                    >
                      <ClipboardDocumentIcon class="w-5 h-5" />
                    </button>
                    <button
                      @click="toggleAvailability(upload.upload_id, upload.is_available)"
                      :class="[
                        'p-2 rounded-lg transition-colors',
                        upload.is_available 
                          ? 'text-gray-400 hover:text-orange-600 hover:bg-orange-50'
                          : 'text-gray-400 hover:text-green-600 hover:bg-green-50'
                      ]"
                      :title="upload.is_available ? 'Make unavailable' : 'Make available'"
                    >
                      <EyeSlashIcon v-if="upload.is_available" class="w-5 h-5" />
                      <EyeIcon v-else class="w-5 h-5" />
                    </button>
                    <button
                      @click="deleteUpload(upload.upload_id)"
                      class="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                      title="Delete upload"
                    >
                      <TrashIcon class="w-5 h-5" />
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Reverse Share Tab -->
          <div v-if="activeTab === 'reverse'">
            <div class="mb-6">
              <div class="flex items-center justify-between mb-4">
                <div>
                  <h3 class="text-lg font-medium text-gray-900">Reverse Share Tokens</h3>
                  <p class="text-sm text-gray-600">Allow others to upload files to your account without registering</p>
                </div>
                <button
                  @click="showCreateToken = true"
                  class="inline-flex items-center px-4 py-2 bg-green-600 hover:bg-green-700 text-white font-medium rounded-lg transition-colors"
                >
                  <PlusIcon class="w-4 h-4 mr-2" />
                  Create Token
                </button>
              </div>

              <!-- Create Token Form -->
              <div v-if="showCreateToken" class="bg-gray-50 rounded-xl p-6 border border-gray-200 mb-6">
                <h4 class="text-lg font-medium text-gray-900 mb-4">Create New Token</h4>
                <form @submit.prevent="createReverseToken" class="space-y-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">Token Name</label>
                    <input
                      v-model="newToken.name"
                      type="text"
                      required
                      placeholder="e.g., Client Uploads"
                      class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
                    />
                  </div>
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">Max Uses</label>
                      <select
                        v-model.number="newToken.max_uses"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
                      >
                        <option :value="-1">Unlimited</option>
                        <option :value="1">1 use</option>
                        <option :value="5">5 uses</option>
                        <option :value="10">10 uses</option>
                        <option :value="25">25 uses</option>
                        <option :value="100">100 uses</option>
                      </select>
                    </div>
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">Expires In</label>
                      <select
                        v-model="newToken.expires_in"
                        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500 focus:border-green-500"
                      >
                        <option value="">Never</option>
                        <option value="7days">7 days</option>
                        <option value="1month">1 month</option>
                        <option value="6months">6 months</option>
                        <option value="1year">1 year</option>
                      </select>
                    </div>
                  </div>
                  <div class="flex space-x-3">
                    <button
                      type="submit"
                      :disabled="isLoading"
                      class="inline-flex items-center px-4 py-2 bg-green-600 hover:bg-green-700 text-white font-medium rounded-lg transition-colors disabled:opacity-50"
                    >
                      Create Token
                    </button>
                    <button
                      type="button"
                      @click="showCreateToken = false; resetNewToken()"
                      class="inline-flex items-center px-4 py-2 bg-gray-300 hover:bg-gray-400 text-gray-700 font-medium rounded-lg transition-colors"
                    >
                      Cancel
                    </button>
                  </div>
                </form>
              </div>

              <!-- Tokens List -->
              <div v-if="reverseTokens.length === 0" class="text-center py-12">
                <ShareIcon class="w-16 h-16 text-gray-300 mx-auto mb-4" />
                <h3 class="text-lg font-medium text-gray-900 mb-2">No reverse share tokens</h3>
                <p class="text-gray-500 mb-6">Create a token to allow others to upload files to your account.</p>
              </div>

              <div v-else class="space-y-4">
                <div
                  v-for="token in reverseTokens"
                  :key="token.id"
                  class="bg-gray-50/50 rounded-xl p-6 border border-gray-200/50"
                >
                  <div class="flex items-start justify-between">
                    <div class="flex-1">
                      <h4 class="font-medium text-gray-900 mb-2">{{ token.name }}</h4>
                      <div class="space-y-2 text-sm text-gray-600">
                        <div class="flex items-center space-x-4">
                          <span>Uses: {{ token.used_count }}{{ token.max_uses !== -1 ? ` / ${token.max_uses}` : ' (unlimited)' }}</span>
                          <span v-if="token.expires_at">Expires: {{ formatDate(token.expires_at) }}</span>
                          <span v-else>Never expires</span>
                        </div>
                        <div class="bg-gray-100 rounded p-2 font-mono text-xs break-all">
                          Upload URL: {{ getBaseUrl() }}/reverse-upload/{{ token.token }}
                        </div>
                      </div>
                    </div>
                    <div class="flex items-center space-x-2">
                      <button
                        @click="copyToClipboard(`${getBaseUrl()}/reverse-upload/${token.token}`)"
                        class="p-2 text-gray-400 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
                        title="Copy upload URL"
                      >
                        <ClipboardDocumentIcon class="w-5 h-5" />
                      </button>
                      <button
                        @click="deleteReverseToken(token.id)"
                        class="p-2 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                        title="Delete token"
                      >
                        <TrashIcon class="w-5 h-5" />
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Admin Settings Tab -->
          <div v-if="activeTab === 'admin' && isAdmin">
            <div class="space-y-8">
              <div class="text-center">
                <h2 class="text-2xl font-bold text-gray-900 mb-2">Admin Settings</h2>
                <p class="text-gray-600">Admin panel functionality is being loaded...</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Success Message for Copy -->
    <div
      v-if="showCopySuccess"
      class="fixed bottom-4 right-4 bg-green-500 text-white px-6 py-3 rounded-lg shadow-lg z-50 transition-all"
    >
      <div class="flex items-center">
        <ClipboardDocumentIcon class="w-5 h-5 mr-2" />
        Link copied to clipboard!
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useAuth } from '../../composables/useAuth'
import {
  UserIcon,
  ArrowLeftIcon,
  CloudArrowUpIcon,
  ArchiveBoxIcon,
  ClockIcon,
  FolderIcon,
  CogIcon,
  DocumentIcon,
  LinkIcon,
  ClipboardDocumentIcon,
  TrashIcon,
  CalendarIcon,
  ShieldCheckIcon,
  ShareIcon,
  PlusIcon,
  EyeIcon,
  EyeSlashIcon
} from '@heroicons/vue/24/outline'

// For infinity icon, we'll use a simple component or text
const InfinityIcon = { template: '<span class="text-lg">∞</span>' }

const { user, uploads, isAuthenticated, isAdmin, isLoading, fetchCurrentUser, fetchUploads, deleteUpload } = useAuth()

const activeTab = ref('uploads')
const showCreateToken = ref(false)
const reverseTokens = ref<ReverseToken[]>([])
const newToken = ref({
  name: '',
  max_uses: -1,
  expires_in: ''
})

// Success message for copy operations
const showCopySuccess = ref(false)

interface ReverseToken {
  id: number
  token: string
  name: string
  used_count: number
  max_uses: number
  created_at: string
  expires_at: string | null
}

const activeUploads = computed(() => {
  return uploads.value.filter(upload => {
    if (!upload.expires_at) return true // Never expires
    return new Date(upload.expires_at) > new Date()
  })
})

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

const formatTotalSize = () => {
  const totalBytes = uploads.value.reduce((sum, upload) => sum + upload.total_size, 0)
  return formatFileSize(totalBytes)
}

const isExpiringSoon = (expiresAt: string) => {
  if (!expiresAt) return false
  const expiry = new Date(expiresAt)
  const now = new Date()
  const timeDiff = expiry.getTime() - now.getTime()
  const daysDiff = timeDiff / (1000 * 3600 * 24)
  return daysDiff <= 1 && daysDiff > 0
}

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    showCopySuccess.value = true
    setTimeout(() => {
      showCopySuccess.value = false
    }, 2000)
  } catch (err) {
    console.error('Failed to copy: ', err)
    // Fallback for older browsers
    const textArea = document.createElement('textarea')
    textArea.value = text
    document.body.appendChild(textArea)
    textArea.select()
    document.execCommand('copy')
    document.body.removeChild(textArea)
  }
}

const getBaseUrl = () => {
  return window.location.origin
}

const toggleAvailability = async (uploadId: string, currentAvailability: boolean) => {
  try {
    const response = await fetch(`http://localhost:8080/uploads/${uploadId}/availability`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include',
      body: JSON.stringify({ is_available: !currentAvailability })
    })

    if (response.ok) {
      // Refresh uploads to get updated data
      await fetchUploads()
    } else {
      throw new Error('Failed to update availability')
    }
  } catch (error) {
    console.error('Error toggling availability:', error)
    alert('Failed to update file availability')
  }
}

const fetchReverseTokens = async () => {
  try {
    const response = await fetch('http://localhost:8080/reverse-tokens', {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include'
    })

    if (response.ok) {
      const data = await response.json()
      reverseTokens.value = data.tokens || []
    }
  } catch (error) {
    console.error('Error fetching reverse tokens:', error)
  }
}

const createReverseToken = async () => {
  try {
    // Add validation before sending
    if (!newToken.value.name || newToken.value.name.trim() === '') {
      alert('Please enter a token name')
      return
    }

    // Ensure max_uses is sent as a number
    const tokenData = {
      name: newToken.value.name.trim(),
      max_uses: parseInt(newToken.value.max_uses.toString()),
      expires_in: newToken.value.expires_in
    }

    console.log('Creating token with data:', tokenData)
    
    const response = await fetch('http://localhost:8080/reverse-tokens', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include',
      body: JSON.stringify(tokenData)
    })

    const responseData = await response.json()
    console.log('Response:', responseData)

    if (response.ok) {
      await fetchReverseTokens()
      showCreateToken.value = false
      resetNewToken()
    } else {
      throw new Error(responseData.error || 'Failed to create token')
    }
  } catch (error) {
    console.error('Error creating reverse token:', error)
    const errorMessage = error instanceof Error ? error.message : 'Unknown error'
    alert(`Failed to create reverse share token: ${errorMessage}`)
  }
}

const deleteReverseToken = async (tokenId: number) => {
  if (!confirm('Are you sure you want to delete this token?')) return

  try {
    const response = await fetch(`http://localhost:8080/reverse-tokens/${tokenId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      credentials: 'include'
    })

    if (response.ok) {
      await fetchReverseTokens()
    } else {
      throw new Error('Failed to delete token')
    }
  } catch (error) {
    console.error('Error deleting reverse token:', error)
    alert('Failed to delete reverse share token')
  }
}

const resetNewToken = () => {
  newToken.value = {
    name: '',
    max_uses: -1,
    expires_in: ''
  }
}

onMounted(async () => {
  if (isAuthenticated.value) {
    await fetchCurrentUser()
    await fetchUploads()
    await fetchReverseTokens()
  }
})
</script>
