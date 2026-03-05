<script lang="ts" module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Button from './Button.svelte';
	import Dialog from './Dialog.svelte';

	const { Story } = defineMeta({
		title: 'Components/Dialog',
		component: Dialog
	});
	let dialog: ReturnType<typeof Dialog> | null = $state(null);
</script>

<Button onclick={() => dialog?.showModal()}>Open Dialog</Button>

<Story name="Default" asChild>
	<Dialog bind:this={dialog}>
		<p>This is a dialog content.</p>
		<Button onclick={() => dialog?.close()}>Close</Button>
	</Dialog>
</Story>
