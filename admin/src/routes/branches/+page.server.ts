import { getBranches } from "$lib"

export async function load() {
	const branches = await getBranches()

	console.log(branches)

	return {
		branches
	}
}
