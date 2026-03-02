<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		variant?: 'primary' | 'outlined' | 'subtle' | 'subtle-outlined' | 'danger';
		href?: string;
		children: Snippet;
	} & svelteHTML.IntrinsicElements['button'];

	let { children, variant = 'primary', href, disabled, ...props }: Props = $props();
	let isBusy = $derived(Boolean(props['aria-busy']));
</script>

{#snippet Spinner()}
	<span class="loader" aria-hidden="true">
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
		>
			<path d="M21 12a9 9 0 1 1-6.219-8.56" />
		</svg>
	</span>
{/snippet}

{#snippet ButtonBase()}
	<button data-variant={variant} disabled={Boolean(disabled || isBusy)} {...props}>
		{#if isBusy}
			{@render Spinner()}
		{/if}
		{@render children()}
	</button>
{/snippet}

{#if href}
	<a {href}>
		{@render ButtonBase()}
	</a>
{:else}
	{@render ButtonBase()}
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

	button:disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}

	button[data-variant='primary'] {
		background-color: var(--color-dark-purple);
		color: var(--color-text-inverted);
	}
	button[data-variant='primary']:hover:not(:disabled) {
		opacity: 0.9;
	}
	button[data-variant='outlined'] {
		background-color: transparent;
		color: var(--color-dark-purple);
		border: 1px solid var(--color-dark-purple);
	}
	button[data-variant='outlined']:hover:not(:disabled) {
		background-color: var(--color-light-purple);
	}

	button[data-variant='subtle'] {
		background-color: transparent;
		color: var(--color-text-secondary);
	}
	button[data-variant='subtle']:hover:not(:disabled) {
		background-color: var(--color-skeleton-gray);
		opacity: 0.8;
	}
	button[data-variant='subtle'][data-active='true'] {
		background-color: var(--color-light-purple);
		color: var(--color-dark-purple);
	}
	button[data-variant='subtle'][data-active='true']:hover:not(:disabled) {
		opacity: 0.8;
	}

	button[data-variant='subtle-outlined'] {
		background-color: transparent;
		color: var(--color-text-secondary);
		border: 1px solid var(--color-text-secondary);
	}
	button[data-variant='subtle-outlined']:hover:not(:disabled) {
		background-color: var(--color-skeleton-gray);
		opacity: 0.8;
	}

	button[data-variant='danger'] {
		background-color: transparent;
		color: var(--color-error-red);
		border: 1px solid var(--color-error-red);
	}
	button[data-variant='danger']:hover:not(:disabled) {
		background-color: var(--color-error-light-red);
	}

	.loader {
		display: inline-flex;
		width: 1em;
		height: 1em;
		animation: spin 1s linear infinite;
	}

	.loader svg {
		width: 100%;
		height: 100%;
	}

	@keyframes spin {
		to {
			transform: rotate(360deg);
		}
	}
</style>
