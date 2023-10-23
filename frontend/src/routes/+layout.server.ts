import mock from "$lib/mock"


export async function load () {
	const identity = await mock.GetIdentity()

	return {
		identity: identity
	} 
} 
