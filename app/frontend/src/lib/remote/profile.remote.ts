import * as v from 'valibot';
import { error, invalid } from '@sveltejs/kit';
import { form, query } from '$app/server';
import { updateProfileSchema, type Profile } from '$lib/models/profile';
import { apiGet, apiPut } from '$lib/fetcher';
import { requireAuth } from '$lib/require-auth';

export const getProfiles = query(async () => {
	const { token } = requireAuth();
	const response = await apiGet('/profiles', token);
	if (!response.ok) {
		error(response.status, `Error fetching profiles: ${response.statusText}`);
	}
	const data = (await response.json()) ?? [];

	return data as Profile[];
});

export const getProfile = query(v.string(), async (profileID) => {
	const { token } = requireAuth();
	const response = await apiGet(`/profiles/${profileID}`, token);
	if (!response.ok) {
		error(response.status, `Error fetching profile: ${response.statusText}`);
	}
	const data = await response.json();

	return data as Profile;
});

export const updateProfile = form(updateProfileSchema, async (profile) => {
	try {
		const { token } = requireAuth();
		// TODO: replace profileId with the actual profile ID
		const profileId = '7';
		const response = await apiPut(`/profiles/${profileId}`, token, {
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
