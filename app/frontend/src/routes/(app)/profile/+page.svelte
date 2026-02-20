<script lang="ts">
	import { getContext, onMount } from 'svelte';
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import Mockup from '$lib/components/Mockup.svelte';
	import { getProfile, updateProfile } from '$lib/remote/profile.remote';
	import ProfileForm from './ProfileForm.svelte';
	import type { UpdateProfile } from '$lib/models/profile';

	const profileID = getContext('profileID') as string;
	const profile = await getProfile(profileID);

	onMount(() => {
		updateProfile.fields.set(profile);
	});

	const displayProfile = $derived.by(() => {
		const formValues = updateProfile.fields;
		return {
			first_name: formValues.first_name.value(),
			last_name: formValues.last_name.value(),
			email: formValues.email.value(),
			avatar_url: formValues.avatar_url.value(),
			links: profile.links
		} as UpdateProfile;
	});
</script>

<div class="desktop-only">
	<Card style="block-size: 100%;">
		<Mockup profile={displayProfile} />
	</Card>
</div>

<Card>
	{#snippet header()}
		<h2 class="heading">Profile Details</h2>
		<p class="description">Add your details to create a personal touch to your profile.</p>
	{/snippet}

	<ProfileForm {updateProfile} />

	{#snippet footer()}
		<div class="end">
			<Button type="submit" form="profile-form">Save</Button>
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
