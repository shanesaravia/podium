<script lang="ts">
  import { SvelteToast } from '@zerodevx/svelte-toast'
  import * as bootstrap from 'bootstrap'
  import { onMount } from 'svelte'
  import { Router } from 'svelte-navigator'

  import Routes from 'src/components/Routes.svelte'
  import Sidebar from 'src/components/Sidebar.svelte'

  const options = {
    theme: {
      '--toastBarHeight': 0,
    },
    intro: { x: 192 },
  }

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
  <SvelteToast {options} />
</main>

<style>
  :root {
    --toastBorderRadius: 10px;
    --toastContainerTop: auto;
    --toastContainerRight: 40px;
    --toastContainerBottom: 20px;
    --toastContainerLeft: auto;
  }
  :global(.info) {
    --toastBackground: #41a3e2;
  }
  :global(.success) {
    --toastBackground: #4caf4f;
  }
  :global(.warn) {
    --toastBackground: #ff9700;
  }
  :global(.error) {
    --toastBackground: #e61b21;
  }

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
