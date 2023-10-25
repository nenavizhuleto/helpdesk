import { getIdentity } from "$lib/api/auth"
import mock from "$lib/mock"


export async function load () {
	//const identity = await mock.GetIdentity()
	const identity = getIdentity()

	return {
		identity: identity
	} 
} 
