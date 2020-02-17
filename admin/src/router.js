import Vue from 'vue'
import Router from 'vue-router'
import login from './views/login.vue'
import statistics from './views/statistics.vue'
import noticeList from './views/notice-list.vue'
import noticeAdd from './views/notice-add.vue'

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
            name: 'statistics',
            component: statistics
        },
        {
            path: '/notice-list',
            name: 'noticeList',
            component: noticeList
        },
        {
            path: '/notice-add',
            name: 'noticeAdd',
            component: noticeAdd
        },

    ]
})
