import { createRouter, createWebHistory } from 'vue-router'
import { useAuth } from '../composables/useAuth'
import SettingsPage from '../pages/settings/Settings.vue'
import AccountPage from '../pages/account/Account.vue'
import HomePage from '../pages/main/Home.vue'
import AuthPage from '../pages/auth/Auth.vue'

const routes = [
  {
    path: '/',
    name: 'Upload',
    component: HomePage,
    meta: { requiresAuth: true }
  },
  // Download page route (public)
  {
    path: '/download/:id',
    name: 'Download',
    component: () => import('../pages/download/Download.vue'), // Lazy loading
  },
  {
    path: '/settings',
    name: 'Settings',
    component: SettingsPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/account',
    name: 'Account',
    component: AccountPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/auth',
    name: 'Auth',
    component: AuthPage,
    meta: { requiresGuest: true }
  },
  // Redirect all unknown routes to auth
  {
    path: '/:pathMatch(.*)*',
    redirect: '/auth'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Global navigation guard
router.beforeEach(async (to, _from, next) => {
  const { isAuthenticated, fetchCurrentUser } = useAuth()
  
  // Try to fetch current user if we have a token but no user data
  if (!isAuthenticated.value) {
    await fetchCurrentUser()
  }
  
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const requiresGuest = to.matched.some(record => record.meta.requiresGuest)
  
  if (requiresAuth && !isAuthenticated.value) {
    // Redirect to auth if authentication is required but user is not authenticated
    next('/auth')
  } else if (requiresGuest && isAuthenticated.value) {
    // Redirect to home if guest route is accessed but user is authenticated
    next('/')
  } else {
    next()
  }
})

export default router