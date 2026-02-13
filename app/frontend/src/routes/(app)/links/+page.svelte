<script>
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
</script>

<Card>mock</Card>

<Card>
	{#snippet header()}
		<h2 class="heading">Customize Your Links</h2>
		<p class="description">
			Add/edit/remove links below and then share all profiles with the world!
		</p>
	{/snippet}

	form here

	{#snippet footer()}
		<Button>Save</Button>
	{/snippet}
</Card>

<style>
	.heading {
		margin-block-end: var(--spacing-2);
	}
	.description {
		margin: 0;
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}
</style>
