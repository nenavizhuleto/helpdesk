import { redirect } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"

export const load = (async ({ parent }) => {
	throw redirect(302, "/system/profile")
}) satisfies PageServerLoad
