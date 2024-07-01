<script setup lang="ts">
import { Button } from "@/components/ui/button";
const props = defineProps<{ class?: string; activeSection: string }>();
import { ref, onMounted, watch } from "vue";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { JWTPayload } from "@/types";
import { isAuthenticated as isAuthenticatedMethod } from "@/lib/utils";
import { Menu } from "lucide-vue-next";
import {
  Sheet,
  SheetContent,
  SheetClose,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

const active = ref<boolean>(
  props.activeSection === "services" || props.activeSection === "contact"
);
const auth = ref<JWTPayload>({} as JWTPayload);

watch(
  () => props.activeSection,
  () => {
    active.value =
      props.activeSection === "services" || props.activeSection === "contact";
  }
);

onMounted(() => {
  const payload: JWTPayload | null = isAuthenticatedMethod();
  if (payload && Object.keys(payload).length > 0) {
    auth.value = payload;
  }
});
</script>

<template>
  <nav
    :class="`flex w-screen z-10 justify-between sm:justify-between p-3 fixed ${props.class}`"
  >
    <div class="flex">
      <a
        :class="{ 'active-nav': active }"
        class="nav-link text-white text-[10px] text-nowrap sm:text-xs font-oleragie pt-5 sm:pt-4 mr-2 font-[900]"
        href="#home"
        >SEA Salon</a
      >
      <div class="hidden sm:flex sm:ml-5 gap-5">
        <a
          :class="{ 'active-nav': active }"
          class="nav-link text-white pt-1 text-xl"
          href="#services"
          >Services</a
        >
        <a
          :class="{ 'active-nav': active }"
          class="nav-link text-white pt-1 text-xl"
          href="#testimonials"
          >Testimonials</a
        >
        <a
          :class="{ 'active-nav': active }"
          class="nav-link text-white pt-1 text-xl"
          href="#contact"
          >Contact</a
        >
      </div>
    </div>

    <Sheet>
      <SheetTrigger as-child>
        <Button
          variant="ghost"
          :class="{ 'active-nav': active }"
          class="sm:hidden text-white bg-transparent m-0 p-0 mr-2"
        >
          <Menu />
        </Button>
      </SheetTrigger>
      <SheetContent>
        <SheetHeader>
          <SheetTitle class="font-oleragie pt-4">SEA Salon</SheetTitle>
          <SheetClose as-child>
            <a
              class="nav-link text-right text-primary pt-1 text-xl"
              href="#services"
              >Services</a
            >
          </SheetClose>
          <SheetClose as-child>
            <a
              class="nav-link text-right text-primary pt-1 text-xl"
              href="#testimonials"
              >Testimonials</a
            >
          </SheetClose>
          <SheetClose as-child>
            <a
              class="nav-link text-right text-primary pt-1 text-xl"
              href="#contact"
              >Contact</a
            >
          </SheetClose>

          <RouterLink to="/customer" v-if="Object.keys(auth).length > 0 && auth.role === 'Customer'">
            <p class="text-right text-primary pt-1 text-xl">Dashboard</p>
          </RouterLink>
          <RouterLink to="/admin" v-else-if="Object.keys(auth).length > 0 && auth.role === 'Admin'">
            <p class="text-right text-primary pt-1 text-xl">Dashboard</p>
          </RouterLink>
          <RouterLink to="/signup" v-else>
            <p class="text-right text-primary pt-1 text-xl">Become a member</p>
          </RouterLink>
        </SheetHeader>
      </SheetContent>
    </Sheet>

    <div class="hidden sm:pr-2 sm:block">
      <RouterLink
        to="/customer"
        v-if="Object.keys(auth).length > 0 && auth.role === 'Customer'"
      >
        <Avatar>
          <AvatarImage
            :src="
              'https://api.dicebear.com/9.x/initials/svg?seed=' + auth.full_name
            "
            alt=""
          />
          <AvatarFallback>{{ auth.full_name }}</AvatarFallback>
        </Avatar>
      </RouterLink>
      <RouterLink
        to="/admin"
        v-else-if="Object.keys(auth).length > 0 && auth.role === 'Admin'"
      >
        <Avatar>
          <AvatarImage
            :src="
              'https://api.dicebear.com/9.x/initials/svg?seed=' + auth.full_name
            "
            alt=""
          />
          <AvatarFallback>{{ auth.full_name }}</AvatarFallback>
        </Avatar>
      </RouterLink>
      <RouterLink to="/signup" v-else>
        <Button
          variant="ghost"
          :class="{ 'active-nav': active }"
          class="text-white bg-transparent active"
          >Become a member</Button
        >
      </RouterLink>
    </div>
  </nav>
</template>

<style scoped>
.active-nav {
  color: black !important;
}
</style>
