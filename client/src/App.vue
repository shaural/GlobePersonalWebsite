<template>
  <div id="app" style="background-color: #000000;">
    <Navigation />
    <Globe />
  </div>
</template>

<script>
import Globe from "./components/Globe.vue";
import Navigation from "./components/Navigation.vue";

export default {
  name: "App",
  components: {
    Navigation,
    Globe
  }
};
</script>
