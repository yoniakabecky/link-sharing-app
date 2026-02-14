<script lang="ts">
	import type { IconName } from './Icon.svelte';
	import Icon from './Icon.svelte';

	type Props = {
		label?: string;
		placeholder?: string;
		leftIcon?: IconName;
		errorMessage?: string;
	} & svelteHTML.IntrinsicElements['input'];

	let uid = $props.id();
	let { label, placeholder, leftIcon, errorMessage, ...props }: Props = $props();
</script>

{#if label}
	<label for={uid}>{label}</label>
{/if}

<div class="input-container">
	{#if leftIcon}
		<div class="icon left">
			<Icon name={leftIcon} size={12} />
		</div>
	{/if}
	<input type="text" {placeholder} id={uid} data-left-icon={!!leftIcon} {...props} />
</div>

{#if !!props['aria-invalid'] && errorMessage}
	<div id={`${uid}-error-message`} class="error-message" aria-live="polite">
		{errorMessage}
	</div>
{/if}

<style>
	label {
		display: inline-block;
		margin-bottom: var(--spacing-1);
		font-size: var(--font-size-xs);
		color: var(--color-text-secondary);
	}

	.input-container {
		position: relative;
	}

	.icon.left {
		position: absolute;
		left: var(--spacing-3);
		top: 50%;
		transform: translateY(-45%);
		color: var(--color-text-secondary);
	}

	input {
		background: var(--color-card);
		border: 1px solid var(--color-border-gray);
		border-radius: var(--radius-sm);
		padding-inline: var(--spacing-3);
		padding-block: var(--spacing-2-5);
		inline-size: 100%;
		font-size: var(--font-size-sm);
		color: var(--color-text-primary);
		box-sizing: border-box;
	}
	input:focus {
		outline: none;
		border-color: var(--color-dark-purple);
		box-shadow: 0px 0px 10px 2px rgba(93, 61, 245, 0.25);
	}
	input::placeholder {
		color: var(--color-text-disabled);
	}

	input[aria-invalid='true'] {
		border-color: var(--color-error-red);
		box-shadow: 0px 0px 10px 2px rgba(255, 99, 71, 0.25);
	}
	input[data-left-icon='true'] {
		padding-left: var(--spacing-8);
	}

	.error-message {
		color: var(--color-error-red);
		font-size: var(--font-size-xs);
		font-weight: 500;
		padding-top: var(--spacing-1);
	}
</style>
