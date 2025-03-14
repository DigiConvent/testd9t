<template>
   <div></div>
</template>

<script lang="ts" setup>
import { CustomTreeNode } from "@/api/iam/permission/types"
import { onMounted, ref } from "vue"

const props = defineProps<{ node: CustomTreeNode }>()

const node = ref(props.node)
onMounted(() => {
   console.log(node.value)
})
</script>
