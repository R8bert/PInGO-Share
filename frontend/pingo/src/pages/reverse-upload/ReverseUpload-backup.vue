<template>
  <div class="min-h-screen" :class="isDark ? 'bg-gray-900' : 'bg-gray-50'">
    <div class="py-8 px-4">
      <div class="max-w-2xl mx-auto">
        
        <!-- Header -->
        <div class="text-center mb-8">
          <div class="w-12 h-12 mx-auto mb-4 rounded-xl bg-blue-500 flex items-center justify-center">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"/>
            </svg>
          </div>
          <h1 class="text-2xl font-semibold mb-2" :class="isDark ? 'text-white' : 'text-gray-900'">
            Reverse Share Upload
          </h1>
          <p class="text-sm" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
            Upload files using your reverse share token
          </p>
        </div>

        <!-- Upload Form -->
        <div class="rounded-lg border p-6"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-white border-gray-200'">
          
          <form @submit.prevent="handleUpload" class="space-y-6">
            
            <!-- Token Input -->
            <div>
              <label for="token" class="block text-sm font-medium mb-2" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                Reverse Share Token
              </label>
              <input 
                v-model="token"
                type="text" 
                id="token" 
                placeholder="Enter your reverse share token"
                class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                :class="isDark 
                  ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400' 
                  : 'bg-gray-50 border-gray-300 text-gray-900 placeholder-gray-500'"
                required
              />
            </div>

            <!-- Email Input -->
            <div>
              <label for="email" class="block text-sm font-medium mb-2" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                Email Address <span class="text-xs text-gray-500">(optional)</span>
              </label>
              <input 
                v-model="email"
                type="email" 
                id="email" 
                placeholder="your.email@example.com"
                class="w-full px-4 py-3 rounded-lg border focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                :class="isDark 
                  ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400' 
                  : 'bg-gray-50 border-gray-300 text-gray-900 placeholder-gray-500'"
              />
            </div>

            <!-- Expiration Selector -->
            <div>
              <label class="block text-sm font-medium mb-3" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                File Expiration
              </label>
              <div class="grid grid-cols-3 gap-3">
                <button
                  v-for="option in expirationOptions"
                  :key="option.value"
                  type="button"
                  @click="validity = option.value"
                  class="px-4 py-2 rounded-lg text-sm font-medium border transition-colors"
                  :class="validity === option.value
                    ? 'bg-blue-500 text-white border-blue-500'
                    : (isDark 
                        ? 'bg-gray-700 text-gray-300 border-gray-600 hover:bg-gray-600' 
                        : 'bg-gray-50 text-gray-700 border-gray-300 hover:bg-gray-100')"
                >
                  {{ option.label }}
                </button>
              </div>
            </div>

            <!-- File Upload -->
            <div>
              <label class="block text-sm font-medium mb-3" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                Select Files
              </label>
              <div class="border-2 border-dashed rounded-lg p-8 text-center"
                   :class="isDark 
                     ? 'border-gray-600 bg-gray-700/50' 
                     : 'border-gray-300 bg-gray-50'">
                <input
                  type="file"
                  id="files"
                  multiple
                  @change="handleFileSelect"
                  class="hidden"
                />
                <label for="files" class="cursor-pointer">
                  <div class="w-12 h-12 mx-auto mb-4 rounded-lg bg-blue-100 flex items-center justify-center"
                       :class="isDark ? 'bg-blue-900/50' : 'bg-blue-100'">
                    <svg class="w-6 h-6 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"/>
                    </svg>
                  </div>
                  <p class="text-sm font-medium mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                    Click to upload files
                  </p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-500'">
                    or drag and drop files here
                  </p>
                </label>
              </div>
              
              <!-- Selected Files -->
              <div v-if="selectedFiles.length > 0" class="mt-4">
                <p class="text-sm font-medium mb-2" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                  Selected Files ({{ selectedFiles.length }})
                </p>
                <div class="space-y-2">
                  <div v-for="file in selectedFiles" :key="file.name"
                       class="flex items-center justify-between p-3 rounded-lg border"
                       :class="isDark ? 'bg-gray-700 border-gray-600' : 'bg-gray-50 border-gray-200'">
                    <div class="flex items-center gap-3">
                      <div class="w-8 h-8 rounded bg-blue-100 flex items-center justify-center"
                           :class="isDark ? 'bg-blue-900/50' : 'bg-blue-100'">
                        <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                        </svg>
                      </div>
                      <div>
                        <p class="text-sm font-medium" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                          {{ file.name }}
                        </p>
                        <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-500'">
                          {{ formatFileSize(file.size) }}
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="!token || selectedFiles.length === 0 || isUploading"
              class="w-full py-3 px-4 rounded-lg font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              :class="isDark 
                ? 'bg-blue-600 hover:bg-blue-700 text-white' 
                : 'bg-blue-500 hover:bg-blue-600 text-white'"
            >
              <span v-if="isUploading">Uploading...</span>
              <span v-else>Upload Files</span>
            </button>
          </form>
        </div>

        <!-- Upload Progress -->
        <div v-if="isUploading" class="mt-6 p-4 rounded-lg border"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-blue-50 border-blue-200'">
          <div class="flex items-center gap-3">
            <div class="w-5 h-5 border-2 border-blue-500/30 border-t-blue-500 rounded-full animate-spin"></div>
            <span class="text-sm" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
              Uploading files...
            </span>
          </div>
        </div>

        <!-- Success Result -->
        <div v-if="uploadResult && uploadResult.success" class="mt-6 p-6 rounded-lg border"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-green-50 border-green-200'">
          <div class="flex items-start gap-4">
            <div class="w-10 h-10 rounded-lg bg-green-500 flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="font-medium mb-3" :class="isDark ? 'text-green-400' : 'text-green-800'">
                Upload Successful
              </h3>
              <div class="space-y-3 text-sm" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <p class="font-medium">Files Uploaded</p>
                    <p class="text-xs">{{ uploadResult.data.files.join(', ') }}</p>
                  </div>
                  <div>
                    <p class="font-medium">Total Files</p>
                    <p class="text-xs">{{ uploadResult.data.count }}</p>
                  </div>
                  <div>
                    <p class="font-medium">Expires</p>
                    <p class="text-xs">{{ uploadResult.data.expires_at ? new Date(uploadResult.data.expires_at).toLocaleDateString() : 'Never' }}</p>
                  </div>
                </div>
                <div class="p-3 rounded border" :class="isDark ? 'bg-gray-900 border-gray-600' : 'bg-white border-gray-300'">
                  <p class="font-medium mb-1">Download URL</p>
                  <code class="text-xs break-all" :class="isDark ? 'text-green-400' : 'text-green-600'">
                    {{ uploadResult.data.download_url }}
                  </code>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Error Result -->
        <div v-if="uploadResult && !uploadResult.success" class="mt-6 p-6 rounded-lg border"
             :class="isDark ? 'bg-gray-800 border-gray-700' : 'bg-red-50 border-red-200'">
          <div class="flex items-start gap-4">
            <div class="w-10 h-10 rounded-lg bg-red-500 flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </div>
            <div>
              <h3 class="font-medium mb-2" :class="isDark ? 'text-red-400' : 'text-red-800'">
                Upload Failed
              </h3>
              <p class="text-sm" :class="isDark ? 'text-gray-300' : 'text-gray-700'">
                {{ uploadResult.error }}
              </p>
            </div>
          </div>
        </div>
        
      </div>
    </div>
  </div>
</template>
                required
                :disabled="isUploading"
              >
              <div class="absolute inset-y-0 right-0 pr-4 flex items-center">
                <div class="w-2 h-2 bg-indigo-500 rounded-full animate-pulse"></div>
              </div>
            </div>
          </div>
          
          <!-- Email Input -->
          <div class="mb-8 animate-slide-up" style="animation-delay: 0.4s;">
            <label for="email" class="block text-sm font-semibold mb-3" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
              📧 Email Address <span class="text-xs font-normal text-indigo-500">(optional)</span>
            </label>
            <div class="relative group">
              <div class="absolute inset-0 bg-gradient-to-r from-blue-500/20 to-cyan-500/20 rounded-2xl blur opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
              <input 
                v-model="email"
                type="email" 
                id="email" 
                placeholder="your.email@example.com"
                class="relative w-full px-6 py-4 rounded-2xl border-2 transition-all duration-300 focus:outline-none focus:ring-4 focus:ring-blue-500/20 focus:border-blue-500"
                :class="isDark 
                  ? 'bg-gray-800/50 border-gray-600 text-white placeholder-gray-400 focus:bg-gray-800/70' 
                  : 'bg-white/80 border-gray-200 text-gray-900 placeholder-gray-500 focus:bg-white focus:border-blue-400'"
                :disabled="isUploading"
              >
              <div class="absolute inset-y-0 right-0 pr-4 flex items-center">
                <div class="w-2 h-2 bg-blue-500 rounded-full animate-pulse" style="animation-delay: 0.5s;"></div>
              </div>
            </div>
          </div>
          
          <!-- Expiration Selector -->
          <div class="mb-8 animate-slide-up" style="animation-delay: 0.5s;">
            <label class="block text-sm font-semibold mb-4" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
              ⏱️ File Expiration
            </label>
            <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
              <button
                v-for="(option, index) in expirationOptions"
                :key="option.value"
                type="button"
                @click="validity = option.value"
                class="relative px-4 py-3 rounded-xl text-sm font-medium transition-all duration-300 transform hover:scale-105 focus:outline-none focus:ring-4"
                :class="validity === option.value
                  ? 'bg-gradient-to-r from-indigo-500 to-purple-600 text-white shadow-lg shadow-indigo-500/25 ring-indigo-500/20'
                  : (isDark ? 'bg-gray-800/50 border border-gray-600 text-gray-300 hover:bg-gray-700/50 hover:border-gray-500 focus:ring-gray-500/20' : 'bg-white/80 border border-gray-200 text-gray-700 hover:bg-gray-50 hover:border-gray-300 shadow-sm focus:ring-indigo-500/20')"
                :disabled="isUploading"
                :style="`animation-delay: ${0.6 + index * 0.1}s`"
              >
                <span class="relative z-10">{{ option.label }}</span>
                <div v-if="validity === option.value" class="absolute inset-0 bg-gradient-to-r from-indigo-400/20 to-purple-400/20 rounded-xl animate-pulse"></div>
              </button>
            </div>
          </div>
          
          <!-- File Upload -->
          <div class="mb-10 animate-slide-up" style="animation-delay: 0.9s;">
            <label class="block text-sm font-semibold mb-4" :class="isDark ? 'text-gray-200' : 'text-gray-800'">
              📁 Select Files to Share
            </label>
            <div class="relative">
              <input 
                @change="handleFileSelect"
                type="file" 
                id="files" 
                multiple
                class="sr-only"
                required
                :disabled="isUploading"
              >
              <label 
                for="files"
                class="group relative flex flex-col items-center justify-center w-full h-48 border-2 border-dashed rounded-3xl cursor-pointer transition-all duration-300 hover:scale-105 upload-zone"
                :class="isDark 
                  ? 'border-gray-600 bg-gray-800/30 hover:bg-gray-700/30 hover:border-indigo-500/50' 
                  : 'border-gray-300 bg-gradient-to-br from-gray-50/50 to-indigo-50/50 hover:from-indigo-50 hover:to-purple-50 hover:border-indigo-400'"
              >
                <!-- Upload Icon with Animation -->
                <div class="relative mb-6">
                  <div class="absolute inset-0 bg-gradient-to-r from-indigo-500 to-purple-600 rounded-full blur-xl opacity-20 group-hover:opacity-40 transition-opacity duration-300"></div>
                  <div class="relative w-16 h-16 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-2xl flex items-center justify-center shadow-xl group-hover:shadow-2xl transition-all duration-300 transform group-hover:scale-110">
                    <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"/>
                    </svg>
                  </div>
                </div>
                
                <div class="text-center">
                  <p class="text-lg font-semibold mb-2" :class="isDark ? 'text-white' : 'text-gray-900'">
                    {{ selectedFiles.length > 0 ? `${selectedFiles.length} file(s) selected` : 'Drop files here or click to browse' }}
                  </p>
                  <p class="text-sm" :class="isDark ? 'text-gray-400' : 'text-gray-600'">
                    Support for any file type • Maximum 100MB per file
                  </p>
                </div>
                
                <!-- Animated Border -->
                <div class="absolute inset-0 rounded-3xl bg-gradient-to-r from-indigo-500/20 via-purple-500/20 to-pink-500/20 opacity-0 group-hover:opacity-100 transition-opacity duration-300 -z-10 blur-sm"></div>
              </label>
            </div>
            
            <!-- Selected Files with Animation -->
            <div v-if="selectedFiles.length > 0" class="mt-6 space-y-3">
              <div 
                v-for="(file, index) in selectedFiles" 
                :key="index"
                class="flex items-center gap-4 p-4 rounded-2xl border transition-all duration-300 hover:shadow-lg file-item"
                :class="isDark ? 'bg-gray-800/40 border-gray-600/50 hover:bg-gray-700/40' : 'bg-white/80 border-gray-200/60 hover:bg-white'"
                :style="`animation-delay: ${1 + index * 0.1}s`"
              >
                <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-blue-500 to-indigo-600 flex items-center justify-center shadow-lg">
                  <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
                  </svg>
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-sm font-semibold truncate" :class="isDark ? 'text-white' : 'text-gray-900'">{{ file.name }}</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-500'">{{ formatFileSize(file.size) }}</p>
                </div>
                <div class="w-3 h-3 bg-green-500 rounded-full animate-pulse"></div>
              </div>
            </div>
          </div>

          <!-- Upload Button -->
          <button 
            type="submit" 
            class="w-full py-4 px-8 rounded-2xl font-semibold text-white transition-all duration-300 transform hover:scale-105 focus:outline-none focus:ring-4 focus:ring-indigo-500/20 disabled:scale-100 disabled:opacity-50 relative overflow-hidden upload-button animate-slide-up"
            :class="isUploading || !selectedFiles.length || !token
              ? 'bg-gray-400 cursor-not-allowed' 
              : 'bg-gradient-to-r from-indigo-600 via-purple-600 to-indigo-600 hover:from-indigo-700 hover:via-purple-700 hover:to-indigo-700 shadow-xl hover:shadow-2xl shadow-indigo-500/25'"
            :disabled="isUploading || !selectedFiles.length || !token"
            style="animation-delay: 1.2s;"
          >
            <!-- Button Shimmer Effect -->
            <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/20 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
            
            <span v-if="isUploading" class="flex items-center justify-center gap-3">
              <div class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
              Uploading files...
            </span>
            <span v-else class="flex items-center justify-center gap-3">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M9 19l3 3m0 0l3-3m-3 3V10"/>
              </svg>
              Share Files Securely
            </span>
          </button>
        </form>
        
        <!-- Upload Progress -->
        <div v-if="isUploading" class="mt-8 p-6 rounded-2xl border backdrop-blur-sm modern-status"
             :class="isDark ? 'bg-gray-800/60 border-gray-700/50' : 'bg-blue-50/80 border-blue-200/60'"
             style="animation-delay: 0.2s;">
          <div class="flex items-center gap-4">
            <div class="relative">
              <div class="w-12 h-12 border-4 border-blue-500/30 border-t-blue-500 rounded-full animate-spin"></div>
              <div class="absolute inset-0 w-12 h-12 border-4 border-transparent border-t-purple-500 rounded-full animate-spin" style="animation-direction: reverse; animation-duration: 1.5s;"></div>
            </div>
            <div>
              <p class="font-semibold text-lg" :class="isDark ? 'text-blue-400' : 'text-blue-800'">
                Uploading files...
              </p>
              <p class="text-sm" :class="isDark ? 'text-blue-300' : 'text-blue-600'">
                Please wait while we securely process your files
              </p>
            </div>
          </div>
        </div>
        
        <!-- Success Result -->
        <div v-if="uploadResult && uploadResult.success" class="mt-8 p-8 rounded-2xl border backdrop-blur-sm modern-status"
             :class="isDark ? 'bg-gray-800/60 border-gray-700/50' : 'bg-green-50/80 border-green-200/60'"
             style="animation-delay: 0.3s;">
          <div class="flex items-start gap-6">
            <div class="w-16 h-16 rounded-2xl bg-gradient-to-br from-green-500 to-emerald-600 flex items-center justify-center shadow-xl">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="font-bold text-xl mb-4" :class="isDark ? 'text-green-400' : 'text-green-800'">
                🎉 Upload Successful!
              </h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                <div class="p-4 rounded-xl" :class="isDark ? 'bg-gray-900/50' : 'bg-white/60'">
                  <p class="font-semibold text-sm mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">Files Uploaded</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-600'">{{ uploadResult.data.files.join(', ') }}</p>
                </div>
                <div class="p-4 rounded-xl" :class="isDark ? 'bg-gray-900/50' : 'bg-white/60'">
                  <p class="font-semibold text-sm mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">Total Files</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-600'">{{ uploadResult.data.count }}</p>
                </div>
                <div class="p-4 rounded-xl" :class="isDark ? 'bg-gray-900/50' : 'bg-white/60'">
                  <p class="font-semibold text-sm mb-1" :class="isDark ? 'text-gray-300' : 'text-gray-700'">Expires</p>
                  <p class="text-xs" :class="isDark ? 'text-gray-400' : 'text-gray-600'">{{ uploadResult.data.expires_at ? new Date(uploadResult.data.expires_at).toLocaleDateString() : 'Never' }}</p>
                </div>
              </div>
              <div class="p-6 rounded-2xl border-2 border-dashed" :class="isDark ? 'bg-gray-900/30 border-green-500/30' : 'bg-white/80 border-green-300/50'">
                <p class="text-sm font-semibold mb-3" :class="isDark ? 'text-gray-300' : 'text-gray-700'">📎 Download URL</p>
                <code class="text-sm font-mono break-all block p-3 rounded-lg" 
                      :class="isDark ? 'bg-gray-800/50 text-green-400' : 'bg-gray-100 text-green-600'">
                  {{ uploadResult.data.download_url }}
                </code>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Error Result -->
        <div v-if="uploadResult && !uploadResult.success" class="mt-8 p-8 rounded-2xl border backdrop-blur-sm modern-status"
             :class="isDark ? 'bg-gray-800/60 border-gray-700/50' : 'bg-red-50/80 border-red-200/60'"
             style="animation-delay: 0.3s;">
          <div class="flex items-start gap-6">
            <div class="w-16 h-16 rounded-2xl bg-gradient-to-br from-red-500 to-pink-600 flex items-center justify-center shadow-xl">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </div>
            <div>
              <h3 class="font-bold text-xl mb-3" :class="isDark ? 'text-red-400' : 'text-red-800'">
                ❌ Upload Failed
              </h3>
              <p class="text-sm mb-4" :class="isDark ? 'text-red-300' : 'text-red-700'">
                {{ uploadResult.error }}
              </p>
              <div class="p-4 rounded-xl" :class="isDark ? 'bg-gray-900/30' : 'bg-white/60'">
                <p class="text-xs font-medium" :class="isDark ? 'text-red-400' : 'text-red-600'">
                  💡 Please check your reverse token and try again
                </p>
              </div>
            </div>
          </div>
        </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useTheme } from '../../composables/useTheme'

const { isDark } = useTheme()
const route = useRoute()

const token = ref('')
const email = ref('')
const validity = ref('7days')
const selectedFiles = ref<File[]>([])
const isUploading = ref(false)
const uploadResult = ref<{ success: boolean; data?: any; error?: string } | null>(null)

const expirationOptions = [
  { value: '7days', label: '7 Days' },
  { value: '1month', label: '1 Month' },
  { value: '6months', label: '6 Months' },
  { value: '1year', label: '1 Year' },
  { value: 'never', label: 'Never' }
]

onMounted(() => {
  // Pre-fill token from URL parameter if provided
  if (route.params.token) {
    token.value = route.params.token as string
  }
})

const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const handleFileSelect = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files) {
    selectedFiles.value = Array.from(target.files)
  }
}

const handleUpload = async () => {
  if (!token.value || !selectedFiles.value.length) {
    alert('Please enter a token and select files')
    return
  }

  isUploading.value = true
  uploadResult.value = null

  const formData = new FormData()
  
  // Add files
  for (const file of selectedFiles.value) {
    formData.append('files', file)
  }
  
  // Add metadata
  if (email.value) formData.append('email', email.value)
  formData.append('validity', validity.value)

  try {
    const response = await fetch(`http://localhost:8080/reverse-upload/${token.value}`, {
      method: 'POST',
      body: formData
    })

    const data = await response.json()

    if (response.ok) {
      uploadResult.value = {
        success: true,
        data
      }
      // Reset form
      email.value = ''
      validity.value = '7days'
      selectedFiles.value = []
      // Reset file input
      const fileInput = document.getElementById('files') as HTMLInputElement
      if (fileInput) fileInput.value = ''
    } else {
      throw new Error(data.error || 'Upload failed')
    }
  } catch (error) {
    uploadResult.value = {
      success: false,
      error: error instanceof Error ? error.message : 'Upload failed'
    }
  } finally {
    isUploading.value = false
  }
}
</script>

<style scoped>
/* Animated background elements */
.floating-orb {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(45deg, rgba(59, 130, 246, 0.1), rgba(147, 51, 234, 0.1));
  backdrop-filter: blur(1px);
  animation: float 8s ease-in-out infinite;
}

.floating-orb:nth-child(1) {
  width: 100px;
  height: 100px;
  top: 10%;
  left: 10%;
  animation-delay: 0s;
}

.floating-orb:nth-child(2) {
  width: 150px;
  height: 150px;
  top: 60%;
  right: 10%;
  animation-delay: 2s;
}

.floating-orb:nth-child(3) {
  width: 80px;
  height: 80px;
  bottom: 20%;
  left: 20%;
  animation-delay: 4s;
}

.floating-orb:nth-child(4) {
  width: 120px;
  height: 120px;
  top: 30%;
  right: 30%;
  animation-delay: 6s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) translateX(0px) rotate(0deg);
  }
  25% {
    transform: translateY(-20px) translateX(10px) rotate(90deg);
  }
  50% {
    transform: translateY(-10px) translateX(-10px) rotate(180deg);
  }
  75% {
    transform: translateY(-30px) translateX(15px) rotate(270deg);
  }
}

/* Modern animations */
.fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
  opacity: 0;
  transform: translateY(30px);
}

@keyframes fadeInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.slide-in {
  animation: slideIn 1s ease-out forwards;
  opacity: 0;
  transform: translateX(-50px);
}

@keyframes slideIn {
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.scale-in {
  animation: scaleIn 0.6s ease-out forwards;
  opacity: 0;
  transform: scale(0.9);
}

@keyframes scaleIn {
  to {
    opacity: 1;
    transform: scale(1);
  }
}

.pulse-glow {
  animation: pulseGlow 2s ease-in-out infinite;
}

@keyframes pulseGlow {
  0%, 100% {
    box-shadow: 0 0 20px rgba(59, 130, 246, 0.3);
  }
  50% {
    box-shadow: 0 0 40px rgba(59, 130, 246, 0.6);
  }
}

.shimmer {
  position: relative;
  overflow: hidden;
}

.shimmer::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  animation: shimmer 2s ease-in-out infinite;
}

@keyframes shimmer {
  0% {
    left: -100%;
  }
  100% {
    left: 100%;
  }
}

/* Hover effects */
.modern-input:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

.modern-input:focus-within {
  transform: translateY(-3px);
  box-shadow: 0 15px 35px rgba(59, 130, 246, 0.2);
}

.modern-button:hover {
  transform: translateY(-3px);
  box-shadow: 0 15px 35px rgba(59, 130, 246, 0.3);
}

.modern-button:active {
  transform: translateY(-1px);
}

.modern-upload-zone:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(59, 130, 246, 0.2);
}

.modern-status {
  animation: fadeInUp 0.8s ease-out forwards;
  opacity: 0;
  transform: translateY(20px);
}

/* Icon animations */
.bounce-icon {
  animation: bounceIcon 2s ease-in-out infinite;
}

@keyframes bounceIcon {
  0%, 20%, 50%, 80%, 100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-10px);
  }
  60% {
    transform: translateY(-5px);
  }
}

/* Progress ring animation */
.progress-ring {
  animation: progressRing 2s ease-in-out infinite;
}

@keyframes progressRing {
  0% {
    stroke-dasharray: 0 100;
  }
  100% {
    stroke-dasharray: 100 0;
  }
}

/* Gradient text animation */
.gradient-text {
  background: linear-gradient(-45deg, #3b82f6, #8b5cf6, #06b6d4, #3b82f6);
  background-size: 400% 400%;
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: gradientShift 4s ease infinite;
}

@keyframes gradientShift {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

/* Responsive animations */
@media (prefers-reduced-motion: reduce) {
  .floating-orb,
  .fade-in-up,
  .slide-in,
  .scale-in,
  .pulse-glow,
  .shimmer,
  .bounce-icon,
  .progress-ring,
  .gradient-text {
    animation: none;
  }
  
  .modern-input:hover,
  .modern-button:hover,
  .modern-upload-zone:hover {
    transform: none;
  }
}

/* Glass morphism effects */
.glass-card {
  backdrop-filter: blur(16px) saturate(180%);
  -webkit-backdrop-filter: blur(16px) saturate(180%);
  border: 1px solid rgba(255, 255, 255, 0.125);
}

/* Smooth transitions */
* {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

input, select, button {
  transition: all 0.2s ease;
}
</style>
