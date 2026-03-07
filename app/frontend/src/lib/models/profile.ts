import * as v from 'valibot';
import type { Link } from './link';

export type Profile = {
	id: number;
	user_id: number;
	nickname: string;
	first_name: string;
	last_name: string;
	email: string;
	avatar_url: string;
	created_at: string;
	updated_at: string;
	links: Link[];
};

export const updateProfileSchema = v.object({
	id: v.string(),
	first_name: v.pipe(v.string(), v.nonEmpty('Please enter your first name.')),
	last_name: v.pipe(v.string(), v.nonEmpty('Please enter your last name.')),
	email: v.optional(v.pipe(v.string(), v.email('Please enter a valid email address.'))),
	avatar_url: v.optional(v.pipe(v.string(), v.string())),
	avatar: v.optional(v.file())
});

export type UpdateProfile = v.InferOutput<typeof updateProfileSchema>;

export const createProfileSchema = v.object({
	nickname: v.pipe(
		v.string(),
		v.nonEmpty('Please enter your nickname.'),
		v.minLength(2, 'Nickname must be at least 2 characters.')
	),
	first_name: v.pipe(v.string(), v.nonEmpty('Please enter your first name.')),
	last_name: v.pipe(v.string(), v.nonEmpty('Please enter your last name.'))
});
