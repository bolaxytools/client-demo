<template>
  <div class="card-deck mt-4 mb-3">
    <div class="card shadow-sm">
      <div class="card-header">
        {{ $t('transferTitle') }}
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-5">
            <div class="form-group">
              <label>{{$t('inputFrom')}}</label>
              <b-input
                readonly
                :value="from"
              />
            </div>
            <div class="form-group">
              <label>{{ $t('inputBalance') }}</label>
              <b-input
                readonly
                :value="balance"
              />
              <small class="form-text text-muted">tips: 1=10**18</small>
            </div>
            <b-button @click="transfer" :disabled="disabled">发送</b-button>
          </div>
          <div class="col-6">
            <div class="form-group">
              <label>{{$t('inputTo')}}</label>
              <b-input v-model="target" />
            </div>
            <div class="form-group">
              <label>{{$t('inputValue')}}</label>
              <b-input
                v-model="value"
                aria-describedby="amountHelp"
              />
              <small
                id="amountHelp"
                class="form-text text-muted"
              >tips:
                发送数额不能超过余额.1=10**18</small>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { Decimal } from 'decimal.js'
const unit = new Decimal(10).pow(18)
export default {
  data: function () {
    return {
      value: 0.0,
      disabled: false
    }
  },
  props: {
    target: {
      type: String,
      default: ''
    }
  },
  methods: {
    async transfer () {
      this.disabled = true
      const signTx = window.signTx
      let amount = new Decimal(this.value)
      amount = amount.mul(unit)
      // console.log('tx-> to:', this.target, 'value:', this.value,
      // 'amount:', amount.toString(), 'from:', this.from, 'balance:', this.balance)
      try {
        const signedTx = await new Promise((resolve, reject) => {
          signTx({
            to: this.target,
            value: amount.toString(),
            hexkey: this.hexkey,
            nonce: this.nonce
          }, (err, data) => err ? reject(err) : resolve(data))
        })

        const host = this.$store.state.data.host
        const response = await fetch(host + 'rawtx', {
          method: 'POST',
          body: signedTx
        })
        // {"Data":{"txHash":"0x7a900c8b9e545319e1652033f4e9f946e4dfdd2117d7c8039ce9f4f989562836"},"Err":""}
        const result = await response.json()

        const txhash = result.Data.txHash
        this.$parent.makeToast({ title: '交易发送成功', body: `交易hash: ${txhash}`, variant: 'info' })
        this.disabled = false

        // console.log('after transfer ...', txhash);
        let counter = 0
        const inter = setInterval(() => {
          fetch(host + 'tx/' + txhash)
            .then((resp) => resp.json())
            .then(async (ret) => {
              if (ret.Data !== null) {
                clearInterval(inter)
                // console.log('query txhash -> ', result.Data);
                const status = ret.Data.status
                const tx = { txhash: txhash, to: this.target, value: this.value, status: status }
                await this.$store.dispatch('data/addTxAsync', tx)
                this.$parent.makeToast({ title: '交易已被确认', body: `交易hash: ${txhash}`, variant: 'success' })
                return
              }

              counter += 1
              if (counter > 17) {
                clearInterval(inter)
                this.$parent.makeToast({ title: '交易确认超时', body: `交易hash: ${txhash}`, variant: 'warning' })
              }
            })
            .catch(e => {
              clearInterval(inter)
              this.$parent.makeToast({ title: '查询超时', body: `fetch txhash info failed: ${e}`, variant: 'danger' })
            })
        }, 2000)
        // {"Data":null,"Err":"Key not found"}
      } catch (e) {
        this.$parent.makeToast({ title: '交易失败', body: `transfer failed: ${e}`, variant: 'danger' })
      }
    }
  },
  computed: {
    from () {
      const account = this.$store.state.event.account
      if (Object.prototype.hasOwnProperty.call(account, 'address')) {
        return account.address
      }
      return ''
    },
    balance () {
      const account = this.$store.state.event.account
      if (Object.prototype.hasOwnProperty.call(account, 'balance')) {
        return account.balance
      }
      return ''
    },
    nonce () {
      const account = this.$store.state.event.account
      if (Object.prototype.hasOwnProperty.call(account, 'nonce')) {
        return account.nonce
      }
      return 0
    },
    hexkey () {
      const account = this.$store.state.event.account
      if (Object.prototype.hasOwnProperty.call(account, 'hexkey')) {
        return account.hexkey
      }
      return ''
    }
  }
}
</script>
