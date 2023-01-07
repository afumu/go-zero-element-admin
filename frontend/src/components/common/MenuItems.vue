<template>
  <el-menu-item v-if="menu.children == null || menu.children.length == 0" :key="menu.index" :index="menu.url">
    <i :class="menu.icon"></i>
    <span slot="title">{{menu.name}}</span>
  </el-menu-item>
  <el-submenu v-else :index="menu.url">
    <template slot="title">
      <i :class="menu.icon"></i>
      <span slot="title">{{menu.name}}</span>
    </template>
    <MenuItems v-for="child in menu.children" :menu="child" :key="child.index"/>
  </el-submenu>
</template>

<script>
export default {
  name: 'MenuItems',
  props: {
    menu: {
      type: Object,
      required: true
    }
  }
}
</script>
