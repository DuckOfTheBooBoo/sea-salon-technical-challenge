<script setup lang="ts">
import ServiceCard from "@/components/ServiceCard.vue";
import { Icon } from "@iconify/vue";
import { useToast } from "@/components/ui/toast/use-toast";
import LandingHeader from "@/components/landingPage/Header.vue";
import LandingTestimonial from "@/components/landingPage/Testimonial.vue";

interface Service {
  src: string;
  src2: string | undefined;
  title: string;
  description: string;
}

const { toast } = useToast();

const options = {
  anchors: ['section1', 'section2', 'section3', 'section4'],
}

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
  
  <LandingHeader />
  <full-page ref="fullpage" id="fullpage" :options="options">
    <div class="section">
      <div class="flex w-full h-screen justify-center items-center">
        <div class="bg-white flex flex-col w-[50%] justify-center">
          <p class="text-8xl text-center font-semibold">SEA Salon</p>
          <p class="text-3xl text-center">Beauty and Elegance, Redefined</p>
        </div>
        <div class="w-[50%]">
          <img
            class="w-full h-screen object-cover"
            src="@/assets/images/home-page-img-comp.webp"
            alt=""
          />
        </div>
      </div>
    </div>

    <div class="section">
      <div class="my-2 pt-4" id="services">
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
    </div>

    <div class="section">
      <!-- Review -->
      <LandingTestimonial :toast="toast" />
    </div>

    <div class="section">
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

<style scoped></style>
