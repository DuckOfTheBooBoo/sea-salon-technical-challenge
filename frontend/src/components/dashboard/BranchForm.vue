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
import { ArrowLeft, Plus, MapPin } from "lucide-vue-next";
import { Textarea } from "@/components/ui/textarea";
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { type Coordinate, type Branch } from "@/types";
import * as z from "zod";
import { ref, Ref, watch } from 'vue';
import { addBranch } from "@/service/branchApi";
import { toast } from "../ui/toast";

const props = defineProps<{
  currentView: string;
  location: Coordinate;
}>();

const emit = defineEmits<{
  (e: "drop-pin"): void;
  (e: "update:view", view: string): void;
  (e: "update:branch", branch: Branch): void;
  (e: "update:branch-name", name: string): void;
  (e: "reset:location"): void;
}>();

const dropPinToggle: Ref<boolean> = ref(false);

const branchFormSchema = toTypedSchema(
  z.object({
    name: z.string().min(1, { message: "Branch name is required" }),
    address: z.string().min(1, { message: "Branch address is required" }),
    lat: z.number().min(-90).max(90),
    lng: z.number().min(-180).max(180),
  })
);

const branchForm = useForm({
  validationSchema: branchFormSchema,
});

// We could watch for props.location to change, which indicates the pin drop has been done
watch(() => props.location, () => {
  dropPinToggle.value = false;
  branchForm.setFieldValue("lat", props.location.lat);
  branchForm.setFieldValue("lng", props.location.lng);
})

watch(() => branchForm.values.name, () => {
  if (!branchForm.values.name) return;
  emit("update:branch-name", branchForm.values.name);
})

const onBranchFormSubmit = branchForm.handleSubmit(async (values) => {
  const request: Branch = {
    branch_name: values.name,
    branch_address: values.address,
    lat: values.lat,
    lng: values.lng
  }

  try {
    const resp: Branch = await addBranch(request);
    emit('update:branch', resp);
    emit('update:view', 'branches')
    toast({
      title: 'Success!',
      description: `Successfully added ${request.branch_name}`
    })
    emit('reset:location')
  } catch (error: unknown) {
    const err: Error = error as Error;
    toast({
      title: 'Something went wrong',
      description: err.message
    })
  }
});
</script>

<template>
  <div class="flex items-center gap-2">
    <Button
      @click="emit('update:view', 'branches')"
      class="rounded-full w-fit h-fit"
      variant="ghost"
    >
      <ArrowLeft class="w-fit h-fit" />
    </Button>
    <h1 class="text-xl font-bold">Add new branch</h1>
  </div>
  <ScrollArea class="rounded-md h-[560px]">
    <form @submit.prevent="onBranchFormSubmit" class="px-2">
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
          <FormLabel>
            Branch address

          </FormLabel>
          <FormControl>
            <Textarea
              placeholder="Enter the branch's address details"
              class="resize-none"
              v-bind="componentField"
            />
          </FormControl>
          <FormMessage class="text-xs h-4" />
          <FormDescription class="w-fit text-xs">
            e.g. Pondok Indah Mall 2, Level 1 Jl. Metro Pondok Indah Blok III-B
            Pondok Indah, Jakarta Selatan DKI Jakarta 12310 Indonesia
          </FormDescription>
        </FormItem>
      </FormField>
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
      >
        <Plus /> Add new branch
      </Button>
    </form>
  </ScrollArea>
</template>
