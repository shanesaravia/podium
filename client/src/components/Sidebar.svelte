<script lang="ts">
  import { push } from 'svelte-spa-router'

  import { urls } from 'src/constants'
</script>

<nav id="sidebar">
  <div>
    <img
      src="/images/icons/podiumLogo.png"
      alt="podium-logo"
      class="home-icon game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-title="Home"
      on:click={() => push(urls.home)}
    />
    <div class="divider" />
    <img
      src="/images/icons/valorant.png"
      alt="valorant-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-title="Valorant"
      on:click={() => push(urls.valorant)}
    />
    <img
      src="/images/icons/rocketleague.png"
      alt="rocketleague-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Rocket League<br>(Coming Soon)"
    />
    <img
      src="/images/icons/csgo.png"
      alt="csgo-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="CS:GO<br>(Coming Soon)"
    />
    <img
      src="/images/icons/mw2.png"
      alt="modern-warefare-2-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Modern Warefare 2<br>(Coming Soon)"
    />
    <img
      src="/images/icons/r6.png"
      alt="rainbow-six-siege-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Rainbow Six Siege<br>(Coming Soon)"
    />
    <img
      src="/images/icons/halo.png"
      alt="halo-infinite-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Halo Infinite<br>(Coming Soon)"
    />
    <img
      src="/images/icons/overwatch2.png"
      alt="overwatch-2-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Overwatch 2<br>(Coming Soon)"
    />
    <img
      src="/images/icons/lol.png"
      alt="league-of-legends-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="League Of Legends<br>(Coming Soon)"
    />
    <img
      src="/images/icons/fortnite.webp"
      alt="fortnite-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Fortnite<br>(Coming Soon)"
    />
  </div>
  <div
    class="person-icon-wrapper"
    data-bs-toggle="tooltip"
    data-bs-placement="right"
    data-bs-title="Profile"
    on:click={() => push(urls.profile)}
  >
    <i class="bi bi-person-fill fa-2x person-icon" />
  </div>
</nav>

<style lang="scss">
  #sidebar {
    display: flex;
    flex-direction: column;
    width: 80px;
    justify-content: space-between;
    align-items: center;
    background-color: $dark-grey;
  }
  .home-icon {
    display: flex;
    padding-top: 10px;
    padding-bottom: 10px;
  }
  .game-icon {
    display: flex;
    width: 80px;
    padding: 15px 10px;
    justify-content: center;
  }
  .game-icon:hover,
  .person-icon-wrapper:hover {
    opacity: 0.5;
    cursor: pointer;
  }
  .person-icon-wrapper {
    display: flex;
    width: 100%;
    padding: 10px;
    justify-content: center;
    margin-bottom: 15px;
  }
  .person-icon {
    display: flex;
    color: white;
    font-size: 40px;
  }
  .divider {
    border-bottom: 1px solid grey;
  }
</style>
