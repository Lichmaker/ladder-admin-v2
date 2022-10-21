<template>
  <div>
    <div class="gva-form-box">
      <el-form
        :model="formData"
        ref="elFormRef"
        label-position="right"
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
            placeholder="选择日期"
            :clearable="true"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="上一次统计时间:" prop="statUpdatedAt">
          <el-date-picker
            v-model="formData.statUpdatedAt"
            type="date"
            placeholder="选择日期"
            :clearable="true"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="账号激活状态:" prop="activeState">
          <el-select
            v-model="formData.activeState"
            placeholder="请选择"
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "V2rayUser",
};
</script>

<script setup>
import { createV2rayUser, updateV2rayUser, findV2rayUser } from "@/api/v2rayUser";

// 自动获取字典
import { getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { ref, reactive } from "vue";
const route = useRoute();
const router = useRouter();

const type = ref("");
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

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.adminUserId) {
    const res = await findV2rayUser({ adminUserId: route.query.adminUserId });
    if (res.code === 0) {
      formData.value = res.data.rev2rayUser;
      type.value = "update";
    }
  } else {
    type.value = "create";
  }
  v2rayClientStatusOptions.value = await getDictFunc("v2rayClientStatus");
};

init();
// 保存按钮
const save = async () => {
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
    }
  });
};

// 返回按钮
const back = () => {
  router.go(-1);
};
</script>

<style></style>
