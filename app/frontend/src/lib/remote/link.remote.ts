import * as v from 'valibot';
import { error } from '@sveltejs/kit';
import { query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import type { Link } from '$lib/models/link';

export const getLinks = query(v.string(), async (profileID) => {
	const response = await fetch(`${API_BASE_URL}/links/${profileID}`);
	if (!response.ok) {
		error(response.status, `Error fetching links: ${response.statusText}`);
	}
	const data = (await response.json()) ?? [];

	return data as Link[];
});
