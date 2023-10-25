import { apiGET, apiPOST } from ".";
import type { Task } from "./types";


export async function getUserTasks(user_id: string): Promise<Task[]> {
	const tasks = await apiGET(`/users/${user_id}/tasks`) as Task[];
	return tasks;
}

export async function createUserTask(user_id: string, name: string, subject: string): Promise<Task> {
	const task = await apiPOST(`/users/${user_id}/tasks`, { name, subject }) as Task;
	return task;
}
