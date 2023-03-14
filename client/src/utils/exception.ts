// export function FetchUserException(message: string) {
//   this.message = message
//   this.name = 'FetchUserException'
// }

export class FetchUserException extends Error {
  constructor(message: string) {
    super(message)
    this.name = 'FetchUserException'
  }
}
