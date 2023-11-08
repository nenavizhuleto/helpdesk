import { call } from "$lib"

export async function load() {
	const branches = await call("GET", "/branch")


	return {
		branches: branches.data
	}
}
