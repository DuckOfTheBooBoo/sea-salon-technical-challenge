<script setup lang="ts">
import {
  ComboboxAnchor,
  ComboboxInput,
  ComboboxPortal,
  ComboboxRoot,
} from "radix-vue";
import {
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  TagsInput,
  TagsInputInput,
  TagsInputItem,
  TagsInputItemDelete,
  TagsInputItemText,
} from "@/components/ui/tags-input";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import * as z from "zod";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { useToast } from "../ui/toast";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { ref, type Ref, computed, watch, onMounted } from "vue";
import { type Service } from "@/types";
import { addService, getServices } from "@/service/serviceApi";
import { Plus } from "lucide-vue-next";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

const { toast } = useToast();
const isServiceFormOpen: Ref<boolean> = ref(false);
const services: Ref<Service[]> = ref([]);

const serviceFormSchema = toTypedSchema(
  z.object({
    serviceName: z.string().min(1, { message: "Service name is required" }),
    duration: z.number().min(1, { message: "Duration is required" }),
  })
);

const serviceForm = useForm({
  validationSchema: serviceFormSchema,
});

const emit = defineEmits<{
  (e: "update:services", value: string[]): void;
}>();

const props = defineProps<{
  services?: Service[];
}>();

const modelValue = ref<string[]>([]);
const open: Ref<boolean> = ref(false);
const searchTerm: Ref<string> = ref("");

const filteredServices = computed(() =>
  services.value.filter((i) => !modelValue.value.includes(i.service_name))
);

watch(modelValue.value, () => {
  emit("update:services", modelValue.value);
});

const onSubmit = serviceForm.handleSubmit(async (values) => {
  const request: Service = {
    service_name: values.serviceName,
    service_code: values.serviceName.replace(" ", "-").toLowerCase(),
    duration: values.duration,
  };

  try {
    const resp = await addService(request);
    serviceForm.resetForm();
    isServiceFormOpen.value = false;
    toast({
      title: "Success!",
      description: "Successfully added new service.",
    });
    services.value.push(resp);
  } catch (error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message,
    });
  }
});

onMounted(async () => {
  try {
    const resp = await getServices();
    services.value = resp;
  } catch (error: unknown) {
    const err: Error = error as Error;
    toast({
      title: "Something went wrong",
      description: err.message,
    });
  }
});

onMounted(() => {
  if (props.services) {
    modelValue.value = props.services.map((s) => s.service_name);
  }
});

const handleBlur = () => {
  open.value = false;
};
</script>
<template>
  <div class="flex gap-2 mb-5">
    <TagsInput class="px-0 gap-0 w-80" :model-value="modelValue">
      <div class="flex gap-2 flex-wrap items-center px-3">
        <TagsInputItem v-for="item in modelValue" :key="item" :value="item">
          <TagsInputItemText />
          <TagsInputItemDelete />
        </TagsInputItem>
      </div>

      <ComboboxRoot
        v-model="modelValue"
        v-model:open="open"
        v-model:searchTerm="searchTerm"
        @click="open = true"
        class="w-full"
      >
        <ComboboxAnchor as-child>
          <ComboboxInput placeholder="Services..." as-child>
            <TagsInputInput
              class="w-full px-3"
              :class="modelValue.length > 0 ? 'mt-2' : ''"
              @blur="handleBlur"
              @keydown.enter.prevent
            />
          </ComboboxInput>
        </ComboboxAnchor>

        <ComboboxPortal>
          <CommandList
            position="popper"
            class="w-[--radix-popper-anchor-width] rounded-md mt-2 border bg-popover text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2"
          >
            <CommandEmpty />
            <CommandGroup>
              <CommandItem
                v-for="service in filteredServices"
                :key="service.service_code"
                :value="service.service_name"
                @select.prevent="
                  (ev) => {
                    if (typeof ev.detail.value === 'string') {
                      searchTerm = '';
                      modelValue.push(ev.detail.value);
                      console.log(modelValue);
                    }

                    if (filteredServices.length === 0) {
                      open = false;
                    }
                  }
                "
              >
                {{ service.service_name }}
              </CommandItem>
            </CommandGroup>
          </CommandList>
        </ComboboxPortal>
      </ComboboxRoot>
    </TagsInput>

    <Dialog v-model:open="isServiceFormOpen">
      <DialogTrigger as-child>
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger as-child>
              <Button variant="outline" class="h-full" @click="isServiceFormOpen = true"><Plus /></Button>
            </TooltipTrigger>
            <TooltipContent>
              <p>Add new service</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Service Form</DialogTitle>
        </DialogHeader>
        <form @submit.prevent="onSubmit">
          <FormField v-slot="{ componentField }" name="serviceName">
            <FormItem>
              <FormLabel>Service name</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter the new service name"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <FormField v-slot="{ componentField }" name="duration">
            <FormItem>
              <FormLabel>Duration</FormLabel>
              <FormControl>
                <Input
                  type="number"
                  placeholder="Enter the new service duration"
                  v-bind="componentField"
                />
              </FormControl>
              <FormDescription> In minutes e.g. 60 = 1 hour </FormDescription>
              <FormMessage class="text-xs h-4" />
            </FormItem>
          </FormField>
          <DialogFooter>
            <Button type="submit">Submit</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  </div>
</template>
