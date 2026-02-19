<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import Icon, { type IconName } from '$lib/components/Icon.svelte';
	import Select from '$lib/components/Select.svelte';
	import TextInput from '$lib/components/TextInput.svelte';
	import type { Link } from '$lib/models/link';
	import { getPlatforms } from '$lib/remote/platform.remote';

	const defaultLink = {
		url: '',
		platform_id: ''
	};

	let { links } = $props();
	const platforms = await getPlatforms();
	const platformOptions = $derived(
		platforms.map((platform) => ({
			label: platform.name,
			value: platform.id.toString(),
			icon: platform.icon as IconName
		}))
	);

	$effect(() => {
		if (links.length === 0) {
			links = [defaultLink];
		}
	});

	const onAddLink = () => {
		links.push(defaultLink);
	};
	const onRemoveLink = (index: number) => {
		links = links.filter((_: Link, i: number) => index !== i);
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
					options={platformOptions}
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
