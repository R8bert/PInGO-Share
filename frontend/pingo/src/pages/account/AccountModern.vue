<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50 pt-20">
    <!-- Top Navigation -->
    <nav class="bg-white/80 backdrop-blur-sm border-b border-white/20 fixed top-16 left-0 right-0 z-40">
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
              <div class="flex items-center space-x-2">
                <span class="text-gray-700 font-medium">{{ user?.username }}</span>
                <span v-if="isAdmin" class="admin-badge">
                  <span class="admin-badge-text">ADMIN</span>
                </span>
              </div>
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
                <div class="flex items-center space-x-3">
          <h1 class="text-2xl font-bold text-gray-900">
            {{ user?.username }}'s Dashboard
          </h1>
          <span v-if="isAdmin" class="admin-badge">
            <span class="admin-badge-text">ADMIN</span>
          </span>
        </div>
        <p class="text-xl text-gray-600 mb-4">{{ user?.username }}</p>
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
              @click="activeTab = 'statistics'"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === 'statistics'
                  ? 'border-purple-500 text-purple-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              <div class="flex items-center space-x-2">
                <ChartBarIcon class="w-4 h-4" />
                <span>Statistics</span>
              </div>
            </button>
            <button
              v-if="isAdmin"
              @click="activeTab = 'users'"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === 'users'
                  ? 'border-orange-500 text-orange-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              <div class="flex items-center space-x-2">
                <UsersIcon class="w-4 h-4" />
                <span>Users</span>
              </div>
            </button>
            <button
              v-if="isAdmin"
              @click="activeTab = 'settings'"
              :class="[
                'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === 'settings'
                  ? 'border-blue-500 text-blue-600'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
              ]"
            >
              <div class="flex items-center space-x-2">
                <CogIcon class="w-4 h-4" />
                <span>Settings</span>
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
                    
                    <!-- Expiration Dropdown -->
                    <div class="relative">
                      <button
                        @click="toggleExpirationDropdown(upload.upload_id)"
                        class="p-2 text-gray-400 hover:text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
                        title="Change expiration"
                      >
                        <ClockIcon class="w-5 h-5" />
                      </button>
                      <div v-if="showExpirationDropdown[upload.upload_id]" class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg border border-gray-200 z-10">
                        <div class="py-1">
                          <button
                            @click="changeExpiration(upload.upload_id, '7days')"
                            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                          >
                            7 days
                          </button>
                          <button
                            @click="changeExpiration(upload.upload_id, '1month')"
                            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                          >
                            1 month
                          </button>
                          <button
                            @click="changeExpiration(upload.upload_id, '6months')"
                            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                          >
                            6 months
                          </button>
                          <button
                            @click="changeExpiration(upload.upload_id, '1year')"
                            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                          >
                            1 year
                          </button>
                          <button
                            @click="changeExpiration(upload.upload_id, 'never')"
                            class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                          >
                            Never expires
                          </button>
                        </div>
                      </div>
                    </div>
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

          <!-- Statistics Tab -->
          <div v-if="activeTab === 'statistics' && isAdmin">
            <!-- Loading skeleton -->
            <div v-if="isLoadingAdminSettings" class="space-y-6">
              <div class="bg-gray-200 animate-pulse rounded-xl p-6 border">
                <div class="h-6 bg-gray-300 rounded w-1/3 mb-4"></div>
                <div class="space-y-3">
                  <div class="h-4 bg-gray-300 rounded w-full"></div>
                  <div class="h-4 bg-gray-300 rounded w-2/3"></div>
                </div>
              </div>
            </div>
            
            <!-- Statistics Content -->
            <div v-else class="space-y-8">
              <div>
                <h3 class="text-xl font-bold text-gray-900 mb-6">Platform Statistics</h3>
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                  <div class="bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl p-6 border border-blue-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <h3 class="text-lg font-semibold text-gray-900">Total Users</h3>
                        <p class="text-3xl font-bold text-blue-600">{{ adminStats.totalUsers }}</p>
                      </div>
                      <UserGroupIcon class="w-12 h-12 text-blue-500 opacity-20" />
                    </div>
                  </div>
                  
                  <div class="bg-gradient-to-r from-green-50 to-emerald-50 rounded-xl p-6 border border-green-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <h3 class="text-lg font-semibold text-gray-900">Total Uploads</h3>
                        <p class="text-3xl font-bold text-green-600">{{ adminStats.totalUploads }}</p>
                      </div>
                      <CloudArrowUpIcon class="w-12 h-12 text-green-500 opacity-20" />
                    </div>
                  </div>
                  
                  <div class="bg-gradient-to-r from-purple-50 to-pink-50 rounded-xl p-6 border border-purple-200">
                    <div class="flex items-center justify-between">
                      <div>
                        <h3 class="text-lg font-semibold text-gray-900">Storage Used</h3>
                        <p class="text-3xl font-bold text-purple-600">{{ formatBytes(adminStats.totalStorage) }}</p>
                      </div>
                      <ArchiveBoxIcon class="w-12 h-12 text-purple-500 opacity-20" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Users Tab -->
          <div v-if="activeTab === 'users' && isAdmin">
            <!-- Loading skeleton -->
            <div v-if="isLoadingAdminSettings" class="space-y-6">
              <div class="bg-gray-200 animate-pulse rounded-xl p-6 border">
                <div class="h-6 bg-gray-300 rounded w-1/4 mb-4"></div>
                <div class="space-y-3">
                  <div class="h-4 bg-gray-300 rounded w-full"></div>
                  <div class="h-4 bg-gray-300 rounded w-3/4"></div>
                </div>
              </div>
            </div>
            
            <!-- User Management Content -->
            <div v-else class="space-y-8">
              <div>
                <h3 class="text-xl font-bold text-gray-900 mb-6">User Management</h3>
                <div class="bg-white rounded-xl p-6 border border-gray-200">
                  <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200">
                      <thead class="bg-gray-50">
                        <tr>
                          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">User</th>
                          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Uploads</th>
                          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Storage</th>
                          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                        </tr>
                      </thead>
                      <tbody class="bg-white divide-y divide-gray-200">
                        <tr v-for="user in adminUsers" :key="user.id" class="hover:bg-gray-50">
                          <td class="px-6 py-4 whitespace-nowrap">
                            <div class="flex items-center">
                              <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mr-3">
                                <UserIcon class="w-4 h-4 text-blue-600" />
                              </div>
                              <div>
                                <div class="text-sm font-medium text-gray-900">{{ user.username }}</div>
                                <div class="text-sm text-gray-500">{{ user.email }}</div>
                              </div>
                            </div>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ user.uploadCount || 0 }}</td>
                          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ formatBytes(user.storageUsed || 0) }}</td>
                          <td class="px-6 py-4 whitespace-nowrap">
                            <span :class="[
                              'inline-flex px-2 py-1 text-xs font-semibold rounded-full',
                              user.isBlocked ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800'
                            ]">
                              {{ user.isBlocked ? 'Blocked' : 'Active' }}
                            </span>
                          </td>
                          <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                            <button
                              @click="toggleUserBlock(user.id, !user.isBlocked)"
                              :class="[
                                'px-3 py-1 rounded-md text-sm font-medium transition-colors',
                                user.isBlocked 
                                  ? 'bg-green-600 hover:bg-green-700 text-white' 
                                  : 'bg-red-600 hover:bg-red-700 text-white'
                              ]"
                            >
                              {{ user.isBlocked ? 'Unblock' : 'Block' }}
                            </button>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Settings Tab -->
          <div v-if="activeTab === 'settings' && isAdmin">
            <!-- Loading skeleton -->
            <div v-if="isLoadingAdminSettings" class="space-y-6">
              <div class="bg-gray-200 animate-pulse rounded-xl p-6 border">
                <div class="h-6 bg-gray-300 rounded w-1/3 mb-4"></div>
                <div class="space-y-3">
                  <div class="h-4 bg-gray-300 rounded w-full"></div>
                  <div class="h-4 bg-gray-300 rounded w-2/3"></div>
                </div>
              </div>
            </div>
            
            <!-- Settings Content -->
            <div v-else class="space-y-8">
              <div>
                <h3 class="text-xl font-bold text-gray-900 mb-6">System Settings</h3>
                <div class="bg-white rounded-xl p-6 border border-gray-200">
                  <h4 class="text-lg font-semibold text-gray-900 mb-4">Quick Settings</h4>
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div class="flex items-center justify-between p-4 bg-gray-50 rounded-lg">
                      <div>
                        <p class="font-medium text-gray-900">User Registration</p>
                        <p class="text-sm text-gray-600">Allow new user signups</p>
                      </div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input
                          type="checkbox"
                          v-model="quickSettings.allowRegistration"
                          @change="updateQuickSetting('allowRegistration', quickSettings.allowRegistration)"
                          class="sr-only peer"
                        />
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600"></div>
                      </label>
                    </div>
                  </div>
                </div>
                
                <!-- Advanced Settings -->
                <div class="bg-white rounded-xl p-6 border border-gray-200">
                  <h4 class="text-lg font-semibold text-gray-900 mb-4">Advanced Settings</h4>
                  <div class="space-y-6">
                    <!-- Navbar Title -->
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">
                        Navbar Title
                      </label>
                      <input
                        type="text"
                        v-model="advancedSettings.navbarTitle"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                        placeholder="PInGO Share"
                      />
                    </div>
                    
                    <!-- Logo Upload -->
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">
                        Logo
                      </label>
                      <input
                        type="file"
                        ref="logoInput"
                        @change="handleLogoChange"
                        accept="image/*"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      />
                      <p class="text-sm text-gray-500 mt-1">Upload a new logo (PNG, JPG, SVG)</p>
                    </div>
                    
                    <!-- Max Upload Size -->
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">
                        Maximum Upload Size (MB)
                      </label>
                      <input
                        type="number"
                        v-model="advancedSettings.maxUploadSize"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                        min="1"
                        max="1000"
                      />
                    </div>
                    
                    <!-- Default Validity -->
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">
                        Default File Validity
                      </label>
                      <select
                        v-model="advancedSettings.maxValidity"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      >
                        <option value="7days">7 Days</option>
                        <option value="1month">1 Month</option>
                        <option value="6months">6 Months</option>
                        <option value="1year">1 Year</option>
                        <option value="never">Never</option>
                      </select>
                    </div>
                    
                    <!-- Theme -->
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">
                        Theme
                      </label>
                      <select
                        v-model="advancedSettings.theme"
                        class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500"
                      >
                        <option value="light">Light</option>
                        <option value="dark">Dark</option>
                      </select>
                    </div>
                    
                    <!-- Save Button -->
                    <div class="flex justify-end pt-4">
                      <button
                        @click="saveAllSettings"
                        :disabled="isSavingSettings"
                        class="px-6 py-2 bg-blue-600 hover:bg-blue-700 disabled:bg-blue-400 text-white font-medium rounded-lg transition-colors flex items-center"
                      >
                        <span v-if="isSavingSettings" class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin mr-2"></span>
                        {{ isSavingSettings ? 'Saving...' : 'Save Settings' }}
                      </button>
                    </div>
                  </div>
                </div>
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
import { ref, computed, onMounted, watch } from 'vue'
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
  EyeSlashIcon,
  ChartBarIcon,
  UsersIcon,
  UserGroupIcon
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

// Admin functionality
const isLoadingAdminSettings = ref(true)
const adminStats = ref({
  totalUsers: 0,
  totalUploads: 0,
  totalStorage: 0
})
const quickSettings = ref({
  allowRegistration: true
})
const advancedSettings = ref({
  maxUploadSize: 100, // in MB
  maxValidity: '7days',
  theme: 'light',
  navbarTitle: 'PInGO Share'
})
const isSavingSettings = ref(false)
const logoFile = ref<File | null>(null)
const adminUsers = ref<any[]>([])

// Expiration dropdown state
const showExpirationDropdown = ref<Record<string, boolean>>({})

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

// Admin functions
const formatBytes = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const updateQuickSetting = async (setting: string, value: boolean) => {
  try {
    const response = await fetch('http://localhost:8080/admin/quick-settings', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      body: JSON.stringify({
        setting,
        value
      })
    })

    if (!response.ok) {
      console.error('Failed to update setting')
    }
  } catch (error) {
    console.error('Error updating setting:', error)
  }
}


const toggleUserBlock = async (userId: number, isBlocked: boolean) => {
  try {
    const response = await fetch(`http://localhost:8080/admin/users/${userId}/block`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      body: JSON.stringify({
        blocked: isBlocked
      })
    })

    if (response.ok) {
      // Update local state
      const userIndex = adminUsers.value.findIndex(u => u.id === userId)
      if (userIndex !== -1) {
        adminUsers.value[userIndex].isBlocked = isBlocked
      }
    } else {
      console.error('Failed to toggle user block status')
    }
  } catch (error) {
    console.error('Error toggling user block:', error)
  }
}

// Change upload expiration
const changeExpiration = async (uploadId: string, validity: string) => {
  try {
    const response = await fetch(`http://localhost:8080/uploads/${uploadId}/expiration`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      body: JSON.stringify({
        validity
      })
    })

    if (response.ok) {
      // Refresh uploads to get updated expiration dates
      await fetchUploads()
      // Close dropdown
      showExpirationDropdown.value[uploadId] = false
    } else {
      console.error('Failed to update expiration')
    }
  } catch (error) {
    console.error('Error updating expiration:', error)
  }
}

// Toggle expiration dropdown
const toggleExpirationDropdown = (uploadId: string) => {
  // Close all other dropdowns
  Object.keys(showExpirationDropdown.value).forEach(key => {
    if (key !== uploadId) {
      showExpirationDropdown.value[key] = false
    }
  })
  
  // Toggle the current dropdown
  showExpirationDropdown.value[uploadId] = !showExpirationDropdown.value[uploadId]
}

const fetchAdminData = async () => {
  if (!isAdmin.value) return
  
  try {
    isLoadingAdminSettings.value = true
    
    // Fetch admin statistics
    const statsResponse = await fetch('http://localhost:8080/admin/stats', {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      }
    })
    
    if (statsResponse.ok) {
      const stats = await statsResponse.json()
      adminStats.value = {
        totalUsers: stats.totalUsers,
        totalUploads: stats.totalUploads,
        totalStorage: stats.storageUsed
      }
    }

    // Fetch quick settings from main settings endpoint
    const settingsResponse = await fetch('http://localhost:8080/settings')
    if (settingsResponse.ok) {
      const settings = await settingsResponse.json()
      quickSettings.value = {
        allowRegistration: settings.allowRegistration
      }
      advancedSettings.value = {
        maxUploadSize: Math.round(settings.maxUploadSize / (1024 * 1024)), // Convert bytes to MB
        maxValidity: settings.maxValidity,
        theme: settings.theme,
        navbarTitle: settings.navbarTitle || 'PInGO Share'
      }
    }
    
    // Fetch admin users
    const usersResponse = await fetch('http://localhost:8080/admin/users', {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      }
    })
    
    if (usersResponse.ok) {
      const users = await usersResponse.json()
      adminUsers.value = users
    }
    
  } catch (error) {
    console.error('Error fetching admin data:', error)
  } finally {
    isLoadingAdminSettings.value = false
  }
}

onMounted(async () => {
  if (isAuthenticated.value) {
    await fetchCurrentUser()
    await fetchUploads()
    await fetchReverseTokens()
    if (isAdmin.value) {
      await fetchAdminData()
    }
  }
})

// Watch for tab changes to load admin data when needed
watch(activeTab, async (newTab: string) => {
  if ((newTab === 'statistics' || newTab === 'users' || newTab === 'settings') && isAdmin.value && adminStats.value.totalUsers === 0) {
    await fetchAdminData()
  }
})

// Handle logo file change
const handleLogoChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files[0]) {
    logoFile.value = target.files[0]
  }
}

// Save all settings
const saveAllSettings = async () => {
  isSavingSettings.value = true
  try {
    const formData = new FormData()
    
    // Add all advanced settings
    formData.append('theme', advancedSettings.value.theme)
    formData.append('navbarTitle', advancedSettings.value.navbarTitle)
    formData.append('maxValidity', advancedSettings.value.maxValidity)
    formData.append('maxUploadSize', (advancedSettings.value.maxUploadSize * 1024 * 1024).toString())
    
    // Add logo file if selected
    if (logoFile.value) {
      formData.append('logo', logoFile.value)
    }
    
    // Add quick settings
    formData.append('allowRegistration', quickSettings.value.allowRegistration.toString())

    const response = await fetch('http://localhost:8080/admin/settings', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('auth_token') || ''}`
      },
      body: formData
    })

    if (response.ok) {
      // Show success message or reload settings
      await fetchAdminData()
      logoFile.value = null
      // Reset file input
      const logoInput = document.querySelector('input[type="file"]') as HTMLInputElement
      if (logoInput) {
        logoInput.value = ''
      }
    } else {
      console.error('Failed to save settings')
    }
  } catch (error) {
    console.error('Error saving settings:', error)
  } finally {
    isSavingSettings.value = false
  }
}

</script>
