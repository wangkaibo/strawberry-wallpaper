import Vue from 'vue'
import ElementUI from 'element-ui'
import App from './App.vue'
import router from './router'
import store from './store'
import '../node_modules/element-ui/lib/theme-chalk/index.css'
import './assets/css/index.css'
import localStorage from './assets/js/local-storage.js'

Vue.config.productionTip = false
Vue.prototype.$localStorage = localStorage

Vue.use(ElementUI)

const mainVue = new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')


export default mainVue
