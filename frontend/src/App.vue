<template>
    <Sidebar />

  <router-view />
</template>

<script lang="ts">
import Sidebar from "./components/SideBar.vue";
export default {
  name: 'App',
  components: {
    Sidebar
  }
}
</script>
