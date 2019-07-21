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
import SessionService from './services/session/service.js'

import router from './router'

Vue.use($session);

Vue.component('ErrorBlock', ErrorBlock);

Vue.directive('permission', {
  bind(el, binding) {
    let clvl = SessionService.getCachedUser().access || -1;

    if (binding.value > clvl) {
      el.style.display = 'none';
    }
  }
});

function createApp() {
  new Vue({
    router,
    render: h => h(App)
  }).$mount('#app')
}

SessionService.whoami(createApp);