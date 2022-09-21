//该文件用于创建Vuex中最为核心的store
import Vue from 'vue'
//引入Vuex
import Vuex from 'vuex'
//使用Vuex插件
Vue.use(Vuex)

//准备actions——用于响应组件中的动作
const actions = {
    incrementOdd(context, value) {
        if (context.state.sum % 2) {
            context.commit('Increment', value)
        }
    },
    incrementWait(context, value) {
        setTimeout(() => {
            context.commit('Increment', value)
        }, 500)
    },
}

//准备mutations——用于操作数据（State）
const mutations = {
    Increment(state, value) {
        state.sum += value
    },
    Decrement(state, value) {
        state.sum -= value
    },
}

//准备state——用于存储数据
const state = {
    sum: 0, //当前的和
}

//创建并导出store
export default new Vuex.Store({
    actions,
    mutations,
    state
})
