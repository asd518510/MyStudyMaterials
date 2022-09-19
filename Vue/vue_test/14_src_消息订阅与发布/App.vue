<template>
  <div class="app">
    <h1>{{msg}}</h1>
    <School />
    <Student />
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
        msg:'你好啊！'
    }
  }
}
</script>

<style scoped>
    .app{
        background-color: gray;
        padding: 5px;
    }
</style>