import eslint from "@eslint/js"
import eslintConfigPrettier from "eslint-config-prettier"
import eslintPluginVue from "eslint-plugin-vue"
import globals from "globals"
import typescriptEslint from "typescript-eslint"

export default typescriptEslint.config(
   { ignores: ["*.d.ts", "**/coverage", "**/dist"] },
   {
      extends: [
         eslint.configs.recommended,
         ...typescriptEslint.configs.recommended,
         ...eslintPluginVue.configs["flat/recommended"],
      ],
      files: ["**/*.{ts,vue}"],
      languageOptions: {
         ecmaVersion: "latest",
         sourceType: "module",
         globals: globals.browser,
         parserOptions: {
            parser: typescriptEslint.parser,
         },
      },
      rules: {
         "vue/multi-word-component-names": "off",
         "vue/no-reserved-component-names": "off",
         "vue/no-unused-vars": "error",
         "vue/no-unused-components": "error",
         "vue/prop-name-casing": ["error", "snake_case"],
         "vue/no-multiple-template-root": "error",
         "@typescript-eslint/indent": ["error", 3],
         "@typescript-eslint/naming-convention": [
            "error",
            {
               selector: "variableLike",
               format: ["snake_case"],
            },
            {
               selector: "function",
               format: ["snake_case"],
            },
         ],
         "@typescript-eslint/no-explicit-any": "off",
         "comma-dangle": ["error", "never"],
         "@typescript-eslint/comma-dangle": ["error", "never"],
      },
   },
   eslintConfigPrettier,
)
