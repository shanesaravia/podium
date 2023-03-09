<script context="module" lang="ts">
  export const DEFAULT_STAT = 'averageScore'
</script>

<script lang="ts">
  import SvelteTooltip from 'svelte-tooltip'

  import Loading from 'src/components/Loading.svelte'
  import type { UserStats } from 'src/services/api/valorant/stats'
  import { useAPIGetStats } from 'src/services/api/valorant/stats'
  import {
    valorantStore,
    valorantPodium,
    prevCount,
    selectedStat,
  } from 'src/store/valorant'
  import type { Player } from 'src/types'

  type Stat =
    | 'averageScore'
    | 'averageDamagePerMatch'
    | 'eloChange'
    | 'headshotPercentage'
    | 'killDeathAssistRatio'
    | 'killDeathRatio'
    | 'rankRating'

  const isUser = ({ username, tag }: Player) =>
    username.toLowerCase() === $valorantStore.username.toLowerCase() &&
    tag.toLowerCase() === $valorantStore.tag.toLowerCase()

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

  const selectNewStat = (stat: Stat) => selectedStat.set(stat)

  const calculatePodium = (playerData: UserStats[], stat: Stat) => {
    if (!playerData) return []
    let sortedData = playerData.length ? [...playerData] : []
    sortedData.sort((a, b) => b?.stats?.[stat] - a?.stats?.[stat])
    valorantPodium.set(sortedData)
    return sortedData
  }

  selectedStat.subscribe(stat => calculatePodium($valorantPodium, stat))

  const removeFromPodium = ({ username, tag }: Player) => {
    valorantPodium.update(data => [
      ...data.filter(
        player =>
          !(
            player?.summary?.username.toLowerCase() ===
              username.toLowerCase() &&
            player?.summary?.tag.toLowerCase() === tag.toLowerCase()
          )
      ),
    ])
  }

  const handleRemoveUser = () => {
    const username = $valorantStore.username
    const tag = $valorantStore.tag
    valorantStore.update(data => ({
      ...data,
      username: undefined,
      tag: undefined,
    }))
    removeFromPodium({ username, tag })
  }

  const handleRemoveFriend = ({ username, tag }: Player) => {
    valorantStore.update(data => ({
      ...data,
      friends: data.friends.filter(
        friend =>
          !(
            friend.username.toLowerCase() === username.toLowerCase() &&
            friend.tag.toLowerCase() === tag.toLowerCase()
          )
      ),
    }))
    removeFromPodium({ username, tag })
  }
</script>

{#if isLoading}
  <Loading />
{:else if isError}
  <span>Error</span>
{:else}
  {#if $valorantPodium.length}
    <div class="podium">
      <div class="second">
        <h2 class="podium-user">
          {$valorantPodium[1]?.summary.username || '?'}
        </h2>
        <div class="bar2" />
      </div>
      <div class="first">
        <h2 class="podium-user">
          {$valorantPodium[0]?.summary.username || '?'}
        </h2>
        <div class="bar1" />
      </div>
      <div class="third">
        <h2 class="podium-user">
          {$valorantPodium[2]?.summary.username || '?'}
        </h2>
        <div class="bar3" />
      </div>
    </div>
  {/if}
  <div class="podium-table-container">
    <table class="podium-table">
      <tr class="heading-row">
        <th />
        <th>Username</th>
        <th>Tag</th>
        <th
          class:pointer={!($selectedStat === 'averageScore')}
          on:click={() => selectNewStat('averageScore')}
        >
          <SvelteTooltip tip="Average Score" bottom>
            <span class:selected={$selectedStat === 'averageScore'}>AS</span>
          </SvelteTooltip>
        </th>
        <th
          class:pointer={!($selectedStat === 'averageDamagePerMatch')}
          on:click={() => selectNewStat('averageDamagePerMatch')}
        >
          <SvelteTooltip tip="Average Damage Per Match" bottom>
            <span class:selected={$selectedStat === 'averageDamagePerMatch'}>
              ADM
            </span>
          </SvelteTooltip>
        </th>
        <th
          class:pointer={!($selectedStat === 'eloChange')}
          on:click={() => selectNewStat('eloChange')}
        >
          <SvelteTooltip tip="Elo Change" bottom>
            <span class:selected={$selectedStat === 'eloChange'}>
              Elo +/-
            </span>
          </SvelteTooltip>
        </th>
        <th
          class:pointer={!($selectedStat === 'headshotPercentage')}
          on:click={() => selectNewStat('headshotPercentage')}
        >
          <SvelteTooltip tip="Headshot Percentage" bottom>
            <span class:selected={$selectedStat === 'headshotPercentage'}
              >HS %</span
            >
          </SvelteTooltip>
        </th>
        <th
          class:pointer={!($selectedStat === 'killDeathAssistRatio')}
          on:click={() => selectNewStat('killDeathAssistRatio')}
        >
          <SvelteTooltip tip="Kill/Death/Assist Ratio" bottom>
            <span class:selected={$selectedStat === 'killDeathAssistRatio'}
              >KDA</span
            >
          </SvelteTooltip>
        </th>
        <th
          class:pointer={!($selectedStat === 'killDeathRatio')}
          on:click={() => selectNewStat('killDeathRatio')}
        >
          <SvelteTooltip tip="Kill/Death Ratio" bottom>
            <span class:selected={$selectedStat === 'killDeathRatio'}>KD</span>
          </SvelteTooltip>
        </th>
        <th
          class:pointer={!($selectedStat === 'rankRating')}
          on:click={() => selectNewStat('rankRating')}
        >
          <span class:selected={$selectedStat === 'rankRating'}>Rank</span>
        </th>
        <th />
      </tr>
      {#each $valorantPodium || [] as player}
        <div class="player-row-gap" />
        <tr class="player-row">
          <td>
            <button
              class="toText"
              on:click={() => {
                const { username, tag } = player.summary
                if (isUser({ username, tag })) {
                  handleRemoveUser()
                } else {
                  handleRemoveFriend({ username, tag })
                }
              }}>x</button
            >
          </td>
          <td>{player.summary.username}</td>
          <td>{player.summary.tag}</td>
          <td class:selected={$selectedStat === 'averageScore'}
            >{player.stats.averageScore}</td
          >
          <td class:selected={$selectedStat === 'averageDamagePerMatch'}
            >{player.stats.averageDamagePerMatch}</td
          >
          <td class:selected={$selectedStat === 'eloChange'}
            >{player.stats.eloChange}</td
          >
          <td class:selected={$selectedStat === 'headshotPercentage'}
            >{player.stats.headshotPercentage}</td
          >
          <td class:selected={$selectedStat === 'killDeathAssistRatio'}
            >{player.stats.killDeathAssistRatio}</td
          >
          <td class:selected={$selectedStat === 'killDeathRatio'}
            >{player.stats.killDeathRatio}</td
          >
          <td class:selected={$selectedStat === 'rankRating'}
            >{player.stats.rankRating}</td
          >
          <td />
        </tr>
      {/each}
    </table>
  </div>
{/if}

<style lang="scss">
  .podium {
    display: flex;
    height: 30%;
    align-self: center;
    max-width: 600px;
    width: 90%;
  }
  .podium-user {
    z-index: 9999;
    align-self: center;
  }
  .bar1,
  .bar2,
  .bar3 {
    width: 100%;
    border-radius: 10px 10px 0 0;
  }
  .first,
  .second,
  .third {
    height: 100%;
    width: 33%;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    align-items: center;
  }
  .bar1 {
    height: 80%;
    background-color: #fd4654;
  }
  .bar2 {
    height: 60%;
    background-color: #0160b7;
  }
  .bar3 {
    height: 40%;
    background-color: #f2b404;
  }

  td:first-child,
  th:first-child {
    border-radius: 10px 0 0 10px;
  }
  td:last-child,
  th:last-child {
    border-radius: 0 10px 10px 0;
  }

  tr:first-child td {
    border-top-style: solid;
  }
  tr td:first-child {
    border-left-style: solid;
  }
  .podium-table-container {
    justify-content: center;
    overflow-x: auto;
    margin: 20px 0;
  }
  .podium-table {
    width: 100%;
    border-collapse: collapse;
  }

  .heading-row {
    background-color: $dark;
  }

  .player-row {
    background-color: $table-grey;
    border: solid 1px #000;
    border-style: none solid solid none;
  }
  .player-row-gap {
    height: 10px;
  }
  th {
    text-align: center;
  }
  td {
    padding: 13px;
    text-align: center;
  }

  .toText,
  .toText:active {
    background: none;
    border: none;
    color: red;
  }
  .selected {
    color: #fd4654;
  }
  .pointer {
    cursor: pointer;
  }
  @media only screen and (max-width: 770px) {
    .podium-user {
      font-size: medium;
    }
    .heading-row,
    .player-row {
      font-size: small;
    }
    .player-row td {
      padding: 5px;
    }
  }
  @media only screen and (max-width: 400px) {
    .podium-user {
      font-size: small;
    }
  }
</style>
