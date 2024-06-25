<script setup lang="ts">
import { ref, Ref, watch } from "vue";
import { type Coordinate } from "@/types";
import { Button } from "@/components/ui/button";
import BranchList from "@/components/dashboard/BranchList.vue";
import GoogleMap from "@/components/dashboard/GoogleMapComp.vue";
import BranchForm from "@/components/dashboard/BranchForm.vue";
import { Plus } from "lucide-vue-next";

const branches: Ref<Coordinate[]> = ref([]);
const currentView: Ref<string> = ref("branches");

const location: Ref<Coordinate> = ref({} as Coordinate);
const branchName: Ref<string> = ref("");

const onClick = () => {
  currentView.value = currentView.value === "branches" ? "form" : "branches";
};

const handleViewChange = (newView: string) => {
  currentView.value = newView;
};

const handleBranchChange = (newBranches: Coordinate[]) => {
  branches.value = newBranches;
};

const handleNameChange = (newName: string) => {
  branchName.value = newName;
};

const registerClickListener: Ref<boolean> = ref(false);

const enableDropPin = () => {
  registerClickListener.value = true;
};

const handleLocationChange = (newLocation: Coordinate) => {
  location.value = newLocation;
  registerClickListener.value = false;
};
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
          <BranchList :branches="branches" />
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
        />
      </div>
      <div class="w-full h-full">
        <GoogleMap
          :branches="branches"
          @update:branches="handleBranchChange"
          :location="location"
          @update:location="handleLocationChange"
          :register-click-listener="registerClickListener"
          :branch-name="branchName ? branchName : ''"
        />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
