<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet;
		class?: string;
		preventEscape?: boolean;
		open?: boolean;
		onclose?: () => void;
	} & svelteHTML.IntrinsicElements['dialog'];

	let {
		children,
		class: className = '',
		preventEscape = false,
		open = false,
		onclose,
		...props
	}: Props = $props();
	let dialog: HTMLDialogElement | null = null;

	$effect(() => {
		if (!dialog) return;
		if (open) {
			dialog.showModal();
		} else {
			dialog.close();
		}
	});

	$effect(() => {
		if (!dialog) return;
		const handleClose = () => onclose?.();
		const handleCancel = (e: Event) => e.preventDefault();
		const handleKeydown = (e: KeyboardEvent) => {
			if (e.key === 'Escape' && preventEscape) e.preventDefault();
		};
		dialog.addEventListener('close', handleClose);
		dialog.addEventListener('cancel', handleCancel);
		dialog.addEventListener('keydown', handleKeydown);
		return () => {
			dialog?.removeEventListener('close', handleClose);
			dialog?.removeEventListener('cancel', handleCancel);
			dialog?.removeEventListener('keydown', handleKeydown);
		};
	});
</script>

<dialog bind:this={dialog} class={className} {...props}>
	{@render children()}
</dialog>

<style>
	dialog {
		background-color: var(--color-card);
		border-radius: var(--radius-md);
		padding: var(--spacing-5);
		border: none;
		box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2);
	}

	dialog::backdrop {
		background-color: rgba(0, 0, 0, 0.3);
		backdrop-filter: blur(10px);
	}
</style>
