export interface User extends Friend {
  friends: User[]
}

export interface Friend {
  username: string
  tag: string
}
