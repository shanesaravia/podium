<script lang="ts">
  import * as bootstrap from 'bootstrap'
  import { onMount } from 'svelte'

  import Router from 'src/components/Router.svelte'
  import Sidebar from 'src/components/Sidebar.svelte'

  onMount(() => {
    const tooltipTriggerList = document.querySelectorAll(
      '[data-bs-toggle="tooltip"]'
    )
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const tooltipList = [...tooltipTriggerList].map(
      tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl)
    )
  })
</script>

<main>
  <Sidebar />
  <Router />
</main>

<style>
  main {
    display: flex;
    flex-direction: row;
    max-width: 240px;
    margin: 0 auto;
    height: 100%;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
