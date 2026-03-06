<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import Button from '$lib/components/Button.svelte';
	import Drawer from '$lib/components/Drawer.svelte';
	import type { Profile } from '$lib/models/profile';
	import { logout } from '$lib/remote/auth.remote';
	import { getProfiles } from '$lib/remote/profile.remote';
	import { getProfileID, setProfileID } from '$lib/state.svelte';
	import CreateProfileForm from './CreateProfileForm.svelte';
	import DeleteProfileDialog from './DeleteProfileDialog.svelte';

	const profiles = $derived(await getProfiles());
	const message = $derived(
		profiles.length > 0
			? 'Select a profile to edit or create a new one.'
			: 'No profiles found. Create a new profile to get started.'
	);

	const profileID = getProfileID();

	let open = $state(false);
	let deletingProfile = $state<Profile | null>(null);
	let isDialogOpen = $state(false);

	const onSelectProfile = async (id: string) => {
		setProfileID(id);
		await goto(resolve('/profile'));
	};

	const onDeleteProfile = (profile: Profile) => {
		deletingProfile = profile;
		isDialogOpen = true;
	};
</script>

<main>
	<h1>Hello, <span class="user">{page.data.user?.email}</span></h1>

	<p class="description">{message}</p>

	<ul class="profiles">
		{#each profiles as profile (profile.id)}
			<li class="profile-item">
				<Button
					variant={profile.id.toString() === profileID ? 'primary' : 'outlined'}
					class="full-width"
					onclick={() => onSelectProfile(profile.id.toString())}
				>
					<span class="nickname">
						{profile.nickname}
					</span>
				</Button>
				<Button variant="danger" onclick={() => onDeleteProfile(profile)}>Delete</Button>
			</li>
		{/each}

		<li class="create-button">
			<Button variant="subtle-outlined" class="full-width" onclick={() => (open = !open)}>
				+ Create A New Profile
			</Button>
		</li>
		<hr class="full-width" />
		<li>
			<form {...logout}>
				<Button variant="danger" class="full-width" type="submit">Logout</Button>
			</form>
		</li>
	</ul>
</main>

<Drawer {open}>
	{#snippet header()}
		<h2>Create A New Profile</h2>
	{/snippet}

	<CreateProfileForm />
</Drawer>

<DeleteProfileDialog open={isDialogOpen} profile={deletingProfile} />

<style>
	main {
		display: flex;
		align-items: center;
		flex-direction: column;
		gap: var(--spacing-5);
		padding: var(--spacing-5);
	}

	h1 {
		margin-block-start: 0;
	}

	.description {
		margin: 0;
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}

	.profiles {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
		inline-size: clamp(20rem, 40%, 40rem);
		list-style: none;
		padding-inline-start: 0;
	}

	.profile-item {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: var(--spacing-2);
	}

	.nickname {
		text-transform: uppercase;
	}

	.create-button {
		margin-block-start: var(--spacing-5);
	}
</style>
