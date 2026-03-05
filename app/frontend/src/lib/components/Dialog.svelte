<script lang="ts">
	let { children, class: className = '', preventEscape = false } = $props();
	let dialog: HTMLDialogElement | null = null;

	export const showModal = () => {
		dialog?.showModal();
	};

	export const close = () => {
		dialog?.close();
	};

	$effect(() => {
		dialog?.addEventListener('cancel', (e) => {
			e.preventDefault();
		});
		dialog?.addEventListener('keydown', (e) => {
			if (e.key === 'Escape' && preventEscape) {
				e.preventDefault();
			}
		});
	});
</script>

<dialog bind:this={dialog}>
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
