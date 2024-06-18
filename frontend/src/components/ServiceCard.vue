<script setup lang="ts">
import {Button} from '@/components/ui/button'
defineProps<{
    src: string,
    src2: string | undefined,
    title: string,
    description: string    
}>()
</script>

<template>
    <div class="flex flex-col rounded-xl border overflow-hidden w-[360px] shadow-xl">
        <img v-if="!src2" :src="src" alt="" class="w-full h-96 object-cover">
        <div v-else class="flex">
            <img class="w-1/2 h-96 object-cover" :src="src" alt="">
            <img class="w-1/2 h-96 object-cover" :src="src2" alt="">
        </div>
        <div class="pb-3 px-2">
            <!-- Title -->
            <p class="text-center text-xl mt-2 font-semibold">{{ title }}</p>
            <!-- Description -->
            <p class="text-center">{{ description }}</p>
        </div>

        <Button class="w-fit m-auto mb-2">Learn more</Button>
    </div>
</template>

<style scoped>
</style>