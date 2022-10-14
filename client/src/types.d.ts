export interface User extends Friend {
  friends: FriendsList
}

export interface Friend {
  username: string
  tag: string
}

type FriendsList = Friend[]
