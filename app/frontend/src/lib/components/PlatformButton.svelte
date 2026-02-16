<script lang="ts">
	import Icon, { type IconName } from './Icon.svelte';
	import type { Link } from '$lib/models/link';

	type Props = {
		link: Link;
		size?: 'sm' | 'md';
	};
	let { link, size = 'md' }: Props = $props();
	const iconSize = $derived(size === 'sm' ? 14 : 18);
</script>

<a href={link.url} target="_blank" rel="noopener noreferrer">
	<button data-size={size} style="background-color: {link.platform.color};">
		<div class="platform-info">
			<Icon name={link.platform.icon as IconName} size={iconSize} />
			<div>{link.platform.name}</div>
		</div>
		<Icon name="arrow_right" size={iconSize} />
	</button>
</a>

<style>
	a {
		text-decoration: none;
	}

	button {
		display: flex;
		align-items: center;
		justify-content: space-between;
		border: none;
		border-radius: var(--radius-sm);
		inline-size: 100%;
		block-size: 100%;
		padding-inline: var(--spacing-3);
		color: var(--color-text-inverted);
		cursor: pointer;
	}
	button[data-size='sm'] {
		padding-block: var(--spacing-3);
		font-size: var(--font-size-xs);
	}
	button[data-size='md'] {
		padding-block: var(--spacing-4);
		font-size: var(--font-size-sm);
	}

	.platform-info {
		display: flex;
		align-items: center;
		gap: var(--spacing-2);
	}
</style>
