<script lang="ts">
	import type { Snippet } from 'svelte';
	import { fade } from 'svelte/transition';
	import Icon from './Icon.svelte';

	type Props = {
		open?: boolean;
		header?: Snippet;
		children: Snippet;
		showCloseButton?: boolean;
		direction?: 'left' | 'right';
	};

	let {
		children,
		open = $bindable(false),
		header,
		showCloseButton = true,
		direction = 'right'
	}: Props = $props();

	const close = () => {
		open = false;
	};

	const MIN_WIDTH = 320;
	const DURATION = 250;

	const slide = (node: HTMLElement, params: { direction: 'left' | 'right' }) => {
		let w: number | null = null;
		return {
			duration: DURATION,
			css: (t: number) => {
				if (w === null) w = node.offsetWidth || MIN_WIDTH;
				const x = params.direction === 'right' ? w : -w;
				return `transform: translateX(${x * (1 - t)}px);`;
			}
		};
	};
</script>

{#if open}
	<div
		class="backdrop"
		role="presentation"
		aria-hidden="true"
		onclick={close}
		transition:fade={{ duration: DURATION }}
	></div>
	<div
		class="panel"
		data-direction={direction}
		role="dialog"
		aria-modal="true"
		transition:slide={{ direction }}
	>
		<div class="header-row">
			{#if header}
				<div class="header-slot">
					{@render header()}
				</div>
			{/if}
			{#if showCloseButton}
				<button class="close-btn" aria-label="Close" onclick={close}>
					<Icon name="close" />
				</button>
			{/if}
		</div>

		{@render children()}
	</div>
{/if}

<style>
	.backdrop {
		position: fixed;
		inset: 0;
		z-index: 40;
		background-color: rgba(0, 0, 0, 0.4);
		cursor: pointer;
	}

	.panel {
		position: fixed;
		top: 0;
		left: 0;
		z-index: 41;
		inline-size: max-content;
		min-inline-size: 20rem; /* 320px */
		block-size: 100%;
		background-color: var(--color-card);
		border-right: 1px solid var(--color-border-gray);
		display: flex;
		flex-direction: column;
		overflow: hidden;
	}

	.panel[data-direction='right'] {
		left: auto;
		right: 0;
		border-right: none;
		border-left: 1px solid var(--color-border-gray);
	}

	.header-row {
		position: sticky;
		top: 0;
		display: flex;
		align-items: flex-start;
		justify-content: space-between;
		gap: var(--spacing-4);
		padding-block-start: var(--spacing-5);
		padding-inline: var(--spacing-5);
		flex-shrink: 0;
	}

	.header-slot {
		flex: 1;
		min-width: 0;
	}

	.close-btn {
		position: absolute;
		right: var(--spacing-4);
		top: var(--spacing-4);
		flex-shrink: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: var(--spacing-2);
		border: none;
		background: transparent;
		color: var(--color-text-primary);
		cursor: pointer;
		border-radius: var(--radius-sm);
		transition:
			background-color 0.2s ease,
			color 0.2s ease;
	}
	.close-btn:hover {
		background-color: var(--color-hover-gray);
	}
	.close-btn:focus-visible {
		outline: 2px solid var(--color-dark-purple);
		outline-offset: 2px;
	}
</style>
