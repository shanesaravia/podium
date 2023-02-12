import { useQuery, UseQueryOptions } from '@sveltestack/svelte-query'

import { apiUrls } from 'src/constants'

interface User {
  username: string
  tag: string
}

interface StoreData extends User {
  friends: User[]
}

export interface UserStats {
  [key: User['username']]: {
    stats: Stats
    summary: Summary
  }
}

interface Stats {
  averageDamagePerMatch: number
  averageScore: number
  eloChange: number
  headshotPercentage: number
  killDeathAssistRatio: number
  killDeathRatio: number
  rankRating: number
}

interface Summary {
  accountLevel: number
  currentRR: number
  currentRank: string
  lastUpdate: string
  lastUpdateRaw: number
  puuid: string
  region: string
  tag: string
  username: string
}

const GET_STATS_KEY = 'GET_STATS'
export const getStats = async (userList: User[]) => {
  const response = await fetch(`${apiUrls.baseUrl}${apiUrls.valorantStats}`, {
    method: 'POST',
    body: JSON.stringify(userList),
  })
  if (!response.ok) {
    throw new Error('Failed to get user data')
  }
  const resp = await response.json()
  // console.log('resp: ', resp)
  return resp
  // return await response.json()
}

export const useAPIGetStats = (
  storeData: StoreData,
  options?: UseQueryOptions<User[], Error, UserStats[]>
) => {
  const { username, tag, friends } = storeData
  const allUsers = [{ username, tag }, ...friends]

  return useQuery<User[], Error, UserStats[]>(
    GET_STATS_KEY,
    () => getStats(allUsers),
    { ...options, refetchOnWindowFocus: false }
  )
}

// export const useData = async (id: string) => {
//   return useQuery(["data", id], () => getDataById(id));

// const getPostById = async (id) => {
//     const { data } = await axios.get(
//         `https://jsonplaceholder.typicode.com/posts/${id}`
//     )
//     return data
// }

// const usePost = (postId: string) => {
//     return useQuery<{ title: string; body: string }, AxiosError>(
//         ['post', postId],
//         () => getPostById(postId),
//         {
//             enabled: !!postId,
//         })
// }

// export default usePost

// Example.svelte
//  import { useQuery } from '@sveltestack/svelte-query'

//  const queryResult = useQuery('repoData', () =>
//    fetch('https://api.github.com/repos/SvelteStack/svelte-query').then(res =>
//      res.json()
//    )
//  )

//   import { useQuery } from '@sveltestack/svelte-query'
// import axios, { AxiosError } from 'axios'
