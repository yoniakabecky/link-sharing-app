<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import PublicProfile from '$lib/components/PublicProfile.svelte';
	import { getProfile } from '$lib/remote/profile.remote';
	import { getProfileID } from '$lib/state.svelte';

	const profileID = getProfileID();
	const { profile } = await getProfile(profileID);
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
		min-height: 0;
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

	@media (max-width: 480px) {
		.background-accent {
			display: none;
		}

		main {
			align-items: stretch;
			justify-content: stretch;
			width: 100%;
		}

		main :global(.card) {
			min-block-size: 100%;
			border-radius: 0;
			box-shadow: none !important;
		}
	}
</style>
