<template>
   <div class="flex flex-col gap-1" @click="toggle">
      <Button :severity="'secondary'" class="!py-2 !px-1" :disabled="readonly">
         <div
            class="w-full -m-1 rounded-md"
            :class="{ 'bg-gray-100': !modelValue, 'bg-white': modelValue }"
         >
            <Fa v-if="icon_on && modelValue" :icon="icon_on" />
            <Fa v-if="icon_off != undefined && !modelValue" :icon="icon_off" />
            {{ modelValue ? $t(label_on) : $t(label_off) }}
            <ProgressSpinner
               v-if="loading"
               class="inline"
               style="height: 10px; width: 10px"
               :class="{ hidden: !loading }"
               stroke-width="8"
            />
         </div>
      </Button>
   </div>
</template>

<script lang="ts" setup>
import { computed } from "vue"

const props = defineProps<{
   // eslint-disable-next-line vue/prop-name-casing
   modelValue: boolean
   label_on: string
   label_off: string
   readonly?: boolean
   loading?: boolean
   icon_on?: string
   icon_off?: string
}>()
const emit = defineEmits(["update:modelValue"])

const value_facade = computed<boolean>({
   get: () => props.modelValue,
   set: (value) => {
      if (!props.readonly) emit("update:modelValue", value)
   },
})
function toggle() {
   if (props.loading || props.readonly) return
   value_facade.value = !value_facade.value
}
</script>
