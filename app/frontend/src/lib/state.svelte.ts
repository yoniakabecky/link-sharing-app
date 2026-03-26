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
	const stateID = profileID; // always read $state first so $derived can track it
	if (browser) {
		return sessionStorage.getItem('profileID') || stateID;
	}
	return stateID;
};
