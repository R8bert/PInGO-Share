import { createRouter, createWebHistory } from 'vue-router'
// import UploadPage from '../components/Upload.vue' // Componenta ta existentă
import SettingsPage from '../pages/settings/Settings.vue'
import AccountPage from '../pages/account/Account.vue'
import HomePage from '../pages/main/Home.vue'

// Definirea imaginilor de fundal pentru pagini

const routes = [
  {
    path: '/',
    name: 'Upload',
    component: HomePage, // Componenta ta existentă
  },
  // Adaugă alte rute aici, ex. pentru o pagină de descărcare
  {
    path: '/download/:id',
    name: 'Download',
    component: () => import('../components/Upload.vue'), // Lazy loading
  },
  {
    path: '/settings',
    name: 'Settings',
    component: SettingsPage,
  },
  {
    path: '/account',
    name: 'Account',
    component: AccountPage,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router