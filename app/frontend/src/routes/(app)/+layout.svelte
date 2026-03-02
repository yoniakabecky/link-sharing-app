<script lang="ts">
	import Header from './Header.svelte';

	let { children } = $props();
</script>

<div class="app">
	<Header />

	{@render children()}
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-inline-size: 100vw;
		block-size: 100vh;
		overflow: hidden;
	}
</style>
