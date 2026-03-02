<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import PublicProfile from '$lib/components/PublicProfile.svelte';
	import { getProfile } from '$lib/remote/profile.remote';
	import { globalState } from '$lib/state.svelte';

	const profile = await getProfile(globalState.profileID);
</script>

<div class="background-accent"></div>

<main>
	<Card shadow={true}>
		<PublicProfile {profile} />
	</Card>
</main>

<style>
	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.background-accent {
		position: fixed;
		inset-block-start: 0;
		inset-inline: 0;
		block-size: clamp(10rem, 40vh, 19rem);
		border-bottom-left-radius: var(--radius-md);
		border-bottom-right-radius: var(--radius-md);
		background-color: var(--color-dark-purple);
		z-index: -1;
	}
</style>
