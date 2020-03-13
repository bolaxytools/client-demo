<template>
  <div class="row">
    <Keystore/>
    <KeyItems @selected="selectTo"/>
  </div>
</template>

<script>
import Keystore from './Keystore'
import KeyItems from './KeyItems'
export default {
  components: {
    Keystore,
    KeyItems
  },
  methods: {
    selectTo (e) {
      this.$emit('notifyParent', e)
    }
  }
}
</script>
