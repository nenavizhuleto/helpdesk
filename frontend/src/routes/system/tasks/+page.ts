import { goto } from "$app/navigation";
import { getTasks } from "$lib/api";
import type { PageLoad } from "./$types";


export const load = (async () => {
	const tasks = await getTasks();
	if (!tasks.status) {
		return await goto("/register");
	}

	return {
		tasks: tasks.data,
	}

}) satisfies PageLoad;
