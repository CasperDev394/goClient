<script>
import { defineComponent, onMounted, ref } from 'vue';
import getPing from './components/getPing.vue'

export default defineComponent({
  name: 'App',
  components:{
    getPing
  }
  //setup() {
  //  onMounted(async () => {
  //      const response = await fetch('/bot/ping', {
  //        method: 'GET',
  //      });

  //      const body = await response.json();

  //      ping.value = body
  //    })

  //    return {
  //      ping
  //    }
  //}
})
</script>

<template>
  <div class="root">
    <get-ping></get-ping>
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
