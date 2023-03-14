import { writable } from 'svelte/store'

import { DEFAULT_STAT } from 'src/pages/Valorant/components/PlayerTable.svelte'
import type { UserStats, Stats } from 'src/services/api/valorant/stats'
import type { User } from 'src/types'

const storeKey = 'valorant'
// Get the value out of storage on load.
const stored = localStorage.getItem(storeKey)
// Set the stored value or a sane default.
export const valorantStore = writable<User>(JSON.parse(stored) || {})
// Anytime the store changes, update the local storage value.
valorantStore.subscribe(value =>
  localStorage.setItem(storeKey, JSON.stringify(value))
)

export const valorantPodium = writable<UserStats[]>([])

export const prevCount = writable<number>(0)

export const selectedStat = writable<keyof Stats>(DEFAULT_STAT)
