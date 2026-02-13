<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		variant?: 'primary' | 'outlined' | 'subtle';
		href?: string;
		children: Snippet;
	} & svelteHTML.IntrinsicElements['button'];

	let { children, variant = 'primary', href, ...props }: Props = $props();
</script>

{#if href}
	<a {href}>
		<button data-variant={variant} {...props}>
			{@render children()}
		</button>
	</a>
{:else}
	<button data-variant={variant} {...props}>
		{@render children()}
	</button>
{/if}

<style>
	button {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		gap: var(--spacing-2);
		padding: var(--spacing-2-5) var(--spacing-5);
		border-radius: var(--radius-sm);
		border: none;
		font-size: var(--font-size-sm);
		font-weight: 600;
		text-decoration: none;
		cursor: pointer;
	}

	button[data-variant='primary'] {
		background-color: var(--color-dark-purple);
		color: var(--color-text-inverted);
	}
	button[data-variant='primary']:hover {
		opacity: 0.9;
	}

	button[data-variant='outlined'] {
		background-color: transparent;
		color: var(--color-dark-purple);
		border: 1px solid var(--color-dark-purple);
	}
	button[data-variant='outlined']:hover {
		background-color: var(--color-light-purple);
	}

	button[data-variant='subtle'] {
		background-color: transparent;
		color: var(--color-text-secondary);
	}
	button[data-variant='subtle']:hover {
		background-color: var(--color-skeleton-gray);
		opacity: 0.8;
	}
	button[data-variant='subtle'][data-active='true'] {
		background-color: var(--color-light-purple);
		color: var(--color-dark-purple);
	}
	button[data-variant='subtle'][data-active='true']:hover {
		opacity: 0.8;
	}
</style>
