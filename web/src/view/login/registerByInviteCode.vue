<template>
  <div
    id="userLayout"
    class="w-full h-full relative"
  >
    <div
      class="rounded-lg flex items-center justify-evenly w-full h-full bg-white md:w-screen md:h-screen md:bg-[#194bfb]"
    >
      <div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
        <div class="oblique h-[130%] w-3/5 bg-white transform -rotate-12 absolute -ml-52" />
        <!-- 分割斜块 -->
        <div class="z-[999] pt-12 pb-10 md:w-96 w-full  rounded-lg flex flex-col justify-between box-border">
          <div>
            <div class="flex items-center justify-center">

              <!-- <img
                  class="w-24"
                  :src="$GIN_VUE_ADMIN.appLogo"
                  alt
                > -->
            </div>
            <div class="mb-9">
              <p class="text-center text-4xl font-bold">{{ $GIN_VUE_ADMIN.appName }}</p>
              <p class="text-center text-sm font-normal text-gray-500 mt-2.5">
                若无法注册可联系 lich.wu2014@gmail.com
              </p>
            </div>
            <el-form
              ref="loginForm"
              :model="registerFormData"
              :rules="rules"
              :validate-on-rule-change="false"
              @keyup.enter="submitForm"
            >
              <el-form-item
                prop="inviteCode"
                class="mb-6"
              >
                <el-input
                  v-model="registerFormData.inviteCode"
                  size="large"
                  placeholder="请输入邀请码"
                />
              </el-form-item>
              <el-form-item
                prop="email"
                class="mb-6"
              >

                <el-input
                  v-model="registerFormData.email"
                  size="large"
                  placeholder="请输入邮箱"
                  suffix-icon="user"
                />
              </el-form-item>
              <el-form-item
                prop="password"
                class="mb-6"
              >
                <el-input
                  v-model="registerFormData.password"
                  show-password
                  size="large"
                  type="password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item
                prop="passwordConfirm"
                class="mb-6"
              >
                <el-input
                  v-model="registerFormData.passwordConfirm"
                  show-password
                  size="large"
                  type="password"
                  placeholder="请再次输入密码"
                />
              </el-form-item>
              <el-form-item
                v-if="registerFormData.openCaptcha"
                prop="captcha"
                class="mb-6"
              >
                <div class="flex w-full justify-between">
                  <el-input
                    v-model="registerFormData.captcha"
                    placeholder="请输入验证码"
                    size="large"
                    class="flex-1 mr-5"
                  />
                  <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                    <img
                      v-if="picPath"
                      class="w-full h-full"
                      :src="picPath"
                      alt="请输入验证码"
                      @click="loginVerify()"
                    >
                  </div>
                </div>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="submitForm"
                >注 册</el-button>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="toLogin()"
                >返回登陆</el-button>
              </el-form-item>
              <el-form-item
                v-if="showInit"
                class="mb-6"
              >
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="checkInit"
                >前往初始化</el-button>

              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
      <!-- <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]"><img
          class="h-full"
          src="@/assets/login_right_banner.jpg"
          alt="banner"
        ></div> -->
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto  w-full z-20">
      <!-- <div class="links items-center justify-center gap-2 hidden md:flex">
          <a
            href="http://doc.henrongyi.top/"
            target="_blank"
          >
            <img
              src="@/assets/docs.png"
              class="w-8 h-8"
              alt="文档"
            >
          </a>
          <a
            href="https://support.qq.com/product/371961"
            target="_blank"
          >
            <img
              src="@/assets/kefu.png"
              class="w-8 h-8"
              alt="客服"
            >
          </a>
          <a
            href="https://github.com/flipped-aurora/gin-vue-admin"
            target="_blank"
          >
            <img
              src="@/assets/github.png"
              class="w-8 h-8"
              alt="github"
            >
          </a>
          <a
            href="https://space.bilibili.com/322210472"
            target="_blank"
          >
            <img
              src="@/assets/video.png"
              class="w-8 h-8"
              alt="视频站"
            >
          </a>
        </div> -->
    </BottomInfo>
  </div>
</template>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import { reactive, ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import { registerWithInviteCode } from '@/api/user'

defineOptions({
  name: 'Login',
})

const router = useRouter()
const route = useRoute()
// 验证函数
// const checkUsername = (rule, value, callback) => {
//   if (value.length < 5) {
//     return callback(new Error('请输入正确的用户名'))
//   } else {
//     callback()
//   }
// }
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}
const checkPasswordConfirm = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  }
  if (value !== registerFormData.password) {
    return callback(new Error('两次输入的密码不一致'))
  }

  callback()
}
const checkInviteCode = (rule, value, callback) => {
  if (value.length < 8) {
    return callback(new Error('请输入正确的邀请码'))
  } else {
    callback()
  }
}

onMounted(() => {
  if (route.query.inviteCode) {
    registerFormData.inviteCode = route.query.inviteCode
  }

  checkShowInit()
})

// 获取验证码
const loginVerify = () => {
  captcha({}).then(async(ele) => {
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur',
    })
    picPath.value = ele.data.picPath
    registerFormData.captchaId = ele.data.captchaId
    registerFormData.openCaptcha = ele.data.openCaptcha
  })
}
loginVerify()

// 登录相关操作
const loginForm = ref(null)
const picPath = ref('')
const registerFormData = reactive({
  email: '',
  inviteCode: '',
  password: '',
  passwordConfirm: '',
  captcha: '',
  captchaId: '',
  openCaptcha: false,
})
const rules = reactive({
  email: [
    { required: true, pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/, message: '请输入正确的邮箱', trigger: 'blur' },
  ],
  inviteCode: [{ validator: checkInviteCode, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  passwordConfirm: [{ validator: checkPasswordConfirm, trigger: 'blur' }],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})
const showInit = ref(false)

const userStore = useUserStore()
const register = async() => {
  return await registerWithInviteCode(registerFormData)
}
const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (v) {
      const res = await register()
      if (res.code !== 0) {
        loginVerify()
      } else {
        ElMessage({
          type: 'success',
          message: '注册成功',
        })
        toLogin({
          username: registerFormData.email,
        })
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

// 跳转初始化
const checkInit = async() => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: '已配置数据库信息，无法初始化',
      })
    }
  }
}

const toLogin = (query) => {
  router.push({ name: 'Login', query })
}

const checkShowInit = async() => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      showInit.value = true
    }
  }
}

</script>

