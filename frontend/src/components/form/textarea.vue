<template>
   <div class="flex flex-col gap-1">
      <FloatLabel variant="in">
         <InputText
            :id="name"
            :value="modelValue"
            :name="name"
            type="text"
            fluid
            @input="emit('update:modelValue', $event)"
         />
         <label :for="name">{{ $t("iam.pg.update.name") }}</label>
      </FloatLabel>
      <Message v-if="error">{{ error }}</Message>
   </div>
</template>

<script lang="ts" setup>
// eslint-disable-next-line vue/prop-name-casing
defineProps<{ modelValue: string; name: string; error: string }>()
const emit = defineEmits(["update:modelValue"])
</script>
