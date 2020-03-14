<template>
  <div class="col-md-6">
    <div class="card-deck">
      <div class="card shadow-sm">
        <div class="card-header">
          <div class="d-flex justify-content-between align-items-baseline w-100">
            导入Keystore
          </div>

        </div>
        <div class="card-body">
          <!-- <div class="form-group">
          <label for="keystore-name">密钥简称</label>
          <input class="form-control"/>
        </div> -->
          <div class="form-group">
            <label for="keystore-liternal">Keystore 文本</label>
            <textarea
              id="keystore-liternal"
              name="keystore-literal"
              class="form-control"
              v-model="encryptedJson"
            ></textarea>
          </div>
          <div class="form-group">
            <label for="keystore-password">Keystore 解密密码</label>
            <input
              id="keystore-password"
              name="keystore-password"
              type="password"
              class="form-control"
              v-model="password"
            />
          </div>
          <b-button
            v-b-tooltip.hover
            title="从加密的Keystore格式字符串导入"
            variant="primary"
            :disabled="isDisabled"
            @click="importKey"
            v-html="impBtnHtml"
          >
          </b-button>
          <b-button
            v-b-tooltip.hover
            title="点击即可创建新密钥"
            variant="success"
            :disabled="isDisabled"
            class="ml-3"
            @click="newKey"
            v-html="newBtnHtml"
          >
            新建
          </b-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
const impBtnBefore = '导入'
const newBtnBefore = '新建'
const btnProcess = '<span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>导入中...'
export default {
  data: function () {
    return {
      encryptedJson: '',
      password: '',
      impBtnHtml: impBtnBefore,
      newBtnHtml: newBtnBefore
    }
  },
  computed: {
    ...mapGetters('data', ['isDisabled'])
  },
  methods: {
    newKey () {
      this.newBtnHtml = btnProcess
      const generateKey = window.generateKey
      new Promise((resolve, reject) => {
        generateKey((err, data) => err ? reject(err) : resolve(data))
      }).then(async (data) => {
        await this.$store.dispatch('data/addKeyAsync', data)
        this.$parent.$parent.makeToast({ title: '私钥新建成功', body: `账户地址：${data.address}`, variant: 'success' })
      }).catch(e => {
        this.$parent.$parent.makeToast({ title: '新建私钥出错', body: e, variant: 'warning' })
      }).finally(() => {
        this.newBtnHtml = newBtnBefore
      })
    },
    importKey () {
      this.impBtnHtml = btnProcess
      // importKeyStore 由wasm提供注册的全局函数
      const importKeyStore = window.importKeyStore
      new Promise((resolve, reject) => {
        importKeyStore(this.encryptedJson, this.password,
          (err, data) => err ? reject(err) : resolve(data))
      }).then(async (data) => {
        await this.$store.dispatch('data/addKeyAsync', data)
      }).catch(e => {
        this.$parent.$parent.makeToast({ title: '导入Keystore出错', body: e, variant: 'warning' })
      }).finally(() => {
        this.impBtnHtml = impBtnBefore
      })
    }
  }
}
</script>
