<template>
  <div class="card-deck mt-4 mb-5">
    <div class="card shadow-sm">
      <div class="card-header">
        <h5 class="">最近20条交易记录</h5>
      </div>
      <div
        id="txlist"
        class="card-body"
      >
        <!-- 重复单元 -->
        <div
          class="media-body pb-2 mb-1 small lh-125 border-bottom border-gray"
          v-for="tx in txs"
          :key="tx.index"
        >
          <div class="d-flex justify-content-between align-items-center w-100">
            <strong class="text-gray-dark">TXHASH: {{ tx.txhash }}</strong>
            <span class="badge" :class="[tx.status !== 1 ? failed : success]">{{ tx.status !== 1 ? 'Failed' : 'Success'}}</span>
          </div>
          <div class="d-flex justify-content-between align-items-center w-100">
            <span class="">TO: {{ tx.to }}</span>
            <span>{{ tx.value }} BUSD</span>
          </div>
        </div>
        <!-- 重复单元 -->
      </div>
    </div>
  </div>
</template>
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
