import type { PageServerLoad } from './$types';
import type { Actions } from '@sveltejs/kit';
import * as api from '$lib/api'

export const load = (async ({ parent }) => {
	const data = await parent()
	const identity = data.identity!;
	const [tasks, error] = await api.getUserTasks(identity.user.id)
	if (error) {
		return {
			error
		}
	}
	return {
		identity: identity,
		tasks: tasks!,
	};
}) satisfies PageServerLoad;


export const actions: Actions = {
	task: async ({ request }) => {
		const data = await request.formData()

		const name = data.get("name")?.toString()!
		const subject = data.get("subject")?.toString()!
		console.log(name, subject)

		let [identity, i_error] = await api.getIdentity()
		if (i_error) {
			return {
				error: i_error
			}
		}

		let [task, error] = await api.createUserTask(identity!.user.id, name, subject)
		if (error) {
			return {
				error
			}
		}
		console.log(task)
		return {
			success: true,
			task: task!
		}
	}
}
