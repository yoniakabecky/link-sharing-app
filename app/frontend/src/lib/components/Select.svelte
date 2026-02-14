<script lang="ts">
	import Icon, { type IconName } from './Icon.svelte';

	type Props = {
		label?: string;
		placeholder?: string;
		errorMessage?: string;
		leftIcon?: IconName;
		options: { label: string; value: string; icon?: IconName }[];
	} & svelteHTML.IntrinsicElements['select'];

	let uid = $props.id();
	let { label, placeholder, errorMessage, options, leftIcon, ...props }: Props = $props();
</script>

{#if label}
	<label for={uid}>{label}</label>
{/if}

<div class="select-container">
	<select id={uid} {...props}>
		<option value="">{placeholder ?? 'Please select an option'}</option>
		{#each options as option}
			<option value={option.value}>
				{#if option.icon}
					<span class="icon" aria-hidden="true">
						<Icon name={option.icon} size={20} />
					</span>
				{/if}
				<span class="option-label">{option.label}</span>
			</option>
		{/each}
	</select>

	<div class="picker-icon">
		<Icon name="down" size={16} />
	</div>
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

	.select-container {
		position: relative;
	}

	.icon {
		pointer-events: none;
		margin-block-start: 2px;
		color: var(--color-text-secondary);
	}
	.picker-icon {
		position: absolute;
		top: 50%;
		right: var(--spacing-3);
		pointer-events: none;
		transform: translateY(-45%);
		color: var(--color-dark-purple);
	}

	select,
	::picker(select) {
		appearance: base-select;
	}
	select {
		background: var(--color-card);
		border: 1px solid var(--color-border-gray);
		border-radius: var(--radius-sm);
		padding-inline: var(--spacing-3);
		padding-block: var(--spacing-2-5);
		inline-size: 100%;
		font-size: var(--font-size-sm);
		transition: 0.4s;
	}
	select:hover {
		background: var(--color-hover-gray);
	}
	select:focus {
		outline: none;
		border-color: var(--color-dark-purple);
		box-shadow: 0px 0px 10px 2px rgba(93, 61, 245, 0.25);
	}
	select::picker-icon {
		display: none;
	}
	select[aria-invalid='true'] {
		border-color: var(--color-error-red);
		box-shadow: 0px 0px 10px 2px rgba(255, 99, 71, 0.25);
	}

	option {
		display: flex;
		justify-content: flex-start;
		align-items: center;
		gap: var(--spacing-5);
		padding-inline: var(--spacing-2);
		padding-block: var(--spacing-1);
		block-size: var(--spacing-8);
	}
	option:hover,
	option:focus {
		background-color: var(--color-hover-gray);
	}
	option:checked {
		font-weight: bold;
		background-color: var(--color-light-purple);
	}
	::picker(select) {
		border: 1px solid var(--color-border-gray);
		border-radius: var(--radius-sm);
	}

	.error-message {
		color: var(--color-error-red);
		font-size: var(--font-size-xs);
		font-weight: 500;
		padding-top: var(--spacing-1);
	}
</style>
