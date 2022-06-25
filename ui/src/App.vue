<template>
  <div>
    <Nav></Nav>
    <router-view></router-view>
  </div>
</template>

<script>
import Nav from "@/components/Nav"
export default {
  components: {
    Nav
  },
  Name: 'App'
}
</script>
<!-- <template>
  <h2>router</h2>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
}
</script> -->