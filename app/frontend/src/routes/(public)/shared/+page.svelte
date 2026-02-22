<script lang="ts">
	import { redirect } from '@sveltejs/kit';
	import { page } from '$app/state';
	import Card from '$lib/components/Card.svelte';
	import PublicProfile from '$lib/components/PublicProfile.svelte';
	import { getProfile } from '$lib/remote/profile.remote';

	const profileID = page.url.searchParams.get('id') ?? '';
	// TODO: handle invalid profile ID
	const profile = await getProfile(profileID);
</script>

<div class="centered">
	<Card shadow={true}>
		<PublicProfile {profile} />
	</Card>
</div>

<style>
	.centered {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}
</style>
