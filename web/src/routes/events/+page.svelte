<script lang="ts">
	import { onMount } from 'svelte';
	import Time from "svelte-time/src/Time.svelte";
	import Footer from '../../lib/components/Footer.svelte';
	import Header from '../../lib/components/Header.svelte';
	import Sidebar from "../../lib/components/Sidebar.svelte";
	import { apiData, apiLink, events } from '../../store/store';
	import '../../styles.css';

	onMount(async () => {
		fetch(`${apiLink}/events`)
			.then((response) => response.json())
			.then((data) => {
				apiData.set(data);
			})
			.catch((error) => {
				error.message;
				return [];
			});
	});

</script>

<Header />

<div class=" lg:pt-40 pt-32 mb-10 container mx-auto">
	<div class=" flex lg:flex-row flex-col px-4">

		<Sidebar/>

		<div class="lg:w-9/12  w-full">
			<h3 class=" text-2xl px-4 mb-8">{$events.length} results found</h3>
			<div
				class=" grid lg:grid-cols-4 md:grid-cols-3 px-4 md:gap-4 md:px-0  mb-10 md:mb-0 grid-cols-1">
				{#each $events as event}
					<div class="	 lg:mb-10 mb-10">
						<a href={`/events${event.slug}`}>
							<div class="bg-white shadow-md 	rounded md:max-w-sm w-full ">
								<img class="rounded-t-md w-full" src={`${event.coverImage}`} alt="event" />

								<div class="px-5 pb-5">
									<h3 class="text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight">
										{event.name}
									</h3>
									<p class=" text-gray-600 capitalize text-sm mb-3">{event.venue}</p>
									<div class=" flex flex-row justify-between">
										<p class=" text-gray-700 text-xs font-extrabold">
											<Time timestamp={event.date} format="ddd, D MMM YYYY" />
										</p>
										<p class=" text-gray-700 text-xs font-extrabold">
											<Time timestamp={event.eventTime} format="h:mm A" />
										</p>
									</div>
								</div>
							</div>
						</a>
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>

<Footer />
