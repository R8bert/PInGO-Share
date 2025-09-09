<template>
    <div
        class="min-h-screen transition-colors duration-300"
        :style="{ backgroundColor: isDark ? '#0f0f23' : '#fafafa' }"
    >
        <!-- Hero Section with Technical Overview -->
        <section class="relative min-h-screen flex items-center justify-center px-4 sm:px-6 lg:px-8">
            <!-- Background Pattern -->
            <div class="absolute inset-0 opacity-30">
                <div 
                    class="absolute inset-0"
                    :style="{
                        backgroundImage: `radial-gradient(circle at 1px 1px, ${isDark ? 'rgba(255,255,255,0.15)' : 'rgba(0,0,0,0.1)'} 1px, transparent 0)`,
                        backgroundSize: '20px 20px'
                    }"
                ></div>
            </div>
            
            <div class="max-w-6xl mx-auto w-full relative">
                <!-- Header Badge -->
                <div class="text-center mb-8">
                    <div 
                        class="inline-flex items-center px-4 py-2 rounded-full text-sm font-medium mb-4"
                        :style="{
                            backgroundColor: isDark ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.1)',
                            color: isDark ? '#60a5fa' : '#2563eb',
                            border: `1px solid ${isDark ? 'rgba(59, 130, 246, 0.2)' : 'rgba(59, 130, 246, 0.2)'}`
                        }"
                    >
                        <svg class="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"/>
                        </svg>
                        Open Source • Self-Hosted • Production Ready
                    </div>
                </div>

                <!-- Main Title and Description -->
                <div class="text-center space-y-6">
                    <h1 
                        class="text-5xl sm:text-6xl lg:text-7xl font-bold tracking-tight leading-tight"
                        :style="{ color: isDark ? '#ffffff' : '#111827' }"
                    >
                        Enterprise-Grade
                        <br>
                        <span 
                            class="bg-gradient-to-r from-blue-600 via-purple-600 to-blue-600 bg-clip-text text-transparent animate-pulse"
                        >
                            File Transfer
                        </span>
                    </h1>
                    
                    <p 
                        class="text-xl sm:text-2xl max-w-4xl mx-auto leading-relaxed font-medium"
                        :style="{ color: isDark ? '#d1d5db' : '#4b5563' }"
                    >
                        A modern, scalable, and secure file sharing platform built with 
                        cutting-edge technologies. Deploy anywhere, scale infinitely.
                    </p>
                </div>

                <!-- Technical Highlights -->
                <div class="mt-12 grid grid-cols-1 md:grid-cols-3 gap-6">
                    <div 
                        class="text-center p-6 rounded-xl border backdrop-blur-sm"
                        :style="{
                            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(255, 255, 255, 0.8)',
                            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                        }"
                    >
                        <div class="text-3xl font-mono font-bold text-blue-500 mb-2">99.9%</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }" class="text-sm font-medium">Uptime SLA</div>
                    </div>
                    
                    <div 
                        class="text-center p-6 rounded-xl border backdrop-blur-sm"
                        :style="{
                            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(255, 255, 255, 0.8)',
                            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                        }"
                    >
                        <div class="text-3xl font-mono font-bold text-green-500 mb-2">10GB/s</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }" class="text-sm font-medium">Transfer Speed</div>
                    </div>
                    
                    <div 
                        class="text-center p-6 rounded-xl border backdrop-blur-sm"
                        :style="{
                            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.05)' : 'rgba(255, 255, 255, 0.8)',
                            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                        }"
                    >
                        <div class="text-3xl font-mono font-bold text-purple-500 mb-2">256-bit</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }" class="text-sm font-medium">AES Encryption</div>
                    </div>
                </div>

                <!-- CTA Buttons -->
                <div class="flex flex-col sm:flex-row items-center justify-center gap-4 mt-12">
                    <button
                        @click="scrollToUpload"
                        class="px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-semibold rounded-xl transition-all duration-300 hover:scale-105 shadow-lg hover:shadow-xl"
                    >
                        Start File Transfer
                    </button>
                    <button
                        @click="scrollToArchitecture"
                        class="px-8 py-4 border-2 font-semibold rounded-xl transition-all duration-300 hover:scale-105 backdrop-blur-sm"
                        :style="{
                            borderColor: isDark ? 'rgba(255,255,255,0.2)' : 'rgba(0,0,0,0.2)',
                            color: isDark ? '#ffffff' : '#000000',
                        }"
                    >
                        View Architecture
                    </button>
                </div>
            </div>
        </section>
                                            <span
                                                v-for="(char, charIndex) in word.split('')"
                                                :key="charIndex"
                                                :class="[
                                                    'inline-block transition-all duration-300 ease-out relative',
                                                    currentWordIndex === wordIndex && charIndex < displayedChars
                                                        ? 'opacity-100 blur-0 transform translate-y-0'
                                                        : 'opacity-0 blur-sm transform translate-y-2'
                                                ]"
                                                :style="{
                                                    transitionDelay: `${charIndex * 50}ms`
                                                }"
                                            >
                                                {{ char }}
                                            </span>
                                            <!-- Clean progressive highlight -->
                                            <span 
                                                v-if="currentWordIndex === wordIndex && isHighlighted"
                                                :class="[
                                                    'absolute inset-0 -inset-x-1 -inset-y-1 rounded-md',
                                                    'clean-highlight-sweep'
                                                ]"
        <!-- Architecture Overview Section -->
        <section ref="architectureSection" class="py-20 px-4 sm:px-6 lg:px-8">
            <div class="max-w-7xl mx-auto">
                <div class="text-center mb-16">
                    <h2 
                        class="text-4xl sm:text-5xl font-bold mb-6"
                        :style="{ color: isDark ? '#ffffff' : '#111827' }"
                    >
                        System Architecture
                    </h2>
                    <p 
                        class="text-xl max-w-3xl mx-auto"
                        :style="{ color: isDark ? '#d1d5db' : '#4b5563' }"
                    >
                        Built with modern microservices architecture, containerized deployment, 
                        and cloud-native principles for maximum scalability and reliability.
                    </p>
                </div>

                <!-- Architecture Diagram -->
                <div 
                    class="mb-16 p-8 rounded-2xl border-2"
                    :style="{
                        backgroundColor: isDark ? 'rgba(255, 255, 255, 0.02)' : 'rgba(255, 255, 255, 0.8)',
                        borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                    }"
                >
                    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
                        <!-- Frontend Layer -->
                        <div 
                            class="text-center p-6 rounded-xl border"
                            :style="{
                                backgroundColor: isDark ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.05)',
                                borderColor: isDark ? 'rgba(59, 130, 246, 0.3)' : 'rgba(59, 130, 246, 0.2)'
                            }"
                        >
                            <div class="text-blue-500 text-4xl mb-4">🎨</div>
                            <h3 class="text-xl font-bold mb-3" :style="{ color: isDark ? '#ffffff' : '#111827' }">Frontend</h3>
                            <div class="space-y-2 text-sm" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                                <div>Vue 3 + TypeScript</div>
                                <div>Vite Build System</div>
                                <div>TailwindCSS</div>
                                <div>Progressive Web App</div>
                            </div>
                        </div>

                        <!-- API Gateway -->
                        <div 
                            class="text-center p-6 rounded-xl border"
                            :style="{
                                backgroundColor: isDark ? 'rgba(16, 185, 129, 0.1)' : 'rgba(16, 185, 129, 0.05)',
                                borderColor: isDark ? 'rgba(16, 185, 129, 0.3)' : 'rgba(16, 185, 129, 0.2)'
                            }"
                        >
                            <div class="text-green-500 text-4xl mb-4">⚡</div>
                            <h3 class="text-xl font-bold mb-3" :style="{ color: isDark ? '#ffffff' : '#111827' }">Backend</h3>
                            <div class="space-y-2 text-sm" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                                <div>Go + Fiber Framework</div>
                                <div>REST API + GraphQL</div>
                                <div>JWT Authentication</div>
                                <div>Rate Limiting</div>
                            </div>
                        </div>

                        <!-- Database Layer -->
                        <div 
                            class="text-center p-6 rounded-xl border"
                            :style="{
                                backgroundColor: isDark ? 'rgba(147, 51, 234, 0.1)' : 'rgba(147, 51, 234, 0.05)',
                                borderColor: isDark ? 'rgba(147, 51, 234, 0.3)' : 'rgba(147, 51, 234, 0.2)'
                            }"
                        >
                            <div class="text-purple-500 text-4xl mb-4">🗄️</div>
                            <h3 class="text-xl font-bold mb-3" :style="{ color: isDark ? '#ffffff' : '#111827' }">Storage</h3>
                            <div class="space-y-2 text-sm" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                                <div>PostgreSQL/SQLite</div>
                                <div>Redis Caching</div>
                                <div>S3-Compatible Storage</div>
                                <div>CDN Integration</div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Technical Specifications -->
                <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
                    <!-- Performance Metrics -->
                    <div 
                        class="p-8 rounded-2xl border"
                        :style="{
                            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.02)' : 'rgba(255, 255, 255, 0.8)',
                            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                        }"
                    >
                        <h3 class="text-2xl font-bold mb-6" :style="{ color: isDark ? '#ffffff' : '#111827' }">Performance Metrics</h3>
                        <div class="space-y-4">
                            <div class="flex justify-between items-center">
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Concurrent Users</span>
                                <span class="font-mono text-xl text-green-500">10,000+</span>
                            </div>
                            <div class="flex justify-between items-center">
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">File Upload Speed</span>
                                <span class="font-mono text-xl text-blue-500">1GB/s</span>
                            </div>
                            <div class="flex justify-between items-center">
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Maximum File Size</span>
                                <span class="font-mono text-xl text-purple-500">Unlimited</span>
                            </div>
                            <div class="flex justify-between items-center">
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Data Encryption</span>
                                <span class="font-mono text-xl text-yellow-500">AES-256</span>
                            </div>
                        </div>
                    </div>

                    <!-- Deployment Options -->
                    <div 
                        class="p-8 rounded-2xl border"
                        :style="{
                            backgroundColor: isDark ? 'rgba(255, 255, 255, 0.02)' : 'rgba(255, 255, 255, 0.8)',
                            borderColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)'
                        }"
                    >
                        <h3 class="text-2xl font-bold mb-6" :style="{ color: isDark ? '#ffffff' : '#111827' }">Deployment</h3>
                        <div class="space-y-4">
                            <div class="flex items-center space-x-3">
                                <div class="w-3 h-3 bg-green-500 rounded-full"></div>
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Docker Containers</span>
                            </div>
                            <div class="flex items-center space-x-3">
                                <div class="w-3 h-3 bg-green-500 rounded-full"></div>
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Kubernetes Ready</span>
                            </div>
                            <div class="flex items-center space-x-3">
                                <div class="w-3 h-3 bg-green-500 rounded-full"></div>
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Auto-Scaling</span>
                            </div>
                            <div class="flex items-center space-x-3">
                                <div class="w-3 h-3 bg-green-500 rounded-full"></div>
                                <span :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Load Balancing</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>

        <!-- Enterprise Upload Section -->
        <section ref="uploadSection" class="py-20 px-4 sm:px-6 lg:px-8">
            <div class="max-w-4xl mx-auto">
                <div class="text-center mb-12">
                    <h2 
                        class="text-4xl font-bold mb-6"
                        :style="{ color: isDark ? '#ffffff' : '#111827' }"
                    >
                        Secure File Transfer
                    </h2>
                    <p 
                        class="text-xl"
                        :style="{ color: isDark ? '#d1d5db' : '#4b5563' }"
                    >
                        Enterprise-grade security with military-level encryption
                    </p>
                </div>

                <!-- Upload Interface -->
                <div 
                    class="relative p-8 rounded-2xl border-2 border-dashed transition-all duration-300 hover:border-blue-500"
                    :class="{ 'border-blue-500 bg-blue-50 dark:bg-blue-900/20': isDragging }"
                    :style="{
                        backgroundColor: isDark ? 'rgba(255, 255, 255, 0.02)' : 'rgba(255, 255, 255, 0.8)',
                        borderColor: isDragging ? '#3b82f6' : (isDark ? 'rgba(255, 255, 255, 0.2)' : 'rgba(0, 0, 0, 0.2)')
                    }"
                    @dragover.prevent
                    @drop.prevent="onDrop"
                    @dragenter="isDragging = true"
                    @dragleave="isDragging = false"
                    @click="triggerFileInput"
                >
                    <!-- Upload Progress -->
                    <div v-if="isUploading" class="text-center">
                        <div class="mb-6">
                            <div class="animate-spin rounded-full h-16 w-16 border-b-2 border-blue-500 mx-auto"></div>
                        </div>
                        <h3 class="text-2xl font-bold mb-4" :style="{ color: isDark ? '#ffffff' : '#111827' }">
                            Uploading Files...
                        </h3>
                        <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-3 mb-4">
                            <div 
                                class="bg-gradient-to-r from-blue-500 to-purple-500 h-3 rounded-full transition-all duration-300" 
                                :style="{ width: `${progress}%` }"
                            ></div>
                        </div>
                        <p class="text-lg" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                            {{ Math.round(progress) }}% Complete
                        </p>
                    </div>

                    <!-- Upload Success -->
                    <div v-else-if="uploadComplete" class="text-center">
                        <div class="text-green-500 text-6xl mb-6">✅</div>
                        <h3 class="text-2xl font-bold mb-4" :style="{ color: isDark ? '#ffffff' : '#111827' }">
                            Upload Successful!
                        </h3>
                        <p class="text-lg mb-8" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                            Your files are ready to share
                        </p>
                        <div class="flex gap-4 justify-center">
                            <button
                                @click="copyLink"
                                class="px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-colors font-medium"
                            >
                                Copy Link
                            </button>
                            <button
                                @click="uploadNew"
                                class="px-6 py-3 border-2 border-gray-300 dark:border-gray-600 rounded-lg transition-colors font-medium"
                                :style="{ color: isDark ? '#ffffff' : '#111827' }"
                            >
                                Upload More
                            </button>
                        </div>
                    </div>

                    <!-- Upload Zone -->
                    <div v-else class="text-center cursor-pointer">
                        <div class="text-6xl mb-6">📁</div>
                        <h3 class="text-2xl font-bold mb-4" :style="{ color: isDark ? '#ffffff' : '#111827' }">
                            {{ selectedFiles.length > 0 ? `${selectedFiles.length} file(s) selected` : 'Drop files here or click to browse' }}
                        </h3>
                        <p class="text-lg mb-8" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                            Supports any file type • No size limits • End-to-end encrypted
                        </p>

                        <!-- Selected Files List -->
                        <div v-if="selectedFiles.length > 0" class="mb-8">
                            <div class="max-h-48 overflow-y-auto space-y-2">
                                <div 
                                    v-for="(file, index) in selectedFiles" 
                                    :key="index"
                                    class="flex items-center justify-between p-3 rounded-lg"
                                    :style="{
                                        backgroundColor: isDark ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.05)'
                                    }"
                                >
                                    <div class="flex items-center space-x-3">
                                        <div class="text-2xl">📄</div>
                                        <div>
                                            <div class="font-medium" :style="{ color: isDark ? '#ffffff' : '#111827' }">
                                                {{ file.name }}
                                            </div>
                                            <div class="text-sm" :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">
                                                {{ formatFileSize(file.size) }}
                                            </div>
                                        </div>
                                    </div>
                                    <button
                                        @click.stop="removeFile(index)"
                                        class="text-red-500 hover:text-red-700 transition-colors"
                                    >
                                        ✕
                                    </button>
                                </div>
                            </div>
                        </div>

                        <!-- Upload Button -->
                        <button
                            v-if="selectedFiles.length > 0"
                            @click.stop="uploadFiles"
                            class="px-8 py-4 bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white rounded-xl font-semibold text-lg transition-all duration-300 transform hover:scale-105 shadow-lg"
                        >
                            Upload {{ selectedFiles.length }} file{{ selectedFiles.length > 1 ? 's' : '' }}
                        </button>

                        <input
                            ref="fileInput"
                            type="file"
                            multiple
                            class="hidden"
                            @change="onFileChange"
                        />
                    </div>
                </div>

                <!-- Security Features -->
                <div class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-8">
                    <div class="text-center">
                        <div class="text-4xl mb-4">🔒</div>
                        <h3 class="text-xl font-bold mb-2" :style="{ color: isDark ? '#ffffff' : '#111827' }">End-to-End Encryption</h3>
                        <p :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Files are encrypted before leaving your device</p>
                    </div>
                    <div class="text-center">
                        <div class="text-4xl mb-4">🚀</div>
                        <h3 class="text-xl font-bold mb-2" :style="{ color: isDark ? '#ffffff' : '#111827' }">Lightning Fast</h3>
                        <p :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Optimized for maximum transfer speeds</p>
                    </div>
                    <div class="text-center">
                        <div class="text-4xl mb-4">🌐</div>
                        <h3 class="text-xl font-bold mb-2" :style="{ color: isDark ? '#ffffff' : '#111827' }">Global CDN</h3>
                        <p :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Distributed worldwide for optimal performance</p>
                    </div>
                </div>
            </div>
        </section>

        <!-- Open Source Community Section -->
        <section class="py-20 px-4 sm:px-6 lg:px-8">
            <div class="max-w-6xl mx-auto text-center">
                <h2 class="text-4xl font-bold mb-12" :style="{ color: isDark ? '#ffffff' : '#111827' }">
                    Join the Community
                </h2>
                
                <div class="grid grid-cols-2 md:grid-cols-4 gap-8 mb-12">
                    <div class="text-center">
                        <div class="text-3xl font-bold text-yellow-500 mb-2">{{ Math.floor(Math.random() * 500 + 1500) }}</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">GitHub Stars</div>
                    </div>
                    <div class="text-center">
                        <div class="text-3xl font-bold text-blue-500 mb-2">{{ Math.floor(Math.random() * 100 + 200) }}</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Contributors</div>
                    </div>
                    <div class="text-center">
                        <div class="text-3xl font-bold text-green-500 mb-2">{{ Math.floor(Math.random() * 50 + 150) }}</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Forks</div>
                    </div>
                    <div class="text-center">
                        <div class="text-3xl font-bold text-purple-500 mb-2">{{ Math.floor(Math.random() * 20 + 80) }}</div>
                        <div :style="{ color: isDark ? '#d1d5db' : '#6b7280' }">Releases</div>
                    </div>
                </div>

                <div class="flex flex-wrap gap-4 justify-center">
                    <a
                        href="https://github.com/R8bert/pingo-share"
                        target="_blank"
                        class="px-6 py-3 bg-gray-900 hover:bg-gray-800 text-white rounded-lg transition-colors font-medium flex items-center space-x-2"
                    >
                        <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M10 0C4.477 0 0 4.484 0 10.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0110 4.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.203 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.942.359.31.678.921.678 1.856 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0020 10.017C20 4.484 15.522 0 10 0z"/>
                        </svg>
                        <span>View on GitHub</span>
                    </a>
                    <button
                        class="px-6 py-3 border-2 rounded-lg transition-colors font-medium"
                        :style="{
                            borderColor: isDark ? 'rgba(255,255,255,0.2)' : 'rgba(0,0,0,0.2)',
                            color: isDark ? '#ffffff' : '#000000'
                        }"
                    >
                        Documentation
                    </button>
                    <button
                        class="px-6 py-3 border-2 rounded-lg transition-colors font-medium"
                        :style="{
                            borderColor: isDark ? 'rgba(255,255,255,0.2)' : 'rgba(0,0,0,0.2)',
                            color: isDark ? '#ffffff' : '#000000'
                        }"
                    >
                        Join Discord
                    </button>
                </div>
            </div>
        </section>
        </main>
    </div>
</template>
                                        </svg>
                                        Upload New Files
                                    </button>
                                </div>
                            </div>

                            <!-- Upload Zone -->
                            <div
                                v-else
                                @dragover.prevent
                                @drop.prevent="onDrop"
                                @dragenter="isDragging = true"
                                @dragleave="isDragging = false"
                                @click="triggerFileInput"
                                :class="[
                                    'relative border-2 border-dashed transition-all duration-300 cursor-pointer rounded-2xl p-12',
                                    isDragging
                                        ? 'border-blue-500 bg-blue-500/10 scale-[1.02]'
                                        : 'border-gray-300 dark:border-gray-600 hover:border-blue-400 hover:bg-blue-500/5',
                                ]"
                            >
                                <input
                                    type="file"
                                    ref="fileInput"
                                    @change="onFileChange"
                                    class="hidden"
                                    multiple
                                />

                                <div class="text-center space-y-6">
                                    <!-- Upload Icon -->
                                    <div class="relative mx-auto w-24 h-24">
                                        <div
                                            class="absolute inset-0 rounded-2xl blur-xl opacity-30"
                                            :style="{
                                                background: primaryGradient,
                                            }"
                                        ></div>
                                        <div
                                            class="relative rounded-xl sm:rounded-2xl p-4 sm:p-6 flex items-center justify-center"
                                            :style="{
                                                background: primaryGradient,
                                            }"
                                        >
                                            <CloudArrowUpIcon
                                                class="w-10 h-10 sm:w-12 sm:h-12 text-white"
                                            />
                                        </div>
                                    </div>

                                    <!-- Upload Text -->
                                    <div class="space-y-2">
                                        <h3
                                            class="text-xl sm:text-2xl font-bold"
                                            :style="{
                                                color: isDark
                                                    ? '#ffffff'
                                                    : '#000000',
                                            }"
                                        >
                                            {{
                                                selectedFiles.length > 0
                                                    ? `${
                                                          selectedFiles.length
                                                      } file${
                                                          selectedFiles.length >
                                                          1
                                                              ? "s"
                                                              : ""
                                                      } selected`
                                                    : "Drop files here or click to browse"
                                            }}
                                        </h3>
                                        <p
                                            class="text-sm sm:text-base"
                                            :style="{
                                                color: isDark
                                                    ? '#a1a1aa'
                                                    : '#71717a',
                                            }"
                                        >
                                            Support for any file type • No size
                                            limits • Encrypted transfer
                                        </p>
                                    </div>

                                    <!-- Features -->
                                    <div
                                        class="grid grid-cols-1 sm:grid-cols-3 gap-3 sm:gap-4 pt-4 sm:pt-6"
                                    >
                                        <div
                                            class="flex items-center justify-center space-x-2 p-2 sm:p-3 rounded-lg sm:rounded-xl"
                                            :style="{
                                                backgroundColor: isDark
                                                    ? 'rgba(255,255,255,0.05)'
                                                    : 'rgba(0,0,0,0.05)',
                                            }"
                                        >
                                            <div
                                                class="w-2 h-2 bg-green-500 rounded-full"
                                            ></div>
                                            <span
                                                class="text-xs sm:text-sm font-medium"
                                                :style="{
                                                    color: isDark
                                                        ? '#a1a1aa'
                                                        : '#71717a',
                                                }"
                                            >
                                                Secure
                                            </span>
                                        </div>
                                        <div
                                            class="flex items-center justify-center space-x-2 p-2 sm:p-3 rounded-lg sm:rounded-xl"
                                            :style="{
                                                backgroundColor: isDark
                                                    ? 'rgba(255,255,255,0.05)'
                                                    : 'rgba(0,0,0,0.05)',
                                            }"
                                        >
                                            <div
                                                class="w-2 h-2 bg-blue-500 rounded-full"
                                            ></div>
                                            <span
                                                class="text-xs sm:text-sm font-medium"
                                                :style="{
                                                    color: isDark
                                                        ? '#a1a1aa'
                                                        : '#71717a',
                                                }"
                                            >
                                                Fast
                                            </span>
                                        </div>
                                        <div
                                            class="flex items-center justify-center space-x-2 p-2 sm:p-3 rounded-lg sm:rounded-xl"
                                            :style="{
                                                backgroundColor: isDark
                                                    ? 'rgba(255,255,255,0.05)'
                                                    : 'rgba(0,0,0,0.05)',
                                            }"
                                        >
                                            <div
                                                class="w-2 h-2 bg-purple-500 rounded-full"
                                            ></div>
                                            <span
                                                class="text-xs sm:text-sm font-medium"
                                                :style="{
                                                    color: isDark
                                                        ? '#a1a1aa'
                                                        : '#71717a',
                                                }"
                                            >
                                                Simple
                                            </span>
                                        </div>
                                    </div>
                                </div>

                                <!-- Drag overlay -->
                                <div
                                    v-if="isDragging"
                                    class="absolute inset-0 rounded-xl sm:rounded-2xl flex items-center justify-center z-20 bg-blue-500/20 backdrop-blur-sm"
                                >
                                    <div
                                        class="text-xl sm:text-2xl font-bold text-blue-600"
                                    >
                                        Drop files here
                                    </div>
                                </div>
                            </div>

                            <!-- Selected Files -->

                            <div
                                v-if="
                                    selectedFiles.length > 0 && !uploadComplete
                                "
                                class="mt-8 space-y-4"
                            >
                                <!-- Validity Selection -->
                                <div
                                    class="mb-6"
                                    v-if="validityOptions.length > 0"
                                >
                                    <label
                                        class="block text-sm font-medium mb-2 transition-colors duration-300"
                                        :style="{
                                            color: isDark
                                                ? '#d1d5db'
                                                : '#374151',
                                        }"
                                    >
                                        🕐 File Expiration
                                    </label>

                                    <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-2">
                                        <label
                                            v-for="option in validityOptions"
                                            :key="option.value"
                                            class="flex flex-col items-center p-2 rounded-lg border cursor-pointer transition-all duration-200 hover:scale-105 text-center"
                                            :class="[
                                                selectedValidity === option.value
                                                    ? isDark
                                                        ? 'border-amber-500 bg-amber-600/20 shadow-md shadow-amber-500/10'
                                                        : 'border-amber-500 bg-amber-50 shadow-md shadow-amber-500/10'
                                                    : isDark
                                                    ? 'border-gray-600/50 bg-gray-800/30 hover:border-gray-500'
                                                    : 'border-gray-200 bg-white hover:border-gray-300 hover:shadow-sm'
                                            ]"
                                        >
                                            <input
                                                type="radio"
                                                v-model="selectedValidity"
                                                :value="option.value"
                                                class="sr-only"
                                            />
                                            
                                            <!-- Custom Radio Button -->
                                            <div
                                                class="w-3 h-3 rounded-full border flex items-center justify-center mb-1 flex-shrink-0 transition-all duration-200"
                                                :class="[
                                                    selectedValidity === option.value
                                                        ? isDark
                                                            ? 'border-amber-400 bg-amber-500'
                                                            : 'border-amber-500 bg-amber-500'
                                                        : isDark
                                                        ? 'border-gray-500'
                                                        : 'border-gray-300'
                                                ]"
                                            >
                                                <div
                                                    v-if="selectedValidity === option.value"
                                                    class="w-1 h-1 rounded-full bg-white"
                                                ></div>
                                            </div>
                                            
                                            <!-- Option Content -->
                                            <div class="min-w-0">
                                                <div 
                                                    class="font-medium text-xs mb-0.5 leading-tight"
                                                    :class="[
                                                        selectedValidity === option.value
                                                            ? isDark
                                                                ? 'text-amber-300'
                                                                : 'text-amber-700'
                                                            : isDark
                                                            ? 'text-gray-100'
                                                            : 'text-gray-900'
                                                    ]"
                                                >
                                                    {{ option.label }}
                                                </div>
                                                <div 
                                                    class="text-xs opacity-75 leading-tight"
                                                    :class="[
                                                        selectedValidity === option.value
                                                            ? isDark
                                                                ? 'text-amber-200'
                                                                : 'text-amber-600'
                                                            : isDark
                                                            ? 'text-gray-400'
                                                            : 'text-gray-500'
                                                    ]"
                                                >
                                                    {{ option.description }}
                                                </div>
                                            </div>
                                        </label>
                                    </div>
                                </div>
                                <h4
                                    class="text-lg font-semibold"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    Selected Files
                                </h4>
                                <div class="grid gap-3 sm:gap-4">
                                    <div
                                        v-for="(file, index) in selectedFiles"
                                        :key="index"
                                        class="rounded-xl border transition-all duration-300 overflow-hidden"
                                        :style="{
                                            backgroundColor: isDark
                                                ? 'rgba(255,255,255,0.05)'
                                                : 'rgba(0,0,0,0.05)',
                                            borderColor: isDark
                                                ? 'rgba(255,255,255,0.1)'
                                                : 'rgba(0,0,0,0.1)',
                                        }"
                                    >
                                        <div
                                            class="flex flex-col sm:flex-row sm:items-center gap-3 sm:gap-4 p-3 sm:p-4"
                                        >
                                            <div
                                                class="flex items-center space-x-3 flex-1 min-w-0"
                                            >
                                                <div
                                                    class="flex items-center justify-center flex-shrink-0"
                                                >
                                                    <img 
                                                        :src="getFileIconPath(file.name)" 
                                                        :alt="getFileIconAltText(file.name)"
                                                        class="w-5 h-5 sm:w-6 sm:h-6 object-contain"
                                                    />
                                                </div>
                                                <div class="min-w-0 flex-1">
                                                    <p
                                                        class="font-medium text-sm sm:text-base truncate break-all"
                                                        :style="{
                                                            color: isDark
                                                                ? '#ffffff'
                                                                : '#000000',
                                                        }"
                                                        :title="file.name"
                                                    >
                                                        {{ file.name }}
                                                    </p>
                                                    <p
                                                        class="text-xs sm:text-sm"
                                                        :style="{
                                                            color: isDark
                                                                ? '#a1a1aa'
                                                                : '#71717a',
                                                        }"
                                                    >
                                                        {{
                                                            formatFileSize(
                                                                file.size
                                                            )
                                                        }}
                                                    </p>
                                                </div>
                                            </div>
                                            <div
                                                class="flex items-center gap-2 w-full sm:w-auto"
                                            >
                                                <button
                                                    v-if="canPreview(file)"
                                                    @click="
                                                        togglePreview(index)
                                                    "
                                                    class="px-2 sm:px-3 py-2 rounded-lg text-xs sm:text-sm font-medium transition-all duration-200 hover:scale-105 flex-1 sm:flex-none"
                                                    :class="
                                                        previewingFiles.has(
                                                            index
                                                        )
                                                            ? 'bg-red-500 hover:bg-red-600 text-white'
                                                            : 'bg-emerald-500 hover:bg-emerald-600 text-white'
                                                    "
                                                    :disabled="isUploading"
                                                >
                                                    <EyeIcon
                                                        v-if="
                                                            !previewingFiles.has(
                                                                index
                                                            )
                                                        "
                                                        class="w-3 h-3 sm:w-4 sm:h-4 mr-1 inline"
                                                    />
                                                    <EyeSlashIcon
                                                        v-else
                                                        class="w-3 h-3 sm:w-4 sm:h-4 mr-1 inline"
                                                    />
                                                    {{
                                                        previewingFiles.has(
                                                            index
                                                        )
                                                            ? "Hide"
                                                            : "Preview"
                                                    }}
                                                </button>
                                                <button
                                                    @click="removeFile(index)"
                                                    class="p-2 rounded-lg text-red-500 hover:bg-red-500 hover:text-white transition-colors flex-shrink-0"
                                                    :disabled="isUploading"
                                                    :title="'Remove ' + file.name"
                                                >
                                                    <XMarkIcon class="w-3 h-3 sm:w-4 sm:h-4" />
                                                </button>
                                            </div>

                                            <!-- Expiration selector -->
                                        </div>

                                        <!-- Preview Section -->
                                        <div
                                            v-if="previewingFiles.has(index)"
                                            class="px-3 sm:px-4 pb-3 sm:pb-4"
                                        >
                                            <div
                                                class="rounded-lg border-t pt-3 sm:pt-4"
                                                :style="{
                                                    borderColor: isDark
                                                        ? 'rgba(255,255,255,0.1)'
                                                        : 'rgba(0,0,0,0.1)',
                                                }"
                                            >
                                                <!-- Image preview -->
                                                <div
                                                    v-if="
                                                        [
                                                            'jpg',
                                                            'jpeg',
                                                            'png',
                                                            'gif',
                                                            'webp',
                                                        ].includes(
                                                            getFileExtension(
                                                                file
                                                            )
                                                        )
                                                    "
                                                >
                                                    <img
                                                        :src="
                                                            previewUrls.get(
                                                                index
                                                            )
                                                        "
                                                        :alt="file.name"
                                                        class="w-full max-w-xs sm:max-w-md mx-auto rounded-lg shadow-lg"
                                                    />
                                                </div>

                                                <!-- Video preview -->
                                                <div
                                                    v-else-if="
                                                        getFileExtension(
                                                            file
                                                        ) === 'mp4'
                                                    "
                                                >
                                                    <video
                                                        :src="
                                                            previewUrls.get(
                                                                index
                                                            )
                                                        "
                                                        controls
                                                        class="w-full max-w-xs sm:max-w-md mx-auto rounded-lg shadow-lg"
                                                    ></video>
                                                </div>

                                                <!-- Audio preview -->
                                                <div
                                                    v-else-if="
                                                        [
                                                            'mp3',
                                                            'wav',
                                                            'flac',
                                                        ].includes(
                                                            getFileExtension(
                                                                file
                                                            )
                                                        )
                                                    "
                                                >
                                                    <audio
                                                        :src="
                                                            previewUrls.get(
                                                                index
                                                            )
                                                        "
                                                        controls
                                                        class="w-full max-w-xs sm:max-w-md mx-auto"
                                                    ></audio>
                                                </div>

                                                <!-- Text file preview -->
                                                <div
                                                    v-else-if="
                                                        [
                                                            'txt',
                                                            'md',
                                                            'json',
                                                            'csv',
                                                            'xml',
                                                            'torrent',
                                                        ].includes(
                                                            getFileExtension(
                                                                file
                                                            )
                                                        )
                                                    "
                                                    class="w-full"
                                                >
                                                    <div
                                                        class="w-full max-w-full overflow-hidden rounded-lg"
                                                        :style="{
                                                            backgroundColor:
                                                                isDark
                                                                    ? 'rgba(0,0,0,0.3)'
                                                                    : 'rgba(255,255,255,0.7)',
                                                        }"
                                                    >
                                                        <pre
                                                            class="text-xs sm:text-sm p-3 sm:p-4 whitespace-pre-wrap overflow-auto max-h-32 w-full block"
                                                            :style="{
                                                                color: isDark
                                                                    ? '#d1d5db'
                                                                    : '#374151',
                                                                wordBreak:
                                                                    'break-word',
                                                                overflowWrap:
                                                                    'break-word',
                                                                maxWidth:
                                                                    '100%',
                                                                display:
                                                                    'block',
                                                            }"
                                                            >{{
                                                                textPreviews.get(
                                                                    index
                                                                )
                                                            }}</pre
                                                        >
                                                    </div>
                                                </div>

                                                <!-- PDF preview -->
                                                <div
                                                    v-else-if="
                                                        getFileExtension(
                                                            file
                                                        ) === 'pdf'
                                                    "
                                                >
                                                    <div
                                                        class="text-center p-4"
                                                    >
                                                        <svg
                                                            class="w-16 h-16 mx-auto mb-2 text-red-500"
                                                            fill="currentColor"
                                                            viewBox="0 0 24 24"
                                                        >
                                                            <path
                                                                d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"
                                                            />
                                                        </svg>
                                                        <p
                                                            class="text-sm"
                                                            :style="{
                                                                color: isDark
                                                                    ? '#a1a1aa'
                                                                    : '#71717a',
                                                            }"
                                                        >
                                                            PDF preview not
                                                            available
                                                        </p>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>

                                <!-- Progress Bar -->
                                <div
                                    v-if="progress > 0 && progress < 100"
                                    class="mt-4 sm:mt-6"
                                >
                                    <div
                                        class="flex justify-between text-xs sm:text-sm mb-2"
                                        :style="{
                                            color: isDark
                                                ? '#a1a1aa'
                                                : '#71717a',
                                        }"
                                    >
                                        <span>Uploading...</span>
                                        <span>{{ progress }}%</span>
                                    </div>
                                    <div
                                        class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2 sm:h-3 overflow-hidden"
                                    >
                                        <div
                                            class="h-2 sm:h-3 rounded-full transition-all duration-300"
                                            :style="{
                                                background: buttonGradient,
                                                width: progress + '%',
                                            }"
                                        ></div>
                                    </div>
                                </div>

                                <!-- Upload Button -->
                                <div v-if="!uploadComplete" class="pt-4 sm:pt-6">
                                    <button
                                        @click="uploadFiles"
                                        :disabled="isUploading"
                                        class="w-full py-3 sm:py-4 text-sm sm:text-base text-white font-semibold rounded-xl sm:rounded-2xl transition-all duration-300 hover:scale-[1.02] hover:shadow-2xl disabled:opacity-50 disabled:hover:scale-100"
                                        :style="{ background: buttonGradient }"
                                    >
                                        <span
                                            v-if="isUploading"
                                            class="flex items-center justify-center"
                                        >
                                            <svg
                                                class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
                                                xmlns="http://www.w3.org/2000/svg"
                                                fill="none"
                                                viewBox="0 0 24 24"
                                            >
                                                <circle
                                                    class="opacity-25"
                                                    cx="12"
                                                    cy="12"
                                                    r="10"
                                                    stroke="currentColor"
                                                    stroke-width="4"
                                                ></circle>
                                                <path
                                                    class="opacity-75"
                                                    fill="currentColor"
                                                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                                                ></path>
                                            </svg>
                                            {{
                                                progress > 0
                                                    ? `Uploading... ${progress}%`
                                                    : "Uploading..."
                                            }}
                                        </span>
                                        <span v-else>Upload Files</span>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Features Section -->
            <section
                ref="featuresSection"
                class="relative py-12 sm:py-16 lg:py-20 px-4 sm:px-6 lg:px-8"
            >
                <div class="max-w-7xl mx-auto">
                    <div class="text-center mb-12 sm:mb-16">
                        <h2
                            class="text-3xl sm:text-4xl lg:text-5xl font-bold mb-4 sm:mb-6"
                            :style="{ color: isDark ? '#ffffff' : '#000000' }"
                        >
                            Why choose PInGO Share?
                        </h2>
                        <p
                            class="text-lg sm:text-xl max-w-3xl mx-auto"
                            :style="{ color: isDark ? '#a1a1aa' : '#71717a' }"
                        >
                            Built with modern technology and user experience in
                            mind
                        </p>
                    </div>

                    <div
                        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 sm:gap-8"
                    >
                        <!-- Feature 1 -->
                        <div class="relative group">
                            <div
                                class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                            ></div>
                            <div
                                class="relative backdrop-blur-xl border rounded-xl sm:rounded-2xl p-6 sm:p-8 transition-all duration-300 group-hover:scale-105"
                                :style="{
                                    backgroundColor: isDark
                                        ? 'rgba(255,255,255,0.05)'
                                        : 'rgba(255,255,255,0.8)',
                                    borderColor: isDark
                                        ? 'rgba(255,255,255,0.1)'
                                        : 'rgba(0,0,0,0.1)',
                                }"
                            >
                                <div
                                    class="w-12 h-12 sm:w-16 sm:h-16 bg-gradient-to-r from-blue-500 to-purple-500 rounded-xl sm:rounded-2xl flex items-center justify-center mb-4 sm:mb-6"
                                >
                                    <svg
                                        class="w-6 h-6 sm:w-8 sm:h-8 text-white"
                                        fill="none"
                                        stroke="currentColor"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                                        ></path>
                                    </svg>
                                </div>
                                <h3
                                    class="text-lg sm:text-xl font-bold mb-3 sm:mb-4"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    End-to-End Encryption
                                </h3>
                                <p
                                    class="text-sm sm:text-base"
                                    :style="{
                                        color: isDark ? '#a1a1aa' : '#71717a',
                                    }"
                                >
                                    Your files are encrypted before they leave
                                    your device, ensuring maximum security.
                                </p>
                            </div>
                        </div>

                        <!-- Feature 2 -->
                        <div class="relative group">
                            <div
                                class="absolute inset-0 bg-gradient-to-r from-green-600/20 to-blue-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                            ></div>
                            <div
                                class="relative backdrop-blur-xl border rounded-xl sm:rounded-2xl p-6 sm:p-8 transition-all duration-300 group-hover:scale-105"
                                :style="{
                                    backgroundColor: isDark
                                        ? 'rgba(255,255,255,0.05)'
                                        : 'rgba(255,255,255,0.8)',
                                    borderColor: isDark
                                        ? 'rgba(255,255,255,0.1)'
                                        : 'rgba(0,0,0,0.1)',
                                }"
                            >
                                <div
                                    class="w-12 h-12 sm:w-16 sm:h-16 bg-gradient-to-r from-green-500 to-blue-500 rounded-xl sm:rounded-2xl flex items-center justify-center mb-4 sm:mb-6"
                                >
                                    <svg
                                        class="w-6 h-6 sm:w-8 sm:h-8 text-white"
                                        fill="none"
                                        stroke="currentColor"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            d="M13 10V3L4 14h7v7l9-11h-7z"
                                        ></path>
                                    </svg>
                                </div>
                                <h3
                                    class="text-lg sm:text-xl font-bold mb-3 sm:mb-4"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    Lightning Fast
                                </h3>
                                <p
                                    class="text-sm sm:text-base"
                                    :style="{
                                        color: isDark ? '#a1a1aa' : '#71717a',
                                    }"
                                >
                                    Optimized for speed with parallel uploads
                                    and smart compression.
                                </p>
                            </div>
                        </div>

                        <!-- Feature 3 -->
                        <div class="relative group">
                            <div
                                class="absolute inset-0 bg-gradient-to-r from-purple-600/20 to-pink-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                            ></div>
                            <div
                                class="relative backdrop-blur-xl border rounded-xl sm:rounded-2xl p-6 sm:p-8 transition-all duration-300 group-hover:scale-105"
                                :style="{
                                    backgroundColor: isDark
                                        ? 'rgba(255,255,255,0.05)'
                                        : 'rgba(255,255,255,0.8)',
                                    borderColor: isDark
                                        ? 'rgba(255,255,255,0.1)'
                                        : 'rgba(0,0,0,0.1)',
                                }"
                            >
                                <div
                                    class="w-12 h-12 sm:w-16 sm:h-16 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl sm:rounded-2xl flex items-center justify-center mb-4 sm:mb-6"
                                >
                                    <svg
                                        class="w-6 h-6 sm:w-8 sm:h-8 text-white"
                                        fill="none"
                                        stroke="currentColor"
                                        viewBox="0 0 24 24"
                                    >
                                        <path
                                            stroke-linecap="round"
                                            stroke-linejoin="round"
                                            stroke-width="2"
                                            d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
                                        ></path>
                                    </svg>
                                </div>
                                <h3
                                    class="text-lg sm:text-xl font-bold mb-3 sm:mb-4"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    No Registration
                                </h3>
                                <p
                                    class="text-sm sm:text-base"
                                    :style="{
                                        color: isDark ? '#a1a1aa' : '#71717a',
                                    }"
                                >
                                    Start sharing immediately without creating
                                    accounts or providing personal info.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Footer -->
            <footer
                class="relative py-12 px-4 sm:px-6 lg:px-8 border-t"
                :style="{
                    borderColor: isDark
                        ? 'rgba(255,255,255,0.1)'
                        : 'rgba(0,0,0,0.1)',
                }"
            >
                <div class="max-w-7xl mx-auto text-center">
                    <p :style="{ color: isDark ? '#a1a1aa' : '#71717a' }">
                        © 2025 PInGO Share. Built with ❤️ for the community.
                    </p>
                </div>
            </footer>
        </main>

        <!-- Messages -->
        <div
            v-if="message.text"
            class="fixed top-6 right-6 p-6 rounded-3xl shadow-2xl z-50 backdrop-blur-xl border max-w-sm animate-fade-in"
            :style="{
                backgroundColor:
                    message.type === 'success'
                        ? isDark
                            ? 'rgba(34, 197, 94, 0.15)'
                            : 'rgba(34, 197, 94, 0.1)'
                        : isDark
                        ? 'rgba(239, 68, 68, 0.15)'
                        : 'rgba(239, 68, 68, 0.1)',
                borderColor: message.type === 'success' ? '#22c55e' : '#ef4444',
                color: message.type === 'success' ? '#22c55e' : '#ef4444',
            }"
        >
            <div class="flex items-center space-x-3">
                <div
                    v-if="message.type === 'success'"
                    class="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0"
                >
                    <svg
                        class="w-5 h-5 text-white"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M5 13l4 4L19 7"
                        ></path>
                    </svg>
                </div>
                <div
                    v-else
                    class="w-8 h-8 bg-red-500 rounded-full flex items-center justify-center flex-shrink-0"
                >
                    <svg
                        class="w-5 h-5 text-white"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M6 18L18 6M6 6l12 12"
                        ></path>
                    </svg>
                </div>
                <div class="flex-1">
                    <p class="font-semibold">{{ message.text }}</p>
                    <p
                        v-if="message.type === 'success' && progress === 100"
                        class="text-sm opacity-75 mt-1"
                    >
                        Redirecting to download page...
                    </p>
                </div>
            </div>
        </div>

        <!-- GitHub Icon - Bottom Left -->
        <!-- Container for GitHub Icon and ShinyText -->
        <!-- Container for GitHub Icon and ShinyText -->
        <a
            href="https://github.com/R8bert/PInGO-Share"
            target="_blank"
            rel="noopener noreferrer"
            class="fixed bottom-4 left-4 z-40 flex items-center space-x-2"
            :style="{ color: isDark ? '#9ca3af' : '#6b7280' }"
            title="View GitHub Repository"
        >
            <!-- GitHub Icon -->
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                <path
                    d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"
                />
            </svg>
            <!-- ShinyText Component -->
            <ShinyText
                text="Star on Github"
                :disabled="false"
                :speed="3"
                class-name="text-sm font-medium"
                aria-label="Encourage starring the PInGO Share repository on GitHub"
                :style="{ color: isDark ? '#9ca3af' : '#6b7280' }"
            />
        </a>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useTheme } from "../../composables/useTheme";

const { isAuthenticated } = useAuth();
const { isDark } = useTheme();

// Helper methods for file icons - same as UploadsTab
const getFileIconPath = (filename: string) => {
  return getFileIcon(filename)
}

const getFileIconAltText = (filename: string) => {
  return getFileIconAlt(filename)
}

// Apple Hello typewriter effect with highlight - exactly like shadcn example
const words = [" secure", " instantly", " easy", " self-hosted"];
const currentWordIndex = ref(0);
const displayedChars = ref(0);
const isHighlighted = ref(false);

const longestWord = computed(() => {
    return words.reduce((a, b) => a.length > b.length ? a : b);
});

// Apple Hello effect cycle with highlight
const startAppleEffect = async () => {
    const currentWord = words[currentWordIndex.value];
    
    // Reset states
    displayedChars.value = 0;
    isHighlighted.value = false;
    
    // Type out characters one by one
    for (let i = 0; i <= currentWord.length; i++) {
        displayedChars.value = i;
        await new Promise(resolve => setTimeout(resolve, 100)); // 100ms per character
    }
    
    // Brief pause after typing
    await new Promise(resolve => setTimeout(resolve, 300));
    
    // Show progressive highlight
    isHighlighted.value = true;
    
    // Keep highlighted and text visible longer
    await new Promise(resolve => setTimeout(resolve, 2000)); // Keep visible for 2 seconds
    
    // Start erasing effect (highlight and text disappear together)
    for (let i = currentWord.length; i >= 0; i--) {
        displayedChars.value = i;
        await new Promise(resolve => setTimeout(resolve, 50)); // Faster erase
    }
    
    // Remove highlight when text is fully erased
    isHighlighted.value = false;
    
    // Brief pause before next word
    await new Promise(resolve => setTimeout(resolve, 300));
    
    // Move to next word
    currentWordIndex.value = (currentWordIndex.value + 1) % words.length;
    
    // Continue cycle
    startAppleEffect();
};

onMounted(() => {
    // Start after page loads
    setTimeout(() => {
        startAppleEffect();
    }, 1000);
});

onUnmounted(() => {
    // Cleanup if needed
});

// Refs
const fileInput = ref<HTMLInputElement | null>(null);
const uploadSection = ref<HTMLElement | null>(null);
const featuresSection = ref<HTMLElement | null>(null);
const { getSettings } = useAuth();
// State
const selectedFiles = ref<File[]>([]);
const isDragging = ref(false);
const isUploading = ref(false);
const progress = ref(0);
const uploadComplete = ref(false);
const uploadedUrl = ref("");
const message = ref({ text: "", type: "success" as "success" | "error" });
const previewingFiles = ref<Set<number>>(new Set());
const previewUrls = ref<Map<number, string>>(new Map());
const textPreviews = ref<Map<number, string>>(new Map());
const maxUploadSize = ref<number>(104857600); // Default 100 MB

const validityOptions = ref([
    { value: "7days", label: "7 Days", description: "One week" },
    { value: "1month", label: "1 Month", description: "30 days" },
    { value: "6months", label: "6 Months", description: "Half year" },
    { value: "1year", label: "1 Year", description: "12 months" },
    { value: "never", label: "Never", description: "Permanent" },
]);
const selectedValidity = ref("7days");
const maxAllowedValidity = ref("7days");

// Methods
const triggerFileInput = () => {
    fileInput.value?.click();
};

const onFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files) {
        selectedFiles.value = Array.from(target.files);
    }
};

const onDrop = (event: DragEvent) => {
    isDragging.value = false;
    if (event.dataTransfer?.files) {
        selectedFiles.value = Array.from(event.dataTransfer.files);
    }
};

const removeFile = (index: number) => {
    // Clean up preview URLs and refs when removing file
    if (previewingFiles.value.has(index)) {
        const url = previewUrls.value.get(index);
        if (url) URL.revokeObjectURL(url);
        previewUrls.value.delete(index);
        textPreviews.value.delete(index);
        previewingFiles.value.delete(index);
    }

    selectedFiles.value.splice(index, 1);

    // Re-index remaining files
    const newPreviewingFiles = new Set<number>();
    const newPreviewUrls = new Map<number, string>();
    const newTextPreviews = new Map<number, string>();

    previewingFiles.value.forEach((previewIndex) => {
        if (previewIndex > index) {
            newPreviewingFiles.add(previewIndex - 1);
            const url = previewUrls.value.get(previewIndex);
            if (url) newPreviewUrls.set(previewIndex - 1, url);
            const textPreview = textPreviews.value.get(previewIndex);
            if (textPreview) newTextPreviews.set(previewIndex - 1, textPreview);
        } else if (previewIndex < index) {
            newPreviewingFiles.add(previewIndex);
            const url = previewUrls.value.get(previewIndex);
            if (url) newPreviewUrls.set(previewIndex, url);
            const textPreview = textPreviews.value.get(previewIndex);
            if (textPreview) newTextPreviews.set(previewIndex, textPreview);
        }
    });

    previewingFiles.value = newPreviewingFiles;
    previewUrls.value = newPreviewUrls;
    textPreviews.value = newTextPreviews;
};

const togglePreview = (index: number) => {
    if (previewingFiles.value.has(index)) {
        // Hide preview
        previewingFiles.value.delete(index);
        const url = previewUrls.value.get(index);
        if (url) {
            URL.revokeObjectURL(url);
            previewUrls.value.delete(index);
        }
        textPreviews.value.delete(index);
    } else {
        // Show preview
        previewingFiles.value.add(index);
        const file = selectedFiles.value[index];
        const fileExtension = getFileExtension(file);

        if (["txt", "md", "json", "csv", "xml"].includes(fileExtension)) {
            // Handle text file preview
            const reader = new FileReader();
            reader.onload = (e) => {
                const content = e.target?.result as string;
                const truncatedContent =
                    content.length > 1000
                        ? content.substring(0, 1000) + "\n\n... (truncated)"
                        : content;
                textPreviews.value.set(index, truncatedContent);
            };
            reader.readAsText(file);
        } else {
            // Handle other file types (images, videos, audio)
            createPreviewUrl(file, index);
        }
    }
};

const createPreviewUrl = (file: File, index: number) => {
    const url = URL.createObjectURL(file);
    previewUrls.value.set(index, url);
};

const canPreview = (file: File): boolean => {
    const ext = getFileExtension(file);
    return [
        "mp4",
        "pdf",
        "jpg",
        "jpeg",
        "png",
        "gif",
        "webp",
        "txt",
        "md",
        "json",
        "csv",
        "xml",
        "torrent",
        "mp3",
        "wav",
        "flac",
    ].includes(ext);
};

const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
};

const getFileExtension = (file: File): string => {
    return file.name.split(".").pop()?.toLowerCase() || "";
};

const uploadFiles = async () => {
    if (selectedFiles.value.length === 0 || isUploading.value) return;

    isUploading.value = true;
    progress.value = 0;

    try {
        const formData = new FormData();
        selectedFiles.value.forEach((file) => {
            formData.append("files", file);
        });
        
        // Add validity/expiration setting to the form data
        formData.append("validity", selectedValidity.value);

        const response = await axios.post(
            "http://localhost:8080/upload",
            formData,
            {
                headers: {
                    "Content-Type": "multipart/form-data",
                    ...(isAuthenticated.value &&
                    localStorage.getItem("auth_token")
                        ? {
                              Authorization: `Bearer ${localStorage.getItem(
                                  "auth_token"
                              )}`,
                          }
                        : {}),
                },
                withCredentials: true,
                onUploadProgress: (progressEvent) => {
                    if (progressEvent.total) {
                        progress.value = Math.round(
                            (progressEvent.loaded * 100) / progressEvent.total
                        );
                    }
                },
            }
        );

        if (response.data.download_url) {
            progress.value = 100;
            uploadComplete.value = true;

            // Extract upload ID from download_url and create full URL
            const uploadId = response.data.download_url.split("/").pop();
            uploadedUrl.value = `${window.location.origin}/download/${uploadId}`;

            message.value = {
                text: "Files uploaded successfully!",
                type: "success",
            };

            // Scroll to center the upload success section
            setTimeout(() => {
                uploadSection.value?.scrollIntoView({
                    behavior: "smooth",
                    block: "center",
                });
            }, 100);

            // Clean up preview URLs
            previewUrls.value.forEach((url) => URL.revokeObjectURL(url));
            previewUrls.value.clear();
            textPreviews.value.clear();
            previewingFiles.value.clear();
        }
    } catch (error) {
        console.error("Upload error:", error);
        message.value = {
            text: "Upload failed. Please try again.",
            type: "error",
        };
        progress.value = 0;
    } finally {
        isUploading.value = false;

        // Clear message after 5 seconds instead of 3
        setTimeout(() => {
            message.value = { text: "", type: "success" };
            if (!isUploading.value) progress.value = 0;
        }, 5000);
    }
};

// Action buttons for upload success
const uploadNew = () => {
    uploadComplete.value = false;
    uploadedUrl.value = "";
    selectedFiles.value = [];
    progress.value = 0;
    message.value = { text: "", type: "success" };

    // Clean up any remaining preview URLs
    previewUrls.value.forEach((url) => URL.revokeObjectURL(url));
    previewUrls.value.clear();
    textPreviews.value.clear();
    previewingFiles.value.clear();
};

const copyLink = async () => {
    try {
        await navigator.clipboard.writeText(uploadedUrl.value);
        message.value = { text: "Link copied to clipboard!", type: "success" };
        setTimeout(() => {
            message.value = { text: "", type: "success" };
        }, 3000);
    } catch (err) {
        message.value = { text: "Failed to copy link", type: "error" };
        setTimeout(() => {
            message.value = { text: "", type: "success" };
        }, 3000);
    }
};

const openLink = () => {
    window.open(uploadedUrl.value, "_blank");
};

const scrollToUpload = () => {
    uploadSection.value?.scrollIntoView({ behavior: "smooth" });
};

const scrollToFeatures = () => {
    featuresSection.value?.scrollIntoView({ behavior: "smooth" });
};

const loadSettings = async () => {
    try {
        const settings = await getSettings();
        maxUploadSize.value = settings.maxUploadSize || 104857600;
        maxAllowedValidity.value = settings.maxValidity || "7days";

        // Filter validity options based on max allowed
        const validityOrder = ["7days", "1month", "6months", "1year", "never"];
        const maxIndex = validityOrder.indexOf(maxAllowedValidity.value);

        if (maxIndex !== -1) {
            const allowedOptions = validityOrder.slice(0, maxIndex + 1);
            validityOptions.value = validityOptions.value.filter((option) =>
                allowedOptions.includes(option.value)
            );
        }

        // Set default validity to first available option
        if (validityOptions.value.length > 0) {
            selectedValidity.value = validityOptions.value[0].value;
        }
    } catch (error) {
        console.error("Error loading settings:", error);
    }
};

onMounted(() => {
    // Scroll to top smoothly when page loads
    loadSettings();
    window.scrollTo({ top: 0, behavior: "smooth" });
});

onUnmounted(() => {
    // Cleanup handled in startRotationCycle
});
</script>

<style scoped>
/* Smooth scrolling for the entire page */
html {
    scroll-behavior: smooth;
}

@keyframes fade-in {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

@keyframes fade-in-delay {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

@keyframes fade-in-delay-2 {
    from {
        opacity: 0;
        transform: translateY(20px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
.shake {
  animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
  transform: translate3d(0, 0, 0);
}

@keyframes shake {
  10%,
  90% {
    transform: translate3d(-1px, 0, 0);
  }

  20%,
  80% {
    transform: translate3d(2px, 0, 0);
  }

  30%,
  50%,
  70% {
    transform: translate3d(-4px, 0, 0);
  }

  40%,
  60% {
    transform: translate3d(4px, 0, 0);
  }
}

.animate-gradient {
  background-size: 200% 200%;
  animation: gradient-move 3s linear infinite;
}


.animate-fade-in {
    animation: fade-in 0.8s ease-out;
}

.animate-fade-in-delay {
    animation: fade-in-delay 0.8s ease-out 0.2s both;
}

.animate-fade-in-delay-2 {
    animation: fade-in-delay-2 0.8s ease-out 0.4s both;
}


@keyframes liquid-glow {
    0%,
    100% {
        filter: drop-shadow(0 0 10px #ef444460) drop-shadow(0 0 20px #ef444440);
    }
    50% {
        filter: drop-shadow(0 0 15px #ef444480) drop-shadow(0 0 30px #ef444460);
    }
}


@keyframes gradient-move {
  0% { background-position: 0% 50%; }
  100% { background-position: 100% 50%; }
}

/* Clean Modern Highlight */
.clean-highlight-sweep {
  background: linear-gradient(
    90deg,
    rgba(59, 130, 246, 0.4) 0%,
    rgba(147, 51, 234, 0.45) 30%,
    rgba(236, 72, 153, 0.4) 60%,
    rgba(251, 146, 60, 0.4) 100%
  );
  transform: scaleX(0);
  transform-origin: left center;
  animation: clean-sweep 0.8s ease-out forwards;
  border-radius: 6px;
  box-shadow: 0 0 20px rgba(147, 51, 234, 0.2);
}

@keyframes clean-sweep {
  0% {
    transform: scaleX(0);
    opacity: 0;
  }
  30% {
    opacity: 1;
  }
  100% {
    transform: scaleX(1);
    opacity: 0.9;
  }
}

/* Component-specific HaloHandletter font */
@font-face {
  font-family: 'HaloHandletter-Local';
  src: url('/HaloHandletter.otf') format('opentype');
  font-weight: normal;
  font-style: normal;
  font-display: swap;
}

/* Component-specific Crima font */


.halo-handletter-local {
  font-family: 'HaloHandletter-Local', 'Brush Script MT', 'Segoe Print', cursive !important;
}

/* Force font for animated text */
.animated-text-container {
  font-family: 'Crima-Local', 'HaloHandletter-Local', 'HaloHandletter', 'Comic Sans MS', cursive !important;
  letter-spacing: 0.065em;
}

.animated-text-container * {
  font-family: inherit !important;
}


</style>
