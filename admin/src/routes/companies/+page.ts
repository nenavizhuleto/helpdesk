import { call } from "$lib"
import { error } from "@sveltejs/kit"

export async function load({ }) {
	const companies = await call("GET", "/company")

	return {
		companies: companies.data
	}
}
