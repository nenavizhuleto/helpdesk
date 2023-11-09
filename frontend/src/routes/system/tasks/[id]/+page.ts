import { getTask, getTaskComments } from "$lib/api";
import { error } from "@sveltejs/kit";
import type { PageLoad } from "./$types";

export const load = (async ({ params }) => {
	const id = params.id;
	const task = await getTask(id)
	if (!task.status) {
		throw error(404);
	}

	const comments = await getTaskComments(id)
	if (!comments.status) {
		throw error(404)
	}

	console.log(comments)

	return {
		task: task.data,
		comments: comments.data
	}

}) satisfies PageLoad;
