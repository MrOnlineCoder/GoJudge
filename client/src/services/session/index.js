import SessionService from './service.js'

const plugin = {
  install(_Vue) {
    _Vue.prototype.$session = SessionService;
  }
};

export default plugin;