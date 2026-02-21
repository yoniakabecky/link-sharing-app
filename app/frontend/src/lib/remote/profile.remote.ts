import * as v from 'valibot';
import { error, invalid } from '@sveltejs/kit';
import { form, query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import { updateProfileSchema, type Profile } from '$lib/models/profile';

export const getProfile = query(v.string(), async (profileID) => {
	const response = await fetch(`${API_BASE_URL}/profiles/${profileID}`);
	if (!response.ok) {
		error(response.status, `Error fetching profile: ${response.statusText}`);
	}
	const data = await response.json();

	return data as Profile;
});

export const updateProfile = form(updateProfileSchema, async (profile) => {
	try {
		// TODO: replace profileId with the actual profile ID
		const profileId = '7';
		const response = await fetch(`${API_BASE_URL}/profiles/${profileId}`, {
			method: 'PUT',
			body: JSON.stringify(profile)
		});
		if (!response.ok) {
			invalid(`Error updating profile: ${response.statusText}`);
		}
		return { success: true };
	} catch (err) {
		console.error(err);
		error(500, `Error updating links: ${err instanceof Error ? err.message : 'Unknown error'}`);
	}
});
