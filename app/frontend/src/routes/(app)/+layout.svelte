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
		block-size: 100vh;
		overflow: hidden;
	}

	.app main {
		flex: 1;
		display: grid;
		grid-template-columns: 1fr 60%;
		grid-template-rows: minmax(0, 1fr);
		gap: var(--spacing-5);
		padding: var(--spacing-5);
		padding-block-start: 0;
		min-height: 0;
	}

	@media (max-width: 480px) {
		.app main {
			grid-template-columns: 1fr;
		}
	}
</style>
