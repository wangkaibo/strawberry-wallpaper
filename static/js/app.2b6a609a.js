(function(t){function e(e){for(var n,o,r=e[0],l=e[1],c=e[2],d=0,f=[];d<r.length;d++)o=r[d],s[o]&&f.push(s[o][0]),s[o]=0;for(n in l)Object.prototype.hasOwnProperty.call(l,n)&&(t[n]=l[n]);u&&u(e);while(f.length)f.shift()();return i.push.apply(i,c||[]),a()}function a(){for(var t,e=0;e<i.length;e++){for(var a=i[e],n=!0,r=1;r<a.length;r++){var l=a[r];0!==s[l]&&(n=!1)}n&&(i.splice(e--,1),t=o(o.s=a[0]))}return t}var n={},s={app:0},i=[];function o(e){if(n[e])return n[e].exports;var a=n[e]={i:e,l:!1,exports:{}};return t[e].call(a.exports,a,a.exports,o),a.l=!0,a.exports}o.m=t,o.c=n,o.d=function(t,e,a){o.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:a})},o.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},o.t=function(t,e){if(1&e&&(t=o(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var a=Object.create(null);if(o.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var n in t)o.d(a,n,function(e){return t[e]}.bind(null,n));return a},o.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return o.d(e,"a",e),e},o.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},o.p="/";var r=window["webpackJsonp"]=window["webpackJsonp"]||[],l=r.push.bind(r);r.push=e,r=r.slice();for(var c=0;c<r.length;c++)e(r[c]);var u=l;i.push([0,"chunk-vendors"]),a()})({0:function(t,e,a){t.exports=a("56d7")},"034f":function(t,e,a){"use strict";var n=a("9e74"),s=a.n(n);s.a},"56d7":function(t,e,a){"use strict";a.r(e);a("5c07"),a("53da"),a("2556"),a("d0f8");var n=a("6e6d"),s=a("bea0"),i=a.n(s),o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"sw-app",attrs:{id:"app"}},[a("router-view")],1)},r=[],l=(a("034f"),a("17cc")),c={},u=Object(l["a"])(c,o,r,!1,null,null,null),d=u.exports,f=a("1e6f"),m=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"sw-login"},[a("div",{staticClass:"sw-login_content"},[a("el-form",{attrs:{size:"medium","label-position":"right","label-width":"60px",model:t.form}},[a("el-form-item",{attrs:{label:"用户名"}},[a("el-input",{model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1),a("el-form-item",{attrs:{label:"密码"}},[a("el-input",{model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1),a("el-form-item",[a("el-button",{on:{click:function(e){return t.submit()}}},[t._v("登录")])],1)],1)],1)])},v=[],p="http://strawberry.wangkaibo.com",_=a("7f43"),w=_.create({baseURL:p,timeout:1e3});function h(t){var e=t.data;return 0===e.code?Promise.resolve(e.data):Promise.reject(e)}var b=function(t){return w({url:"/login",method:"post",data:t}).then(function(t){return h(t)})},g=function(){return w.get("/statistic").then(function(t){return h(t)})},y={kk:"1234"},C={name:"login",data:function(){return{form:{name:"",password:""}}},mounted:function(){console.log(y)},methods:{submit:function(){b(this.form).then(function(t){console.log("登录成功")})}}},x=C,S=Object(l["a"])(x,m,v,!1,null,"43fef86c",null),O=S.exports,k=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"sw-main"},[a("div",{staticClass:"sw-main_header"},[a("div",{staticClass:"sw-main_header_h3"},[t._v("Strawberry Wallpaper 后台统计平台")]),a("div",{staticClass:"sw-main_header_icon"}),a("div",{staticClass:"sw-main_header_date"},[t._v(t._s(t.currentDate))])]),a("div",{staticClass:"sw-row sw-row_one"},[a("div",{staticClass:"sw-main_data"},[a("div",{staticClass:"sw-modal-bk"}),a("div",{staticClass:"sw-modal-content"},[a("div",{staticClass:"sw-main_title"},[t._v("安装数量统计")]),a("div",{staticClass:"sw-main_data_title"},[a("div",{staticClass:"sw-main_data_title_t"},[t._v("安装总量:")]),a("div",{staticClass:"sw-main_data_title_c"},[t._v(t._s(t.totalNum))])]),a("div",{staticClass:"sw-main_data_title"},[a("div",{staticClass:"sw-main_data_title_t"},[t._v("最新的日活量:")]),a("div",{staticClass:"sw-main_data_title_c"},[t._v(t._s(t.activeNum))])])])]),a("div",{staticClass:"sw-main_system"},[a("div",{staticClass:"sw-modal-bk"}),a("div",{staticClass:"sw-modal-content"},[a("div",{staticClass:"sw-main_title"},[t._v("安装平台统计")]),a("div",{staticClass:"sw-main_system_content"},[a("div",{ref:"ref_system",staticClass:"sw-main_system_platform"})])])])]),a("div",{staticClass:"sw-row sw-main_add"},[a("div",{staticClass:"sw-modal-bk"}),a("div",{staticClass:"sw-modal-content"},[a("div",{staticClass:"sw-main_title"},[t._v("每日新增用户量统计")]),a("div",{ref:"ref_add",staticClass:"sw-main_add_content sw-row_content"})])]),a("div",{staticClass:"sw-row sw-main_active"},[a("div",{staticClass:"sw-modal-bk"}),a("div",{staticClass:"sw-modal-content"},[a("div",{staticClass:"sw-main_title"},[t._v("每日活跃用户量统计")]),a("div",{ref:"ref_active",staticClass:"sw-main_active_contentm sw-row_content"})])])])},M=[],$=(a("612f"),a("24ce")),j=a.n($),E=(a("5f33"),a("f91a"),function(t,e){var a={"M+":t.getMonth()+1,"d+":t.getDate(),"h+":t.getHours(),"m+":t.getMinutes(),"s+":t.getSeconds(),"q+":Math.floor((t.getMonth()+3)/3),"S+":t.getMilliseconds()};for(var n in/(y+)/i.test(e)&&(e=e.replace(RegExp.$1,"".concat(t.getFullYear()).substr(4-RegExp.$1.length))),a)new RegExp("(".concat(n,")")).test(e)&&(e=e.replace(RegExp.$1,1===RegExp.$1.length?a[n]:"00".concat(a[n]).substr("".concat(a[n]).length)));return e}),P=null,I=null,L=null,A=function(t,e){return{grid:{left:60,right:60,bottom:80,top:20},tooltip:{trigger:"axis",position:function(t){return[t[0],"10%"]}},xAxis:{type:"category",boundaryGap:!1,data:t,axisLine:{lineStyle:{color:"#aaa"}},axisTick:{lineStyle:{color:"#aaa"}},axisLabel:{color:"#aaa"}},yAxis:{type:"value",axisLine:{lineStyle:{color:"#aaa"}},axisTick:{lineStyle:{color:"#aaa"}},axisLabel:{color:"#aaa"}},dataZoom:[{type:"slider",start:0,end:100}],series:[{symbol:"none",sampling:"average",itemStyle:{color:"#2cade2"},data:e,type:"line"}]}},N=function(t){return{color:["#24feb4","#005eff"],tooltip:{trigger:"item",formatter:"{a} <br/>{b}: {c} ({d}%)"},series:[{name:"安装平台",type:"pie",radius:["50%","70%"],avoidLabelOverlap:!1,label:{normal:{show:!0,fontSize:20,formatter:"{b}:{c}台\n{d}%",textStyle:{fontSize:"20",fontWeight:"bold"}},emphasis:{show:!0,position:"center",textStyle:{fontSize:"30",fontWeight:"bold"}}},labelLine:{show:!0},data:t}]}},R={Macintosh:"Mac",Windows:"Win"},D={name:"swMain",data:function(){return{totalNum:"0",activeNum:"0",currentDate:""}},created:function(){var t=this;window.setInterval(function(){t.currentDate=E(new Date,"YYYY-MM-dd hh-mm-ss")},1e3)},mounted:function(){var t=this;g().then(function(e){t.echartsSystemInit(e.platform),t.echartsAddInit(e.register),t.echartsActiveInit(e.active),t.totalNum=e.total_num,t.activeNum=e.active_num}).catch(function(t){console.log(t)})},methods:{echartsSystemInit:function(t){var e=t.map(function(t){return{value:t.count,name:R[t.platform]||t.platform}});null===P&&(P=j.a.init(this.$refs.ref_system)),P.setOption(N(e))},echartsAddInit:function(t){var e=[],a=[];t.forEach(function(t){e.push(t.date),a.push(t.count)}),null===I&&(I=j.a.init(this.$refs.ref_add)),I.setOption(A(e,a))},echartsActiveInit:function(t){var e=[],a=[];t.forEach(function(t){e.push(t.date),a.push(t.count)}),null===L&&(L=j.a.init(this.$refs.ref_active)),L.setOption(A(e,a))}}},T=D,W=Object(l["a"])(T,k,M,!1,null,"4c4956b6",null),Y=W.exports;n["default"].use(f["a"]);var z=new f["a"]({routes:[{path:"/login",name:"/login",component:O},{path:"/",name:"main",component:Y}]}),J=a("52c1");n["default"].use(J["a"]);var q=new J["a"].Store({state:{},mutations:{},actions:{}});a("5b17"),a("a2f0");n["default"].config.productionTip=!1,n["default"].use(i.a),new n["default"]({router:z,store:q,render:function(t){return t(d)}}).$mount("#app")},"9e74":function(t,e,a){},a2f0:function(t,e,a){}});
//# sourceMappingURL=app.2b6a609a.js.map