<script setup lang="ts">
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Toggle } from "@/components/ui/toggle";
import { ArrowLeft, Plus, MapPin, Pencil } from "lucide-vue-next";
import { Textarea } from "@/components/ui/textarea";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import ServiceInput from "./ServiceInput.vue";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { type Coordinate, type Branch } from "@/types";
import * as z from "zod";
import { ref, Ref, watch, onMounted } from "vue";
import { addBranch, updateBranch } from "@/service/branchApi";
import { useToast } from "../ui/toast";

const props = defineProps<{
  currentView: string;
  location: Coordinate;
  branchData?: Branch;
}>();

const { toast } = useToast();

const emit = defineEmits<{
  (e: "drop-pin"): void;
  (e: "update:view", view: string): void;
  (e: "update:branch", branch: Branch, oldBranch?: Branch): void;
  (e: "update:branch-name", name: string): void;
  (e: "reset:location"): void;
}>();

const dropPinToggle: Ref<boolean> = ref(false);
const timeRegex = /^([0-1][0-9]|2[0-3]):([0-5][0-9])$/;

const branchFormSchema = toTypedSchema(
  z.object({
    name: z.string().min(1, { message: "Branch name is required" }),
    address: z.string().min(1, { message: "Branch address is required" }),
    lat: z.number().min(-90).max(90),
    lng: z.number().min(-180).max(180),
    services: z.array(z.string()).nonempty(),
    openTime: z
      .string()
      .min(1, { message: "Open time is required" })
      .regex(timeRegex, {
        message: "Open time must be in format hh:mm",
      }),
    closeTime: z
      .string()
      .min(1, { message: "Close time is required" })
      .regex(timeRegex, {
        message: "Close time must be in format hh:mm",
      }),
  })
);

const branchForm = useForm({
  validationSchema: branchFormSchema,
});

// We could watch for props.location to change, which indicates the pin drop has been done
watch(
  () => props.location,
  () => {
    dropPinToggle.value = false;
    branchForm.setFieldValue("lat", props.location.lat);
    branchForm.setFieldValue("lng", props.location.lng);
  }
);

watch(
  () => branchForm.values.name,
  () => {
    if (!branchForm.values.name) return;
    emit("update:branch-name", branchForm.values.name);
  }
);

const handleServiceInput = (values: string[]) => {
  // @ts-ignore
  branchForm.setFieldValue("services", values);
};

const onBranchFormSubmit = branchForm.handleSubmit(async (values) => {
  const request = {
    branch_name: values.name,
    branch_address: values.address,
    lat: values.lat,
    lng: values.lng,
    open_time: values.openTime,
    close_time: values.closeTime,
    services: values.services.map((service: string) => {
      let normalized_string = service.toLowerCase();
      normalized_string = normalized_string.replace(/[^a-z0-9\s]/g, ' ');
      normalized_string = normalized_string.replace(/\s+/g, '-').trim().replace(/-{2,}/g, '-');
      return normalized_string.replace(/\s+/g, '-').replace(/\s/g, '-');
    }),
  };

  // Update
  if (
    props.branchData !== undefined &&
    Object.keys(props.branchData).length > 0
  ) {
    try {
      const resp: Branch = await updateBranch(props.branchData, request);
      emit("update:branch", resp, props.branchData);
      emit("update:view", "branches");
      toast({
        title: "Success!",
        description: `Successfully updated ${request.branch_name}`,
      });
      emit("reset:location");
    } catch (error: unknown) {
      const err: Error = error as Error;
      toast({
        title: "Something went wrong",
        description: err.message,
      });
    }

    return;
  }

  try {
    // @ts-ignore
    const resp: Branch = await addBranch(request);
    emit("update:branch", resp);
    emit("update:view", "branches");
    toast({
      title: "Success!",
      description: `Successfully added ${request.branch_name}`,
    });
    emit("reset:location");
  } catch (error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message,
    });
  }
});

onMounted(() => {
  if (
    props.branchData !== undefined &&
    Object.keys(props.branchData).length > 0
  ) {
    branchForm.setFieldValue("name", props.branchData.branch_name);
    branchForm.setFieldValue("address", props.branchData.branch_address);
    branchForm.setFieldValue("lat", props.branchData.lat);
    branchForm.setFieldValue("lng", props.branchData.lng);
    branchForm.setFieldValue(
      "openTime",
      props.branchData.open_time.substring(0, 5)
    );
    branchForm.setFieldValue(
      "closeTime",
      props.branchData.close_time.substring(0, 5)
    );
    branchForm.setFieldValue(
      "services",
      // @ts-ignore
      props.branchData.services.map((s) => s.service_name)
    );
  }
});
</script>

<template>
  <div class="w-full h-full">
    <div class="flex items-center gap-2 py-2 h-[3rem]">
      <Button
        @click="
          () => {
            emit('update:view', 'branches');
            branchForm.resetForm();
          }
        "
        class="rounded-full w-fit h-fit ml-0 pl-0"
        variant="ghost"
      >
        <ArrowLeft class="w-fit h-fit" />
      </Button>
      <h1 class="text-xl font-bold">Add new branch</h1>
    </div>
    <ScrollArea class="rounded-md w-full h-[35rem]">
      <form @submit.prevent="onBranchFormSubmit" class="">
        <FormField v-slot="{ componentField }" name="name">
          <FormItem>
            <FormLabel>Branch name</FormLabel>
            <FormControl>
              <Input
                type="text"
                placeholder="Enter the new branch name"
                v-bind="componentField"
              />
            </FormControl>
            <FormDescription class="w-fit text-xs">
              e.g. Branch-Pondok Indah Mall II
            </FormDescription>
            <FormMessage class="text-xs h-4" />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="address">
          <FormItem>
            <FormLabel> Branch address </FormLabel>
            <FormControl>
              <Textarea
                placeholder="Enter the branch's address details"
                class="resize-none"
                v-bind="componentField"
              />
            </FormControl>
            <FormMessage class="text-xs h-4" />
            <FormDescription class="w-fit text-xs">
              e.g. Pondok Indah Mall 2, Level 1 Jl. Metro Pondok Indah Blok
              III-B Pondok Indah, Jakarta Selatan DKI Jakarta 12310 Indonesia
            </FormDescription>
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="services" class="mb-5">
          <FormItem>
            <FormLabel> Services </FormLabel>
            <FormControl>
              <ServiceInput
                v-bind="componentField"
                @update:services="handleServiceInput"
                :services="branchData!.services"
                class="w-full"
              />
            </FormControl>
            <FormDescription class="w-fit text-xs">
              Services available in this branch
            </FormDescription>
            <FormMessage class="text-xs h-4" />
          </FormItem>
        </FormField>
        <div class="flex sm:flex-row flex-col w-full gap-2">
          <FormField v-slot="{ componentField }" name="openTime">
            <FormItem>
              <FormLabel>Open time</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter open time in 24H"
                  v-bind="componentField"
                />
              </FormControl>
              <FormDescription class="w-fit text-xs">
                eg. 09:00
              </FormDescription>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="closeTime">
            <FormItem>
              <FormLabel>Close time</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter close time in 24H"
                  v-bind="componentField"
                />
              </FormControl>
              <FormDescription class="w-fit text-xs">
                eg. 19:00
              </FormDescription>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
        </div>

        <FormField v-slot="{ componentField }" name="lat">
          <FormItem>
            <FormLabel>Latitude</FormLabel>
            <FormControl>
              <Input
                type="number"
                step="any"
                placeholder="Enter the branch latitude"
                v-bind="componentField"
              />
            </FormControl>
            <FormDescription class="w-fit text-xs">
              eg. Latitude: 7.331748
            </FormDescription>
            <FormMessage class="text-xs h-4" />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="lng">
          <FormItem>
            <FormLabel>Longitude</FormLabel>
            <FormControl>
              <Input
                type="number"
                step="any"
                placeholder="Enter the branch longitude"
                v-bind="componentField"
              />
            </FormControl>
            <FormDescription class="w-fit text-xs">
              e.g Longitude: 26.717810
            </FormDescription>
            <FormMessage class="text-xs h-4" />
          </FormItem>
        </FormField>
        <Separator />
        <Toggle
          v-model:pressed="dropPinToggle"
          class="w-full my-2 flex gap-3"
          @click="emit('drop-pin')"
          variant="outline"
        >
          <MapPin /> Drop a pin
        </Toggle>
        <Button
          type="submit"
          :disabled="false"
          class="w-full flex gap-3"
          variant="outline"
          v-if="Object.keys(branchData!).length === 0"
        >
          <Plus /> Add new branch
        </Button>
        <Button
          type="submit"
          :disabled="false"
          class="w-full flex gap-3"
          variant="outline"
          v-else
        >
          <Pencil /> Update branch
        </Button>
      </form>
    </ScrollArea>
  </div>
</template>
