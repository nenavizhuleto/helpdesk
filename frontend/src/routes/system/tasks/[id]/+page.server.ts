import * as api from "$lib/api"
import type { Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";



export const load = (async ({ params }) => {
	const [task, error] = await api.getTaskById(params.id)
	if (error) {
		return {
			error
		}
	}
	return {
		task: task!,
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	comment: async ({ request }) => {
		const data = await request.formData()
		const task_id = data.get("task_id")!
		const message = data.get("message")!

		const direction = "from"

		const [comment, error] = await api.commentTask(task_id.toString(), message, direction)
		if (error) {
			return {
				error
			}
		}

		return {
			comment
		}


	}
}
