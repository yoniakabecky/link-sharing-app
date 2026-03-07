<script lang="ts">
	import { onDestroy } from 'svelte';
	import { toast } from 'svelte-sonner';
	import Icon from '$lib/components/Icon.svelte';
	import TextInput from '$lib/components/TextInput.svelte';
	import type { UpdateProfile } from '$lib/models/profile';
	import type { RemoteForm } from '@sveltejs/kit';

	type Props = {
		updateProfile: RemoteForm<UpdateProfile, { success: boolean }>;
	};

	let { updateProfile }: Props = $props();

	let selectedFile = $state<File | null>(null);
	let blobUrl = $state<string | null>(null);

	const thumbnailUrl = $derived(blobUrl ?? updateProfile.fields.avatar_url.value() ?? '');
	const haveImage = $derived(!!thumbnailUrl);
	const fileInputMessage = $derived(haveImage ? 'Change Image' : '+ Upload Image');

	$effect(() => {
		const file = selectedFile;
		if (file) {
			const url = URL.createObjectURL(file);
			blobUrl = url;
			return () => {
				URL.revokeObjectURL(url);
				blobUrl = null;
			};
		} else {
			blobUrl = null;
		}
	});

	onDestroy(() => {
		if (blobUrl) {
			URL.revokeObjectURL(blobUrl);
		}
	});

	const handleFileChange = (e: Event) => {
		const input = e.target as HTMLInputElement;
		selectedFile = input.files?.[0] ?? null;
	};
</script>

<form
	id="profile-form"
	enctype="multipart/form-data"
	novalidate
	{...updateProfile.enhance(async ({ submit }) => {
		try {
			await submit();
			if (updateProfile.result?.success) {
				selectedFile = null;
				toast.success('Profile updated successfully!');
			} else {
				toast.warning('Please check the fields and try again...');
			}
		} catch (error) {
			console.error(error);
			toast.error('Failed to update profile.');
		}
	})}
>
	<div class="profile-picture">
		<label for="profile-picture-input">Profile Picture</label>

		<div class="profile-picture-wrapper">
			{#if thumbnailUrl}
				<img class="profile-picture-img" src={thumbnailUrl} alt="Your profile" />
			{/if}
			<input
				id="profile-picture-input"
				accept="image/png,image/jpeg,image/bmp"
				{...updateProfile.fields.avatar.as('file')}
				onchange={handleFileChange}
			/>
			<label
				for="profile-picture-input"
				class="profile-picture-overlay"
				data-have-image={haveImage}
			>
				<Icon name="pic_line" size={32} />
				<div>{fileInputMessage}</div>
			</label>
		</div>
		<div>
			<small>Image must be below 1024x1024px.</small>
			<small>Use PNG, JPG, or BMP format.</small>
		</div>
	</div>

	<input type="hidden" {...updateProfile.fields.avatar_url.as('text')} />
	<input type="hidden" {...updateProfile.fields.id.as('text')} />

	<div class="profile-details">
		<div class="profile-details-item">
			<label for="first-name">First name*</label>
			<TextInput id="first-name" {...updateProfile.fields.first_name.as('text')} />
			{#each updateProfile.fields.first_name.issues() as issue (issue.message)}
				<small class="issue">{issue.message}</small>
			{/each}
		</div>
		<div class="profile-details-item">
			<label for="last-name">Last name*</label>
			<TextInput id="last-name" {...updateProfile.fields.last_name.as('text')} />
			{#each updateProfile.fields.last_name.issues() as issue (issue.message)}
				<span></span>
				<small class="issue">{issue.message}</small>
			{/each}
		</div>
		<div class="profile-details-item">
			<label for="email">Email</label>
			<TextInput id="email" {...updateProfile.fields.email.as('email')} />
			{#each updateProfile.fields.email.issues() as issue (issue.message)}
				<span></span>
				<small class="issue">{issue.message}</small>
			{/each}
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

	.profile-picture-overlay:not([data-have-image='true']) {
		background: var(--color-light-purple);
		color: var(--color-dark-purple);
	}

	.profile-picture-overlay:not([data-have-image='true']):hover {
		opacity: 0.8;
	}

	input[type='file'] {
		display: none;
	}

	.profile-details-item {
		display: grid;
		grid-template-columns: 30% 1fr;
		align-items: center;
	}

	.issue {
		margin-block-start: var(--spacing-1);
		color: var(--color-error-red);
	}

	@media (max-width: 768px) {
		.profile-picture {
			grid-template-columns: 1fr;
			gap: var(--spacing-2);
		}
		.profile-details-item {
			grid-template-columns: 1fr;
		}

		.profile-details-item label {
			margin-block-end: var(--spacing-1);
		}
	}
</style>
