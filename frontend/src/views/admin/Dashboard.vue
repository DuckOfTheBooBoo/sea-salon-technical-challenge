<script setup lang="ts">
import { ref, Ref, onMounted } from "vue";
import { type Branch, type Coordinate } from "@/types";
import { Button } from "@/components/ui/button";
import BranchList from "@/components/dashboard/BranchList.vue";
import GoogleMap from "@/components/dashboard/GoogleMapComp.vue";
import BranchForm from "@/components/dashboard/BranchForm.vue";
import { getBranches } from "@/service/branchApi";
import { Plus, Info } from "lucide-vue-next";
import { toast } from "@/components/ui/toast";

const branches: Ref<Branch[]> = ref([]);
const currentView: Ref<string> = ref("branches");
const center: Ref<Coordinate> = ref({ lat: -6.310533833522174, lng: 106.80416367147718 });
const location: Ref<Coordinate> = ref({} as Coordinate);
const branchName: Ref<string> = ref("");
const tempBranch: Ref<Branch> = ref({} as Branch);

const onClick = () => {
  currentView.value = currentView.value === "branches" ? "form" : "branches";
};

const handleBranchEdit = (branch: Branch) => {
  currentView.value = "form";
  tempBranch.value = branch;
}

const handleViewChange = (newView: string) => {
  currentView.value = newView;
  // Prune Branch form fields populated by editing an existing branch
  tempBranch.value = {} as Branch;
};

const handleNewBranch = (newBranch: Branch, oldBranch?: Branch) => {
  if (oldBranch) {
    const index = branches.value.findIndex((b) => b.branch_name === oldBranch.branch_name);
    branches.value[index] = newBranch;
    return
  }

  branches.value.push(newBranch)
}

const handleBranchDelete = (branch: Branch) => {
  branches.value = branches.value.filter((b) => b.branch_name !== branch.branch_name);
}

const handleNameChange = (newName: string) => {
  branchName.value = newName;
};

const handleCenterChange = (newCenter: Coordinate) => {
  center.value = newCenter;
};

const registerClickListener: Ref<boolean> = ref(false);

const enableDropPin = () => {
  registerClickListener.value = true;
};

const handleLocationChange = (newLocation: Coordinate) => {
  location.value = newLocation;
  registerClickListener.value = false;
};

onMounted(async () => {
  try {
    branches.value = await getBranches();
  } catch(error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message
    })
  }
})
</script>

<template>
  <div class="flex flex-col h-screen w-screen">
    <div class="py-2 px-3">
      <h1><span class="font-oleragie font-semibold text-[0.5rem]">SEA Salon</span> Admin Dashboard</h1>
    </div>
    <div class="flex w-full h-full flex-col-reverse sm:flex-row">
      <div class="w-full h-1/2 sm:w-1/3 px-3 max-h-full">
        <div
          class="w-full h-full flex flex-col justify-between pb-2"
          v-if="currentView === 'branches'"
        >
          <BranchList 
            :branches="branches" 
            @branch:focus="handleCenterChange" 
            @branch:delete="handleBranchDelete"
            @branch:edit="handleBranchEdit" />
          <Button @click="onClick" class="w-full flex gap-3" variant="outline">
            <Plus />Add new branch
          </Button>
        </div>
        <BranchForm
          v-else
          :current-view="currentView"
          @update:view="handleViewChange"
          :location="location"
          @drop-pin="enableDropPin"
          @update:branch-name="handleNameChange"
          @update:branch="handleNewBranch"
          :branch-data="tempBranch"
          @reset:location="location = {} as Coordinate"
          />
      </div>
      <div class="w-full h-1/2 sm:w-full sm:h-full">
        <div class="absolute z-10 flex gap-2 rounded-xl py-2 px-2 my-1 mx-2 bg-blue-100">
          <Info />
          <p class="text-base font-semibold">Click branch to see details.</p>
        </div>
        <GoogleMap
          :branches="branches"
          :location="location"
          @update:location="handleLocationChange"
          :register-click-listener="registerClickListener"
          :branch-name="branchName ? branchName : ''"
          :center="center"
        />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
