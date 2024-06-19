<script setup lang="ts">
import ServiceCard from "@/components/ServiceCard.vue";
import {
  CalendarDate,
  DateFormatter,
  getLocalTimeZone,
  parseDate,
  today,
} from "@internationalized/date";
import { toDate } from "radix-vue/date";
import { Calendar as CalendarIcon } from "lucide-vue-next";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Icon } from "@iconify/vue";
import { Ref, ref, computed } from "vue";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import { ScrollArea } from "@/components/ui/scroll-area";
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
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input, StarRating } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Toaster from "@/components/ui/toast/Toaster.vue";
import Review from "@/components/Review.vue";
import { useToast } from "@/components/ui/toast/use-toast";
import { cn } from "@/lib/utils";
import { addReservation } from '@/service/reservationApi';
import { ReservationRequest } from "@/types";

interface Service {
  src: string;
  src2: string | undefined;
  title: string;
  description: string;
}

interface ReviewType {
  name: string;
  rating: number;
  comment: string;
}

const { toast } = useToast();

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

// Reservation
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});

const placeholder = ref();

const reservationFormSchema = toTypedSchema(
  z.object({
    name: z.string().min(1, { message: "Name is required" }),
    phoneNumber: z.string().min(1, { message: "Phone number is required" }),
    service: z.string().min(1, { message: "Service is required" }),
    date: z.string().date().min(1, { message: "Date is required" }),
    time: z.string().min(1, { message: "Time is required" }),
  })
);

const reservationForm = useForm({
  validationSchema: reservationFormSchema,
});

const dateValue = computed({
  get: () =>
    reservationForm.values.date
      ? parseDate(reservationForm.values.date)
      : undefined,
  set: (val) => val,
});

const availableTime = computed(() =>
  Array.from({ length: 20 - 9 + 1 }, (_, i) => i + 9)
);

const isReservationFormOpen: Ref<boolean> = ref(false);

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

// Temporary reviews array state
const reviews: Ref<ReviewType[]> = ref([
  {
    name: "Emily Johnson",
    rating: 5,
    comment:
      "I loved my haircut at SEA Salon! The stylist was very skilled and understood exactly what I wanted.",
  },
  {
    name: "David Smith",
    rating: 5,
    comment:
      "Outstanding service! The pedicure and manicure were done professionally. My nails look perfect.",
  },
  {
    name: "Sophia Lee",
    rating: 4,
    comment:
      "The facial treatment was refreshing and relaxing. I felt rejuvenated afterwards.",
  },
  {
    name: "Michael Brown",
    rating: 3,
    comment:
      "Decent salon overall. Had a good haircut, but the waiting time was a bit long.",
  },
  {
    name: "Emma Davis",
    rating: 4,
    comment:
      "Great experience! The ambiance was pleasant and the staff were friendly and attentive.",
  },
]);

const onReviewSubmit = reviewForm.handleSubmit((values) => {
  const newReview: ReviewType = {
    name: values.name,
    rating: values.rating,
    comment: values.comment,
  };
  reviews.value.unshift(newReview);
  toast({
    title: "Success!",
    description: "Your review has been submitted",
  });
  reviewForm.resetForm();
});

const onReservationSubmit = reservationForm.handleSubmit(async (values): Promise<void> => {
    // Construct the request body
    const requestBody: ReservationRequest = {
      full_name: values.name,
      phoneNumber: values.phoneNumber,
      service: values.service,
      date: values.date,
      time: values.time,
    };

    try {
      const responseData = await addReservation(requestBody);
      toast({
        title: "Success!",
        description: "Your reservation has been submitted",
      });
      reservationForm.resetForm();
      isReservationFormOpen.value = false;
    } catch (error: unknown) {
      const err = error as Error;
      toast({
        title: "Something went wrong!",
        description: err.message,
      });
    }
  }
);
</script>

<template>
  <Toaster />

  <nav class="flex justify-between p-3">
    <div class="flex gap-5">
      <span>SEA Salon</span>
      <a href="#services">Our Services</a>
      <a href="#reviews">Testimonials</a>
    </div>

    <div>
      <Dialog v-model:open="isReservationFormOpen">
        <DialogTrigger as-child>
          <Button>Make a Reservation</Button>
        </DialogTrigger>
        <DialogContent>
          <form @submit.prevent="onReservationSubmit">
            <DialogHeader>
              <DialogTitle>Reservation Form</DialogTitle>
            </DialogHeader>
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
                <FormDescription
                  >This will be your reservation name.</FormDescription
                >
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
            <FormField v-slot="{ componentField }" name="phoneNumber">
              <FormItem>
                <FormLabel>Your active phone number</FormLabel>
                <FormControl>
                  <Input
                    type="tel"
                    placeholder="Enter your active phone number here"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormDescription
                  >This will be your phone number.</FormDescription
                >
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
            <FormField v-slot="{ componentField }" name="service">
              <FormItem>
                <FormLabel>Service type</FormLabel>
                <Select v-bind="componentField">
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="Select a service" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectGroup>
                      <SelectLabel>Services</SelectLabel>
                      <SelectItem value="hairstyle">
                        Haircut & styling
                      </SelectItem>
                      <SelectItem value="manipedi">
                        Manicure & pedicure
                      </SelectItem>
                      <SelectItem value="facial">
                        Facial treatments
                      </SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
                <FormDescription
                  >This will be the service you want to
                  reserve.</FormDescription
                >
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
            <div class="flex">
              <!-- DATE -->
              <FormField name="date">
                <FormItem>
                  <FormLabel>Reservation date</FormLabel>
                  <Popover>
                    <PopoverTrigger as-child>
                      <FormControl>
                        <Button
                          variant="outline"
                          :class="
                            cn(
                              'w-[240px] ps-3 text-start font-normal',
                              !dateValue && 'text-muted-foreground'
                            )
                          "
                        >
                          <span>{{
                            dateValue
                              ? df.format(toDate(dateValue))
                              : "Pick a date"
                          }}</span>
                          <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                        </Button>
                        <input hidden />
                      </FormControl>
                    </PopoverTrigger>
                    <PopoverContent class="w-auto p-0">
                      <Calendar
                        v-model:placeholder="placeholder"
                        v-model="dateValue"
                        calendar-label="Reservation Date"
                        initial-focus
                        :min-value="today(getLocalTimeZone())"
                        @update:model-value="
                          (v) => {
                            if (v) {
                              reservationForm.setFieldValue(
                                'date',
                                v.toString()
                              );
                            } else {
                              reservationForm.setFieldValue('date', undefined);
                            }
                          }
                        "
                      />
                    </PopoverContent>
                  </Popover>
                  <FormDescription class="w-fit"
                    >This will be your reservation date.</FormDescription
                  >
                  <FormMessage class="text-xs h-4" />
                </FormItem>
              </FormField>

              <!-- TIME -->
              <FormField v-slot="{ componentField }" name="time">
                <FormItem>
                  <FormLabel>Time</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select time" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem
                          v-for="hour in availableTime"
                          :value="`${hour}:00`"
                          :key="hour"
                        >
                          {{ `${hour}:00` }}
                        </SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                  <FormDescription>Each reservation is one hour long.</FormDescription
                  >
                  <FormMessage class="text-xs h-4" />
                </FormItem>
              </FormField>
            </div>
            <DialogFooter>
              <Button type="submit">Submit</Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  </nav>

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

  <div class="my-2" id="services">
    <h1 class="text-center text-4xl my-10 font-semibold">Our Services</h1>
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
  <div class="h-[530px] m-h-[530px]" id="reviews">
    <h1 class="text-center text-4xl my-10 font-semibold">
      What Others Said About Us
    </h1>
    <div class="flex px-32 w-full justify-between gap-4">
      <Card class="w-2/3">
        <form @submit="onReviewSubmit">
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
      <Card class="w-full">
        <CardHeader>
          <CardTitle>Testimonials</CardTitle>
        </CardHeader>
        <CardContent class="">
          <ScrollArea class="rounded-md p-4 h-[300px]">
            <Review
              v-for="(review, index) in reviews"
              :key="index"
              :name="review.name"
              :rating="review.rating"
              :comment="review.comment"
            />
          </ScrollArea>
        </CardContent>
      </Card>
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
