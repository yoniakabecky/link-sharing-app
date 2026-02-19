import * as v from 'valibot';
import { error } from '@sveltejs/kit';
import { query } from '$app/server';
import { API_BASE_URL } from '$env/static/private';
import type { Profile } from '$lib/models/profile';

export const getProfile = query(v.string(), async (profileID) => {
	const response = await fetch(`${API_BASE_URL}/profiles/${profileID}`);
	if (!response.ok) {
		error(response.status, `Error fetching profile: ${response.statusText}`);
	}
	const data = await response.json();

	return data as Profile;
});
