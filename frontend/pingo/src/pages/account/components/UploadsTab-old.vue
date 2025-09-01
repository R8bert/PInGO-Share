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
                  <img src="/src/assets/svg/icons/files_uploaded.svg" alt="Total uploads" class="w-6 h-6" />
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
                  <img src="/src/assets/svg/icons/total_files_icon.svg" alt="Total files" class="w-6 h-6" />
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
          <img src="/src/assets/svg/icons/files_uploaded2.svg" alt="No uploads" class="w-12 h-12 opacity-50" />
        </div>
        <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">No uploads found</h3>
        <p class="text-gray-600 dark:text-gray-400 max-w-md mx-auto">
          {{ deletionFilter === 'all' ? 'You haven\'t uploaded any files yet.' : `No ${deletionFilter} uploads found.` }}
        </p>
      </div>

      <!-- Uploads Grid -->
      <div v-else class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        <transition-group name="list" tag="div" class="contents">
          <div
            v-for="upload in displayedUploads"
            :key="upload.upload_id"
            class="group bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 hover:shadow-md transition-all duration-200 overflow-hidden"
            :class="selectedUpload?.id === upload.id ? 'ring-2 ring-blue-500' : ''"
          >
            <!-- Upload Header -->
            <div class="p-6 pb-4">
              <div class="flex items-start justify-between mb-4">
                <div class="flex items-center gap-3">
                  <!-- File Type Icon -->
                  <div class="w-12 h-12 rounded-lg bg-gray-100 dark:bg-gray-700 flex items-center justify-center">
                    <img 
                      v-if="getFileTypeSvg(upload) === 'pdf'" 
                      src="/src/assets/svg/icons/files_pdf.svg" 
                      alt="PDF" 
                      class="w-8 h-8"
                    />
                    <img 
                      v-else-if="getFileTypeSvg(upload) === 'text'" 
                      src="/src/assets/svg/icons/word_docx_doc.svg" 
                      alt="Document" 
                      class="w-8 h-8"
                    />
                    <svg v-else-if="getFileTypeSvg(upload) === 'image'" class="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                    </svg>
                    <svg v-else-if="getFileTypeSvg(upload) === 'video'" class="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                    </svg>
                    <svg v-else class="w-8 h-8 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  
                  <div>
                    <h3 class="font-semibold text-gray-900 dark:text-white">Upload {{ upload.upload_id }}</h3>
                    <p class="text-sm text-gray-500 dark:text-gray-400">{{ JSON.parse(upload.files).length }} files</p>
                  </div>
                </div>

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

              <!-- Upload Details -->
              <div class="space-y-2 text-sm text-gray-600 dark:text-gray-400">
                <div class="flex items-center justify-between">
                  <span>Created:</span>
                  <span>{{ formatDate(upload.created_at) }}</span>
                </div>
                <div v-if="upload.expires_at" class="flex items-center justify-between">
                  <span>Expires:</span>
                  <span>{{ formatDate(upload.expires_at) }}</span>
                </div>
                <div class="flex items-center justify-between">
                  <span>Size:</span>
                  <span>{{ formatFileSize(upload.size || 0) }}</span>
                </div>
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="px-6 pb-6">
              <div class="flex gap-2">
                <button
                  @click="openFile(upload)"
                  class="flex-1 bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center justify-center gap-2"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                  </svg>
                  Open
                </button>
                
                <button
                  @click="$emit('copyToClipboard', `${getBaseUrl()}/download/${upload.upload_id}`)"
                  class="bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center justify-center"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                  </svg>
                </button>
                
                <button
                  @click="selectUpload(upload)"
                  class="bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 text-gray-700 dark:text-gray-300 px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center justify-center"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>
              </div>
            </div>

            <!-- Expanded Details (when selected) -->
            <div v-if="selectedUpload?.id === upload.id" class="border-t border-gray-200 dark:border-gray-700 p-6 bg-gray-50 dark:bg-gray-900/50">
              <h4 class="font-medium text-gray-900 dark:text-white mb-3">Files in this upload:</h4>
              <div class="space-y-2 max-h-40 overflow-y-auto">
                <div v-for="(file, fileIndex) in JSON.parse(upload.files)" :key="fileIndex" 
                     class="flex items-center gap-3 p-2 bg-white dark:bg-gray-800 rounded-lg">
                  <div class="w-8 h-8 bg-gray-100 dark:bg-gray-700 rounded flex items-center justify-center">
                    <img 
                      v-if="getFileSvg(file) === 'pdf'" 
                      src="/src/assets/svg/icons/files_pdf.svg" 
                      alt="PDF" 
                      class="w-5 h-5"
                    />
                    <img 
                      v-else-if="getFileSvg(file) === 'text'" 
                      src="/src/assets/svg/icons/word_docx_doc.svg" 
                      alt="Document" 
                      class="w-5 h-5"
                    />
                    <svg v-else class="w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  <span class="text-sm text-gray-900 dark:text-white truncate">{{ file }}</span>
                </div>
              </div>
              
              <!-- Additional Actions -->
              <div class="flex gap-2 mt-4">
                <button
                  @click="downloadSingleFile(upload)"
                  class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                  Download
                </button>
                
                <button
                  v-if="!upload.is_deleted"
                  @click="$emit('toggleAvailability', upload.upload_id, upload.is_available)"
                  :class="[
                    'px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2',
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
                  {{ upload.is_available ? 'Hide' : 'Show' }}
                </button>
                
                <button
                  @click="$emit('deleteUpload', upload.upload_id)"
                  class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 flex items-center gap-2"
                >
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                  Delete
                </button>
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
        <div class="relative">
          <div class="w-20 h-20 border-4 border-purple-200 dark:border-purple-800 rounded-full animate-spin border-t-purple-600"></div>
          <div class="absolute inset-0 w-20 h-20 border-4 border-pink-200 dark:border-pink-800 rounded-full animate-ping"></div>
        </div>
        <p class="text-2xl font-bold mt-6 bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent">
          🌟 Loading your cosmic files...
        </p>
        <div class="flex space-x-1 mt-4">
          <div class="w-2 h-2 bg-purple-500 rounded-full animate-bounce" style="animation-delay: 0ms"></div>
          <div class="w-2 h-2 bg-pink-500 rounded-full animate-bounce" style="animation-delay: 150ms"></div>
          <div class="w-2 h-2 bg-blue-500 rounded-full animate-bounce" style="animation-delay: 300ms"></div>
        </div>
      </div>
      
      <!-- Magical Empty State -->
      <div v-else-if="uploads.length === 0" class="text-center py-20">
        <div class="relative inline-block mb-8">
          <div class="w-32 h-32 bg-gradient-to-br from-purple-500 via-pink-500 to-blue-500 rounded-3xl flex items-center justify-center mx-auto shadow-2xl transform rotate-3 hover:rotate-0 transition-transform duration-500">
            <div class="absolute inset-0 bg-gradient-to-br from-purple-400 via-pink-400 to-blue-400 rounded-3xl blur-xl opacity-70 animate-pulse"></div>
            <div class="relative text-6xl">🚀</div>
          </div>
          <div class="absolute -top-2 -right-2 text-3xl animate-bounce">✨</div>
          <div class="absolute -bottom-2 -left-2 text-2xl animate-pulse">🌟</div>
        </div>
        <h3 class="text-4xl font-black mb-4 bg-gradient-to-r from-purple-600 via-pink-500 to-blue-500 bg-clip-text text-transparent">
          Ready for Liftoff!
        </h3>
        <p class="text-xl text-gray-600 dark:text-gray-300 mb-8">
          Your digital space station is empty. Time to launch some files into orbit! 🌌
        </p>
        <router-link
          to="/"
          class="group inline-flex items-center space-x-3 px-8 py-4 bg-gradient-to-r from-purple-600 via-pink-600 to-blue-600 hover:from-purple-700 hover:via-pink-700 hover:to-blue-700 text-white font-bold text-lg rounded-2xl transition-all duration-300 shadow-2xl hover:shadow-purple-500/25 transform hover:scale-105"
        >
          <span class="text-2xl group-hover:animate-bounce">🚀</span>
          <span>Launch Your First Upload</span>
          <span class="text-xl">✨</span>
        </router-link>
      </div>

      <!-- No Results Space -->
      <div v-else-if="displayedUploads.length === 0" class="text-center py-16">
        <div class="w-24 h-24 bg-gradient-to-br from-gray-400 to-gray-600 rounded-2xl flex items-center justify-center mx-auto mb-6 opacity-50">
          <div class="text-4xl">🔍</div>
        </div>
        <h3 class="text-2xl font-bold mb-2 text-gray-600 dark:text-gray-300">
          No files found in this dimension
        </h3>
        <p class="text-lg text-gray-500 dark:text-gray-400">
          Try exploring different filter galaxies to discover more uploads! 🌌
        </p>
      </div>

      <!-- FILE LIST VIEW -->
      <div v-else class="space-y-4">
        <transition-group name="list" tag="div" class="space-y-4">
          <div
            v-for="upload in displayedUploads"
            :key="upload.upload_id"
            class="group relative overflow-hidden rounded-2xl transition-all duration-300 hover:shadow-xl cursor-pointer"
            :style="{ background: getStatusGradient(upload) }"
            @click="selectUpload(upload)"
            :class="selectedUpload?.id === upload.id ? 'ring-4 ring-blue-500 ring-opacity-50' : ''"
          >
          <!-- Main Content Row -->
          <div class="relative p-6">
            <!-- Background overlay -->
            <div class="absolute inset-0 bg-gradient-to-r from-black/20 via-transparent to-black/10 opacity-60"></div>
            
            <div class="relative flex items-center space-x-6">
              <!-- File Type Icon & Count -->
              <div class="flex-shrink-0">
                <div class="relative">
                  <div class="w-16 h-16 rounded-2xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                    <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path v-if="getFileTypeSvg(upload) === 'image'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      <path v-else-if="getFileTypeSvg(upload) === 'video'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                      <path v-else-if="getFileTypeSvg(upload) === 'audio'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zM9 10l12-3" />
                      <path v-else-if="getFileTypeSvg(upload) === 'pdf'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                      <path v-else-if="getFileTypeSvg(upload) === 'archive'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                      <path v-else-if="getFileTypeSvg(upload) === 'text'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                      <path v-else-if="getFileTypeSvg(upload) === 'code'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
                      <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                  </div>
                  <div class="absolute -top-2 -right-2 bg-blue-600 text-white text-xs font-bold px-2 py-1 rounded-full shadow-lg">
                    {{ JSON.parse(upload.files).length }}
                  </div>
                </div>
              </div>

              <!-- File Details -->
            <!-- File Details -->
            <div class="flex-1 min-w-0">
              <div class="flex items-start justify-between">
                <div class="flex-1 min-w-0 pr-4">
                  <div class="flex items-center gap-3 mb-2">
                    <h3 class="text-xl font-bold text-white">
                      Upload {{ upload.upload_id }}
                    </h3>
                    <!-- Expiration Warning Badge - moved here -->
                    <div v-if="!upload.is_deleted && upload.expires_at && isExpiringSoon(upload.expires_at)" 
                         class="bg-yellow-600 text-white text-xs font-bold px-2 py-1 rounded-full animate-pulse">
                      <svg class="w-3 h-3 inline mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                      </svg>
                      Expiring Soon
                    </div>
                  </div>                    <!-- File Names Strip -->
                    <div class="flex flex-wrap gap-2 mb-3">
                      <span v-for="(filename, fileIndex) in JSON.parse(upload.files).slice(0, 4)" :key="fileIndex"
                            class="inline-flex items-center space-x-2 px-3 py-1 bg-white/20 backdrop-blur-sm rounded-lg text-white/90 text-sm font-medium border border-white/20">
                        <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path v-if="getFileSvg(filename) === 'image'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                          <path v-else-if="getFileSvg(filename) === 'video'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z" />
                          <path v-else-if="getFileSvg(filename) === 'audio'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zM9 10l12-3" />
                          <path v-else-if="getFileSvg(filename) === 'pdf'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                          <path v-else-if="getFileSvg(filename) === 'archive'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                          <path v-else-if="getFileSvg(filename) === 'code'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
                          <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                        </svg>
                        <span class="truncate max-w-32">{{ filename }}</span>
                      </span>
                      <span v-if="JSON.parse(upload.files).length > 4"
                            class="inline-flex items-center px-3 py-1 bg-gray-600/70 backdrop-blur-sm rounded-lg text-white text-sm font-medium border border-white/20">
                        +{{ JSON.parse(upload.files).length - 4 }} more files
                      </span>
                    </div>

                    <!-- Metadata Row -->
                    <div class="flex items-center space-x-4 text-white/70 text-sm">
                      <div class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                        <span>{{ formatDate(upload.created_at) }}</span>
                      </div>
                      <div class="flex items-center space-x-1">
                        <svg v-if="upload.is_deleted" class="w-4 h-4 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                        <svg v-else-if="!upload.is_available" class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21" />
                        </svg>
                        <svg v-else-if="upload.expires_at && isExpiringSoon(upload.expires_at)" class="w-4 h-4 text-orange-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <svg v-else class="w-4 h-4 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <span>{{ getStatusText(upload) }}</span>
                      </div>
                      <div v-if="!upload.is_deleted && upload.expires_at" class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                        <span>Expires {{ formatDate(upload.expires_at) }}</span>
                      </div>
                      <div v-else-if="!upload.is_deleted && !upload.expires_at" class="flex items-center space-x-1">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                        </svg>
                        <span>Never expires</span>
                      </div>
                    </div>
                  </div>

                  <!-- Action Buttons -->
                  <div class="flex-shrink-0 flex items-center space-x-2 opacity-0 group-hover:opacity-100 transition-all duration-300 translate-x-4 group-hover:translate-x-0">
                    <button v-if="!upload.is_deleted"
                            @click.stop="openFile(upload)"
                            class="action-btn-list bg-gray-700 hover:bg-gray-800"
                            title="Open file">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                      </svg>
                    </button>

                    <button v-if="!upload.is_deleted"
                            @click.stop="downloadSingleFile(upload)"
                            class="action-btn-list bg-indigo-600 hover:bg-indigo-700"
                            title="Download files">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                      </svg>
                    </button>

                    <button v-if="!upload.is_deleted"
                            @click.stop="$emit('copyToClipboard', `${getBaseUrl()}/download/${upload.upload_id}`)"
                            class="action-btn-list bg-blue-500 hover:bg-blue-600"
                            title="Copy link">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                      </svg>
                    </button>
                    
                    <button v-if="!upload.is_deleted"
                            @click.stop="$emit('toggleAvailability', upload.upload_id, upload.is_available)"
                            :class="upload.is_available ? 'action-btn-list bg-orange-500 hover:bg-orange-600' : 'action-btn-list bg-green-500 hover:bg-green-600'"
                            :title="upload.is_available ? 'Hide' : 'Show'">
                      <svg v-if="upload.is_available" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21" />
                      </svg>
                      <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                      </svg>
                    </button>
                    
                    <button v-if="!upload.is_deleted"
                            @click.stop="$emit('deleteUpload', upload.upload_id)"
                            class="action-btn-list bg-red-500 hover:bg-red-600"
                            title="Delete upload permanently">
                      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        </transition-group>
      </div>
      
      <!-- Pagination Space Controls -->
      <div v-if="filteredUploads.length > itemsPerPage" class="mt-12 text-center">
        <div class="inline-flex items-center space-x-4 bg-white/10 backdrop-blur-sm rounded-2xl p-4 border border-white/20">
          <button @click="currentPage = 0"
                  :disabled="currentPage === 0"
                  class="pagination-btn"
                  :class="currentPage === 0 ? 'opacity-50 cursor-not-allowed' : 'hover:scale-110'">
            ⏮️
          </button>
          
          <button @click="currentPage--"
                  :disabled="currentPage === 0"
                  class="pagination-btn"
                  :class="currentPage === 0 ? 'opacity-50 cursor-not-allowed' : 'hover:scale-110'">
            ⬅️
          </button>
          
          <div class="px-6 py-2 bg-gradient-to-r from-purple-500 to-pink-500 text-white font-bold rounded-xl">
            🚀 {{ currentPage + 1 }} / {{ totalPages }}
          </div>
          
          <button @click="currentPage++"
                  :disabled="currentPage >= totalPages - 1"
                  class="pagination-btn"
                  :class="currentPage >= totalPages - 1 ? 'opacity-50 cursor-not-allowed' : 'hover:scale-110'">
            ➡️
          </button>
          
          <button @click="currentPage = totalPages - 1"
                  :disabled="currentPage >= totalPages - 1"
                  class="pagination-btn"
                  :class="currentPage >= totalPages - 1 ? 'opacity-50 cursor-not-allowed' : 'hover:scale-110'">
            ⏭️
          </button>
        </div>
        
        <p class="text-sm text-gray-600 dark:text-gray-400 mt-4">
          Showing {{ currentPage * itemsPerPage + 1 }} to {{ Math.min((currentPage + 1) * itemsPerPage, filteredUploads.length) }} of {{ filteredUploads.length }} cosmic uploads
        </p>
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
const itemsPerPage = ref(12) // Show more in grid view
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

// Visual helper methods
const getStatusGradient = (upload: any) => {
  // Red gradient for deleted or expired uploads
  if (upload.is_deleted || !upload.is_available) {
    return 'linear-gradient(90deg, #ef4444 0%, #dc2626 100%)'
  }
  
  // Green gradient for all active uploads (including expiring soon)
  return 'linear-gradient(90deg, #22c55e 0%, #16a34a 100%)'
}

const getFileTypeSvg = (upload: any) => {
  const files = JSON.parse(upload.files)
  const extensions = files.map((file: string) => getFileExtension(file).toLowerCase())
  
  if (extensions.some((ext: string) => ['jpg', 'jpeg', 'png', 'gif', 'webp'].includes(ext))) return 'image'
  if (extensions.some((ext: string) => ['mp4', 'mov', 'avi', 'mkv'].includes(ext))) return 'video'
  if (extensions.some((ext: string) => ['mp3', 'wav', 'flac', 'ogg'].includes(ext))) return 'audio'
  if (extensions.some((ext: string) => ['pdf'].includes(ext))) return 'pdf'
  if (extensions.some((ext: string) => ['zip', 'rar', '7z'].includes(ext))) return 'archive'
  if (extensions.some((ext: string) => ['txt', 'md', 'doc', 'docx'].includes(ext))) return 'text'
  if (extensions.some((ext: string) => ['js', 'ts', 'html', 'css', 'vue', 'jsx', 'tsx'].includes(ext))) return 'code'
  return 'file'
}

const getFileSvg = (filename: string) => {
  const ext = getFileExtension(filename).toLowerCase()
  
  if (['jpg', 'jpeg', 'png', 'gif', 'webp'].includes(ext)) return 'image'
  if (['mp4', 'mov', 'avi', 'mkv'].includes(ext)) return 'video'
  if (['mp3', 'wav', 'flac', 'ogg'].includes(ext)) return 'audio'
  if (['pdf'].includes(ext)) return 'pdf'
  if (['zip', 'rar', '7z'].includes(ext)) return 'archive'
  if (['txt', 'md', 'doc', 'docx'].includes(ext)) return 'text'
  if (['js', 'ts', 'html', 'css', 'vue', 'jsx', 'tsx'].includes(ext)) return 'code'
  return 'file'
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

// Single file download function
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

    // Get the filename from Content-Disposition header or use a default
    const contentDisposition = response.headers.get('Content-Disposition')
    let filename = 'download'
    if (contentDisposition) {
      const matches = contentDisposition.match(/filename="([^"]*)"/)
      if (matches) {
        filename = matches[1]
      }
    } else {
      // If multiple files, create a zip name, otherwise use the single file name
      const files = JSON.parse(upload.files)
      filename = files.length === 1 ? files[0] : `upload_${upload.upload_id}.zip`
    }

    // Create blob and download
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

// Methods
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
/* List transition animations for smooth delete */
.list-enter-active,
.list-leave-active {
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
}

.list-enter-from {
  opacity: 0;
  transform: translateY(-20px) scale(0.95);
}

.list-leave-to {
  opacity: 0;
  transform: translateX(-100%) scale(0.95);
}

.list-move {
  transition: transform 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
}

/* Magical Cosmic Styles */
.grid-pattern {
  background-image: 
    radial-gradient(circle at 1px 1px, rgba(255,255,255,0.15) 1px, transparent 0);
  background-size: 20px 20px;
  animation: grid-float 20s ease-in-out infinite;
}

.stats-bubble {
  padding: 1.5rem;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(4px);
  border-radius: 1rem;
  border: 1px solid rgba(255, 255, 255, 0.2);
  text-align: center;
  transition: all 0.3s ease;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.stats-bubble:hover {
  transform: scale(1.05);
  background: rgba(255, 255, 255, 0.2);
}

.action-btn {
  width: 3rem;
  height: 3rem;
  border-radius: 0.75rem;
  color: white;
  font-weight: bold;
  transition: all 0.3s ease;
  transform: scale(1);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.action-btn:hover {
  transform: scale(1.1);
}

.action-btn-list {
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 0.75rem;
  color: white;
  font-weight: bold;
  transition: all 0.3s ease;
  transform: scale(1);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.action-btn-list:hover {
  transform: scale(1.1);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
}

.pagination-btn {
  width: 3rem;
  height: 3rem;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(4px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 0.75rem;
  font-size: 1.5rem;
  font-weight: bold;
  color: #6b7280;
  transition: all 0.3s ease;
}

/* Masonry columns responsive behavior */
@media (max-width: 768px) {
  .columns-1 {
    column-count: 1;
  }
}

@media (min-width: 768px) and (max-width: 1280px) {
  .md\:columns-2 {
    column-count: 2;
  }
}

@media (min-width: 1280px) and (max-width: 1536px) {
  .xl\:columns-3 {
    column-count: 3;
  }
}

@media (min-width: 1536px) {
  .\32xl\:columns-4 {
    column-count: 4;
  }
}

/* Animations */
@keyframes grid-float {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  33% { transform: translateY(-10px) rotate(1deg); }
  66% { transform: translateY(-5px) rotate(-1deg); }
}

/* Card hover effects */
.group:hover .group-hover\:scale-110 {
  transform: scale(1.1);
}

.group:hover .group-hover\:opacity-100 {
  opacity: 1;
}

.group:hover .group-hover\:translate-y-0 {
  transform: translateY(0);
}

/* Smooth transitions for all interactive elements */
* {
  transition: all 0.2s ease;
}

/* Custom scrollbar for better aesthetics */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: linear-gradient(45deg, #667eea, #764ba2);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(45deg, #764ba2, #667eea);
}

/* Fix for column breaks */
.break-inside-avoid {
  break-inside: avoid;
  page-break-inside: avoid;
}

/* Enhanced file card animations */
.group {
  transform-origin: center;
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

.group:hover {
  transform: translateY(-2px);
  filter: brightness(1.05);
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
}

/* Magical glow effects */
.action-btn:hover {
  box-shadow: 0 0 20px rgba(255, 255, 255, 0.3);
}

/* Text gradient animations */
.bg-clip-text {
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  animation: gradient-shift 3s ease-in-out infinite;
}

@keyframes gradient-shift {
  0%, 100% { filter: hue-rotate(0deg); }
  50% { filter: hue-rotate(10deg); }
}

/* Enhanced stats bubbles */
.stats-bubble:hover {
  background: rgba(255, 255, 255, 0.2);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  transform: scale(1.05) rotate(1deg);
}

/* Pulse animation for loading dots */
@keyframes bounce {
  0%, 20%, 50%, 80%, 100% { transform: translateY(0); }
  40% { transform: translateY(-10px); }
  60% { transform: translateY(-5px); }
}

/* Filter pill animations */
button {
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}

button:active {
  transform: scale(0.95);
}

/* Card content animations */
.group .text-4xl {
  transition: all 0.3s ease;
}

.group:hover .text-4xl {
  transform: rotate(10deg) scale(1.1);
}

/* Backdrop blur support */
.backdrop-blur-xl {
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
}

.backdrop-blur-sm {
  backdrop-filter: blur(4px);
  -webkit-backdrop-filter: blur(4px);
}

/* Enhanced ring animations for selected cards */
.ring-4 {
  animation: ring-pulse 2s ease-in-out infinite;
}

@keyframes ring-pulse {
  0%, 100% { 
    box-shadow: 0 0 0 4px rgba(168, 85, 247, 0.5);
  }
  50% { 
    box-shadow: 0 0 0 8px rgba(168, 85, 247, 0.3);
  }
}
</style>
