<script lang="ts">
  import Loading from 'src/components/Loading.svelte'
  import { useAPIGetProfile } from 'src/services/api/valorant/profile'
  import { useAPIGetStats } from 'src/services/api/valorant/stats'
  import type { UserStats } from 'src/services/api/valorant/stats'
  import {
    valorantStore,
    valorantPodium,
    prevCount,
    selectedStat,
  } from 'src/store/valorant'

  import AddUserInput from './components/AddUserInput.svelte'
  import PlayerTable, {
    Stat,
    DEFAULT_STAT,
  } from './components/PlayerTable.svelte'
  import Podium from './components/Podium.svelte'

  const getProfileQuery = useAPIGetProfile()
  $: ({ isLoading: isProfileLoading } = $getProfileQuery)
  const getStatsQuery = useAPIGetStats()
  $: ({ isLoading, isError } = $getStatsQuery)

  const getUserList = () => {
    const { username, tag, friends } = $valorantStore
    if (username && tag) {
      return [{ username, tag }, ...(friends || [])]
    } else {
      return [...(friends || [])]
    }
  }

  const calculatePodium = (playerData: UserStats[], stat: Stat) => {
    if (!playerData) return []
    let sortedData = playerData.length ? [...playerData] : []
    sortedData.sort((a, b) => b?.stats?.[stat] - a?.stats?.[stat])
    valorantPodium.set(sortedData)
    return sortedData
  }

  valorantStore.subscribe(() => {
    const userList = getUserList()
    if ($prevCount > userList.length) {
      prevCount.set(userList.length)
      return calculatePodium($valorantPodium, DEFAULT_STAT)
    }

    $getStatsQuery.mutate(
      { userList },
      {
        onSuccess: stats => {
          calculatePodium(stats, $selectedStat)
        },
        onError: () => console.error('Unable to retrieve stats'),
      }
    )
    prevCount.set(userList.length)
  })

  selectedStat.subscribe(stat => calculatePodium($valorantPodium, stat))
</script>

<div class="container">
  <h1 class="heading-container">Valorant</h1>
  {#if isLoading || isProfileLoading}
    <Loading />
  {:else if isError}
    <span>Error</span>
  {:else}
    <AddUserInput {getProfileQuery} />
    <Podium />
    <PlayerTable />
  {/if}
</div>

<style lang="scss">
  .container {
    display: flex;
    flex-direction: column;
    width: calc(100% - 80px);
    padding: 0;
  }
  .heading-container {
    display: flex;
    justify-content: center;
    margin: 30px 0;
  }
</style>
