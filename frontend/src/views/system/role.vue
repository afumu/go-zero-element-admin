<template>
  <TableLayout>
  <!-- <TableLayout> -->
    <!-- 搜索表单 -->
    <div slot="search-form">
      <el-collapse v-model="activeNames" accordion>
        <el-collapse-item title="查询条件" name="1">
          <template slot="title">
            <div class="kj-title">
              <span style="margin: 0 10px">查询条件</span>
              <section>
                <el-button
                  type="primary"
                  size="mini"
                  icon="el-icon-search"
                  @click.stop="search"
                  >搜索</el-button
                >
                <el-button @click.stop="reset" size="mini" icon="el-icon-search"
                  >重置</el-button
                >
              </section>
            </div>
          </template>
          <el-form
            ref="searchForm"
            :model="searchForm"
            label-width="100px"
            inline
          >
            <el-form-item label="角色名称" prop="roleName">
              <searchInput
                v-model="searchForm.roleName"
                placeholder="请输入角色名称"
                @keypress.enter.native="search"
              ></searchInput>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
      <!-- <div style="height: 10px"></div> -->
    </div>
    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button
            type="primary"
            @click="$refs.operaRoleWindow.open('新建角色')"
            icon="el-icon-plus"
            >新建</el-button
          >
        </li>
        <!-- <li><el-button @click="deleteByIdInBatch" icon="el-icon-delete">删除</el-button></li> -->
      </ul>
      <el-table
          border
        v-loading="isWorking.search"
        :data="tableData.list"
        :default-sort="{ prop: 'createTime', order: 'descending' }"
        stripe
        @selection-change="handleSelectionChange"
        @sort-change="handleSortChange"
      >
        <el-table-column
          prop="code"
          label="角色编码"
          fixed="left"
          min-width="100px"
        ></el-table-column>
        <el-table-column
          prop="name"
          label="角色名称"
          fixed="left"
          min-width="100px"
        ></el-table-column>
        <el-table-column
          prop="createdAt"
          label="创建时间"
          min-width="140px"
        ></el-table-column>
        <el-table-column
          label="操作"
          align="center"
          min-width="150"
          fixed="right"
        >
          <template slot-scope="{ row }">
            <el-button
              type="text"
              @click="$refs.operaRoleWindow.open('编辑角色', row)"
              icon="el-icon-edit"
              >编辑</el-button
            >
            <el-button
              type="text"
              @click="$refs.selectUser.open(row)"
              icon="el-icon-user"
              >用户</el-button
            >
            <el-button
              type="text"
              @click="$refs.menuConfigWindow.open(row)"
              icon="el-icon-menu"
              >授权</el-button
            >
            <el-button
              v-if="!row.fixed"
              type="text"
              @click="deleteById(row)"
              icon="el-icon-delete"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <pagination
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
        :pagination="tableData.pagination"
      ></pagination>
    </template>
    <!-- 新建/修改 -->
    <OperaRoleWindow
      ref="operaRoleWindow"
      @fail="fail"
      @success="handlePageChange(tableData.pagination.pageIndex)"
    />
    <!-- 配置权限 -->
    <PermissionConfigWindow
      ref="permissionConfigWindow"
      @success="handlePageChange(tableData.pagination.pageIndex)"
    />
    <!-- 授权菜单 -->
    <MenuConfigWindow
      ref="menuConfigWindow"
      @success="handlePageChange(tableData.pagination.pageIndex)"
    />

    <!-- 用户 -->
    <selectUser
      ref="selectUser"
      @success="handlePageChange(tableData.pagination.pageIndex)"
    />
  </TableLayout>
</template>

<script>
import Pagination from '@/components/common/Pagination'
import TableLayout from '@/layouts/TableLayout'
import BaseTable from '@/components/base/BaseTable'
import OperaRoleWindow from '@/components/system/role/OperaRoleWindow'
import PermissionConfigWindow from '@/components/system/role/PermissionConfigWindow'
import MenuConfigWindow from '@/components/system/role/MenuConfigWindow'
import selectUser from '@/components/system/role/selectUser'

import searchInput from '@/components/common/searchInput'
export default {
  name: 'SystemRole',
  extends: BaseTable,
  components: {
    MenuConfigWindow,
    PermissionConfigWindow,
    OperaRoleWindow,
    TableLayout,
    Pagination,
    searchInput,
    selectUser,
  },
  data() {
    return {
      // 搜索
      searchForm: {
        roleCode: '',
        roleName: '',
      },
      activeNames: '1',
    }
  },
  created() {
    this.config({
      module: '角色',
      api: '/system/role',
      'field.main': 'roleName',
    })
    this.search()
  },
  methods: {
    fail(){
    //   this.$tip.warning("编码重复")
    }
  },
}
</script>
