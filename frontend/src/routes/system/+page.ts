import { goto } from "$app/navigation";
import { getProfile, getTasks } from "$lib/api";
import type { PageLoad } from "./$types"

export const load = (async () => {
	const profile = await getProfile();
	if (!profile.status) {
		return goto("/register");
	}

	const tasks = await getTasks();
	if (!tasks.status) {
		return goto("/register");
	}


	return {
		profile: profile.data,
		tasks: tasks.data,
	}

}) satisfies PageLoad
