import Vue from 'vue';
import VueRouter from 'vue-router';
import VueMaterial from 'vue-material';
import App from './app';
import routes from './routes';
import VueSocketio from 'vue-socket.io';
import VueHttp from './http';
import VueThemes from './themes';

import 'roboto-fontface/css/roboto/roboto-fontface.css';
import 'material-design-icons/iconfont/material-icons.css';
import 'vue-material/dist/vue-material.css';

Vue.config.productionTip = false;
Vue.use(VueRouter);
Vue.use(VueMaterial);
Vue.use(VueHttp, {
  baseURL: '/api',
});
Vue.use(VueSocketio, '/');
Vue.material.registerTheme(VueThemes);

new Vue({
  el: '#app',
  template: '<App/>',
  components: {
    App,
  },
  router: new VueRouter({
    routes,
  }),
});
