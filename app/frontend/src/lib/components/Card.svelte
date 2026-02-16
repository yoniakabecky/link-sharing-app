<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		header?: Snippet;
		children: Snippet;
		footer?: Snippet;
		shadow?: boolean;
	} & svelteHTML.IntrinsicElements['div'];

	let { children, header, footer, shadow = false, ...props }: Props = $props();
</script>

<div class="card" data-shadow={shadow} {...props}>
	{#if header}
		<header>
			{@render header()}
		</header>
	{/if}

	<main>
		{@render children()}
	</main>

	{#if footer}
		<footer>
			{@render footer()}
		</footer>
	{/if}
</div>

<style>
	.card {
		display: flex;
		flex-direction: column;
		background-color: var(--color-card);
		border-radius: var(--radius-md);
	}

	.card header,
	.card main,
	.card footer {
		padding-inline: var(--spacing-8);
		padding-block: var(--spacing-5);
	}

	.card header {
		position: sticky;
		top: 0;
		border-top-left-radius: var(--radius-md);
		border-top-right-radius: var(--radius-md);
		background-color: var(--color-card);
	}

	.card main {
		flex: 1;
		min-height: 0;
		overflow-y: auto;
	}

	.card footer {
		position: sticky;
		bottom: 0;
		display: flex;
		justify-content: flex-end;
		align-items: center;
		border-bottom-left-radius: var(--radius-md);
		border-bottom-right-radius: var(--radius-md);
		background-color: var(--color-card);
		border-top: 1px solid var(--color-border-gray);
	}

	.card[data-shadow='true'] {
		box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.1);
	}
</style>
