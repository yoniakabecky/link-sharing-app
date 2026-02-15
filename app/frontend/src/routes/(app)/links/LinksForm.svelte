<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import Select from '$lib/components/Select.svelte';
	import TextInput from '$lib/components/TextInput.svelte';

	// TODO: get platforms from API
	const platforms = [
		{ label: 'Github', value: 'github' },
		{ label: 'Twitter', value: 'twitter' },
		{ label: 'LinkedIn', value: 'linkedin' },
		{ label: 'Instagram', value: 'instagram' },
		{ label: 'Facebook', value: 'facebook' }
	];

	const defaultLink = {
		url: '',
		platform_id: ''
	};

	// TODO: get links from API, if links is empty add a empty link object
	let links = $state([defaultLink]);

	const onAddLink = () => {
		links.push(defaultLink);
	};
	const onRemoveLink = (index: number) => {
		links = links.filter((_, i) => index !== i);
	};
</script>

<form id="links-form">
	<Button type="button" variant="outlined" class="full-width" onclick={onAddLink}>
		+ Add new link
	</Button>

	{#each links as link, index}
		<fieldset>
			<div class="link-header">
				<div class="link-header-title">
					<span class="drag-icon">
						<Icon name="drag_handle" size={20} />
					</span>
					<legend>Link #{index + 1}</legend>
				</div>
				<button class="remove-link" onclick={() => onRemoveLink(index)}>Remove</button>
			</div>

			<div>
				<Select
					label="Platform"
					placeholder="Select a platform"
					options={platforms}
					value={link.platform_id}
				/>
			</div>
			<div>
				<TextInput label="URL" placeholder="Enter your URL" leftIcon="link" value={link.url} />
			</div>
		</fieldset>
	{/each}
</form>

<style>
	form {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
		overflow-y: auto;
	}

	fieldset {
		display: flex;
		flex-direction: column;
		gap: var(--spacing-3);
		border: none;
		border-radius: var(--radius-md);
		background-color: var(--color-canvas);
		padding-inline: var(--spacing-4);
		padding-block: var(--spacing-5);
	}

	.link-header {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: var(--spacing-2);
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}

	.link-header-title {
		display: flex;
		align-items: center;
		gap: var(--spacing-1);
		font-weight: 700;
	}

	.drag-icon {
		block-size: var(--spacing-5);
		cursor: grab;
	}

	.remove-link {
		background: none;
		border: none;
		padding: 0;
		margin: 0;
		cursor: pointer;
		font-size: var(--font-size-sm);
		color: var(--color-text-secondary);
	}
	.remove-link:focus {
		outline: none;
	}
	.remove-link:hover {
		text-decoration: underline;
	}
</style>
