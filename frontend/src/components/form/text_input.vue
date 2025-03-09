<template>
   <div class="flex flex-col gap-1">
      <FloatLabel variant="in">
         <InputText
            :id="name"
            v-model="value_facade"
            :name="name"
            type="text"
            fluid
            @input="handle_input($event)"
         />
         <label :for="name">{{ $t(label + "." + name) }}</label>
      </FloatLabel>
      <Message v-if="error">{{ error }}</Message>
   </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"

// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ modelValue: string; name: string; error: string; label: string }>()
const emit = defineEmits(["update:modelValue"])

const value_facade = computed({
   get: () => props.modelValue,
   set: (value) => emit("update:modelValue", value),
})

const handle_input = (event: any) => {
   console.log(event)
   emit("update:modelValue", event.target!.value)
}
</script>
