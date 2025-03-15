<template>
   <div class="card flex justify-center">
      <Form class="flex flex-col gap-4">
         <div class="flex items-center gap-1 mb-4">
            <FloatLabel variant="in">
               <InputText id="name" class="flex-auto" autocomplete="off" />
               <label for="name" class="font-semibold w-24">{{
                  $t("iam.user_status.new.name")
               }}</label>
            </FloatLabel>
         </div>
         <div class="flex items-center gap-4 mb-4">
            <FloatLabel variant="in">
               <InputText id="abbr" class="flex-auto" autocomplete="off" />
               <label for="abbr" class="font-semibold w-24">{{
                  $t("iam.user_status.new.abbr")
               }}</label>
            </FloatLabel>
         </div>
         <div class="flex items-center gap-4 mb-4">
            <FloatLabel>
               <label for="description" class="font-semibold w-24">{{
                  $t("iam.user_status.new.description")
               }}</label>
               <InputText id="description" class="flex-auto" autocomplete="off" />
            </FloatLabel>
         </div>
         <div class="flex items-center gap-4 mb-8">
            <ToggleSwitch />
            <label for="archived" class="font-semibold w-24">{{
               $t("iam.user_status.new.archived")
            }}</label>
         </div>
         <div class="flex justify-end gap-2">
            <Button
               type="button"
               :label="$t('iam.user_status.new.submit')"
               @click="create_user_status"
            ></Button>
         </div>
      </Form>
   </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import * as v from "valibot"
import { useI18n } from "vue-i18n"
import { api } from "@/api"

const t = useI18n().t

const name = ref<string>("")
const name_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user_status.new.required", { field: t("iam.user_status.new.name") })),
)

const abbr = ref<string>("")
const abbr_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user_status.new.required", { field: t("iam.user_status.new.abbr") })),
)

const description = ref<string>("")
const description_check = v.pipe(
   v.string(),
   v.nonEmpty(t("iam.user_status.new.required", { field: t("iam.user_status.new.description") })),
)

const archived = ref<boolean>(false)
const archived_check = v.boolean()

const emit = defineEmits(["created"])

async function create_user_status() {
   const re = v.safeParse(
      v.object({
         name: name_check,
         abbr: abbr_check,
         description: description_check,
         archived: archived_check,
      }),
      {
         name: name.value,
         abbr: abbr.value,
         description: description.value,
         archived: archived.value,
      },
   )

   if (re.success) {
      ;(
         await api.iam.user_status.create({
            name: re.output["name"],
            abbr: re.output["abbr"],
            description: re.output["description"],
            archived: re.output["archived"],
         })
      ).fold(
         (l) => {
            console.log(l)
         },
         (id: string) => {
            console.log(id)
            emit("created", id)
         },
      )
   }
}
</script>
