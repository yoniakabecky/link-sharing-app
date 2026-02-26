<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import LogoLink from './LogoLink.svelte';

	const { Story } = defineMeta({
		title: 'Components/LogoLink',
		component: LogoLink
	});
</script>

<Story name="Default" />
