<script module>
	import { defineMeta } from '@storybook/addon-svelte-csf';
	import PublicProfile from './PublicProfile.svelte';

	const { Story } = defineMeta({
		title: 'Components/PublicProfile',
		component: PublicProfile
	});

	const profile = {
		id: 1,
		user_id: 1,
		first_name: 'John',
		last_name: 'Doe',
		email: 'john.doe@example.com',
		avatar_url: 'https://picsum.photos/id/40/100/100',
		links: [
			{
				id: 1,
				profile_id: 1,
				platform_id: 1,
				url: 'https://example.com',
				platform: {
					id: 1,
					name: 'Github',
					icon: 'github',
					color: '#000000'
				}
			},
			{
				id: 2,
				profile_id: 1,
				platform_id: 2,
				url: 'https://example.com',
				platform: {
					id: 2,
					name: 'LinkedIn',
					icon: 'linkedin',
					color: '#0077B5'
				}
			}
		]
	};
</script>

<Story name="Default" args={{ profile }} />
