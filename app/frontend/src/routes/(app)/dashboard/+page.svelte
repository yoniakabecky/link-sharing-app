<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import Button from '$lib/components/Button.svelte';
	import { getProfiles } from '$lib/remote/profile.remote';
	import { globalState } from '$lib/state.svelte';

	const profiles = await getProfiles();
	const message = $derived(
		profiles.length > 0
			? 'Select a profile to edit or create a new one.'
			: 'No profiles found. Create a new profile to get started.'
	);

	const onSelectProfile = async (profileID: string) => {
		globalState.profileID = profileID;
		// await goto('/profile');
	};
</script>

<main>
	<h1>Hello, <span class="user">{page.data.user?.email}</span></h1>

	<p class="description">{message}</p>

	<ul class="profiles">
		{#each profiles as profile}
			<li>
				<Button
					variant="outlined"
					class="full-width"
					onclick={() => onSelectProfile(profile.id.toString())}
				>
					<span class="nickname">
						{profile.nickname}
					</span>
				</Button>
			</li>
		{/each}
		<li><Button variant="outlined" class="full-width">+ Create A New Profile</Button></li>
		<hr class="full-width" />
		<li><Button variant="outlined" class="full-width">Logout</Button></li>
	</ul>
</main>

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

	.nickname {
		text-transform: uppercase;
	}
</style>
