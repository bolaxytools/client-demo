<template>
  <div class="col-md-6">
    <div class="card-deck">
      <div class="card shadow-sm">
        <div class="card-header">
          <div class="d-flex justify-content-between align-items-baseline w-100">
            <h5 class="font-weight-normal">导入Keystore</h5>
            <b-button
              variant="success"
              :disabled="isDisabled"
              @click="newKey"
            >
              新建
            </b-button>
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
            variant="primary"
            :disabled="isDisabled"
            @click="importKey"
          >导入
          </b-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  data: function () {
    return {
      encryptedJson: '',
      password: ''
    }
  },
  computed: {
    ...mapGetters('data', ['isDisabled'])
  },
  methods: {
    newKey () {
      const generateKey = window.generateKey
      new Promise((resolve, reject) => {
        generateKey((err, data) => {
          if (err) {
            reject(err)
            return
          }

          resolve(data)
        })
      }).then(async (data) => {
        await this.$store.dispatch('data/addKeyAsync', data)
      }).catch(e => {
        console.log(e)
      })
    },
    importKey () {
      const importKeyStore = window.importKeyStore
      new Promise((resolve, reject) => {
        // importKeyStore 由wasm提供注册的全局函数
        importKeyStore(this.encryptedJson, this.password, (err, data) => {
          if (err) {
            reject(err)
            return
          }

          resolve(data)
        })
      }).then(async (data) => {
        await this.$store.dispatch('data/addKeyAsync', data)
      }).catch(e => {
        console.log(e)
      })
    }
  }
}
</script>
