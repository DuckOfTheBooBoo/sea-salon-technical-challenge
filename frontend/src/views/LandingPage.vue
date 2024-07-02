<script setup lang="ts">
import ServiceCard from "@/components/ServiceCard.vue";
import { Phone } from "lucide-vue-next";
import { useToast } from "@/components/ui/toast/use-toast";
import LandingHeader from "@/components/landingPage/Header.vue";
import LandingTestimonial from "@/components/landingPage/Testimonial.vue";
import { Ref, ref } from "vue";
import HaircutImg from '@/assets/images/haircut-img-comp.webp'
import ManicureImg from '@/assets/images/manicure-img-comp.webp'
import PedicureImg from '@/assets/images/pedicure-img-comp.webp'
import FacialImg from '@/assets/images/facial-img-comp.webp'

interface Service {
  src: string;
  src2: string | undefined;
  title: string;
  description: string;
}

const { toast } = useToast();
const activeSection: Ref<string> = ref("home");

const options = {
  anchors: ["home", "services", "testimonials", "contact"],
  licenseKey: "gplv3-license",
  // @ts-ignore
  afterLoad: (origin, destination) => {
    activeSection.value = destination.anchor;
  },
};

// Service list
const services: Service[] = [
  {
    src: HaircutImg,
    src2: undefined,
    title: "Haircut",
    description:
      "Expert cuts that flatter you. Effortless style for every occasion.",
  },
  {
    src: ManicureImg,
    src2: PedicureImg,
    title: "Manicure & Pedicure",
    description:
      "Indulge in luxurious manicures & pedicures. Flawless polish, lasting confidence.",
  },
  {
    src: FacialImg,
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
    <div
      class="section bg-gradient-to-r from-primary to-[#373735] 
      justify-start sm:justify-center"
      data-anchor="home"
    >
      <div class="flex flex-col sm:flex-row w-full justify-between sm:justify-center items-center">
        <div
          class="flex flex-col w-full sm:w-[40%] h-full sm:pl-10 justify-center mt-[10rem]"
        >
          <p
            class="text-3xl text-center sm:text-left sm:text-7xl font-[900] font-oleragie text-white z-10 w-full sm:w-fit h-[4.5rem] align-baseline mt-4 leading-[4rem] sm:leading-none sm:text-nowrap"
          >
            SEA Salon
          </p>
          <p class="text-xl text-center sm:text-left sm:text-4xl text-white z-20 sm:text-nowrap">
            Beauty and Elegance, Redefined
          </p>
        </div>
        <div class="w-full overflow-hidden">
          <img
            class="w-full h-[29rem] sm:h-screen object-cover transform scale-[1.2] translate-y-[2rem] translate-x-[-1rem] drop-shadow-xl"
            src="@/assets/images/file.png"
            alt=""
          />
        </div>
      </div>
    </div>

    <div
      class="section"
      data-anchor="services"
      data-percentage="80"
      data-centered="true"
    >
      <h1 class="text-center text-4xl mt-10 mb-4 font-semibold">
        Our Services
      </h1>
      <div
        class="flex flex-col sm:flex-row gap-10 sm:gap-2 px-2 justify-center"
      >
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

    <div
      class="section bg-gradient-to-r from-primary to-[#373735]"
      data-anchor="testimonials"
    >
      <!-- Review -->
      <LandingTestimonial
        :toast="toast"
        class="bg-gradient-to-r from-primary to-[#373735]"
      />
    </div>

    <div class="section" data-anchor="contact">
      <h1 class="text-center text-4xl mb-3 font-semibold mt-10">Contact Us</h1>
      <p class="text-center mb-12">
        Got questions? We're all ears and excited to help! Reach out anytime!
      </p>
      <div class="flex flex-col sm:flex-row gap-10 justify-center items-center">
        <div class="flex flex-col items-center">
          <img
            src="@/assets/images/thomas-crop.jpeg"
            alt=""
            class="min-w-[10rem] max-w-[20rem] rounded-full"
          />
          <p class="text-center text-xl font-semibold mt-4">Thomas</p>
          <p class="text-center">
            <Phone class="inline w-4 h-4 align-middle mr-2" />08123456789
          </p>
        </div>
        <div class="flex flex-col items-center mb-10 sm:mb-0">
          <img
            src="@/assets/images/sekar-crop.jpeg"
            alt=""
            class="min-w-[10rem] max-w-[20rem] rounded-full"
          />
          <p class="text-center text-xl font-semibold mt-4">Sekar</p>
          <p class="text-center">
            <Phone class="inline w-4 h-4 align-middle mr-2" />08164829372
          </p>
        </div>
      </div>
    </div>
  </full-page>
</template>
