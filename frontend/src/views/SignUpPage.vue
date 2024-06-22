<script setup lang="ts">
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  FormDescription
} from "@/components/ui/form";
import {
  Card,
  CardDescription,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";
import {SignUpRequest, SignUpResponse} from "@/types";
import { addNewUser } from "@/service/authApi";
import { useToast } from "@/components/ui/toast/use-toast";

const { toast } = useToast();

const signUpFormSchema = toTypedSchema(
  z.object({
    firstName: z.string().min(1, { message: "First name is required" }),
    lastName: z.string().optional(),
    email: z.string().email({ message: "Invalid email address" }),
    phoneNumber: z.string().min(1, { message: "Phone number is required" }),
    password: z
      .string()
      .min(8, { message: "Password must be at least 8 characters long" })
      .regex(/[a-z]/, {
        message: "Password must contain at least one lowercase letter",
      })
      .regex(/[A-Z]/, {
        message: "Password must contain at least one uppercase letter",
      })
      .regex(/[0-9]/, {
        message: "Password must contain at least one number",
      })
      .regex(/[@$!%*?&]/, {
        message: "Password must contain at least one special character",
      }),
    passwordConfirm: z.string().min(8, {
      message: "Confirm password must be at least 8 characters long",
    }),
  // Check if passwords match
  }).refine((data) => data.password === data.passwordConfirm, {
    message: "Passwords do not match",
    path: ["passwordConfirm"],

  // Check if phone number contains whitespace or hypens
  }).refine((data) => !/[ ^\s-]/.test(data.phoneNumber), {
    message: "Phone number cannot contain whitespace or hypens",
    path: ["phoneNumber"],
  })
);

const signUpForm = useForm({
  validationSchema: signUpFormSchema,
});

const onSignUpSubmit = signUpForm.handleSubmit(async (values) => {
  const request: SignUpRequest = {
    first_name: values.firstName,
    last_name: values.lastName,
    email: values.email,
    password: values.password,
    phone_number: values.phoneNumber,
  };
  console.log(request)
  try {
    const response: SignUpResponse = await addNewUser(request);

  } catch (error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message,
    })
  }
});
</script>

<template>
  <Toaster />

  <div class="w-screen h-screen flex justify-center items-center">
    <Card class="max-w-xl">
      <form @submit.prevent="onSignUpSubmit">
        <CardHeader>
          <CardTitle>Sign Up</CardTitle>
          <CardDescription
            >Create your account to become a member.</CardDescription
          >
        </CardHeader>
        <CardContent>
          <div class="flex gap-4">
            <FormField
              v-slot="{ componentField }"
              name="firstName"
              class="w-1/2"
            >
              <FormItem>
                <FormLabel>First name</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Enter your first name here"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
            <FormField
              v-slot="{ componentField }"
              name="lastName"
              class="w-1/2"
            >
              <FormItem>
                <FormLabel>Last name</FormLabel>
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Enter your last name here"
                    v-bind="componentField"
                  />
                </FormControl>
                <FormMessage class="text-xs h-4" />
              </FormItem>
            </FormField>
          </div>
          <FormField v-slot="{ componentField }" name="email">
            <FormItem>
              <FormLabel>Your e-mail</FormLabel>
              <FormControl>
                <Input
                  type="email"
                  placeholder="Enter your e-mail address here"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="phoneNumber">
            <FormItem>
              <FormLabel>Your active phone number</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter your phone number here"
                  v-bind="componentField"
                />
              </FormControl>
              <FormDescription>e.g +62 812 3456 7890</FormDescription>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="password">
            <FormItem>
              <FormLabel>Your password</FormLabel>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Enter your password here"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="passwordConfirm">
            <FormItem>
              <FormLabel>Confirm your password</FormLabel>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Re-enter your password here"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
        </CardContent>
        <CardFooter>
          <Button type="submit">Submit</Button>
        </CardFooter>
      </form>
    </Card>
  </div>
</template>
