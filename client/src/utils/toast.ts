import { toast } from '@zerodevx/svelte-toast'

export const successToast = () =>
  toast.push('User successfully saved', { classes: ['success'] })

export const errorToast = () =>
  toast.push('Unable to add user. Please try again later', {
    classes: ['error'],
  })

export const duplicateToast = () =>
  toast.push('User already exists', { classes: ['error'] })

export const emptyFieldToast = () =>
  toast.push('Username or tag is empty', { classes: ['warn'] })

export const limitToast = () =>
  toast.push("You've exceeded the limit", { classes: ['warn'] })

export const unableToFetchUserToast = () =>
  toast.push('Unable to fetch user profile', { classes: ['error'] })
