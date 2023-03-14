import { useMutation } from '@sveltestack/svelte-query'

import { apiUrls } from 'src/constants'
import type { Player } from 'src/types'
import { FetchUserException } from 'src/utils/exception'

export const getProfile = async ({ username, tag }: Player) => {
  const response = await fetch(`${apiUrls.baseUrl}${apiUrls.valorantProfile}`, {
    method: 'POST',
    body: JSON.stringify({ username, tag }),
  })
  if (!response.ok) {
    throw new FetchUserException('Failed to get user data')
  }
  return await response.json()
}

export const useAPIGetProfile = () => {
  return useMutation<boolean, Error, any>(args => getProfile({ ...args }))
}
