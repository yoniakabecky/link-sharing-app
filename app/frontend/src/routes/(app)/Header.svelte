<script lang="ts">
	import { page } from '$app/state';
	import logo from '$lib/assets/logo.svg';

	let isPreview = $derived(page.url.pathname === '/preview');
</script>

{#snippet MainNav()}
	<a href="/" class="logo-link">
		<img src={logo} alt="Logo" /><span class="logo-text">devlinks</span>
	</a>

	<nav>
		<ul>
			<li><a href="/links">LINKS</a></li>
			<li><a href="/profile">PROFILE</a></li>
		</ul>
	</nav>
	<div><a href="/preview">PREVIEW</a></div>
{/snippet}

{#snippet PreviewNav()}
	<div><a href="/links">BACK TO EDITOR</a></div>
	<div><a href="/">SHARE LINK</a></div>
{/snippet}

<header>
	{#if isPreview}
		{@render PreviewNav()}
	{:else}
		{@render MainNav()}
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

	.logo-link {
		display: flex;
		align-items: center;
		gap: var(--spacing-1);
		text-decoration: none;
		color: inherit;
	}

	.logo-text {
		font-size: 1.375rem;
		font-weight: 700;
		letter-spacing: -0.03em;
	}

	nav ul {
		display: flex;
		gap: var(--spacing-4);
		list-style: none;
		margin: 0;
		padding: 0;
	}
</style>
