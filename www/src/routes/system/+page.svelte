<script>
    import Header from '$lib/Header.svelte';
    import Registration from './Registration.svelte';
    import System from './System.svelte';
    import DataNotFound from './DataNotFound.svelte';

    async function getIdentity() {
        const res = await fetch("/api/identity");
        if (res.status == 200) {
            return res.json()
        } else {
            return undefined
        }
    }


</script>

{#await getIdentity()}
    <img src="assets/oval.svg" alt="">
{:then identity}
    <Header {identity} />
    {#if identity}
        <System {identity} />
    {:else}
        <Registration />
    {/if}
{:catch}
    <DataNotFound />
{/await}
