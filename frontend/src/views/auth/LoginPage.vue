<script setup lang="ts">
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
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
import {LogInRequest, LogInResponse} from "@/types";
import { logIn } from "@/service/authApi";
import { useToast } from "@/components/ui/toast/use-toast";
import { useRouter } from "vue-router";
import { AuthError } from "@/errors/auth";
import { HTTP_UNAUTHORIZED, HTTP_NOT_FOUND } from "@/constants";

const { toast } = useToast();
const router = useRouter();

const logInFormSchema = toTypedSchema(
  z.object({
    email: z.string().email({ message: "Invalid email address" }),
    password: z
      .string()
      .min(1, { message: "Password is required" })
  })
);

const logInForm = useForm({
  validationSchema: logInFormSchema,
});

const onLogInSubmit = logInForm.handleSubmit(async (values) => {
  const request: LogInRequest = {
    email: values.email,
    password: values.password,
  };
  try {
    const response: LogInResponse = await logIn(request);
    localStorage.setItem("token", response.token);
    router.push({path: response.redirect_path})
  } catch (error: unknown) {
    const err: AuthError = error as AuthError;
    console.log(err)
    if (err.code === HTTP_UNAUTHORIZED) {
      logInForm.setFieldError("password", err.message);
      logInForm.setFieldError("email", err.message);
    } else if (err.code === HTTP_NOT_FOUND) {
      logInForm.setFieldError("email", err.message);
    }
    toast({
      title: "Something went wrong",
      description: err.message,
    })
  }
});
</script>

<template>
  <div class="w-screen h-screen flex justify-center items-center">
    <Card class="max-w-xl">
      <form @submit.prevent="onLogInSubmit">
        <CardHeader>
          <CardTitle>Log In</CardTitle>
          <CardDescription
            >Enter your email and password to log in.</CardDescription
          >
        </CardHeader>
        <CardContent>
          <FormField v-slot="{ componentField }" name="email">
            <FormItem v-auto-animate>
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
          <FormField v-slot="{ componentField }" name="password">
            <FormItem v-auto-animate>
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
        </CardContent>
        <CardFooter class="flex flex-col">
          <Button type="submit" class="w-1/2">Login</Button>
          <RouterLink to="/signup" class="text-base my-2">Create a new account?</RouterLink>
        </CardFooter>
      </form>
    </Card>
  </div>
</template>
