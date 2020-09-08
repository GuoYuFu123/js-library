import Vue from 'vue';

import App from './App';
import '@/plugins';

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
}).$mount('#app');
