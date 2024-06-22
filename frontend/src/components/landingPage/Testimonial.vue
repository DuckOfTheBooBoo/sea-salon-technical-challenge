<script setup lang="ts">
import { ScrollArea } from "@/components/ui/scroll-area";
import {
  FormControl,
  // FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input, StarRating } from "@/components/ui/input";
import { addReview, getReviews } from '@/service/reviewApi';
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import {ref, Ref, onBeforeMount} from 'vue';
import { Button } from "@/components/ui/button";
import { Review as ReviewType } from "@/types";
import Review from "@/components/Review.vue";


const props = defineProps<{toast: Function}>();

const reviewFormSchema = toTypedSchema(
  z.object({
    name: z.string().min(1, 'Name is required'),
    rating: z.number().min(1).max(5),
    comment: z.string().min(1, 'Comment is required'),
  })
);

const reviewForm = useForm({
  validationSchema: reviewFormSchema
});

const reviews: Ref<ReviewType[]> = ref([]);

const onReviewSubmit = reviewForm.handleSubmit(async (values) => {
  const newReview: ReviewType = {
    full_name: values.name,
    rating: values.rating,
    comment: values.comment,
  };

  try {
    const response = await addReview(newReview);
    console.log(response)
    reviews.value.unshift(response);
    props.toast({
      title: "Success!",
      description: "Your review has been submitted",
    });
    reviewForm.resetForm();
  } catch (error: unknown) {
    const err: Error = error as Error;
    props.toast({
      title: "Something went wrong",
      description: err.message,
    });
  }
});

onBeforeMount(async () => {
  try {
    const response: ReviewType[] = await getReviews();
    reviews.value = response;
    reviews.value.reverse();
  } catch (error: unknown) {
    const err = error as Error;
    props.toast({
      title: "Something went wrong!",
      description: err.message,
    });
  }
});
</script>

<template>
  <div class="h-[530px] m-h-[530px]" id="reviews">
    <h1 class="text-center text-4xl my-10 font-semibold">
      What Others Said About Us
    </h1>
    <div class="flex px-32 w-full justify-between gap-4">
      <Card class="w-2/3">
        <form @submit.prevent="onReviewSubmit">
          <CardHeader>
            <CardTitle>Leave a review</CardTitle>
          </CardHeader>
          <CardContent>
            <FormField v-slot="{ componentField }" name="name">
              <FormItem>
                <FormLabel>Your name</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Enter your full name here"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
            <FormField v-slot="{ componentField }" name="rating">
              <FormItem>
                <FormLabel>Your rating</FormLabel>
                <FormControl>
                  <StarRating
                    v-bind="componentField"
                    :model-value="0"
                    :isWriteable="true"
                  />
                </FormControl>
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
            <FormField v-slot="{ componentField }" name="comment">
              <FormItem>
                <FormLabel>Comment</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Enter your comment"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage class="text-xs" />
              </FormItem>
            </FormField>
          </CardContent>
          <CardFooter>
            <Button type="submit">Submit</Button>
          </CardFooter>
        </form>
      </Card>
      <Card class="w-full">
        <CardHeader>
          <CardTitle>Testimonials</CardTitle>
        </CardHeader>
        <CardContent class="">
          <ScrollArea class="rounded-md p-4 h-[300px]">
            <Review
              v-for="(review, index) in reviews"
              :key="index"
              :name="review.full_name"
              :rating="review.rating"
              :comment="review.comment"
            />
          </ScrollArea>
        </CardContent>
      </Card>
    </div>
  </div>
</template>
