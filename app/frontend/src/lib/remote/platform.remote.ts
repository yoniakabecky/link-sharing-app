import { error } from '@sveltejs/kit';
import { query } from '$app/server';
import type { Platform } from '$lib/models/platform';
import { apiGet } from '$lib/fetcher';
import { requireAuth } from '$lib/require-auth';

export const getPlatforms = query(async () => {
	const { token } = requireAuth();
	const response = await apiGet(`/platforms`, token);
	if (!response.ok) {
		error(response.status, `Error fetching platforms: ${response.statusText}`);
	}
	const data = (await response.json()) ?? [];

	return data as Platform[];
});
