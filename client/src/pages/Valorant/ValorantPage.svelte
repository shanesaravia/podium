<script lang="ts">
  import AddUserInput from './components/AddUserInput.svelte'
  import UserList from './components/UserList.svelte'
</script>

<div class="container">
  <h1 class="heading-container">Valorant</h1>
  <AddUserInput />
  <UserList />
</div>

<style lang="scss">
  .container {
    flex-direction: column;
    width: 100%;
  }
  .heading-container {
    display: flex;
    justify-content: center;
    margin: 30px 0;
  }
</style>
