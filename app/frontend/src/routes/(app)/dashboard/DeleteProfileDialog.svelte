<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Dialog from '$lib/components/Dialog.svelte';
	import type { Profile } from '$lib/models/profile';
	import { deleteProfile } from '$lib/remote/profile.remote';
	import { getProfileID, setProfileID } from '$lib/state.svelte';
	import { toast } from 'svelte-sonner';

	let dialog: Dialog;

	type Props = {
		open: boolean;
		profile: Profile | null;
	};
	let { open, profile }: Props = $props();

	const profileID = getProfileID();
	let errorMessage = $state('');

	$effect(() => {
		if (profile && open) {
			dialog?.showModal();
		}
	});
</script>

<Dialog bind:this={dialog}>
	<form
		{...deleteProfile.enhance(async ({ submit }) => {
			try {
				await submit();
				const result = deleteProfile.result;
				if (result?.success) {
					if (profileID === profile?.id.toString()) {
						setProfileID('');
					}
					dialog?.close();
					toast.success('Profile deleted successfully!');
				} else {
					errorMessage = 'Failed to delete profile.';
				}
			} catch (err) {
				console.error(err);
				errorMessage = 'Failed to delete profile.';
			}
		})}
	>
		<h2>Delete Profile</h2>
		<div>
			<p>Are you sure you want to delete this profile?</p>
			<b>Profile Name: {profile?.nickname}</b>
		</div>
		<input type="hidden" name="profileID" value={profile?.id.toString()} />

		{#if errorMessage}
			<p class="error">Error:{errorMessage}</p>
		{/if}

		<div class="buttons">
			<Button variant="subtle" onclick={() => dialog?.close()}>Cancel</Button>
			<Button variant="danger" type="submit">Delete</Button>
		</div>
	</form>
</Dialog>

<style>
	form {
		display: flex;
		flex-direction: column;
	}

	h2 {
		margin-block-end: var(--spacing-2);
		color: var(--color-text-primary);
	}

	p {
		margin-block: var(--spacing-2);
		color: var(--color-text-secondary);
	}

	.buttons {
		display: flex;
		gap: var(--spacing-2);
		justify-content: flex-end;
		margin-block-start: var(--spacing-5);
	}

	.error {
		margin-block: var(--spacing-4);
		color: var(--color-error-red);
		font-size: var(--font-size-sm);
		font-weight: 600;
	}
</style>
