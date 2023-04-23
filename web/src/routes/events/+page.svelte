<script lang="ts">
	import { onMount } from 'svelte';
	import Time from 'svelte-time/src/Time.svelte';
	import Footer from '../../lib/components/Footer.svelte';
	import Header from '../../lib/components/Header.svelte';
	import Sidebar from '../../lib/components/Sidebar.svelte';
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
		<Sidebar />

		<div class="lg:w-9/12  w-full">
			<h3 class=" text-2xl px-4 mb-8">{$events.length} results found</h3>
			<div
				class=" grid lg:grid-cols-4 md:grid-cols-3 px-4 md:gap-4 md:px-0  mb-10 md:mb-0 grid-cols-1">
				{#each $events as event}
					<div class="	 lg:mb-10 mb-10">
						<a href={`/events${event.slug}/${event.id}`}>
							<div class="bg-white shadow-md 	rounded md:max-w-sm w-full ">
								<img
									class="rounded-t-md w-full"
									src={`${event.coverImage}`}
									alt="event" />

								<div class="px-5 pb-5">
									<h3
										class="text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight">
										{event.name}
									</h3>
									<p class=" text-gray-600 capitalize text-sm mb-3">
										{event.venue}
									</p>
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
			<div class="flex justify-center my-8-center">
				<div class="flex items-center justify-center mr-8 gap-2">
					<svg
						class="w-5 h-5"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z"
							clip-rule="evenodd" /></svg>
					Prev
				</div>
				<div class="flex items-center gap-2">
					Next
					<svg
						class="w-5 h-5"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z"
							clip-rule="evenodd" /></svg>
				</div>
			</div>
		</div>
	</div>
</div>

<Footer />
