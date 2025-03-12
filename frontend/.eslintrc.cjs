module.exports = {
   extends: ["eslint:recommended", "plugin:vue/vue3-recommended"],
   rules: {
      "vue/no-unused-vars": "error",
      quotes: ["error", "double"],
   },
   plugins: ["@typescript-eslint", "prettier"],
}
