<template>
  <TableLayout class="menu-layout">
     <!-- <TableLayout class="menu-layout"> -->
    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button type="primary" @click="$refs.operaMenuWindow.open('新建一级菜单')" icon="el-icon-plus">新建</el-button>
        </li>
        <!-- <li><el-button @click="deleteByIdInBatch" icon="el-icon-delete">批量删除</el-button></li> -->
        <!-- <li><el-button @click="sort('top')" :loading="isWorking.sort" icon="el-icon-sort-up">上移</el-button></li>
        <li><el-button @click="sort('bottom')" :loading="isWorking.sort" icon="el-icon-sort-down">下移</el-button></li> -->
      </ul>
      <el-table border ref="table" v-loading="isWorking.search" :data="tableData.list" :tree-props="{ children: 'children', hasChildren: 'hasChildren' }" row-key="id" stripe @selection-change="handleSelectionChange">
        <el-table-column align="center" type="selection" width="55" fixed="left"></el-table-column>
        <el-table-column prop="name" label="菜单名称" fixed="left" min-width="120px"></el-table-column>
        <el-table-column prop="icon" label="图标" min-width="80px" class-name="table-column-icon">
          <template slot-scope="{ row }">
            <i v-if="row.icon != null && row.icon !== ''" :class="{ [row.icon]: true }"></i>
            <template v-else>未设置</template>
          </template>
        </el-table-column>
        <el-table-column prop="menuType" label="菜单类型" min-width="120px">
          <template slot-scope="{ row }">
            <span v-if="row.menuType == 2">按钮/权限</span>
            <span v-else>菜单</span>
          </template>
        </el-table-column>
        <el-table-column prop="url" label="访问路径" min-width="140px"></el-table-column>
        <el-table-column prop="sort" label="排序" min-width="50px"></el-table-column>
        <el-table-column prop="description" label="备注" min-width="140px"></el-table-column>
        <el-table-column prop="createdAt" label="创建时间" min-width="140px"></el-table-column>
        <el-table-column prop="updatedAt" label="更新时间" min-width="140px"></el-table-column>
        <el-table-column label="操作" min-width="220" align="center" fixed="right">
          <template slot-scope="{ row }">
            <el-button style="margin-left: 0;!important" type="text" icon="el-icon-edit" @click="$refs.operaMenuWindow.open('编辑菜单', row)">编辑</el-button>
            <el-button style="margin-left:15px;!important" type="text" icon="el-icon-plus" @click="$refs.operaMenuWindow.open('新建子菜单', null, row)">新建子菜单</el-button>
            <el-button style="margin-left: 15px;!important" v-if="!row.fixed" type="text" icon="el-icon-delete" @click="deleteById(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </template>
    <!-- 新建/修改 -->
    <OperaMenuWindow ref="operaMenuWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
  </TableLayout>
</template>

<script>
import TableLayout from "@/layouts/TableLayout";
import BaseTable from "@/components/base/BaseTable";
import OperaMenuWindow from "@/components/system/permission/OperaMenuWindow";
import { fetchTree, updateStatus, sort } from "@/api/system/menu";
export default {
  extends: BaseTable,
  components: { OperaMenuWindow, TableLayout },
  data() {
    return {
      // 是否正在处理中
      isWorking: {
        sort: false,
      },
    };
  },
  methods: {
    // 查询数据
    handlePageChange() {
      this.isWorking.search = true;
      fetchTree()
        .then((records) => {
          console.info('----------------records-----------',records)
          this.tableData.list = records.menuList;
          // console.log('this.tableData.list', this.tableData.list)
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        })
        .finally(() => {
          this.isWorking.search = false;
        });
    },
    // 排序
    sort(direction) {
      if (this.isWorking.sort) {
        return;
      }
      if (this.tableData.selectedRows.length === 0) {
        this.$tip.warning("请选择一条数据");
        return;
      }
      if (this.tableData.selectedRows.length > 1) {
        this.$tip.warning("排序时仅允许选择一条数据");
        return;
      }
      const menuId = this.tableData.selectedRows[0].id;
      // 找到菜单范围
      let menuPool;
      for (const rootMenu of this.tableData.list) {
        const parent = this.__findParent(menuId, rootMenu);
        if (parent != null) {
          menuPool = parent.children;
        }
      }
      menuPool = menuPool || this.tableData.list;
      const menuIndex = menuPool.findIndex((menu) => menu.id === menuId);
      // 上移校验
      if (direction === "top" && menuIndex === 0) {
        this.$tip.warning("菜单已到顶部");
        return;
      }
      // 下移校验
      if (direction === "bottom" && menuIndex === menuPool.length - 1) {
        this.$tip.warning("菜单已到底部");
        return;
      }
      this.isWorking.sort = true;
      sort({
        id: this.tableData.selectedRows[0].id,
        direction,
      })
        .then(() => {
          if (direction === "top") {
            menuPool.splice(menuIndex, 0, menuPool.splice(menuIndex - 1, 1)[0]);
          } else {
            menuPool.splice(menuIndex, 0, menuPool.splice(menuIndex + 1, 1)[0]);
          }
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        })
        .finally(() => {
          this.isWorking.sort = false;
        });
    },
    // 启用/禁用菜单
    switchDisabled(row) {
      if (!row.disabled) {
        this.__updateMenuStatus(row);
        return;
      }
      this.$dialog
        .disableConfirm(`确认禁用 ${row.name} 菜单吗？`)
        .then(() => {
          this.__updateMenuStatus(row);
        })
        .catch(() => {
          row.disabled = !row.disabled;
        });
    },
    // 查询父节点
    __findParent(id, parent) {
      if (parent.children === 0) {
        return;
      }
      for (const menu of parent.children) {
        if (menu.id === id) {
          return parent;
        }
        if (menu.children.length > 0) {
          const m = this.__findParent(id, menu);
          if (m != null) {
            return m;
          }
        }
      }
      return null;
    },
    // 修改菜单状态
    __updateMenuStatus(row) {
      updateStatus({
        id: row.id,
        parentId: row.parentId,
        disabled: row.disabled,
      })
        .then(() => {
          this.$tip.apiSuccess("修改成功");
        })
        .catch((e) => {
          row.disabled = !row.disabled;
          this.$tip.apiFailed(e);
        });
    },
  },
  created() {
    this.config({
      module: "菜单",
      api: "/system/menu",
    });
    this.search();
  },
};
</script>

<style lang="scss" scoped>
@import "@/assets/style/variables.scss";
.menu-layout {
  /deep/ .table-content {
    margin-top: 0;
  }
}
// 图标列
.table-column-icon {
  // element-ui图标
  i {
    background-color: $primary-color;
    opacity: 0.72;
    font-size: 20px;
    color: #fff;
    padding: 4px;
    border-radius: 50%;
  }
}
</style>
