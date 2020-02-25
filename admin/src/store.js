import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        noticeDetail:{}
    },
    mutations: {
        setNoticeDetail(state, data){
            state.noticeDetail=data
        }
    },
    actions: {

    }
})
