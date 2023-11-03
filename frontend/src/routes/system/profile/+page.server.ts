import { redirect } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"
import * as api from '$lib/api'

export const load = (async ({ parent }) => {
	//const data = await parent()
	//console.log(data)

	//if (!data.identity) {
	//	throw redirect(300, "/register")
	//}
	//const [tasks, error] = await api.getUserTasks(data.identity.id)
	//if (error) {
	//	return {
	//		error
	//	}
	//}
	//return {
	//	identity: data.identity!,
	//	tasks: tasks!,
	//};
}) satisfies PageServerLoad
