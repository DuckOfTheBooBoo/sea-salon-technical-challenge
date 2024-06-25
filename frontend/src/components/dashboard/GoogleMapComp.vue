<script setup lang="ts">
import {
  GoogleMap,
  Marker,
  MarkerCluster,
  CustomMarker,
} from "vue3-google-map";
import { Store } from "lucide-vue-next";
import { Ref, ref, watch } from "vue";
import { type Coordinate } from "@/types";

const mapRef: Ref<any> = ref(null);
const props = defineProps<{
  branches: Coordinate[];
  location: Coordinate;
  registerClickListener: boolean;
}>();
let clickListener: any = null;
const API_KEY: string = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
const center: Coordinate = { lat: 40.689247, lng: -74.044502 };
const emit = defineEmits<{
  (e: "update:location", location: Coordinate): void;
  (e: "update:branches", branches: Coordinate[]): void;
}>();

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
      // const newBranches = props.branches;
      // newBranches.push(coordinate);
      // emit("update:branches", newBranches);
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
    addClickListener();
  }
);
</script>

<template>
  <GoogleMap
    ref="mapRef"
    :api-key="API_KEY"
    :center="center"
    :zoom="15"
    class="h-full"
    :fullscreen-control="false"
    :clickable-icons="false"
    :street-view-control="false"
    :map-type-control="false"
  >
    <MarkerCluster>
      <CustomMarker
        v-for="(branch, i) in branches"
        :key="branch.lat + branch.lng"
        :options="{ position: branch }"
      >
        <div class="text-center">
          <div class="text-lg text-center font-semibold">Branch {{ i }}</div>
          <Store class="w-6 h-6" />
        </div>
      </CustomMarker>
      <CustomMarker v-if="location" :options="{ position: location }">
        <div class="flex flex-col justify-center">
          <div class="text-lg text-center font-semibold">Branch</div>
          <Store class="w-full h-6 text-center" />
        </div>
      </CustomMarker>
    </MarkerCluster>
  </GoogleMap>
</template>
