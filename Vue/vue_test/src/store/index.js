//该文件用于创建Vuex中最为核心的store
import Vue from 'vue'
//引入Vuex
import Vuex from 'vuex'
//使用Vuex插件
Vue.use(Vuex)

//求和相关的配置
const countOptions = {
    namespaced: true,
    actions: {
        incrementOdd(context, value) {
            if (context.state.sum % 2) {
                context.commit('Increment', value)
            }
        },
        incrementWait(context, value) {
            setTimeout(() => {
                context.commit('Increment', value)
            }, 500)
        }
    },
    mutations: {
        Increment(state, value) {
            state.sum += value
        },
        Decrement(state, value) {
            state.sum -= value
        },
    },
    state: {
        sum: 0, //当前的和
        province: '广东省',
        city: '深圳市',
    },
    getters: {
        bigSum(state) {
            return state.sum * 10
        }
    }
}

//人员相关的配置
const personOptions = {
    namespaced: true,
    actions: {},
    mutations: {
        Add_Person(state, value) {
            state.personList.unshift(value)
        }
    },
    state: {
        personList: [
            { id: '001', name: '张三' }
        ]
    },
    getters: {}
}

//创建并导出store
export default new Vuex.Store({
    modules: {
        countAbout: countOptions,
        personAbout: personOptions
    }
})
