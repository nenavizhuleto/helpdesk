import { getCompanies } from "$lib"

export async function load({}) {
	const companies = await getCompanies()

	console.log(companies)

	return {
		companies: companies
	}
}
