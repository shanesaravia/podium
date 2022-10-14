<script lang="ts">
  import Router from 'svelte-spa-router'

  import Home from 'src/pages/Home/HomePage.svelte'
  import PageNotFound from 'src/pages/PageNotFound.svelte'
  import ProfilePage from 'src/pages/Profile/ProfilePage.svelte'
  import ValorantPage from 'src/pages/Valorant/ValorantPage.svelte'

  const routes = {
    '/': Home,
    '/profile': ProfilePage,
    '/valorant': ValorantPage,
    '*': PageNotFound, // catchall,
  }
</script>

<Router {routes} />
