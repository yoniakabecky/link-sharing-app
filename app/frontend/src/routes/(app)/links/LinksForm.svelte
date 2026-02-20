<script lang="ts">
	import type { RemoteForm } from '@sveltejs/kit';
	import Button from '$lib/components/Button.svelte';
	import Icon, { type IconName } from '$lib/components/Icon.svelte';
	import Select from '$lib/components/Select.svelte';
	import TextInput from '$lib/components/TextInput.svelte';
	import { defaultLink } from '$lib/constants/default-link';
	import type { UpdateLinks } from '$lib/models/link';
	import type { Platform } from '$lib/models/platform';

	type Props = {
		updateLinks: RemoteForm<UpdateLinks, void>;
		platforms: Platform[];
	};
	let { updateLinks, platforms }: Props = $props();

	const platformOptions = $derived(
		platforms.map((platform) => ({
			label: platform.name,
			value: platform.id.toString(),
			icon: platform.icon as IconName
		}))
	);

	const onAddLink = () => {
		const currentLength = updateLinks.fields.links.value().length;
		updateLinks.fields.links[currentLength].set(defaultLink);
	};

	const onRemoveLink = (index: number) => {
		const newLinks = updateLinks.fields.links.value().filter((_: any, i: number) => index !== i);
		updateLinks.fields.links.set(newLinks);
	};
</script>

<form
	id="links-form"
	{...updateLinks.enhance(async ({ submit }) => {
		await submit();
	})}
>
	<Button type="button" variant="outlined" class="full-width" onclick={onAddLink}>
		+ Add new link
	</Button>

	{#each updateLinks.fields.links.value() as _, index}
		<fieldset>
			<div class="link-header">
				<div class="link-header-title">
					<span class="drag-icon">
						<Icon name="drag_handle" size={20} />
					</span>
					<legend>Link #{index + 1}</legend>
				</div>
				<button class="remove-link" type="button" onclick={() => onRemoveLink(index)}>Remove</button
				>
			</div>

			<input type="hidden" {...updateLinks.fields.links[index].id.as('text')} />

			<div>
				<Select
					label="Platform"
					placeholder="Select a platform"
					options={platformOptions}
					{...updateLinks.fields.links[index].platform_id.as('select')}
				/>
			</div>
			<div>
				<TextInput
					label="URL"
					placeholder="Enter your URL"
					leftIcon="link"
					{...updateLinks.fields.links[index].url.as('text')}
				/>
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
