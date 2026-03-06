import * as v from 'valibot';
import { error, invalid } from '@sveltejs/kit';
import { form, query } from '$app/server';
import { updateLinksSchema, type Link } from '$lib/models/link';
import { apiGet, apiPut } from '$lib/fetcher';
import { requireAuth } from '$lib/require-auth';

export const getLinks = query(v.string(), async (profileID) => {
	if (!profileID) return [];

	const { token } = requireAuth();
	const response = await apiGet(`/links/${profileID}`, token);
	if (!response.ok) {
		error(response.status, `Error fetching links: ${response.statusText}`);
	}
	const data = (await response.json()) ?? [];

	return data as Link[];
});

export const updateLinks = form(updateLinksSchema, async ({ profileID, links }) => {
	try {
		const { token } = requireAuth();
		const body = links.map((link) => ({
			id: Number(link.id),
			platform_id: Number(link.platform_id),
			url: link.url
		}));

		const response = await apiPut(`/links/${profileID}`, token, {
			body: JSON.stringify(body)
		});
		if (!response.ok) {
			invalid(`Error updating links: ${response.statusText}`);
		}
		return { success: true };
	} catch (err) {
		console.error(err);
		error(500, `Error updating links: ${err instanceof Error ? err.message : 'Unknown error'}`);
	}
});
