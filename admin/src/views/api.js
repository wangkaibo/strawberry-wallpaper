const apiBaseUrl = 'http://strawberry.wangkaibo.com'
const axios = require('axios')
import $LocalStorage from '../assets/js/local-storage'
import qs from 'qs';

const instance = axios.create({
    baseURL: apiBaseUrl,
    timeout: 0
});

const noticeInstance = axios.create({
    baseURL: apiBaseUrl,
    timeout: 0,
})

// // 添加响应拦截器
// axios.interceptors.response.use(
//     (response) => {
//     // 对响应数据做点什么
//         const { data: result } = response;
//         if (result.code === 0) {
//             return Promise.resolve(result.data);
//         } 
//         return Promise.reject(result);
//     },
//     error => 
//     // 对响应错误做点什么
//         // eslint-disable-next-line implicit-arrow-linebreak
//         Promise.reject(error)

// );

function afterAxios(response) {
    const { data: result } = response;
    if (result.code === 0) {
        return Promise.resolve(result.data);
    }
    return Promise.reject(result);
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
        "TOKEN": `${$LocalStorage.getStore('token')}`
    }
}).then(response => {
    return new Promise((resolve, reject) => {
        afterAxios(response).then(res => {
            const now = (new Date).getTime()
            const result = res.map(item => {
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
    })
})

/**
 * 增加
 */
export const apiAddNotice = () => noticeInstance.post('/notice', {
    headers: {
        "Authorization": $LocalStorage.getStore('token')
    }
}).then(response => afterAxios(response))

/**
 * 删除公告
 */
export const apiDeleteNotice = () => noticeInstance.delete('/notice', {
    headers: {
        "TOKEN": $LocalStorage.getStore('token')
    }
}).then(response => afterAxios())
