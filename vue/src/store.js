import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

// root state object.
// each Vuex instance is just a single state tree.
const state = {
  count: 0,
  chosenHeroList : [ ],
  heroes: [
    { name: "Superman" },
    { name: "Batman" },
    { name: "Aquaman" },
    { name: "Wonder Woman" },
    { name: "Green Lantern" },
    { name: "Martian Manhunter" },
    { name: "Flash" }
  ]
}

// mutations are operations that actually mutate the state.
// each mutation handler gets the entire state tree as the
// first argument, followed by additional payload arguments.
// mutations must be synchronous and can be recorded by plugins
// for debugging purposes.
const mutations = {
  increment (state) {
    state.count++
  },
  decrement (state) {
    state.count--
  },
  addChosenHeroToChosenHeroList (state, hero) {
      state.chosenHeroList.push(hero)
  },
  removeChosenHeroFromChosenHeroList (state, hero) {
    state.chosenHeroList = state.chosenHeroList.filter(h => h !== hero)
  }
}

// actions are functions that cause side effects and can involve
// asynchronous operations.
const actions = {
  increment: ({ commit }) => commit('increment'),
  decrement: ({ commit }) => commit('decrement'),
  addChosenHeroToChosenHeroList: ({ commit }, hero) => commit('addChosenHeroToChosenHeroList', hero),
  removeChosenHeroFromChosenHeroList: ({commit}, hero) => commit('removeChosenHeroFromChosenHeroList', hero),
  incrementIfOdd ({ commit, state }) {
    if ((state.count + 1) % 2 === 0) {
      commit('increment')
    }
  },
  incrementAsync ({ commit }) {
    return new Promise((resolve) => {
      setTimeout(() => {
        commit('increment')
        resolve()
      }, 1000)
    })
  }
}

// getters are functions.
const getters = {
  evenOrOdd: state => state.count % 2 === 0 ? 'even' : 'odd',
  chosenHerosList: state => state.chosenHeroList,
}

// A Vuex instance is created by combining the state, mutations, actions,
// and getters.
export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})