<template>
  <div class="col-md-6">
    <b-card
      no-body
      :header="$t('accountTitle')"
      class="shadow-sm"
    >
      <b-list-group flush>
        <b-list-group-item
          class="d-flex justify-content-between align-items-center"
          v-for="key in keys"
          :key="key.index"
          href="javascript:void(0)"
        >
          <p
            class="address"
            v-b-tooltip.hover
            :title="key.address"
          >{{ key.address }}</p>
          <b-dropdown
            size="sm"
            :text="$t('accountOpt')"
            dropleft
          >
            <b-dropdown-item @click="toSend(key)">{{ $t('optSend')}}</b-dropdown-item>
            <b-dropdown-item @click="copyToClipboard(key)">{{ $t('optSendTo')}}</b-dropdown-item>
            <b-dropdown-item @click="removeKey(key)">{{ $t('optRemove')}}</b-dropdown-item>
          </b-dropdown>
        </b-list-group-item>
      </b-list-group>
    </b-card>
  </div>
</template>

<style scoped>
@import '../assets/styles/global.css';
</style>

<script>
import { mapState } from 'vuex'
import { Decimal } from 'decimal.js'

const unit = new Decimal(10).pow(18)

const fetchAccount = async function (host, address) {
  const FETCH_TIMEOUT = 5000
  try {
    const response = await new Promise((resolve, reject) => {
      const timeout = setTimeout(() => {
        // eslint-disable-next-line prefer-promise-reject-errors
        reject('Request timed out')
      }, FETCH_TIMEOUT)

      fetch(host + 'account/' + address)
        .then((response) => {
          clearTimeout(timeout)
          resolve(response)
        })
        // eslint-disable-next-line prefer-promise-reject-errors
        .catch((err) => { reject(err) })
    })

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
        this.$parent.$parent.makeToast({ title: '获取账户信息出错', body: `${result.err} 请检查你的节点服务地址是否正确`, variant: 'danger' })
      }
    },
    copyToClipboard (account) {
      this.$emit('selected', account.address)
    },
    async removeKey (account) {
      await this.$store.dispatch('data/removeKeyAsync', account.address)
      this.$parent.$parent.makeToast({ title: '账户已删除', body: account.address, variant: 'info' })
    }
  }
}
</script>
