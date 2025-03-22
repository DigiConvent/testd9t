<template>
   <Card>
      <template #header>
         {{ variant }}
      </template>
      <template #content>
         <label>
            {{ `@/assets/${variant}.jpg?v=${l}` }}
            <img :src="`@/assets/${variant}.jpg?v=${l}`" alt="" />
            <input type="file" name="logo" accept="image/jpg" @input="on_upload" />
         </label>
      </template>
   </Card>
</template>

<script lang="ts" setup>
import { api } from "@/api"
import { error, info, success } from "@/composables/toast"
import { ref } from "vue"
import { useI18n } from "vue-i18n"

const props = defineProps<{ variant: "small" | "large" }>()

const t = useI18n().t
let file: File
const l = ref(new Date().getTime())
const on_upload = async (event: any) => {
   file = event.target.files[0]
   info("Success", "File Uploaded" + event.files, 3000)
   ;(await api.sys.upload_logo(props.variant, file)).fold(
      (err: string) => {
         error(err)
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
