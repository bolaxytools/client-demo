const state = {
  account: {}
}

const mutations = {
  setAccount (state, acc) {
    state.account = acc
  }
}

export default {
  namespaced: true,
  state,
  mutations
}
