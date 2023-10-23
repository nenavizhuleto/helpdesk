import type { PageServerLoad } from './$types';
import mock from '$lib/mock';

export const load = (async () => {
	const identity = await mock.GetIdentity()
	const tasks = await mock.GetTasks()
	return {
		identity: identity,
		tasks: tasks,
	};
}) satisfies PageServerLoad;
