<template>
  <div class="container">
    <Category title="美食">
      <img src="https://www.gomaji.com/blog/wp-content/uploads/2020/04/Da-Tung-Food-Banner-e1588216225724.jpg" alt="" />
    </Category>
    <Category title="游戏">
      <ul>
        <li v-for="(g, index) in games" :key="index">{{ g }}</li>
      </ul>
    </Category>
    <Category title="电影">
      <video controls src="http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4"></video>
    </Category>
  </div>
</template>

<script>
import Category from './components/Category'
export default {
  name: 'App',
  components: { Category },
  data() {
    return {
      foods: ['火锅', '烧烤', '小龙虾', '牛排'],
      games: ['LOL', '劲舞团', 'CF'],
      films: ['《教父》', '《砖家》', '《狗屁砖家》'],
    }
  },
}
</script>
<style>
.container {
  display: flex;
  justify-content: space-around;
}
</style>
