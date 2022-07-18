<template>
  <div class="container">
    <h1>{{ message }}</h1>
  </div>
</template>

<script>
import { ref } from 'vue'
import { onMounted } from '@vue/runtime-core'
import axios from 'axios'
import { useStore } from 'vuex'
export default {
  name: "Home-item",
  setup() {
    const message = ref('You are not logged in!')
    const store = useStore()

    onMounted(async () => {
      // user情報を取得
      // ログイン情報はCookieに保存しているのでリクエストするだけでOK

      try {
        const { data } = await axios.get('user')
        message.value = `Hi ${data.first_name} ${data.last_name}`

        // actionsに指定したパラメターを設定
        await store.dispatch('setAuth', true)
      } catch(e) {
        await store.dispatch('setAuth', false)
      }
    })

    return {
      store,
      message
    }
  }
}
</script>