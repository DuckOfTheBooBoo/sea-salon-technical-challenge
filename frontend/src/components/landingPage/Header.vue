<script setup lang="ts">
import {
  //  
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
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
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
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { addReservation } from '@/service/reservationApi';
import { ReservationRequest } from "@/types";
import { Input } from "@/components/ui/input";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import {ref, computed, Ref} from 'vue';

const props = defineProps<{toast: Function}>();

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

const availableTime = Array.from({ length: 20 - 9 + 1 }, (_, i) => i + 9);

const isReservationFormOpen: Ref<boolean> = ref(false);

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
      await addReservation(requestBody);
      props.toast({
        title: "Success!",
        description: "Your reservation has been submitted",
      });
      reservationForm.resetForm();
      isReservationFormOpen.value = false;
    } catch (error: unknown) {
      const err = error as Error;
      props.toast({
        title: "Something went wrong!",
        description: err.message,
      });
    }
  }
);
</script>

<template>
  <nav class="flex w-screen bg-white justify-between p-3 fixed">
    <div class="flex gap-5">
      <span>SEA Salon</span>
      <a href="#services">Our Services</a>
      <a href="#reviews">Testimonials</a>
      <RouterLink to="/login">Login</RouterLink>
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
                  <FormDescription
                    >Each reservation is one hour long.</FormDescription
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
</template>
