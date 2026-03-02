<script lang="ts">
	let { children } = $props();
</script>

<main>
	{@render children()}
</main>

<style>
	main {
		flex: 1;
		display: grid;
		grid-template-columns: 1fr minmax(0, 60%);
		grid-template-rows: minmax(0, 1fr);
		gap: var(--spacing-5);
		padding: var(--spacing-5);
		padding-block-start: 0;
		min-height: 0;
	}

	@media (max-width: 768px) {
		main {
			grid-template-columns: 1fr;
		}
	}
</style>
