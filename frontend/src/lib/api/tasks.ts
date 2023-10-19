import { apiGET } from ".";
import type { Task } from "./types";


export async function getUserTasks(user_id: string): Promise<Task[]> {
	const tasks = await apiGET(`/users/${user_id}/tasks`) as Task[];
	return tasks;
}
