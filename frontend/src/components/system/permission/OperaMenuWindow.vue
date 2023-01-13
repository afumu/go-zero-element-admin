<template>
  <GlobalWindow class="handle-table-dialog" :title="title" :visible.sync="visible" :confirm-working="isWorking" @confirm="confirm">
    <p class="tip" v-if="form.parent != null && form.id == null">
      为 <em>{{ parentName }}</em> 新建子菜单
    </p>
    <el-form :model="form" ref="form" :rules="rules">
      <el-form-item>
        <el-radio-group v-model="form.menuType">
          <el-radio :label="0">一级菜单</el-radio>
          <el-radio :label="1">子菜单</el-radio>
          <el-radio :label="2">按钮/权限</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item v-if="form.menuType != 0" label="上级菜单" prop="parentId">
        <MenuSelect v-if="visible" v-model="form.parentId" placeholder="请选择上级菜单" :exclude-id="excludeMenuId" clearable :inline="false" />
      </el-form-item>

      <el-form-item :label="form.menuType != 2 ? '菜单名称' : '按钮/权限'" prop="name" required>
        <el-input v-model="form.name" :placeholder="
            form.menuType != 2 ? '请输入菜单名称' : '请输入按钮/权限'
          " v-trim maxlength="50" />
      </el-form-item>

      <el-form-item v-if="form.menuType == 2" label="授权标识" prop="perms">
        <el-input v-model="form.perms" placeholder="请输入授权标识" v-trim maxlength="200" />
      </el-form-item>

      <el-form-item v-if="form.menuType == 2" label="授权策略">
        <el-radio-group v-model="form.permsType">
          <el-radio style="margin-top: 10px;" label="1">可见/可访问(授权后可见/可访问)</el-radio><br />
          <el-radio style="margin-top: 10px;" label="2">可编辑(未授权时禁用)</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item v-if="form.menuType == 2" label="状态">
        <el-radio-group v-model="form.status">
          <el-radio label="0">无效</el-radio>
          <el-radio label="1">有效</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item v-if="form.menuType != 2" label="访问路径" prop="url">
        <el-input v-model="form.url" placeholder="请输入访问路径" v-trim maxlength="200" />
      </el-form-item>

      <el-form-item v-if="form.menuType != 2" label="排序" prop="url">
        <el-input-number style="width: 150px;" v-model="form.sort" :min="1" :max="100"></el-input-number>
      </el-form-item>

      <el-form-item v-if="form.menuType != 2" label="图标" prop="icon" class="form-item-icon">
        <el-radio-group v-model="form.icon">
          <el-radio :label="icon" v-for="icon in icons" :key="icon">
            <i :class="{ [icon]: true }"></i>
          </el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="备注" prop="description">
        <el-input type="textarea" v-model="form.description" placeholder="请输入菜单备注" v-trim :rows="3" maxlength="500" />
      </el-form-item>
    </el-form>
  </GlobalWindow>
</template>

<script>
import BaseOpera from "@/components/base/BaseOpera";
import GlobalWindow from "@/components/common/GlobalWindow";
import MenuSelect from "@/components/common/MenuSelect";
import icons from "@/utils/icons";
export default {
  name: "OperaMenuWindow",
  extends: BaseOpera,
  components: { MenuSelect, GlobalWindow },
  data() {
    return {
      icons,
      // 上级菜单名称
      parentName: "",
      // 需排除选择的菜单ID
      excludeMenuId: null,
      // 表单数据
      form: {
        id: undefined,
        parentId: undefined,
        url: "",
        name: "",
        icon: "",
        description: "",
        sort: 1,
        // route: true,
        // keepAlive: false,
        // internalOrExternal: false,
        // hidden: false,
        // alwaysShow: false,
      },
      // 验证规则
      rules: {
        name: [{ required: true, message: "请输入菜单名称" }],
        parentId: [{ required: true, message: "请选择" }],
        url: [{ required: true, message: "请输入" }],
      },
    };
  },
  watch: {
    // 'form.parentId'(val) {
    //   if (val) {
    //     this.form.menuType = 1
    //   } else {
    //     this.form.menuType = 0
    //   }
    // },
  },
  methods: {
    result() {
      this.form = {
        id: undefined,
        parentId: undefined,
        url: "",
        name: "",
        icon: "",
        description: "",
        sort: 1,
        // route: true,
        // keepAlive: false,
        // internalOrExternal: false,
        // hidden: false,
        // alwaysShow: false,
      };
    },
    /**
     * @title: 窗口标题
     * @target: 编辑的菜单对象
     * @parent: 新建时的上级菜单
     */
    open(title, target, parent) {
      this.result();
      this.title = title;
      this.visible = true;
      if (parent) {
        if (parent.menuType == 1) {
          this.form.menuType = 2;
        } else {
          this.form.menuType = 1;
        }
      }
      // 新建，menu为空时表示新建菜单
      if (target == null) {
        this.excludeMenuId = null;
        this.$nextTick(() => {
          this.$refs.form.resetFields();
          this.form.id = undefined;
          this.form.parentId = parent == undefined ? undefined : parent.id;
          this.parentName = parent == undefined ? undefined : parent.name;
        });
        return;
      } else {
        this.form = Object.assign(this.form, parent);
      }
      // 编辑
      this.$nextTick(() => {
        this.excludeMenuId = target.id;
        for (const key in this.form) {
          this.form[key] = target[key];
        }
      });
    },
  },
  created() {
    this.config({
      api: "/system/menu",
    });
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/style/variables";
$icon-background-color: $primary-color;
.global-window {
  .tip {
    margin-bottom: 12px;
    em {
      font-style: normal;
      color: $primary-color;
      font-weight: bold;
    }
  }
  // 图标
  /deep/ .form-item-icon {
    .el-form-item__content {
      height: 193px;
      overflow-y: auto;
    }
    .el-radio-group {
      background: $icon-background-color;
      padding: 4px;
      .el-radio {
        margin-right: 0;
        color: #fff;
        padding: 6px;
        &.is-checked {
          background: $icon-background-color - 30;
          border-radius: 10px;
        }
        .el-radio__input.is-checked + .el-radio__label {
          color: #fff;
        }
      }
      .el-radio__input {
        display: none;
      }
      .el-radio__label {
        padding-left: 0;
        // element-ui图标
        i {
          font-size: 30px;
        }
        // 自定义图标
        [class^="eva-icon-"],
        [class*=" eva-icon-"] {
          width: 30px;
          height: 30px;
          background-size: 25px;
          vertical-align: bottom;
        }
      }
      .el-radio--small {
        height: auto;
        padding: 8px;
        margin-left: 0;
      }
    }
  }
}
</style>
