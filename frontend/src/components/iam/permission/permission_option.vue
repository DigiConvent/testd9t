<template>
   <div>
      <label
         class="select-none block text-2xl"
         :class="`${is_checked && !is_indeterminate ? 'text-green-700' : is_indeterminate ? 'text-orange-600' : 'text-gray-500'} ${multiple && hovered ? 'bg-blue-100' : ''}`"
         @mouseenter="hovered = true"
         @mouseleave="hovered = false"
      >
         <template v-if="props.multiple">
            <span
               v-if="(is_checked && !is_indeterminate) || parent_hovered"
               class="pi pi-check-circle"
               :class="`${parent_hovered && !is_checked ? 'text-gray-500' : 'text-green-700'}`"
            ></span>
            <span v-else-if="is_indeterminate" class="pi pi-minus-circle text-orange-600"></span>
            <span v-else class="pi pi-circle text-slate-400"></span>
         </template>
         <input
            class="hidden"
            type="checkbox"
            :checked="is_checked"
            :indeterminate.prop="is_indeterminate"
            @change="handle_change"
         />
         {{ node.label }}
      </label>
      <div v-if="node.children" class="pl-4">
         <PermissionOption
            v-for="(child, index) in node.children"
            :key="index"
            :parent_hovered="parent_hovered || hovered"
            :node="child"
            :multiple="props.multiple"
            @update:checked="update_child_checked(index, $event)"
            @selected="emit('selected', $event)"
         />
      </div>
   </div>
</template>

<script setup lang="ts">
import type { CustomTreeNode } from "@/api/iam/permission/types"
import { computed, ref, watch } from "vue"
import PermissionOption from "./permission_option.vue"

const props = defineProps<{
   node: CustomTreeNode
   parent_hovered: boolean
   multiple: boolean
}>()
const hovered = ref(false)

const emit = defineEmits<{
   (event: "update:checked", value: boolean | null): void
   (event: "selected", value: CustomTreeNode): void
}>()

const node = ref(props.node)

const is_checked = computed(() => node.value.checked === true)

const is_indeterminate = computed(() => {
   if (!node.value.children) return false
   return node.value.partially_checked() && !node.value.fully_checked()
})

const update_node_and_children = (node: CustomTreeNode, is_checked: boolean | null) => {
   if (is_checked == null) return
   node.checked = is_checked
   if (node.children) {
      node.children.forEach((child: CustomTreeNode) => update_node_and_children(child, is_checked))
   }
}

const handle_change = (event: Event) => {
   const target = event.target as HTMLInputElement
   const is_checked = target.checked
   update_node_and_children(node.value, is_checked)
   emit("update:checked", is_checked)
}

const update_child_checked = (index: number, is_checked: boolean | null) => {
   if (is_checked == null) return
   if (node.value.children) {
      node.value.children[index].checked = is_checked
      update_parent_state(node.value)
   }
}

const update_parent_state = (node: CustomTreeNode) => {
   if (!node.children) return
   const all_checked = node.children.every((child) => child.checked === true)
   const some_checked = node.partially_checked()
   node.checked = all_checked ? true : some_checked ? false : false
}

watch(
   () => node.value.children,
   () => {
      update_parent_state(node.value)
   },
   { deep: true },
)
</script>
