<script lang="ts">
  // import { toast } from '@zerodevx/svelte-toast'
  import { get } from 'svelte/store'

  import { valorantStore } from 'src/store/valorant'
  import type { FriendsList } from 'src/types'
  import {
    duplicateToast,
    emptyFieldToast,
    errorToast,
    limitToast,
    successToast
  } from 'src/utils/toast'

  let username: string
  let tag: string
  let storeData: Record<string, any> = get(valorantStore) || {}

  // const successToast = () =>
  //   toast.push('User successfully saved', { classes: ['success'] })

  // const errorToast = () =>
  //   toast.push('Unable to add user. Please try again later', {
  //     classes: ['error'],
  //   })

  // const duplicateToast = () =>
  //   toast.push('User already exists', { classes: ['error'] })

  // const emptyFieldToast = () =>
  //   toast.push('Username or tag is empty', { classes: ['warn'] })

  // const limitToast = () =>
  //   toast.push("You've exceeded the limit", { classes: ['warn'] })

  const existsInFriends = (friends: FriendsList) => {
    return friends.find(
      friend => friend.username === username && friend.tag === tag
    )
  }

  const clearInputFields = () => {
    username = ''
    tag = ''
  }

  const handleSave = () => {
    try {
      if (!(username && tag)) {
        emptyFieldToast()
        return
      }

      const currentData = get(valorantStore)
      if (existsInFriends(currentData?.friends || [])) {
        duplicateToast()
        return
      }

      successToast()
      valorantStore.update(data => ({ ...data, username, tag }))
      storeData = get(valorantStore)
      clearInputFields()
    } catch (error) {
      errorToast()
      console.error(error.message)
    }
  }

  const handleAddFriend = () => {
    try {
      if (!(username && tag)) {
        emptyFieldToast()
        return
      }

      const currentData = get(valorantStore)
      const friends = currentData?.friends || []
      if (friends.length > 10) {
        limitToast()
        return
      }
      if (existsInFriends(friends) || currentData.username === username) {
        duplicateToast()
        return
      }

      valorantStore.update(data => {
        successToast()
        friends.push({ username, tag })
        storeData = get(valorantStore)

        return { ...data, friends }
      })
      clearInputFields()
    } catch (error) {
      errorToast()
      console.error(error.message)
    }
  }
</script>

<div class="container">
  <h1 class="heading-container">Valorant</h1>
  <div class="game-input-container">
    <div class="form-floating username-wrapper">
      <input
        type="text"
        class="form-control username"
        id="floatingInput"
        placeholder="username"
        maxlength="16"
        bind:value={username}
      />
      <label for="floatingInput">Username</label>
    </div>
    <div class="form-floating tag-wrapper">
      <input
        type="text"
        class="form-control tag"
        id="floatingInput"
        placeholder="username"
        maxlength="5"
        bind:value={tag}
      />
      <label for="floatingInput">Tag</label>
    </div>
    <div class="button-container">
      <button
        class="btn btn-primary submit save"
        class:disabled={storeData?.username && storeData?.tag}
        type="button"
        on:click={handleSave}
      >
        Save
      </button>
      <button
        class="btn btn-secondary submit friend"
        type="button"
        on:click={handleAddFriend}
      >
        Add Friend
      </button>
    </div>
  </div>
  {#if storeData.username && storeData.tag}
    {storeData.username}
    {storeData.tag}
  {/if}
  {#each get(valorantStore)?.friends || [] as friendData}
    <p>
      {friendData.username}
      {friendData.tag}
    </p>
  {/each}
</div>

<style lang="scss">
  .container {
    flex-direction: column;
    width: 100%;
  }
  .heading-container {
    display: flex;
    justify-content: center;
    margin: 30px 0;
  }
  .form-floating {
    color: $dark;
  }
  .game-input-container {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
  .username-wrapper {
    width: 30%;
    min-width: 180px;
    margin: 5px;
  }
  .tag-wrapper {
    width: 10%;
    min-width: 80px;
    margin: 5px;
  }
  .username,
  .tag {
    height: 100% !important;
  }
  .submit,
  .submit:hover,
  .submit:active,
  .submit:focus {
    color: white !important;
  }
  .button-container {
    display: flex;
    min-width: 165px;
    margin: 5px;
  }
  .save {
    border-radius: $btn-border-radius 0 0 $btn-border-radius;
  }
  .friend {
    border-radius: 0 $btn-border-radius $btn-border-radius 0;
    flex-wrap: nowrap;
  }
</style>
