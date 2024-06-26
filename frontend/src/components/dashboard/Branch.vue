<script setup lang="ts">
import {type Branch} from '@/types';
import { Store } from "lucide-vue-next";

defineProps<{branch: Branch}>();
</script>

<template>
    <Store class="w-fit h-fit" />
    <p class="text-base">{{ branch.branch_name }}</p>
</template>