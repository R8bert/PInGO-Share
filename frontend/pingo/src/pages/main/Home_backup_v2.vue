<template>
    <div
        class="min-h-screen transition-colors duration-300"
        :style="{ backgroundColor: isDark ? '#000000' : '#ffffff' }"
    >
        <!-- WebGL Background -->
        <div class="fixed inset-0 pointer-events-none overflow-hidden z-0">
            <WebGLBackground
                :use-settings-color="true"
                :noise-intensity="isDark ? 0.03 : 0.01"
                :scanline-intensity="isDark ? 0.1 : 0.05"
                :speed="0.3"
                :scanline-frequency="isDark ? 0.5 : 0.2"
                :warp-amount="0.1"
                :resolution-scale="0.8"
                class="opacity-40"
            />
        </div>

        <!-- Main Content -->
        <main class="relative z-10">
            <!-- Hero Section -->
            <section
                class="relative min-h-screen flex items-center justify-center px-4 sm:px-6 lg:px-8"
            >
                <div class="max-w-7xl mx-auto w-full">
                    <div class="text-center space-y-8">
                        <!-- Badge -->
                        <div
                            class="inline-flex items-center space-x-2 px-4 py-2 rounded-full border backdrop-blur-sm animate-fade-in"
                            :style="{
                                backgroundColor: isDark
                                    ? 'rgba(255,255,255,0.05)'
                                    : 'rgba(0,0,0,0.05)',
                                borderColor: isDark
                                    ? 'rgba(255,255,255,0.1)'
                                    : 'rgba(0,0,0,0.1)',
                                color: isDark ? '#a1a1aa' : '#71717a',
                            }"
                        >
                            <div
                                class="w-2 h-2 rounded-full bg-green-500 animate-pulse"
                            ></div>
                            <span class="text-sm font-medium"
                                >Secure & Fast File Transfer</span
                            >
                        </div>

                        <!-- Main Headline -->
                        <div class="space-y-6 animate-fade-in-delay">
                            <h1
                                class="text-5xl sm:text-6xl lg:text-7xl font-bold tracking-tight"
                            >
                                <span
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                    >Share files</span
                                >
                                <br />
                                <RotatingText
                                    :texts="[
                                        'Secure',
                                        'Instantly',
                                        'Free',
                                        'Easy!',
                                    ]"
                                    mainClassName="inline-block px-2 sm:px-2 md:px-3 bg-green-300 text-black overflow-hidden py-0.5 sm:py-1 md:py-2 justify-center rounded-lg"
                                    :initial="{ y: '100%' }"
                                    :animate="{ y: 0 }"
                                    :exit="{ y: '-120%' }"
                                    :staggerDuration="0.025"
                                    splitLevelClassName="overflow-hidden pb-0.5 sm:pb-1 md:pb-1"
                                    :transition="{
                                        type: 'spring',
                                        damping: 30,
                                        stiffness: 400,
                                    }"
                                    :rotationInterval="3500"
                                />
                            </h1>

                            <p
                                class="text-xl sm:text-2xl max-w-3xl mx-auto leading-relaxed"
                                :style="{
                                    color: isDark ? '#a1a1aa' : '#71717a',
                                }"
                            >
                                Transfer files of any size securely and
                                instantly. No registration required, no limits,
                                just pure simplicity.
                            </p>
                        </div>

                        <!-- CTA Buttons -->
                        <div
                            class="flex flex-col sm:flex-row items-center justify-center gap-4 animate-fade-in-delay-2"
                        >
                            <button
                                @click="scrollToUpload"
                                class="group relative px-8 py-4 text-white font-semibold rounded-2xl transition-all duration-300 hover:scale-105 hover:shadow-2xl"
                                :style="{ background: buttonGradient }"
                            >
                                <span class="relative z-10"
                                    >Start Uploading</span
                                >
                                <div
                                    class="absolute inset-0 rounded-2xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                                    :style="{ background: hoverGradient }"
                                ></div>
                            </button>

                            <button
                                @click="scrollToFeatures"
                                class="px-8 py-4 border-2 font-semibold rounded-2xl transition-all duration-300 hover:scale-105 backdrop-blur-sm"
                                :style="{
                                    borderColor: isDark
                                        ? 'rgba(255,255,255,0.2)'
                                        : 'rgba(0,0,0,0.2)',
                                    color: isDark ? '#ffffff' : '#000000',
                                }"
                            >
                                Learn More
                            </button>
                        </div>
                    </div>
                </div>

                <!-- Scroll Indicator -->
                <div
                    class="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce"
                >
                    <div
                        class="w-6 h-10 border-2 rounded-full flex justify-center"
                        :style="{
                            borderColor: isDark
                                ? 'rgba(255,255,255,0.3)'
                                : 'rgba(0,0,0,0.3)',
                        }"
                    >
                        <div
                            class="w-1 h-3 bg-gradient-to-b from-blue-500 to-purple-500 rounded-full mt-2 animate-pulse"
                        ></div>
                    </div>
                </div>
            </section>

            <!-- Upload Section -->
            <section
                ref="uploadSection"
                class="relative py-20 px-4 sm:px-6 lg:px-8"
            >
                <div class="max-w-4xl mx-auto">
                    <!-- Upload Card -->
                    <div class="relative">
                        <!-- Background Glow -->
                        <div
                            class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-3xl blur-3xl"
                        ></div>

                        <!-- Main Card -->
                        <div
                            class="relative backdrop-blur-xl border rounded-3xl p-8 sm:p-12"
                            :style="{
                                backgroundColor: isDark
                                    ? 'rgba(255,255,255,0.05)'
                                    : 'rgba(255,255,255,0.8)',
                                borderColor: isDark
                                    ? 'rgba(255,255,255,0.1)'
                                    : 'rgba(0,0,0,0.1)',
                            }"
                        >
                            <!-- Upload Success Screen -->
                            <div v-if="uploadComplete" class="text-center">
                                <div
                                    class="w-20 h-20 mx-auto mb-6 bg-gradient-to-br from-green-400 to-emerald-600 rounded-full flex items-center justify-center"
                                >
                                    <svg
                                        class="w-10 h-10 text-white"
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

                                <h3
                                    class="text-2xl font-semibold text-gray-900 dark:text-white mb-2"
                                >
                                    Upload Complete!
                                </h3>
                                <p
                                    class="text-gray-600 dark:text-gray-300 mb-8"
                                >
                                    Your files have been successfully uploaded
                                    and are ready to share.
                                </p>

                                <!-- Action Buttons -->
                                <div class="space-y-3">
                                    <!-- Upload New -->
                                    <button
                                        @click="openLink"
                                        class="w-full bg-gradient-to-r from-green-500 to-green-600 hover:from-blue-600 hover:to-purple-700 backdrop-blur-sm border border-white/30 text-gray-800 dark:text-white font-medium py-3 px-6 rounded-xl transition-all duration-200 flex items-center justify-center gap-2"
                                    >
                                        <svg
                                            class="w-5 h-5"
                                            fill="none"
                                            stroke="currentColor"
                                            viewBox="0 0 24 24"
                                        >
                                            <path
                                                stroke-linecap="round"
                                                stroke-linejoin="round"
                                                stroke-width="2"
                                                d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                                            ></path>
                                        </svg>
                                        Open Download Page
                                    </button>
                                    <button
                                        @click="copyLink"
                                        class="w-full bg-white/20 hover:bg-white/30 backdrop-blur-sm border border-white/30 text-gray-800 dark:text-white font-medium py-3 px-6 rounded-xl transition-all duration-200 flex items-center justify-center gap-2"
                                    >
                                        <svg
                                            class="w-5 h-5"
                                            fill="none"
                                            stroke="currentColor"
                                            viewBox="0 0 24 24"
                                        >
                                            <path
                                                stroke-linecap="round"
                                                stroke-linejoin="round"
                                                stroke-width="2"
                                                d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                                            ></path>
                                        </svg>
                                        Copy Share Link
                                    </button>
                                    <button
                                        @click="uploadNew"
                                        class="w-full bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-medium py-3 px-6 rounded-xl transition-all duration-200 transform hover:scale-105 flex items-center justify-center gap-2"
                                    >
                                        <svg
                                            class="w-5 h-5"
                                            fill="none"
                                            stroke="currentColor"
                                            viewBox="0 0 24 24"
                                        >
                                            <path
                                                stroke-linecap="round"
                                                stroke-linejoin="round"
                                                stroke-width="2"
                                                d="M12 4v16m8-8H4"
                                            ></path>
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
                                            class="relative rounded-2xl p-6 flex items-center justify-center"
                                            :style="{
                                                background: primaryGradient,
                                            }"
                                        >
                                            <CloudArrowUpIcon
                                                class="w-12 h-12 text-white"
                                            />
                                        </div>
                                    </div>

                                    <!-- Upload Text -->
                                    <div class="space-y-2">
                                        <h3
                                            class="text-2xl font-bold"
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
                                        class="grid grid-cols-1 sm:grid-cols-3 gap-4 pt-6"
                                    >
                                        <div
                                            class="flex items-center justify-center space-x-2 p-3 rounded-xl"
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
                                                class="text-sm font-medium"
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
                                            class="flex items-center justify-center space-x-2 p-3 rounded-xl"
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
                                                class="text-sm font-medium"
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
                                            class="flex items-center justify-center space-x-2 p-3 rounded-xl"
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
                                                class="text-sm font-medium"
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
                                    class="absolute inset-0 rounded-2xl flex items-center justify-center z-20 bg-blue-500/20 backdrop-blur-sm"
                                >
                                    <div
                                        class="text-2xl font-bold text-blue-600"
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
                                        class="block text-sm font-medium mb-3 transition-colors duration-300"
                                        :style="{
                                            color: isDark
                                                ? '#d1d5db'
                                                : '#374151',
                                        }"
                                    >
                                        File expiration
                                    </label>

                                    <div
                                        class="grid grid-cols-2 sm:grid-cols-3 gap-3"
                                    >
                                        <label
                                            v-for="option in validityOptions"
                                            :key="option.value"
                                            class="relative flex items-center justify-center px-3 py-2 rounded-lg cursor-pointer transition-all text-sm"
                                            :style="{
                                                borderWidth: '2px',
                                                borderColor:
                                                    selectedValidity ===
                                                    option.value
                                                        ? isDark
                                                            ? '#60a5fa'
                                                            : '#3b82f6'
                                                        : isDark
                                                        ? '#4b5563'
                                                        : '#d1d5db',
                                                backgroundColor:
                                                    selectedValidity ===
                                                    option.value
                                                        ? isDark
                                                            ? '#1e3a8a20'
                                                            : '#dbeafe'
                                                        : isDark
                                                        ? '#1f2937'
                                                        : '#ffffff',
                                                color:
                                                    selectedValidity ===
                                                    option.value
                                                        ? isDark
                                                            ? '#93c5fd'
                                                            : '#1d4ed8'
                                                        : isDark
                                                        ? '#d1d5db'
                                                        : '#374151',
                                            }"
                                        >
                                            <input
                                                type="radio"
                                                v-model="selectedValidity"
                                                :value="option.value"
                                                class="sr-only"
                                            />
                                            <div class="text-center">
                                                <div class="font-medium">
                                                    {{ option.label }}
                                                </div>
                                                <div class="text-xs opacity-75">
                                                    {{ option.description }}
                                                </div>
                                            </div>
                                        </label>
                                    </div>
                                </div>

                                <!-- Additional Upload Options -->
                                <div class="mt-6 space-y-4">
                                    <!-- Password Protection -->
                                    <div>
                                        <label
                                            class="block text-sm font-medium mb-3 transition-colors duration-300"
                                            :style="{
                                                color: isDark
                                                    ? '#d1d5db'
                                                    : '#374151',
                                            }"
                                        >
                                            <svg class="w-4 h-4 inline mr-2" viewBox="0 0 20 20" fill="currentColor">
                                                <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" />
                                            </svg>
                                            Password Protection (optional)
                                        </label>
                                        <div class="relative">
                                            <input
                                                v-model="uploadPassword"
                                                :type="showPassword ? 'text' : 'password'"
                                                placeholder="Set a password to protect your files"
                                                class="w-full px-4 py-3 rounded-lg transition-all duration-200 pr-12"
                                                :style="{
                                                    backgroundColor: isDark
                                                        ? 'rgba(255,255,255,0.05)'
                                                        : 'rgba(0,0,0,0.05)',
                                                    borderWidth: '2px',
                                                    borderColor: isDark
                                                        ? '#4b5563'
                                                        : '#d1d5db',
                                                    color: isDark ? '#ffffff' : '#000000',
                                                }"
                                            />
                                            <button
                                                @click="showPassword = !showPassword"
                                                type="button"
                                                class="absolute right-3 top-1/2 -translate-y-1/2 p-2 rounded hover:bg-black/10 transition-colors"
                                            >
                                                <svg v-if="showPassword" class="w-5 h-5" :style="{ color: isDark ? '#9ca3af' : '#6b7280' }" viewBox="0 0 20 20" fill="currentColor">
                                                    <path fill-rule="evenodd" d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" />
                                                    <path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z" />
                                                </svg>
                                                <svg v-else class="w-5 h-5" :style="{ color: isDark ? '#9ca3af' : '#6b7280' }" viewBox="0 0 20 20" fill="currentColor">
                                                    <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                                                    <path fill-rule="evenodd" d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" />
                                                </svg>
                                            </button>
                                        </div>
                                    </div>

                                    <!-- Transfer Message -->
                                    <div>
                                        <label
                                            class="block text-sm font-medium mb-3 transition-colors duration-300"
                                            :style="{
                                                color: isDark
                                                    ? '#d1d5db'
                                                    : '#374151',
                                            }"
                                        >
                                            <svg class="w-4 h-4 inline mr-2" viewBox="0 0 20 20" fill="currentColor">
                                                <path fill-rule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" />
                                            </svg>
                                            Message (optional)
                                        </label>
                                        <textarea
                                            v-model="transferMessage"
                                            placeholder="Add a message for the recipient..."
                                            rows="3"
                                            class="w-full px-4 py-3 rounded-lg transition-all duration-200 resize-none"
                                            :style="{
                                                backgroundColor: isDark
                                                    ? 'rgba(255,255,255,0.05)'
                                                    : 'rgba(0,0,0,0.05)',
                                                borderWidth: '2px',
                                                borderColor: isDark
                                                    ? '#4b5563'
                                                    : '#d1d5db',
                                                color: isDark ? '#ffffff' : '#000000',
                                            }"
                                        ></textarea>
                                    </div>

                                    <!-- Email Notification -->
                                    <div>
                                        <label
                                            class="block text-sm font-medium mb-3 transition-colors duration-300"
                                            :style="{
                                                color: isDark
                                                    ? '#d1d5db'
                                                    : '#374151',
                                            }"
                                        >
                                            <svg class="w-4 h-4 inline mr-2" viewBox="0 0 20 20" fill="currentColor">
                                                <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                                                <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                                            </svg>
                                            Email Notification (optional)
                                        </label>
                                        <input
                                            v-model="recipientEmail"
                                            type="email"
                                            placeholder="recipient@example.com"
                                            class="w-full px-4 py-3 rounded-lg transition-all duration-200"
                                            :style="{
                                                backgroundColor: isDark
                                                    ? 'rgba(255,255,255,0.05)'
                                                    : 'rgba(0,0,0,0.05)',
                                                borderWidth: '2px',
                                                borderColor: isDark
                                                    ? '#4b5563'
                                                    : '#d1d5db',
                                                color: isDark ? '#ffffff' : '#000000',
                                            }"
                                        />
                                        <p
                                            class="mt-2 text-xs"
                                            :style="{
                                                color: isDark ? '#9ca3af' : '#6b7280',
                                            }"
                                        >
                                            The recipient will receive a notification when your files are ready
                                        </p>
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
                                <div class="grid gap-4">
                                    <div
                                        v-for="(file, index) in selectedFiles"
                                        :key="index"
                                        class="rounded-xl border transition-all duration-300"
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
                                            class="flex items-center justify-between p-4"
                                        >
                                            <div
                                                class="flex items-center space-x-3"
                                            >
                                                <div
                                                    class="w-10 h-10 bg-gradient-to-r from-blue-500 to-purple-500 rounded-lg flex items-center justify-center overflow-hidden"
                                                >
                                                    <span
                                                        class="text-white text-xs font-bold leading-none text-center px-1 truncate max-w-full"
                                                    >
                                                        {{
                                                            getFileExtension(
                                                                file
                                                            ).toUpperCase()
                                                                .length > 4
                                                                ? getFileExtension(
                                                                      file
                                                                  )
                                                                      .toUpperCase()
                                                                      .substring(
                                                                          0,
                                                                          3
                                                                      ) + "."
                                                                : getFileExtension(
                                                                      file
                                                                  ).toUpperCase()
                                                        }}
                                                    </span>
                                                </div>
                                                <div>
                                                    <p
                                                        class="font-medium"
                                                        :style="{
                                                            color: isDark
                                                                ? '#ffffff'
                                                                : '#000000',
                                                        }"
                                                    >
                                                        {{ file.name }}
                                                    </p>
                                                    <p
                                                        class="text-sm"
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
                                                class="flex items-center space-x-2"
                                            >
                                                <button
                                                    v-if="canPreview(file)"
                                                    @click="
                                                        togglePreview(index)
                                                    "
                                                    class="px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200 hover:scale-105"
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
                                                        class="w-4 h-4 mr-1 inline"
                                                    />
                                                    <EyeSlashIcon
                                                        v-else
                                                        class="w-4 h-4 mr-1 inline"
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
                                                    class="p-2 rounded-full hover:bg-red-500/20 text-red-500 transition-colors"
                                                >
                                                    <XMarkIcon
                                                        class="w-5 h-5"
                                                    />
                                                </button>
                                            </div>

                                            <!-- Expiration selector -->
                                        </div>

                                        <!-- Preview Section -->
                                        <div
                                            v-if="previewingFiles.has(index)"
                                            class="px-4 pb-4"
                                        >
                                            <div
                                                class="rounded-lg border-t pt-4"
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
                                                        class="w-full max-w-md mx-auto rounded-lg shadow-lg"
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
                                                        class="w-full max-w-md mx-auto rounded-lg shadow-lg"
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
                                                        class="w-full max-w-md mx-auto"
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
                                                            class="text-xs p-4 whitespace-pre-wrap overflow-auto max-h-32 w-full block"
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
                                    class="mt-6"
                                >
                                    <div
                                        class="flex justify-between text-sm mb-2"
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
                                        class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-3 overflow-hidden"
                                    >
                                        <div
                                            class="h-3 rounded-full transition-all duration-300"
                                            :style="{
                                                background: buttonGradient,
                                                width: progress + '%',
                                            }"
                                        ></div>
                                    </div>
                                </div>

                                <!-- Upload Button -->
                                <div v-if="!uploadComplete" class="pt-6">
                                    <button
                                        @click="uploadFiles"
                                        :disabled="isUploading"
                                        class="w-full py-4 text-white font-semibold rounded-2xl transition-all duration-300 hover:scale-[1.02] hover:shadow-2xl disabled:opacity-50 disabled:hover:scale-100"
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

                                    <!-- Upload Stats -->
                                    <div
                                        v-if="isUploading && progress > 0"
                                        class="mt-4 flex justify-between text-sm"
                                        :style="{
                                            color: isDark ? '#9ca3af' : '#6b7280',
                                        }"
                                    >
                                        <span class="flex items-center">
                                            <svg class="w-4 h-4 mr-1.5" viewBox="0 0 20 20" fill="currentColor">
                                                <path fill-rule="evenodd" d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" />
                                            </svg>
                                            {{ uploadSpeed }}
                                        </span>
                                        <span class="flex items-center">
                                            <svg class="w-4 h-4 mr-1.5" viewBox="0 0 20 20" fill="currentColor">
                                                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" />
                                            </svg>
                                            {{ timeRemaining }}
                                        </span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <!-- Features Section -->
            <section
                ref="featuresSection"
                class="relative py-20 px-4 sm:px-6 lg:px-8"
            >
                <div class="max-w-7xl mx-auto">
                    <div class="text-center mb-16">
                        <h2
                            class="text-4xl sm:text-5xl font-bold mb-6"
                            :style="{ color: isDark ? '#ffffff' : '#000000' }"
                        >
                            Why choose PInGO Share?
                        </h2>
                        <p
                            class="text-xl max-w-3xl mx-auto"
                            :style="{ color: isDark ? '#a1a1aa' : '#71717a' }"
                        >
                            Built with modern technology and user experience in
                            mind
                        </p>
                    </div>

                    <div
                        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"
                    >
                        <!-- Feature 1 -->
                        <div class="relative group">
                            <div
                                class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-purple-600/20 rounded-2xl blur-xl opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                            ></div>
                            <div
                                class="relative backdrop-blur-xl border rounded-2xl p-8 transition-all duration-300 group-hover:scale-105"
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
                                    class="w-16 h-16 bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl flex items-center justify-center mb-6"
                                >
                                    <svg
                                        class="w-8 h-8 text-white"
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
                                    class="text-xl font-bold mb-4"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    End-to-End Encryption
                                </h3>
                                <p
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
                                class="relative backdrop-blur-xl border rounded-2xl p-8 transition-all duration-300 group-hover:scale-105"
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
                                    class="w-16 h-16 bg-gradient-to-r from-green-500 to-blue-500 rounded-2xl flex items-center justify-center mb-6"
                                >
                                    <svg
                                        class="w-8 h-8 text-white"
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
                                    class="text-xl font-bold mb-4"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    Lightning Fast
                                </h3>
                                <p
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
                                class="relative backdrop-blur-xl border rounded-2xl p-8 transition-all duration-300 group-hover:scale-105"
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
                                    class="w-16 h-16 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl flex items-center justify-center mb-6"
                                >
                                    <svg
                                        class="w-8 h-8 text-white"
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
                                    class="text-xl font-bold mb-4"
                                    :style="{
                                        color: isDark ? '#ffffff' : '#000000',
                                    }"
                                >
                                    No Registration
                                </h3>
                                <p
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
import { ref, onMounted, onUnmounted } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useTheme } from "../../composables/useTheme";
import { useUIColors } from "../../composables/useUIColors";
import axios from "axios";
import {
    CloudArrowUpIcon,
    XMarkIcon,
    EyeIcon,
    EyeSlashIcon,
} from "@heroicons/vue/24/outline";
import WebGLBackground from "../../components/WebGLBackground.vue";
import RotatingText from "../../blocks/TextAnimations/RotatingText/RotatingText.vue";
import GlitchText from "../../blocks/TextAnimations/GlitchText/GlitchText.vue";
import ShinyText from "../../blocks/TextAnimations/ShinyText/ShinyText.vue";

const { isAuthenticated } = useAuth();
const { isDark } = useTheme();
const {
    primaryGradient,
    buttonGradient,
    hoverGradient,
    primaryColor,
    secondaryColor,
    accentColor,
} = useUIColors();

// Rotating text logic
const rotatingWords = ref([
    "effortlessly",
    "securely",
    "instantly",
    "freely",
    "seamlessly",
]);
const currentWordIndex = ref(0);
let rotationInterval: number | null = null;

const getWordStyle = (index: number) => {
    const colors = [
        "#ef4444", // effortlessly - red
        "#10b981", // securely - emerald
        "#8b5cf6", // instantly - purple
        "#3b82f6", // freely - blue
        "#ec4899", // seamlessly - pink
    ];

    const isActive = index === currentWordIndex.value;

    return {
        color: colors[index] || "#ef4444",
        textShadow: isActive
            ? `0 0 20px ${colors[index]}40, 0 0 40px ${colors[index]}20`
            : "none",
        filter: isActive ? `drop-shadow(0 0 10px ${colors[index]}60)` : "none",
    };
};

const startRotation = () => {
    rotationInterval = window.setInterval(() => {
        currentWordIndex.value =
            (currentWordIndex.value + 1) % rotatingWords.value.length;
    }, 2500);
};

const stopRotation = () => {
    if (rotationInterval) {
        clearInterval(rotationInterval);
        rotationInterval = null;
    }
};

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

// New upload options
const uploadPassword = ref('');
const transferMessage = ref('');
const recipientEmail = ref('');
const showPassword = ref(false);

// Upload statistics
const uploadSpeed = ref('0 MB/s');
const timeRemaining = ref('Calculating...');
const startTime = ref(0);
const lastLoaded = ref(0);
const lastTime = ref(0);
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
    startTime.value = Date.now();
    lastLoaded.value = 0;
    lastTime.value = Date.now();
    uploadSpeed.value = '0 MB/s';
    timeRemaining.value = 'Calculating...';

    try {
        const formData = new FormData();
        selectedFiles.value.forEach((file) => {
            formData.append("files", file);
        });

        // Add upload options
        if (uploadPassword.value) {
            formData.append("password", uploadPassword.value);
        }
        if (transferMessage.value) {
            formData.append("message", transferMessage.value);
        }
        if (recipientEmail.value) {
            formData.append("recipient_email", recipientEmail.value);
        }
        if (selectedValidity.value) {
            formData.append("validity_hours", selectedValidity.value.toString());
        }

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

                        // Calculate upload speed and time remaining
                        const currentTime = Date.now();
                        const timeElapsed = (currentTime - lastTime.value) / 1000; // seconds
                        
                        if (timeElapsed > 0.5) { // Update every 0.5 seconds
                            const bytesUploaded = progressEvent.loaded - lastLoaded.value;
                            const speed = bytesUploaded / timeElapsed; // bytes per second
                            const speedMB = (speed / (1024 * 1024)).toFixed(2);
                            uploadSpeed.value = `${speedMB} MB/s`;

                            const remainingBytes = progressEvent.total - progressEvent.loaded;
                            const remainingSeconds = remainingBytes / speed;
                            
                            if (remainingSeconds < 60) {
                                timeRemaining.value = `${Math.ceil(remainingSeconds)}s remaining`;
                            } else if (remainingSeconds < 3600) {
                                timeRemaining.value = `${Math.ceil(remainingSeconds / 60)}m remaining`;
                            } else {
                                timeRemaining.value = `${Math.ceil(remainingSeconds / 3600)}h remaining`;
                            }

                            lastLoaded.value = progressEvent.loaded;
                            lastTime.value = currentTime;
                        }
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

    // Start the rotating text animation
    startRotation();
});

onUnmounted(() => {
    // Clean up the rotation interval
    stopRotation();
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


</style>