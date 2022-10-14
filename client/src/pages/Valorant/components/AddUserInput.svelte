<script lang="ts">
  import { valorantStore } from 'src/store/valorant'
  import type { FriendsList } from 'src/types'
  import {
    duplicateToast,
    emptyFieldToast,
    errorToast,
    limitToast,
    successToast,
  } from 'src/utils/toast'

  let username: string
  let tag: string

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

      if (existsInFriends($valorantStore?.friends || [])) {
        duplicateToast()
        return
      }

      successToast()
      valorantStore.update(data => ({ ...data, username, tag }))
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

      const friends = $valorantStore?.friends || []
      if (friends.length > 10) {
        limitToast()
        return
      }
      if (existsInFriends(friends) || $valorantStore.username === username) {
        duplicateToast()
        return
      }

      valorantStore.update(data => {
        successToast()
        friends.push({ username, tag })

        return { ...data, friends }
      })
      clearInputFields()
    } catch (error) {
      errorToast()
      console.error(error.message)
    }
  }
</script>

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
      class:disabled={$valorantStore?.username && $valorantStore?.tag}
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

<style lang="scss">
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
