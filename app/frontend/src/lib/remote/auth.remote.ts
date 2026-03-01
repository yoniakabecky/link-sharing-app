import { form, getRequestEvent, query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import { authInputSchema } from '$lib/models/auth';
import { invalid, redirect } from '@sveltejs/kit';

export const register = form(authInputSchema, async (data, issue) => {
	const { cookies } = getRequestEvent();

	const response = await fetch(`${API_BASE_URL}/auth/register`, {
		method: 'POST',
		body: JSON.stringify(data)
	});
	if (!response.ok) {
		const message = await response.text();
		if (message.includes('Duplicate entry')) {
			invalid(issue.email('The email is already in use'));
		} else {
			console.error('response error:', message);
			invalid(`Error registering user: ${message}`);
		}
	}
	const { token } = await response.json();
	cookies.set('token', token, { path: '/' });

	redirect(302, '/dashboard');
});

export const login = form(authInputSchema, async (data) => {
	const { cookies } = getRequestEvent();

	const response = await fetch(`${API_BASE_URL}/auth/login`, {
		method: 'POST',
		body: JSON.stringify(data)
	});

	if (!response.ok) {
		console.error('response error:', response.statusText);
		invalid(`Error logging in: ${response.statusText}`);
	}
	const { token } = await response.json();
	cookies.set('token', token, { path: '/' });

	await new Promise((resolve) => setTimeout(resolve, 500));

	redirect(302, '/dashboard');
});

export const hasToken = query(async () => {
	const { cookies } = getRequestEvent();
	const token = cookies.get('token');
	return !!token;
});
