<template>
   <div class="flex flex-col gap-1" @click="toggle">
      <label :for="name">{{ $t(label + "." + name) }}</label>
      <ToggleSwitch
         :id="name"
         v-model="value_facade"
         :name="name"
         fluid
         :readonly="readonly"
         @input="toggle"
      />
   </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"

const props = defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: boolean
   name: string
   label: string
   readonly?: boolean
}>()
const emit = defineEmits(["update:modelValue"])

const value_facade = computed<boolean>({
   get: () => props.modelValue,
   set: (value) => emit("update:modelValue", value),
})
function toggle() {
   value_facade.value = !value_facade.value
}
</script>
