<template>
  <div class="col-md-6">
    <div class="card-deck">
      <div class="card shadow-sm">
        <div class="card-header">
          <h5 class="font-weight-normal">账户列表</h5>
        </div>
        <div
          id="key-list"
          class="card-body"
        >
          <!-- 重复单元 -->
          <b-list-group>
            <b-list-group-item
              class="d-flex justify-content-between align-items-center"
              v-for="key in keys"
              :key="key.index"
            >
              <p class="address">{{ key.address }}</p>
              <b-dropdown
                size="sm"
                text="操作"
                dropleft
              >
                <b-dropdown-item @click="toSend(key)">交易</b-dropdown-item>
                <b-dropdown-item @click="copyToClipboard(key)">发送给</b-dropdown-item>
                <b-dropdown-item @click="removeKey(key)">删除</b-dropdown-item>
              </b-dropdown>
            </b-list-group-item>
          </b-list-group>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.address {
  margin-bottom: 0;
  max-width: 80%;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  -o-text-overflow: ellipsis;
  -ms-text-overflow: ellipsis;
}
</style>

<script>
import { mapState } from 'vuex'
import { Decimal } from 'decimal.js'

const unit = new Decimal(10).pow(18)

const fetchAccount = async function (host, address) {
  let response
  try {
    response = await fetch(host + 'account/' + address, { Origin: 'http://localhost:8080/' })
    const result = await response.json()
    if (result.Err !== '') {
      return { err: result.Err }
    }

    let balance = new Decimal(result.Data.balance)
    balance = balance.div(unit)
    const nonce = result.Data.nonce
    return { err: null, balance: balance.toFixed(18), nonce }
  } catch (e) {
    return { err: e }
  }
}

export default {
  computed: {
    ...mapState('data', ['keys'])
  },
  methods: {
    async toSend (account) {
      const host = this.$store.state.data.host
      const result = await fetchAccount(host, account.address)
      if (result.err === null) {
        this.$store.commit('event/setAccount', {
          balance: result.balance,
          nonce: result.nonce,
          address: account.address,
          hexkey: account.hexkey
        })
      } else {
        console.log('获取账户信息出错', result.err)
      }
    },
    copyToClipboard (account) {
      this.$emit('selected', account.address)
    },
    removeKey (account) {
      this.$store.dispatch('data/removeKeyAsync', account.address)
    }
  }
}
</script>
