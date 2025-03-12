export default class Either<E, T> {
   private _type: string = ""
   private _dataL: E | undefined
   private _dataR: T | undefined
   constructor() {}

   fold(left: (data: E) => void, right: (data: T) => void) {
      if (this._type == "left") {
         left(this._dataL!)
      } else {
         right(this._dataR!)
      }
   }

   left(data: E): Either<E, T> {
      this._type = "left"
      this._dataL = data
      return this
   }
   right(data: T): Either<E, T> {
      this._type = "right"
      this._dataR = data
      return this
   }

   isLeft(): boolean {
      return this._type == "left"
   }

   isRight(): boolean {
      return this._type == "right"
   }

   getLeft(): E | undefined {
      return this._dataL
   }

   getRight(): T | undefined {
      return this._dataR
   }
}
