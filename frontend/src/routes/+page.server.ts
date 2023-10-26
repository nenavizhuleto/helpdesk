import { redirect } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load = (async ({ parent }) => {
	const data = await parent()
	console.log(data)

	if (data.error) {
		throw redirect(300, "/register")
	}

	return {
		identity: data.identity
	}
}) satisfies PageServerLoad
