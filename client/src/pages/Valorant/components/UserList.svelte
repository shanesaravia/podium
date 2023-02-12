<script lang="ts">
  // import { onMount } from 'svelte'

  import Loading from 'src/components/Loading.svelte'
  import type { UserStats } from 'src/services/api/valorant/stats'
  import { useAPIGetStats, getStats } from 'src/services/api/valorant/stats'
  import { valorantStore, valorantPodium } from 'src/store/valorant'
  import type { Friend } from 'src/types'

  type Stat =
    | 'averageScore'
    | 'averageDamagePerMatch'
    | 'eloChange'
    | 'headshotPercentage'
    | 'killDeathAssistRatio'
    | 'killDeathRatio'
    | 'rankRating'

  // onMount(async () => {
  //   const allUsers = getAllUsers()
  //   const stats = await getStats(allUsers)
  //   console.log('stats: ', stats)
  //   // const res = await fetch(`/tutorial/api/album`)
  //   // photos = await res.json()
  // })

  // const getAllUsers = () => {
  //   const { username, tag, friends } = $valorantStore
  //   return [{ username, tag }, ...friends]
  // }

  const queryResult = useAPIGetStats($valorantStore)
  $: ({ data, isLoading, error } = $queryResult)
  $: console.log('query data: ', data)
  $: console.log('query loading: ', isLoading)

  // Any time the store changes it will run this code. If we added a user (current count is higher than previous count)
  valorantStore.subscribe(async user => {
    console.log('RUNNINGGGGGGG: ')
    const { username, tag, friends } = user
    const allUsers = [{ username, tag }, ...friends]
    // const currentUsers = getAllUsers()
    // calculatePodium(stats, 'averageScore')

    // return user
    // if (currentUsers.length > allUsers.length) {
    const stats = await getStats(allUsers)
    await calculatePodium(stats, 'averageScore')
    // }
  })

  const calculatePodium = (playerData: UserStats[], stat: Stat) => {
    //   // valorantPodium.set(['test'])
    console.log('order podium stat: ', stat)
    //   console.log('apiData: ', apiData)
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
  // console.log('valorant podium: ', $valorantPodium)
  // // data.sort((a, b) => a.value - b.value);

  // let podiumData
  // if (data) {
  //   console.log('data here: ', data)
  //   podiumData = calculatePodium(data, 'averageScore')
  //   console.log('podiumData: ', podiumData)
  // }
  // console.log('podiumData: ', podiumData)

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

{#if isLoading || !data}
  <Loading />
{:else if error}
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
