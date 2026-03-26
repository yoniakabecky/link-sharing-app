import * as v from 'valibot';
import { error, invalid } from '@sveltejs/kit';
import { form, query } from '$app/server';
import { createProfileSchema, updateProfileSchema, type Profile } from '$lib/models/profile';
import { apiDelete, apiGet, apiPost, apiPut } from '$lib/fetcher';
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
		return {
			success: false,
			status: response.status,
			message: `Error fetching profile: ${response.statusText}`
		};
	}
	const data = await response.json();

	return {
		success: true,
		status: response.status,
		profile: data as Profile
	};
});

const uploadAvatar = async (profileID: string, avatar: File, token: string) => {
	const formData = new FormData();
	formData.append('avatar', avatar);
	const response = await apiPost(`/profiles/${profileID}/avatar`, token, { body: formData });
	if (!response.ok) {
		console.error(response.statusText);
		invalid(`Failed to upload avatar: ${response.statusText}`);
	}
	const { avatar_url } = await response.json();
	return avatar_url as string;
};

export const updateProfile = form(updateProfileSchema, async (profile) => {
	try {
		const { token } = requireAuth();
		const { id, avatar, ...profileData } = profile;
		let avatar_url = profileData.avatar_url;

		if (avatar && typeof avatar === 'object' && 'size' in avatar && avatar.size > 0) {
			avatar_url = await uploadAvatar(id, avatar as File, token);
		}

		const response = await apiPut(`/profiles/${id}`, token, {
			body: JSON.stringify({
				...profileData,
				avatar_url
			})
		});
		if (!response.ok) {
			invalid(`Error updating profile: ${response.statusText}`);
		}
		return { success: true };
	} catch (err) {
		console.error(err);
		error(500, `Error updating profile: ${err instanceof Error ? err.message : 'Unknown error'}`);
	}
});

export const createProfile = form(createProfileSchema, async (data) => {
	try {
		const { token } = requireAuth();
		const response = await apiPost('/profiles', token, {
			body: JSON.stringify(data)
		});
		if (!response.ok) {
			invalid(`Error creating profile: ${response.statusText}`);
		}
		const newProfile = await response.json();

		return { success: true, profile: newProfile };
	} catch (err) {
		console.error(err);
		error(500, `Error creating profile: ${err instanceof Error ? err.message : 'Unknown error'}`);
	}
});

export const deleteProfile = form('unchecked', async ({ profileID }) => {
	try {
		const { token } = requireAuth();
		const response = await apiDelete(`/profiles/${profileID}`, token);
		if (!response.ok) {
			invalid(`Error deleting profile: ${response.statusText}`);
		}

		return { success: true };
	} catch (err) {
		console.error(err);
		error(500, `Error deleting profile: ${err instanceof Error ? err.message : 'Unknown error'}`);
	}
});
