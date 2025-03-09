<template>
  <div v-if="profile">
    <pre>{{ JSON.stringify(profile, null, 2) }} </pre>
  </div>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import type { PermissionGroupProfile } from "@/api/iam/permission_group/types"
import { useToast } from "primevue"
import { ref } from "vue"

const props = defineProps(["modelValue"])

const toast = useToast()

const profile = ref<PermissionGroupProfile | null>(null)
function load() {
  api.iam.permission_group.get_profile(props.modelValue).then((result) => {
    result.fold(
      (err: string) => {
        toast.add({
          severity: "error",
          summary: "Error",
          detail: err
        })
      },
      (data: PermissionGroupProfile) => {
        profile.value = data
      }
    )
  })
}

load()
</script>
