<script setup lang="ts">
import { Button } from "@/components/ui/button";
const props = defineProps<{class?: string, activeSection: string}>()
import { ref, onBeforeMount, watch } from 'vue'
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { JWTPayload } from "@/types";
import {isAuthenticated as isAuthenticatedMethod} from '@/lib/utils'

const active = ref<boolean>(props.activeSection === 'services' || props.activeSection === 'contact')
const auth = ref<JWTPayload>({} as JWTPayload)

watch(() => props.activeSection, () => {
  active.value = props.activeSection === 'services' || props.activeSection === 'contact'
})

onBeforeMount(() => {
  const isAuthenticated: JWTPayload | null = isAuthenticatedMethod()
  if (isAuthenticated) {
    auth.value = isAuthenticated;
  }
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
      <RouterLink to="/customer" v-if="auth && auth.role === 'Customer'">
        <Avatar>
            <AvatarImage :src="'https://api.dicebear.com/9.x/initials/svg?seed='+auth.full_name" alt="" />
            <AvatarFallback>{{auth.full_name}}</AvatarFallback>
        </Avatar>
      </RouterLink>
      <RouterLink to="/admin" v-else-if="auth && auth.role === 'Admin'">
        <Avatar>
            <AvatarImage :src="'https://api.dicebear.com/9.x/initials/svg?seed='+auth.full_name" alt="" />
            <AvatarFallback>{{auth.full_name}}</AvatarFallback>
        </Avatar>
      </RouterLink>
      <RouterLink to='/signup' v-else>
        <Button variant="ghost" :class="{'active-nav': active}" class="text-white bg-transparent active">Become a member</Button>
      </RouterLink>
    </div>
  </nav>
</template>

<style scoped>
.active-nav {
  color: black !important;
}
</style>