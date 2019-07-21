/*
	Vue 
*/
import Vue from 'vue'

Vue.config.productionTip = false

/*
	Bootstrap Vue
*/

import BootstrapVue from 'bootstrap-vue'

Vue.use(BootstrapVue)

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

/*
	Font Awesome Icons
*/

import { library } from '@fortawesome/fontawesome-svg-core'

import { fas } from '@fortawesome/free-solid-svg-icons'

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(fas);

Vue.component('font-awesome-icon', FontAwesomeIcon);

/*
	Application
*/

import App from './App.vue'

import ErrorBlock from './components/ErrorBlock.vue'

import $session from './services/session'

import router from './router'

Vue.use($session);

Vue.component('ErrorBlock', ErrorBlock);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
