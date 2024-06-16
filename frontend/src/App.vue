<script setup lang="ts">
import ServiceCard from "@/components/ServiceCard.vue";
import { Icon } from "@iconify/vue";
import { Ref, ref } from "vue";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input, StarRating } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Toaster from '@/components/ui/toast/Toaster.vue'
import { useToast } from '@/components/ui/toast/use-toast'

interface Service {
  src: string;
  src2: string | undefined;
  title: string;
  description: string;
}

interface Review {
  name: string;
  rating: number;
  comment: string;
}

const { toast } = useToast()

const reviewFormSchema = toTypedSchema(
  z.object({
    name: z.string().nonempty({ message: "Name is required" }),
    rating: z
      .number()
      .min(1, { message: "Rating is required" })
      .max(5, { message: "Rating must be between 1 and 5" }),
    comment: z.string().nonempty({ message: "Comment is required" }),
  })
);

const reviewForm = useForm({
  validationSchema: reviewFormSchema,
});

const services: Service[] = [
  {
    src: "src/assets/images/haircut-img-comp.webp",
    src2: undefined,
    title: "Haircut",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
  },
  {
    src: "src/assets/images/manicure-img-comp.webp",
    src2: "src/assets/images/pedicure-img-comp.webp",
    title: "Manicure & Pedicure",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
  },
  {
    src: "src/assets/images/facial-img-comp.webp",
    src2: undefined,
    title: "Facial Treatment",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
  },
];

const reviews: Ref<Review[]> = ref([
  {
    name: "Emily Johnson",
    rating: 4,
    comment: "I loved my haircut at SEA Salon! The stylist was very skilled and understood exactly what I wanted."
  },
  {
    name: "David Smith",
    rating: 5,
    comment: "Outstanding service! The pedicure and manicure were done professionally. My nails look perfect."
  },
  {
    name: "Sophia Lee",
    rating: 4,
    comment: "The facial treatment was refreshing and relaxing. I felt rejuvenated afterwards."
  },
  {
    name: "Michael Brown",
    rating: 3,
    comment: "Decent salon overall. Had a good haircut, but the waiting time was a bit long."
  },
  {
    name: "Emma Davis",
    rating: 4,
    comment: "Great experience! The ambiance was pleasant and the staff were friendly and attentive.",
  }
]);

const onSubmit = reviewForm.handleSubmit((values) => {
  const newReview: Review = {
    name: values.name,
    rating: values.rating,
    comment: values.comment,
  }
  reviews.value.push(newReview)
  toast({
    title: 'Success!',
    description: 'Your review has been submitted',
  })
  reviewForm.resetForm()
});
</script>

<template>
  <Toaster />
  <div class="flex w-screen h-screen justify-center items-center">
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
  <div class="my-2">
    <h1 class="text-center text-4xl mb-3 font-semibold">Our Services</h1>
    <div class="flex gap-10 px-2 justify-center">
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

  <!-- Review -->
  <div class="h-[500px]">
    <h1 class="text-center text-4xl mb-3 font-semibold">What Others Said About Us</h1>
    <div class="flex px-32 h-full items-center">
      <div class="w-1/2">
        <Card>
          <form @submit="onSubmit">
            <CardHeader>
              <CardTitle>Leave a review</CardTitle>
            </CardHeader>
            <CardContent>
              <FormField v-slot="{ componentField }" name="name">
                <FormItem>
                  <FormLabel>Your name</FormLabel>
                  <FormControl>
                    <Input
                      type="text"
                      placeholder="Enter your full name here"
                      v-bind="componentField"
                    />
                  </FormControl>
                  <FormMessage class="text-xs h-4" />
                </FormItem>
              </FormField>
              <FormField v-slot="{ componentField }" name="rating">
                <FormItem>
                  <FormLabel>Your rating</FormLabel>
                  <FormControl>
                    <StarRating
                      v-bind="componentField"
                      :model-value="0"
                      :isWriteable="true"
                    />
                  </FormControl>
                  <FormMessage class="text-xs h-4" />
                </FormItem>
              </FormField>
              <FormField v-slot="{ componentField }" name="comment">
                <FormItem>
                  <FormLabel>Comment</FormLabel>
                  <FormControl>
                    <Input
                      type="text"
                      placeholder="Enter your comment"
                      v-bind="componentField"
                    />
                  </FormControl>
                  <FormMessage class="text-xs" />
                </FormItem>
              </FormField>
            </CardContent>
            <CardFooter>
              <Button type="submit">Submit</Button>
            </CardFooter>
          </form>
        </Card>
      </div>
      <div class="w-full"></div>
    </div>
  </div>

  <div class="">
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
</template>

<style scoped></style>
