import type Either from "@/api/core/either"
import type { Ref, VNode } from "vue"

export interface FieldOptions {
   key: string
   label: string
   required?: boolean
   disabled?: boolean
   [key: string]: any
}

export interface Field {
   label: string
   key: string
   value: Ref<any>
   error: Ref<string | null>
   v_node: Ref<VNode>
   update_value: (value: any) => void
   validate: () => boolean
}

export type ApiGetById<T> = (id: string) => Promise<Either<string, T>>
export type ApiSaveById<T> = (id: string, data: T) => Promise<Either<string, boolean>>
export type ApiSaveNew<T> = (data: T) => Promise<Either<string, string>>
export type ApiProvideFresh<T> = () => Promise<Either<string, T>>

export type FormInput<T> = (
   | { id: string; data?: undefined; fetch_endpoint: ApiGetById<T> } // fetch data first
   | { id?: undefined; data: T; fetch_endpoint?: undefined } // data is already present
) &
   (
      | { save_endpoint: ApiSaveNew<T>; update_endpoint?: undefined } // update instance
      | { save_endpoint?: undefined; update_endpoint: ApiSaveById<T> } // create new instance
   ) & {
      t: string
      on_failure?: (...args: any) => void
      on_success?: (...args: any) => void
   }
