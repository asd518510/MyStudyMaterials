//人员相关的配置
import axios from 'axios'
import { nanoid } from 'nanoid'

export default {
    namespaced: true,
    actions: {
        addPresonWang(context, value) {
            if (value.name.indexOf('王') === 0) {
                context.commit('Add_Person', value)
            } else {
                alert('添加的人必须姓王！')
            }
        },
        addPersonServer(context) {
            axios.get('https://api.uixsj.cn/hitokoto/get?type=social').then(
                response => {
                    context.commit('Add_Person', { id: nanoid(), name: response.data })
                },
                error => {
                    alert(error.message)
                }
            )
        }
    },
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
    getters: {
        firstPersonName(state) {
            return state.personList[0].name
        }
    }
}