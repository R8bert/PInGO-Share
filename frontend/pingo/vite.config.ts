import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'
import tailwindcss from '@tailwindcss/vite'
import Icons from 'unplugin-icons/vite'
import * as path from 'path'
import { fileURLToPath } from 'url'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(), 
    VueDevTools(), 
    tailwindcss(),
    Icons({
      compiler: 'vue3',
      autoInstall: true,
    }),
  ],
    resolve: {
        alias: {
        '@': path.resolve(__dirname, './src'),
        'images': path.resolve(__dirname, './src/assets/images'),
        },
    },
    assetsInclude: ['**/*.svg'],
    server: {
        proxy: {
            '/api': {
                target: 'http://localhost:8080',
                changeOrigin: true,
                // Don't rewrite the path - backend expects /api prefix
            },
        }
    },

})
