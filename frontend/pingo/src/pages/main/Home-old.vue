<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 via-white to-indigo-50">
    <!-- Navigation Bar -->
    <nav class="relative z-50 bg-white/80 backdrop-blur-md border-b border-gray-200/50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center space-x-3">
            <img v-if="logoPath" :src="`http://localhost:8080${logoPath}`" class="h-8 w-8" alt="Logo" @error="handleImageError" />
            <CloudArrowUpIcon v-else class="h-8 w-8 text-blue-600" />
            <span class="text-xl font-bold text-gray-900">PinGO</span>
          </div>
          <router-link to="/settings" class="text-gray-600 hover:text-gray-900 transition-colors">
            <CogIcon class="h-5 w-5" />
          </router-link>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="relative">
      <!-- Background Pattern -->
      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <div class="absolute -top-40 -right-40 w-80 h-80 bg-blue-100 rounded-full opacity-30"></div>
        <div class="absolute -bottom-40 -left-40 w-96 h-96 bg-indigo-100 rounded-full opacity-20"></div>
      </div>

      <div class="relative max-w-4xl mx-auto px-4 py-20">
        <!-- Hero Section -->
        <div class="text-center mb-16">
          <h1 class="text-5xl font-bold text-gray-900 mb-6 animate-fade-in">
            Share files <span class="text-blue-600">simply</span>
          </h1>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto animate-fade-in-delay">
            Transfer files of any size securely and quickly. No registration required.
          </p>
        </div>

        <!-- Upload Area -->
        <div class="bg-white rounded-2xl shadow-xl border border-gray-200/50 overflow-hidden animate-slide-up">
          <!-- Upload Zone -->
          <div 
            @dragover.prevent 
            @drop.prevent="onDrop" 
            @dragenter="isDragging = true" 
            @dragleave="isDragging = false"
            @click="triggerFileInput"
            :class="[
              'relative border-2 border-dashed transition-all duration-300 cursor-pointer',
              isDragging ? 'border-blue-400 bg-blue-50' : 'border-gray-300 hover:border-blue-400 hover:bg-gray-50'
            ]"
            class="p-12"
          >
            <input type="file" ref="fileInput" @change="onFileChange" class="hidden" />
            
            <div class="text-center">
              <div class="mb-6">
                <CloudArrowUpIcon 
                  :class="[
                    'w-16 h-16 mx-auto transition-all duration-300',
                    isDragging ? 'text-blue-500 scale-110' : 'text-gray-400'
                  ]" 
                />
              </div>
              
              <h3 class="text-2xl font-semibold text-gray-900 mb-2">
                {{ selectedFile ? 'File selected' : 'Drop your files here' }}
              </h3>
              
              <p class="text-gray-600 mb-6">
                {{ selectedFile ? selectedFile.name : 'or click to browse' }}
              </p>
              
              <button 
                v-if="!selectedFile"
                class="inline-flex items-center px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors duration-200"
              >
                Select files
              </button>
            </div>

            <!-- Drag overlay -->
            <div 
              v-if="isDragging" 
              class="absolute inset-0 bg-blue-100/50 rounded-lg flex items-center justify-center"
            >
              <div class="text-blue-600 font-semibold text-lg">Drop files here</div>
            </div>
          </div>

          <!-- File Details & Controls -->
          <div v-if="selectedFile" class="border-t border-gray-200 bg-gray-50/50">
            <div class="p-6">
              <!-- File Info -->
              <div class="flex items-center justify-between mb-6">
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center">
                    <img :src="getFileIcon()" alt="File type" class="w-8 h-8" />
                  </div>
                  <div>
                    <h4 class="font-medium text-gray-900">{{ selectedFile.name }}</h4>
                    <p class="text-sm text-gray-600">{{ formatFileSize(selectedFile.size) }}</p>
                  </div>
                </div>
                <button 
                  @click="clearFile"
                  class="p-2 text-gray-400 hover:text-gray-600 transition-colors"
                >
                  <XMarkIcon class="w-5 h-5" />
                </button>
              </div>

              <!-- Size Warning -->
              <div v-if="selectedFile.size > maxUploadSize" class="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg">
                <p class="text-red-800 text-sm">
                  File size exceeds the maximum limit of {{ formatFileSize(maxUploadSize) }}
                </p>
              </div>

              <!-- Email Input -->
              <div class="mb-6">
                <label class="block text-sm font-medium text-gray-700 mb-2">
                  Send to (optional)
                </label>
                <div class="relative">
                  <EnvelopeIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
                  <input 
                    v-model="email"
                    type="email" 
                    placeholder="recipient@example.com"
                    class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                  />
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="flex space-x-3">
                <button 
                  @click="handleUpload"
                  :disabled="isUploading || selectedFile.size > maxUploadSize"
                  :class="[
                    'flex-1 py-3 px-6 rounded-lg font-medium transition-all duration-200',
                    isUploading || selectedFile.size > maxUploadSize
                      ? 'bg-gray-300 text-gray-500 cursor-not-allowed'
                      : 'bg-blue-600 hover:bg-blue-700 text-white hover:scale-105'
                  ]"
                >
                  <span v-if="isUploading">Uploading...</span>
                  <span v-else>Transfer</span>
                </button>
                
                <button 
                  @click="clearFile"
                  class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors"
                >
                  Clear
                </button>
              </div>

              <!-- Progress Bar -->
              <div v-if="progress > 0 && progress < 100" class="mt-6">
                <div class="flex justify-between text-sm text-gray-600 mb-2">
                  <span>Uploading...</span>
                  <span>{{ progress }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div 
                    class="bg-blue-600 h-2 rounded-full transition-all duration-300"
                    :style="{ width: progress + '%' }"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Success Message -->
        <div 
          v-if="message && message.type === 'success'" 
          class="mt-8 bg-green-50 border border-green-200 rounded-lg p-6 animate-slide-up"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center space-x-3">
              <div class="w-8 h-8 bg-green-100 rounded-full flex items-center justify-center">
                <svg class="w-5 h-5 text-green-600" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                </svg>
              </div>
              <div>
                <h3 class="font-medium text-green-900">Upload successful!</h3>
                <p class="text-green-700 text-sm">Your file has been uploaded and is ready to share.</p>
              </div>
            </div>
            <a 
              v-if="downloadUrl"
              :href="downloadUrl" 
              target="_blank"
              class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-lg font-medium transition-colors"
            >
              Open link
            </a>
          </div>
        </div>

        <!-- Error Message -->
        <div 
          v-if="message && message.type === 'error'" 
          class="mt-8 bg-red-50 border border-red-200 rounded-lg p-6 animate-slide-up"
        >
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-red-100 rounded-full flex items-center justify-center">
              <XMarkIcon class="w-5 h-5 text-red-600" />
            </div>
            <div>
              <h3 class="font-medium text-red-900">Upload failed</h3>
              <p class="text-red-700 text-sm">{{ message.text }}</p>
            </div>
          </div>
        </div>

        <!-- Features Section -->
        <div class="mt-20 grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="text-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900 mb-2">Secure</h3>
            <p class="text-gray-600 text-sm">Your files are encrypted and automatically deleted after download.</p>
          </div>
          
          <div class="text-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900 mb-2">Fast</h3>
            <p class="text-gray-600 text-sm">Quick uploads and downloads with optimized performance.</p>
          </div>
          
          <div class="text-center">
            <div class="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mx-auto mb-4">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
              </svg>
            </div>
            <h3 class="font-semibold text-gray-900 mb-2">Any file type</h3>
            <p class="text-gray-600 text-sm">Upload documents, images, videos, and more.</p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

          <!-- Preview section -->
          <div v-if="selectedFile" class="space-y-4">
            <button v-if="['mp4', 'pdf'].includes(getFileExtension())" @click="togglePreview"
              class="w-full text-white py-2 rounded-lg text-sm font-semibold hover:opacity-90 transition-all duration-200"
              :class="isPreviewing ? 'bg-red-500 hover:bg-red-600' : 'bg-green-500 hover:bg-green-600'"
              :disabled="isUploading" aria-label="Toggle preview">
              {{ isPreviewing ? 'Hide Preview' : 'Show Preview' }}
            </button>

            <!-- Video preview for .mp4 -->
            <div v-if="isPreviewing && getFileExtension() === 'mp4'" class="w-full">
              <video ref="videoPreview" controls class="w-full rounded-lg shadow-md" :src="previewUrl" @loadedmetadata="updateVideoMetadata">
                Your browser does not support the video tag.
              </video>
            </div>

            <!-- PDF preview for .pdf -->
            <div v-if="isPreviewing && getFileExtension() === 'pdf'" class="w-full">
              <object :data="previewUrl" type="application/pdf" class="w-full h-96 rounded-lg shadow-md">
                <p>Your browser does not support PDF preview. <a :href="previewUrl" download>Download the PDF</a> to view it.</p>
              </object>
            </div>
          </div>

          <!-- Email field -->
          <div class="relative">
            <input v-model="email" type="email" placeholder="Recipient's email (optional)"
              class="w-full border border-gray-300 px-3 py-2 rounded-lg text-sm focus:ring-2 focus:ring-blue-400 focus:border-blue-400 outline-none transition-all duration-200"
              aria-label="Recipient's email" />
            <EnvelopeIcon class="absolute right-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
          </div>

          <!-- Upload and Clear buttons -->
          <div class="flex space-x-2">
            <button :disabled="!selectedFile || isUploading" @click="handleUpload"
              class="flex-1 bg-gradient-to-r from-[var(--button-from)] to-[var(--button-to)] text-white py-2.5 rounded-lg text-sm font-semibold hover:opacity-90 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300 transform hover:-translate-y-0.5"
              aria-label="Upload file">
              {{ isUploading ? 'Uploading...' : 'Share Now' }}
            </button>
            
            <button v-if="selectedFile && !isUploading" @click="clearFile"
              class="p-2.5 bg-red-500 hover:bg-red-600 text-white rounded-lg transition-all duration-300 hover:scale-105"
              aria-label="Clear selected file">
              <XMarkIcon class="w-5 h-5" />
            </button>
          </div>

          <!-- Progress bar -->
          <div v-if="progress > 0 && progress < 100" class="w-full bg-gray-100 rounded-full h-2 overflow-hidden">
            <div class="bg-gradient-to-r from-[var(--button-from)] to-[var(--button-to)] h-full transition-all duration-300"
              :style="{ width: progress + '%' }" />
          </div>

          <!-- Upload Size Slider -->
          <div v-if="selectedFile && maxUploadSize > 0" class="space-y-5">
            <!-- Enhanced header with more sophistication -->
            <div class="flex items-center justify-between p-4 bg-gradient-to-r from-slate-50 to-gray-50 rounded-2xl border border-gray-200/50 shadow-sm">
              <div class="flex items-center space-x-3">
                <div class="relative">
                  <div class="w-3 h-3 rounded-full bg-gradient-to-r from-blue-500 to-purple-500 animate-pulse shadow-lg"></div>
                  <div class="absolute inset-0 w-3 h-3 rounded-full bg-gradient-to-r from-blue-400 to-purple-400 animate-ping opacity-30"></div>
                </div>
                <div>
                  <h3 class="text-sm font-bold text-gray-800 tracking-wide">File Analysis</h3>
                  <p class="text-xs text-gray-500 font-medium">Smart size detection</p>
                </div>
              </div>
              <div class="flex items-center space-x-2">
                <div class="px-3 py-1.5 bg-white rounded-xl shadow-sm border border-gray-200">
                  <span class="text-xs font-mono font-semibold text-gray-700">{{ formatFileSize(selectedFile.size) }}</span>
                </div>
                <div class="px-2 py-1 rounded-lg text-xs font-semibold"
                  :class="{
                    'bg-emerald-100 text-emerald-700': selectedFile.size <= maxUploadSize * 0.7,
                    'bg-amber-100 text-amber-700': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                    'bg-red-100 text-red-700': selectedFile.size > maxUploadSize
                  }"
                >
                  {{ Math.round((selectedFile.size / maxUploadSize) * 100) }}%
                </div>
              </div>
            </div>
            
            <!-- Enhanced progress container with more details -->
            <div class="relative p-6 bg-gradient-to-br from-gray-50 via-white to-gray-50 rounded-3xl border border-gray-200/50 shadow-lg">
              <!-- Enhanced floating label with icon -->
              <div class="absolute -top-3 left-6 px-3 py-1 bg-white rounded-2xl shadow-md border border-gray-200 flex items-center space-x-2">
                <svg class="w-3 h-3 text-gray-500" fill="currentColor" viewBox="0 0 20 20">
                  <path d="M3 4a1 1 0 011-1h12a1 1 0 011 1v2a1 1 0 01-1 1H4a1 1 0 01-1-1V4zM3 10a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H4a1 1 0 01-1-1v-6zM14 9a1 1 0 00-1 1v6a1 1 0 001 1h2a1 1 0 001-1v-6a1 1 0 00-1-1h-2z" />
                </svg>
                <span class="text-xs font-semibold text-gray-600">Usage Monitor</span>
              </div>
              
              <!-- Progress track with enhanced styling -->
              <div class="relative w-full h-3 bg-gradient-to-r from-gray-200 via-gray-100 to-gray-200 rounded-full overflow-hidden shadow-inner border border-gray-300/50 mt-4">
                <!-- Multiple animated layers -->
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/50 to-transparent animate-shimmer-modern"></div>
                <div class="absolute inset-0 bg-gradient-to-r from-transparent via-blue-200/20 to-transparent animate-shimmer-modern" style="animation-delay: 1s;"></div>
                
                <!-- Main progress bar with enhanced gradients -->
                <div
                  class="h-full rounded-full relative overflow-hidden transition-all duration-1000 ease-out shadow-lg border border-white/50"
                  :class="{
                    'bg-gradient-to-r from-emerald-400 via-emerald-500 to-green-500': selectedFile.size <= maxUploadSize * 0.7,
                    'bg-gradient-to-r from-amber-400 via-orange-500 to-red-500': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                    'bg-gradient-to-r from-red-500 via-pink-500 to-red-600 animate-pulse-glow': selectedFile.size > maxUploadSize
                  }"
                  :style="{ 
                    width: `${Math.min((selectedFile.size / maxUploadSize) * 100, 100)}%`,
                    filter: selectedFile.size > maxUploadSize ? 'drop-shadow(0 0 12px rgba(239, 68, 68, 0.7))' : 'drop-shadow(0 3px 6px rgba(0, 0, 0, 0.15))',
                    boxShadow: selectedFile.size > maxUploadSize ? 'inset 0 1px 0 rgba(255,255,255,0.3), 0 0 20px rgba(239, 68, 68, 0.4)' : 'inset 0 1px 0 rgba(255,255,255,0.3)'
                  }"
                >
                  <!-- Enhanced animated highlights -->
                  <div class="absolute inset-0 bg-gradient-to-r from-transparent via-white/40 to-transparent animate-slide-modern"></div>
                  <div class="absolute top-0 left-0 right-0 h-1/2 bg-gradient-to-b from-white/30 to-transparent rounded-t-full"></div>
                  
                  <!-- Enhanced progress indicator dot with ripple effect -->
                  <div class="absolute right-0 top-1/2 transform translate-x-1/2 -translate-y-1/2">
                    <div class="relative">
                      <div class="w-5 h-5 rounded-full bg-white shadow-xl border-3 transition-all duration-300"
                        :class="{
                          'border-emerald-500 shadow-emerald-200': selectedFile.size <= maxUploadSize * 0.7,
                          'border-orange-500 shadow-orange-200': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                          'border-red-500 animate-bounce shadow-red-200': selectedFile.size > maxUploadSize
                        }"
                      ></div>
                      <!-- Ripple effect -->
                      <div class="absolute inset-0 rounded-full animate-ping opacity-40"
                        :class="{
                          'bg-emerald-400': selectedFile.size <= maxUploadSize * 0.7,
                          'bg-orange-400': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                          'bg-red-400': selectedFile.size > maxUploadSize
                        }"
                      ></div>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Enhanced progress details -->
              <div class="flex justify-between items-center mt-4">
                <div class="text-xs text-gray-500 font-medium">0%</div>
                <div class="flex items-center space-x-3">
                  <div class="px-3 py-1.5 bg-white rounded-xl shadow-sm border border-gray-200">
                    <span class="text-xs font-bold"
                      :class="{
                        'text-emerald-600': selectedFile.size <= maxUploadSize * 0.7,
                        'text-orange-600': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                        'text-red-600': selectedFile.size > maxUploadSize
                      }"
                    >
                      {{ Math.min(Math.round((selectedFile.size / maxUploadSize) * 100), 100) }}%
                    </span>
                  </div>
                  <div class="text-xs text-gray-400 font-mono">
                    {{ formatFileSize(maxUploadSize) }} max
                  </div>
                </div>
                <div class="text-xs text-gray-500 font-medium">100%</div>
              </div>
              
              <!-- Speed indicator -->
              <div class="mt-3 flex items-center justify-center">
                <div class="px-4 py-2 bg-gradient-to-r from-blue-50 to-indigo-50 rounded-xl border border-blue-200/50">
                  <div class="flex items-center space-x-2">
                    <div class="w-2 h-2 rounded-full bg-blue-500 animate-pulse"></div>
                    <span class="text-xs font-semibold text-blue-700">
                      Upload speed: ~{{ Math.round(selectedFile.size / 1024 / 1024 * 8) }} Mbps
                    </span>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Enhanced status card with more sophisticated design -->
            <div class="relative overflow-hidden rounded-2xl transition-all duration-500 hover:scale-[1.02] transform"
              :class="{
                'bg-gradient-to-br from-emerald-50 via-green-50 to-emerald-100 border-2 border-emerald-200': selectedFile.size <= maxUploadSize * 0.7,
                'bg-gradient-to-br from-amber-50 via-orange-50 to-yellow-100 border-2 border-orange-200': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                'bg-gradient-to-br from-red-50 via-pink-50 to-red-100 border-2 border-red-200 animate-pulse-border': selectedFile.size > maxUploadSize
              }"
            >
              <!-- Animated background pattern -->
              <div class="absolute inset-0 opacity-30">
                <div class="absolute top-0 right-0 w-32 h-32 rounded-full transform translate-x-16 -translate-y-16"
                  :class="{
                    'bg-gradient-to-br from-emerald-200 to-green-300': selectedFile.size <= maxUploadSize * 0.7,
                    'bg-gradient-to-br from-orange-200 to-amber-300': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                    'bg-gradient-to-br from-red-200 to-pink-300': selectedFile.size > maxUploadSize
                  }"
                ></div>
              </div>
              
              <div class="relative p-5">
                <div class="flex items-center space-x-4">
                  <!-- Enhanced status icon with multiple layers -->
                  <div class="relative flex-shrink-0">
                    <div class="w-12 h-12 rounded-2xl flex items-center justify-center shadow-lg transform transition-transform duration-300 hover:scale-110"
                      :class="{
                        'bg-gradient-to-br from-emerald-400 to-green-600': selectedFile.size <= maxUploadSize * 0.7,
                        'bg-gradient-to-br from-orange-400 to-amber-600': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                        'bg-gradient-to-br from-red-400 to-pink-600': selectedFile.size > maxUploadSize
                      }"
                    >
                      <svg v-if="selectedFile.size <= maxUploadSize * 0.7" class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                      <svg v-else-if="selectedFile.size <= maxUploadSize" class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
                      </svg>
                      <svg v-else class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                      </svg>
                    </div>
                    <!-- Glow effect -->
                    <div class="absolute inset-0 rounded-2xl blur-lg opacity-50"
                      :class="{
                        'bg-emerald-400': selectedFile.size <= maxUploadSize * 0.7,
                        'bg-orange-400': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                        'bg-red-400': selectedFile.size > maxUploadSize
                      }"
                    ></div>
                  </div>
                  
                  <!-- Enhanced status text -->
                  <div class="flex-1">
                    <div class="flex items-center space-x-2 mb-1">
                      <p class="text-base font-bold"
                        :class="{
                          'text-emerald-800': selectedFile.size <= maxUploadSize * 0.7,
                          'text-orange-800': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                          'text-red-800': selectedFile.size > maxUploadSize
                        }"
                      >
                        <span v-if="selectedFile.size <= maxUploadSize * 0.7">✨ Optimal file size</span>
                        <span v-else-if="selectedFile.size <= maxUploadSize">⚠️ Large file detected</span>
                        <span v-else>🚫 File size exceeded</span>
                      </p>
                      <div class="px-2 py-0.5 rounded-full text-xs font-bold"
                        :class="{
                          'bg-emerald-200 text-emerald-800': selectedFile.size <= maxUploadSize * 0.7,
                          'bg-orange-200 text-orange-800': selectedFile.size > maxUploadSize * 0.7 && selectedFile.size <= maxUploadSize,
                          'bg-red-200 text-red-800': selectedFile.size > maxUploadSize
                        }"
                      >
                        {{ selectedFile.size > maxUploadSize ? 'BLOCKED' : 'ALLOWED' }}
                      </div>
                    </div>
                    
                    <div class="space-y-1">
                      <p class="text-sm text-gray-700 font-medium">
                        {{ formatFileSize(Math.abs(maxUploadSize - selectedFile.size)) }} 
                        {{ selectedFile.size > maxUploadSize ? 'over the limit' : 'remaining space' }}
                      </p>
                      <div class="flex items-center space-x-4 text-xs text-gray-600">
                        <span>📁 {{ getFileExtension().toUpperCase() }} format</span>
                        <span>⏱️ Est. {{ Math.ceil(selectedFile.size / 1024 / 1024) }}s upload</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Enhanced File Info section -->
        <div v-if="selectedFile" class="mt-6 sm:mt-0 sm:absolute sm:top-8 sm:right-8 w-full sm:w-auto sm:max-w-sm">
          <!-- Ultra-modern card container with glassmorphism -->
          <div :class="fileDetailsClasses" :style="fileDetailsStyle">
            <!-- Animated background layers -->
            <div class="absolute inset-0 bg-gradient-to-br from-white/10 via-transparent to-white/10"></div>
            <div class="absolute inset-0 bg-gradient-to-tl from-white/5 via-transparent to-white/5"></div>
            
            <!-- Premium glassmorphism header -->
            <div class="relative border-b border-white/20" :style="{
              backgroundColor: `rgba(255, 255, 255, ${(100 - uploadBoxTransparency) / 100 * 0.05})`,
              backdropFilter: blurIntensity > 0 ? `blur(${blurIntensity * 0.5}px)` : 'none',
              '-webkit-backdrop-filter': blurIntensity > 0 ? `blur(${blurIntensity * 0.5}px)` : 'none'
            } as any">
              <div class="flex items-center justify-between gap-8 p-6">
                <div class="flex items-center space-x-3">
                  <div class="relative">
                    <div class="w-4 h-4 rounded-full bg-white/40 animate-pulse shadow-lg"></div>
                    <div class="absolute inset-0 w-4 h-4 rounded-full bg-white/20 animate-ping opacity-30"></div>
                  </div>
                  <div>
                    <h2 class="text-lg font-bold text-white tracking-tight">File Details</h2>
                    <p class="text-xs text-white/70 font-medium">Smart analysis results</p>
                  </div>
                </div>
                
                <button @click="showDetails = !showDetails" 
                  class="group relative p-2.5 rounded-3xl bg-white/10 hover:bg-white/20 backdrop-blur-sm transition-all duration-500 border border-white/30 hover:border-white/40"
                  :aria-label="showDetails ? 'Hide details' : 'Show details'">
                  <svg class="w-5 h-5 text-white transition-all duration-500" 
                    :class="{ 'rotate-180': !showDetails }" 
                    fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 15l7-7 7 7" />
                  </svg>
                  <!-- Button glow effect -->
                  <div class="absolute inset-0 rounded-2xl bg-white/5 opacity-0 group-hover:opacity-100 transition-opacity duration-500 blur-sm"></div>
                </button>
              </div>
              <!-- Subtle animated overlay -->
              <div class="absolute inset-0 bg-gradient-to-br from-white/5 via-transparent to-white/5 pointer-events-none"></div>
            </div>
            
            <!-- Enhanced collapsible content -->
            <div v-show="showDetails" class="relative p-6 space-y-5 animate-slide-down">
              <!-- Premium file type indicator with enhanced animations -->
              <div class="relative">
                <div class="flex items-center justify-center p-6 rounded-3xl border border-white/20 shadow-inner" :style="fileDetailsItemStyle">
                  <div class="text-center">
                    <div class="relative mx-auto mb-4">
                      <!-- Main icon container -->
                      <div class="w-16 h-16 mx-auto rounded-3xl bg-white/20 backdrop-blur-sm flex items-center justify-center shadow-2xl border border-white/30">
                        <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z" clip-rule="evenodd" />
                        </svg>
                      </div>
                      <!-- Floating particles -->
                      <div class="absolute -top-2 -right-2 w-3 h-3 rounded-full bg-yellow-400/70 animate-bounce opacity-70"></div>
                      <div class="absolute -bottom-1 -left-1 w-2 h-2 rounded-full bg-green-400/70 animate-pulse opacity-60"></div>
                      <!-- Glow effect -->
                      <div class="absolute inset-0 rounded-3xl bg-white/10 blur-xl opacity-50"></div>
                    </div>
                    <div class="space-y-2">
                      <p class="text-lg font-bold text-white">.{{ getFileExtension().toUpperCase() }} Document</p>
                      <div class="flex items-center justify-center space-x-2">
                        <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
                        <p class="text-sm font-semibold text-white/80">{{ formatFileSize(selectedFile.size) }}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Enhanced data grid with transparent styling -->
              <div class="space-y-3">
                <!-- File name with transparent styling -->
                <div :class="fileDetailsItemClasses" :style="fileDetailsItemStyle">
                  <div class="relative flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M7 3a1 1 0 000 2h6a1 1 0 100-2H7zM4 7a1 1 0 011-1h10a1 1 0 110 2H5a1 1 0 01-1-1zM2 11a2 2 0 012-2h12a2 2 0 012 2v4a2 2 0 01-2 2H4a2 2 0 01-2-2v-4z" />
                        </svg>
                      </div>
                      <span class="text-sm font-bold text-white uppercase tracking-wider">File Name</span>
                    </div>
                    <div class="flex-1 text-right">
                      <p class="text-sm font-bold text-white truncate max-w-36" :title="selectedFile.name">
                        {{ selectedFile.name }}
                      </p>
                    </div>
                  </div>
                </div>
                
                <!-- File type with transparent design -->
                <div :class="fileDetailsItemClasses" :style="fileDetailsItemStyle">
                  <div class="relative flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M3 4a1 1 0 011-1h4a1 1 0 010 2H6.414l2.293 2.293a1 1 0 01-1.414 1.414L5 6.414V8a1 1 0 01-2 0V4zm9 1a1 1 0 010-2h4a1 1 0 011 1v4a1 1 0 01-2 0V6.414l-2.293 2.293a1 1 0 11-1.414-1.414L13.586 5H12z" clip-rule="evenodd" />
                        </svg>
                      </div>
                      <span class="text-sm font-bold text-white uppercase tracking-wider">MIME Type</span>
                    </div>
                    <p class="text-sm font-bold text-white">{{ selectedFile.type || 'application/octet-stream' }}</p>
                  </div>
                </div>
                
                <!-- Modified date with transparent styling -->
                <div :class="fileDetailsItemClasses" :style="fileDetailsItemStyle">
                  <div class="relative flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd" />
                        </svg>
                      </div>
                      <span class="text-sm font-bold text-white uppercase tracking-wider">Last Modified</span>
                    </div>
                    <p class="text-sm font-bold text-white">{{ formatDate(selectedFile.lastModified) }}</p>
                  </div>
                </div>
                
                <!-- Video duration with transparent styling -->
                <div v-if="videoMetadata.duration" :class="fileDetailsItemClasses" :style="fileDetailsItemStyle">
                  <div class="relative flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z" clip-rule="evenodd" />
                        </svg>
                      </div>
                      <span class="text-sm font-bold text-white uppercase tracking-wider">Duration</span>
                    </div>
                    <p class="text-sm font-bold text-white">{{ formatDuration(videoMetadata.duration) }}</p>
                  </div>
                </div>
                
                <!-- Recipient with transparent design -->
                <div v-if="email" class="relative overflow-hidden p-5 bg-white/10 backdrop-blur-sm rounded-2xl border border-white/20">
                  <div class="relative flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                          <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                        </svg>
                      </div>
                      <span class="text-sm font-bold text-white uppercase tracking-wider">Recipient</span>
                    </div>
                    <p class="text-sm font-bold text-white truncate max-w-32" :title="email">{{ email }}</p>
                  </div>
                </div>
                
                <!-- Security & performance indicator -->
                <div class="relative overflow-hidden p-5 bg-white/10 backdrop-blur-sm rounded-2xl border border-white/20">
                  <div class="relative flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 rounded-xl bg-white/20 backdrop-blur-sm flex items-center justify-center border border-white/30">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                        </svg>
                      </div>
                      <span class="text-sm font-bold text-white uppercase tracking-wider">Security</span>
                    </div>
                    <div class="flex items-center space-x-2">
                      <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
                      <p class="text-sm font-bold text-white">Encrypted</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Ultra-premium success/error message -->
            <div v-if="message" class="p-6 pt-0">
              <div class="relative overflow-hidden rounded-3xl p-6 animate-slide-up"
                :class="{
                  'bg-gradient-to-br from-emerald-50 via-green-100 to-emerald-200 border-2 border-emerald-300/50': message.type === 'success',
                  'bg-gradient-to-br from-red-50 via-pink-100 to-red-200 border-2 border-red-300/50': message.type === 'error'
                }"
              >
                <!-- Enhanced animated background pattern -->
                <div class="absolute inset-0 opacity-30">
                  <div class="w-full h-full bg-gradient-to-br from-white/60 via-transparent to-white/30 animate-shimmer-slow"></div>
                </div>
                
                <!-- Floating decorative elements -->
                <div class="absolute top-2 right-2 w-4 h-4 rounded-full opacity-20 animate-bounce"
                  :class="{
                    'bg-emerald-400': message.type === 'success',
                    'bg-red-400': message.type === 'error'
                  }"
                  :style="message.type === 'success' ? `background-color: var(--primary-from)` : ''"
                ></div>
                <div class="absolute bottom-2 left-2 w-3 h-3 rounded-full opacity-15 animate-pulse"
                  :class="{
                    'bg-green-400': message.type === 'success',
                    'bg-pink-400': message.type === 'error'
                  }"
                  :style="message.type === 'success' ? `background-color: var(--secondary-from)` : ''"
                ></div>
                
                <div class="relative text-center">
                  <!-- Enhanced status icon with glow -->
                  <div class="relative mx-auto mb-4">
                    <div class="w-14 h-14 mx-auto rounded-2xl flex items-center justify-center shadow-2xl"
                      :class="{
                        'bg-gradient-to-br from-emerald-400 via-green-500 to-emerald-600': message.type === 'success',
                        'bg-gradient-to-br from-red-400 via-pink-500 to-red-600': message.type === 'error'
                      }"
                      :style="message.type === 'success' ? `background: linear-gradient(to br, var(--primary-from), var(--secondary-from), var(--accent-from))` : ''"
                    >
                      <svg v-if="message.type === 'success'" class="w-7 h-7 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                      <svg v-else class="w-7 h-7 text-white" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                      </svg>
                    </div>
                    <!-- Glow effect -->
                    <div class="absolute inset-0 rounded-2xl blur-lg opacity-60"
                      :class="{
                        'bg-emerald-400': message.type === 'success',
                        'bg-red-400': message.type === 'error'
                      }"
                      :style="message.type === 'success' ? `background-color: var(--primary-from)` : ''"
                    ></div>
                  </div>
                  
                  <p class="font-bold text-lg mb-3"
                    :class="{
                      'text-emerald-800': message.type === 'success',
                      'text-red-800': message.type === 'error'
                    }"
                  >
                    {{ message.text }}
                  </p>
                  
                  <a v-if="message.type === 'success' && downloadUrl" :href="downloadUrl" target="_blank"
                    class="inline-flex items-center px-6 py-3 bg-white/90 hover:bg-white rounded-2xl text-blue-600 hover:text-blue-700 text-sm font-bold transition-all duration-500 hover:scale-110 shadow-lg border border-white/50 hover:shadow-xl group"
                  >
                    <svg class="w-5 h-5 mr-2 transition-transform duration-300 group-hover:scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                    Open Download Link
                  </a>
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
import { ref, onMounted, onUnmounted, watch, computed } from 'vue'
import axios from 'axios'
import { CloudArrowUpIcon, EnvelopeIcon, XMarkIcon } from '@heroicons/vue/24/outline'

// Import icons
import fileMp4Icon from './images/train/icons/file-mp4.png'
import filePdfIcon from './images/train/icons/file-pdf.png'
import fileFolderIcon from './images/train/icons/file-folder.png'
import filePngIcon from './images/train/icons/file-png.png'
import fileJpgIcon from './images/train/icons/file-jpg.png'
import fileDocxIcon from './images/train/icons/file-docx.png'

interface Message {
  text: string
  type: 'success' | 'error'
}

interface VideoMetadata {
  duration: number
}

const selectedFile = ref<File | null>(null)
const fileInput = ref<HTMLInputElement | null>(null)
const email = ref('')
const downloadUrl = ref('')
const isUploading = ref(false)
const progress = ref(0)
const isDragging = ref(false)
const message = ref<Message | null>(null)
const useFolderIcon = ref(false)
const isPreviewing = ref(false)
const previewUrl = ref<string | undefined>(undefined)
const videoPreview = ref<HTMLVideoElement | null>(null)
const backgroundImage = ref<HTMLImageElement | null>(null)
const buttonFromColor = ref<string>('')
const buttonToColor = ref<string>('#6366f1')
const videoMetadata = ref<VideoMetadata>({ duration: 0 })
const logoPath = ref<string | null>(null)
const backgroundPath = ref<string | null>(null)
const maxUploadSize = ref<number>(104857600) // Default 100 MB
const uploadBoxTransparency = ref<number>(80) // Default 80% transparent
const blurIntensity = ref<number>(12) // Default medium blur
const showDetails = ref<boolean>(true)

// Computed properties for dynamic styling
const uploadBoxClasses = computed(() => {
  return `max-w-md w-full shadow-xl rounded-xl p-6 space-y-6 animate-fade-in relative overflow-hidden ml-0 mt-4`
})

const uploadBoxStyle = computed(() => {
  // 0% transparency = fully opaque (opacity 1)
  // 100% transparency = fully transparent (opacity 0)
  const opacity = (100 - uploadBoxTransparency.value) / 100
  const blurValue = blurIntensity.value
  
  return {
    backgroundColor: `rgba(255, 255, 255, ${opacity})`,
    backdropFilter: blurValue > 0 ? `blur(${blurValue}px)` : 'none',
    '-webkit-backdrop-filter': blurValue > 0 ? `blur(${blurValue}px)` : 'none' // Safari support
  } as any
})

// File details styling
const fileDetailsClasses = computed(() => {
  return `relative rounded-3xl shadow-2xl border border-white/30 overflow-hidden`
})

const fileDetailsStyle = computed(() => {
  const opacity = (100 - uploadBoxTransparency.value) / 100 // Use same transparency as upload box
  const blurValue = blurIntensity.value
  return {
    backgroundColor: `rgba(255, 255, 255, ${opacity * 0.05})`,
    backdropFilter: blurValue > 0 ? `blur(${blurValue}px)` : 'none',
    '-webkit-backdrop-filter': blurValue > 0 ? `blur(${blurValue}px)` : 'none'
  } as any
})

const fileDetailsItemClasses = computed(() => {
  return `relative overflow-hidden p-5 rounded-2xl border border-white/20`
})

const fileDetailsItemStyle = computed(() => {
  const opacity = (100 - uploadBoxTransparency.value) / 100
  const blurValue = blurIntensity.value * 0.5 // Use half the blur for items
  return {
    backgroundColor: `rgba(255, 255, 255, ${opacity * 0.1})`,
    backdropFilter: blurValue > 0 ? `blur(${blurValue}px)` : 'none',
    '-webkit-backdrop-filter': blurValue > 0 ? `blur(${blurValue}px)` : 'none'
  } as any
})

// Computed property for text contrast based on transparency
const textContrastStyle = computed(() => {
  const opacity = (100 - uploadBoxTransparency.value) / 100
  // When background is very transparent (low opacity), we need stronger text shadows
  const shadowStrength = opacity < 0.3 ? 0.9 : 0.5
  return {
    textShadow: `2px 2px 4px rgba(255,255,255,${shadowStrength}), -1px -1px 2px rgba(0,0,0,${shadowStrength * 0.6})`
  }
})

const iconMap: Record<string, string> = {
  pdf: filePdfIcon,
  png: filePngIcon,
  jpg: fileJpgIcon,
  jpeg: fileJpgIcon,
  docx: fileDocxIcon,
  mp4: fileMp4Icon,
}

const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 Bytes'
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

const formatDuration = (seconds: number) => {
  const minutes = Math.floor(seconds / 60)
  const remainingSeconds = Math.floor(seconds % 60)
  return `${minutes}:${remainingSeconds < 10 ? '0' : ''}${remainingSeconds}`
}

const formatDate = (timestamp: number) => {
  const date = new Date(timestamp)
  return date.toLocaleDateString() + ' ' + date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

const triggerFileInput = () => {
  fileInput.value?.click()
}

const onFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (target.files?.length) {
    selectedFile.value = target.files[0]
    useFolderIcon.value = false
    isDragging.value = false
    message.value = null
    updatePreviewUrl()
    videoMetadata.value.duration = 0
  }
}

const onDrop = (e: DragEvent) => {
  if (e.dataTransfer?.files.length) {
    selectedFile.value = e.dataTransfer.files[0]
    useFolderIcon.value = false
    isDragging.value = false
    message.value = null
    updatePreviewUrl()
    videoMetadata.value.duration = 0
  }
}

const handleUpload = async () => {
  if (!selectedFile.value) return

  isUploading.value = true
  progress.value = 0
  message.value = null

  const formData = new FormData()
  formData.append('file', selectedFile.value)
  formData.append('email', email.value)

  try {
    const res = await axios.post('http://localhost:8080/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (event) => {
        if (event.total) progress.value = Math.round((event.loaded * 100) / event.total)
      },
    })

    downloadUrl.value = `http://localhost:8080${res.data.download_url}`
    message.value = { text: 'File uploaded successfully!', type: 'success' }
  } catch (err) {
    console.error('Upload error:', err)
    message.value = { text: 'Upload failed. Please try again.', type: 'error' }
  } finally {
    isUploading.value = false
  }
}

const getFileIcon = () => {
  if (!selectedFile.value) return ''
  const ext = selectedFile.value.name.split('.').pop()?.toLowerCase() || ''
  return useFolderIcon.value ? fileFolderIcon : (iconMap[ext] || fileFolderIcon)
}

const getFileExtension = () => {
  return selectedFile.value?.name.split('.').pop()?.toLowerCase() || ''
}

const updatePreviewUrl = () => {
  if (selectedFile.value) previewUrl.value = URL.createObjectURL(selectedFile.value)
  else previewUrl.value = undefined
}

const togglePreview = () => {
  isPreviewing.value = !isPreviewing.value
  if (isPreviewing.value && videoPreview.value) videoPreview.value.play().catch(err => console.error('Autoplay blocked:', err))
  else if (!isPreviewing.value && videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
}

const updateVideoMetadata = () => {
  if (videoPreview.value && getFileExtension() === 'mp4') videoMetadata.value.duration = videoPreview.value.duration
}

const clearFile = () => {
  selectedFile.value = null
  useFolderIcon.value = false
  isDragging.value = false
  message.value = null
  downloadUrl.value = ''
  progress.value = 0
  isPreviewing.value = false
  videoMetadata.value.duration = 0
  
  // Reset file input
  if (fileInput.value) {
    fileInput.value.value = ''
  }
  
  // Clean up preview URL
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    previewUrl.value = undefined
  }
  
  // Stop video if playing
  if (videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
}

const handleImageError = () => {
  console.error('Image failed to load:', backgroundPath.value || logoPath.value)
  if (!backgroundPath.value) backgroundPath.value = 'https://4kwallpapers.com/images/wallpapers/jungle-tree-dark-3840x2160-22695.jpg'
  if (!logoPath.value) logoPath.value = null
}

const onBackgroundImageLoad = () => {
  console.log('Background image loaded, extracting colors...')
  // Small delay to ensure the image is properly rendered in the DOM
  setTimeout(() => {
    extractBackgroundColors()
  }, 50)
}

// Cleanup
watch(selectedFile, () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    updatePreviewUrl()
  }
  if (!isPreviewing.value && videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
  videoMetadata.value.duration = 0
})

// Watch for background changes and re-extract colors
watch(backgroundPath, async (newPath) => {
  if (newPath && backgroundImage.value) {
    console.log('Background path changed to:', newPath)
    // Wait a bit for the image to load
    setTimeout(() => {
      extractBackgroundColors()
    }, 300)
  }
}, { immediate: true })

onUnmounted(() => {
  if (previewUrl.value) URL.revokeObjectURL(previewUrl.value)
  if (videoPreview.value) {
    videoPreview.value.pause()
    videoPreview.value.currentTime = 0
  }
})

onMounted(async () => {
  await loadSettings()
  // Color extraction will happen automatically when the background image loads via @load event
  // But also try to extract colors after a short delay in case the image is already loaded
  setTimeout(() => {
    if (backgroundImage.value && backgroundPath.value) {
      console.log('Attempting color extraction from onMounted...')
      extractBackgroundColors()
    }
  }, 500)
})

// Watch for changes in settings
watch(uploadBoxTransparency, (newVal) => {
  console.log('Transparency changed to:', newVal)
})

watch(blurIntensity, (newVal) => {
  console.log('Blur intensity changed to:', newVal)
})

const loadSettings = async () => {
  try {
    const response = await axios.get('http://localhost:8080/settings')
    console.log('Loaded settings:', response.data)
    logoPath.value = response.data.logo || null
    backgroundPath.value = response.data.backgroundImage || null
    maxUploadSize.value = response.data.maxUploadSize || 104857600 // Default 100 MB
    
    // Handle transparency and blur with proper fallbacks
    uploadBoxTransparency.value = response.data.uploadBoxTransparency !== undefined ? response.data.uploadBoxTransparency : 80 // Default 80%
    blurIntensity.value = response.data.blurIntensity !== undefined ? response.data.blurIntensity : 12 // Default medium blur
    
    console.log('Final transparency:', uploadBoxTransparency.value, 'blur:', blurIntensity.value)
  } catch (error) {
    console.error('Error loading settings:', error)
  }
}

const extractBackgroundColors = () => {
  if (!backgroundImage.value || !backgroundPath.value) {
    console.log('No background image available for color extraction')
    return
  }

  console.log('Extracting colors from background image...')
  
  try {
    // Ensure the image is loaded and has dimensions
    if (!backgroundImage.value.complete || backgroundImage.value.naturalWidth === 0) {
      console.log('Background image not fully loaded yet')
      return
    }
    
    // Get dominant color for button gradients
    const canvas = document.createElement('canvas')
    const ctx = canvas.getContext('2d')
    if (!ctx) return

    canvas.width = backgroundImage.value.width
    canvas.height = backgroundImage.value.height
    ctx.drawImage(backgroundImage.value, 0, 0)

    const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height)
    const data = imageData.data
    let r = 0, g = 0, b = 0, pixelCount = 0

    for (let i = 0; i < data.length; i += 4) {
      r += data[i]
      g += data[i + 1]
      b += data[i + 2]
      pixelCount++
    }

    r = Math.round(r / pixelCount)
    g = Math.round(g / pixelCount)
    b = Math.round(b / pixelCount)
    
    console.log('Dominant color extracted:', `rgb(${r}, ${g}, ${b})`)
    
    const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255
    const isDark = luminance < 0.5

    const hsl = rgbToHsl(r, g, b)
    let complementaryHue = (hsl.h + 180) % 360
    const saturation = hsl.s
    const lightness = isDark ? Math.min(0.7, hsl.l + 0.2) : Math.max(0.3, hsl.l - 0.2)

    const fromColor = hslToRgb(complementaryHue, saturation, lightness)
    const toColor = hslToRgb((complementaryHue + 30) % 360, saturation, lightness)
    buttonFromColor.value = `rgb(${fromColor.r}, ${fromColor.g}, ${fromColor.b})`
    buttonToColor.value = `rgb(${toColor.r}, ${toColor.g}, ${toColor.b})`
    
    console.log('Color extraction completed successfully!')
    
  } catch (error) {
    console.error('Error extracting colors from background:', error)
  }
}

function rgbToHsl(r: number, g: number, b: number) {
  r /= 255; g /= 255; b /= 255
  const max = Math.max(r, g, b), min = Math.min(r, g, b)
  let h = 0, s = 0, l = (max + min) / 2

  if (max === min) h = s = 0
  else {
    const d = max - min
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min)
    switch (max) {
      case r: h = (g - b) / d + (g < b ? 6 : 0); break
      case g: h = (b - r) / d + 2; break
      case b: h = (r - g) / d + 4; break
    }
    h /= 6
  }
  return { h: h * 360, s, l }
}

function hslToRgb(h: number, s: number, l: number) {
  h /= 360; s /= 100; l /= 100
  let r, g, b
  if (s === 0) r = g = b = l
  else {
    const hue2rgb = (p: number, q: number, t: number) => {
      if (t < 0) t += 1
      if (t > 1) t -= 1
      if (t < 1/6) return p + (q - p) * 6 * t
      if (t < 1/2) return q
      if (t < 2/3) return p + (q - p) * (2/3 - t) * 6
      return p
    }
    const q = l < 0.5 ? l * (1 + s) : l + s - l * s
    const p = 2 * l - q
    r = hue2rgb(p, q, h + 1/3)
    g = hue2rgb(p, q, h)
    b = hue2rgb(p, q, h - 1/3)
  }
  return { r: Math.round(r * 255), g: Math.round(g * 255), b: Math.round(b * 255) }
}
</script>

<style scoped>
/* Fade in and slide up animations */
.animate-fade-in { animation: fadeIn 0.6s ease-in-out; }
@keyframes fadeIn { from { opacity: 0; transform: translateY(15px); } to { opacity: 1; transform: translateY(0); } }
.animate-slide-up { animation: slideUp 0.4s ease-in-out; }
@keyframes slideUp { from { opacity: 0; transform: translateY(15px); } to { opacity: 1; transform: translateY(0); } }
.animate-slide-down { animation: slideDown 0.3s ease-in-out; }
@keyframes slideDown { from { opacity: 0; transform: translateY(-10px); } to { opacity: 1; transform: translateY(0); } }

/* Continuous diagonal scroll effect */
@keyframes diagonal-scroll { 0% { background-position: 0 0; } 100% { background-position: -240px 240px; } }
.animate-diagonal-scroll { animation: diagonal-scroll 10s linear infinite; }

/* Enhanced Upload size slider animations */
@keyframes shimmer-modern {
  0% { transform: translateX(-100%); opacity: 0; }
  50% { opacity: 1; }
  100% { transform: translateX(100%); opacity: 0; }
}
.animate-shimmer-modern { animation: shimmer-modern 3s ease-in-out infinite; }

@keyframes slide-modern {
  0% { transform: translateX(-100%) skewX(-15deg); opacity: 0; }
  50% { opacity: 1; }
  100% { transform: translateX(150%) skewX(-15deg); opacity: 0; }
}
.animate-slide-modern { animation: slide-modern 2.5s ease-in-out infinite; }

@keyframes pulse-glow {
  0%, 100% { box-shadow: 0 0 5px rgba(239, 68, 68, 0.3); }
  50% { box-shadow: 0 0 20px rgba(239, 68, 68, 0.6), 0 0 30px rgba(239, 68, 68, 0.4); }
}
.animate-pulse-glow { animation: pulse-glow 2s ease-in-out infinite; }

@keyframes pulse-border {
  0%, 100% { border-color: rgb(254 226 226); }
  50% { border-color: rgb(248 113 113); }
}
.animate-pulse-border { animation: pulse-border 2s ease-in-out infinite; }

/* Enhanced glassmorphism animations */
@keyframes gradient-shift {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}
.animate-gradient-shift { animation: gradient-shift 8s ease-in-out infinite; background-size: 200% 200%; }

@keyframes shimmer-slow {
  0% { transform: translateX(-100%) rotate(-45deg); opacity: 0; }
  50% { opacity: 0.6; }
  100% { transform: translateX(300%) rotate(-45deg); opacity: 0; }
}
.animate-shimmer-slow { animation: shimmer-slow 6s ease-in-out infinite; }

/* Enhanced shadow effects */
.shadow-3xl {
  box-shadow: 0 35px 60px -12px rgba(0, 0, 0, 0.25), 0 0 0 1px rgba(255, 255, 255, 0.1);
}

/* Enhanced border utilities */
.border-3 { border-width: 3px; }

/* Ensure overlay covers the entire dialog */
.absolute.inset-0 { top: 0; left: 0; right: 0; bottom: 0; }

/* Styles for icons and overlay */
.bg-repeat { background-repeat: repeat; }

/* Extended container to cover corners when rotated */
.absolute.inset-100p { top: -100%; left: -100%; right: -100%; bottom: -100%; }

/* Responsive adjustments */
@media (max-width: 640px) {
  .sm\:absolute { position: relative; top: auto; right: auto; width: 100%; margin-top: 1.5rem; }
  .max-w-md { max-width: 100%; }
}

/* Dynamic button styles */
button {
  background: var(--button-from);
  background: linear-gradient(to right, var(--button-from), var(--button-to));
}
button:hover { opacity: 0.9; }

/* Enhanced hover effects */
.group:hover .group-hover\:scale-110 { transform: scale(1.1); }
.group:hover .group-hover\:rotate-3 { transform: rotate(3deg); }
</style>