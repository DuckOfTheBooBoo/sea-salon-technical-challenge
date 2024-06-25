<script setup lang="ts">
import { ScrollArea } from "@/components/ui/scroll-area";
import { ref, watch, Ref } from "vue";
import { Separator } from "@/components/ui/separator";
import {type Coordinate} from '@/types';

defineProps<{branches: Coordinate[]}>();
</script>

<template>
    <div class="w-full">
        <p class="text-lg font-semibold">Branches</p>
        <Separator />
      
        <ScrollArea class="rounded-md px-2 py-1 h-[520px]">
          <p v-for="i in 100" :key="i">Branch {{ i }}</p>
        </ScrollArea>
    </div>
</template>
