<script setup lang="ts">
import { ScrollArea } from "@/components/ui/scroll-area";
import { ref, watch, Ref } from "vue";
import { Separator } from "@/components/ui/separator";
import { type Branch as BranchType, type Coordinate} from "@/types";
import Branch from "@/components/dashboard/Branch.vue";
import { Toggle } from "@/components/ui/toggle";

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

const compareBranch = (branch: BranchType): boolean =>  {
  return selectedBranch.value === branch
}
</script>

<template>
  <div class="w-full">
    <p class="text-lg font-semibold">Branches</p>
    <Separator />

    <ScrollArea class="rounded-md px-2 py-1 h-[520px]">
      <Toggle
        v-for="branch in branches"
        :key="`${branch.lat} + ${branch.lng}`"
        class="flex gap-2 my-2 hover:cursor-pointer"
        v-bind:pressed="compareBranch(branch)"
        @click="selectedBranch = branch" 
      >
        <Branch :branch="branch"/>
      </Toggle>
    </ScrollArea>
  </div>
</template>
