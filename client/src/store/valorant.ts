import { writable } from 'svelte/store'

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
