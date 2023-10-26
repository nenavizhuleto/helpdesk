import * as api from "$lib/api"
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
