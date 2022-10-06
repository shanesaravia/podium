<div id="sidebar">
  <div>
    <img
      src="/images/podiumLogo.png"
      alt="podium-logo"
      class="home-icon game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-title="Home"
    />
    <div class="divider" />
    <img
      src="/images/valorant.png"
      alt="valorant-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-title="Valorant"
    />
    <img
      src="/images/rocketleague.png"
      alt="rocketleague-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Rocket League<br>(Coming Soon)"
    />
    <img
      src="/images/csgo.png"
      alt="csgo-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="CS:GO<br>(Coming Soon)"
    />
    <img
      src="/images/mw2.png"
      alt="mw2-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Modern Warefare 2<br>(Coming Soon)"
    />
    <img
      src="/images/r6.png"
      alt="r6-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Rainbow Six Siege<br>(Coming Soon)"
    />
    <img
      src="/images/halo.png"
      alt="halo-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Halo Infinite<br>(Coming Soon)"
    />
    <img
      src="/images/overwatch2.png"
      alt="overwatch2-logo"
      class="game-icon"
      data-bs-toggle="tooltip"
      data-bs-placement="right"
      data-bs-html="true"
      data-bs-title="Overwatch 2<br>(Coming Soon)"
    />
  </div>
  <div
    class="person-icon-wrapper"
    data-bs-toggle="tooltip"
    data-bs-placement="right"
    data-bs-title="Account"
  >
    <i class="bi bi-person-fill fa-2x person-icon" />
  </div>
</div>

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
    padding: 20px 10px 20px 10px;
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
