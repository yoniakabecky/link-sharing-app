<script>
	import { goto } from '$app/navigation';
	import Button from '$lib/components/Button.svelte';
	import TextInput from '$lib/components/TextInput.svelte';
	import { createProfile } from '$lib/remote/profile.remote';
	import { globalState } from '$lib/state.svelte';
	import { toast } from 'svelte-sonner';
</script>

<form
	novalidate
	{...createProfile.enhance(async ({ submit, form }) => {
		try {
			await submit();
			const result = createProfile.result;
			if (result?.success) {
				form.reset();
				globalState.profileID = String(result.profile.id);
				toast.success('Profile created successfully!');
				goto('/profile');
			} else {
				toast.error('Failed to create profile.');
			}
		} catch (err) {
			console.error(err);
			toast.error('Failed to create profile.');
		}
	})}
>
	<div>
		<TextInput
			label="Nickname*"
			placeholder="Enter your nickname"
			required
			{...createProfile.fields.nickname.as('text')}
		/>
		{#each createProfile.fields.nickname.issues() as issue}
			<small class="issue">{issue.message}</small>
		{/each}
	</div>
	<div>
		<TextInput
			label="First Name*"
			placeholder="Enter your first name"
			required
			{...createProfile.fields.first_name.as('text')}
		/>
		{#each createProfile.fields.first_name.issues() as issue}
			<small class="issue">{issue.message}</small>
		{/each}
	</div>
	<div>
		<TextInput
			label="Last Name*"
			placeholder="Enter your last name"
			required
			{...createProfile.fields.last_name.as('text')}
		/>
		{#each createProfile.fields.last_name.issues() as issue}
			<small class="issue">{issue.message}</small>
		{/each}
	</div>

	<Button type="submit" class="w-full">Create</Button>
</form>

<style>
	form {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
		overflow-y: auto;
		padding: var(--spacing-5);
		min-inline-size: 24rem;
	}

	.issue {
		margin-block-start: var(--spacing-1);
		color: var(--color-error-red);
	}
</style>
