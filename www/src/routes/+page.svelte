<script>
    import Header from '$lib/Header.svelte';
    import Registration from '$lib/pages/Registration.svelte';
    import System from '$lib/pages/System.svelte';

    async function getIdentity() {
        const res = await fetch("http://localhost:3000/api/identity");
        if (res.status == 200) {
            return res.json()
        } else {
            return undefined
        }
    }

</script>

{#await getIdentity()}
    <h1>loading...</h1>
{:then identity}
    <Header {identity} />
    {#if identity}
        <System {identity} />
    {:else}
        <Registration />
    {/if}
{/await}
