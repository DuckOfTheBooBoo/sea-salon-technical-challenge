<script setup lang="ts">
import { GoogleMap, Marker, MarkerCluster } from "vue3-google-map";
import {Ref, ref, watch} from 'vue'
import {type Coordinate} from '@/types'
import {cn} from '@/lib/utils'

const mapRef: Ref<any> = ref(null);
defineProps<{branches: Coordinate[]}>();
const API_KEY: string = import.meta.env.VITE_GOOGLE_MAPS_API_KEY;
const center: Coordinate = { lat: 40.689247, lng: -74.044502 };

// watch(
//   () => mapRef.value?.ready,
//   (ready: boolean) => {
//     if (!ready) return;

//     mapRef.value.api.event.addListener(
//       mapRef.value.map,
//       "click",
//       (event: any) => {
//         branches.value.push({
//           lat: event.latLng.lat(),
//           lng: event.latLng.lng(),
//         });
//       }
//     );
//   }
// );
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
              :options="{ position: location }"
            />
          </MarkerCluster>
        </GoogleMap>
</template>