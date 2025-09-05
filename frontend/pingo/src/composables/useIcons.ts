import { computed } from 'vue'

// Import SVG icons as URLs using Vite's asset handling
import filesPdfIcon from '@/assets/svg/icons/files_pdf.svg?url'
import wordDocIcon from '@/assets/svg/icons/word_docx_doc.svg?url'
import pptFileIcon from '@/assets/svg/icons/ppt_pptx_file.svg?url'
import filesArchiveIcon from '@/assets/svg/icons/files_archive.svg?url'
import filesUploadedIcon from '@/assets/svg/icons/files_uploaded.svg?url'
import filesUploaded2Icon from '@/assets/svg/icons/files_uploaded2.svg?url'
import totalFilesIcon from '@/assets/svg/icons/total_files_icon.svg?url'
import unknownFileIcon from '@/assets/svg/icons/unknown_file_formats.svg?url'

// Define file type to icon mapping using imported assets
const FILE_TYPE_ICONS = {
  // Document files
  pdf: filesPdfIcon,
  doc: wordDocIcon,
  docx: wordDocIcon,
  ppt: pptFileIcon,
  pptx: pptFileIcon,
  
  // Image files
  jpg: unknownFileIcon,
  jpeg: unknownFileIcon,
  png: unknownFileIcon,
  gif: unknownFileIcon,
  svg: unknownFileIcon,
  webp: unknownFileIcon,
  
  // Video files
  mp4: unknownFileIcon,
  mov: unknownFileIcon,
  avi: unknownFileIcon,
  mkv: unknownFileIcon,
  wmv: unknownFileIcon,
  flv: unknownFileIcon,
  webm: unknownFileIcon,
  
  // Audio files
  mp3: unknownFileIcon,
  wav: unknownFileIcon,
  flac: unknownFileIcon,
  aac: unknownFileIcon,
  ogg: unknownFileIcon,
  
  // Archive files
  zip: filesArchiveIcon,
  rar: filesArchiveIcon,
  '7z': filesArchiveIcon,
  tar: filesArchiveIcon,
  gz: filesArchiveIcon,
  
  // Text files
  txt: unknownFileIcon,
  md: unknownFileIcon,
  rtf: unknownFileIcon,
  
  // Spreadsheet files
  xls: unknownFileIcon,
  xlsx: unknownFileIcon,
  csv: unknownFileIcon,
  
  // Code files
  js: unknownFileIcon,
  ts: unknownFileIcon,
  html: unknownFileIcon,
  css: unknownFileIcon,
  json: unknownFileIcon,
  xml: unknownFileIcon,
  vue: unknownFileIcon,
  py: unknownFileIcon,
  java: unknownFileIcon,
  cpp: unknownFileIcon,
  c: unknownFileIcon,
  php: unknownFileIcon,
  rb: unknownFileIcon,
  go: unknownFileIcon,
} as const

// System icons for UI components
const SYSTEM_ICONS = {
  // Upload/file management
  filesUploaded: filesUploadedIcon,
  filesUploaded2: filesUploaded2Icon,
  totalFiles: totalFilesIcon,
  multipleFiles: unknownFileIcon,
  fileFolder: unknownFileIcon,
  
  // File type icons (for missing SVGs, we'll fall back to default)
  filesImage: unknownFileIcon,
  filesAudio: unknownFileIcon,
  filesArchive: filesArchiveIcon,
  filesText: unknownFileIcon,
  filesExcel: unknownFileIcon,
  filesCode: unknownFileIcon,
  
  // File operations
  download: unknownFileIcon,
  upload: unknownFileIcon,
  share: unknownFileIcon,
  copy: unknownFileIcon,
  delete: unknownFileIcon,
  
  // Status icons
  success: unknownFileIcon,
  error: unknownFileIcon,
  warning: unknownFileIcon,
  info: unknownFileIcon,
  
  // Navigation
  home: unknownFileIcon,
  settings: unknownFileIcon,
  account: unknownFileIcon,
  
  // Default fallback
  default: unknownFileIcon,
} as const

export type FileExtension = keyof typeof FILE_TYPE_ICONS
export type SystemIcon = keyof typeof SYSTEM_ICONS

export function useIcons() {
  /**
   * Get file extension from filename
   */
  const getFileExtension = (filename: string): string => {
    const parts = filename.split('.')
    return parts.length > 1 ? parts.pop()?.toLowerCase() || '' : ''
  }

  /**
   * Get icon path for a specific file based on its extension
   */
  const getFileIcon = (filename: string): string => {
    const extension = getFileExtension(filename) as FileExtension
    return FILE_TYPE_ICONS[extension] || SYSTEM_ICONS.default
  }

  /**
   * Get icon alt text for a file
   */
  const getFileIconAlt = (filename: string): string => {
    const extension = getFileExtension(filename)
    
    // Special cases for better alt text
    const altTextMap: Record<string, string> = {
      pdf: 'PDF document',
      doc: 'Word document',
      docx: 'Word document',
      ppt: 'PowerPoint presentation',
      pptx: 'PowerPoint presentation',
      xls: 'Excel spreadsheet',
      xlsx: 'Excel spreadsheet',
      jpg: 'Image file',
      jpeg: 'Image file',
      png: 'Image file',
      gif: 'Animated image',
      mp4: 'Video file',
      mov: 'Video file',
      mp3: 'Audio file',
      wav: 'Audio file',
      zip: 'Archive file',
      txt: 'Text file',
      js: 'JavaScript file',
      ts: 'TypeScript file',
      vue: 'Vue component',
      html: 'HTML file',
      css: 'CSS file',
    }
    
    return altTextMap[extension] || `${extension.toUpperCase()} file` || 'File'
  }

  /**
   * Get icon for upload based on files count and types
   */
  const getUploadIcon = (files: string[]): string => {
    // If multiple files, use multiple files icon
    if (files.length > 1) {
      return SYSTEM_ICONS.multipleFiles
    }
    
    // For single file, use specific file type icon
    return getFileIcon(files[0])
  }

  /**
   * Get icon alt text for upload
   */
  const getUploadIconAlt = (files: string[]): string => {
    if (files.length > 1) {
      return `${files.length} files`
    }
    
    return getFileIconAlt(files[0])
  }

  /**
   * Get system icon path
   */
  const getSystemIcon = (iconName: SystemIcon): string => {
    return SYSTEM_ICONS[iconName]
  }

  /**
   * Get file type category for styling purposes
   */
  const getFileCategory = (filename: string): string => {
    const extension = getFileExtension(filename)
    
    if (['jpg', 'jpeg', 'png', 'gif', 'svg', 'webp'].includes(extension)) {
      return 'image'
    }
    if (['mp4', 'mov', 'avi', 'mkv', 'wmv', 'flv', 'webm'].includes(extension)) {
      return 'video'
    }
    if (['mp3', 'wav', 'flac', 'aac', 'ogg'].includes(extension)) {
      return 'audio'
    }
    if (['pdf', 'doc', 'docx', 'ppt', 'pptx', 'xls', 'xlsx'].includes(extension)) {
      return 'document'
    }
    if (['zip', 'rar', '7z', 'tar', 'gz'].includes(extension)) {
      return 'archive'
    }
    if (['js', 'ts', 'html', 'css', 'json', 'xml', 'vue', 'py', 'java', 'cpp', 'c', 'php', 'rb', 'go'].includes(extension)) {
      return 'code'
    }
    if (['txt', 'md', 'rtf'].includes(extension)) {
      return 'text'
    }
    
    return 'other'
  }

  /**
   * Check if file is an image
   */
  const isImageFile = (filename: string): boolean => {
    return getFileCategory(filename) === 'image'
  }

  /**
   * Check if file is a video
   */
  const isVideoFile = (filename: string): boolean => {
    return getFileCategory(filename) === 'video'
  }

  /**
   * Check if file is an audio file
   */
  const isAudioFile = (filename: string): boolean => {
    return getFileCategory(filename) === 'audio'
  }

  /**
   * Get file size color class based on size
   */
  const getFileSizeColor = (bytes: number, isDark: boolean): string => {
    const mb = bytes / (1024 * 1024)
    
    if (mb < 1) {
      return isDark ? 'text-green-400' : 'text-green-600'
    } else if (mb < 10) {
      return isDark ? 'text-yellow-400' : 'text-yellow-600'
    } else if (mb < 100) {
      return isDark ? 'text-orange-400' : 'text-orange-600'
    } else {
      return isDark ? 'text-red-400' : 'text-red-600'
    }
  }

  return {
    // Core functions
    getFileExtension,
    getFileIcon,
    getFileIconAlt,
    getUploadIcon,
    getUploadIconAlt,
    getSystemIcon,
    
    // Category functions
    getFileCategory,
    isImageFile,
    isVideoFile,
    isAudioFile,
    
    // Utility functions
    getFileSizeColor,
    
    // Constants for external use
    FILE_TYPE_ICONS: computed(() => FILE_TYPE_ICONS),
    SYSTEM_ICONS: computed(() => SYSTEM_ICONS),
  }
}
