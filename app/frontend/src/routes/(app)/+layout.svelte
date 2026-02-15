<script lang="ts">
	import Header from './Header.svelte';

	let { children } = $props();
</script>

<div class="app">
	<Header />

	<main>
		{@render children()}
	</main>
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-inline-size: 100vw;
		min-block-size: 100vh;
	}

	.app main {
		flex: 1;
		display: grid;
		grid-template-columns: 1fr 60%;
		gap: var(--spacing-5);
		padding: var(--spacing-5);
		padding-block-start: 0;
	}

	@media (max-width: 480px) {
		.app main {
			grid-template-columns: 1fr;
		}
	}
</style>
