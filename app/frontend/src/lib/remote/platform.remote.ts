import { error } from '@sveltejs/kit';
import { query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import type { Platform } from '$lib/models/platform';

export const getPlatforms = query(async () => {
	const response = await fetch(`${API_BASE_URL}/platforms`);
	if (!response.ok) {
		error(response.status, `Error fetching platforms: ${response.statusText}`);
	}
	const data = (await response.json()) ?? [];

	return data as Platform[];
});
