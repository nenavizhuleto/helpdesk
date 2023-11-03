import type { Token, Profile, Task, Comment } from "./types";
const url = "http://172.16.223.26:3000/api";

interface Error {
	type: string,
	body: any,
}

interface Response<T> {
	status: boolean,
	data: T,
	error: Error | undefined,
}

async function call<T>(
	method: "GET" | "POST" | "PUT" | "DELETE",
	path: string,
	body: any
): Promise<Response<T>> {
	const response = await fetch(url + path, {
		method: method,
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(body),
	});

	const data = await response.json();

	return data;
}

export async function GET<T>(path: string): Promise<Response<T>> {
	return await call<T>("GET", path, undefined);
}
export async function POST<T>(path: string, body: any): Promise<Response<T>> {
	return await call<T>("POST", path, body);
}

export async function getToken(): Promise<Response<Token>> {
	return await GET<Token>("/auth/token");
}

export async function register(name: string, phone: string): Promise<Response<Token>> {
	return await POST<Token>("/auth/register", { name, phone });
}


export async function getProfile(): Promise<Response<Profile>> {
	return await GET<Profile>("/helpdesk/profile");
}

type TaskFilterOptions = "branch" | "company";

export async function getTasks(filter?: TaskFilterOptions): Promise<Response<Task[]>> {
	const url = filter ? "/helpdesk/tasks?filter=" + filter : "/helpdesk/tasks";
	return await GET<Task[]>(url);
}

export async function getTask(id: string): Promise<Response<Task>> {
	return await GET<Task>("/helpdesk/tasks/" + id);
}

export async function newTask(name: string, subject: string): Promise<Response<string>> {
	return await POST<string>("/helpdesk/tasks", { name, subject });
}

export async function getTaskComments(id: string): Promise<Response<Comment[]>> {
	return await GET<Comment[]>("/helpdesk/tasks/" + id + "/comments");
}

export async function newTaskComment(id: string, content: string): Promise<Response<string>> {
	return await POST<string>("/helpdesk/tasks/" + id + "/comments", { content });
}
