<template>
  <div class="table-layout">
    <!-- 头部 -->
    <!-- <div v-if="withBreadcrumb" class="table-header">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item v-for="path in paths" :key="path">{{path}}</el-breadcrumb-item>
      </el-breadcrumb>
    </div> -->
    <div v-if="withBreadcrumb" class="table-header">
        <el-tabs :tab-position="tabPosition" closable  type="card" @tab-remove="removeTab" v-model="editableTabsValue">
            <template v-for="item in paths">
                <el-tab-pane :key="item.path" :name="item.name">
                    <span slot="label" @click="toPath(item.path)">{{item.title}}</span>
                </el-tab-pane>
            </template>
        </el-tabs>
    </div>
    <Profile :roles="roles" :permissions="permissions">
      <!-- 搜索表单部分 -->
      <div class="table-search-form">
        <div class="form-wrap">
          <slot name="search-form"></slot>
        </div>
      </div>
      <slot name="space"></slot>
      <!-- 列表和分页部分 -->
      <div class="table-content">
        <div class="table-wrap">
          <slot name="table-wrap"></slot>
        </div>
      </div>
      <slot></slot>
    </Profile>
  </div>
</template>
<script>
import Profile from '../components/common/Profile'
export default {
  name: 'TableLayout',
  components: { Profile },
  props: {
    // 角色
    roles: {
      type: Array
    },
    // 权限
    permissions: {
      type: Array
    },
    // 是否展示头部面包屑
    withBreadcrumb: {
      type: Boolean,
      default: true
    },
    tabPosition:{
        type:String,
        default: 'top'
    },
  },
  computed: {
    paths () {
      return this.$store.state.paths
    }
  },
    created(){
        let path = {
          path:this.$route.path,
          name:this.$route.meta.title,
          title:this.$route.meta.title
        }
      this.$store.commit('paths',path)
      this.editableTabsValue = path.title
    },
    data(){
        return{
            editableTabsValue:''
        }
    },
    methods:{
        toPath(path){
            if(path == this.$route.path) return
            this.$router.push({path})
        },
        removeTab(targetName){
            this.paths.forEach((item,index) => {
                if(this.paths.length>1){
                    if(targetName === item.name){
                        this.editableTabsValue = item.title
                        this.toPath(this.paths[index-1].path)
                    }
                }
                
            });
            this.$store.commit('remove',targetName)
        },
    }
}
</script>

<style lang="scss">
@import "@/assets/style/variables.scss";
.table-layout {
  height: 100%;
  display: flex;
  flex-direction: column;
  .not-allow-wrap {
    padding-top: 0;
  }
}
// 头部
.table-header {
  overflow: hidden;
  padding: 0 16px;
  flex-shrink: 0;
  background: #fff;
  margin-bottom: 12px;
  // 页面路径
  .el-breadcrumb {
    .el-breadcrumb__item {
      .el-breadcrumb__inner {
        color: #ABB2BE;
        font-size: 12px;
      }
      &:last-of-type .el-breadcrumb__inner {
        color: #606263;
        font-size: 14px;
      }
    }
  }
}
// 搜索
.table-search-form {
  display: flex;
  flex-wrap: wrap;
  padding: 0 16px;
  .form-wrap {
    padding: 16px 16px 0 16px;
    width: 100%;
    background: #fff;
    &:empty {
      padding: 0;
    }
  }
  section {
    display: inline-block;
    margin-left: 16px;
    margin-bottom: 18px;
  }
}
// 列表和分页
.table-content {
  // margin-top: 10px;
  padding: 0 16px;
  .table-wrap {
    padding: 16px 16px 0 16px;
    background: #fff;
    // 工具栏
    .toolbar {
      border-bottom: 1px solid #eee;
      padding-bottom: 15px;
      li {
        display: inline-block;
        margin-right: 6px;
      }
    }
    // 表格
    .el-table {
      th {
        .cell {
          color: #666;
        }
      }
      // 复选框列
      .el-table-column--selection {
        .cell {
          text-align: center !important;
        }
      }
      // 多值字段
      .table-column-strings {
        ul {
          li {
            display: inline-block;
            background: #eee;
            border-radius: 3px;
            padding: 0 3px;
            margin-right: 3px;
            margin-bottom: 3px;
          }
        }
      }
      // 树视觉调整
      [class*=el-table__row--level] .el-table__expand-icon {
        position: relative;
        left: -6px;
        margin-right: 0;
      }
    }
    // 分页
    .table-pagination {
      padding: 16px 0;
      text-align: left;
    }
  }
}
</style>
