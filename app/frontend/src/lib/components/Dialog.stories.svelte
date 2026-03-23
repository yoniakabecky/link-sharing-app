<script lang="ts" module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Button from './Button.svelte';
	import Dialog from './Dialog.svelte';

	const { Story } = defineMeta({
		title: 'Components/Dialog',
		component: Dialog
	});
	let dialogOpen = $state(false);
</script>

<Button onclick={() => (dialogOpen = true)}>Open Dialog</Button>

<Story name="Default" asChild>
	<Dialog open={dialogOpen} onclose={() => (dialogOpen = false)}>
		<p>This is a dialog content.</p>
		<Button onclick={() => (dialogOpen = false)}>Close</Button>
	</Dialog>
</Story>
