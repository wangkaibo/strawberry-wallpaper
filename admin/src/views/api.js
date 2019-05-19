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

function afterAxios(response){
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
