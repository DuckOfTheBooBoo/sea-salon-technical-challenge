<script setup lang="ts">
import ServiceCard from '@/components/ServiceCard.vue'
import { Icon } from '@iconify/vue'

interface Service {
  src: string,
  src2: string | undefined,
  title: string,
  description: string
}

const services: Service[] = [
  {
    src: 'src/assets/images/haircut-img-comp.webp',
    src2: undefined,
    title: 'Haircut',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.'
  },
  {
    src: 'src/assets/images/manicure-img-comp.webp',
    src2: 'src/assets/images/pedicure-img-comp.webp',
    title: 'Manicure & Pedicure',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.'
  },
  {
    src: 'src/assets/images/facial-img-comp.webp',
    src2: undefined,
    title: 'Facial Treatment',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.'
  },
]
</script>

<template>
  <div class="flex w-screen h-screen justify-center items-center">
    <div class="bg-white flex flex-col w-[50%] justify-center">
        <p class="text-8xl text-center">SEA Salon</p>
        <p class="text-3xl text-center">Beauty and Elegance, Redefined</p>
    </div>
    <div class="w-[50%]">
        <img class="w-full h-screen object-cover" src="@/assets/images/home-page-img-comp.webp" alt="">
    </div>
  </div>
  <div class="my-2">
    <h1 class="text-center text-4xl mb-3">Our Services</h1>
    <div class="flex gap-10 px-2 justify-center">
      <ServiceCard v-for="service in services" :key="service.title" :src="service.src" :src2="service.src2" :title="service.title" :description="service.description"/>
    </div>
  </div>
  <div>
    <h1 class="text-center text-4xl mb-3">Contact Us</h1>
    <div class="flex gap-5 justify-center">
      <div class="">
        <img src="" alt="" class="rounded-full">
        <p class="text-center"><Icon icon="mdi:account" class="inline" />Thomas</p>
        <p class="text-center"><Icon icon="ph:phone-fill" class="inline" />08123456789</p>
      </div>
      <div class="">
        <img src="" alt="">
        <p class="text-center"><Icon icon="mdi:account" class="inline" />Sekar</p>
        <p class="text-center"><Icon icon="ph:phone-fill" class="inline" />08164829372</p>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
