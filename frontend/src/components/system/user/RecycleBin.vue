<template>
  <el-dialog title="回收站" :visible.sync="visible" width="70%">
    <el-table border v-loading="isWorking.search" :data="deleteList" :default-sort="{prop: 'createTime'}" stripe>
      <el-table-column label="#" fixed="left" min-width="50px" align="center">
        <template slot-scope="scope">
          <label>{{scope.$index+1}}</label>
        </template>
      </el-table-column>
      <el-table-column prop="username" label="账号" fixed="left" min-width="100px"></el-table-column>
      <el-table-column prop="realname" label="姓名" fixed="left" min-width="100px"></el-table-column>
      <el-table-column prop="orgCode" label="部门" min-width="150px"></el-table-column>
      <el-table-column label="操作" min-width="100" fixed="right" align="center">
        <template slot-scope="{row}">
          <el-button type="text" @click="deleteBack(row.id)" icon="el-icon-refresh-right">还原用户</el-button>
          <el-button type="text" @click="deletePhysic(row.id)" icon="el-icon-delete">彻底删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import { deleteRecycleBin, putRecycleBin } from "../../../api/system/user";
export default {
  name: "DictDeleteList",
  extends: BaseOpera,
  data() {
    return {
      visible: false,
      deleteList: [],
    };
  },
  methods: {
    // 查询回收站的列表
    search() {
      this.api
        .recycleBin()
        .then((data) => {
          this.deleteList = data;
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        });
    },
    // 回复
    deleteBack(id) {
      this.api
        .putRecycleBin({ userIds: id })
        .then(() => {
          this.$tip.apiSuccess("成功恢复");
          this.search()
          this.$emit('success')
        })
        .catch((e) => {
          this.$tip.apiFailed(e);
        });
    },
    // 彻底删除
    deletePhysic(id) {
      let message =
        "您确定要彻底删除这 1 个用户吗？注意：彻底删除后将无法恢复，请谨慎操作！";
      this.$dialog.deleteConfirm(message).then(() => {
        this.api
          .deleteRecycleBin(id)
          .then(() => {
            this.$tip.apiSuccess("成功删除");
            this.search();
          })
          .catch((e) => {
            this.$tip.apiFailed(e);
          });
      });
    },
    open() {
      this.visible = true;
      this.search();
    },
  },
  //生命周期 - 创建完成（访问当前this实例）
  created() {
    this.config({
      api: "/system/user",
    });
  },
};
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
</style>
