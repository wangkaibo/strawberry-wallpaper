<template>
  <div class="hello">
    <el-button type="" @click="onClick">点击按钮请求</el-button>
  </div>
</template>

<script>
import {postRegister} from './api.js'
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  methods: {
    onClick(){
      postRegister({
        uid:'123456'
      })
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
