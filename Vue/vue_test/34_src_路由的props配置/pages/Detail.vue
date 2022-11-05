<template>
  <div>
    <ul>
      <li>消息编号：{{ id }}</li>
      <li>消息标题：{{ title }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'Detail',
  props: ['id', 'title'],
  computed: {
    // id(){
    //     return this.$route.query.id
    // },
    // title(){
    //     return this.$route.query.title
    // }
  },
  mounted() {},
}
</script>
