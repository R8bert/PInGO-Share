// Icon utilities for consistent icon usage across components
import { useIcons } from '../composables/useIcons'

/**
 * Icon size presets for consistent sizing
 */
export const ICON_SIZES = {
  xs: 'w-3 h-3',      // 12px
  sm: 'w-4 h-4',      // 16px  
  md: 'w-6 h-6',      // 24px
  lg: 'w-8 h-8',      // 32px
  xl: 'w-12 h-12',    // 48px
  '2xl': 'w-16 h-16', // 64px
} as const

export type IconSize = keyof typeof ICON_SIZES

/**
 * Common icon props interface
 */
export interface IconProps {
  size?: IconSize
  className?: string
  alt?: string
}

/**
 * Get Tailwind classes for icon sizing
 */
export function getIconSizeClasses(size: IconSize = 'md'): string {
  return ICON_SIZES[size]
}

/**
 * Create a reusable icon component props object
 */
export function createIconProps(
  src: string, 
  alt: string, 
  size: IconSize = 'md', 
  additionalClasses: string = ''
): Record<string, any> {
  return {
    src,
    alt,
    class: `${getIconSizeClasses(size)} object-contain ${additionalClasses}`.trim(),
    style: getFixedIconStyles(size)
  }
}

/**
 * Get fixed icon styles to ensure consistent sizing
 */
export function getFixedIconStyles(size: IconSize = 'md'): string {
  const sizeMap = {
    xs: '12px',
    sm: '16px', 
    md: '24px',
    lg: '32px',
    xl: '48px',
    '2xl': '64px',
  }
  
  const pixels = sizeMap[size]
  return `width: ${pixels} !important; height: ${pixels} !important;`
}

/**
 * Utility function to get upload icon with size
 */
export function getUploadIconWithSize(files: string[], size: IconSize = 'md') {
  const { getUploadIcon, getUploadIconAlt } = useIcons()
  
  return createIconProps(
    getUploadIcon(files),
    getUploadIconAlt(files),
    size
  )
}

/**
 * Utility function to get file icon with size
 */
export function getFileIconWithSize(filename: string, size: IconSize = 'md') {
  const { getFileIcon, getFileIconAlt } = useIcons()
  
  return createIconProps(
    getFileIcon(filename),
    getFileIconAlt(filename),
    size
  )
}

/**
 * Utility function to get system icon with size
 */
export function getSystemIconWithSize(iconName: string, size: IconSize = 'md', alt: string = '') {
  const { getSystemIcon } = useIcons()
  
  return createIconProps(
    getSystemIcon(iconName as any),
    alt || iconName,
    size
  )
}

/**
 * File type color schemes for consistent theming
 */
export const FILE_TYPE_COLORS = {
  image: {
    light: 'bg-purple-100 text-purple-600',
    dark: 'bg-purple-900/30 text-purple-400'
  },
  video: {
    light: 'bg-red-100 text-red-600', 
    dark: 'bg-red-900/30 text-red-400'
  },
  audio: {
    light: 'bg-green-100 text-green-600',
    dark: 'bg-green-900/30 text-green-400' 
  },
  document: {
    light: 'bg-blue-100 text-blue-600',
    dark: 'bg-blue-900/30 text-blue-400'
  },
  archive: {
    light: 'bg-orange-100 text-orange-600',
    dark: 'bg-orange-900/30 text-orange-400'
  },
  code: {
    light: 'bg-gray-100 text-gray-600',
    dark: 'bg-gray-800 text-gray-400'
  },
  text: {
    light: 'bg-gray-100 text-gray-600', 
    dark: 'bg-gray-800 text-gray-400'
  },
  other: {
    light: 'bg-gray-100 text-gray-600',
    dark: 'bg-gray-800 text-gray-400'
  }
} as const

/**
 * Get color classes for file type based on theme
 */
export function getFileTypeColors(filename: string, isDark: boolean): string {
  const { getFileCategory } = useIcons()
  const category = getFileCategory(filename) as keyof typeof FILE_TYPE_COLORS
  
  return FILE_TYPE_COLORS[category]?.[isDark ? 'dark' : 'light'] || FILE_TYPE_COLORS.other[isDark ? 'dark' : 'light']
}
