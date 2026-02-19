<script lang="ts">
	import { getContext } from 'svelte';
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import Mockup from '$lib/components/Mockup.svelte';
	import { getLinks } from '$lib/remote/link.remote';
	import LinksForm from './LinksForm.svelte';

	const profileID = getContext('profileID') as string;
	const links = await getLinks(profileID);
</script>

<div class="desktop-only">
	<Card style="block-size: 100%;">
		<Mockup {links} />
	</Card>
</div>

<Card>
	{#snippet header()}
		<h2 class="heading">Customize Your Links</h2>
		<p class="description">
			Add/edit/remove links below and then share all profiles with the world!
		</p>
	{/snippet}

	<LinksForm {links} />

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
