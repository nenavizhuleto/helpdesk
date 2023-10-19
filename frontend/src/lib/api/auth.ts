import { apiGET, apiPOST } from ".";
import type { Identity } from "./types";

export async function getIdentity(): Promise<Identity | null> {
	const identity = await apiGET("/identity") as Identity;
	return identity;
}

export async function register(name: string, phone: string): Promise<any> {
	const res = await apiPOST("/register", { name, phone });
	return res;
}
