import * as v from 'valibot';
import { error } from '@sveltejs/kit';
import { form, query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import { updateLinksSchema, type Link } from '$lib/models/link';

export const getLinks = query(v.string(), async (profileID) => {
	const response = await fetch(`${API_BASE_URL}/links/${profileID}`);
	if (!response.ok) {
		error(response.status, `Error fetching links: ${response.statusText}`);
	}
	const data = (await response.json()) ?? [];

	return data as Link[];
});

export const updateLinks = form(updateLinksSchema, async ({ links }) => {
	try {
		const body = links.map((link) => ({
			id: Number(link.id),
			platform_id: Number(link.platform_id),
			url: link.url
		}));
		// TODO: replace profileId with the actual profile ID
		const profileId = '7';
		const response = await fetch(`${API_BASE_URL}/links/${profileId}`, {
			method: 'PUT',
			body: JSON.stringify(body)
		});
		if (!response.ok) {
			error(response.status, `Error updating links: ${response.statusText}`);
		}
	} catch (err) {
		console.error(err);
		error(500, `Error updating links: ${err instanceof Error ? err.message : 'Unknown error'}`);
	}
});
