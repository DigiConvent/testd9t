import "./assets/main.css"
import 'primeicons/primeicons.css'

import { createApp } from "vue"
import { createI18n } from 'vue-i18n'
import { createPinia } from "pinia"
import PrimeVue from "primevue/config"
import Aura from "@primevue/themes/aura"

import App from "./app/App.vue"
import router from "./router"

import Toast from "primevue/toast"
import { Form } from "@primevue/forms"
import { Accordion, AccordionContent, AccordionHeader, AccordionPanel, Badge, Button, Card, Checkbox, Dialog, FloatLabel, InputGroup, InputGroupAddon, InputMask, InputText, Knob, Listbox, Menubar, Message, OrganizationChart, Popover, ProgressBar, Ripple, Select, ToastService, ToggleSwitch, TreeSelect } from "primevue"

const app = createApp(App)

app.use(createPinia())
app.use(PrimeVue, { theme: { preset: Aura } })

import de from './locales/de.json'
import en from './locales/en.json'
import jp from "./locales/jp.json"
import { is_mini_app } from "./auth/telegram"
import JwtAuthenticator from "./auth/jwt"

app.use(createI18n({
    locale: 'de',
    fallbackLocale: 'de',
    messages: {
        de,
        en,
        jp
    }
}))
app.use(router)


app.component("Accordion", Accordion);
app.component("AccordionPanel", AccordionPanel);
app.component("AccordionHeader", AccordionHeader);
app.component("AccordionContent", AccordionContent);
app.component("Badge", Badge)
app.component("Button", Button)
app.component("Card", Card)
app.component("Dialog", Dialog)
app.component("Checkbox", Checkbox)
app.component("FloatLabel", FloatLabel)
app.component("Form", Form)
app.component("InputGroupAddon", InputGroupAddon)
app.component("InputGroup", InputGroup)
app.component("InputMask", InputMask)
app.component("InputText", InputText)
app.component("Knob", Knob)
app.component("Listbox", Listbox)
app.component("Menubar", Menubar)
app.component("Message", Message)
app.component("OrganizationChart", OrganizationChart)
app.component("Popover", Popover)
app.component("ProgressBar", ProgressBar)
app.component("Toast", Toast)
app.component("ToggleSwitch", ToggleSwitch)
app.component("TreeSelect", TreeSelect)
app.component("Select", Select)

app.directive('ripple', Ripple)

app.use(ToastService);

const auth = JwtAuthenticator.get_instance()
if (is_mini_app()) {
    auth.login_using_telegram().then(() => {
        app.mount("#app")
    })
} else {
    auth.load_permissions().then(() => {
        app.mount("#app")
    });
}
