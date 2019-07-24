import Vue from 'vue'
import Router from 'vue-router'

import Home from '@/views/Home.vue'
import Problems from '@/views/Problems.vue'
import Submissions from '@/views/Submissions.vue'
import Login from '@/views/Login.vue'
import Admin from '@/views/Admin.vue'
import Problem from '@/views/Problem.vue'
import Submit from '@/views/Submit.vue'
import Settings from '@/views/Settings.vue'

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
      path: '/problem/:idx',
      name: 'problem',
      component: Problem
    },
    {
      path: '/submit',
      name: 'submit',
      component: Submit
    },
    {
      path: '/submit/:idx',
      name: 'submit_idx',
      component: Submit
    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin
    },
    {
      path: '/settings',
      name: 'settings',
      component: Settings
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
