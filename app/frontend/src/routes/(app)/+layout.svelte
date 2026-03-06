<script lang="ts">
	import { getProfileID, setProfileID } from '$lib/state.svelte';
	import { onMount } from 'svelte';
	import Header from './Header.svelte';

	let { children } = $props();

	onMount(() => {
		const profileID = getProfileID();
		const sessionProfileID = sessionStorage.getItem('profileID');
		if (profileID !== sessionProfileID && !profileID && sessionProfileID) {
			setProfileID(sessionProfileID);
		}
	});
</script>

<div class="app">
	<Header />

	{@render children()}
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-inline-size: 100vw;
		block-size: 100vh;
		overflow: hidden;
	}
</style>
