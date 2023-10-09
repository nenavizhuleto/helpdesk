<script>
    import Header from '$lib/Header.svelte';
    import Registration from '$lib/Registration.svelte';
    import System from '$lib/System.svelte';
    import { onMount } from 'svelte';

    let identity;
    let loggined = false;
    onMount(async () => {
        const res = await fetch("http://localhost:3000/api/identity");
        identity = await res.json();
        if (identity?.status == 503) {
            loggined = false
        } else {
            loggined = true
        }
        console.log(identity)
    })

</script>

<Header {identity} />
{#if loggined}
<System {identity} />
{:else}
<Registration {identity} />
{/if}
