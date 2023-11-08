// place files you want to import through the `$lib` alias in this folder.

export const url: string = "http://127.0.0.1:3000/api/admin";

export interface Response {
	status: boolean,
	data: any,
	error: any,
}

export async function call(method: "GET" | "POST" | "DELETE", path: string, body?: any): Promise<Response> {
	const res = await fetch(url + path, {
		method: method,
		headers: {
			"token": "password",
		},
		body: JSON.stringify(body)
	})

	return await res.json() as Response
}

export async function getCompanies() {
	return await call("GET", "/company")
}

export async function getCompany(id: string) {
	return await call("GET", "/company/" + id)
}

export async function newCompany(itn: string, name: string) {
	return await call("POST", "/company", { itn, name })
}


export async function getBranches() {
	return await call("GET", "/branch")
}

export async function getNetworks() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/network")
	return res.json()
}

export async function getUsers() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/users")
	return res.json()
}

export async function getDevices() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/devices")
	return res.json()
}

export async function getTasks() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/tasks")
	return res.json()
}
