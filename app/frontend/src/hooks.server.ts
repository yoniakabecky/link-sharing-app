import { apiGet, apiPost } from '$lib/fetcher';
import type { AuthResponse } from '$lib/models/auth';
import type { Handle, RequestEvent } from '@sveltejs/kit';

const setSessionFromTokens = async (event: RequestEvent, data: AuthResponse) => {
	event.cookies.set('token', data.access_token, { path: '/' });
	event.cookies.set('refresh_token', data.refresh_token, { path: '/' });
	event.locals.user = data.user;
};

const clearSession = (event: RequestEvent) => {
	event.cookies.delete('token', { path: '/' });
	event.cookies.delete('refresh_token', { path: '/' });
	event.locals.user = null;
};

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.user = null;

	const accessToken = event.cookies.get('token');
	const refreshToken = event.cookies.get('refresh_token');

	if (!accessToken && !refreshToken) {
		return await resolve(event);
	}

	// Validate session with access token
	if (accessToken) {
		try {
			const res = await apiGet('/auth/session', accessToken);
			if (res.ok) {
				const data = (await res.json()) as AuthResponse;
				event.locals.user = data.user;
				return await resolve(event);
			}
		} catch {
			clearSession(event);
			return await resolve(event);
		}
	}

	// Try refresh
	if (refreshToken) {
		try {
			const res = await apiPost('/auth/refresh', refreshToken, {
				body: JSON.stringify({ refresh_token: refreshToken })
			});
			if (res.ok) {
				const data = (await res.json()) as AuthResponse;
				await setSessionFromTokens(event, data);
				return await resolve(event);
			}
		} catch {
			clearSession(event);
		}
	} else {
		clearSession(event);
	}

	return await resolve(event);
};
