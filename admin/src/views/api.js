const apiBaseUrl = 'http://strawberry.wangkaibo.com';
const axios = require('axios');

const instance = axios.create({
    baseURL: apiBaseUrl,
    timeout: 1000
});

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
    data
}).then(response => afterAxios(response))

/**
 * 获取统计
 */
export const apiStatistic = () => instance.get('/statistic').then(response => afterAxios(response))

/**
 * 获取列表
 */
export const apiGetNoticeList = () => instance.get('/notice_list').then(response => {
    return new Promise((resolve, reject) => {
        afterAxios(response).then(res => {
   
            const now = (new Date).getTime()
            const result = res.map(item => {
                let status = '2'
                if (now > item.expire_at && is_publish === 1) {
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
            resolve([])
        })
    })
})

/**
 * 增加
 */
export const apiAddNotice = () => instance.post('/notice').then(response => afterAxios(response))

/**
 * 删除公告
 */
export const apiDeleteNotice = () => instance.delete('/notice').then(response => afterAxios())
