<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { ArrowLeft } from "lucide-vue-next";
import {useRouter} from "vue-router";

const router = useRouter();
</script>

<template>
  <div class="flex w-screen h-screen flex-col justify-center items-center">
    <span class="font-semibold text-lg">403</span>
    <h1 class="font-semibold text-5xl my-5">Forbidden</h1>
    <h2 class="text-xl text-gray-500 mb-5">Sorry, but you don't have permission to access this page.</h2>
    <Button class="w-[10rem]" @click="router.go(-1)">
      <ArrowLeft />
      Go back
    </Button>
  </div>
</template>
