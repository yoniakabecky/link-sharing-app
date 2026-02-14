<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import Mockup from '$lib/components/Mockup.svelte';
	import type { Link } from '$lib/models/link';
	import type { Profile } from '$lib/models/profile';

	const dummyProfile = {
		id: 1,
		user_id: 1,
		first_name: 'John',
		last_name: 'Doe',
		email: 'john.doe@example.com',
		avatar_url: 'https://picsum.photos/id/40/100/100'
	} as Profile;

	const dummyLinks = [
		{
			id: 1,
			url: 'https://example.com',
			platform: {
				id: 1,
				name: 'Github',
				icon: 'github',
				color: '#000000'
			}
		}
	] as Link[];
</script>

<Card>
	<Mockup profile={dummyProfile} links={dummyLinks} />
</Card>

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
