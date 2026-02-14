<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import Mockup from './Mockup.svelte';

	const { Story } = defineMeta({
		title: 'Components/Mockup',
		component: Mockup
	});

	const profile = {
		id: 1,
		user_id: 1,
		first_name: 'John',
		last_name: 'Doe',
		email: 'john.doe@example.com',
		avatar_url: 'https://picsum.photos/id/40/100/100'
	};

	const links = [
		{
			id: 1,
			url: 'https://example.com',
			platform: {
				id: 1,
				name: 'Github',
				icon: 'github',
				color: '#000000'
			}
		}
	];
</script>

<Story name="Default" args={{ profile, links }} />

<Story name="Skeleton" args={{ showSkeleton: true }} />
