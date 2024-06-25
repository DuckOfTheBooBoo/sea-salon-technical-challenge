<script setup lang="ts">
import { GoogleMap, Marker, MarkerCluster } from "vue3-google-map";
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
          lat: event.latLng.lat(),
          lng: event.latLng.lng(),
        };
        const newBranches = props.branches;
        newBranches.push(coordinate);
        emit("update:branches", newBranches);
        emit("update:location", coordinate);
        removeListener();
      }
    );
}

const removeListener = () => {
  if (mapRef.value?.ready && clickListener) {
    mapRef.value.api.event.removeListener(clickListener);
    clickListener = null;
  }
}

watch(() => props.registerClickListener, (newVal: boolean) => {
  if (!newVal) return;
  addClickListener();
})
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
      <Marker
        v-for="(branch, i) in branches"
        :key="i"
        :options="{ position: branch }"
      />
    </MarkerCluster>
  </GoogleMap>
</template>
