export interface User extends Player {
  friends: Player[]
}

export interface Player {
  username: string
  tag: string
}
