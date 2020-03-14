<template>
  <div class="card-deck mt-4 mb-5">
    <div class="card shadow-sm">
      <div class="card-header">
        最近20条交易记录
      </div>
      <div
        id="txlist"
        class="card-body"
      >
        <!-- 重复单元 -->
        <div
          class="media-body mt-1 small lh-125 border-bottom border-gray"
          v-for="tx in txs"
          :key="tx.index"
        >
          <div class="container-fluid">
            <div class="row row-cols-2">
              <div class="col-9">
                <strong class="d-inline">TXHASH:</strong>
                <p class="text-gray-dark address"> {{ tx.txhash }}</p>
              </div>
              <div class="col-3 text-right">
                <p
                  class="badge"
                  :class="[tx.status !== 1 ? failed : success]"
                >{{ tx.status !== 1 ? 'Failed' : 'Success'}}</p>
              </div>
              <div class="col-9">
                <p class="address">TO: {{ tx.to }}</p>
              </div>
              <div class="col-3 text-right">
                <p>{{ tx.value }} BUSD</p>
              </div>
            </div>

          </div>

        </div>
        <!-- 重复单元 -->
      </div>
    </div>
  </div>
</template>
<style scoped>
@import '../assets/styles/global.css';
</style>
<script>
import { mapState } from 'vuex'
export default {
  data: function () {
    return {
      success: 'badge-success',
      failed: 'badge-secondary'
    }
  },
  computed: {
    ...mapState('data', ['txs'])
  }
}
</script>
