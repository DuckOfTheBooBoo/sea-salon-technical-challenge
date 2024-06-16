<script setup lang="ts">
defineProps<{
    src: string,
    src2: string | undefined,
    title: string,
    description: string    
}>()
</script>

<template>
    <div class="flex flex-col rounded-xl overflow-hidden w-[360px] shadow-xl">
        <img v-if="!src2" :src="src" alt="" class="w-full h-96 object-cover">
        <div v-else class="flex">
            <img class="w-1/2 h-96 object-cover" :src="src" alt="">
            <img class="w-1/2 h-96 object-cover" :src="src2" alt="">
        </div>
        <div class="pb-3">
            <!-- Title -->
            <p class="text-center text-xl mt-2">{{ title }}</p>
            <!-- Description -->
            <p class="text-center">{{ description }}</p>
        </div>
    </div>
</template>

<style scoped>
</style>