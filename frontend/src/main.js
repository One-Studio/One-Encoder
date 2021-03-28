import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';

// import Vant from 'vant';
// import 'vant/lib/index.css';
// import ElementUI from 'element-ui';
// import 'element-ui/lib/theme-chalk/index.css';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';

// Vue.use(Vant);
// Vue.use(ElementUI);
Vue.use(Antd);

// i18n
import VueI18n from 'vue-i18n';

Vue.use(VueI18n);

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
    new Vue({
        render: h => h(App)
    }).$mount('#app');
});
