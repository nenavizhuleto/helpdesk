import { getUsers } from "$lib"

export async function load() {
	const users = await getUsers()

	console.log(users)

	return {
		users
	}
}
