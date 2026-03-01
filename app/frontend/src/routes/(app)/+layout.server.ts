import { redirect } from '@sveltejs/kit';

export function load({ locals }) {
	if (!locals.user) {
		redirect(307, '/login');
	}
	return { user: locals.user };
}
