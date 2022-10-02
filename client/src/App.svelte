<script lang="ts">
	import Sidebar from 'src/components/Sidebar.svelte'
	export let name: string;
	const test = "Shane"
</script>

<main>
	<Sidebar />
	<h1>Hello {name} {test}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<img src="/images/podiumLogo.png" alt="podium-logo" />
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
