<script lang="ts">
  import { valorantStore } from 'src/store/valorant'
  import type { Friend } from 'src/types'

  const handleRemoveUser = () => {
    valorantStore.update(data => ({
      ...data,
      username: undefined,
      tag: undefined,
    }))
  }

  const handleRemoveFriend = ({ username, tag }: Friend) => {
    valorantStore.update(data => ({
      ...data,
      friends: data.friends.filter(
        friend => !(friend.username === username && friend.tag === tag)
      ),
    }))
  }
</script>

{#if $valorantStore?.username && $valorantStore?.tag}
  {$valorantStore.username}
  {$valorantStore.tag}
  <button class="toText" on:click={handleRemoveUser}>x</button>
{/if}
{#each $valorantStore?.friends || [] as friendData}
  <p>
    {friendData.username}
    {friendData.tag}
    <button
      class="toText"
      on:click={() => handleRemoveFriend({ ...friendData })}>x</button
    >
  </p>
{/each}

<style>
  .toText,
  .toText:active {
    background: none;
    border: none;
    color: red;
  }
</style>
