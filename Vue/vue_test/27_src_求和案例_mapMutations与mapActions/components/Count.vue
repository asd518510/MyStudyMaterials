<template>
  <div>
    <h1>当前求和为：{{ sum }}</h1>
    <h3>当前求和放大10倍为：{{ bigSum }}</h3>
    <h3>我在{{ province }}{{ city }}</h3>
    <select v-model.number="n">
      <option value="1">1</option>
      <option value="2">2</option>
      <option value="3">3</option>
    </select>
    <button @click="increment(n)">+</button>
    <button @click="decrement(n)">-</button>
    <button @click="incrementOdd(n)">当前求和为奇数再加</button>
    <button @click="incrementWait(n)">等一等再加</button>
  </div>
</template>

<script>
import { mapState, mapGetters, mapMutations, mapActions } from 'vuex'
export default {
  name: 'Count',
  data() {
    return {
      n: 1, //用户选择的数字
    }
  },
  computed: {
    //借助mapState生成计算属性，从state中读取数据。（对象写法）
    // ...mapState({ sum: 'sum', province: 'province', city: 'city' }),

    //借助mapState生成计算属性，从state中读取数据。（数组写法）
    ...mapState(['sum', 'province', 'city']),

    /* ************************* */
    ...mapGetters(['bigSum']),
  },
  methods: {
    /* increment() {
      this.$store.commit('Increment', this.n)
    },
    decrement() {
      this.$store.commit('Decrement', this.n)
    }, */
    //借助mapMutations生成对应的方法，方法中会调用commit去联系mutations
    ...mapMutations({ increment: 'Increment', decrement: 'Decrement' }),

    /* incrementOdd() {
      this.$store.dispatch('incrementOdd', this.n)
    },
    incrementWait() {
      this.$store.dispatch('incrementWait', this.n)
    }, */
    ...mapActions(['incrementOdd', 'incrementWait']),
  },
  mounted() {},
}
</script>
<style lang="css">
button {
  margin-left: 5px;
}
</style>
