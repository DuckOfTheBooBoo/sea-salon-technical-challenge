<script setup lang="ts">
import Toaster from "@/components/ui/toast/Toaster.vue";
</script>

<template>
  <Toaster />
  <!-- <nav>
    <RouterLink to="/">Home</RouterLink>
    <RouterLink to="/login">Login</RouterLink>
    <RouterLink to="/signup">Sign Up</RouterLink>
    <RouterLink to="/admin">Admin</RouterLink>
    <RouterLink to="/customer">Customer</RouterLink>
  </nav> -->
  <main>
    <RouterView />
  </main>
</template>

<style scoped></style>
