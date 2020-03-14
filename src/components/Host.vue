<template>
  <div class="card-deck mb-4">
    <div class="card shadow-sm">
      <div class="card-header">
        节点服务地址
      </div>
      <div class="card-body">
        <div class="form-group">
          <b-form-input
            :value="host"
            @change="notify"
            :placeholder="placeholder"
          ></b-form-input>
          <small class="form-text text-muted pt-1">
            <span class="badge badge-warning mr-2">注意</span>使用前请修改为你的节点服务地址
          </small>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  props: {
    host: String,
    placeholder: {
      type: String,
      default: 'http://localhost:8080/'
    }
  },
  methods: {
    notify (value) {
      this.$store.dispatch('data/setHostAsync', value)
    }
  }
}
</script>
