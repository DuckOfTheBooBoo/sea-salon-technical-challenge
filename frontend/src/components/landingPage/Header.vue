<script setup lang="ts">
import { Button } from "@/components/ui/button";
const props = defineProps<{class?: string, activeSection: string}>()
import { ref, onMounted, watch } from 'vue'

const active = ref(props.activeSection === 'services' || props.activeSection === 'contact')


watch(() => props.activeSection, () => {
  active.value = props.activeSection === 'services' || props.activeSection === 'contact'
})
</script>

<template>
  <nav :class="`flex w-screen z-10 justify-between p-3 fixed ${props.class}`">
    <div class="flex gap-5">
      <a :class="{'active-nav': active}" class="nav-link text-white text-xs font-oleragie pt-4 mr-2 font-[900]" href="#home">SEA Salon</a>
      <a :class="{'active-nav': active}" class="nav-link text-white pt-1 text-xl" href="#services">Our Services</a>
      <a :class="{'active-nav': active}" class="nav-link text-white pt-1 text-xl" href="#testimonials">Testimonials</a>
      <a :class="{'active-nav': active}" class="nav-link text-white pt-1 text-xl" href="#contact">Contact</a>
    </div>

    <div>
      <Button variant="ghost" :class="{'active-nav': active}" class="text-white bg-transparent active">Become a member</Button>
    </div>
  </nav>
</template>

<style scoped>
.active-nav {
  color: black !important;
}
</style>