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
import {ArrowLeft, Plus, MapPin} from 'lucide-vue-next'
import { Textarea } from "@/components/ui/textarea"
import { Separator } from "@/components/ui/separator";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { type Coordinate } from "@/types";
import * as z from "zod";

defineProps<{
  currentView: string;
  location: Coordinate;
}>();

const emit = defineEmits<{
  (e: 'drop-pin'): void;
  (e: 'update:view', view: string): void;
}>();

const branchFormSchema = toTypedSchema(
  z.object({
    name: z.string().min(1, { message: "Branch name is required" }),
    address: z.string().min(1, { message: "Branch address is required" }),
    lat: z.number().min(1, { message: "Latitude is required" }),
    lng: z.number().min(1, { message: "Longitude is required" }),
  })
);

const branchForm = useForm({
  validationSchema: branchFormSchema,
}); 

const onBranchFormSubmit = branchForm.handleSubmit(async (values) => {
  
})
</script>

<template>
  <div class="flex items-center gap-2">
    <Button @click="emit('update:view', 'branches')" class="rounded-full w-fit h-fit" variant="ghost">
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
          <FormLabel>Branch address</FormLabel>
          <FormControl>
            <Textarea
              placeholder="Enter the branch's address details"
              class="resize-none"
              v-bind="componentField"
            />
          </FormControl>
          <FormDescription class="w-fit text-xs">
            e.g. Pondok Indah Mall 2, Level 1 Jl. Metro Pondok Indah Blok III-B
            Pondok Indah, Jakarta Selatan DKI Jakarta 12310 Indonesia
          </FormDescription>
          <FormMessage class="text-xs h-4" />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="lat">
        <FormItem>
          <FormLabel>Latitude</FormLabel>
          <FormControl>
            <Input
              type="text"
              placeholder="Enter the branch latitude"
              v-bind="componentField"
              v-model="location.lat"
            />
          </FormControl>
          <FormDescription class="w-fit text-xs">
            eg. Latitude: 7.331748
          </FormDescription>
          <FormMessage class="text-xs h-4" />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="long">
        <FormItem>
          <FormLabel>Longitude</FormLabel>
          <FormControl>
            <Input
              type="text"
              placeholder="Enter the branch longitude"
              v-bind="componentField"
              v-model="location.lng"
            />
          </FormControl>
          <FormDescription class="w-fit text-xs">
            e.g Longitude: 26.717810
          </FormDescription>
          <FormMessage class="text-xs h-4" />
        </FormItem>
      </FormField>
      <Separator />
      <Button class="w-full my-2 flex gap-3" variant="outline" @click="emit('drop-pin')">
        <MapPin /> Drop a pin
      </Button>
      <Button type="submit" :disabled="false" class="w-full flex gap-3" variant="outline">
        <Plus /> Add new branch
      </Button>
    </form>
  </ScrollArea>
</template>
