<template>
  <div class="card-deck mb-4">
    <div class="card shadow-sm">
      <div class="card-header d-flex justify-content-between align-items-center">
        {{ $t('hostTitle') }}
        <b-dropdown
          size="sm"
          text="Language"
          dropleft
        >
          <b-dropdown-item-button @click="language(1)">中文</b-dropdown-item-button>
          <b-dropdown-item-button @click="language(2)">English</b-dropdown-item-button>
        </b-dropdown>
      </div>
      <div class="card-body">
        <div class="form-group">
          <b-form-input
            :value="host"
            @change="notify"
            :placeholder="placeholder"
          ></b-form-input>
          <small class="form-text text-muted pt-1">
            <span class="badge badge-warning mr-2">{{ $t('hostTip')}}</span>{{$t('hostNotice')}}
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
    },
    language (flag) {
      flag === 2 ? this.$i18n.locale = 'en' : this.$i18n.locale = 'zh-cn'
    }
  }
}
</script>
