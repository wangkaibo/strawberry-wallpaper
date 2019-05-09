const apiBaseUrl = "http://strawberry.wangkaibo.com";
const axios = require("axios");

const instance = axios.create({
  baseURL: apiBaseUrl,
  timeout: 1000
});

// 添加响应拦截器
axios.interceptors.response.use(
  function(response) {
    // 对响应数据做点什么
    const { data: result } = response;
    if (result.code === 0) {
      return Promise.resolve(result);
    } else {
      return Promise.reject(result);
    }
  },
  function(error) {
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);

/**
 * 登录
 * @param {*} data
 */
export const apiLogin = data => {
    return instance({
        url: "/login",
        method: "post",
        data
      })
};

/**
 * 获取统计
 */
export const apiStatistic = () => {
    return instance.get("/statistic")
};
