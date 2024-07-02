<script setup lang="ts">
import { Icon } from '@iconify/vue'
import { Separator } from '@/components/ui/separator'
import { Ref, ref } from 'vue'

defineProps<{
  name: string;
  rating: number;
  comment: string;
}>();

const stars: Ref<number[]> = ref([1,2,3,4,5])
</script>

<template>
  <div class="my-2">
    <div class="flex justify-between">
      <p class="text-lg font-semibold">{{ name }}</p>
      <div class="flex gap-1 hover:cursor-pointer w-fit">
          <Icon
            :class="['text-2xl', star <= rating ? 'text-yellow-400' : '']"
            v-for="star in stars"
            :key="star"
            :icon="star <= rating ? 'radix-icons:star-filled' : 'radix-icons:star'"
          />
      </div>
    </div>
    <p class="text-sm">{{ comment }}</p>
  </div>
  <Separator />
</template>

<style scoped></style>
