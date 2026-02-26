<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import TextInput from '$lib/components/TextInput.svelte';

	type Props = {
		buttonText: string;
	};

	let { buttonText }: Props = $props();
</script>

<form>
	<div>
		<TextInput label="Email" placeholder="Email" />
	</div>
	<div>
		<TextInput label="Password" placeholder="Password" />
	</div>
	<Button type="submit">{buttonText}</Button>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
		inline-size: 100%;
	}

	.register-link {
		display: flex;
		justify-content: center;
		gap: var(--spacing-2);
		margin-block: var(--spacing-4);
		color: var(--color-text-secondary);
		font-size: var(--font-size-xs);
	}

	.register-link a {
		color: var(--color-dark-purple);
		text-decoration: none;
	}

	.register-link a:hover {
		text-decoration: underline;
	}
</style>
