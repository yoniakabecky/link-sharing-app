<script lang="ts">
	import { onMount } from 'svelte';
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import Mockup from '$lib/components/Mockup.svelte';
	import type { Profile } from '$lib/models/profile';
	import { getProfile, updateProfile } from '$lib/remote/profile.remote';
	import { getProfileID } from '$lib/state.svelte';
	import ProfileForm from './ProfileForm.svelte';

	const profileID = getProfileID();
	const profile = await getProfile(profileID);

	onMount(() => {
		if (profile) {
			updateProfile.fields.set({ ...profile, id: String(profile.id) });
		}
	});

	const displayProfile = $derived.by(() => {
		const formValues = updateProfile.fields;
		return {
			id: formValues.id.value(),
			first_name: formValues.first_name.value(),
			last_name: formValues.last_name.value(),
			email: formValues.email.value(),
			avatar_url: formValues.avatar_url.value(),
			links: profile?.links ?? []
		} as unknown as Partial<Profile>;
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
		<div class="footer">
			<Button type="submit" form="profile-form" class="full-width">Save</Button>
		</div>
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

		.footer {
			inline-size: 100%;
		}
	}
</style>
