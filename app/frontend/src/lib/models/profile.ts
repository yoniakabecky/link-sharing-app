import * as v from 'valibot';
import type { Link } from './link';

export type Profile = {
	id: number;
	user_id: number;
	first_name: string;
	last_name: string;
	email: string;
	avatar_url: string;
	created_at: string;
	updated_at: string;
	links: Link[];
};

export const updateProfileSchema = v.object({
	first_name: v.pipe(v.string(), v.nonEmpty('Please enter your first name.')),
	last_name: v.pipe(v.string(), v.nonEmpty('Please enter your last name.')),
	email: v.optional(v.pipe(v.string(), v.email('Please enter a valid email address.'))),
	avatar_url: v.optional(v.pipe(v.string(), v.url('Please enter a valid URL.'))),
	avatar: v.optional(v.file())
});

export type UpdateProfile = v.InferOutput<typeof updateProfileSchema>;
