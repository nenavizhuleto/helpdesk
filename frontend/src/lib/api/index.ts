
const baseURL = "http://127.0.0.1:3000/api/v3"

async function apiCall(method: "GET" | "POST" | "PUT" | "DELETE", url: string, body: any): Promise<any> {
	const res = await fetch(baseURL + url, {
		method: method,
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(body)
	})

	const data = await res.json()

	return data
}

export async function apiGET(url: string, body?: any): Promise<any> {
	return await apiCall("GET", url, body)
}

export async function apiPOST(url: string, body?: any): Promise<any> {
	return await apiCall("POST", url, body)
}

export async function apiPUT(url: string, body?: any): Promise<any> {
	return await apiCall("PUT", url, body)
}

export async function apiDELETE(url: string, body?: any): Promise<any> {
	return await apiCall("PUT", url, body)
}
