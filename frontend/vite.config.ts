import { fileURLToPath, URL } from "node:url"
import { defineConfig } from "vite"
import vue from "@vitejs/plugin-vue"
import tailwindcss from "@tailwindcss/vite"

// https://vitejs.dev/config/
export default defineConfig({
   plugins: [vue(), tailwindcss()],
   resolve: {
      alias: {
         "@": fileURLToPath(new URL("./src", import.meta.url)),
      },
   },
   build: {
      cssCodeSplit: true,
      minify: "esbuild",
      rollupOptions: {
         output: {
            assetFileNames: "assets/[name].[ext]",
            chunkFileNames: "chunks/[name].js",
            manualChunks: (id) => {
               if (id.includes("node_modules")) return "node_modules/"
               if (id.includes("/components"))
                  return "c/" + id.split("/components/")[1].split(".")[0]
               if (id.includes("/views/")) return "v/" + id.split("/views/")[1].split(".")[0]
               return "main"
            },
         },
      },
   },
})
