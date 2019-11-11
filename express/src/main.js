import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import {Icon, Button, Form, Tabs, Input, Checkbox, Layout, } from 'ant-design-vue';

Vue.config.productionTip = false;
Vue.use(Button)
Vue.use(Icon)
Vue.use(Form)
Vue.use(Tabs)
Vue.use(Input)
Vue.use(Checkbox)
Vue.use(Layout)

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount("#app");
