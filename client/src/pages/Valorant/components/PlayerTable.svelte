<script context="module" lang="ts">
  export const DEFAULT_STAT = 'averageScore'
  export type Stat =
    | 'averageScore'
    | 'averageDamagePerMatch'
    | 'eloChange'
    | 'headshotPercentage'
    | 'killDeathAssistRatio'
    | 'killDeathRatio'
    | 'rankRating'
</script>

<script lang="ts">
  import SvelteTooltip from 'svelte-tooltip'

  import {
    valorantPodium,
    selectedStat,
    valorantStore,
  } from 'src/store/valorant'
  import type { Player } from 'src/types'

  const isUser = ({ username, tag }: Player) =>
    username.toLowerCase() === $valorantStore.username.toLowerCase() &&
    tag.toLowerCase() === $valorantStore.tag.toLowerCase()
  const selectNewStat = (stat: Stat) => selectedStat.set(stat)

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
          <span class:selected={$selectedStat === 'eloChange'}> Elo +/- </span>
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

<style lang="scss">
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
    .heading-row,
    .player-row {
      font-size: small;
    }
    .player-row td {
      padding: 5px;
    }
  }
</style>
