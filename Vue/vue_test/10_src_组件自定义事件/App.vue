<template>
  <div class="app">
    <h1>{{msg}},学生姓名是：{{studentName}}</h1>

    <!--通过父组件给子组件传递函数类型的props实现：子给父传递数据-->
    <School :getSchoolName="getSchoolName"></School>

    <!--通过父组件给子组件绑定一个自定义事件实现：子给父传递数据(第一种写法：使用@或v-on)-->
    <!-- <Student @studyTest="getStudetName" @demo="m1"></Student> -->
    
    <!--通过父组件给子组件绑定一个自定义事件实现：子给父传递数据(第二种写法：使用ref)-->
    <Student ref="student" @click.native="show"/>
  </div>
</template>

<script>
import School from '../src/components/School.vue'
import Student from '../src/components/Student.vue'


export default {
  name: 'App',
  components: {
    School,
    Student,
  },
  data() {
    return {
        msg:'你好啊！',
        studentName:''
    }
  },
  methods: {
    getSchoolName(name){
        console.log('App收到学校名：',name)
    },
    getStudetName(name,...params){
        console.log('App收到学生名：',name,params)
        this.studentName=name
    },
    m1(){
        console.log('demo事件被触发了！')
    },
    show(){
        alert(123)
    }
  },
  mounted() {
    // setTimeout(()=>{
        this.$refs.student.$on('studyTest',this.getStudetName)
    // },3000)
    // this.$refs.student.$once('studyTest',this.getStudetName)//绑定自定义事件（一次性）
  },
}
</script>

<style scoped>
    .app{
        background-color: gray;
        padding: 5px;
    }
</style>