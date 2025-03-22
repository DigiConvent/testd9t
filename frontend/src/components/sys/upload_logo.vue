<template>
   <Card>
      <template #title>
         <label :for="`upload-${variant}`">
            {{ label }}
         </label>
      </template>
      <template #content>
         <label>
            <img :src="`/assets/${variant}.jpg?v=${l}`" alt="" />
            <input
               :id="`upload-${variant}`"
               type="file"
               name="logo"
               accept="image/jpg"
               @input="on_upload"
            />
         </label>
      </template>
   </Card>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { error, success } from "@/composables/toast"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const props = defineProps<{ variant: "small" | "large"; label: string }>()

const t = useI18n().t
let file: File
const l = ref(new Date().getTime())
const on_upload = async (event: any) => {
   file = event.target.files[0]
   ;(await api.sys.upload_logo(props.variant, file)).fold(
      (err: string) => {
         error(t("sys.upload_logo.fail"), err)
      },
      (res) => {
         if (res) {
            l.value = new Date().getTime()
            success(t("sys.upload_logo.success"))
         } else error(t("sys.upload_logo.fail"))
      },
   )
}
</script>
