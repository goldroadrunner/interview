<template>
  <div>
    <select v-model="chosenHero">
      <!-- placeholder value -->
      <option :value="null">Select a hero</option>

      <!-- available heroes -->
      <option v-for="hero in heroes.filter(h => ! $store.getters.chosenHerosList.find(hh => h.name === hh.name))"
              :key="hero.name"
              :value="hero">
        {{ hero.name }}
      </option>
    </select>
    <span>&nbsp;</span>
    <button @click="addHero(chosenHero)"
            :disabled="chosenHero === null">Add Hero</button>
    <br>
    <h3>EMORY Chosen Heroes</h3>
    <div class="chosen-heroes">
      <div v-for="(hero, i) in $store.getters.chosenHerosList"
           :key="hero.name">
        <strong>Slot {{ i + 1 }}:</strong>
        <Hero :hero="hero"
              @removeHero="removeHero(hero)" />
      </div>
    </div>
  </div>
</template>

<script>
import Hero from "./Hero";

export default {
  components: {
    Hero
  },
  props: ["heroes"],
  data() {
    return {
      chosenHero: null,
      chosenHeroes: []
    };
  },
  methods: {
    addHero(name) {
      this.$store.dispatch('addChosenHeroToChosenHeroList', name);
      this.chosenHero = null;
    },
    removeHero(hero) {
      this.$store.dispatch('removeChosenHeroFromChosenHeroList', hero);
    }
  }
};
</script>

<style scoped>
.chosen-heroes {
  display: flex;
  flex-flow: column;
  align-items: center;
}
</style>


