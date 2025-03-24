<template>
   <div class="flex flex-col gap-1">
      <FloatLabel variant="in">
         <Textarea
            :id="name"
            v-model="value_facade"
            :name="name"
            :label="$t(label + '.' + name)"
            error=""
            fluid
            @on-update:model-value="emit('update:modelValue', $event)"
         />
         <label :for="name">{{ $t(label + "." + name) }}</label>
      </FloatLabel>
   </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"

// eslint-disable-next-line vue/prop-name-casing
const props = defineProps<{ modelValue: string; name: string; label: string }>()
const emit = defineEmits(["update:modelValue"])

const value_facade = computed({
   get: () => props.modelValue,
   set: (value) => emit("update:modelValue", value),
})
</script>
