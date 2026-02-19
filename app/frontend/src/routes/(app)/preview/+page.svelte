<script lang="ts">
	import Card from '$lib/components/Card.svelte';
	import PublicProfile from '$lib/components/PublicProfile.svelte';
	import { getProfile } from '$lib/remote/profile.remote';

	// TODO: replace it with the actual profile ID
	const profileID = '7';
	const profile = await getProfile(profileID);
</script>

<Card shadow={true}>
	<PublicProfile {profile} />
</Card>

<style>
</style>
