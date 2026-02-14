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
