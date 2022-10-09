<script lang="ts">
  import { get } from 'svelte/store'

  import { valorantStore } from 'src/store/valorant'
  // import type { Friend } from 'src/types'

  // interface FriendsList {
  //   friends: Friend[]
  // }

  let username: string
  let tag: string
  // let friends: FriendsList

  const handleSave = () => {
    console.log('val store before: ', get(valorantStore))
    valorantStore.update(data => ({ ...data, username, tag }))
    console.log('val store after: ', get(valorantStore))
  }
  // const handleAddFriend = () => {
  //   console.log('val store before: ', get(valorantStore))
  //   valorantStore.update(data => {
  //     const friendsData = [...(data.friends || [])]
  //     const exists = friendsData.find(
  //       friend => friend.username === username && friend.tag === tag
  //     )
  //     if (!exists) {
  //       friends.push({ username, tag })
  //     }
  //     return { ...data, friends }
  //   })
  //   console.log('val store after: ', get(valorantStore))
  // }
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
        type="button"
        on:click={handleSave}
        >Save
      </button>
      <button class="btn btn-secondary submit friend" type="button"
        >Add Friend
      </button>
    </div>
  </div>
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
