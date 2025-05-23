<!-- eslint-disable vue/no-use-v-if-with-v-for -->
<template>
   <div class="text-2xl">
      <label
         class="select-none block"
         :class="`${is_checked && !is_indeterminate ? 'text-green-700' : is_indeterminate ? 'text-orange-600' : 'text-gray-500'} ${multiple && hovered && !readonly ? 'cursor-pointer bg-blue-100 underline' : ''}`"
         @mouseenter="hovered = true"
         @mouseleave="hovered = false"
      >
         <template v-if="props.multiple">
            <Fa
               v-if="(is_checked && !is_indeterminate) || parent_hovered"
               icon="check-circle"
               class="fa-fw"
               :class="`${parent_hovered && !is_checked ? 'text-gray-500' : 'text-green-700'}`"
            />
            <Fa v-else-if="is_indeterminate" icon="minus-circle" class="text-orange-600 fa-fw" />
            <Fa
               v-else
               :icon="
                  (props.preselected || []).includes(props.node.key) ? 'check-circle' : 'circle'
               "
               class="text-slate-400 fa-fw"
            />
         </template>
         <input
            class="hidden"
            type="checkbox"
            :checked="is_checked"
            :indeterminate.prop="is_indeterminate"
            @change="handle_change($event)"
         />
         {{ node.label }}
         <span
            v-if="(props.preselected || []).includes(props.node.key)"
            class="text-grass-200 text-xs"
            >({{ $t("inherited") }})</span
         >
      </label>
      <div v-if="node.children && !(preselected || []).includes(node.key)" class="pl-4">
         <!-- this is a summarised view of the child nodes (only show the first and last child in a non-collapsed way) -->
         <template
            v-for="sorted of [children]"
            v-if="
               // show if the parent is not checked and no summary is needed
               (is_checked && node.parent != null && !node.parent!.checked && show == false) ||
               // or if summary is needed when the node is checked and visible
               (summarised && node.checked && show != true)
            "
            :key="sorted"
         >
            <template v-if="is_checked && node.children.length > 2">
               <PermissionOption
                  :multiple="props.multiple"
                  :node="sorted[0]"
                  :parent_hovered="parent_hovered || hovered"
                  :preselected="preselected"
                  :summarised="false"
                  :readonly="(props.preselected || []).includes(props.node.key) || readonly"
                  @update:checked="update_child_checked(0, $event)"
               />
               <div class="text-2xl">
                  <div @click="show = true">
                     <Fa icon="ellipsis-v" class="text-green-700 fa-fw" />
                     <label class="text-lg cursor-pointer hover:underline">show more</label>
                  </div>
               </div>
               <PermissionOption
                  :multiple="props.multiple"
                  :node="sorted[1]"
                  :parent_hovered="parent_hovered || hovered"
                  :preselected="preselected"
                  :summarised="false"
                  :readonly="(props.preselected || []).includes(props.node.key) || readonly"
                  @update:checked="update_child_checked(node.children.length - 1, $event)"
               />
            </template>
         </template>
         <template
            v-else-if="((node.parent != null && !node.parent.checked) || show) && !node.leaf"
         >
            <PermissionOption
               v-for="(child, index) in node.children"
               :key="index"
               :multiple="props.multiple"
               :node="child"
               :parent_hovered="parent_hovered || hovered"
               :preselected="preselected"
               :summarised="!summarised"
               :readonly="readonly"
               @update:checked="update_child_checked(index, $event)"
            />
            <div
               v-if="node.parent != null && node.checked"
               class="pl-5 cursor-pointer hover:underline text-lg"
               @click="show = false"
            >
               show fewer
            </div>
         </template>
         <template v-else></template>
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
   summarised: boolean | null
   preselected: string[]
   readonly: boolean
}>()
const hovered = ref(false)

const show = ref<boolean | null>(null)
const emit = defineEmits<{
   (event: "update:checked", value: boolean | null): void
}>()
const children = computed(() => {
   const sorted = props.node.children
   if (!is_checked.value) return sorted
   if (sorted.length > 2) {
      // eslint-disable-next-line vue/no-side-effects-in-computed-properties
      if (show.value == null) show.value = false
      if (show.value == false) return [sorted[0], sorted[sorted.length - 1]]
   }
   return sorted
})

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
   if (node.children.every((child: CustomTreeNode) => child.checked === true)) {
      show.value = null
   }
}

const handle_change = (event: Event) => {
   if (event.target == null) return
   const target = event.target as HTMLInputElement
   if (target.readOnly) return
   const previous = node.value.checked
   const is_checked = target.checked
   update_node_and_children(node.value, is_checked)
   emit("update:checked", is_checked)
   if (previous != is_checked) {
      show.value = false
   }
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
   const previous = node.checked
   node.checked = all_checked ? true : some_checked ? false : false
   if (previous != node.checked) {
      show.value = false
   }
}

watch(
   node.value.children,
   () => {
      update_parent_state(node.value)
   },
   { deep: true },
)

if ((props.preselected || []).includes(props.node.key)) {
   update_node_and_children(node.value, is_checked.value)
}
</script>
