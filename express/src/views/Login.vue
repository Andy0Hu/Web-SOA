<template>
  <div class="container-login">
    <a-layout class="container-login">
      <a-layout-content class="main">
        <div class="login-title">Express</div>
        <a-form class="user-layout-login" @submit="handleSubmit" :form="form">
          <a-form-item>
            <a-input
              size="large"
              type="text"
              placeholder="账户: admin"
              v-decorator="[
                'username',
                {rules: [{ required: true, message: '请输入帐户名' }], validateTrigger: 'change'}
              ]"
            >
              <a-icon slot="prefix" type="user" :style="{ color: 'rgba(0,0,0,.25)' }" />
            </a-input>
          </a-form-item>
          <a-form-item>
            <a-input
              size="large"
              type="password"
              autocomplete="false"
              placeholder="密码: admin or ant.design"
              v-decorator="[
                'password',
                {rules: [{ required: true, message: '请输入密码' }], validateTrigger: 'blur'}
              ]"
            >
              <a-icon slot="prefix" type="lock" :style="{ color: 'rgba(0,0,0,.25)' }" />
            </a-input>
          </a-form-item>

          <a-form-item style="margin-top:24px">
            <a-button size="large" type="primary" htmlType="submit" class="login-button">确定</a-button>
          </a-form-item>
        </a-form>
        <router-link class="register" :to="{ name: 'register' }">注册账户</router-link>
      </a-layout-content>
      <a-layout-footer>Web SOA Express</a-layout-footer>
    </a-layout>
  </div>
</template>

<script>
export default {
  name: "login",
  date() {
    return {
      form: this.$form.createForm(this, {name: "login"})
    };
  },
  methods: {
    handleSubmit() {}
  }
};
</script>

<style lang="less">
.container-login {
  height: 100%;
  width: 100%;
  .main {
    min-width: 260px;
    width: 368px;
    margin: auto;
    margin-top: 15%;
    .login-title {
      text-align: center;
      margin: auto;
      font-size: 30px;
      margin-bottom: 10px;
    }
  }
  .user-layout-login {
    label {
      font-size: 14px;
    }
    button.login-button {
      padding: 0 15px;
      font-size: 16px;
      height: 40px;
      width: 100%;
    }
    .user-login-other {
      text-align: left;
      margin-top: 24px;
      line-height: 22px;
      .item-icon {
        font-size: 24px;
        color: rgba(0, 0, 0, 0.2);
        margin-left: 16px;
        vertical-align: middle;
        cursor: pointer;
        transition: color 0.3s;
        &:hover {
          color: #1890ff;
        }
      }
      .register {
        float: right;
      }
    }
  }
}
</style>
