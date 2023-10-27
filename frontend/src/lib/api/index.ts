import type {APIError, Identity, User, Task} from "./types";
const baseURL = "http://172.16.222.31:3000/api/v3";

type Response<T> = [T | undefined, APIError | undefined];

async function apiCall(
	method: "GET" | "POST" | "PUT" | "DELETE",
	url: string,
	body: any
): Promise<Response<any>> {
	const res = await fetch(baseURL + url, {
		method: method,
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(body),
	});

	const data = await res.json();

	if (res.status == 500) {
		return [undefined, data as APIError];
	}

	return [data, undefined];
}

export async function apiGET(url: string, body?: any): Promise<Response<any>> {
	return await apiCall("GET", url, body);
}

export async function apiPOST(url: string, body?: any): Promise<Response<any>> {
	return await apiCall("POST", url, body);
}

export async function apiPUT(url: string, body?: any): Promise<Response<any>> {
	return await apiCall("PUT", url, body);
}

export async function apiDELETE(
	url: string,
	body?: any
): Promise<Response<any>> {
	return await apiCall("PUT", url, body);
}

export async function getIdentity(): Promise<Response<Identity>> {
	const [identity, err] = await apiGET("/identity");
	return [identity as Identity, err];
}

export async function RegisterUser(
	name: string,
	phone: string
): Promise<Response<User>> {
	const [user, err] = await apiPOST("/register", {name, phone});
	return [user as User, err];
}

export async function getUserTasks(user_id: string): Promise<Response<Task[]>> {
	const [tasks, error] = await apiGET(`/users/${user_id}/tasks`);
	return [tasks as Task[], error];
}

export async function getTaskById(task_id: string): Promise<Response<Task>> {
	const [task, error] = await apiGET(`/tasks/${task_id}`);
	return [task as Task, error];
}

export async function createUserTask(
	user_id: string,
	name: string,
	subject: string
): Promise<Response<Task>> {
	const [task, error] = await apiPOST(`/users/${user_id}/tasks`, {
		name,
		subject,
	});
	return [task as Task, error];
}
