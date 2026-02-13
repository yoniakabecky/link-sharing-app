<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Button from './Button.svelte';

	const { Story } = defineMeta({
		title: 'Components/Button',
		component: Button
	});
</script>

<Story name="Primary" args={{ 'data-variant': 'primary' }}>Primary Button</Story>

<Story name="Outlined" args={{ 'data-variant': 'outlined' }}>Outlined Button</Story>

<Story name="Subtle" args={{ 'data-variant': 'subtle' }}>Subtle Button</Story>

<Story name="Subtle Active" args={{ 'data-variant': 'subtle', 'data-active': true }}>
	Subtle Active Button
</Story>
