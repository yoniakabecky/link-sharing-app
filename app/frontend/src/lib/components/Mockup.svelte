<script lang="ts">
	import mock from '$lib/assets/mobile-mock.svg';
	import type { Link } from '$lib/models/link';
	import type { Profile } from '$lib/models/profile';
	import PlatformButton from './PlatformButton.svelte';

	type Props = {
		profile?: Profile;
		links?: Partial<Link>[];
		showSkeleton?: boolean;
	};
	let { profile, links, showSkeleton }: Props = $props();
	const maxDisplayLinks = 5;
</script>

{#snippet displayProfile(p: Profile)}
	<div class="profile">
		<div class="avatar">
			<img src={p.avatar_url} alt="Avatar" />
		</div>
		<div class="name">{p.first_name} {p.last_name}</div>
		<div class="email">{p.email}</div>
	</div>
{/snippet}

{#snippet profileSkeleton()}
	<div class="skeleton">
		<div class="avatar"></div>
		<div class="name"></div>
		<div class="email"></div>
	</div>
{/snippet}

<div class="mockup">
	<div class="mobile-mockup">
		<img src={mock} alt="Mobile Mockup" />
	</div>
	<div class="content">
		{#if !profile || showSkeleton}
			{@render profileSkeleton()}
		{:else}
			{@render displayProfile(profile)}
		{/if}
		<div class="links">
			{#if links && links.length > 0}
				{#each links as link}
					<div class="link">
						<PlatformButton link={link as Link} size="sm" />
					</div>
				{/each}
			{:else if showSkeleton && (links ?? []).length < maxDisplayLinks}
				{#each Array(maxDisplayLinks - (links ?? []).length) as _}
					<div class="link skeleton"></div>
				{/each}
			{/if}
		</div>
	</div>
</div>

<style>
	.mockup {
		position: relative;
		display: flex;
		justify-content: center;
		align-items: center;
		block-size: 100%;
		inline-size: 100%;
		min-block-size: 540px;
		min-inline-size: 280px;
	}

	.mobile-mockup {
		position: absolute;
	}

	.content {
		block-size: 474px;
		inline-size: 240px;
		padding: var(--spacing-5);
		z-index: 1;
	}

	.profile,
	.links,
	.skeleton {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.avatar {
		width: 80px;
		height: 80px;
		border-radius: var(--radius-full);
		overflow: hidden;
		outline: 3px solid var(--color-dark-purple);
		background: var(--color-skeleton-gray);
		margin: var(--spacing-1);
	}

	.avatar img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.name {
		margin-block-start: var(--spacing-4);
		font-size: var(--font-size-lg);
		font-weight: 700;
	}

	.email {
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
		margin-block-start: var(--spacing-2);
	}

	.skeleton .avatar {
		outline: none;
		background: var(--color-skeleton-gray);
	}

	.skeleton .name {
		inline-size: 70%;
		block-size: var(--spacing-4);
		border-radius: var(--radius-full);
		background: var(--color-skeleton-gray);
	}

	.skeleton .email {
		inline-size: 30%;
		block-size: var(--spacing-2);
		border-radius: var(--radius-full);
		background: var(--color-skeleton-gray);
	}

	.links {
		gap: var(--spacing-4);
		margin-block-start: var(--spacing-9);
	}

	.link {
		block-size: 38px;
		inline-size: 200px;
	}

	.link.skeleton {
		background: var(--color-skeleton-gray);
	}
</style>
