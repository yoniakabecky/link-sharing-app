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
