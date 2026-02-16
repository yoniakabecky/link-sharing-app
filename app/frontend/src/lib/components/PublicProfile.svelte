<script lang="ts">
	import PlatformButton from './PlatformButton.svelte';

	let { profile } = $props();
</script>

<div class="public-profile">
	<div class="profile">
		<div class="avatar">
			<img src={profile.avatar_url} alt="Avatar" />
		</div>
		<div class="name">{profile.first_name} {profile.last_name}</div>
		<div class="email">{profile.email}</div>
	</div>

	<div class="links">
		{#each profile.links as link}
			<div class="link">
				<PlatformButton {link} />
			</div>
		{/each}
	</div>
</div>

<style>
	.public-profile {
		padding-block: var(--spacing-5);
		inline-size: 100%;
		block-size: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: var(--spacing-10);
	}

	.profile {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.avatar {
		width: 88px;
		height: 88px;
		border-radius: var(--radius-full);
		overflow: hidden;
		outline: 4px solid var(--color-dark-purple);
		background: var(--color-skeleton-gray);
		margin: var(--spacing-1);
	}

	.avatar img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}

	.name {
		margin-block-start: var(--spacing-7);
		font-size: var(--font-size-xl);
		font-weight: 700;
	}

	.email {
		margin-block-start: var(--spacing-3);
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}

	.links {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-4);
	}

	.link {
		block-size: 48px;
		inline-size: 200px;
	}
</style>
