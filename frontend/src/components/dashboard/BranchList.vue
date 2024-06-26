<script setup lang="ts">
import { ScrollArea } from "@/components/ui/scroll-area";
import { ref, watch, Ref } from "vue";
import { Separator } from "@/components/ui/separator";
import { type Branch as BranchType, type Coordinate } from "@/types";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Button } from "@/components/ui/button"


defineProps<{ branches: BranchType[] }>();
const emit = defineEmits<{
  (e: "branch:focus", coordinate: Coordinate): void;
}>();

const selectedBranch: Ref<BranchType> = ref({} as BranchType);

watch(
  () => selectedBranch.value,
  () => {
    console.log(selectedBranch.value);
    emit("branch:focus", {
      lat: selectedBranch.value.lat,
      lng: selectedBranch.value.lng,
    } as Coordinate);
  }
);
</script>

<template>
  <div class="w-full">
    <p class="text-lg font-semibold">Branches</p>
    <Separator />

    <ScrollArea class="rounded-md px-2 py-1 h-[520px]">
      <Accordion
        type="single"
        class="w-full"
        collapsible
      >
        <AccordionItem
        v-for="branch in branches"
        :key="`${branch.lat} + ${branch.lng}`"
        :value="branch.branch_name"
        @click="selectedBranch = branch"
        class="w-full"
        >
          <AccordionTrigger>
            <p class="text-base">{{ branch.branch_name }}</p>
          </AccordionTrigger>
          <AccordionContent class="">
            {{ branch.branch_address }}
            <div class="flex w-full gap-2">
              <Button class="w-1/2">Edit</Button>
              <Button class="w-1/2">Delete</Button>
            </div>
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </ScrollArea>
  </div>
</template>
