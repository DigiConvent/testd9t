import type Either from "./either"

export type ApiCall<T> = (...args: any) => Promise<Either<string, T>>

export interface ApiGetById<T> extends ApiCall<T> {
   (id: string): Promise<Either<string, T>>
}

export interface ApiSaveById<I, U> extends ApiCall<U> {
   (id: string, data: I): Promise<Either<string, U>>
}

export type ApiCollection = {
   [key: string]: any | ApiCall<any>
}
