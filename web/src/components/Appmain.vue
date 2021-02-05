<!--  -->
<template>
  <div>
    <transition>
      <keep-alive>
        <router-view :key="key" />
      </keep-alive>
    </transition>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },

  components: {},

  computed: {
    key() {
      return this.$route.name !== undefined
        ? this.$route.name + Math.random()
        : this.$route + Math.random();
    }
  },

  //   mounted: {},

  methods: {}
};
</script>
<style lang="scss" scoped></style>
