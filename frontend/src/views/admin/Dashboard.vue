<script setup lang="ts">
import { ref, Ref, onMounted } from "vue";
import { type Branch, type Coordinate } from "@/types";
import { Button } from "@/components/ui/button";
import BranchList from "@/components/dashboard/BranchList.vue";
import GoogleMap from "@/components/dashboard/GoogleMapComp.vue";
import BranchForm from "@/components/dashboard/BranchForm.vue";
import { getBranches } from "@/service/branchApi";
import { Plus } from "lucide-vue-next";
import { toast } from "@/components/ui/toast";

const branches: Ref<Branch[]> = ref([]);
const currentView: Ref<string> = ref("branches");
const center: Ref<Coordinate> = ref({ lat: -6.310533833522174, lng: 106.80416367147718 });
const location: Ref<Coordinate> = ref({} as Coordinate);
const branchName: Ref<string> = ref("");

const onClick = () => {
  currentView.value = currentView.value === "branches" ? "form" : "branches";
};

const handleViewChange = (newView: string) => {
  currentView.value = newView;
};

const handleNewBranch = (newBranch: Branch) => {
  branches.value.push(newBranch)
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
      <h1>SEA Salon Admin Dashboard</h1>
    </div>
    <div class="flex w-full h-full">
      <div class="w-1/3 px-3 max-h-full">
        <div
          class="w-full h-full flex flex-col justify-between pb-2"
          v-if="currentView === 'branches'"
        >
          <BranchList :branches="branches" @branch:focus="handleCenterChange" />
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
          @reset:location="location = {} as Coordinate"
        />
      </div>
      <div class="w-full h-full">
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
