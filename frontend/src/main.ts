import "./assets/main.css"

import { createApp } from "vue"
import { createI18n } from "vue-i18n"
import PrimeVue from "primevue/config"
import Aura from "@primevue/themes/aura"

import { library, config } from "@DigiConvent/ff/fontawesome-svg-core"
import { FontAwesomeIcon, FontAwesomeLayers } from "@DigiConvent/ff/vue-fontawesome"
import { fasds } from "@DigiConvent/ff/sharp-duotone-solid-svg-icons"
import { faGithub } from "@DigiConvent/ff/free-brands-svg-icons"

import App from "./app/App.vue"
import router from "./router"

import Toast from "primevue/toast"
import { Form } from "@primevue/forms"
import {
   Accordion,
   AccordionContent,
   AccordionHeader,
   AccordionPanel,
   Badge,
   Button,
   Card,
   Checkbox,
   Dialog,
   Drawer,
   Fieldset,
   FloatLabel,
   InputGroup,
   InputGroupAddon,
   InputMask,
   InputText,
   Listbox,
   Menu,
   Menubar,
   Message,
   MeterGroup,
   OrganizationChart,
   Popover,
   ProgressBar,
   ProgressSpinner,
   Ripple,
   Select,
   Skeleton,
   Splitter,
   SplitterPanel,
   Textarea,
   Timeline,
   ToastService,
   ToggleButton,
   TreeSelect,
} from "primevue"

const app = createApp(App)

app.use(PrimeVue, { theme: { preset: Aura } })

import de from "./locales/de.json"
import en from "./locales/en.json"
import jp from "./locales/jp.json"
import { is_mini_app } from "./auth/telegram"
import JwtAuthenticator from "./auth/jwt"
import Auth from "./components/auth.vue"

app.use(
   createI18n({
      locale: "de",
      fallbackLocale: "de",
      messages: {
         de,
         en,
         jp,
      },
   }),
)

config.familyDefault = "sharp-duotone"
library.add(fasds)
library.add(faGithub)

app.use(router)

app.component("Accordion", Accordion)
app.component("AccordionPanel", AccordionPanel)
app.component("AccordionHeader", AccordionHeader)
app.component("AccordionContent", AccordionContent)
app.component("Badge", Badge)
app.component("Button", Button)
app.component("Card", Card)
app.component("Dialog", Dialog)
app.component("Drawer", Drawer)
app.component("Checkbox", Checkbox)
app.component("Fieldset", Fieldset)
app.component("FloatLabel", FloatLabel)
app.component("Form", Form)
app.component("InputGroupAddon", InputGroupAddon)
app.component("InputGroup", InputGroup)
app.component("InputMask", InputMask)
app.component("InputText", InputText)
app.component("Listbox", Listbox)
app.component("Menu", Menu)
app.component("Menubar", Menubar)
app.component("MeterGroup", MeterGroup)
app.component("Message", Message)
app.component("OrganizationChart", OrganizationChart)
app.component("Popover", Popover)
app.component("ProgressBar", ProgressBar)
app.component("ProgressSpinner", ProgressSpinner)
app.component("Textarea", Textarea)
app.component("Timeline", Timeline)
app.component("Toast", Toast)
app.component("ToggleButton", ToggleButton)
app.component("TreeSelect", TreeSelect)
app.component("Select", Select)
app.component("Skeleton", Skeleton)
app.component("Splitter", Splitter)
app.component("SplitterPanel", SplitterPanel)
app.component("NeedsPermission", Auth)

app.directive("ripple", Ripple)

app.component("Fa", FontAwesomeIcon).component("fal", FontAwesomeLayers)

let remember = window.location.href.replace(window.location.origin, "")
app.use(ToastService)

const auth = JwtAuthenticator.get_instance()
if (is_mini_app()) {
   auth.login_using_telegram().then(() => {
      mount()
   })
} else {
   auth.load_permissions().then(() => {
      if (!auth.is_authenticated.value) {
         remember = "/home"
      }
      mount()
   })
}

function mount() {
   app.mount("#app")
   if (auth.is_authenticated.value) {
      router.replace({ path: remember })
   } else {
      router.replace({ name: "home" })
   }
}
