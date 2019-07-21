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

Vue.filter('formatTimelimit', (val) => {
  if (val < 1000) {
    return `${val} ms.`;
  } else {
    let secs = Math.floor(val / 1000);
    let left = val % 1000;

    if (left == 0) {
      return `${secs} s.`;
    } else {
      return `${secs} s. ${left} ms.`;
    }
  }
});

Vue.filter('formatMemlimit', (val) => {
  if (val < 1024) {
    return `${val} KB`;
  } else {
    let megs = Math.floor(val / 1024);
    let left = val % 1024;

    if (left == 0) {
      return `${megs} MB.`;
    } else {
      return `${megs} MB ${left} KB.`; //ugly case
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