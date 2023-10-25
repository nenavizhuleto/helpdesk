import type { PageServerLoad } from './$types';
import mock from '$lib/mock';
import { apiGET } from '$lib/api';
import { createUserTask, getUserTasks } from '$lib/api/tasks';
import type { Actions } from '@sveltejs/kit';
import { getIdentity } from '$lib/api/auth';

export const load = (async ({ parent }) => {
	//const identity = await mock.GetIdentity()
	//const tasks = await mock.GetTasks()
	const data = await parent()
	const identity = data.identity!;
	const tasks = await getUserTasks(identity.user.id)
	return {
		identity: identity,
		tasks: tasks,
	};
}) satisfies PageServerLoad;


export const actions: Actions = {
	task: async ({ request }) => {
		const data = await request.formData()

		const name = data.get("name")?.toString()!
		const subject = data.get("subject")?.toString()!
		console.log(name, subject)

		const identity = (await getIdentity())!

		const task = await createUserTask(identity.user.id, name, subject)
		console.log(task)
		return {
			success: true,
			task
		}
	}
}
