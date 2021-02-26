import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
// import Antd from 'ant-design-vue';
// import 'ant-design-vue/dist/antd.css';

import Quasar from 'quasar'
import 'quasar/dist/quasar.css'

Vue.use(Quasar);
// import {
// 	Quasar,
// 	QAjaxBar
// } from 'quasar'
//
// Vue.use(Quasar, {
// 	components: {
// 		QAjaxBar
// 	}
// })

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

// Vue.use(Antd);

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
