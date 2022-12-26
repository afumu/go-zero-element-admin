<template>
  <userCommon :permissions="['system:user:query']"></userCommon>
</template>

<script>
export default {
  components: {
    userCommon: (resolve) =>
      require(['@/components/system/user/UserCommon'], resolve),
  },
  data() {
    return {}
  },
}
</script>

<style scoped lang="scss"></style>
