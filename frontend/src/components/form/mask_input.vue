<template>
   <div class="flex flex-col gap-1">
      <FloatLabel variant="in">
         <InputMask
            :id="name"
            v-model="value_facade"
            :name="name"
            type="text"
            fluid
            :mask="mask"
            :slot-char="slot_char"
            @input="handle_input($event)"
         />
         <label :for="name">{{ $t(label + "." + name) }}</label>
      </FloatLabel>
   </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"

const props = defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: string
   name: string
   label: string
   readonly?: boolean
   mask?: string
   slot_char?: string
}>()
const emit = defineEmits(["update:modelValue"])

const value_facade = computed({
   get: () => props.modelValue,
   set: (value) => emit("update:modelValue", value),
})

const handle_input = (event: any) => {
   emit("update:modelValue", event.target!.value)
}
</script>
