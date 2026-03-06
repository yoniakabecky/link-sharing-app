import { browser } from '$app/environment';

let profileID = $state('');

export const setProfileID = (id: string) => {
	if (id !== profileID) profileID = id;

	if (browser && id) {
		sessionStorage.setItem('profileID', id);
	} else if (browser) {
		sessionStorage.removeItem('profileID');
	}
};

export const getProfileID = () => {
	if (browser) {
		return sessionStorage.getItem('profileID') || profileID;
	}
	return profileID;
};
