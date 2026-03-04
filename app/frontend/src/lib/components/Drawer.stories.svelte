<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Button from './Button.svelte';
	import Drawer from './Drawer.svelte';

	const { Story } = defineMeta({
		title: 'Components/Drawer',
		component: Drawer
	});

	let open = $state(false);
</script>

<Story name="Default" asChild>
	<div>
		<Button onclick={() => (open = true)}>Open Drawer</Button>
		<Drawer bind:open>
			<p style="width: 320px;">Drawer content. 320px width.</p>
		</Drawer>
	</div>
</Story>

<Story name="With header" asChild>
	<div>
		<Button onclick={() => (open = true)}>Open Drawer</Button>
		<Drawer bind:open>
			{#snippet header()}
				<h3>Header</h3>
			{/snippet}
			<p style="width: 320px;">Drawer content. 320px width.</p>
		</Drawer>
	</div>
</Story>
