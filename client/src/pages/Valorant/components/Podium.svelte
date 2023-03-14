<script lang="ts">
  import { valorantPodium } from 'src/store/valorant'
</script>

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

  @media only screen and (max-width: 770px) {
    .podium-user {
      font-size: medium;
    }
  }
  @media only screen and (max-width: 400px) {
    .podium-user {
      font-size: small;
    }
  }
</style>
