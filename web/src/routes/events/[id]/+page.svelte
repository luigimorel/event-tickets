<script lang="ts">
	import { page } from '$app/stores';

	import { onMount } from 'svelte';
	import Footer from '../../../lib/components/Footer.svelte';
	import Header from '../../../lib/components/Header.svelte';
	import { apiData, apiLink, singleEvent } from '../../../store/store';
	import '../../../styles.css';

	let id = $page.params?.id;

	onMount(async () => {
		fetch(`${apiLink}/events/${id}`)
			.then((response) => response.json())
			.then((data) => {
				apiData.set(data);
			})

			.catch((error) => {
				console.log(error);
				return [];
			});
	});
</script>

<Header />

<div class=" pt-40">
	<p>{console.log(singleEvent)}</p>
	{#each singleEvent as singoEvent}
		<p>{singoEvent.name}</p>
	{/each}
</div>
<Footer />
