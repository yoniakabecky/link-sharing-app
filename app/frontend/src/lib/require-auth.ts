import { getRequestEvent } from '$app/server';
import { redirect } from '@sveltejs/kit';

export const requireAuth = () => {
	const { locals, cookies } = getRequestEvent();

	// TODO: add session check

	const token = cookies.get('token');
	if (!token) {
		redirect(307, '/login');
	}

	return {
		user: locals.user,
		token: token
	};
};
