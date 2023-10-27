<script>
	import "../app.postcss";
	import { page } from "$app/stores";
	import {
		Sidebar,
		SidebarGroup,
		SidebarItem,
		SidebarWrapper,
	} from "flowbite-svelte";

	$: activeUrl = $page.url.pathname;
</script>

<div class="flex">
	<div class="h-screen">
	<Sidebar {activeUrl}>
		<SidebarWrapper>
			<SidebarGroup>
				<SidebarItem label="Dashboard" href="/" />
				<SidebarItem label="Companies" href="/companies" />
				<SidebarItem label="Branches" href="/branches" />
				<SidebarItem label="Networks" href="/networks" />
				<SidebarItem label="Users" href="/users" />
				<SidebarItem label="Devices" href="/devices" />
				<SidebarItem label="Tasks" href="/tasks" />
			</SidebarGroup>
		</SidebarWrapper>
	</Sidebar>

	</div>
	<div class="p-8 grow">
		<slot />
	</div>
</div>
