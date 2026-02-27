<script lang="ts">
	import type { RemoteFormFields } from '@sveltejs/kit';
	import Button from '$lib/components/Button.svelte';
	import TextInput from '$lib/components/TextInput.svelte';
	import type { AuthInput } from '$lib/models/auth';

	type Props = {
		buttonText: string;
		pending: boolean;
		fields: RemoteFormFields<AuthInput>;
	} & svelteHTML.IntrinsicElements['form'];

	let showPassword = $state(false);

	let { buttonText, fields, pending, ...props }: Props = $props();
</script>

<form novalidate {...props}>
	<div>
		<TextInput label="Email" placeholder="Email" {...fields?.email?.as('email')} />
		{#each fields?.email?.issues() as issue}
			<small class="issue">{issue.message}</small>
		{/each}
	</div>
	<div>
		<TextInput
			label="Password"
			placeholder="Password"
			{...fields?.password?.as('password')}
			type={showPassword ? 'text' : 'password'}
		/>
		{#each fields?.password?.issues() as issue}
			<small class="issue">{issue.message}</small>
		{/each}
		<div class="show-password">
			<input
				type="checkbox"
				id="show-password"
				value={showPassword}
				onchange={() => (showPassword = !showPassword)}
			/>
			<label for="show-password">Show password</label>
		</div>
	</div>
	<Button type="submit" aria-busy={pending}>{buttonText}</Button>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
		inline-size: 100%;
	}

	.show-password {
		display: flex;
		gap: var(--spacing-1);
		margin-inline: var(--spacing-1);
		margin-block-start: var(--spacing-2);
		font-size: var(--font-size-xs);
		color: var(--color-text-secondary);
	}

	.issue {
		margin-block-start: var(--spacing-1);
		margin-inline-start: var(--spacing-1);
		color: var(--color-error-red);
	}
</style>
