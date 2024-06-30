<script setup lang="ts">
import { Button } from "@/components/ui/button";
const props = defineProps<{class?: string}>()
</script>

<template>
  <nav :class="`flex w-screen z-10 bg-white justify-between p-3 fixed ${props.class}`">
    <div class="flex gap-5">
      <span>SEA Salon</span>
      <a href="#services">Our Services</a>
      <a href="#reviews">Testimonials</a>
    </div>

    <div>
      <Button>Make a Reservation</Button>
    </div>
  </nav>
</template>
