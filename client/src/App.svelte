<script lang="ts">
  import * as bootstrap from 'bootstrap'
  import { onMount } from 'svelte'
  import { Router } from 'svelte-navigator'

  import Routes from 'src/components/Routes.svelte'
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
  <Router>
    <Sidebar />
    <Routes />
  </Router>
</main>

<style>
  main {
    display: flex;
    flex-direction: row;
    margin: 0 auto;
    height: 100%;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
