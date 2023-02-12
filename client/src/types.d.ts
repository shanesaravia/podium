export interface User extends Friend {
  friends: FriendsList
}

export interface Friend {
  username: string
  tag: string
}

export interface Player {
  username: string
  tag: string
}

type FriendsList = Friend[]
