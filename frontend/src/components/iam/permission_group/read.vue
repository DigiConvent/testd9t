<template>
   <ProgressBar v-if="loading" mode="indeterminate"></ProgressBar>
   <div v-else-if="permission_group" v-permission="'iam.permission_group.read'">
      <table>
         <tr>
            <td>Name</td>
            <td>{{ permission_group.name }}</td>
         </tr>
         <tr>
            <td>Abbr</td>
            <td>{{ permission_group.abbr }}</td>
         </tr>
         <tr>
            <td>Description</td>
            <td>{{ permission_group.description }}</td>
         </tr>
         <tr v-if="permission_group.meta != null">
            <td>Type</td>
            <td>{{ permission_group.meta }}</td>
         </tr>
      </table>
   </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupRead } from "@/api/iam/permission_group/types"
import type { IdOrData } from "@/components/form/form"
import { error } from "@/composables/toast"
import { ref } from "vue"

const loading = ref(true)
const props = defineProps<IdOrData<PermissionGroupRead>>()

const permission_group = ref<PermissionGroupRead | null>(null)

async function load() {
   loading.value = true
   if (props.id != undefined) {
      ;(await api.iam.permission_group.get(props.id)).fold(
         (err: string) => {
            error(err)
         },
         (data: PermissionGroupRead) => {
            permission_group.value = data
            loading.value = false
         },
      )
   } else {
      permission_group.value = props.data
   }
   loading.value = false
}

load()
</script>
