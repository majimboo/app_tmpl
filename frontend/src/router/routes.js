import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import UserList from '../views/UserList.vue'
import UserEdit from '../views/UserEdit.vue'
import AppLayout from '../layouts/AppLayout.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { hideLayout: true }
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: { hideLayout: true }
  },
  {
    path: '/',
    component: AppLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: Dashboard
      },
      {
        path: '/users',
        name: 'UserList',
        component: UserList
      },
      {
        path: '/users/:id/edit',
        name: 'UserEdit',
        component: UserEdit
      }
    ]
  }
]

export default routes