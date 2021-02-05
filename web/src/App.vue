<template>
  <v-app id="inspire">
    <router-view/>
  </v-app>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: "App",
  computed: {
    ...mapGetters(["isAuthenticated"])
  },
  created() {
    const currentPath = this.$router.history.current.path;
    if (this.isAuthenticated === false) {
      this.$router.push("/login");
    }

    if (currentPath === "/" || currentPath === "/app") {
      this.$router.push("/app/dashboard")
    }
  }
}
</script>