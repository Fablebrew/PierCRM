import { createRouter, createWebHistory } from 'vue-router'

import Workspace from '../components/Workspace.vue'
import LoginView from '../components/Auth.vue'
import Register from '../components/Register.vue'
import Projects from '@/components/Projects.vue'
import Team from '@/components/Team.vue'
import OtherExpenses from '@/components/OtherExpenses.vue'
import Tarif from '@/components/Tarif.vue'

const routes = [
  {
    path: '/register',
    name: 'register',
    component: Register
  },
  {
    path: '/tarif',
    name: 'tarif',
    component: Tarif
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/workspace',
    name: 'workspace',
    meta: { requiresAuth: true },
    component: Workspace
  },
  {
    path: '/projects',
    name: 'projects',
    meta: { requiresAuth: true },
    component: Projects
  },
  {
    path: '/team',
    name: 'team',
    component: Team
  },
  {
    path: '/expenses',
    name: 'expenses',
    component: OtherExpenses
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Проверка срока действия токена
function isTokenExpired(token) {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    return payload.exp * 1000 < Date.now()
  } catch (e) {
    return true // если не можем прочитать - считаем просроченным
  }
}

router.beforeEach((to) => {
  const token = localStorage.getItem('token')

  // Проверяем токен на срок действия
  if (token && isTokenExpired(token)) {
    localStorage.removeItem('token')
  }

  const validToken = localStorage.getItem('token')

  // Если маршрут требует авторизацию и токена нет
  if (to.meta.requiresAuth && !validToken) {
    return '/login'
  }

  // Если пользователь уже авторизован и идёт на login
  if (to.path === '/login' && validToken) {
    return '/workspace'
  }

  // Иначе разрешаем переход
  return true
})

export default router