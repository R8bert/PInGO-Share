/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class', // Enable class-based dark mode
  theme: {
    extend: {
      colors: {
        // Custom color scheme for better dark mode support - much blacker
        background: {
          light: '#ffffff',
          dark: '#000000',
        },
        surface: {
          light: '#f8fafc',
          dark: '#0a0a0a',
        },
        card: {
          light: '#ffffff',
          dark: '#1a1a1a',
        },
        text: {
          primary: {
            light: '#1e293b',
            dark: '#ffffff',
          },
          secondary: {
            light: '#64748b',
            dark: '#e5e5e5',
          },
          muted: {
            light: '#94a3b8',
            dark: '#a3a3a3',
          },
        },
        border: {
          light: '#e2e8f0',
          dark: '#262626',
        },
        accent: {
          blue: {
            light: '#3b82f6',
            dark: '#60a5fa',
          },
          green: {
            light: '#10b981',
            dark: '#34d399',
          },
          red: {
            light: '#ef4444',
            dark: '#f87171',
          },
          purple: {
            light: '#8b5cf6',
            dark: '#a78bfa',
          },
        }
      },
      animation: {
        'fade-in': 'fadeIn 0.5s ease-in-out',
        'slide-up': 'slideUp 0.3s ease-out',
        'theme-transition': 'themeTransition 0.3s ease-in-out',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0', transform: 'translateY(10px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        slideUp: {
          '0%': { opacity: '0', transform: 'translateY(20px)' },
          '100%': { opacity: '1', transform: 'translateY(0)' },
        },
        themeTransition: {
          '0%': { opacity: '0.8' },
          '100%': { opacity: '1' },
        },
      },
    },
  },
  plugins: [],
}
