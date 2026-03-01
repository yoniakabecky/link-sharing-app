import { getRequestEvent } from '$app/server';
import { redirect } from '@sveltejs/kit';

export const requireAuth = () => {
	const { locals, cookies } = getRequestEvent();

	if (!locals.user) {
		redirect(307, '/login');
	}

	const token = cookies.get('token');
	return {
		user: locals.user,
		token: token ?? ''
	};
};
