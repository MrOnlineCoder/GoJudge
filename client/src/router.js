import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/views/Home.vue'
import Problems from '@/views/Problems.vue'
import Submissions from '@/views/Submissions.vue'
import Login from '@/views/Login.vue'
import Admin from '@/views/Admin.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/problems',
      name: 'problems',
      component: Problems
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/submissions',
      name: 'submissions',
      component: Submissions
    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
