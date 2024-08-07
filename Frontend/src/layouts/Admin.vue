<template>
  <div>
    <sidebar />
    <div class="relative md:ml-64 bg-blueGray-100 flex h-screen flex-col">
      <admin-navbar />
      <header-stats />
      <div class="px-4 md:px-10 mx-auto w-full">
        <router-view />
      </div>
    </div>
  </div>
</template>
<script>
import AdminNavbar from "@/components/Navbars/AdminNavbar.vue";
import Sidebar from "@/components/Sidebar/Sidebar.vue";
import HeaderStats from "@/components/Headers/HeaderStats.vue";
export default {
  name: "admin-layout",
  components: {
    AdminNavbar,
    Sidebar,
    HeaderStats,
  },
};
</script>
