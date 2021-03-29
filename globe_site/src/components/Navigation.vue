<template>
  <div>
    <div class="navitem" v-on:click="mapFocus('BE')">Belgium</div>
    <div class="navitem" v-on:click="mapFocus('SG')">Singapore</div>
    <div class="navitem" v-on:click="mapFocus('EG')">Egype</div>
  </div>
</template>

<script>
import { mapMutations } from "vuex";
export default {
  name: "Navigation",
  methods: {
    ...mapMutations(["setCountry"]),
    mapFocus: function(location) {
      this.setCountry(location);
    }
  }
};
</script>

<style scoped>
.navitem {
  color: white;
}
</style>
