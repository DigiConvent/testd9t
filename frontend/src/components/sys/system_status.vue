<template>
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
import type { SystemStatus } from "@/api/sys/types"
import { error, info } from "@/composables/toast"

const status = ref<SystemStatus | null>(null)

api.sys.status().then((result: Either<string, SystemStatus>) => {
   result.fold(
      () => {
         error("API is down")
      },
      (data: SystemStatus) => {
         info("Got system status")
         status.value = data
      },
   )
})
</script>
