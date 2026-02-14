<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import TextInput from './TextInput.svelte';

	const { Story } = defineMeta({
		title: 'Components/TextInput',
		component: TextInput
	});
</script>

<Story name="Default" args={{ label: 'Label', placeholder: 'Placeholder' }} />

<Story
	name="Error"
	args={{
		label: 'Label',
		value: 'value',
		'aria-invalid': true,
		errorMessage: 'error-message'
	}}
/>

<Story
	name="With left icon"
	args={{ label: 'Label', placeholder: 'Placeholder', leftIcon: 'link' }}
/>
