<script lang="ts">
	import { getContext, onMount } from 'svelte';
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import Mockup from '$lib/components/Mockup.svelte';
	import { defaultLink } from '$lib/constants/default-link';
	import type { Link } from '$lib/models/link';
	import { getLinks, updateLinks } from '$lib/remote/link.remote';
	import { getPlatforms } from '$lib/remote/platform.remote';
	import LinksForm from './LinksForm.svelte';

	const profileID = getContext('profileID') as string;
	const links = await getLinks(profileID);
	const platforms = await getPlatforms();

	onMount(() => {
		if (links.length === 0) {
			updateLinks.fields.links.set([defaultLink]);
		} else {
			for (let i = 0; i < links.length; i++) {
				updateLinks.fields.links[i].set({
					id: String(links[i].id),
					platform_id: String(links[i].platform_id),
					url: links[i].url
				});
			}
		}
	});

	const displayLinks = $derived.by(() => {
		const formValues = updateLinks.fields.links.value() ?? [];
		return formValues
			.map((link) => ({
				...link,
				platform: platforms.find((p) => p.id === Number(link.platform_id))
			}))
			.filter((link) => link.platform);
	}) as unknown as Partial<Link>[];
</script>

<div class="desktop-only">
	<Card style="block-size: 100%;">
		<Mockup links={displayLinks} />
	</Card>
</div>

<Card>
	{#snippet header()}
		<h2 class="heading">Customize Your Links</h2>
		<p class="description">
			Add/edit/remove links below and then share all profiles with the world!
		</p>
	{/snippet}

	<LinksForm {updateLinks} {platforms} />

	{#snippet footer()}
		<Button type="submit" form="links-form">Save</Button>
	{/snippet}
</Card>

<style>
	.heading {
		margin-block-end: var(--spacing-2);
		font-size: var(--font-size-xl);
	}
	.description {
		margin: 0;
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}

	@media (max-width: 480px) {
		.heading {
			margin-block-start: 0;
			font-size: var(--font-size-lg);
		}
	}
</style>
