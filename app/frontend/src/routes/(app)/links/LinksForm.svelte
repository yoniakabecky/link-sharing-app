<script lang="ts">
	import type { RemoteForm } from '@sveltejs/kit';
	import { toast } from 'svelte-sonner';
	import Button from '$lib/components/Button.svelte';
	import Icon, { type IconName } from '$lib/components/Icon.svelte';
	import Select from '$lib/components/Select.svelte';
	import TextInput from '$lib/components/TextInput.svelte';
	import { defaultLink } from '$lib/constants/default-link';
	import type { UpdateLinks } from '$lib/models/link';
	import type { Platform } from '$lib/models/platform';

	type Props = {
		updateLinks: RemoteForm<UpdateLinks, { success: boolean }>;
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
		updateLinks.fields.links[currentLength].set({
			...defaultLink,
			position: currentLength
		});
	};

	const onRemoveLink = (index: number) => {
		const newLinks = updateLinks.fields.links.value().filter((_: any, i: number) => index !== i);
		updateLinks.fields.links.set(newLinks);
	};

	let dragged: number | null = $state(null);
	let dropTarget: number | null = $state(null);

	const onDragEnd = () => {
		dragged = null;
		dropTarget = null;
	};

	const onDragOver = (e: DragEvent, index: number) => {
		e.preventDefault();
		if (dragged === null) return;
		dropTarget = index;
	};

	const onDragLeave = (e: DragEvent) => {
		const related = e.relatedTarget as Node | null;
		const listEl = e.currentTarget as HTMLElement;
		if (related && listEl.contains(related)) return;
		dropTarget = null;
	};

	const insertAt = (insertIdx: number) => {
		if (dragged === null) return;
		const srcIdx = dragged;
		const link = updateLinks.fields.links.value()[srcIdx];
		const newLinks = [...updateLinks.fields.links.value()];
		newLinks.splice(srcIdx, 1);
		const finalIdx = srcIdx < insertIdx ? insertIdx - 1 : insertIdx;
		newLinks.splice(finalIdx, 0, link);
		updateLinks.fields.links.set(newLinks);
		dragged = null;
		dropTarget = null;
	};

	const onDrop = (e: DragEvent, index: number) => {
		e.preventDefault();
		insertAt(index);
	};

	const showPlaceholderBefore = (index: number) => {
		return dropTarget === index;
	};

	const linksLength = $derived(updateLinks.fields.links.value()?.length ?? 0);
</script>

{#snippet dropTargetPlaceholder(targetIndex: number, isEnd: boolean)}
	<li
		class="drop-target"
		aria-hidden="true"
		ondragover={(e) => {
			e.preventDefault();
			if (isEnd) dropTarget = targetIndex;
		}}
		ondrop={(e) => onDrop(e, targetIndex)}
	></li>
{/snippet}

<form
	id="links-form"
	{...updateLinks.enhance(async ({ submit }) => {
		try {
			await submit();
			if (updateLinks.result?.success) {
				toast.success('Links updated successfully!');
			} else {
				toast.warning('Please check the fields and try again...');
			}
		} catch (error) {
			console.error(error);
			toast.error('Failed to update links.');
		}
	})}
>
	<Button type="button" variant="outlined" class="full-width" onclick={onAddLink}>
		+ Add new link
	</Button>

	<ul
		class="links-list"
		role="list"
		aria-label="Links"
		ondragover={(e) => e.preventDefault()}
		ondragleave={(e) => onDragLeave(e)}
	>
		{#each updateLinks.fields.links.value() as _, index}
			{#if showPlaceholderBefore(index)}
				{@render dropTargetPlaceholder(index, false)}
			{/if}
			<li
				draggable="true"
				ondragstart={() => (dragged = index)}
				ondragend={onDragEnd}
				ondragover={(e) => onDragOver(e, index)}
				ondrop={(e) => onDrop(e, index)}
			>
				<fieldset>
					<div class="link-header">
						<div class="link-header-title">
							<span class="drag-icon">
								<Icon name="drag_handle" size={20} />
							</span>
							<legend>Link #{index + 1}</legend>
						</div>
						<button class="remove-link" type="button" onclick={() => onRemoveLink(index)}>
							Remove
						</button>
					</div>

					<input {...updateLinks.fields.links[index].id.as('text')} type="hidden" />
					<input {...updateLinks.fields.links[index].position.as('number')} type="hidden" />

					<div>
						<Select
							label="Platform"
							placeholder="Select a platform"
							options={platformOptions}
							{...updateLinks.fields.links[index].platform_id.as('select')}
						/>
						{#each updateLinks.fields.links[index].platform_id.issues() as issue}
							<small class="issue">{issue.message}</small>
						{/each}
					</div>
					<div>
						<TextInput
							label="URL"
							placeholder="Enter your URL"
							leftIcon="link"
							{...updateLinks.fields.links[index].url.as('text')}
						/>
						{#each updateLinks.fields.links[index].url.issues() as issue}
							<small class="issue">{issue.message}</small>
						{/each}
					</div>
				</fieldset>
			</li>
		{/each}
		{#if dragged !== null}
			{@render dropTargetPlaceholder(linksLength, true)}
		{/if}
	</ul>
</form>

<style>
	form {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
		overflow-y: auto;
	}

	.links-list {
		padding: 0;
		margin: 0;
		list-style: none;
		display: flex;
		flex-direction: column;
		gap: var(--spacing-5);
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

	.issue {
		margin-block-start: var(--spacing-1);
		color: var(--color-error-red);
	}

	.drop-target {
		border: 1px dashed var(--color-border-gray);
		border-radius: var(--radius-sm);
		min-height: 2rem;
	}
</style>
