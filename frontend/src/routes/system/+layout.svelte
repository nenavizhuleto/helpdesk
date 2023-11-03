<script lang="ts">
	// --- Utils ---
	import { page } from "$app/stores";
	import type { Writable } from "svelte/store";
	import { getContext } from "svelte";

	// --- Types ---
	import type { PageData } from "./$types";

	// --- Components ---
	import {
		Sidebar,
		SidebarWrapper,
		SidebarGroup,
		SidebarBrand,
		SidebarItem,
	} from "flowbite-svelte";

	// -- Icons ---
	import {
		ClipboardListOutline,
		UserCircleOutline,
	} from "flowbite-svelte-icons";

	export let data: PageData;
	import type { User } from "$lib/api/types";
	import { beforeNavigate, goto } from "$app/navigation";
	import { getIdentity } from "$lib/api";
	import { setContext } from "svelte";
	import { userStore } from "$lib";

	beforeNavigate(async (nav) => {
		let [identity, error] = await getIdentity();
		if (identity) {
			userStore.set(identity);
			setContext("user", identity)
		}
		if (error) {
			nav.cancel();
			goto("/register");
		}
	});
	let user = getContext<Writable<User>>("user");
	$: console.log(user);
	// highlight corresponding SideBarItem in relation of current url
	$: activeUrl = $page.url.pathname;
</script>

<div class="flex h-full">
	<div class="bg-gray-50 border-r border-gray-300 px-3 pt-3">
		<Sidebar asideClass="" {activeUrl}>
			<SidebarWrapper>
				<SidebarBrand
					site={{ href: "/", name: "", img: "/Logotype.svg" }}
				/>
				<SidebarGroup>
					<SidebarItem href="/system/profile" label="Профиль">
						<svelte:fragment slot="icon">
							<UserCircleOutline />
						</svelte:fragment>
					</SidebarItem>
					<SidebarItem href="/system/tasks" label="Обращения">
						<svelte:fragment slot="icon">
							<ClipboardListOutline />
						</svelte:fragment>
					</SidebarItem>
				</SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
	</div>
	<slot />
</div>
