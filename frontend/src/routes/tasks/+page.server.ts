import { getIdentity } from '$lib/api/auth';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { getUserTasks } from '$lib/api/tasks';

export const load = (async () => {
	const identity = await getIdentity();
	if (!identity) {
		throw redirect(300, "/register");
	}
	const tasks = await getUserTasks(identity.user.id);
	return {
		identity: identity,
		tasks: tasks,
	};
}) satisfies PageServerLoad;
