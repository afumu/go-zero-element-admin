<template>
  <el-menu-item v-if="menu.children == null || menu.children.length == 0" :key="menu.index" :index="menu.path">
    <i :class="menu.meta.icon"></i>
    <span slot="title">{{menu.meta.title}}</span>
  </el-menu-item>
  <el-submenu v-else :index="menu.path">
    <template slot="title">
      <i :class="menu.meta.icon"></i>
      <span slot="title">{{menu.meta.title}}</span>
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
