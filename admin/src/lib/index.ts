// place files you want to import through the `$lib` alias in this folder.
export async function getCompanies() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/company")
	return res.json()
}

export async function getBranches() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/branch")
	return res.json()
}

export async function getNetworks() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/network")
	return res.json()
}

export async function getUsers() {
	const res = await fetch("http://127.0.0.1:3000/api/v3/users")
	return res.json()
}
