<script lang="ts">
	import Header from '../Header.svelte';

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
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
</style>
