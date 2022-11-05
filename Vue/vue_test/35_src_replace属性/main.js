//引入Vue
import Vue from 'vue'
//引入App
import App from './App.vue'
//引入VueRouter
import VueRouter from 'vue-router'
//引入一个路由器
import router from './router'

//关闭vue的生产提示
Vue.config.productionTip = false
//应用插件
Vue.use(VueRouter)

//创建Vue实例对象
new Vue({
    el: '#app',
    render: (h) => h(App),
    router: router
})
