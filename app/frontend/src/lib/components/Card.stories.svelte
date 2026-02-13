<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Card from './Card.svelte';

	const { Story } = defineMeta({
		title: 'Components/Card',
		component: Card,
		globals: {
			backgrounds: { value: 'dark' }
		}
	});
</script>

<Story name="Default">Content only</Story>

<Story name="With header and footer" asChild>
	<Card>
		{#snippet header()}
			<h2>Title</h2>
		{/snippet}

		Content here

		{#snippet footer()}
			<p>Footer</p>
		{/snippet}
	</Card>
</Story>
