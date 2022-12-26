<template>
  <div class="menu" :class="{ collapse: menuData.collapse }">
    <div class="logo">
      <div><img src="../../../public/logo_title.png" /></div>
      <h1 :class="{ hidden: menuData.collapse }">{{title}}</h1>
    </div>
    <scrollbar>
      <el-menu
        ref="menu"
        :default-active="activeIndex"
        text-color="#fff"
        active-text-color="#fff"
        :unique-opened="true"
        :collapse="menuData.collapse"
        :default-openeds="defaultOpeneds"
        :collapse-transition="false"
        @select="handleSelect"
      >
        <MenuItems
          v-for="menu in menuData.list"
          :key="menu.index"
          :menu="menu"
          :is-root-menu="true"
        />
      </el-menu>
    </scrollbar>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import MenuItems from './MenuItems'
import Scrollbar from './Scrollbar'
export default {
  name: 'Menu',
  components: { Scrollbar, MenuItems },
  computed: {
    ...mapState(['menuData']),
    // 选中的菜单index
    activeIndex () {
      let path = this.$route.path
      if (path.endsWith('/')) {
        path = path.substring(0, path.length - 1)
      }
      const menuConfig = this.__getMenuConfig(path, 'path', this.menuData.list)
      if (menuConfig == null) {
        return null
      }
      return menuConfig.path
    },
    // 默认展开的菜单index
    defaultOpeneds () {
      return this.menuData.list.map((menu) => menu.index)
    }
  },
  methods: {
    // 处理菜单选中
    handleSelect (menuIndex) {
      const menuConfig = this.__getMenuConfig(
        menuIndex,
        'path',
        this.menuData.list
      )
      // 找不到页面
      try {
        require('@/views' + menuConfig.path)
      } catch (e) {
        this.$tip.error(
          '未找到页面文件@/views' +
            menuConfig.path +
            '.vue，请检查菜单路径是否正确'
        )
      }
      // 点击当前菜单不做处理
      if (menuConfig.path === this.$route.path) {
        return
      }
      if (menuConfig.path == null || menuConfig.path.trim().length === 0) {
        return
      }
      this.$router.push(menuConfig.path)
    },
    // 获取菜单配置
    __getMenuConfig (value, key, menus) {
      for (const menu of menus) {
        if (menu[key] === value) {
          return menu
        }
        if (menu.children != null && menu.children.length > 0) {
          const menuConfig = this.__getMenuConfig(value, key, menu.children)
          if (menuConfig != null) {
            return menuConfig
          }
        }
      }
      return null
    }
  },
  data(){
    return{
        title:window._CONFIG['systemTitle']
    }
  }
  // watch: {
  //   $route: {
  //     handler: function (val, oldVal) {
  //       console.log('1321123',val);
  //     },
  //     // 深度观察监听
  //     deep: true,
  //   },
  // },
}
</script>

<style lang="scss" scoped>
@import "@/assets/style/variables.scss";
.menu {
  height: 100%;
  display: flex;
  flex-direction: column;
  background:#001529;
  // LOGO
  .logo {
    height: 60px;
    flex-shrink: 0;
    line-height: 60px;
    overflow: hidden;
    display: flex;
    background: #002140;
    padding: 0 16px;
    & > div {
      width: 32px;
      flex-shrink: 0;
      margin-right: 12px;
      img {
        width: 100%;
        flex-shrink: 0;
        vertical-align: middle;
        position: relative;
        top: -2px;
      }
    }
    h1 {
      font-size: 16px;
      font-weight: 500;
      transition: opacity ease 0.3s;
      overflow: hidden;
      &.hidden {
        opacity: 0;
      }
    }
  }
}
</style>
<style lang="scss">
@import "@/assets/style/variables.scss";
// 菜单样式
.el-menu {
  border-right: 0 !important;
  user-select: none;
  background: #001529 !important;
  .el-menu-item {
    background: #000C17;
    // 选中状态
    &.is-active {
      background: $primary-color !important;
    }
    // 悬浮
    &:hover {
      background-color: $primary-color - 12;
    }
    &:focus {
      background: $primary-color;
    }
  }
  // 子菜单
  .el-submenu {
    .el-submenu__title {
      background-color: #001529;
    }
    &.is-active {
      .el-submenu__title {
        background-color: #001529;
      }
      .el-menu .el-menu-item {
        background-color: #000C17 - 20;
        // 悬浮
        &:hover {
          background-color: #001529 - 30;
        }
      }
    }
    // 菜单上下箭头
    .el-submenu__title i {
      color: #f7f7f7;
    }
  }
  // 菜单图标
  i:not(.el-submenu__icon-arrow) {
    color: #f7f7f7 !important;
    position: relative;
    top: -1px;
    // 自定义图标
    &[class^="eva-icon-"] {
      width: 24px;
      margin-right: 5px;
      //background-size: 15px;
    }
  }
}
</style>
