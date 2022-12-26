<template>
  <div style="width:1663px; margin-left:16px;" >
    <iframe  :id="id" :src="url" frameborder="0" width="100%" height="800px" scrolling="auto"></iframe>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        id:"sqldruid",
        url:""
      }
    },
    mounted () {
      this.goUrl()
    },
    updated () {
      this.goUrl()
    },
    watch: {
      $route(to, from) {
        this.goUrl();
      }
    },
    methods: {
      goUrl () {
        let url = window._CONFIG['domianURL']+'/druid/login.html'
        if (url !== null && url !== undefined) {
          this.url=url;
        }
      }
    }
  }
</script>

<style>
</style>