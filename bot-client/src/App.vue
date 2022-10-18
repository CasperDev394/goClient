<script>
import { defineComponent, onMounted, ref } from 'vue';

export default defineComponent({
  name: 'App',
  setup() {
    onMounted(async () => {
        const response = await fetch('/bot/ping', {
          method: 'GET',
        });

        const body = await response.json();

        ping.value = body
      })

      return {
        ping
      }
  }
})
</script>

<template>
    <h1>App Work!</h1>
    <h3 v-if="ping">This is Ping {{ ping }}</h3>
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
