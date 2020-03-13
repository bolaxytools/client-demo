<template>
  <div class="card-deck mt-4 mb-3">
    <div class="card shadow-sm">
      <div class="card-header">
        <h5 class="">发送BUSD</h5>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-5">
            <div class="form-group">
              <label>发送账户</label>
              <b-input
                readonly
                :value="from"
              />
            </div>
            <div class="form-group">
              <label>BUSD余额</label>
              <b-input
                readonly
                :value="balance"
              />
              <small class="form-text text-muted">tips: 1=10**18</small>
            </div>
            <b-button @click="transfer">发送</b-button>
          </div>
          <div class="col-6">
            <div class="form-group">
              <label>发送给</label>
              <b-input v-model="target" />
            </div>
            <div class="form-group">
              <label>发送数额</label>
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
      value: 0.0
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
      const signTx = window.signTx
      let amount = new Decimal(this.value)
      amount = amount.mul(unit)
      console.log('tx-> to:', this.target, 'value:', this.value,
        'amount:', amount.toString(), 'from:', this.from, 'balance:', this.balance)
      try {
        const signedTx = await new Promise((resolve, reject) => {
          signTx({
            to: this.target,
            value: amount.toString(),
            hexkey: this.hexkey,
            nonce: this.nonce
          }, (err, data) => {
            if (err) {
              reject(err)
              return
            }
            resolve(data)
          })
        })

        const host = this.$store.state.data.host
        const response = await fetch(host + 'rawtx', {
          method: 'POST',
          body: signedTx
        })
        // {"Data":{"txHash":"0x7a900c8b9e545319e1652033f4e9f946e4dfdd2117d7c8039ce9f4f989562836"},"Err":""}
        const result = await response.json()

        const txhash = result.Data.txHash

        // console.log('after transfer ...', txhash);
        let counter = 0
        const inter = setInterval(() => {
          fetch(host + 'tx/' + txhash)
            .then((resp) => resp.json())
            .then((ret) => {
              if (ret.Data !== null) {
                clearInterval(inter)
                // console.log('query txhash -> ', result.Data);
                const status = ret.Data.status
                const tx = { txhash: txhash, to: this.target, value: this.value, status: status }
                this.$store.dispatch('data/addTxAsync', tx)
                return
              }

              counter += 1
              if (counter > 17) {
                clearInterval(inter)
              }
            })
            .catch(e => {
              clearInterval(inter)
              console.log('fetch txhash info failed.', e)
            })
        }, 2000)
        // {"Data":null,"Err":"Key not found"}
      } catch (e) {
        console.log('transfer failed.', e)
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
