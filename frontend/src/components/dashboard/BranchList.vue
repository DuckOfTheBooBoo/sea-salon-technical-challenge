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
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import {deleteBranch} from '@/service/branchApi';
import { useToast } from "../ui/toast";
import { Badge } from "@/components/ui/badge"

defineProps<{ branches: BranchType[] }>();
const emit = defineEmits<{
  (e: "branch:focus", coordinate: Coordinate): void;
  (e: "branch:delete", branch: BranchType): void;
  (e: "branch:edit", branch: BranchType): void;
}>();

const {toast} = useToast()
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

const deleteBranchMethod = async (branch: BranchType) => {
  try {
    await deleteBranch(branch);
    emit("branch:delete", branch);
    toast({
      title: "Success",
      description: `${branch.branch_name} deleted successfully`
    })
  } catch(error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message
    })
  }
}
</script>

<template>
  <div class="w-full">
    <p class="text-lg font-semibold">Branches</p>
    <Separator />

    <ScrollArea class="rounded-md px-2 py-1 h-full sm:h-[520px]">
      <Accordion type="single" class="w-full" collapsible>
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
            <p class="mb-2">{{ branch.branch_address }}</p>
            <div>
              Open at: {{ branch.open_time }} Close at: {{ branch.close_time }}
            </div>
            <div class="flex gap-2 flex-wrap my-1">
              <Badge v-for="service in branch.services">{{ service.service_name }}</Badge>
            </div>
            <div class="flex w-full gap-2">
              <Button class="w-1/2 text-sm my-1" variant="outline" @click="emit('branch:edit', branch)">Edit</Button>

              <AlertDialog>
                <AlertDialogTrigger as-child>
                  <Button class="w-1/2 text-sm my-1" variant="destructive">Delete</Button>
                </AlertDialogTrigger>
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle
                      >Are you absolutely sure?</AlertDialogTitle
                    >
                    <AlertDialogDescription>
                      This action cannot be undone. This will permanently delete
                      {{ branch.branch_name }} and remove the data from our database.
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <AlertDialogFooter>
                    <AlertDialogCancel>Cancel</AlertDialogCancel>
                    <AlertDialogAction @click="deleteBranchMethod(branch)" class="bg-destructive">Continue</AlertDialogAction>
                  </AlertDialogFooter>
                </AlertDialogContent>
              </AlertDialog>
            </div>
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </ScrollArea>
  </div>
</template>
