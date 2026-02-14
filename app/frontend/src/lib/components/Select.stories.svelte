<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Select from './Select.svelte';

	const { Story } = defineMeta({
		title: 'Components/Select',
		component: Select
	});

	const options = [
		{ label: 'Option 1', value: 'option1' },
		{ label: 'Option 2', value: 'option2' },
		{ label: 'Option 3', value: 'option3' }
	];
</script>

<Story name="Default" args={{ options, label: 'Label', placeholder: 'Please select an option' }} />

<Story
	name="Error"
	args={{
		options,
		label: 'Label',
		placeholder: 'Please select an option',
		'aria-invalid': true,
		errorMessage: 'error-message'
	}}
/>
