<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import PlatformButton from './PlatformButton.svelte';

	const { Story } = defineMeta({
		title: 'Components/PlatformButton',
		component: PlatformButton
	});

	const link = {
		url: 'https://example.com',
		platform: {
			name: 'Github',
			icon: 'github',
			color: '#000000'
		}
	};
</script>

<Story name="Small" args={{ size: 'sm', link }} />

<Story name="Medium" args={{ size: 'md', link }} />
