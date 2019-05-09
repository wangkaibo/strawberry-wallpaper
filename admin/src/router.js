import Vue from 'vue'
import Router from 'vue-router'
import login from './views/login.vue'
import main from './views/main.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/login',
            name: '/login',
            component: login
        },
        {
            path: '/',
            name: 'main',
            component: main
        }
    ]
})
