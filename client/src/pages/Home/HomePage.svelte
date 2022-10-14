<script lang="ts">
  import { urls } from 'src/constants'

  import GameTile from './components/GameTile.svelte'
</script>

<div class="container text-center">
  <h1 class="heading">Select Your Game</h1>
  <div class="row">
    <GameTile
      title="Valorant"
      filename="valorant.jpeg"
      alt="Valorant Tile"
      link={urls.valorant}
      active
    />
    <GameTile
      title="Rocket League"
      filename="rocketleague.jpeg"
      alt="RocketLeague Tile"
    />
    <GameTile title="CS:GO" filename="csgo.jpeg" alt="CS:GO Tile" />
    <GameTile
      title="Modern Warefare 2"
      filename="mw2.jpeg"
      alt="Modern Warefare 2 Tile"
    />
    <GameTile
      title="Rainbow Six Siege"
      filename="r6.jpeg"
      alt="Rainbow Six Siege Tile"
    />
    <GameTile
      title="Halo Infinite"
      filename="halo.jpeg"
      alt="Halo Infinite Tile"
    />
    <GameTile
      title="Overwatch 2"
      filename="overwatch2.jpeg"
      alt="Overwatch 2 Tile"
    />
    <GameTile
      title="League of Legends"
      filename="lol.jpeg"
      alt="League of Legends Tile"
    />
    <GameTile title="Fortnite" filename="fortnite.jpeg" alt="Fortnite Tile" />
  </div>
</div>

<style>
  .heading {
    padding-top: 30px;
    padding-bottom: 20px;
  }
</style>
