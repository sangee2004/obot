import { ChatService, getProfile, type AuthProvider } from '$lib/services';
import { Role } from '$lib/services/admin/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	let authProviders: AuthProvider[] = [];
	let profile;

	try {
		profile = await getProfile({ fetch });
	} catch (_err) {
		// unauthorized, no need to do anything with error
		authProviders = await ChatService.listAuthProviders({ fetch });
	}

	return {
		loggedIn: profile?.loaded ?? false,
		isAdmin: profile?.role === Role.ADMIN,
		authProviders
	};
};
