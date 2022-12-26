<template>
  <TableLayout>
    <!-- 搜索表单 -->

    <div slot="search-form">
      <el-collapse v-model="activeNames" accordion>
        <el-collapse-item title="查询条件" name="1">
          <template slot="title">
            <div class="kj-title">
              <span style="margin: 0 10px">查询条件</span>
              <section>
                <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="searchUser">搜索
                </el-button>
                <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置
                </el-button>
              </section>
            </div>
          </template>
          <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
            <el-form-item label="用户账号" prop="username">
              <search-input class="search_width" v-model="searchForm.username" v-trim placeholder="请输入用户账号" @keypress.enter.native="search" />
            </el-form-item>

            <el-form-item label="用户姓名" prop="realname">
              <search-input class="search_width" v-model="searchForm.realname" v-trim placeholder="请输入姓名" @keypress.enter.native="search" />
            </el-form-item>

            <el-form-item label="用户性别" prop="sex">
              <el-select clearable class="search_width" v-model="searchForm.sex" placeholder="请选择用户性别">
                <el-option v-for="item in this.$dict.getAllDictItemByCode(
                    this.$consts.dicConstant.sex
                  )" :key="item.value" :label="item.text" :value="parseInt(item.value)"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item label="账号状态" prop="status">
              <el-select clearable class="search_width" v-model="searchForm.status" placeholder="请选择账号状态">
                <el-option v-for="item in allStatus" :key="item.value" :label="item.text" :value="parseInt(item.value)"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>

    <!-- 表格和分页 -->
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button icon="el-icon-plus" type="primary" @click="addUser">新建</el-button>
        </li>
        <li>
          <el-button icon="el-icon-delete" type="primary" @click="deleteByIdInBatch">删除</el-button>
        </li>
        <li>
          <el-button icon="el-icon-delete" type="primary" @click="$refs.recycleBin.open()">回收站</el-button>
        </li>
      </ul>
      <el-table border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @selection-change="handleSelectionChange" @sort-change="handleSortChange">
        <el-table-column type="selection" width="55"></el-table-column>
        <el-table-column prop="realname" label="用户姓名" min-width="120px" fixed="left"></el-table-column>
        <el-table-column prop="username" label="用户账号" min-width="120px"></el-table-column>
        <el-table-column prop="orgCodeTxt" label="部门" min-width="120px"></el-table-column>
        <el-table-column prop="sex_dictText" label="性别" min-width="80px"></el-table-column>
        <el-table-column prop="birthday" label="出生日期" min-width="120px"></el-table-column>
        <el-table-column prop="phone" label="手机号" min-width="120px"></el-table-column>
        <el-table-column prop="status_dictText" label="状态" min-width="80px"></el-table-column>
        <el-table-column align="center" label="操作" min-width="120px" fixed="right">
          <template slot-scope="{ row }">
            <el-button style="margin-right: 10px !important;" type="text" icon="el-icon-edit" @click="editUser(row)">编辑</el-button>
            <el-dropdown>
              <span class="el-dropdown-link">
                <el-button type="text" style="padding-right: 0px !important">更多</el-button>
                <i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>
                  <el-button type="text" icon="el-icon-user-solid" @click="userDetail(row)">详情</el-button>
                </el-dropdown-item>

                <el-dropdown-item>
                  <el-button type="text" icon="el-icon-edit" @click="$refs.resetPwdWindow.open(row)">修改密码</el-button>
                </el-dropdown-item>

                <el-dropdown-item>
                  <el-button type="text" icon="el-icon-delete" @click="deleteById(row)">删除</el-button>
                </el-dropdown-item>

                <el-dropdown-item>
                  <el-button type="text" icon="el-icon-lock" v-if="row.status === 1" @click="frozenBatchUser(row, 2)">冻结</el-button>
                </el-dropdown-item>

                <el-dropdown-item>
                  <el-button type="text" icon="el-icon-unlock" v-if="row.status === 2" @click="frozenBatchUser(row, 1)">解冻</el-button>
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </el-table>
      <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </template>
    <!-- 新建/修改 -->
    <OperaUserWindow ref="operaUserWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
    <!-- 重置密码 -->
    <ResetPwdWindow ref="resetPwdWindow" />

    <recycle-bin ref="recycleBin" @success="search()"></recycle-bin>
  </TableLayout>
</template>

<script>
import Pagination from "@/components/common/Pagination";
import TableLayout from "@/layouts/TableLayout";
import TreeSelect from "@/components/common/TreeSelect";
import BaseTable from "@/components/base/BaseTable";
import OperaUserWindow from "@/components/system/user/OperaUserWindow";
import RoleConfigWindow from "@/components/system/user/RoleConfigWindow";
import ResetPwdWindow from "@/components/system/user/ResetPwdWindow";
import DepartmentSelect from "@/components/common/DepartmentSelect";
import PositionSelect from "@/components/common/PositionSelect";
import RecycleBin from "./RecycleBin";
import Cookies from "js-cookie";

export default {
  extends: BaseTable,
  components: {
    PositionSelect,
    DepartmentSelect,
    ResetPwdWindow,
    RoleConfigWindow,
    OperaUserWindow,
    TableLayout,
    Pagination,
    TreeSelect,
    RecycleBin,
  },
  data() {
    return {
      // 搜索
      searchForm: {
        order: "desc",
        column: "createTime",
        status: null,
        username: null, // 用户账号
        realname: null, // 用户姓名
        sex: null, // 性别
      },
      allStatus: this.$dict.getAllDictItemByCode(
        this.$consts.dicConstant.userStatus
      ),
      allCompany: [],
      allParty: [],
      allDepart: [],
      dialogVisible: false,
      token: { "X-Access-Token": Cookies.get("X-Access-Token") },
      activeNames: "1",
    };
  },
  props: {
    currentUserCompanyId: {
      type: String,
      default: null,
    },
    companyUserFlag: {
      type: Boolean,
      default: false,
    },
    permissions: {
      type: Array,
    },
  },
  methods: {
    // 填充部门树
    fillDepartData(list, pool) {
      for (const dept of pool) {
        const deptNode = {
          id: dept.orgCode,
          label: dept.title,
        };
        list.push(deptNode);
        if (dept.children != null && dept.children.length > 0) {
          deptNode.children = [];
          this.fillDepartData(deptNode.children, dept.children);
          if (deptNode.children.length === 0) {
            deptNode.children = undefined;
          }
        }
      }
    },
    searchUser() {
      this.search();
    },
    addUser() {
      this.$refs.operaUserWindow.withFooter = true;
      this.$refs.operaUserWindow.editFlag = false;
      this.$refs.operaUserWindow.open("新建用户");
    },
    userDetail(row) {
      if (row.partyGroupId) {
        row.partyId = row.partyGroupId;
      }
      this.beforeOperateUser(row);
      this.$refs.operaUserWindow.withFooter = false;
      this.$refs.operaUserWindow.open("用户详情", row);
    },
    editUser(row) {
      // 如果存在党小组的话，把党小组赋值给 partyId
      if (row.partyGroupId) {
        row.partyId = row.partyGroupId;
      }
      this.$refs.operaUserWindow.withFooter = true;
      this.beforeOperateUser(row);
      this.$refs.operaUserWindow.open("编辑用户", row);
    },
    beforeOperateUser(row) {
      this.$refs.operaUserWindow.editFlag = true;
      this.queryUserRoleByUserId(row.id);
      this.queryUserDepartByUserId(row.id);
    },

    // 解冻用户/冻结用户
    frozenBatchUser(row, status) {
      const parmas = {
        ids: row.id,
        status: status,
      };
      const message = `确认是否${status === 2 ? "冻结" : "解冻"}${
        row.realname
      }`;
      this.$dialog.frozenBatchConfirm(message,status).then(() => {
        this.api.frozenBatch(parmas).then((res) => {
          if (200 === res.code) {
            row.status = status;
            this.$message.success(res.message);
              this.search()
          }
        });
      });
    },
    queryUserRoleByUserId(userId) {
      this.api.queryUserRoleByUserId({ userId: userId }).then((res) => {
        this.$refs.operaUserWindow.form.selectedrolesArr = res;
      });
    },
    queryUserDepartByUserId(userId) {
      this.api.queryUserDepartByUserId({ userId: userId }).then((res) => {
        let result = [];
        if(res){
          res.forEach((e)=>{
            this.pushId(e,result)
          })
        }
        let departList=result
        if (departList.length<2){
          this.$refs.operaUserWindow.form.selecteddeparts =departList[0];
          return
        }
        this.$refs.operaUserWindow.form.selecteddeparts=departList;
      });
    },
    pushId(obj,result){
      if(obj.key){
        result.push(obj.key)
        if(obj.children){
          this.pushId(obj.children,result)
        }
      }
    },
  },

  created() {
    this.config({
      module: '用户', // 当前组件名称
      api: '/system/user', // api目录下的js文件
      'field.main': 'realname',
    })
    this.search()
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/style/variables.scss";
// 列表头像处理
.table-column-avatar {
  img {
    width: 48px;
  }
}

.el-dropdown-link {
  cursor: pointer;
  font-size: 12px;
  
}

.search_width {
  width: 300px;
}
</style>
