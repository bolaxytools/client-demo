<template>
  <div
    id="app"
    class="container-fluid"
  >

    <Head />
    <Host :host="host" />
    <Key @notifyParent="toTx"/>
    <Transaction :target="sendTo"/>
    <LatestRecords />

    <div class="app-mask" :style="{display: progress.display}">
      <div class="container h-100">
        <div class="row h-100 justify-content-center align-items-center">
          <div class="card w-50">
            <div class="card-header">下载WASM</div>
            <div class="card-body">
              <p class="card-text">{{ progress.label }}</p>
              <b-progress :max="progress.max" animated>
                <b-progress-bar
                  :value="progress.value"
                  :label="`${((progress.value / progress.max) * 100).toFixed(1)}%`"
                ></b-progress-bar>
              </b-progress>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="app-mask--backdrop" :style="{display: progress.display}"></div>
  </div>
</template>

<style scoped>
.app-mask--backdrop {
  z-index: 1040;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: black;
  opacity: 0.5;
  overflow: hidden;
}

.app-mask {
  z-index: 1050;
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  width: 100%;
  height: 100%;
  overflow-x: hidden;
  overflow-y: auto;
}
</style>

<script>
// @替代 ./src
import Head from '@/components/Head.vue'
import Host from '@/components/Host.vue'
import Key from '@/components/Key.vue'
import Transaction from '@/components/Transaction.vue'
import LatestRecords from '@/components/LatestRecords'
import { mapState } from 'vuex'

const components = {
  Head,
  Host,
  Key,
  Transaction,
  LatestRecords
}

const data = function () {
  return {
    progress: {
      label: '核心组件wasm文件较大, 下载中请稍等',
      max: 100,
      value: 0,
      display: 'block'
    },
    sendTo: ''
  }
}

const download = async function () {
  const fetchOptions = { cache: 'default' }
  // if (lastModified !== null && lastModified && cacheControl !== null && cacheControl) {
  //   console.log(lastModified)
  //   const headers = new Headers()
  //   headers.append('If-Modified-Since', lastModified)
  //   headers.append('Cache-Control', cacheControl)
  //   fetchOptions.headers = headers
  // }

  const downloadUrl = 'http://localhost:8879/js/bolaxy.wasm'
  let response
  try {
    response = await fetch(downloadUrl, fetchOptions)
  } catch (e) {
    this.progress.label = '下载失败，请检查下载链接是否正常: ' + downloadUrl
    return
  }

  let contentLength = 0
  const chunks = []
  let receivedLength = 0
  if (response.ok) {
    contentLength = +response.headers.get('Content-Length')
    this.progress.max = contentLength

    const reader = response.body.getReader()

    while (true) {
      const { done, value } = await reader.read()
      if (done) {
        this.progress.value = contentLength
        this.progress.label = '下载完毕'
        break
      }

      receivedLength = value.length + receivedLength
      this.progress.value = receivedLength

      chunks.push(value)
    }
  }

  if (chunks.length > 0 && receivedLength > 0) {
    const bytes = new Uint8Array(receivedLength)
    let position = 0
    chunks.forEach((chunk) => {
      bytes.set(chunk, position)
      position += chunk.byteLength
    })
    // console.log('bytes', bytes);

    const go = window.go

    WebAssembly.instantiate(bytes.buffer, go.importObject)
      .then(async (result) => {
        this.progress.display = 'none'
        this.$store.commit('data/setInited')
        await go.run(result.instance)
      })
      .catch(err => {
        console.log(err)
      })
  }
}

const methods = {
  download,
  toTx (address) {
    this.sendTo = address
  }
}

const computed = {
  ...mapState('data', ['host'])
}

export default {
  components,
  data,
  mounted: function () {
    this.download()
  },
  computed,
  methods
}
</script>
