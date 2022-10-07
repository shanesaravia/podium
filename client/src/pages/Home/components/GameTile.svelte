<script lang="ts">
  import { push } from 'svelte-spa-router'

  export let title: string
  export let alt: string
  export let filename: string
  export let active = false
  export let link = ''

  const redirect = () => {
    if (link) push(link)
  }
</script>

<div class="tile-container" on:click={redirect}>
  <div class="tile-content">
    <h2 class="tile-title">{title}</h2>
    {#if !active}
      <h2 class="hide tile-subtitle">(Coming Soon)</h2>
    {/if}
  </div>
  <img src={`/images/tiles/${filename}`} {alt} class="tile" width="100%" />
</div>

<style>
  .hide {
    display: none;
  }
  .tile-container {
    position: relative;
    opacity: 0.7;
    width: 33%;
    padding: 3px;
  }
  .tile-container:hover {
    cursor: pointer;
    opacity: 1;
  }
  .tile-container:hover .tile-subtitle {
    display: block;
  }
  .tile-content {
    position: absolute;
    width: 100%;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  .tile-title {
    color: white;
    font-size: 3vw;
  }
  .tile-subtitle {
    font-size: 2vw;
  }
  .tile {
    border-radius: 10px;
  }
</style>
