<script setup lang="ts">
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
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
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { X } from "lucide-vue-next";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { cn } from "@/lib/utils";
import {
  addReservation,
  deleteReservation,
  getReservationAll,
} from "@/service/reservationApi";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import type { Branch, ReservationRequest, Reservation, Service } from "@/types";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import { ref, computed, Ref, watch, onMounted } from "vue";
import { useToast } from "@/components/ui/toast/use-toast";
import { generateAvailableTimes, parseToLocalDateTime } from "@/lib/utils";

const { toast } = useToast();

const props = defineProps<{
  branches: Branch[];
}>();

const services: Ref<Service[]> = ref([]);
const reservations: Ref<Reservation[]> = ref([]);

// Reservation
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});

const placeholder = ref();

const reservationFormSchema = toTypedSchema(
  z.object({
    branch: z.string().min(1, { message: "Branch is required" }),
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

watch(
  () => reservationForm.values.branch,
  () => {
    const branch = props.branches.find(
      (b) => b.branch_name === reservationForm.values.branch
    );
    if (branch) {
      services.value = branch.services;
    }
  }
);

const availableTime: Ref<string[]> = ref([]);
const timeRegex = /^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}$/;

watch(
  () => reservationForm.values.service,
  () => {
    const branch = props.branches.find(
      (b) => b.branch_name === reservationForm.values.branch
    );
    const service = services.value.find(
      (s) => s.service_name === reservationForm.values.service
    );

    if (branch !== undefined && service !== undefined) {
      const openTime = timeRegex.test(branch.open_time)
        ? branch.open_time.substring(11)
        : branch.open_time;
      const closeTime = timeRegex.test(branch.close_time)
        ? branch.close_time.substring(11)
        : branch.close_time;
      availableTime.value = generateAvailableTimes(
        openTime,
        closeTime,
        service.duration
      );
    }
  }
);

const isReservationFormOpen: Ref<boolean> = ref(false);

const onReservationSubmit = reservationForm.handleSubmit(
  async (values): Promise<void> => {
    // Construct the request body
    const branch = props.branches.find((b) => b.branch_name === values.branch);
    const requestBody: ReservationRequest = {
      service: values.service,
      date: values.date,
      time: values.time,
      branch_id: branch!.ID,
    };

    try {
      const resp = await addReservation(requestBody);
      toast({
        title: "Success!",
        description: "Your reservation has been submitted",
      });
      reservationForm.resetForm();
      isReservationFormOpen.value = false;
      reservations.value.push(resp)
    } catch (error: unknown) {
      const err = error as Error;
      toast({
        title: "Something went wrong!",
        description: err.message,
      });
    }
  }
);

onMounted(async () => {
  try {
    reservations.value = await getReservationAll();
  } catch (error: unknown) {
    const err = error as Error;
    toast({
      title: "Something went wrong!",
      description: err.message,
    });
  }
});

const deleteReservationMethod = async (reservation: Reservation) => {
  try {
    await deleteReservation(reservation);
    toast({
      title: "Success!",
      description: "Your reservation has been deleted",
    });
    reservations.value = reservations.value.filter(
      (r) => r.ID !== reservation.ID
    );
  } catch (error: unknown) {
    const err = error as Error;
    toast({
      title: "Something went wrong!",
      description: err.message,
    });
  }
};
</script>

<template>
  <Card class="w-full max-w-screen-md h-full">
    <CardHeader>
      <CardTitle class="text-center">Your reservations</CardTitle>
    </CardHeader>
    <CardContent class="h-full">
      <ScrollArea class="h-[90%] w-full rounded-md">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead class="w-fit text-xs sm:text-base">Branch</TableHead>
              <TableHead class="text-xs sm:text-base">Date & Time</TableHead>
              <TableHead class="text-xs sm:text-base">Service</TableHead>
              <TableHead class="text-center text-xs sm:text-base">Action</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow v-for="reservation in reservations" :key="reservation.ID">
              <TableCell class="text-xs sm:text-base">{{
                branches.find((b) => b.ID === reservation.branch_id)
                  ?.branch_name
              }}</TableCell>
              <TableCell class="text-wrap text-xs sm:text-base">{{ parseToLocalDateTime(reservation.date).toLocaleString('id-ID', { year: 'numeric', month: 'numeric', day: 'numeric', hour: 'numeric', minute: 'numeric', hour12: false }) }}</TableCell>
              <TableCell>{{ reservation.service }}</TableCell>
              <TableCell class="">
                <AlertDialog>
                  <AlertDialogTrigger as-child>
                    <Button
                      variant="outline"
                      class="py-1 px-1 h-fit"
                    >
                      <X :size="20" color="#ff0000" class="py-0" />
                    </Button>
                  </AlertDialogTrigger>
                  <AlertDialogContent>
                    <AlertDialogHeader>
                      <AlertDialogTitle
                        >Are you absolutely sure?</AlertDialogTitle
                      >
                      <AlertDialogDescription>
                        This action cannot be undone. This will cancel your RSVP and remove the data from our
                        database.
                      </AlertDialogDescription>
                    </AlertDialogHeader>
                    <AlertDialogFooter>
                      <AlertDialogCancel>Cancel</AlertDialogCancel>
                      <AlertDialogAction
                        @click="deleteReservationMethod(reservation)"
                        class="bg-destructive"
                        >Continue</AlertDialogAction
                      >
                    </AlertDialogFooter>
                  </AlertDialogContent>
                </AlertDialog>
              </TableCell>
            </TableRow>
          </TableBody>
        </Table>
        <ScrollBar orientation="horizontal" />
      </ScrollArea>

    </CardContent>
  </Card>
  <Dialog v-model:open="isReservationFormOpen">
      <DialogTrigger as-child>
        <Button variant="outline" class="w-full mt-4 max-w-screen-sm bg-primary text-white">Make reservation</Button>
      </DialogTrigger>
      <DialogContent>
        <form @submit.prevent="onReservationSubmit">
          <DialogHeader>
            <DialogTitle>Reservation Form</DialogTitle>
          </DialogHeader>
          <FormField v-slot="{ componentField }" name="branch">
            <FormItem v-auto-animate>
              <FormLabel>Branch</FormLabel>
              <Select v-bind="componentField">
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a branch" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectItem
                    v-for="branch in branches"
                    :value="branch.branch_name"
                  >
                    {{ branch.branch_name }}
                  </SelectItem>
                </SelectContent>
              </Select>
              <FormDescription
                >This will be the service you want to
                reserve.</FormDescription
              >
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="service">
            <FormItem v-auto-animate>
              <FormLabel>Service type</FormLabel>
              <Select
                v-bind="componentField"
                :disabled="services.length === 0"
              >
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a service" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  <SelectGroup>
                    <SelectLabel>Services</SelectLabel>
                    <SelectItem
                      v-for="service in services"
                      :value="service.service_name"
                    >
                      {{ service.service_name }}
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
          <div class="flex sm:flex-row flex-col sm:justify-between">
            <!-- DATE -->
            <FormField name="date" class="w-full">
              <FormItem v-auto-animate>
                <FormLabel>Reservation date</FormLabel>
                <Popover>
                  <PopoverTrigger as-child>
                    <FormControl>
                      <Button
                        variant="outline"
                        :class="
                          cn(
                            'w-[240px] ps-3 text-start font-normal w-full',
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
                            reservationForm.setFieldValue(
                              'date',
                              undefined
                            );
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
            <FormField v-slot="{ componentField }" name="time" class="w-full">
              <FormItem v-auto-animate>
                <FormLabel>Time</FormLabel>
                <Select
                  v-bind="componentField"
                  :disabled="!reservationForm.values.service"
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="Select time" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectGroup>
                      <SelectItem
                        v-for="hour in availableTime"
                        :value="`${hour}`"
                        :key="hour"
                      >
                        {{ hour }}
                      </SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
                <FormDescription
                  >Your reservation time.</FormDescription
                >
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
          </div>
          <DialogFooter>
            <Button type="submit" class="mt-4">Submit</Button>
          </DialogFooter>
        </form>
      </DialogContent>
  </Dialog>
</template>
