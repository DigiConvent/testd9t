<template>
   <Toast />
   <div v-if="status != null">
      <table>
         <tbody>
            <tr v-for="(data, category) of status" :key="category">
               <td>{{ category }}</td>
               <td>
                  <table>
                     <tbody>
                        <tr v-for="(item, label) of data" :key="label">
                           <td>{{ label }}</td>
                           <td>{{ item }}</td>
                        </tr>
                     </tbody>
                  </table>
               </td>
            </tr>
         </tbody>
      </table>
   </div>
</template>

<script setup lang="ts">
import { ref } from "vue"
import { api } from "./../../api"
import type Either from "../../api/core/either"
import { useToast } from "primevue/usetoast"
import type { SystemStatus } from "@/api/sys/types"

const toast = useToast()

const status = ref<SystemStatus | null>(null)

api.sys.status().then((result: Either<string, SystemStatus>) => {
   result.fold(
      () => {
         toast.add({ severity: "error", summary: "API is down", life: 3000 })
      },
      (data: SystemStatus) => {
         toast.add({ severity: "info", summary: "Got system status", life: 3000 })
         status.value = data
      },
   )
})
</script>
