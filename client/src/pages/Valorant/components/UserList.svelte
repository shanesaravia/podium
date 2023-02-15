<script lang="ts">
  import Loading from 'src/components/Loading.svelte'
  import type { UserStats } from 'src/services/api/valorant/stats'
  import { useAPIGetStats } from 'src/services/api/valorant/stats'
  import { valorantStore, valorantPodium, prevCount } from 'src/store/valorant'
  import type { Friend } from 'src/types'

  type Stat =
    | 'averageScore'
    | 'averageDamagePerMatch'
    | 'eloChange'
    | 'headshotPercentage'
    | 'killDeathAssistRatio'
    | 'killDeathRatio'
    | 'rankRating'

  const getStatsQuery = useAPIGetStats()
  $: ({ isLoading, isError } = $getStatsQuery)
  $: console.log('podium: ', $valorantPodium)

  const getUserList = () => {
    const { username, tag, friends } = $valorantStore
    if (username && tag) {
      return [{ username, tag }, ...friends]
    } else {
      return [...friends]
    }
  }

  // Any time the store changes it will run this code. If we added a user (current count is higher than previous count)
  valorantStore.subscribe(() => {
    const userList = getUserList()
    if ($prevCount > userList.length) {
      prevCount.set(userList.length)
      return calculatePodium($valorantPodium, 'rankRating')
    }

    $getStatsQuery.mutate(
      { userList },
      {
        onSuccess: stats => {
          calculatePodium(stats, 'rankRating')
        },
        onError: () => console.log('failure'),
      }
    )
    prevCount.set(userList.length)
  })

  const calculatePodium = (playerData: UserStats[], stat: Stat) => {
    if (!playerData) return []
    let sortedData = playerData.length ? [...playerData] : []
    sortedData.sort(
      (a, b) =>
        Object.values(b)?.[0]?.stats?.[stat] -
        Object.values(a)?.[0]?.stats?.[stat]
    )
    valorantPodium.set(sortedData)
    console.log('valorant podium: ', $valorantPodium)
    return sortedData
  }

  const updatePodium = ({
    username,
    tag,
  }: {
    username: string
    tag: string
  }) => {
    valorantPodium.update(data => ({
      ...data.filter(
        player =>
          !player[username.toLowerCase()] ||
          !(
            player?.[username]?.summary?.username.toLowerCase() ===
              username.toLowerCase() &&
            player?.[username]?.summary?.tag.toLowerCase() === tag.toLowerCase()
          )
      ),
    }))
  }

  const handleRemoveUser = () => {
    const username = $valorantStore.username
    const tag = $valorantStore.tag
    valorantStore.update(data => ({
      ...data,
      username: undefined,
      tag: undefined,
    }))
    updatePodium({ username, tag })
  }

  const handleRemoveFriend = ({ username, tag }: Friend) => {
    valorantStore.update(data => ({
      ...data,
      friends: data.friends.filter(
        friend => !(friend.username === username && friend.tag === tag)
      ),
    }))
    updatePodium({ username, tag })
  }
</script>

{#if isLoading}
  <Loading />
{:else if isError}
  <span>Error</span>
{:else}
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
{/if}

<style>
  .toText,
  .toText:active {
    background: none;
    border: none;
    color: red;
  }
</style>
