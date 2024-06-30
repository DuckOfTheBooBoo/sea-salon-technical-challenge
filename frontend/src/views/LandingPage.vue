<script setup lang="ts">
import ServiceCard from "@/components/ServiceCard.vue";
import { Icon } from "@iconify/vue";
import { useToast } from "@/components/ui/toast/use-toast";
import LandingHeader from "@/components/landingPage/Header.vue";
import LandingTestimonial from "@/components/landingPage/Testimonial.vue";
import {Ref, ref} from 'vue'

interface Service {
  src: string;
  src2: string | undefined;
  title: string;
  description: string;
}

const { toast } = useToast();
const activeSection: Ref<string> = ref('home')

const options = {
  anchors: ['home', 'services', 'testimonials', 'contact'],
  licenseKey: 'gplv3-license',
  afterLoad: (origin, destination) => {
    activeSection.value = destination.anchor;
  },
};

// Service list
const services: Service[] = [
  {
    src: "src/assets/images/haircut-img-comp.webp",
    src2: undefined,
    title: "Haircut",
    description:
      "Expert cuts that flatter you. Effortless style for every occasion.",
  },
  {
    src: "src/assets/images/manicure-img-comp.webp",
    src2: "src/assets/images/pedicure-img-comp.webp",
    title: "Manicure & Pedicure",
    description:
      "Indulge in luxurious manicures & pedicures. Flawless polish, lasting confidence.",
  },
  {
    src: "src/assets/images/facial-img-comp.webp",
    src2: undefined,
    title: "Facial Treatment",
    description:
      "Revive & restore your skin. Personalized treatments for a radiant you.",
  },
];
</script>

<template>
  
  <LandingHeader :activeSection="activeSection" />
  <full-page ref="fullpage" id="fullpage" :options="options">
    <div class="section bg-gradient-to-r from-primary to-[#373735]" data-anchor="home">
      <div class="flex w-full justify-center items-center">
        <div class="gradient flex flex-col w-[40%] h-full pl-10 justify-center mt-[10rem]">
          <p class="text-7xl font-[900] font-oleragie text-white z-10 text-nowrap w-fit h-[4.5rem] align-baseline mt-4">SEA Salon</p>
          <p class="text-4xl text-white z-20 text-nowrap">Beauty and Elegance, Redefined</p>
        </div>
        <div class="">
          <img
            class="w-full h-screen object-cover"
            src="@/assets/images/file.png"
            alt=""
          />
        </div>
      </div>
    </div>

    <div class="section" data-anchor="services" data-percentage="80" data-centered="true">
      <h1 class="text-center text-4xl mt-10 mb-4 font-semibold">Our Services</h1>
      <div class="flex gap-10 sm:gap-2 px-2 justify-center">
        <ServiceCard
          v-for="service in services"
          :key="service.title"
          :src="service.src"
          :src2="service.src2"
          :title="service.title"
          :description="service.description"
        />
      </div>
    </div>

    <div class="section bg-gradient-to-r from-primary to-[#373735]" data-anchor="testimonials">
      <!-- Review -->
      <LandingTestimonial :toast="toast" class="bg-gradient-to-r from-primary to-[#373735]" />
    </div>

    <div class="section" data-anchor="contact">
        <h1 class="text-center text-4xl mb-3 font-semibold">Contact Us</h1>
        <div class="flex gap-5 justify-center">
          <div class="">
            <img src="" alt="" class="rounded-full" />
            <p class="text-center">
              <Icon icon="mdi:account" class="inline" />Thomas
            </p>
            <p class="text-center">
              <Icon icon="ph:phone-fill" class="inline" />08123456789
            </p>
          </div>
          <div class="">
            <img src="" alt="" />
            <p class="text-center">
              <Icon icon="mdi:account" class="inline" />Sekar
            </p>
            <p class="text-center">
              <Icon icon="ph:phone-fill" class="inline" />08164829372
            </p>
          </div>
        </div>
    </div>
  </full-page>
</template>

<style scoped>
/* .gradient {
  background: linear-gradient(to right, rgba(255,255,255,1), #373735);
} */
</style>
