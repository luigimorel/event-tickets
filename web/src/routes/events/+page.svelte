<script lang="ts">
	import Header from '../../lib/components/Header.svelte';

	import { onMount } from 'svelte';
	import Footer from '../../lib/components/Footer.svelte';
	import { apiData, apiLink, events } from '../../store/store';
	import '../../styles.css';

	onMount(async () => {
		fetch(`${apiLink}/events`)
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

<div class=" pt-40 max-w-full mx-auto">
	<div class=" flex flex-row md:px-14 px-4">
		<div class=" w-3/12 pr-20">
			<p>
				Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto animi dignissimos culpa
				sunt consequatur minima aliquam at eligendi, voluptatum maiores natus, distinctio et, illo
				ut porro. Saepe sint exercitationem vero?
			</p>
		</div>
		<div class="w-9/12">
			<h3 class=" text-2xl mb-8">{$events.length} results found</h3>

			<div
				class=" grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0  mb-10 md:mb-0 grid-cols-1"
			>
				{#each $events as event}
					<div class=" mx-auto lg:mb-10 mb-10">
						<a href={`/events/${event.id}`}>
							<div class="bg-white shadow-md rounded md:max-w-sm w-full ">
								<img class="rounded-t-md w-full " src={`${event.coverImage}`} alt="event" />

								<div class="px-5 pb-5">
									<h3 class="text-gray-900 font-semibold text-lg mt-4 tracking-tight">
										{event.name}
									</h3>
									<p class=" text-gray-600 text-sm mb-3">{event.venue}</p>
									<div class=" flex flex-row justify-between">
										<p class=" text-gray-700 text-xs font-bold">{event.date}</p>
										<p class=" text-gray-700 text-xs font-bold">{event.eventTime}</p>
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
