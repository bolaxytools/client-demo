const storeKey = 'bolaxy'

const initLocalDB = function (state) {
  const localStore = localStorage.getItem(storeKey)
  if (localStore != null) {
    const doc = JSON.parse(localStore)
    Object.keys(doc).forEach((key) => {
      state[key] = doc[key]
    })
  }
}

const setLocalDB = function (state) {
  const data = Object.assign({}, state)
  delete data.inited
  const value = JSON.stringify(data)
  localStorage.setItem(storeKey, value)
}

const state = {
  inited: false,
  host: '',
  keys: [{
    address: '0x3D339aA0598c700aD8F93946fBcA412E65d817dd',
    hexkey: '0x1f9b03a26b50bc8e370606ab55266014cf15eb255191e7188e3189c40dda9888'
  }],
  txs: []
}

initLocalDB(state)

const mutations = {
  setHost (state, host) {
    state.host = host
  },
  setInited (state) {
    state.inited = true
  },
  addKey (state, key) {
    state.keys.push(key)
  },
  removeKey (state, address) {
    state.keys = state.keys.filter((key) => key.address !== address)
  },
  addTx (state, tx) {
    if (state.txs.length >= 20) {
      state.txs.pop()
    }
    state.txs.unshift(tx)
  }
}

const syncToDB = function (state) {
  return new Promise(resolve => {
    setLocalDB(state)
    resolve()
  })
}

const actions = {
  setHostAsync ({ commit, state }, host) {
    commit('setHost', host)
    return syncToDB(state)
  },
  addKeyAsync ({ commit, state }, key) {
    commit('addKey', key)
    return syncToDB(state)
  },
  removeKeyAsync ({ commit, state }, address) {
    commit('removeKey', address)
    return syncToDB(state)
  },
  addTxAsync ({ commit, state }, tx) {
    commit('addTx', tx)
    return syncToDB(state)
  }
}

const getters = {
  isDisabled (state) {
    return !state.inited
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
