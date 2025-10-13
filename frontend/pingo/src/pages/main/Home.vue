<template>
    <div class="min-h-screen relative overflow-hidden xp-desktop">
        <!-- Windows XP Bliss Background -->
        <div class="absolute inset-0 z-0 xp-background"></div>
        
        <!-- Floating animated clouds -->
        <div class="clouds-container">
            <div ref="cloud1" class="cloud cloud-1"></div>
            <div ref="cloud2" class="cloud cloud-2"></div>
            <div ref="cloud3" class="cloud cloud-3"></div>
        </div>

        <!-- Main Content -->
        <div class="relative z-40 flex flex-col lg:flex-row items-center justify-center min-h-screen px-6 lg:px-12 gap-8 py-20">
            <!-- Left Side - Welcome Panel -->
            <div ref="welcomePanel" class="flex-1 max-w-2xl">
                <div class="xp-window welcome-window">
                    <div class="xp-title-bar">
                        <div class="xp-title-bar-text">
                            <span class="xp-window-icon"></span>
                            Welcome to PInGO Share
                        </div>
                        <div class="xp-title-bar-controls">
                            <button class="xp-btn-minimize" title="Minimize"></button>
                            <button class="xp-btn-maximize" title="Maximize"></button>
                            <button class="xp-btn-close" title="Close"></button>
                        </div>
                    </div>
                    <div class="xp-window-content">
                        <div ref="welcomeText" class="welcome-content">
                            <h1 class="xp-title">Share Files Instantly</h1>
                            <p class="xp-subtitle">Fast, secure, and retro-cool file sharing</p>
                            <div class="xp-features">
                                <div ref="feature1" class="xp-feature-item">
                                    <div class="xp-icon-box">📁</div>
                                    <span>No registration needed</span>
                                </div>
                                <div ref="feature2" class="xp-feature-item">
                                    <div class="xp-icon-box">⚡</div>
                                    <span>Lightning fast uploads</span>
                                </div>
                                <div ref="feature3" class="xp-feature-item">
                                    <div class="xp-icon-box">🔒</div>
                                    <span>Secure & encrypted</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Right Side - Upload Window -->
            <div ref="uploadWindow" class="flex-shrink-0 w-full max-w-md lg:max-w-lg">
                <!-- Success State -->
                <div v-if="uploadSuccess && shareableLink" class="xp-window success-window">
                    <div class="xp-title-bar">
                        <div class="xp-title-bar-text">
                            <span class="xp-window-icon success-icon">✓</span>
                            Upload Complete
                        </div>
                        <div class="xp-title-bar-controls">
                            <button class="xp-btn-minimize" title="Minimize"></button>
                            <button class="xp-btn-maximize" title="Maximize"></button>
                            <button class="xp-btn-close" title="Close"></button>
                        </div>
                    </div>
                    <div class="xp-window-content">
                        <div class="success-content">
                            <div class="success-icon-box">
                                <div class="success-checkmark">✓</div>
                            </div>
                            <h3 class="success-title">Files uploaded successfully! 🎉</h3>
                            <p class="success-subtitle">Your files are ready to share</p>
                            
                            <div class="xp-control-group">
                                <label class="xp-label">Shareable Link:</label>
                                <div class="xp-input-group">
                                    <input
                                        v-model="shareableLink"
                                        readonly
                                        class="xp-input xp-input-readonly"
                                    />
                                    <button
                                        @click="copyToClipboard"
                                        class="xp-button xp-button-primary"
                                    >
                                        {{ copied ? '✓ Copied' : 'Copy' }}
                                    </button>
                                </div>
                            </div>
                            
                            <button
                                @click="resetUpload"
                                class="xp-button xp-button-secondary w-full"
                            >
                                Upload more files
                            </button>
                        </div>
                    </div>
                </div>

                <!-- Upload Form -->
                <div v-else class="xp-window upload-window">
                    <div class="xp-title-bar">
                        <div class="xp-title-bar-text">
                            <span class="xp-window-icon">📤</span>
                            File Upload Center
                        </div>
                        <div class="xp-title-bar-controls">
                            <button class="xp-btn-minimize" title="Minimize"></button>
                            <button class="xp-btn-maximize" title="Maximize"></button>
                            <button class="xp-btn-close" title="Close"></button>
                        </div>
                    </div>
                    <div class="xp-window-content">
                        <!-- Header -->
                        <div class="upload-header">
                            <h2 class="xp-heading">Share Your Files</h2>
                            <p class="xp-text">Drop files or click to browse</p>
                        </div>

                        <!-- Upload Area -->
                        <div class="upload-zone">
                            <!-- Drop Zone -->
                            <div
                                @drop="handleFileDrop"
                                @dragover="handleDragOver" 
                                @dragleave="handleDragLeave"
                                @click="triggerFileSelect"
                                ref="dropZone"
                                :class="[
                                    'xp-drop-zone',
                                    isDragging ? 'xp-drop-zone-active' : ''
                                ]"
                            >
                                <input
                                    ref="fileInput"
                                    type="file"
                                    multiple
                                    @change="handleFileSelect"
                                    class="hidden"
                                />
                                
                                <div class="drop-zone-content">
                                    <CloudArrowUpIcon class="drop-zone-icon" />
                                    <div>
                                        <p class="drop-zone-text">
                                            Drop files here or click to browse
                                        </p>
                                        <p class="drop-zone-subtext">
                                            Maximum file size: 100MB
                                        </p>
                                    </div>
                                </div>
                            </div>

                            <!-- Selected Files -->
                            <div v-if="selectedFiles.length > 0" class="files-section">
                                <div class="xp-group-box">
                                    <div class="xp-group-box-header">
                                        <span class="xp-group-box-title">
                                            Selected Files ({{ selectedFiles.length }})
                                        </span>
                                        <button @click="clearFiles" class="xp-link-button">
                                            Clear all
                                        </button>
                                    </div>
                                    
                                    <div class="xp-file-list">
                                        <div
                                            v-for="(file, index) in selectedFiles"
                                            :key="index"
                                            :ref="el => { if (el) fileRefs[index] = el }"
                                            class="xp-file-item"
                                        >
                                            <div class="xp-file-info">
                                                <img :src="getFileIconPath(file.name)" :alt="getFileIconAltText(file.name)" class="xp-file-icon" />
                                                <div class="xp-file-details">
                                                    <p class="xp-file-name">{{ file.name }}</p>
                                                    <p class="xp-file-size">{{ formatFileSize(file.size) }}</p>
                                                </div>
                                            </div>
                                            <button @click="removeFile(index)" class="xp-file-remove" title="Remove">
                                                <XMarkIcon class="w-4 h-4" />
                                            </button>
                                        </div>
                                    </div>
                                </div>

                                <!-- Expiration Settings -->
                                <div class="xp-control-group">
                                    <label class="xp-label">Link expires in:</label>
                                    <select v-model="selectedValidity" class="xp-select">
                                        <option value="1h">1 hour</option>
                                        <option value="1d">1 day</option>
                                        <option value="7d">7 days</option>
                                        <option value="30d">30 days</option>
                                    </select>
                                </div>

                                <!-- Upload Button -->
                                <button
                                    @click="uploadFiles"
                                    :disabled="uploading || selectedFiles.length === 0"
                                    class="xp-button xp-button-primary w-full"
                                    ref="uploadBtn"
                                >
                                    <span v-if="uploading" class="upload-progress">
                                        <span class="spinner"></span>
                                        Uploading... {{ uploadProgress }}%
                                    </span>
                                    <span v-else>Upload Files</span>
                                </button>
                            </div>

                            <!-- Default State Button -->
                            <button
                                v-if="selectedFiles.length === 0"
                                @click="triggerFileSelect"
                                class="xp-button xp-button-primary w-full"
                            >
                                Select Files
                            </button>

                            <!-- Error Message -->
                            <div v-if="errorMessage" class="xp-error-box">
                                <div class="xp-error-content">
                                    <span class="xp-error-icon">⚠️</span>
                                    <span class="xp-error-text">{{ errorMessage }}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Retro XP Taskbar -->
        <div class="xp-taskbar">
            <div class="xp-start-button" ref="startButton">
                <div class="xp-start-icon"></div>
                <span class="xp-start-text">Start</span>
            </div>
            <div class="xp-taskbar-items">
                <div class="xp-taskbar-item xp-taskbar-item-active">
                    <span class="xp-taskbar-icon">📤</span>
                    <span class="xp-taskbar-label">PInGO Share</span>
                </div>
            </div>
            <div class="xp-system-tray">
                <div class="xp-tray-time" ref="clockDisplay">{{ currentTime }}</div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import { useAuth } from "../../composables/useAuth";
import { useIcons } from "../../composables/useIcons";
import { useSettings } from "../../composables/useSettings";
import { getApiUrl } from "../../utils/apiUtils";
import axios from "axios";
import { animate, createTimeline } from "animejs";
import {
    CloudArrowUpIcon,
    XMarkIcon,
} from "@heroicons/vue/24/outline";

const { isAuthenticated } = useAuth();
const { getFileIcon, getFileIconAlt } = useIcons();
const { fetchSettings } = useSettings();

// Refs for animations
const welcomePanel = ref<HTMLElement | null>(null);
const uploadWindow = ref<HTMLElement | null>(null);
const welcomeText = ref<HTMLElement | null>(null);
const feature1 = ref<HTMLElement | null>(null);
const feature2 = ref<HTMLElement | null>(null);
const feature3 = ref<HTMLElement | null>(null);
const cloud1 = ref<HTMLElement | null>(null);
const cloud2 = ref<HTMLElement | null>(null);
const cloud3 = ref<HTMLElement | null>(null);
const dropZone = ref<HTMLElement | null>(null);
const uploadBtn = ref<HTMLElement | null>(null);
const startButton = ref<HTMLElement | null>(null);
const clockDisplay = ref<HTMLElement | null>(null);
const fileRefs = ref<Record<number, any>>({});

// Current time for XP taskbar clock
const currentTime = ref("");
let clockInterval: number | null = null;

const updateClock = () => {
    const now = new Date();
    currentTime.value = now.toLocaleTimeString('en-US', { 
        hour: 'numeric', 
        minute: '2-digit',
        hour12: true 
    });
};

// Helper methods for file icons
const getFileIconPath = (filename: string) => {
    return getFileIcon(filename);
};

const getFileIconAltText = (filename: string) => {
    return getFileIconAlt(filename);
};

// Initialize animations
const initAnimations = () => {
    // Animate clouds floating across the screen
    if (cloud1.value) {
        animate(cloud1.value, {
            translateX: ['-100%', '120%'],
            duration: 45000,
            easing: 'linear',
            loop: true,
            delay: 0
        });
    }
    
    if (cloud2.value) {
        animate(cloud2.value, {
            translateX: ['-100%', '120%'],
            duration: 60000,
            easing: 'linear',
            loop: true,
            delay: 5000
        });
    }
    
    if (cloud3.value) {
        animate(cloud3.value, {
            translateX: ['-100%', '120%'],
            duration: 50000,
            easing: 'linear',
            loop: true,
            delay: 10000
        });
    }

    // Animate welcome panel entrance
    if (welcomePanel.value) {
        animate(welcomePanel.value, {
            opacity: [0, 1],
            translateY: [-50, 0],
            scale: [0.9, 1],
            duration: 1000,
            easing: 'out(3)'
        });
    }

    // Animate upload window entrance
    if (uploadWindow.value) {
        animate(uploadWindow.value, {
            opacity: [0, 1],
            translateY: [50, 0],
            scale: [0.9, 1],
            duration: 1000,
            delay: 200,
            easing: 'out(3)'
        });
    }

    // Stagger feature items
    const timeline = createTimeline({
        defaults: {
            duration: 600
        }
    });

    if (feature1.value && feature2.value && feature3.value) {
        timeline
            .add(feature1.value, {
                opacity: [0, 1],
                translateX: [-30, 0],
                easing: 'out(2)'
            }, 400)
            .add(feature2.value, {
                opacity: [0, 1],
                translateX: [-30, 0],
                easing: 'out(2)'
            }, 200)
            .add(feature3.value, {
                opacity: [0, 1],
                translateX: [-30, 0],
                easing: 'out(2)'
            }, 200);
    }

    // Pulse animation for start button
    if (startButton.value) {
        animate(startButton.value, {
            scale: [1, 1.05, 1],
            duration: 2000,
            easing: 'inOut(2)',
            loop: true
        });
    }
};

// Animate file additions
const animateFileAddition = (index: number) => {
    setTimeout(() => {
        const element = fileRefs.value[index];
        if (element) {
            animate(element, {
                opacity: [0, 1],
                translateX: [-20, 0],
                duration: 400,
                easing: 'out(3)'
            });
        }
    }, 50);
};

// Fetch settings on component mount
onMounted(() => {
    fetchSettings();
    initAnimations();
    updateClock();
    clockInterval = window.setInterval(updateClock, 1000);
});

onUnmounted(() => {
    if (clockInterval) {
        clearInterval(clockInterval);
    }
});

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
const fileInput = ref<HTMLInputElement | null>(null);

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
        
        // Animate new files
        selectedFiles.value.forEach((_, index) => {
            animateFileAddition(index);
        });
    }
};

const handleFileDrop = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = false;
    
    if (event.dataTransfer?.files) {
        selectedFiles.value = Array.from(event.dataTransfer.files);
        uploadSuccess.value = false;
        errorMessage.value = "";
        
        // Animate dropped files with celebration
        if (dropZone.value) {
            animate(dropZone.value, {
                scale: [1.05, 1],
                duration: 300,
                easing: 'out(3)'
            });
        }
        
        selectedFiles.value.forEach((_, index) => {
            animateFileAddition(index);
        });
    }
};

const handleDragOver = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = true;
    
    // Subtle pulse animation on drag over
    if (dropZone.value) {
        animate(dropZone.value, {
            scale: [1, 1.02],
            duration: 200,
            easing: 'out(3)'
        });
    }
};

const handleDragLeave = (event: DragEvent) => {
    event.preventDefault();
    isDragging.value = false;
    
    if (dropZone.value) {
        animate(dropZone.value, {
            scale: [1.02, 1],
            duration: 200,
            easing: 'out(3)'
        });
    }
};

const removeFile = (index: number) => {
    const element = fileRefs.value[index];
    if (element) {
        animate(element, {
            opacity: [1, 0],
            translateX: [0, 20],
            duration: 300,
            easing: 'out(3)',
            onComplete: () => {
                selectedFiles.value.splice(index, 1);
                if (selectedFiles.value.length === 0) {
                    uploadSuccess.value = false;
                    errorMessage.value = "";
                }
            }
        });
    } else {
        selectedFiles.value.splice(index, 1);
        if (selectedFiles.value.length === 0) {
            uploadSuccess.value = false;
            errorMessage.value = "";
        }
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
    const sizes = ["Bytes", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
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
            uploadSuccess.value = true;

            const uploadId = response.data.download_url.split("/").pop();
            shareableLink.value = `${window.location.origin}/download/${uploadId}`;
        }
    } catch (error: any) {
        console.error("Upload error:", error);
        if (error.response?.status === 413) {
            errorMessage.value = "File too large. Maximum file size is 100MB.";
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
/* ============================================
   WINDOWS XP RETRO DESIGN SYSTEM
   ============================================ */

/* XP Desktop Background - Bliss Hill Style */
.xp-desktop {
    background: #5a98d7;
    font-family: 'Tahoma', 'Segoe UI', sans-serif;
}

.xp-background {
    background: linear-gradient(180deg, 
        #5a98d7 0%,
        #6ba4dd 20%,
        #7eb1e3 40%,
        #8ec0ea 60%,
        #9bcff0 80%,
        #a5d8f3 100%
    );
    position: relative;
}

.xp-background::before {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 40%;
    background: linear-gradient(to bottom, 
        transparent 0%,
        rgba(90, 170, 90, 0.3) 50%,
        rgba(100, 180, 100, 0.5) 100%
    );
}

/* Floating Clouds */
.clouds-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 60%;
    overflow: hidden;
    pointer-events: none;
    z-index: 1;
}

.cloud {
    position: absolute;
    background: rgba(255, 255, 255, 0.7);
    border-radius: 100px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.cloud::before,
.cloud::after {
    content: '';
    position: absolute;
    background: rgba(255, 255, 255, 0.7);
    border-radius: 100px;
}

.cloud-1 {
    width: 120px;
    height: 40px;
    top: 15%;
}

.cloud-1::before {
    width: 60px;
    height: 60px;
    top: -30px;
    left: 20px;
}

.cloud-1::after {
    width: 80px;
    height: 50px;
    top: -25px;
    right: 20px;
}

.cloud-2 {
    width: 100px;
    height: 35px;
    top: 35%;
}

.cloud-2::before {
    width: 50px;
    height: 50px;
    top: -25px;
    left: 15px;
}

.cloud-2::after {
    width: 70px;
    height: 45px;
    top: -20px;
    right: 15px;
}

.cloud-3 {
    width: 90px;
    height: 30px;
    top: 55%;
}

.cloud-3::before {
    width: 45px;
    height: 45px;
    top: -22px;
    left: 12px;
}

.cloud-3::after {
    width: 60px;
    height: 40px;
    top: -18px;
    right: 12px;
}

/* XP Window Chrome */
.xp-window {
    background: #ece9d8;
    border: 3px solid;
    border-color: #0054e3 #003c9d #003c9d #0054e3;
    border-radius: 8px 8px 0 0;
    box-shadow: 
        0 0 0 1px rgba(255, 255, 255, 0.5) inset,
        0 10px 30px rgba(0, 0, 0, 0.3),
        0 0 0 1px rgba(0, 0, 0, 0.2);
    overflow: hidden;
    opacity: 0;
}

.welcome-window,
.upload-window,
.success-window {
    opacity: 1;
}

/* XP Title Bar */
.xp-title-bar {
    background: linear-gradient(180deg, 
        #0054e3 0%,
        #0060ff 3%,
        #004ce5 50%,
        #003c9d 100%
    );
    padding: 4px 6px 4px 8px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    min-height: 30px;
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

.xp-title-bar-text {
    color: white;
    font-size: 11px;
    font-weight: bold;
    display: flex;
    align-items: center;
    gap: 6px;
    text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.5);
    letter-spacing: 0.3px;
}

.xp-window-icon {
    font-size: 14px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
}

.success-icon {
    background: #22c55e;
    border-radius: 3px;
    padding: 2px 4px;
}

/* XP Title Bar Controls */
.xp-title-bar-controls {
    display: flex;
    gap: 2px;
}

.xp-btn-minimize,
.xp-btn-maximize,
.xp-btn-close {
    width: 21px;
    height: 21px;
    border: 1px solid;
    border-color: rgba(255, 255, 255, 0.7) rgba(0, 0, 0, 0.4) rgba(0, 0, 0, 0.4) rgba(255, 255, 255, 0.7);
    cursor: pointer;
    position: relative;
    background: linear-gradient(180deg, #ebf2fd 0%, #c4ddfd 100%);
    border-radius: 2px;
    transition: all 0.1s;
}

.xp-btn-minimize:hover,
.xp-btn-maximize:hover,
.xp-btn-close:hover {
    background: linear-gradient(180deg, #fff 0%, #d8e9fc 100%);
    border-color: rgba(255, 255, 255, 0.9) rgba(0, 0, 0, 0.5) rgba(0, 0, 0, 0.5) rgba(255, 255, 255, 0.9);
}

.xp-btn-minimize:active,
.xp-btn-maximize:active,
.xp-btn-close:active {
    background: linear-gradient(180deg, #c4ddfd 0%, #ebf2fd 100%);
    border-color: rgba(0, 0, 0, 0.4) rgba(255, 255, 255, 0.7) rgba(255, 255, 255, 0.7) rgba(0, 0, 0, 0.4);
}

.xp-btn-minimize::before {
    content: '';
    position: absolute;
    bottom: 6px;
    left: 4px;
    right: 4px;
    height: 2px;
    background: #000;
}

.xp-btn-maximize::before {
    content: '';
    position: absolute;
    top: 4px;
    left: 4px;
    right: 4px;
    bottom: 4px;
    border: 2px solid #000;
    border-bottom-width: 3px;
}

.xp-btn-close::before {
    content: '×';
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 16px;
    font-weight: bold;
    color: #000;
    line-height: 1;
}

/* XP Window Content */
.xp-window-content {
    padding: 16px;
    background: #ece9d8;
}

/* Welcome Content */
.welcome-content {
    text-align: center;
}

.xp-title {
    font-size: 28px;
    font-weight: bold;
    color: #003c9d;
    margin-bottom: 12px;
    text-shadow: 1px 1px 2px rgba(255, 255, 255, 0.8);
}

.xp-subtitle {
    font-size: 14px;
    color: #333;
    margin-bottom: 24px;
}

.xp-features {
    display: flex;
    flex-direction: column;
    gap: 16px;
    margin-top: 20px;
}

.xp-feature-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px;
    background: linear-gradient(180deg, #fff 0%, #f0f0f0 100%);
    border: 1px solid;
    border-color: #fff #888 #888 #fff;
    border-radius: 4px;
    box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.1);
    opacity: 0;
}

.xp-icon-box {
    width: 40px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 24px;
    background: linear-gradient(135deg, #0054e3 0%, #003c9d 100%);
    border-radius: 4px;
    border: 2px solid;
    border-color: rgba(255, 255, 255, 0.5) rgba(0, 0, 0, 0.3) rgba(0, 0, 0, 0.3) rgba(255, 255, 255, 0.5);
    box-shadow: inset 0 1px 2px rgba(255, 255, 255, 0.3);
}

/* Upload Section */
.upload-header {
    text-align: center;
    margin-bottom: 20px;
}

.xp-heading {
    font-size: 18px;
    font-weight: bold;
    color: #003c9d;
    margin-bottom: 8px;
}

.xp-text {
    font-size: 12px;
    color: #333;
}

.upload-zone {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

/* XP Drop Zone */
.xp-drop-zone {
    border: 2px solid #7b96bc;
    background: linear-gradient(180deg, #fff 0%, #f5f5f5 100%);
    padding: 32px;
    text-align: center;
    cursor: pointer;
    border-radius: 4px;
    transition: all 0.2s;
    position: relative;
}

.xp-drop-zone::before {
    content: '';
    position: absolute;
    inset: 4px;
    border: 1px dashed #7b96bc;
    border-radius: 2px;
    pointer-events: none;
}

.xp-drop-zone:hover {
    background: linear-gradient(180deg, #fff 0%, #e8f4ff 100%);
    border-color: #0054e3;
}

.xp-drop-zone:hover::before {
    border-color: #0054e3;
}

.xp-drop-zone-active {
    background: linear-gradient(180deg, #e8f4ff 0%, #d0e8ff 100%);
    border-color: #0054e3;
    border-width: 3px;
}

.xp-drop-zone-active::before {
    border-color: #0054e3;
    border-width: 2px;
}

.drop-zone-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
}

.drop-zone-icon {
    width: 48px;
    height: 48px;
    color: #0054e3;
    filter: drop-shadow(2px 2px 3px rgba(0, 0, 0, 0.2));
}

.drop-zone-text {
    font-weight: bold;
    color: #003c9d;
    font-size: 14px;
}

.drop-zone-subtext {
    font-size: 11px;
    color: #666;
}

/* XP Group Box */
.xp-group-box {
    border: 2px solid;
    border-color: #fff #7b96bc #7b96bc #fff;
    background: #fff;
    padding: 12px;
    border-radius: 4px;
}

.xp-group-box-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
    padding-bottom: 8px;
    border-bottom: 1px solid #d3d3d3;
}

.xp-group-box-title {
    font-weight: bold;
    font-size: 11px;
    color: #003c9d;
}

.xp-link-button {
    font-size: 11px;
    color: #0054e3;
    text-decoration: underline;
    cursor: pointer;
    background: none;
    border: none;
    padding: 0;
}

.xp-link-button:hover {
    color: #ff6600;
}

/* XP File List */
.xp-file-list {
    display: flex;
    flex-direction: column;
    gap: 8px;
    max-height: 200px;
    overflow-y: auto;
    padding-right: 4px;
}

.xp-file-list::-webkit-scrollbar {
    width: 16px;
}

.xp-file-list::-webkit-scrollbar-track {
    background: #ece9d8;
    border: 1px solid #7b96bc;
}

.xp-file-list::-webkit-scrollbar-thumb {
    background: linear-gradient(180deg, #d8dae8 0%, #acb9d0 50%, #acb9d0 100%);
    border: 1px solid;
    border-color: #fff #7b96bc #7b96bc #fff;
}

.xp-file-list::-webkit-scrollbar-thumb:hover {
    background: linear-gradient(180deg, #e0e2f0 0%, #b8c5dd 50%, #b8c5dd 100%);
}

.xp-file-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px;
    background: linear-gradient(180deg, #fff 0%, #f8f8f8 100%);
    border: 1px solid #d3d3d3;
    border-radius: 2px;
    opacity: 0;
}

.xp-file-info {
    display: flex;
    align-items: center;
    gap: 10px;
    flex: 1;
    min-width: 0;
}

.xp-file-icon {
    width: 32px;
    height: 32px;
    flex-shrink: 0;
}

.xp-file-details {
    flex: 1;
    min-width: 0;
}

.xp-file-name {
    font-size: 11px;
    font-weight: bold;
    color: #000;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.xp-file-size {
    font-size: 10px;
    color: #666;
}

.xp-file-remove {
    width: 20px;
    height: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(180deg, #fff 0%, #e0e0e0 100%);
    border: 1px solid;
    border-color: #fff #7b96bc #7b96bc #fff;
    border-radius: 2px;
    cursor: pointer;
    color: #d00;
    transition: all 0.1s;
}

.xp-file-remove:hover {
    background: linear-gradient(180deg, #ffe0e0 0%, #ffc0c0 100%);
    border-color: #ffc0c0 #d00 #d00 #ffc0c0;
}

/* XP Controls */
.xp-control-group {
    display: flex;
    flex-direction: column;
    gap: 6px;
}

.xp-label {
    font-size: 11px;
    font-weight: bold;
    color: #000;
}

.xp-select,
.xp-input {
    width: 100%;
    padding: 4px 6px;
    border: 1px solid;
    border-color: #7b96bc #fff #fff #7b96bc;
    background: #fff;
    font-size: 11px;
    font-family: 'Tahoma', sans-serif;
    border-radius: 0;
}

.xp-select:focus,
.xp-input:focus {
    outline: 1px solid #0054e3;
    outline-offset: -1px;
}

.xp-input-readonly {
    background: #f0f0f0;
    color: #666;
}

.xp-input-group {
    display: flex;
    gap: 8px;
}

.xp-input-group .xp-input {
    flex: 1;
}

/* XP Buttons */
.xp-button {
    padding: 6px 16px;
    border: 1px solid;
    border-color: rgba(255, 255, 255, 0.9) rgba(0, 0, 0, 0.4) rgba(0, 0, 0, 0.4) rgba(255, 255, 255, 0.9);
    font-size: 11px;
    font-weight: bold;
    cursor: pointer;
    border-radius: 3px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
    transition: all 0.1s;
    font-family: 'Tahoma', sans-serif;
    position: relative;
    overflow: hidden;
}

.xp-button::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 50%;
    background: linear-gradient(180deg, rgba(255, 255, 255, 0.4) 0%, transparent 100%);
    pointer-events: none;
}

.xp-button-primary {
    background: linear-gradient(180deg, #d7ebff 0%, #77b9f5 50%, #4fa5f3 100%);
    color: #000;
}

.xp-button-primary:hover {
    background: linear-gradient(180deg, #e5f3ff 0%, #88c5ff 50%, #60b0ff 100%);
    border-color: rgba(255, 255, 255, 1) rgba(0, 0, 0, 0.5) rgba(0, 0, 0, 0.5) rgba(255, 255, 255, 1);
}

.xp-button-primary:active {
    background: linear-gradient(180deg, #4fa5f3 0%, #77b9f5 50%, #d7ebff 100%);
    border-color: rgba(0, 0, 0, 0.4) rgba(255, 255, 255, 0.9) rgba(255, 255, 255, 0.9) rgba(0, 0, 0, 0.4);
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.2);
}

.xp-button-primary:disabled {
    background: linear-gradient(180deg, #e8e8e8 0%, #d0d0d0 100%);
    color: #888;
    cursor: not-allowed;
    border-color: #aaa #ddd #ddd #aaa;
}

.xp-button-secondary {
    background: linear-gradient(180deg, #fff 0%, #d4d0c8 50%, #c8c4bc 100%);
    color: #000;
}

.xp-button-secondary:hover {
    background: linear-gradient(180deg, #fff 0%, #e0dcd5 50%, #d4d0c8 100%);
}

.xp-button-secondary:active {
    background: linear-gradient(180deg, #c8c4bc 0%, #d4d0c8 50%, #fff 100%);
    border-color: rgba(0, 0, 0, 0.4) rgba(255, 255, 255, 0.9) rgba(255, 255, 255, 0.9) rgba(0, 0, 0, 0.4);
}

.w-full {
    width: 100%;
}

/* Upload Progress */
.upload-progress {
    display: flex;
    align-items: center;
    gap: 8px;
}

.spinner {
    width: 14px;
    height: 14px;
    border: 2px solid rgba(0, 0, 0, 0.2);
    border-top-color: #000;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
}

@keyframes spin {
    to { transform: rotate(360deg); }
}

/* Success Content */
.success-content {
    text-align: center;
}

.success-icon-box {
    width: 80px;
    height: 80px;
    margin: 0 auto 20px;
    background: linear-gradient(135deg, #22c55e 0%, #16a34a 100%);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    border: 4px solid;
    border-color: rgba(255, 255, 255, 0.5) rgba(0, 0, 0, 0.2) rgba(0, 0, 0, 0.2) rgba(255, 255, 255, 0.5);
    box-shadow: 
        0 4px 8px rgba(0, 0, 0, 0.3),
        inset 0 2px 4px rgba(255, 255, 255, 0.3);
    animation: successPulse 1.5s ease-in-out infinite;
}

@keyframes successPulse {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.05); }
}

.success-checkmark {
    font-size: 48px;
    color: white;
    font-weight: bold;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.success-title {
    font-size: 16px;
    font-weight: bold;
    color: #003c9d;
    margin-bottom: 8px;
}

.success-subtitle {
    font-size: 12px;
    color: #333;
    margin-bottom: 20px;
}

/* Error Box */
.xp-error-box {
    padding: 12px;
    background: linear-gradient(180deg, #fff8f8 0%, #ffe8e8 100%);
    border: 2px solid #d00;
    border-radius: 4px;
}

.xp-error-content {
    display: flex;
    align-items: center;
    gap: 10px;
}

.xp-error-icon {
    font-size: 20px;
    flex-shrink: 0;
}

.xp-error-text {
    font-size: 11px;
    color: #d00;
    font-weight: bold;
}

/* XP Taskbar */
.xp-taskbar {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    height: 40px;
    background: linear-gradient(180deg, 
        #3788e8 0%,
        #2468c8 5%,
        #1f5fb5 50%,
        #1456a8 100%
    );
    border-top: 2px solid #1f5fb5;
    display: flex;
    align-items: center;
    padding: 0 4px;
    z-index: 1000;
    box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.3);
}

/* XP Start Button */
.xp-start-button {
    height: 32px;
    padding: 0 12px;
    background: linear-gradient(180deg, #5eab5c 0%, #3d8b3a 100%);
    border: 1px solid;
    border-color: rgba(255, 255, 255, 0.7) rgba(0, 0, 0, 0.5) rgba(0, 0, 0, 0.5) rgba(255, 255, 255, 0.7);
    border-radius: 4px;
    display: flex;
    align-items: center;
    gap: 6px;
    cursor: pointer;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
    transition: all 0.1s;
}

.xp-start-button:hover {
    background: linear-gradient(180deg, #6ebc6c 0%, #4d9b4a 100%);
}

.xp-start-button:active {
    background: linear-gradient(180deg, #3d8b3a 0%, #5eab5c 100%);
    border-color: rgba(0, 0, 0, 0.5) rgba(255, 255, 255, 0.7) rgba(255, 255, 255, 0.7) rgba(0, 0, 0, 0.5);
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.3);
}

.xp-start-icon {
    width: 20px;
    height: 20px;
    background: radial-gradient(circle, #fff 0%, #f0f0f0 100%);
    border-radius: 50%;
    border: 2px solid rgba(0, 0, 0, 0.2);
}

.xp-start-text {
    color: white;
    font-size: 12px;
    font-weight: bold;
    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
}

/* Taskbar Items */
.xp-taskbar-items {
    flex: 1;
    display: flex;
    gap: 4px;
    margin-left: 8px;
}

.xp-taskbar-item {
    height: 32px;
    padding: 0 12px;
    background: linear-gradient(180deg, #4f94e3 0%, #2a6ec5 100%);
    border: 1px solid;
    border-color: rgba(255, 255, 255, 0.4) rgba(0, 0, 0, 0.3) rgba(0, 0, 0, 0.3) rgba(255, 255, 255, 0.4);
    border-radius: 3px;
    display: flex;
    align-items: center;
    gap: 6px;
    cursor: pointer;
}

.xp-taskbar-item-active {
    background: linear-gradient(180deg, #2a6ec5 0%, #4f94e3 100%);
    border-color: rgba(0, 0, 0, 0.3) rgba(255, 255, 255, 0.4) rgba(255, 255, 255, 0.4) rgba(0, 0, 0, 0.3);
    box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.2);
}

.xp-taskbar-icon {
    font-size: 14px;
}

.xp-taskbar-label {
    color: white;
    font-size: 11px;
    font-weight: bold;
    text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.4);
}

/* System Tray */
.xp-system-tray {
    height: 32px;
    margin-left: auto;
    padding: 0 8px;
    background: linear-gradient(180deg, #1998ff 0%, #0c78e8 100%);
    border: 1px solid;
    border-color: rgba(255, 255, 255, 0.2) rgba(0, 0, 0, 0.4) rgba(0, 0, 0, 0.4) rgba(255, 255, 255, 0.2);
    border-radius: 3px;
    display: flex;
    align-items: center;
    gap: 8px;
}

.xp-tray-time {
    color: white;
    font-size: 11px;
    font-weight: bold;
    text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.4);
}

/* Responsive */
@media (max-width: 1024px) {
    .xp-desktop {
        padding-bottom: 60px;
    }
    
    .xp-title {
        font-size: 24px;
    }
    
    .xp-window {
        max-width: 100%;
    }
}

@media (max-width: 640px) {
    .xp-title {
        font-size: 20px;
    }
    
    .xp-subtitle {
        font-size: 12px;
    }
    
    .xp-window-content {
        padding: 12px;
    }
    
    .xp-taskbar-label {
        display: none;
    }
}
</style>
