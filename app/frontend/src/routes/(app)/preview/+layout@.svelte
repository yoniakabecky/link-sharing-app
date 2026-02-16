<script lang="ts">
	import Header from '../Header.svelte';

	let { children } = $props();
</script>

<div class="app">
	<div class="background-accent"></div>

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

	.background-accent {
		position: fixed;
		inset-block-start: 0;
		inset-inline: 0;
		block-size: clamp(10rem, 40vh, 19rem);
		border-bottom-left-radius: var(--radius-md);
		border-bottom-right-radius: var(--radius-md);
		background-color: var(--color-dark-purple);
		z-index: -1;
	}
</style>
