<script setup lang="ts">
import { Calendar } from "@/components/ui/calendar";
import ReservationList from "@/components/dashboard/ReservationList.vue";
import { ref, type Ref, onMounted } from "vue";
import {type Branch} from "@/types";
import { getBranches } from "@/service/branchApi";
import { useToast } from "@/components/ui/toast/use-toast";

const branches: Ref<Branch[]> = ref([]);

const { toast } = useToast()

onMounted(async () => {
  try {
    branches.value = await getBranches();
  } catch (error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message,
    });
  }
});
</script>

<template>
  <div class="flex w-screen h-screen gap-2">
    <ReservationList :branches="branches" />
  </div>
</template>
