import {
  createRouter,
  createWebHistory,
  type RouteLocationNormalized,
  type NavigationGuardNext,
} from 'vue-router'
import Login from '@/views/LoginForm.vue'
import Signup from '@/views/SignupView.vue'
import AdminDashboard from '@/views/AdminDashboard.vue'

const isAuthenticated = () => {
  const token = localStorage.getItem('token')
  return token && !isTokenExpired(token)
}

const isTokenExpired = (token: string) => {
  const payload = JSON.parse(atob(token.split('.')[1]))
  return Date.now() >= payload.exp * 1000
}

const userRole = () => localStorage.getItem('userRole')

const routes = [
  { path: '/', redirect: '/login' },
  { path: '/login', component: Login },
  { path: '/signup', component: Signup },
  {
    path: '/admin/dashboard',
    component: AdminDashboard,
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext,
    ) => {
      if (isAuthenticated() && userRole() === 'admin') {
        next()
      } else {
        next('/login')
      }
    },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
