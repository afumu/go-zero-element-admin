<!-- 定时任务列表 -->
<template>
  <table-layout :permissions="['system:quartzJobList:query']">
    <!-- 搜索表单 -->
    <div slot="search-form">
      <el-collapse v-model="activeNames" accordion>
        <el-collapse-item title="查询条件" name="1">
          <template slot="title">
            <div class="kj-title">
              <span style="margin: 0 10px">查询条件</span>
              <section>
                <el-button type="primary" size="mini" icon="el-icon-search" @click.stop="search">搜索</el-button>
                <el-button @click.stop="reset" size="mini" icon="el-icon-refresh">重置</el-button>
              </section>
            </div>
          </template>
          <el-form ref="searchForm" :model="searchForm" label-width="100px" inline>
            <el-form-item label="任务类名" prop="jobClassName">
              <search-input v-model="searchForm.jobClassName" v-trim placeholder="请输入任务类名" @keypress.enter.native="search" />
            </el-form-item>
            <el-form-item label="任务状态" prop="status">
              <el-select clearable v-model="searchForm.status" placeholder="请选择任务状态">
                <el-option value="" label="全部"></el-option>
                <el-option value="0" label="正常"></el-option>
                <el-option value="-1" label="停止"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
      <!-- <div style="height: 10px"></div> -->
    </div>
    <template v-slot:table-wrap>
      <ul class="toolbar">
        <li>
          <el-button type="primary" @click="$refs.operaQuartzJobWindow.open('新建')" icon="el-icon-plus">新建</el-button>
        </li>
        <li>
          <el-button @click="deleteByIdInBatch" icon="el-icon-delete">删除</el-button>
        </li>
      </ul>
      <el-table  border v-loading="isWorking.search" :data="tableData.list" :default-sort="{ prop: 'createTime', order: 'descending' }" stripe @sort-change="handleSortChange" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="50px" fixed="left"></el-table-column>
        <!-- <el-table-column label="#" fixed="left" min-width="30px" align="center">
          <template slot-scope="scope">
            <label>{{ scope.$index + 1 }}</label>
          </template>
        </el-table-column> -->
        <el-table-column prop="jobClassName" label="任务类名" fixed="left" min-width="170px"></el-table-column>
        <el-table-column prop="cronExpression" label="cron表达式" fixed="left" min-width="60px"></el-table-column>
        <el-table-column prop="parameter" label="参数" fixed="left" min-width="60px"></el-table-column>
        <el-table-column prop="description" label="描述" fixed="left" min-width="120px"></el-table-column>
        <el-table-column prop="status" label="状态" fixed="left" min-width="50px">
          <template slot-scope="{ row }">
              <el-tag v-if="row.status === -1" type="danger">已暂停</el-tag>
              <el-tag v-else-if="row.status === 0" type="success">已启动</el-tag>
            <!-- <label v-if="row.status === -1" style="color: #fa8c16;" class="status">已暂停</label>
            <label v-else-if="row.status === 0" style="color: green" class="status">已启动</label> -->
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="80" fixed="right" align="center">
          <template slot-scope="{ row }">
            <el-button type="text" v-if="row.status === -1" @click="resume(row)" icon="el-icon-video-play">启动</el-button>
            <el-button type="text" v-else-if="row.status === 0" @click="pause(row)" icon="el-icon-video-pause">暂停</el-button>

            <el-dropdown>
              <span class="el-dropdown-link">
                <el-button type="text" style="padding-right: 0px !important">更多</el-button>
                <i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>
                  <el-button type="text" @click="execute(row)">立即执行</el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button type="text" @click="$refs.operaQuartzJobWindow.open('编辑', row)" icon="el-icon-setting">编辑</el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button type="text" @click="deleteById(row)" icon="el-icon-delete">删除</el-button>
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </el-table>
      <pagination @size-change="handleSizeChange" @current-change="handlePageChange" :pagination="tableData.pagination"></pagination>
    </template>
    <!-- 新建/修改 -->
    <OperaQuartzJobWindow ref="operaQuartzJobWindow" @success="handlePageChange(tableData.pagination.pageIndex)" />
  </table-layout>
</template>

<script>
import BaseTable from "@/components/base/BaseTable";
import Pagination from "@/components/common/Pagination";
import TableLayout from "@/layouts/TableLayout";
import OperaQuartzJobWindow from "@/components/system/quartzjob/OperaQuartzJobWindow";
export default {
  name: "QuzrtzJobList",
  extends: BaseTable,
  components: {
    TableLayout,
    Pagination,
    OperaQuartzJobWindow,
  },
  data() {
    return {
      activeNames: "1",
      searchForm: {
        status: null,
        jobClassName: null,
      },
    };
  },
  methods: {
    // 启用
    resume(row) {
      console.log("是否启动：" + row.jobClassName);
      let message = `是否启动选中的任务?`;
      this.$confirm(message, "确认启动", {
        distinguishCancelAndClose: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(() => {
          this.api
            .resume(row.jobClassName)
            .then(() => {
              this.$tip.apiSuccess("启动成功");
              this.search();
            })
            .catch((e) => {
              this.$tip.apiFailed(e);
            });
        })
        .catch(() => {
          this.$tip.apiSuccess("已取消");
        });
    },

    // 暂停
    pause(row) {
      console.log("是否暂停：" + row.jobClassName);
      let message = `是否停止选中的任务?`;
      this.$confirm(message, "确认暂停", {
        distinguishCancelAndClose: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(() => {
          this.api
            .pause(row.jobClassName)
            .then(() => {
              this.$tip.apiSuccess("暂停成功");
              this.search();
            })
            .catch((e) => {
              this.$tip.apiFailed(e);
            });
        })
        .catch(() => {
          this.$tip.apiSuccess("已取消");
        });
    },

    execute(row) {
      console.log("立即执行：" + row.jobClassName);
      let message = `是否立即执行选中的任务?`;
      this.$confirm(message, "立即执行", {
        distinguishCancelAndClose: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
      })
        .then(() => {
          this.api
            .execute(row.id)
            .then(() => {
              this.$tip.apiSuccess("成功执行");
            })
            .catch((e) => {
              this.$tip.apiFailed(e);
            });
        })
        .catch(() => {});
    },
  },
  //生命周期 - 创建完成（访问当前this实例）
  created() {
    this.config({
      module: "定时任务",
      api: "/system/quartzJob",
      sorts: [
        {
          property: "quartzJob.CREATE_TIME",
          direction: "DESC",
        },
      ],
      "field.main": "jobClassName",
    });
    this.search();
  },
};
</script>
<style scoped lang="scss">
/* @import url(); 引入css类 */
.status {
  border-radius: 4px;
  padding: 0 7px;
  margin: 0;
  font-size: 12px;
  display: inline-block;
  height: auto;
  line-height: 20px;
  color: #fa8c16;
  background: #fff7e6;
  border-color: #ffd591;
}

.el-dropdown-link {
  cursor: pointer;
  font-size: 12px;
  margin-left: 10px;
}
.el-icon-arrow-down {
  font-size: 12px;
}
</style>
