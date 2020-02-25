import qs from 'qs'
import { Message } from 'element-ui'
import axios from 'axios'
import mainVue from '../main'
import $LocalStorage from '../assets/js/local-storage'

const apiBaseUrl = 'http://strawberry.wangkaibo.com'

const instance = axios.create({
    baseURL: apiBaseUrl,
    timeout: 0
})

const noticeInstance = axios.create({
    baseURL: apiBaseUrl,
    timeout: 0,
})

function afterAxios(response) {
    const { data: result } = response
    if (result.code === 0) {
        return Promise.resolve(result.data)
    }
    if (result.code === 401){
        mainVue.$router.replace('/login')
    }
    Message.error(result.message || '')
    return Promise.reject(result)
}

/**
 * 登录
 * @param {*} data
 */
export const apiLogin = data => instance({
    url: '/login',
    method: 'post',
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
    },
    data: qs.stringify(data)
}).then(response => afterAxios(response))

/**
 * 获取统计
 */
export const apiStatistic = () => instance.get('/statistic').then(response => afterAxios(response))

/**
 * 获取列表
 */
export const apiGetNoticeList = () => noticeInstance.get('/notice_list', {
    headers: {
        Authorization: `${$LocalStorage.getStore('token')}`
    }
}).then(response => new Promise((resolve, reject) => {
    afterAxios(response).then((res) => {
        const now = (new Date()).getTime()
        const result = res.map((item) => {
            let status = '2'
            if (now > item.expire_at && item.is_publish === 1) {
                status = '3'
            }
            else if (item.is_publish === 1) {
                status = '1'
            }
            return {
                ...item,
                status
            }
        })
        resolve(result)
    }).catch((err) => {
        console.log(err)
        resolve([])
    })
}))

/**
 * 增加,编辑
 */
export const apiAddNotice = data => noticeInstance.request({
    url: '/notice',
    method: 'POST',
    ...data,
    data: qs.stringify(data.data),
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Authorization: $LocalStorage.getStore('token')
    }
}).then(response => afterAxios(response))

/**
 * 删除公告
 */
export const apiDeleteNotice = id => noticeInstance.request({
    url: `/notice/${id}`,
    method: 'DELETE',
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Authorization: $LocalStorage.getStore('token')
    }
}).then(response => afterAxios(response))

// 发布公告
export const apiPublicNotice = (id, data) => noticeInstance.request({
    url: `/notice/${id}`,
    method: 'PATCH',
    data: qs.stringify(data),
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Authorization: $LocalStorage.getStore('token')
    }
}).then(response => afterAxios(response))
