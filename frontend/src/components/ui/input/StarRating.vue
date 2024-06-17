<script setup lang="ts">
import { Ref, ref, onBeforeMount, watchEffect, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import { useField, FieldContext } from 'vee-validate'

const props = defineProps<{
  modelValue: number,
  isWriteable: boolean
}>()

type ValueType = number

const stars: number[] = [1,2,3,4,5]

const ratingStar: Ref<number> = ref(0)

const { value: rating, validate }: FieldContext<ValueType> = useField('rating', {
  required: true,
  min: 1,
  max: 5,
})

const setRating = (stars: number): void => {
  ratingStar.value = stars
  rating.value = stars
  validate()
}

onBeforeMount(() => {
  console.log(props.modelValue)
  setRating(props.modelValue)
})

watchEffect(() => {
  ratingStar.value = props.modelValue
})

onMounted(() => {
  console.log('REF: ', ratingStar.value)
  console.log('MODEL: ', props.modelValue)
  console.log('FORM: ', rating.value)
})
</script>

<template>
  <div class="flex gap-1 hover:cursor-pointer w-fit">
    <Icon 
      :class="['text-2xl', star <= rating ? 'text-yellow-400' : '']"
      v-for="star in stars"
      :key="star"
      :icon="star <= rating ? 'radix-icons:star-filled' : 'radix-icons:star'"
      @click="isWriteable ? setRating(star) : null"
    />
  </div>
</template>

<style scoped></style>
