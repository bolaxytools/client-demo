import Vue from 'vue'
import Vuex from 'vuex'
import data from './data.js'
import event from './event.js'

Vue.use(Vuex)

const modules = {
  data,
  event
}

export default new Vuex.Store({
  modules
})
