import { getTasks } from "$lib"

export async function load() {
	const tasks = await getTasks()

	console.log(tasks)

	return {
		tasks
	}
}
