<template>
  <!-- Navbar -->
  <nav
    class="absolute top-0 left-0 w-full z-10 bg-transparent md:flex-row md:flex-nowrap md:justify-start flex items-center p-4"
  >
    <div
      class="w-full mx-autp items-center flex justify-between md:flex-nowrap flex-wrap md:px-10 px-4"
    >
      <!-- User -->
      <!-- <ul class="flex-col md:flex-row list-none items-center hidden md:flex">
        <user-dropdown />
      </ul> -->

      <h1 class="text-white text-3xl mt-8"> {{ viewTitle }} </h1>

    </div>
  </nav>
  <!-- End Navbar -->
</template>

<script>
export default {
  computed: {
    viewTitle() {
      return this.$route.meta.title || 'Default Title';
    }
  },
  watch: {
    '$route'() {
      document.title = this.viewTitle;
    }
  },
  mounted() {
    document.title = this.viewTitle;
  }
};
</script>
