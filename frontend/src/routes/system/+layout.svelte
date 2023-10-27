<script lang="ts">
	// --- Utils ---
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";
	import { writable } from "svelte/store";
	import { setContext } from "svelte";

	// --- Types ---
	import type { PageData } from "./$types";
	import type { Identity } from "$lib/api/types";

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
	const identity = writable<Identity>();
	// If couldn't identify user goto registration
	// Else move identity to storage for using in other pages
	$: () => {
		if (!data.identity) {
			goto("/register");
		} else {
			identity.set(data.identity);
			setContext("identity", identity);
		}
	};

	// highlight corresponding SideBarItem in relation of current url
	$: activeUrl = $page.url.pathname;
</script>

<div class="flex h-full">
	<div class="bg-gray-50">
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
