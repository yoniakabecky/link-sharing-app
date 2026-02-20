<script lang="ts">
	import { page } from '$app/state';
	import logo from '$lib/assets/logo.svg';
	import Button from '$lib/components/Button.svelte';
	import Icon from '$lib/components/Icon.svelte';

	let isPreview = $derived(page.url.pathname === '/preview');
	let pathname = $derived(page.url.pathname);
</script>

{#snippet MainNav()}
	<a href="/" class="logo-link">
		<img src={logo} alt="Logo" /><span class="logo-text desktop-only">devlinks</span>
	</a>

	<nav>
		<ul>
			<li>
				<Button variant="subtle" data-active={pathname === '/links'} href="/links">
					<Icon name="link" size={20} />
					<span class="desktop-only">Links</span>
				</Button>
			</li>
			<li>
				<Button variant="subtle" data-active={pathname === '/profile'} href="/profile">
					<Icon name="user" size={20} />
					<span class="desktop-only">Profile Details</span>
				</Button>
			</li>
		</ul>
	</nav>

	<Button variant="outlined" href="/preview">
		<Icon name="eye" size={20} class="mobile-only" />
		<span class="desktop-only">Preview</span>
	</Button>
{/snippet}

{#snippet PreviewNav()}
	<Button variant="outlined" onclick={() => history.back()}>
		<Icon name="arrow_left" size={20} class="mobile-only" />
		<span class="desktop-only">Back to Editor</span>
	</Button>
	<Button href="/">Share Link</Button>
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

	@media (max-width: 480px) {
		header {
			margin: 0;
			margin-block-end: var(--spacing-3);
			padding-inline-start: var(--spacing-4);
			border-radius: 0;
		}
	}
</style>
