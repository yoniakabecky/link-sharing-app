import { form, getRequestEvent, query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import { authInputSchema, type AuthResponse } from '$lib/models/auth';
import { invalid, redirect, type Cookies } from '@sveltejs/kit';

const setAuthCookies = (cookies: Cookies, data: AuthResponse) => {
	cookies.set('token', data.access_token, { path: '/' });
	cookies.set('refresh_token', data.refresh_token, { path: '/' });
};

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
	const body = await response.json();
	setAuthCookies(cookies, body);
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
	const body = await response.json();
	setAuthCookies(cookies, body);
	redirect(302, '/dashboard');
});

export const logout = form(async () => {
	const { cookies } = getRequestEvent();
	cookies.delete('token', { path: '/' });
	cookies.delete('refresh_token', { path: '/' });
	redirect(302, '/');
});

export const hasToken = query(async () => {
	const { cookies } = getRequestEvent();
	const token = cookies.get('token');
	return !!token;
});
