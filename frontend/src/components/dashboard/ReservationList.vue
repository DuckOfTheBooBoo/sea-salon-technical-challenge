<script setup lang="ts">
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
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
import { cn } from "@/lib/utils";
import { addReservation } from "@/service/reservationApi";
import { Branch, ReservationRequest, Service } from "@/types";
import { Input } from "@/components/ui/input";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import { ref, computed, Ref, watch } from "vue";
import { useToast } from "@/components/ui/toast/use-toast";
import {generateAvailableTimes} from '@/lib/utils'

const { toast } = useToast();

const props = defineProps<{
  branches: Branch[];
}>();

// Reservation
const df = new DateFormatter("en-US", {
  dateStyle: "long",
});

const placeholder = ref();

const reservationFormSchema = toTypedSchema(
  z.object({
    branch: z.string().min(1, { message: "Branch is required" }),
    service: z.object({
        service_name: z.string().min(1, { message: "Service is required" }),
        service_code: z.string().min(1, { message: "Service is required" }),
        duration: z.number().min(1, { message: "Duration is required" }),
    }),
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

const getServicesFromBranch = (branchName: string): Service[] => {
  const branch = props.branches.find((b) => b.branch_name === branchName);
  if (!branch) {
    return [] as Service[];
  }
  return branch.services;
};

const availableTime = computed(() => {
  const branch = props.branches.find((b) => b.branch_name === reservationForm.values.branch);
  if (branch) {
      const services = getServicesFromBranch(branch.branch_name);
      const service = services.find((s) => s.service_name === reservationForm.values.service?.service_name)
      return generateAvailableTimes(
        branch.open_time.substring(0, 5),
        branch.close_time.substring(0, 5),
        service?.duration ?? 0
      )
  }
});

watch(() => reservationForm.values.service?.service_name, () => {
  
})

const isReservationFormOpen: Ref<boolean> = ref(false);

const onReservationSubmit = reservationForm.handleSubmit(
  async (values): Promise<void> => {
    // Construct the request body
    const requestBody: ReservationRequest = {
      service: values.service,
      date: values.date,
      time: values.time,
    };

    try {
      await addReservation(requestBody);
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
  <Card class="p-2">
    <CardHeader>
      <CardTitle>Your reservations</CardTitle>
      <CardDescription></CardDescription>
    </CardHeader>
    <CardContent>
      <div class="flex justify-end">
        <Dialog v-model:open="isReservationFormOpen">
          <DialogTrigger as-child>
            <Button variant="outline">Make reservation</Button>
          </DialogTrigger>
          <DialogContent>
            <form @submit.prevent="onReservationSubmit">
              <DialogHeader>
                <DialogTitle>Reservation Form</DialogTitle>
              </DialogHeader>
              <FormField v-slot="{ componentField }" name="branch">
                <FormItem>
                  <FormLabel>Branch</FormLabel>
                  <Select v-bind="componentField">
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder="Select a branch" />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectGroup>
                        <SelectLabel>Branches</SelectLabel>
                        <SelectItem
                          v-for="branch in branches"
                          :value="branch.branch_name"
                        >
                          {{ branch.branch_name }}
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
              <FormField v-slot="{ componentField }" name="service">
                <FormItem>
                  <FormLabel>Service type</FormLabel>
                  <Select
                    v-bind="componentField"
                    :disabled="!reservationForm.values.branch"
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
                          v-for="service in getServicesFromBranch(reservationForm.values.branch!)"
                          :value="service.service_code"
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
                <FormField v-slot="{ componentField }" name="time">
                  <FormItem>
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
                            :value="hour"
                            :key="hour"
                          >
                            {{ hour }}
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

      <ScrollArea class="h-[200px] w-[350px] rounded-md border p-4">
        Jokester began sneaking into the castle in the middle of the night and
        leaving jokes all over the place: under the king's pillow, in his soup,
        even in the royal toilet. The king was furious, but he couldn't seem to
        stop Jokester. And then, one day, the people of the kingdom discovered
        that the jokes left by Jokester were so funny that they couldn't help
        but laugh. And once they started laughing, they couldn't stop.
      </ScrollArea>
    </CardContent>
  </Card>
</template>
