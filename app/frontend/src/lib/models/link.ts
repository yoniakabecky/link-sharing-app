import * as v from 'valibot';
import type { Platform } from './platform';

export type Link = {
	id: number;
	profile_id: number;
	platform_id: number;
	url: string;
	created_at: string;
	updated_at: string;
	platform: Platform;
};

const updateLinkSchema = v.object({
	id: v.optional(v.string()),
	platform_id: v.pipe(v.string(), v.nonEmpty('Please select a platform.')),
	url: v.pipe(
		v.string(),
		v.nonEmpty('Please enter your url.'),
		v.url('The url is badly formatted.')
	)
});

export const updateLinksSchema = v.object({
	links: v.array(updateLinkSchema)
});

export type UpdateLinks = v.InferOutput<typeof updateLinksSchema>;
