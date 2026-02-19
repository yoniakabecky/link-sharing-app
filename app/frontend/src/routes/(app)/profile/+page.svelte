<script lang="ts">
	import { getContext } from 'svelte';
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import Mockup from '$lib/components/Mockup.svelte';
	import { getProfile } from '$lib/remote/profile.remote';
	import ProfileForm from './ProfileForm.svelte';

	const profileID = getContext('profileID') as string;
	const profile = await getProfile(profileID);
</script>

<div class="desktop-only">
	<Card style="block-size: 100%;">
		<Mockup {profile} links={profile.links} />
	</Card>
</div>

<Card>
	{#snippet header()}
		<h2 class="heading">Profile Details</h2>
		<p class="description">Add your details to create a personal touch to your profile.</p>
	{/snippet}

	<ProfileForm {profile} />

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
