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
    // meta: { requiresAuth: true },
    component: Workspace
  },
  {
    path: '/projects',
    name: 'projects',
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
router.beforeEach((to) => {
  const token = localStorage.getItem('token')

  // если маршрут требует авторизацию и токена нет
  if (to.meta.requiresAuth && !token) {
    return '/login'
  }

  // если пользователь уже авторизован и идёт на login
  if (to.path === '/login' && token) {
    return '/workspace'
  }

  // иначе разрешаем переход
  return true
})
export default router