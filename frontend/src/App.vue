<template>
  <div id="app">
    <PageMain />
  </div>
</template>

<script>
import PageMain from './components/PageMain.vue';

export default {
  name: 'App',
  components: {
    PageMain,
  },
};
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: Arial, sans-serif;
  background-color: #f5f5f5;
}
</style>