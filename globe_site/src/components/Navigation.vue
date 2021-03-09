<template>
  <div id="location" v-on:click="mapFocus('Belgium')">Belgium</div>
</template>

<script>
export default {
  name: "Navigation",
  methods: {
    mapFocus: function(location) {
      alert(location);
    }
  }
};
</script>

<style scoped>
#location {
  color: white;
}
</style>
