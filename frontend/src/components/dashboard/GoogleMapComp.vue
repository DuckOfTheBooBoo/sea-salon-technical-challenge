<script setup lang="ts">
import { GoogleMap, MarkerCluster, CustomMarker } from "vue3-google-map";
import { Store } from "lucide-vue-next";
import { Ref, ref, watch } from "vue";
import { type Coordinate, type Branch } from "@/types";
import { useToast } from "@/components/ui/toast/use-toast";
const mapRef: Ref<any> = ref(null);
const props = defineProps<{
  branches: Branch[];
  location: Coordinate;
  registerClickListener: boolean;
  center: Coordinate;
  branchName: string;
}>();

const zoom: Ref<number> = ref(12);
let clickListener: any = null;
const API_KEY: string = 'AIzaSyBlqsn0PcSVEckwY85-9ZxCXK-tJEHUm7c'; // Hapus setelah pengumuman tahap 1, bismillah :D
const emit = defineEmits<{
  (e: "update:location", location: Coordinate): void;
}>();
const { toast } = useToast();


const addClickListener = () => {
  if (!mapRef.value?.ready) return;
  clickListener = mapRef.value.api.event.addListener(
    mapRef.value.map,
    "click",
    (event: any) => {
      const coordinate: Coordinate = {
        lat: Number(event.latLng.lat()),
        lng: Number(event.latLng.lng()),
      };
      emit("update:location", coordinate);
      removeListener();
    }
  );
};

const removeListener = () => {
  if (mapRef.value?.ready && clickListener) {
    mapRef.value.api.event.removeListener(clickListener);
    clickListener = null;
  }
};

watch(
  () => props.registerClickListener,
  (newVal: boolean) => {
    if (!newVal) return;
    toast({
      title: "Drop Pin",
      description: "Click on the map to drop a pin",
    });
    addClickListener();
  }
);

watch(
  () => props.center,
  () => {
    zoom.value = 17;
  }
);
</script>

<template>
  <GoogleMap
    ref="mapRef"
    :api-key="API_KEY"
    :center="center"
    :zoom="12"
    class="h-full"
    :fullscreen-control="false"
    :clickable-icons="false"
    :street-view-control="false"
    :map-type-control="false"
  >
    <MarkerCluster>
      <CustomMarker
        v-for="branch in branches"
        :key="branch.lat + branch.lng"
        class="hover:cursor-pointer"
        :options="{ position: branch }"
      >
        <div class="flex flex-col justify-center gap-2 max-w-[20rem]">
          <div class="text-lg text-center font-semibold bg-white rounded-lg px-2">
            {{ branch.branch_name }}
          </div>
          <Store class="w-full h-6 text-center" />
        </div>
      </CustomMarker>

      <CustomMarker
        v-if="location"
        :options="{ position: location, anchorPoint: 'BOTTOM_CENTER' }"
      >
        <div class="flex flex-col justify-center gap-2 max-w-[20rem]">
          <div
            class="text-lg text-center font-semibold bg-white rounded-lg px-2"
          >
            {{ branchName }}
          </div>
          <Store class="w-full h-6 text-center" />
        </div>
      </CustomMarker>
    </MarkerCluster>
  </GoogleMap>
</template>
