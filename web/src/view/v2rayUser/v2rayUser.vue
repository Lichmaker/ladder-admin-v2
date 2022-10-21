<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户id">
          <el-input v-model="searchInfo.adminUserId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="searchInfo.email" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="账号激活状态" prop="activeState">
          <el-select
            v-model="searchInfo.activeState"
            clearable
            placeholder="请选择"
            @clear="
              () => {
                searchInfo.activeState = undefined;
              }
            "
          >
            <el-option
              v-for="(item, key) in v2rayClientStatusOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button size="small" type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="small"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >删除</el-button
            >
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="用户id" prop="adminUserId" width="120" />
        <el-table-column align="left" label="邮箱" prop="email" width="120" />
        <el-table-column
          align="left"
          label="每月最大的流量使用，单位MB"
          prop="bandwidthUsageMax"
          width="120"
        />
        <el-table-column align="left" label="vemss url" prop="v2rayVmess" width="120" />
        <el-table-column align="left" label="uuid字段" prop="uuid" width="120" />
        <el-table-column align="left" label="账户过期时间" prop="expireAt" width="120" />
        <el-table-column
          align="left"
          label="上一次统计时间"
          prop="statUpdatedAt"
          width="120"
        />
        <el-table-column align="left" label="账号激活状态" prop="activeState" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.activeState, v2rayClientStatusOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              size="small"
              class="table-button"
              @click="updateV2rayUserFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              size="small"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form
        :model="formData"
        label-position="right"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="用户id:" prop="adminUserId">
          <el-input
            v-model.number="formData.adminUserId"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="每月最大的流量使用，单位MB:" prop="bandwidthUsageMax">
          <el-input
            v-model.number="formData.bandwidthUsageMax"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="vemss url:" prop="v2rayVmess">
          <el-input
            v-model="formData.v2rayVmess"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="uuid字段:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="账户过期时间:" prop="expireAt">
          <el-date-picker
            v-model="formData.expireAt"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="上一次统计时间:" prop="statUpdatedAt">
          <el-date-picker
            v-model="formData.statUpdatedAt"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="账号激活状态:" prop="activeState">
          <el-select
            v-model="formData.activeState"
            placeholder="请选择"
            style="width: 100%"
            :clearable="false"
          >
            <el-option
              v-for="(item, key) in v2rayClientStatusOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "V2rayUser",
};
</script>

<script setup>
import {
  createV2rayUser,
  deleteV2rayUser,
  deleteV2rayUserByIds,
  updateV2rayUser,
  findV2rayUser,
  getV2rayUserList,
} from "@/api/v2rayUser";

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

// 自动化生成的字典（可能为空）以及字段
const v2rayClientStatusOptions = ref([]);
const formData = ref({
  adminUserId: 0,
  email: "",
  bandwidthUsageMax: 0,
  v2rayVmess: "",
  uuid: "",
  expireAt: new Date(),
  statUpdatedAt: new Date(),
  activeState: undefined,
});

// 验证规则
const rule = reactive({
  adminUserId: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  email: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  expireAt: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  activeState: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
});

const elFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getV2rayUserList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  v2rayClientStatusOptions.value = await getDictFunc("v2rayClientStatus");
};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteV2rayUserFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.adminUserId);
    });
  const res = await deleteV2rayUserByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateV2rayUserFunc = async (row) => {
  const res = await findV2rayUser({ adminUserId: row.adminUserId });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.rev2rayUser;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteV2rayUserFunc = async (row) => {
  const res = await deleteV2rayUser({ adminUserId: row.adminUserId });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    adminUserId: 0,
    email: "",
    bandwidthUsageMax: 0,
    v2rayVmess: "",
    uuid: "",
    expireAt: new Date(),
    statUpdatedAt: new Date(),
    activeState: undefined,
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createV2rayUser(formData.value);
        break;
      case "update":
        res = await updateV2rayUser(formData.value);
        break;
      default:
        res = await createV2rayUser(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style></style>
