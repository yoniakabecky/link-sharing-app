<script lang="ts">
	import Icon from '$lib/components/Icon.svelte';
	import TextInput from '$lib/components/TextInput.svelte';

	let { profile } = $props();
</script>

<form id="profile-form">
	<div class="profile-picture">
		<label for="profile-picture-input">Profile Picture</label>

		<div class="profile-picture-wrapper">
			{#if profile.avatar_url}
				<img class="profile-picture-img" src={profile.avatar_url} alt="Your profile" />
			{/if}
			<input type="file" id="profile-picture-input" name="profile-picture" accept="image/*" />
			<div class="profile-picture-overlay">
				<Icon name="pic_line" size={32} />
				<div>Change Image</div>
			</div>
		</div>
		<div>
			<small>Image must be below 1024x1024px.</small>
			<small>Use PNG, JPG, or BMP format.</small>
		</div>
	</div>
	<div class="profile-details">
		<div class="profile-details-item">
			<label for="first-name">First name*</label>
			<TextInput id="first-name" name="first-name" value={profile.first_name} />
		</div>
		<div class="profile-details-item">
			<label for="last-name">Last name*</label>
			<TextInput id="last-name" name="last-name" value={profile.last_name} />
		</div>
		<div class="profile-details-item">
			<label for="email">Email</label>
			<TextInput id="email" name="email" value={profile.email} />
		</div>
	</div>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
	}

	.profile-picture,
	.profile-details {
		display: grid;
		gap: var(--spacing-5);
		border-radius: var(--radius-md);
		background-color: var(--color-canvas);
		padding-inline: var(--spacing-4);
		padding-block: var(--spacing-5);
	}

	label {
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}

	small {
		display: inline-block;
		font-size: var(--font-size-xs);
		color: var(--color-text-secondary);
	}

	.profile-picture {
		grid-template-columns: 30% 1fr auto;
		align-items: center;
	}
	.profile-picture-wrapper {
		position: relative;
		width: 164px;
		height: 164px;
		border-radius: var(--radius-md);
		overflow: hidden;
		background: var(--color-skeleton-gray);
	}

	.profile-picture-img {
		position: absolute;
		inset: 0;
		inline-size: 100%;
		block-size: 100%;
		object-fit: cover;
	}

	.profile-picture-overlay {
		position: absolute;
		inset: 0;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		gap: var(--spacing-2);
		cursor: pointer;
		background-color: rgba(0, 0, 0, 0.5);
		inline-size: 100%;
		block-size: 100%;
		color: var(--color-text-inverted);
		font-weight: 500;
		transition: background-color 0.2s ease-in-out;
	}
	.profile-picture-overlay:hover {
		background-color: rgba(0, 0, 0, 0.3);
	}

	input[type='file'] {
		display: none;
	}

	.profile-details-item {
		display: grid;
		grid-template-columns: 30% 1fr;
		align-items: center;
	}
</style>
