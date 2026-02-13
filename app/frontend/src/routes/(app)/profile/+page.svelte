<script>
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
</script>

<Card>mock</Card>

<Card>
	{#snippet header()}
		<h2 class="heading">Profile Details</h2>
		<p class="description">Add your details to create a personal touch to your profile.</p>
	{/snippet}

	form here

	{#snippet footer()}
		<div class="end">
			<Button>Save</Button>
		</div>
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
