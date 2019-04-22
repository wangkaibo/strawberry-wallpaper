### 草莓壁纸接口文档 V0.1

#### 接口返回格式定义
```js
{
    code:200,     // 标准的状态码
    data:[]|{}|'', //返回的数据
    message:'请求成功' //返回的提示信息
}
```

#### 接口基地址 https://api.taoacat.com/strawberry/v0.1/

接口地址|请求方式|接口说明|接口参数
--|--|--|--
/register|POST|第一次注册的时候保留注册信息|os:`String` 系统类型</br> osVersion:`String` 系统版本</br> version：`String` 软件版本 </br> userName: `String` 系统用户名 </br> resTime:`String` 注册时间,unix秒 </br> uid: `String` 生成的软件唯一ID,生成方式 md5_32(`${os}${ovVersion}${userName}${resTime}`)
/statistic|GET|统计正在使用的用户|uid:`String` 用户ID </br>