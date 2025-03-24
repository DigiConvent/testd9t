<template>
   <div class="flex flex-col gap-1">
      <FloatLabel variant="in">
         <InputText
            :id="name"
            v-model="value_facade"
            :name="name"
            type="text"
            fluid
            :readonly="readonly"
            @input="handle_input($event)"
         />
         <label :for="name">{{ $t(label + "." + name) }}</label>
      </FloatLabel>
   </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"

// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ modelValue: string; name: string; label: string; readonly?: boolean }>()
const emit = defineEmits(["update:modelValue"])

const value_facade = computed({
   get: () => props.modelValue,
   set: (value) => emit("update:modelValue", value),
})

const handle_input = (event: any) => {
   emit("update:modelValue", event.target!.value)
}
</script>
