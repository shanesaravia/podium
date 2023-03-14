import { useMutation } from '@sveltestack/svelte-query'

import { apiUrls } from 'src/constants'
import type { Player } from 'src/types'

export interface UserStats {
  stats: Stats
  summary: Summary
}

export interface Stats {
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

interface GetStatsParams {
  userList: Player[]
}

export const getStats = async ({ userList }: GetStatsParams) => {
  const response = await fetch(`${apiUrls.baseUrl}${apiUrls.valorantStats}`, {
    method: 'POST',
    body: JSON.stringify(userList),
  })
  if (!response.ok) {
    throw new Error('Failed to get user data')
  }
  return await response.json()
}

export const useAPIGetStats = () => {
  return useMutation<UserStats[], Error, any>(args => getStats({ ...args }))
}
