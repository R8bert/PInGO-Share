<template>
  <div class="min-h-screen flex flex-col transition-colors duration-300" 
       :style="{ backgroundColor: isDark ? '#000000' : '#f8fafc' }">
    <!-- Main Content -->
    <main class="relative">
      <!-- Hero Section with background effect -->
      <div class="relative overflow-hidden min-h-screen flex items-center justify-center transition-colors duration-300">
        <!-- Simplified background for dark mode -->
        <div v-if="isDark" class="fixed inset-0 pointer-events-none overflow-hidden z-0">
          <!-- Subtle animated dots -->
          <div class="absolute inset-0 opacity-20">
            <div class="absolute top-1/4 left-1/4 w-2 h-2 bg-blue-500 rounded-full animate-pulse"></div>
            <div class="absolute top-1/3 right-1/3 w-1 h-1 bg-purple-500 rounded-full animate-pulse" style="animation-delay: 1s;"></div>
            <div class="absolute bottom-1/3 left-1/3 w-1.5 h-1.5 bg-green-500 rounded-full animate-pulse" style="animation-delay: 2s;"></div>
            <div class="absolute bottom-1/4 right-1/4 w-1 h-1 bg-yellow-500 rounded-full animate-pulse" style="animation-delay: 0.5s;"></div>
          </div>
          
          <!-- Subtle gradient overlay -->
          <div class="absolute inset-0 bg-gradient-to-br from-blue-900/10 via-transparent to-purple-900/10"></div>
        </div>
        
        <!-- Light mode subtle background -->
        <div v-else class="absolute inset-0 pointer-events-none">
          <div class="absolute top-1/4 left-1/4 w-64 h-64 rounded-full blur-3xl opacity-20 bg-gradient-to-r from-blue-400 to-purple-400 animate-pulse"></div>
          <div class="absolute bottom-1/3 right-1/3 w-48 h-48 rounded-full blur-3xl opacity-15 bg-gradient-to-r from-green-400 to-blue-400 animate-pulse" style="animation-delay: 1s;"></div>
        </div>
        
        <div class="relative max-w-6xl mx-auto px-6 py-20 text-center">
          <!-- Main headline -->
          <div class="mb-8">
            <h1 class="text-6xl md:text-7xl lg:text-8xl font-bold mb-4 tracking-tight" 
                :style="{ color: isDark ? '#ffffff' : '#000000' }">
              Share
              <span class="block text-transparent bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 bg-clip-text">
                Simply
              </span>
            </h1>
            
            <p class="text-xl md:text-2xl max-w-3xl mx-auto font-medium opacity-80"
               :style="{ color: isDark ? '#e5e7eb' : '#374151' }">
              Transfer files of any size, securely and instantly. No registration required.
            </p>
          </div>
          <!-- Upload Card -->
          <div class="max-w-2xl mx-auto">
            <div class="backdrop-blur-xl rounded-3xl p-8 border border-white/10 transition-all duration-300"
                 :style="{ 
                   backgroundColor: isDark ? 'rgba(0, 0, 0, 0.2)' : 'rgba(255, 255, 255, 0.8)',
                   boxShadow: isDark 
                     ? '0 25px 50px -12px rgba(0, 0, 0, 0.5), 0 0 0 1px rgba(255, 255, 255, 0.05)' 
                     : '0 25px 50px -12px rgba(0, 0, 0, 0.1), 0 0 0 1px rgba(0, 0, 0, 0.05)'
                 }">
              
              <!-- Upload Zone -->
              <div 
                @dragover.prevent 
                @drop.prevent="onDrop" 
                @dragenter="isDragging = true" 
                @dragleave="isDragging = false"
                @click="triggerFileInput"
                :class="[
                  'relative border-2 border-dashed rounded-2xl p-12 cursor-pointer transition-all duration-300',
                  isDragging 
                    ? 'border-blue-400 bg-blue-500/10 transform scale-105' 
                    : 'border-gray-300 dark:border-gray-600 hover:border-blue-400 dark:hover:border-blue-500'
                ]"
              >
                <input type="file" ref="fileInput" @change="onFileChange" class="hidden" multiple />
                
                <div class="text-center">
                  <div class="mb-6">
                    <div class="w-20 h-20 mx-auto mb-4 rounded-2xl flex items-center justify-center transition-all duration-300"
                         :style="{ 
                           backgroundColor: isDragging 
                             ? 'rgba(59, 130, 246, 0.2)' 
                             : (isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.05)')
                         }">
                      <CloudArrowUpIcon 
                        :class="[
                          'w-10 h-10 transition-all duration-300',
                          isDragging 
                            ? 'text-blue-500 scale-110' 
                            : (isDark ? 'text-gray-400' : 'text-gray-500')
                        ]" 
                      />
                    </div>
                  </div>
                  
                  <h3 class="text-2xl font-bold mb-3 transition-colors duration-300"
                      :style="{ color: isDark ? '#ffffff' : '#111827' }">
                    {{ selectedFiles.length > 0 ? `${selectedFiles.length} file${selectedFiles.length > 1 ? 's' : ''} selected` : 'Drop files here' }}
                  </h3>
                  
                  <p class="text-lg mb-8 opacity-70"
                     :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                    {{ selectedFiles.length > 0 ? selectedFiles.map(f => f.name).slice(0, 2).join(', ') + (selectedFiles.length > 2 ? ` and ${selectedFiles.length - 2} more` : '') : 'or click to browse your device' }}
                  </p>
                  
                  <button 
                    v-if="selectedFiles.length === 0"
                    class="inline-flex items-center px-8 py-4 bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600 text-white font-bold rounded-2xl transition-all duration-300 hover:scale-105 shadow-lg hover:shadow-xl"
                  >
                    <CloudArrowUpIcon class="w-6 h-6 mr-3" />
                    Choose Files
                  </button>
                </div>

                <!-- Drag overlay -->
                <div 
                  v-if="isDragging" 
                  class="absolute inset-0 rounded-2xl flex items-center justify-center backdrop-blur-sm"
                  :style="{ backgroundColor: 'rgba(59, 130, 246, 0.1)' }"
                >
                  <div class="text-center">
                    <CloudArrowUpIcon class="w-12 h-12 mx-auto mb-4 text-blue-500" />
                    <div class="font-bold text-xl text-blue-600 dark:text-blue-400">Drop files here</div>
                  </div>
                </div>
              </div>
              <!-- Upload Actions -->
              <div v-if="selectedFiles.length > 0" class="mt-8 space-y-6">
                <!-- File Summary -->
                <div class="flex items-center justify-between p-4 rounded-xl"
                     :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.03)' }">
                  <div class="flex items-center space-x-3">
                    <img :src="getSummaryIcon()" alt="File type" class="w-8 h-8" />
                    <div>
                      <p class="font-bold" :style="{ color: isDark ? '#ffffff' : '#111827' }">
                        {{ selectedFiles.length }} file{{ selectedFiles.length > 1 ? 's' : '' }}
                      </p>
                      <p class="text-sm opacity-70" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                        {{ formatTotalFileSize() }}
                      </p>
                    </div>
                  </div>
                  <button 
                    @click="clearFile"
                    class="p-2 rounded-lg hover:bg-red-500/10 text-red-500 transition-colors"
                  >
                    <XMarkIcon class="w-5 h-5" />
                  </button>
                </div>

                <!-- Email Input (Optional) -->
                <div>
                  <label class="block text-sm font-bold mb-2 opacity-80"
                         :style="{ color: isDark ? '#d1d5db' : '#374151' }">
                    Send to email (optional)
                  </label>
                  <input 
                    v-model="email"
                    type="email" 
                    placeholder="recipient@example.com"
                    class="w-full px-4 py-3 rounded-xl border-0 transition-all duration-200 focus:ring-2 focus:ring-blue-500"
                    :style="{
                      backgroundColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.05)',
                      color: isDark ? '#f9fafb' : '#111827'
                    }"
                  />
                </div>

                <!-- Upload Button -->
                <button 
                  @click="handleUpload"
                  :disabled="isUploading || selectedFiles.length === 0 || getTotalFileSize() > maxUploadSize"
                  :class="[
                    'w-full py-4 px-6 rounded-2xl font-bold text-lg transition-all duration-300',
                    isUploading || selectedFiles.length === 0 || getTotalFileSize() > maxUploadSize
                      ? 'bg-gray-300 dark:bg-gray-700 text-gray-500 cursor-not-allowed'
                      : 'bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600 text-white hover:scale-105 shadow-lg hover:shadow-xl'
                  ]"
                >
                  <span v-if="isUploading" class="flex items-center justify-center">
                    <svg class="animate-spin -ml-1 mr-3 h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Uploading...
                  </span>
                  <span v-else class="flex items-center justify-center">
                    <CloudArrowUpIcon class="w-6 h-6 mr-3" />
                    Share Files
                  </span>
                </button>

                <!-- Progress Bar -->
                <div v-if="progress > 0 && progress < 100" class="mt-4">
                  <div class="flex justify-between text-sm mb-2" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                    <span>Uploading...</span>
                    <span>{{ progress }}%</span>
                  </div>
                  <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-3 overflow-hidden">
                    <div 
                      class="bg-gradient-to-r from-blue-500 to-purple-500 h-3 rounded-full transition-all duration-300"
                      :style="{ width: progress + '%' }"
                    ></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
          <div v-if="selectedFiles.length > 0" class="space-y-4">
            <div v-for="(file, index) in selectedFiles" :key="index" class="rounded-lg p-4 transition-colors duration-300"
                 :style="{ 
                   backgroundColor: isDark ? '#111111' : '#f9fafb',
                   borderColor: isDark ? '#333333' : '#e5e7eb',
                   borderWidth: '1px'
                 }">
              <div class="flex items-center justify-between mb-3">
                <div class="flex items-center space-x-3">
                  <img :src="getFileIcon(file)" alt="File type" class="w-10 h-10" />
                  <div>
                    <h4 class="font-medium transition-colors duration-300"
                        :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ file.name }}</h4>
                    <p class="text-sm transition-colors duration-300"
                       :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ formatFileSize(file.size) }}</p>
                  </div>
                </div>
                <div class="flex items-center space-x-2">
                  <button 
                    v-if="['mp4', 'pdf', 'jpg', 'jpeg', 'png', 'gif', 'webp', 'txt', 'md', 'json', 'csv', 'xml', 'mp3', 'wav', 'flac'].includes(getFileExtension(file))" 
                    @click="togglePreview(index)"
                    class="px-4 py-2 rounded-xl text-sm font-semibold transition-all duration-300 hover:scale-105 transform active:scale-95"
                    :class="previewingFiles.has(index) 
                      ? 'bg-gradient-to-r from-red-500 to-red-600 hover:from-red-600 hover:to-red-700 text-white shadow-lg' 
                      : 'bg-gradient-to-r from-emerald-500 to-emerald-600 hover:from-emerald-600 hover:to-emerald-700 text-white shadow-lg'"
                    :disabled="isUploading">
                    <EyeIcon v-if="!previewingFiles.has(index)" class="w-4 h-4 mr-1 inline" />
                    <EyeSlashIcon v-else class="w-4 h-4 mr-1 inline" />
                    {{ previewingFiles.has(index) ? 'Hide' : 'Preview' }}
                  </button>
                  <button 
                    @click="removeFile(index)"
                    class="p-1 text-gray-400 dark:text-gray-500 hover:text-red-600 dark:hover:text-red-400 transition-all duration-200"
                  >
                    <XMarkIcon class="w-4 h-4" />
                  </button>
                </div>
              </div>
                 :style="{ 
                   backgroundColor: isDark ? '#111111' : '#f9fafb',
                   borderColor: isDark ? '#333333' : '#e5e7eb',
                   borderWidth: '1px'
                 }">
              <div class="flex items-center justify-between mb-3">
                <div class="flex items-center space-x-3">
                  <img :src="getFileIcon(file)" alt="File type" class="w-10 h-10" />
                  <div>
                    <h4 class="font-medium transition-colors duration-300"
                        :style="{ color: isDark ? '#f9fafb' : '#111827' }">{{ file.name }}</h4>
                    <p class="text-sm transition-colors duration-300"
                       :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ formatFileSize(file.size) }}</p>
                  </div>
                </div>
                <div class="flex items-center space-x-2">
                  <button 
                    v-if="['mp4', 'pdf', 'jpg', 'jpeg', 'png', 'gif', 'webp', 'txt', 'md', 'json', 'csv', 'xml', 'mp3', 'wav', 'flac'].includes(getFileExtension(file))" 
                    @click="togglePreview(index)"
                    class="px-4 py-2 rounded-xl text-sm font-semibold transition-all duration-300 hover:scale-105 transform active:scale-95"
                    :class="previewingFiles.has(index) 
                      ? 'bg-gradient-to-r from-red-500 to-red-600 hover:from-red-600 hover:to-red-700 text-white shadow-lg' 
                      : 'bg-gradient-to-r from-emerald-500 to-emerald-600 hover:from-emerald-600 hover:to-emerald-700 text-white shadow-lg'"
                    :disabled="isUploading">
                    <EyeIcon v-if="!previewingFiles.has(index)" class="w-4 h-4 mr-1 inline" />
                    <EyeSlashIcon v-else class="w-4 h-4 mr-1 inline" />
                    {{ previewingFiles.has(index) ? 'Hide' : 'Preview' }}
                  </button>
                  <button 
                    @click="removeFile(index)"
                    class="p-1 text-gray-400 dark:text-gray-500 hover:text-red-600 dark:hover:text-red-400 transition-all duration-200"
                  >
                    <XMarkIcon class="w-4 h-4" />
                  </button>
                </div>
              </div>

              <!-- Individual Preview -->
              <div v-if="previewingFiles.has(index)" class="mt-3">
                <!-- Image preview for JPG/JPEG/PNG/GIF/WEBP -->
                <div v-if="['jpg', 'jpeg', 'png', 'gif', 'webp'].includes(getFileExtension(file))" class="w-full">
                  <img 
                    :src="previewUrls.get(index)" 
                    :alt="file.name"
                    class="w-full max-w-lg mx-auto rounded-2xl shadow-2xl border"
                    :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }"
                    @error="() => console.error('Image preview failed')">
                </div>

                <!-- Video preview for .mp4 -->
                <div v-else-if="getFileExtension(file) === 'mp4'" class="w-full">
                  <video 
                    :ref="(el) => setVideoRef(index, el)"
                    controls 
                    class="w-full max-w-lg mx-auto rounded-2xl shadow-2xl border" 
                    :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }"
                    :src="previewUrls.get(index)" 
                    @loadedmetadata="updateVideoMetadata">
                    Your browser does not support the video tag.
                  </video>
                </div>

                <!-- PDF preview for .pdf -->
                <div v-else-if="getFileExtension(file) === 'pdf'" class="w-full">
                  <object 
                    :data="previewUrls.get(index)" 
                    type="application/pdf" 
                    class="w-full max-w-lg mx-auto h-96 rounded-2xl shadow-2xl border"
                    :style="{ borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)' }">
                    <div class="flex items-center justify-center h-96 rounded-2xl"
                         :style="{ backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(0, 0, 0, 0.05)' }">
                      <div class="text-center p-8">
                        <div class="w-16 h-16 mx-auto mb-4 rounded-2xl flex items-center justify-center"
                             :style="{ backgroundColor: isDark ? 'rgba(239, 68, 68, 0.2)' : 'rgba(239, 68, 68, 0.1)' }">
                          <svg class="w-8 h-8 text-red-500" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z"/>
                            <path fill-rule="evenodd" d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z" clip-rule="evenodd"/>
                          </svg>
                        </div>
                        <p class="mb-4 font-medium transition-colors duration-300"
                           :style="{ color: isDark ? '#f3f4f6' : '#1f2937' }">PDF preview not supported</p>
                        <p class="text-sm transition-colors duration-300"
                           :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">Download the file to view it</p>
                      </div>
                    </div>
                  </object>
                </div>

                <!-- Text preview for TXT/MD/JSON/CSV/XML -->
                <div v-else-if="['txt', 'md', 'json', 'csv', 'xml'].includes(getFileExtension(file))" class="w-full">
                  <div class="max-w-lg mx-auto rounded-2xl p-6 shadow-2xl border"
                       :style="{ 
                         backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(248, 250, 252, 1)',
                         borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                       }">
                    <pre class="text-sm whitespace-pre-wrap overflow-auto max-h-64 font-mono leading-relaxed"
                         :style="{ color: isDark ? '#e5e7eb' : '#374151' }">{{ textPreviews.get(index) || 'Loading preview...' }}</pre>
                  </div>
                </div>

                <!-- Audio preview for MP3/WAV/FLAC -->
                <div v-else-if="['mp3', 'wav', 'flac'].includes(getFileExtension(file))" class="w-full text-center">
                  <div class="inline-block p-8 rounded-2xl shadow-2xl border"
                       :style="{ 
                         backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(248, 250, 252, 1)',
                         borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                       }">
                    <div class="w-16 h-16 mx-auto mb-4 rounded-2xl flex items-center justify-center"
                         :style="{ backgroundColor: isDark ? 'rgba(34, 197, 94, 0.2)' : 'rgba(34, 197, 94, 0.1)' }">
                      <svg class="w-8 h-8 text-emerald-500" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M9.383 3.076A1 1 0 0110 4v12a1 1 0 01-1.707.707L4.586 13H2a1 1 0 01-1-1V8a1 1 0 011-1h2.586l3.707-3.707a1 1 0 011.09-.217zM15.657 6.343a1 1 0 011.414 0A9.972 9.972 0 0119 12a9.972 9.972 0 01-1.929 5.657 1 1 0 11-1.414-1.414A7.971 7.971 0 0017 12a7.971 7.971 0 00-1.343-4.243 1 1 0 010-1.414z" clip-rule="evenodd"/>
                        <path fill-rule="evenodd" d="M13.828 8.172a1 1 0 011.414 0A5.983 5.983 0 0117 12a5.983 5.983 0 01-1.758 3.828 1 1 0 11-1.414-1.414A3.987 3.987 0 0015 12a3.987 3.987 0 00-1.172-2.828 1 1 0 010-1.414z" clip-rule="evenodd"/>
                      </svg>
                    </div>
                    <audio 
                      controls 
                      class="w-80 max-w-full"
                      :src="previewUrls.get(index)"
                      @error="() => console.error('Audio preview failed')">
                      Your browser does not support the audio tag.
                    </audio>
                  </div>
                </div>
              </div>
            </div>
          </div>

                    <!-- File Details & Controls -->
          <div v-if="selectedFiles.length > 0" class="animate-slide-down transition-colors duration-300"
               :style="{ 
                 borderTopColor: isDark ? '#333333' : '#e5e7eb',
                 borderTopWidth: '1px',
                 backgroundColor: isDark ? '#0f0f0fcc' : '#f9fafbcc'
               }">
            <div class="p-6">
              <!-- Total Files Summary -->
              <div class="flex items-center justify-between mb-6">
                <div class="flex items-center space-x-4">
                  <img :src="getSummaryIcon()" alt="File type" class="w-12 h-12" />
                  <div>
                    <h4 class="font-medium animate-fade-in transition-colors duration-300"
                        :style="{ color: isDark ? '#ffffff' : '#111827' }">{{ selectedFiles.length }} file{{ selectedFiles.length > 1 ? 's' : '' }} selected</h4>
                    <p class="text-sm animate-fade-in-delay transition-colors duration-300"
                       :style="{ color: isDark ? '#9ca3af' : '#6b7280' }">{{ formatTotalFileSize() }}</p>
                  </div>
                </div>
                <button 
                  @click="clearFile"
                  class="p-2 text-gray-400 dark:text-gray-500 hover:text-gray-600 dark:hover:text-gray-300 transition-all duration-200 hover:scale-110"
                >
                  <XMarkIcon class="w-5 h-5" />
                </button>
              </div>

              <!-- Size Warning -->
              <div v-if="getTotalFileSize() > maxUploadSize" class="mb-6 p-4 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg animate-pulse">
                <p class="text-red-800 dark:text-red-400 text-sm font-medium">
                  Total file size exceeds the maximum limit of {{ formatFileSize(maxUploadSize) }}
                </p>
              </div>

              <!-- Email Input -->
              <div class="mb-6">
                <label class="block text-sm font-medium mb-2 transition-colors duration-300"
                       :style="{ color: isDark ? '#d1d5db' : '#374151' }">
                  Send to (optional)
                </label>
                <div class="relative">
                  <EnvelopeIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 transition-colors duration-300"
                                :style="{ color: isDark ? '#6b7280' : '#9ca3af' }" />
                  <input 
                    v-model="email"
                    type="email" 
                    placeholder="recipient@example.com"
                    class="w-full pl-10 pr-4 py-3 rounded-lg focus:ring-2 transition-all duration-200"
                    :style="{
                      borderColor: isDark ? '#4b5563' : '#d1d5db',
                      borderWidth: '1px',
                      backgroundColor: isDark ? '#1f2937' : '#ffffff',
                      color: isDark ? '#f9fafb' : '#111827'
                    }"
                  />
                </div>
              </div>

              <!-- Validity Selection -->
              <div class="mb-6" v-if="validityOptions.length > 0">
                <label class="block text-sm font-medium mb-3 transition-colors duration-300"
                       :style="{ color: isDark ? '#d1d5db' : '#374151' }">
                  File expiration
                </label>
                <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
                  <label
                    v-for="option in validityOptions"
                    :key="option.value"
                    class="relative flex items-center justify-center px-3 py-2 rounded-lg cursor-pointer transition-all text-sm"
                    :style="{
                      borderWidth: '2px',
                      borderColor: selectedValidity === option.value 
                        ? (isDark ? '#60a5fa' : '#3b82f6')
                        : (isDark ? '#4b5563' : '#d1d5db'),
                      backgroundColor: selectedValidity === option.value
                        ? (isDark ? '#1e3a8a20' : '#dbeafe')
                        : (isDark ? '#1f2937' : '#ffffff'),
                      color: selectedValidity === option.value
                        ? (isDark ? '#93c5fd' : '#1d4ed8')
                        : (isDark ? '#d1d5db' : '#374151')
                    }"
                  >
                    <input
                      type="radio"
                      v-model="selectedValidity"
                      :value="option.value"
                      class="sr-only"
                    />
                    <div class="text-center">
                      <div class="font-medium">{{ option.label }}</div>
                      <div class="text-xs opacity-75">{{ option.description }}</div>
                    </div>
                  </label>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="flex space-x-4">
                <button 
                  @click="handleUpload"
                  :disabled="isUploading || selectedFiles.length === 0 || getTotalFileSize() > maxUploadSize"
                  :class="[
                    'flex-1 py-4 px-3 rounded-4xl font-semibold transition-all duration-100',
                    isUploading || selectedFiles.length === 0 || getTotalFileSize() > maxUploadSize
                      ? 'bg-gray-300 dark:bg-gray-600 text-gray-500 dark:text-gray-400 cursor-not-allowed'
                      : 'bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 text-white hover:scale-101 hover:shadow-xl active:scale-95'
                  ]"
                >
                  <span v-if="isUploading" class="flex items-center justify-center">
                    <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Uploading...
                  </span>
                  <span v-else class="flex items-center justify-center">
                    <CloudArrowUpIcon class="w-5 h-5 mr-2" />
                    Transfer Files
                  </span>
                </button>
                
                <button 
                  @click="clearFile"
                  class="px-8 py-4 border-2 border-gray-300 dark:border-gray-600 text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-800 rounded-2xl font-semibold hover:bg-gray-50 dark:hover:bg-gray-700 hover:border-gray-400 dark:hover:border-gray-500 transition-all duration-300 hover:scale-105 transform active:scale-95"
                >
                  <XMarkIcon class="w-5 h-5 mr-2 inline" />
                  Clear All
                </button>
              </div>

              <!-- Progress Bar -->
              <div v-if="progress > 0 && progress < 100" class="mt-6">
                <div class="flex justify-between text-sm text-gray-600 dark:text-gray-400 mb-2">
                  <span>Uploading...</span>
                  <span>{{ progress }}%</span>
                </div>
                <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2 overflow-hidden">
                  <div 
                    class="bg-blue-600 dark:bg-blue-500 h-2 rounded-full transition-all duration-300 animate-pulse"
                    :style="{ width: progress + '%' }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Success Message - WeTransfer Style -->
        <div 
          v-if="message && message.type === 'success'" 
          class="mt-8 bg-white dark:bg-gray-800 rounded-2xl shadow-xl border border-green-200 dark:border-green-800 overflow-hidden animate-scale-in"
        >
          <div class="bg-gradient-to-r from-green-500 to-emerald-600 dark:from-green-600 dark:to-emerald-700 p-6">
            <div class="flex items-center space-x-4">
              <div class="w-12 h-12 bg-white/20 rounded-full flex items-center justify-center">
                <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
              </div>
              <div>
                <h3 class="text-xl font-bold text-white">Transfer complete! 🎉</h3>
                <p class="text-green-100">Your file has been uploaded successfully</p>
              </div>
            </div>
          </div>
          
          <div class="p-6">
            <div class="flex items-center justify-between">
              <div>
                <p class="text-gray-600 dark:text-gray-400 mb-2">Share this link with your recipient:</p>
                <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-3 font-mono text-sm text-gray-800 dark:text-gray-200 break-all">
                  {{ downloadUrl }}
                </div>
              </div>
              <div class="flex space-x-3 ml-6">
                <button 
                  @click="copyToClipboard(downloadUrl)"
                  class="group relative bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 dark:from-blue-600 dark:to-blue-700 dark:hover:from-blue-700 dark:hover:to-blue-800 text-white px-6 py-3 rounded-xl font-semibold transition-all duration-300 hover:scale-105 hover:shadow-lg transform active:scale-95"
                >
                  <span class="flex items-center">
                    <svg class="w-5 h-5 mr-2 transition-transform group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
                    </svg>
                    Copy Link
                  </span>
                  <div class="absolute inset-0 rounded-xl bg-white opacity-0 group-hover:opacity-20 transition-opacity duration-300"></div>
                </button>
                <router-link 
                  v-if="downloadPath"
                  :to="downloadPath"
                  class="group relative bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 dark:from-green-600 dark:to-emerald-700 dark:hover:from-green-700 dark:hover:to-emerald-800 text-white px-6 py-3 rounded-xl font-semibold transition-all duration-300 hover:scale-105 hover:shadow-lg transform active:scale-95"
                >
                  <span class="flex items-center">
                    <svg class="w-5 h-5 mr-2 transition-transform group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                    </svg>
                    View Files
                  </span>
                  <div class="absolute inset-0 rounded-xl bg-white opacity-0 group-hover:opacity-20 transition-opacity duration-300"></div>
                </router-link>
              </div>
            </div>
          </div>
        </div>

        <!-- Error Message -->
        <div 
          v-if="message && message.type === 'error'" 
          class="mt-8 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-2xl p-6 animate-shake"
        >
          <div class="flex items-center space-x-3">
            <div class="w-10 h-10 bg-red-100 dark:bg-red-900/50 rounded-full flex items-center justify-center">
              <XMarkIcon class="w-5 h-5 text-red-600 dark:text-red-400" />
            </div>
            <div>
              <h3 class="font-bold text-red-900 dark:text-red-300">Upload failed</h3>
              <p class="text-red-700 dark:text-red-400">{{ message.text }}</p>
            </div>
          </div>
        </div>
        </div>
      </div>
    </main>
    
    <!-- GitHub Icon - Bottom Left -->
    <a 
      href="https://github.com/R8bert/PInGO-Share" 
      target="_blank" 
      rel="noopener noreferrer" 
      class="fixed bottom-4 left-4 z-40"
      :style="{ color: isDark ? '#9ca3af' : '#6b7280' }"
      title="View GitHub Repository"
    >
      <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
        <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
      </svg>
    </a>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useAuth } from '../../composables/useAuth'
import { useTheme } from '../../composables/useTheme'
import axios from 'axios'
import { CloudArrowUpIcon, EnvelopeIcon, XMarkIcon, EyeIcon, EyeSlashIcon } from '@heroicons/vue/24/outline'

// Import icons
import filePdfIcon from '../../images/icons/files_pdf.svg'
import fileFolderIcon from '../../assets/images/train/icons/file-folder.png'
import pptPptxIcon from '../../images/icons/ppt_pptx_file.svg'
import wordDocIcon from '../../images/icons/word_docx_doc2.svg'
import unknownFileIcon from '../../images/icons/unknown_file_formats.svg'
import totalFilesIcon from '../../images/icons/total_files_icon.svg'

// Use generic file icon for missing specific types
const fileMp4Icon = fileFolderIcon
const filePngIcon = fileFolderIcon  
const fileJpgIcon = fileFolderIcon

interface Message {
  text: string
  type: 'success' | 'error'
}

interface VideoMetadata {
  duration: number
}

// Use auth composable
const { getSettings } = useAuth()
const { isDark } = useTheme()

// State
const selectedFiles = ref<File[]>([])
const fileInput = ref<HTMLInputElement | null>(null)
const email = ref('')
const downloadUrl = ref('')
const downloadPath = ref('')
const isUploading = ref(false)
const progress = ref(0)
const isDragging = ref(false)
const message = ref<Message | null>(null)
const logoPath = ref<string | null>('/logos/004571540fcfd318992384ba96ffb6ae07634920f70e009cf76cf9a1aac603ab.png')
const maxUploadSize = ref<number>(104857600) // Default 100 MB
const previewingFiles = ref<Set<number>>(new Set())
const previewUrls = ref<Map<number, string>>(new Map())
const textPreviews = ref<Map<number, string>>(new Map())
const videoRefs = ref<Map<number, HTMLVideoElement>>(new Map())
const videoMetadata = ref<VideoMetadata>({ duration: 0 })

// Star field data for dark mode animation
interface Star {
  id: number
  x: number
  y: number
  animationDelay: number
  animationDuration: number
  fadeDelay: number
}

const stars = ref<Star[]>([])

// Validity options and selected validity
const validityOptions = ref([
  { value: '7days', label: '7 Days', description: 'One week' },
  { value: '1month', label: '1 Month', description: '30 days' },
  { value: '6months', label: '6 Months', description: 'Half year' },
  { value: '1year', label: '1 Year', description: '12 months' },
  { value: 'never', label: 'Never', description: 'Permanent' }
])
const selectedValidity = ref('7days')
const maxAllowedValidity = ref('7days')

// Icon mapping
const iconMap: Record<string, string> = {
  pdf: filePdfIcon,
  png: filePngIcon,
  jpg: fileJpgIcon,
  jpeg: fileJpgIcon,
  gif: fileJpgIcon, // Use jpg icon for gif
  webp: fileJpgIcon, // Use jpg icon for webp
  docx: wordDocIcon,
  doc: wordDocIcon,
  ppt: pptPptxIcon,
  pptx: pptPptxIcon,
  mp4: fileMp4Icon,
  avi: fileMp4Icon, // Use mp4 icon for avi
  mov: fileMp4Icon, // Use mp4 icon for mov
  txt: wordDocIcon, // Use word icon for text files
  md: wordDocIcon, // Use word icon for markdown
  json: wordDocIcon, // Use word icon for json
  csv: wordDocIcon, // Use word icon for csv
  xml: wordDocIcon, // Use word icon for xml
  mp3: fileMp4Icon, // Use mp4 icon for audio (temporary)
  wav: fileMp4Icon, // Use mp4 icon for audio (temporary)
  flac: fileMp4Icon, // Use mp4 icon for audio (temporary)
}

// Helper functions
const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const getTotalFileSize = () => {
  return selectedFiles.value.reduce((total, file) => total + file.size, 0)
}

const formatTotalFileSize = () => {
  return formatFileSize(getTotalFileSize())
}

const getFileIcon = (file: File) => {
  if (!file) return unknownFileIcon
  const ext = file.name.split('.').pop()?.toLowerCase() || ''
  return iconMap[ext] || unknownFileIcon
}

const getSummaryIcon = () => {
  if (selectedFiles.value.length > 1) {
    return totalFilesIcon
  } else if (selectedFiles.value.length === 1) {
    return getFileIcon(selectedFiles.value[0])
  }
  return unknownFileIcon
}

const getFileExtension = (file: File) => {
  return file?.name.split('.').pop()?.toLowerCase() || ''
}

const createPreviewUrl = (file: File, index: number) => {
  const url = URL.createObjectURL(file)
  previewUrls.value.set(index, url)
  return url
}

const togglePreview = (index: number) => {
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
    // Stop video if playing
    const videoEl = videoRefs.value.get(index)
    if (videoEl) {
      videoEl.pause()
      videoEl.currentTime = 0
    }
    // Clean up preview URL
    const url = previewUrls.value.get(index)
    if (url) {
      URL.revokeObjectURL(url)
      previewUrls.value.delete(index)
    }
    // Clean up text preview
    textPreviews.value.delete(index)
  } else {
    previewingFiles.value.add(index)
    const file = selectedFiles.value[index]
    const fileExtension = getFileExtension(file)
    
    if (['txt', 'md', 'json', 'csv', 'xml'].includes(fileExtension)) {
      // Handle text file preview
      const reader = new FileReader()
      reader.onload = (e) => {
        const content = e.target?.result as string
        // Truncate content if too long (first 1000 characters)
        const truncatedContent = content.length > 1000 
          ? content.substring(0, 1000) + '\n\n... (truncated)'
          : content
        textPreviews.value.set(index, truncatedContent)
      }
      reader.readAsText(file)
    } else {
      // Handle other file types (images, videos, audio)
      createPreviewUrl(file, index)
    }
  }
}

const setVideoRef = (index: number, el: any) => {
  if (el && el instanceof HTMLVideoElement) {
    videoRefs.value.set(index, el)
  }
}

const removeFile = (index: number) => {
  // Clean up preview if active
  if (previewingFiles.value.has(index)) {
    previewingFiles.value.delete(index)
    const url = previewUrls.value.get(index)
    if (url) {
      URL.revokeObjectURL(url)
      previewUrls.value.delete(index)
    }
    textPreviews.value.delete(index)
  }
  
  // Remove the file
  selectedFiles.value.splice(index, 1)
  
  // Reset all preview indices since array shifted
  const newPreviewingFiles = new Set<number>()
  const newPreviewUrls = new Map<number, string>()
  const newTextPreviews = new Map<number, string>()
  const newVideoRefs = new Map<number, HTMLVideoElement>()
  
  previewingFiles.value.forEach(oldIndex => {
    if (oldIndex > index) {
      newPreviewingFiles.add(oldIndex - 1)
      const url = previewUrls.value.get(oldIndex)
      if (url) newPreviewUrls.set(oldIndex - 1, url)
      const textPreview = textPreviews.value.get(oldIndex)
      if (textPreview) newTextPreviews.set(oldIndex - 1, textPreview)
      const videoEl = videoRefs.value.get(oldIndex)
      if (videoEl) newVideoRefs.set(oldIndex - 1, videoEl)
    } else if (oldIndex < index) {
      newPreviewingFiles.add(oldIndex)
      const url = previewUrls.value.get(oldIndex)
      if (url) newPreviewUrls.set(oldIndex, url)
      const textPreview = textPreviews.value.get(oldIndex)
      if (textPreview) newTextPreviews.set(oldIndex, textPreview)
      const videoEl = videoRefs.value.get(oldIndex)
      if (videoEl) newVideoRefs.set(oldIndex, videoEl)
    }
  })
  
  previewingFiles.value = newPreviewingFiles
  previewUrls.value = newPreviewUrls
  textPreviews.value = newTextPreviews
  videoRefs.value = newVideoRefs
}

const updateVideoMetadata = () => {
  // This function can be expanded if needed for multiple video metadata tracking
  console.log('Video metadata updated')
}

const triggerFileInput = () => {
  fileInput.value?.click()
}

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (target.files?.length) {
    selectedFiles.value = Array.from(target.files)
    isDragging.value = false
    message.value = null
    videoMetadata.value.duration = 0
  }
}

const onDrop = (e: DragEvent) => {
  if (e.dataTransfer?.files.length) {
    selectedFiles.value = Array.from(e.dataTransfer.files)
    isDragging.value = false
    message.value = null
    videoMetadata.value.duration = 0
  }
}

const handleUpload = async () => {
  if (selectedFiles.value.length === 0) return

  isUploading.value = true
  progress.value = 0
  message.value = null

  const formData = new FormData()
  selectedFiles.value.forEach((file) => {
    formData.append(`files`, file)
  })
  formData.append('email', email.value)
  formData.append('validity', selectedValidity.value)

  try {
    const res = await axios.post('/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (event) => {
        if (event.total) progress.value = Math.round((event.loaded * 100) / event.total)
      },
    })

    const downloadId = res.data.download_url.split('/').pop()
    downloadUrl.value = `${window.location.origin}/download/${downloadId}`
    downloadPath.value = `/download/${downloadId}`
    message.value = { text: `${selectedFiles.value.length} file${selectedFiles.value.length > 1 ? 's' : ''} uploaded successfully!`, type: 'success' }
  } catch (err: any) {
    console.error('Upload error:', err)
    const errorMessage = err.response?.data?.error || 'Upload failed. Please try again.'
    message.value = { text: errorMessage, type: 'error' }
  } finally {
    isUploading.value = false
  }
}

const clearFile = () => {
  // Clean up all preview URLs
  previewUrls.value.forEach(url => URL.revokeObjectURL(url))
  
  // Stop all videos
  videoRefs.value.forEach(video => {
    video.pause()
    video.currentTime = 0
  })
  
  selectedFiles.value = []
  isDragging.value = false
  message.value = null
  downloadUrl.value = ''
  downloadPath.value = ''
  progress.value = 0
  previewingFiles.value.clear()
  previewUrls.value.clear()
  textPreviews.value.clear()
  videoRefs.value.clear()
  videoMetadata.value.duration = 0
  
  if (fileInput.value) {
    fileInput.value.value = ''
  }
}

const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    message.value = { text: 'Link copied to clipboard!', type: 'success' }
    setTimeout(() => {
      message.value = null
    }, 3000)
  } catch (err) {
    console.error('Failed to copy text: ', err)
    message.value = { text: 'Failed to copy link', type: 'error' }
    setTimeout(() => {
      message.value = null
    }, 3000)
  }
}

const loadSettings = async () => {
  try {
    const settings = await getSettings()
    logoPath.value = settings.logo || null
    maxUploadSize.value = settings.maxUploadSize || 104857600
    maxAllowedValidity.value = settings.maxValidity || '7days'
    
    // Filter validity options based on max allowed
    const validityOrder = ['7days', '1month', '6months', '1year', 'never']
    const maxIndex = validityOrder.indexOf(maxAllowedValidity.value)
    
    if (maxIndex !== -1) {
      const allowedOptions = validityOrder.slice(0, maxIndex + 1)
      validityOptions.value = validityOptions.value.filter(option => 
        allowedOptions.includes(option.value)
      )
    }
    
    // Set default validity to first available option
    if (validityOptions.value.length > 0) {
      selectedValidity.value = validityOptions.value[0].value
    }
  } catch (error) {
    console.error('Error loading settings:', error)
  }
}

// Generate stars for dark mode animation
const generateStars = () => {
  const newStars: Star[] = []
  for (let i = 0; i < 200; i++) {
    newStars.push({
      id: i,
      x: Math.random() * 100,
      y: Math.random() * 100,
      animationDelay: Math.random() * 3,
      animationDuration: 2 + Math.random() * 3,
      fadeDelay: Math.random() * 3
    })
  }
  stars.value = newStars
}

onMounted(() => {
  generateStars()
  loadSettings()
})

// Watch for dark mode changes and regenerate stars
watch(isDark, (newIsDark) => {
  if (newIsDark) {
    generateStars()
  }
})
</script>

<style scoped>
/* Apple-style Star Animations */
.starfield {
  position: fixed;
  inset: 0;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  z-index: -1;
}

.star {
  position: absolute;
  width: 2px;
  height: 2px;
  background: white;
  border-radius: 50%;
  animation: twinkle infinite ease-in-out;
  box-shadow: 0 0 6px rgba(255, 255, 255, 0.8);
}

/* Constellation patterns */
.constellation {
  position: absolute;
  opacity: 0.4;
}

.constellation-1 {
  top: 20%;
  left: 10%;
  width: 120px;
  height: 80px;
  background: 
    radial-gradient(circle 1px at 10px 20px, white 1px, transparent 2px),
    radial-gradient(circle 1px at 50px 10px, white 1px, transparent 2px),
    radial-gradient(circle 1px at 90px 25px, white 1px, transparent 2px),
    radial-gradient(circle 1px at 30px 60px, white 1px, transparent 2px),
    radial-gradient(circle 1px at 70px 50px, white 1px, transparent 2px);
  animation: constellation-float 12s ease-in-out infinite;
}

.constellation-2 {
  top: 60%;
  right: 15%;
  width: 100px;
  height: 100px;
  background: 
    radial-gradient(circle 1px at 20px 20px, #60a5fa 1px, transparent 2px),
    radial-gradient(circle 1px at 60px 30px, #60a5fa 1px, transparent 2px),
    radial-gradient(circle 1px at 40px 70px, #60a5fa 1px, transparent 2px),
    radial-gradient(circle 1px at 80px 60px, #60a5fa 1px, transparent 2px);
  animation: constellation-float 15s ease-in-out infinite reverse;
}

.constellation-3 {
  top: 40%;
  left: 70%;
  width: 80px;
  height: 120px;
  background: 
    radial-gradient(circle 1px at 15px 30px, #a78bfa 1px, transparent 2px),
    radial-gradient(circle 1px at 45px 20px, #a78bfa 1px, transparent 2px),
    radial-gradient(circle 1px at 25px 80px, #a78bfa 1px, transparent 2px),
    radial-gradient(circle 1px at 65px 70px, #a78bfa 1px, transparent 2px);
  animation: constellation-float 18s ease-in-out infinite;
}

/* Shooting stars */
.shooting-stars {
  position: fixed;
  inset: 0;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
  z-index: -1;
}

.shooting-star {
  position: absolute;
  width: 4px;
  height: 4px;
  background: #fff;
  border-radius: 50%;
  box-shadow: 
    0 0 0 4px rgba(255,255,255,0.1),
    0 0 0 8px rgba(255,255,255,0.1),
    0 0 20px rgba(255,255,255,0.1);
  animation: shooting-fall 3s linear infinite;
}

.shooting-star::before {
  content: '';
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 300px;
  height: 1px;
  background: linear-gradient(90deg, #fff, transparent);
}

/* Nebula effects */
.nebula {
  position: absolute;
  border-radius: 50%;
  filter: blur(40px);
  opacity: 0.15;
  animation: nebula-drift infinite ease-in-out;
}

.nebula-1 {
  top: 10%;
  left: 20%;
  width: 300px;
  height: 200px;
  background: radial-gradient(ellipse, #6366f1, transparent 70%);
  animation-duration: 20s;
}

.nebula-2 {
  bottom: 20%;
  right: 30%;
  width: 250px;
  height: 300px;
  background: radial-gradient(ellipse, #8b5cf6, transparent 70%);
  animation-duration: 25s;
  animation-direction: reverse;
}

.nebula-3 {
  top: 50%;
  left: 60%;
  width: 200px;
  height: 250px;
  background: radial-gradient(ellipse, #06b6d4, transparent 70%);
  animation-duration: 30s;
}

/* Star Animation Keyframes */
@keyframes twinkle {
  0%, 100% { 
    opacity: 0.3;
    transform: scale(1);
  }
  50% { 
    opacity: 1;
    transform: scale(1.2);
  }
}

@keyframes shooting-fall {
  0% {
    transform: rotate(315deg) translateX(0);
    opacity: 1;
  }
  70% {
    opacity: 1;
  }
  100% {
    transform: rotate(315deg) translateX(-1000px);
    opacity: 0;
  }
}

@keyframes constellation-float {
  0%, 100% {
    transform: translateY(0) translateX(0);
    opacity: 0.4;
  }
  50% {
    transform: translateY(-10px) translateX(5px);
    opacity: 0.7;
  }
}

@keyframes nebula-drift {
  0%, 100% {
    transform: translateY(0) translateX(0) scale(1);
  }
  33% {
    transform: translateY(-20px) translateX(10px) scale(1.1);
  }
  66% {
    transform: translateY(10px) translateX(-15px) scale(0.9);
  }
}

/* Optimized Animations */
@keyframes fade-in {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes slide-up {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes slide-down {
  from { opacity: 0; transform: translateY(-20px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes scale-in {
  from { opacity: 0; transform: scale(0.8); }
  to { opacity: 1; transform: scale(1); }
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

@keyframes float-delayed {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-15px); }
}

@keyframes float-slow {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-20px) rotate(2deg); }
}

@keyframes float-slower {
  0%, 100% { transform: translateY(0px) rotate(0deg); }
  50% { transform: translateY(-15px) rotate(-1deg); }
}

@keyframes rainbow {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

@keyframes rainbow-circle {
  0% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(255, 0, 0, 0.3) 0%, 
      rgba(255, 64, 0, 0.2) 20%, 
      rgba(255, 128, 0, 0.15) 40%, 
      rgba(255, 192, 0, 0.1) 60%, 
      rgba(255, 255, 0, 0.05) 80%);
    transform: translateY(0px) rotate(0deg) scale(1);
  }
  16.66% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(255, 128, 0, 0.3) 0%, 
      rgba(255, 192, 0, 0.2) 20%, 
      rgba(255, 255, 0, 0.15) 40%, 
      rgba(192, 255, 0, 0.1) 60%, 
      rgba(128, 255, 0, 0.05) 80%);
    transform: translateY(-5px) rotate(60deg) scale(1.05);
  }
  33.33% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(255, 255, 0, 0.3) 0%, 
      rgba(192, 255, 0, 0.2) 20%, 
      rgba(128, 255, 0, 0.15) 40%, 
      rgba(64, 255, 0, 0.1) 60%, 
      rgba(0, 255, 0, 0.05) 80%);
    transform: translateY(-10px) rotate(120deg) scale(1.1);
  }
  50% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(0, 255, 0, 0.3) 0%, 
      rgba(0, 255, 128, 0.2) 20%, 
      rgba(0, 255, 255, 0.15) 40%, 
      rgba(0, 128, 255, 0.1) 60%, 
      rgba(0, 0, 255, 0.05) 80%);
    transform: translateY(-15px) rotate(180deg) scale(1.15);
  }
  66.66% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(0, 0, 255, 0.3) 0%, 
      rgba(64, 0, 255, 0.2) 20%, 
      rgba(128, 0, 255, 0.15) 40%, 
      rgba(192, 0, 255, 0.1) 60%, 
      rgba(255, 0, 255, 0.05) 80%);
    transform: translateY(-10px) rotate(240deg) scale(1.1);
  }
  83.33% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(255, 0, 255, 0.3) 0%, 
      rgba(255, 0, 192, 0.2) 20%, 
      rgba(255, 0, 128, 0.15) 40%, 
      rgba(255, 0, 64, 0.1) 60%, 
      rgba(255, 0, 0, 0.05) 80%);
    transform: translateY(-5px) rotate(300deg) scale(1.05);
  }
  100% { 
    background: radial-gradient(circle at 50% 50%, 
      rgba(255, 0, 0, 0.3) 0%, 
      rgba(255, 64, 0, 0.2) 20%, 
      rgba(255, 128, 0, 0.15) 40%, 
      rgba(255, 192, 0, 0.1) 60%, 
      rgba(255, 255, 0, 0.05) 80%);
    transform: translateY(0px) rotate(360deg) scale(1);
  }
}

@keyframes rainbow-circle-delayed {
  0% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(0, 255, 128, 0.3) 0%, 
      rgba(0, 192, 255, 0.2) 25%, 
      rgba(64, 128, 255, 0.15) 50%, 
      rgba(128, 64, 255, 0.1) 75%, 
      rgba(255, 0, 192, 0.05) 100%);
    transform: translateY(0px) rotate(0deg) scale(1);
  }
  16.66% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(0, 192, 255, 0.3) 0%, 
      rgba(64, 128, 255, 0.2) 25%, 
      rgba(128, 64, 255, 0.15) 50%, 
      rgba(192, 0, 255, 0.1) 75%, 
      rgba(255, 0, 128, 0.05) 100%);
    transform: translateY(-4px) rotate(-60deg) scale(1.03);
  }
  33.33% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(64, 128, 255, 0.3) 0%, 
      rgba(128, 64, 255, 0.2) 25%, 
      rgba(192, 0, 255, 0.15) 50%, 
      rgba(255, 0, 192, 0.1) 75%, 
      rgba(255, 64, 128, 0.05) 100%);
    transform: translateY(-8px) rotate(-120deg) scale(1.08);
  }
  50% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(128, 64, 255, 0.3) 0%, 
      rgba(192, 0, 255, 0.2) 25%, 
      rgba(255, 0, 192, 0.15) 50%, 
      rgba(255, 64, 128, 0.1) 75%, 
      rgba(255, 128, 64, 0.05) 100%);
    transform: translateY(-12px) rotate(-180deg) scale(1.12);
  }
  66.66% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(192, 0, 255, 0.3) 0%, 
      rgba(255, 0, 192, 0.2) 25%, 
      rgba(255, 64, 128, 0.15) 50%, 
      rgba(255, 128, 64, 0.1) 75%, 
      rgba(255, 192, 0, 0.05) 100%);
    transform: translateY(-8px) rotate(-240deg) scale(1.08);
  }
  83.33% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(255, 0, 192, 0.3) 0%, 
      rgba(255, 64, 128, 0.2) 25%, 
      rgba(255, 128, 64, 0.15) 50%, 
      rgba(255, 192, 0, 0.1) 75%, 
      rgba(192, 255, 0, 0.05) 100%);
    transform: translateY(-4px) rotate(-300deg) scale(1.03);
  }
  100% { 
    background: radial-gradient(circle at 30% 70%, 
      rgba(255, 64, 128, 0.3) 0%, 
      rgba(255, 128, 64, 0.2) 25%, 
      rgba(255, 192, 0, 0.15) 50%, 
      rgba(192, 255, 0, 0.1) 75%, 
      rgba(0, 255, 128, 0.05) 100%);
    transform: translateY(0px) rotate(-360deg) scale(1);
  }
}

@keyframes rainbow-circle-fast {
  0% { 
    background: radial-gradient(circle at 70% 30%, 
      rgba(255, 64, 255, 0.25) 0%, 
      rgba(192, 128, 255, 0.2) 20%, 
      rgba(128, 192, 255, 0.15) 40%, 
      rgba(64, 255, 255, 0.1) 60%, 
      rgba(0, 255, 192, 0.08) 80%);
    transform: translateY(0px) rotate(0deg) scale(0.9);
  }
  25% { 
    background: radial-gradient(circle at 70% 30%, 
      rgba(64, 255, 255, 0.25) 0%, 
      rgba(0, 255, 192, 0.2) 20%, 
      rgba(0, 255, 128, 0.15) 40%, 
      rgba(64, 255, 64, 0.1) 60%, 
      rgba(128, 255, 0, 0.08) 80%);
    transform: translateY(-8px) rotate(90deg) scale(1.05);
  }
  50% { 
    background: radial-gradient(circle at 70% 30%, 
      rgba(128, 255, 0, 0.25) 0%, 
      rgba(192, 255, 0, 0.2) 20%, 
      rgba(255, 255, 0, 0.15) 40%, 
      rgba(255, 192, 64, 0.1) 60%, 
      rgba(255, 128, 128, 0.08) 80%);
    transform: translateY(-16px) rotate(180deg) scale(1.2);
  }
  75% { 
    background: radial-gradient(circle at 70% 30%, 
      rgba(255, 128, 128, 0.25) 0%, 
      rgba(255, 64, 192, 0.2) 20%, 
      rgba(255, 0, 255, 0.15) 40%, 
      rgba(192, 64, 255, 0.1) 60%, 
      rgba(128, 128, 255, 0.08) 80%);
    transform: translateY(-8px) rotate(270deg) scale(1.05);
  }
  100% { 
    background: radial-gradient(circle at 70% 30%, 
      rgba(128, 128, 255, 0.25) 0%, 
      rgba(128, 64, 255, 0.2) 20%, 
      rgba(192, 0, 255, 0.15) 40%, 
      rgba(255, 0, 192, 0.1) 60%, 
      rgba(255, 64, 255, 0.08) 80%);
    transform: translateY(0px) rotate(360deg) scale(0.9);
  }
}

/* Animation Classes */
.animate-fade-in {
  animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
  animation: fade-in 0.8s ease-out 0.2s both;
}

.animate-slide-down {
  animation: slide-down 0.4s ease-out;
}

.animate-scale-in {
  animation: scale-in 0.3s ease-out;
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

.animate-float-delayed {
  animation: float-delayed 8s ease-in-out infinite;
}

.animate-float-slow {
  animation: float-slow 8s ease-in-out infinite;
}

.animate-float-slower {
  animation: float-slower 10s ease-in-out infinite;
}

.animate-rainbow {
  background-size: 400% 400%;
  animation: rainbow 3s ease-in-out infinite;
}

.animate-rainbow-circle {
  animation: rainbow-circle 20s ease-in-out infinite;
}

.animate-rainbow-circle-delayed {
  animation: rainbow-circle-delayed 25s ease-in-out infinite;
}

.animate-rainbow-circle-fast {
  animation: rainbow-circle-fast 15s ease-in-out infinite;
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.2, 0, 0.2, 1);
}

/* Enhanced hover effects */
.hover\:scale-105:hover {
  transform: scale(1.05);
}

.hover\:scale-110:hover {
  transform: scale(1.10);
}

/* Progress bar pulse */
.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

/* Custom focus states */
input:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

/* Mobile Optimizations */
@media (max-width: 768px) {
  .text-5xl {
    font-size: 2.5rem;
  }
}

/* Reduced Motion Support */
@media (prefers-reduced-motion: reduce) {
  * {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
</style>
