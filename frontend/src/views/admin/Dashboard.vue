<script setup lang="ts">
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "@/components/ui/resizable";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Toggle } from "@/components/ui/toggle";
import { GoogleMap, Marker, MarkerCluster } from "vue3-google-map";
import { ScrollArea } from "@/components/ui/scroll-area";
import { ref, Ref, watch } from "vue";

interface Location {
  lat: number;
  lng: number;
}

const API_KEY: string = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
const mapRef: Ref<any> = ref(null);

const locations: Ref<Location[]> = ref([]);

watch(() => mapRef.value?.ready, (ready: boolean) => {
  if (!ready) return

  mapRef.value.api.event.addListener(mapRef.value.map, "click", (event: any) => {
    locations.value.push({ lat: event.latLng.lat(), lng: event.latLng.lng() })
    console.log(event)
  });
});

const center = { lat: 40.689247, lng: -74.044502 };
</script>

<template>
  <div class="flex flex-col h-screen w-screen">
    <div class="py-2 px-3">
      <h1>SEA Salon Admin Dashboard</h1>
      <Separator />
    </div>
    <div class="flex w-full h-full">
      <div class="w-1/4 px-3">
        <p>Branches</p>
        <Separator />

        <ScrollArea class="rounded-md p-4 h-[300px]">
          <p v-for="i in 100" :key="i">Branch {{ i }}</p>
        </ScrollArea>

        <Button class="w-full">Add new branch +</Button>
      </div>
      <div class="w-full h-full">
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
              v-for="(location, i) in locations"
              :key="i"
              :options="{ position: location }"
            />
          </MarkerCluster>
        </GoogleMap>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
