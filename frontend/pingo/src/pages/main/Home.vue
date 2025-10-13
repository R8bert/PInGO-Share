<template>
    <div class="min-h-screen xp-desktop">
        <!-- Windows XP Wallpaper Background -->
        <div class="xp-bliss-bg"></div>
        
        <!-- Desktop Icons (Left Side) -->
        <div class="desktop-icons">
            <div class="desktop-icon">
                <div class="icon-image">
                    <svg class="w-12 h-12 text-blue-500" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M20 6h-8l-2-2H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2zm-1 12H5c-.55 0-1-.45-1-1V9c0-.55.45-1 1-1h14c.55 0 1 .45 1 1v8c0 .55-.45 1-1 1z"/>
                    </svg>
                </div>
                <span class="icon-label">My Files</span>
            </div>
            
            <div class="desktop-icon">
                <div class="icon-image">
                    <svg class="w-12 h-12 text-green-600" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                    </svg>
                </div>
                <span class="icon-label">Shared</span>
            </div>
            
            <div class="desktop-icon">
                <div class="icon-image">
                    <svg class="w-12 h-12 text-yellow-500" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M19.35 10.04C18.67 6.59 15.64 4 12 4 9.11 4 6.6 5.64 5.35 8.04 2.34 8.36 0 10.91 0 14c0 3.31 2.69 6 6 6h13c2.76 0 5-2.24 5-5 0-2.64-2.05-4.78-4.65-4.96z"/>
                    </svg>
                </div>
                <span class="icon-label">Upload</span>
            </div>
        </div>

        <!-- Main Upload Window -->
        <div class="xp-window-container">
            <div class="xp-window" ref="uploadWindow">
                <!-- Window Title Bar -->
                <div class="xp-titlebar">
                    <div class="flex items-center gap-2">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M19.35 10.04C18.67 6.59 15.64 4 12 4 9.11 4 6.6 5.64 5.35 8.04 2.34 8.36 0 10.91 0 14c0 3.31 2.69 6 6 6h13c2.76 0 5-2.24 5-5 0-2.64-2.05-4.78-4.65-4.96z"/>
                        </svg>
                        <span class="font-bold text-white text-sm">PInGO Share - File Upload</span>
                    </div>
                    <div class="flex items-center gap-1">
                        <button class="xp-titlebar-btn minimize">
                            <div class="w-2 h-0.5 bg-current"></div>
                        </button>
                        <button class="xp-titlebar-btn maximize">
                            <div class="w-2 h-2 border border-current"></div>
                        </button>
                        <button class="xp-titlebar-btn close" @click="minimizeWindow">
                            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M6 18L18 6M6 6l12 12"></path>
                            </svg>
                        </button>
                    </div>
                </div>

                <!-- Menu Bar -->
                <div class="xp-menubar">
                    <button class="xp-menu-item">File</button>
                    <button class="xp-menu-item">Edit</button>
                    <button class="xp-menu-item">View</button>
                    <button class="xp-menu-item">Help</button>
                </div>

                <!-- Toolbar -->
                <div class="xp-toolbar">
                    <button class="xp-toolbar-btn" title="Back">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                        </svg>
                    </button>
                    <button class="xp-toolbar-btn" title="Forward">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                        </svg>
                    </button>
                    <div class="xp-separator"></div>
                    <button class="xp-toolbar-btn" title="Search">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                        </svg>
                    </button>
                    <button class="xp-toolbar-btn" title="Folders">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M10 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2h-8l-2-2z"/>
                        </svg>
                    </button>
                    <div class="flex-1"></div>
                    <button class="xp-view-btn active">
                        <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M4 14h4v-4H4v4zm0 5h4v-4H4v4zM4 9h4V5H4v4zm5 5h12v-4H9v4zm0 5h12v-4H9v4zM9 5v4h12V5H9z"/>
                        </svg>
                    </button>
                </div>

                <!-- Address Bar -->
                <div class="xp-addressbar">
                    <span class="text-gray-600 text-xs font-semibold">Address</span>
                    <div class="xp-address-input">
                        <svg class="w-4 h-4 text-gray-500" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M10 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2h-8l-2-2z"/>
                        </svg>
                        <span class="text-sm">My Computer/PInGO Share/Upload</span>
                    </div>
                    <button class="xp-go-btn">Go</button>
                </div>

                <!-- Content Area -->
                <div class="xp-content">
                    <!-- Success State -->
                    <div v-if="uploadSuccess && shareableLink" class="xp-success-panel">
                        <div class="flex items-center gap-4 mb-6">
                            <div class="xp-success-icon">
                                <svg class="w-12 h-12 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path>
                                </svg>
                            </div>
                            <div>
                                <h2 class="text-2xl font-bold text-gray-800 mb-1">Upload Complete!</h2>
                                <p class="text-gray-600">Your files have been uploaded successfully</p>
                            </div>
                        </div>

                        <div class="xp-panel mb-4">
                            <div class="xp-panel-header">Share Link</div>
                            <div class="xp-panel-content">
                                <div class="flex gap-2">
                                    <input 
                                        v-model="shareableLink" 
                                        readonly
                                        class="xp-input flex-1 font-mono text-sm"
                                    />
                                    <button @click="copyToClipboard" class="xp-button primary">
                                        <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                                        </svg>
                                        {{ copied ? 'Copied!' : 'Copy' }}
                                    </button>
                                </div>
                            </div>
                        </div>

                        <div class="flex gap-2">
                            <button @click="resetUpload" class="xp-button flex-1">
                                Upload More Files
                            </button>
                        </div>
                    </div>

                    <!-- Upload Area -->
                    <div v-else>
                        <!-- Task Panel (Left Side) -->
                        <div class="xp-task-panel">
                            <div class="xp-task-header">File and Folder Tasks</div>
                            <div class="xp-task-list">
                                <button @click="triggerFileSelect" class="xp-task-item">
                                    <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                                    </svg>
                                    <span>Upload files</span>
                                </button>
                                <button class="xp-task-item" @click="showPreview = selectedFiles.length > 0">
                                    <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                                    </svg>
                                    <span>Preview files</span>
                                </button>
                                <button class="xp-task-item" @click="clearFiles" :disabled="selectedFiles.length === 0">
                                    <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                                    </svg>
                                    <span>Clear all files</span>
                                </button>
                            </div>

                            <div class="xp-task-header mt-6">Other Places</div>
                            <div class="xp-task-list">
                                <button class="xp-task-item">
                                    <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 24 24">
                                        <path d="M10 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2h-8l-2-2z"/>
                                    </svg>
                                    <span>My Documents</span>
                                </button>
                                <button class="xp-task-item">
                                    <svg class="w-5 h-5 text-blue-600" fill="currentColor" viewBox="0 0 24 24">
                                        <path d="M4 6H2v14c0 1.1.9 2 2 2h14v-2H4V6zm16-4H8c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-1 9H9V9h10v2zm-4 4H9v-2h6v2zm4-8H9V5h10v2z"/>
                                    </svg>
                                    <span>Shared Folders</span>
                                </button>
                            </div>
                        </div>

                        <!-- Main Content -->
                        <div class="xp-main-content">
                            <!-- Drop Zone -->
                            <div
                                @drop="handleFileDrop"
                                @dragover="handleDragOver"
                                @dragleave="handleDragLeave"
                                @click="triggerFileSelect"
                                :class="[
                                    'xp-dropzone',
                                    isDragging ? 'xp-dropzone-active' : ''
                                ]"
                            >
                                <input
                                    ref="fileInput"
                                    type="file"
                                    multiple
                                    @change="handleFileSelect"
                                    class="hidden"
                                />
                                <CloudArrowUpIcon class="w-16 h-16 text-blue-500 mb-4" />
                                <h3 class="text-xl font-bold text-gray-700 mb-2">
                                    {{ isDragging ? 'Drop files here' : 'Upload Files' }}
                                </h3>
                                <p class="text-gray-600 text-sm mb-4">
                                    Drag and drop files here or click to browse
                                </p>
                                <button class="xp-button primary">
                                    Browse Files
                                </button>
                            </div>

                            <!-- File List -->
                            <div v-if="selectedFiles.length > 0" class="xp-file-list">
                                <div class="xp-list-header">
                                    <span class="flex-1">Name</span>
                                    <span class="w-24 text-right">Size</span>
                                    <span class="w-24 text-right">Type</span>
                                    <span class="w-16"></span>
                                </div>
                                <div class="xp-list-item" v-for="(file, index) in selectedFiles" :key="index">
                                    <div class="flex items-center gap-2 flex-1 min-w-0">
                                        <img :src="getFileIconPath(file.name)" :alt="getFileIconAltText(file.name)" class="w-6 h-6" />
                                        <span class="text-sm truncate">{{ file.name }}</span>
                                    </div>
                                    <span class="w-24 text-right text-sm text-gray-600">{{ formatFileSize(file.size) }}</span>
                                    <span class="w-24 text-right text-sm text-gray-600">{{ getFileExtension(file.name) }}</span>
                                    <button @click="removeFile(index)" class="w-16 text-right text-red-600 hover:text-red-800">
                                        <XMarkIcon class="w-5 h-5 ml-auto" />
                                    </button>
                                </div>
                            </div>

                            <!-- Settings Panel -->
                            <div v-if="selectedFiles.length > 0" class="xp-panel mt-4">
                                <div class="xp-panel-header">Upload Settings</div>
                                <div class="xp-panel-content space-y-4">
                                    <div>
                                        <label class="block text-sm font-semibold text-gray-700 mb-2">Link Expiration</label>
                                        <select v-model="selectedValidity" class="xp-select">
                                            <option value="1h">1 hour</option>
                                            <option value="1d">1 day</option>
                                            <option value="7d">7 days</option>
                                            <option value="30d">30 days</option>
                                        </select>
                                    </div>

                                    <!-- Upload Progress -->
                                    <div v-if="uploading" class="space-y-2">
                                        <div class="flex justify-between text-sm">
                                            <span class="text-gray-700 font-semibold">Uploading...</span>
                                            <span class="text-blue-600 font-bold">{{ uploadProgress }}%</span>
                                        </div>
                                        <div class="xp-progress-bar">
                                            <div class="xp-progress-fill" :style="{ width: uploadProgress + '%' }"></div>
                                        </div>
                                    </div>

                                    <!-- Action Buttons -->
                                    <div class="flex gap-2 pt-2">
                                        <button 
                                            @click="uploadFiles" 
                                            :disabled="uploading || selectedFiles.length === 0"
                                            class="xp-button primary flex-1"
                                        >
                                            <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                                            </svg>
                                            {{ uploading ? 'Uploading...' : 'Upload Now' }}
                                        </button>
                                    </div>

                                    <!-- Error Message -->
                                    <div v-if="errorMessage" class="xp-error-box">
                                        <svg class="w-5 h-5 text-red-600 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                                            <path fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z" clip-rule="evenodd"/>
                                        </svg>
                                        <div>
                                            <p class="font-semibold text-red-800">Error</p>
                                            <p class="text-sm text-red-700">{{ errorMessage }}</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Status Bar -->
                <div class="xp-statusbar">
                    <div class="flex items-center gap-4">
                        <span class="text-xs text-gray-600">
                            {{ selectedFiles.length }} file(s) selected
                        </span>
                        <span v-if="selectedFiles.length > 0" class="text-xs text-gray-600">
                            Total: {{ totalFileSize }}
                        </span>
                    </div>
                    <div class="flex items-center gap-2">
                        <svg class="w-4 h-4 text-gray-500" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                        </svg>
                        <span class="text-xs text-gray-600">Ready</span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Preview Window (Modal) -->
        <div v-if="showPreview" class="xp-modal-overlay" @click="showPreview = false">
            <div class="xp-window xp-modal-window" @click.stop>
                <!-- Window Title Bar -->
                <div class="xp-titlebar">
                    <div class="flex items-center gap-2">
                        <svg class="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 24 24">
                            <path d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                            <path d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                        </svg>
                        <span class="font-bold text-white text-sm">File Preview</span>
                    </div>
                    <div class="flex items-center gap-1">
                        <button class="xp-titlebar-btn close" @click="showPreview = false">
                            <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M6 18L18 6M6 6l12 12"></path>
                            </svg>
                        </button>
                    </div>
                </div>

                <!-- Content -->
                <div class="xp-modal-content">
                    <div class="space-y-2 max-h-96 overflow-y-auto">
                        <div v-for="(file, index) in selectedFiles" :key="index" class="xp-preview-item">
                            <div class="flex items-center gap-3 flex-1">
                                <img :src="getFileIconPath(file.name)" :alt="getFileIconAltText(file.name)" class="w-10 h-10" />
                                <div class="flex-1 min-w-0">
                                    <p class="font-semibold text-gray-800 truncate">{{ file.name }}</p>
                                    <p class="text-sm text-gray-600">{{ formatFileSize(file.size) }} • {{ file.type || 'Unknown type' }}</p>
                                    <p class="text-xs text-gray-500">Modified: {{ new Date(file.lastModified).toLocaleDateString() }}</p>
                                </div>
                            </div>
                            <button @click="removeFile(index)" class="xp-button">Remove</button>
                        </div>
                    </div>
                    
                    <div class="flex justify-end gap-2 mt-4 pt-4 border-t border-gray-300">
                        <button @click="showPreview = false" class="xp-button primary">Close</button>
                    </div>
                </div>
            </div>
        </div>

        <!-- Taskbar -->
        <div class="xp-taskbar">
            <button class="xp-start-button">
                <svg class="w-6 h-6 text-white" viewBox="0 0 88 88" fill="currentColor">
                    <path d="M0 12.402l35.687-4.86.016 34.423-35.67.203zm35.67 33.529l.028 34.453L.028 75.48.026 45.7zm4.326-39.025L87.314 0v41.527l-47.318.376zm47.329 39.349l-.011 41.34-47.318-6.678-.066-34.739z"/>
                </svg>
                <span class="ml-2 font-bold">start</span>
            </button>
            
            <div class="xp-taskbar-buttons">
                <button class="xp-taskbar-app active">
                    <svg class="w-4 h-4 text-white mr-2" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M19.35 10.04C18.67 6.59 15.64 4 12 4 9.11 4 6.6 5.64 5.35 8.04 2.34 8.36 0 10.91 0 14c0 3.31 2.69 6 6 6h13c2.76 0 5-2.24 5-5 0-2.64-2.05-4.78-4.65-4.96z"/>
                    </svg>
                    <span class="text-white text-sm font-semibold">File Upload</span>
                </button>
            </div>
            
            <div class="xp-system-tray">
                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.536 8.464a5 5 0 010 7.072m2.828-9.9a9 9 0 010 12.728M5.586 15H4a1 1 0 01-1-1v-4a1 1 0 011-1h1.586l4.707-4.707C10.923 3.663 12 4.109 12 5v14c0 .891-1.077 1.337-1.707.707L5.586 15z"></path>
                </svg>
                <span class="xp-clock">{{ currentTime }}</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useIcons } from "../../composables/useIcons";
import { useSettings } from "../../composables/useSettings";
import { getApiUrl } from "../../utils/apiUtils";
import axios from "axios";
import {
    CloudArrowUpIcon,
    XMarkIcon,
} from "@heroicons/vue/24/outline";

const { isAuthenticated } = useAuth();
const { getFileIcon, getFileIconAlt } = useIcons();
const { fetchSettings } = useSettings();

// Refs
const uploadWindow = ref<HTMLElement | null>(null);
const fileInput = ref<HTMLInputElement | null>(null);
const particlesContainer = ref<HTMLElement | null>(null);

// Helper methods for file icons
const getFileIconPath = (filename: string) => {
    return getFileIcon(filename);
};

const getFileIconAltText = (filename: string) => {
    return getFileIconAlt(filename);
};

// State
const selectedFiles = ref<File[]>([]);
const isDragging = ref(false);
const uploading = ref(false);
const uploadProgress = ref(0);
const uploadSuccess = ref(false);
const shareableLink = ref("");
const errorMessage = ref("");
const copied = ref(false);
const selectedValidity = ref("7d");
const showPreview = ref(false);
const currentTime = ref("");

// Computed properties
const totalFileSize = computed(() => {
    const total = selectedFiles.value.reduce((sum, file) => sum + file.size, 0);
    return formatFileSize(total);
});

// Clock update
const updateClock = () => {
    const now = new Date();
    currentTime.value = now.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' });
};

// Fetch settings on component mount
onMounted(() => {
    fetchSettings();
    updateClock();
    setInterval(updateClock, 1000);
});

// Cleanup
onBeforeUnmount(() => {
    if (particlesContainer.value) {
        particlesContainer.value.innerHTML = '';
    }
});

// Methods
const triggerFileSelect = () => {
    fileInput.value?.click();
};

const handleFileSelect = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files) {
        selectedFiles.value = Array.from(target.files);
        uploadSuccess.value = false;
        errorMessage.value = "";
    }
};

const handleFileDrop = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = false;
    
    if (event.dataTransfer?.files) {
        selectedFiles.value = Array.from(event.dataTransfer.files);
        uploadSuccess.value = false;
        errorMessage.value = "";
    }
};

const handleDragOver = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = true;
};

const handleDragLeave = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = false;
};

const removeFile = (index: number) => {
    selectedFiles.value.splice(index, 1);
    if (selectedFiles.value.length === 0) {
        uploadSuccess.value = false;
        errorMessage.value = "";
    }
};

const clearFiles = () => {
    selectedFiles.value = [];
    uploadSuccess.value = false;
    errorMessage.value = "";
};

const resetUpload = () => {
    selectedFiles.value = [];
    uploadSuccess.value = false;
    shareableLink.value = "";
    errorMessage.value = "";
    uploadProgress.value = 0;
};

const formatFileSize = (bytes: number): string => {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB", "TB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
};

const getFileExtension = (filename: string): string => {
    const ext = filename.split('.').pop();
    return ext ? ext.toUpperCase() : 'FILE';
};

const minimizeWindow = () => {
    // Window minimize animation (placeholder)
    console.log('Window minimized');
};

const uploadFiles = async () => {
    if (selectedFiles.value.length === 0 || uploading.value) return;

    uploading.value = true;
    uploadProgress.value = 0;
    errorMessage.value = "";

    try {
        const formData = new FormData();
        selectedFiles.value.forEach((file) => {
            formData.append("files", file);
        });
        
        formData.append("validity", selectedValidity.value);

        const response = await axios.post(
            getApiUrl("/upload"),
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
                        uploadProgress.value = Math.round(
                            (progressEvent.loaded * 100) / progressEvent.total
                        );
                    }
                },
            }
        );

        if (response.data.download_url) {
            uploadProgress.value = 100;
            
            setTimeout(() => {
                uploadSuccess.value = true;
                const uploadId = response.data.download_url.split("/").pop();
                shareableLink.value = `${window.location.origin}/download/${uploadId}`;
            }, 500);
        }
    } catch (error: any) {
        console.error("Upload error:", error);
        if (error.response?.status === 413) {
            errorMessage.value = "File too large. Maximum file size is 100MB per file.";
        } else {
            errorMessage.value = error.response?.data?.error || "Upload failed. Please try again.";
        }
        uploadProgress.value = 0;
    } finally {
        uploading.value = false;
    }
};

const copyToClipboard = async () => {
    try {
        await navigator.clipboard.writeText(shareableLink.value);
        copied.value = true;
        
        setTimeout(() => {
            copied.value = false;
        }, 2000);
    } catch (err) {
        console.error('Failed to copy:', err);
    }
};
</script>

<style scoped>
</style>
