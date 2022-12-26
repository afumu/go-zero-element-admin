<template>
    <TableLayout :permissions="['system:log:query']">
    <el-tabs v-model="activeName" class="table-wrap">
      <el-tab-pane label="登录日志" name="first">
        <loginLog ></loginLog>
      </el-tab-pane>
      <el-tab-pane label="系统日志" name="second">
        <systemLog ></systemLog>
      </el-tab-pane>
    </el-tabs>
  </TableLayout>
</template>

<script>
import BaseTable from "@/components/base/BaseTable";
import TableLayout from "@/layouts/TableLayout";
import loginLog from "@/components/system/log/loginLog";
import systemLog from "@/components/system/log/systemLog";
export default {
  name: "log",
  extends: BaseTable,
  components: { TableLayout, loginLog, systemLog},
  data() {
    return {
      activeName: "first",
    };
  },
  methods: {
  },
};
</script>

<style scoped lang="scss">
.table-wrap{
  margin: 0 16px;
  background: #fff;
}
</style>
