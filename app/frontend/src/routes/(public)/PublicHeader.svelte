<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import LogoLink from '$lib/components/LogoLink.svelte';
	import { hasToken } from '$lib/remote/auth.remote';

	let isAuthenticated = await hasToken();
</script>

<header>
	<LogoLink />

	{#if isAuthenticated}
		<Button variant="outlined" href="/dashboard">Go to Dashboard</Button>
	{:else}
		<Button variant="outlined" href="/register">
			<Icon name="user" size={20} />
			<span>Register</span>
		</Button>
	{/if}
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		block-size: 4rem;
		margin: var(--spacing-5);
		background-color: var(--color-card);
		padding: var(--spacing-3);
		border-radius: var(--radius-md);
	}

	@media (max-width: 480px) {
		header {
			margin: 0;
			margin-block-end: var(--spacing-3);
			padding-inline-start: var(--spacing-4);
			border-radius: 0;
		}
	}
</style>
