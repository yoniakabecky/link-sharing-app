<script lang="ts">
	import { page } from '$app/state';
	import Button from '$lib/components/Button.svelte';
	import Card from '$lib/components/Card.svelte';
	import PublicProfile from '$lib/components/PublicProfile.svelte';
	import { getProfile } from '$lib/remote/profile.remote';

	const profileID = page.url.searchParams.get('id') ?? '';
	const profile = await getProfile(profileID);
</script>

{#if profile.success}
	<div class="centered">
		<Card shadow={true}>
			<PublicProfile {profile} />
		</Card>
	</div>
{:else}
	<div class="centered error">
		<h1>Profile Not Found</h1>
		<p>The profile you are looking for does not exist.</p>
		<Button variant="outlined" href="/">Go to home</Button>
	</div>
{/if}

<style>
	.centered {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	.error {
		flex-direction: column;
		gap: var(--spacing-6);
	}

	.error > h1 {
		margin-block: 0;
	}
</style>
