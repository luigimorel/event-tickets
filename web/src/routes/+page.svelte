<script>
	// @ts-nocheck

	import Time from 'svelte-time';
	import Button from '../lib/components/Button.svelte';
	import Footer from '../lib/components/Footer.svelte';
	import Header from '../lib/components/Header.svelte';
	import Hero from '../lib/components/Hero.svelte';

	import '../styles.css';

	import { onMount } from 'svelte';
	import { apiData, apiLink, events } from '../store/store';

	onMount(async () => {
		fetch(`${apiLink}/homepage_events`, {
			method: 'GET'
		})
			.then((response) => response.json())
			.then((data) => {
				apiData.set(data);
			})
			.catch((error) => {
				return [];
			});
	});
</script>

<Header />

<Hero />

<div class=" flex flex-col pt-32 max-w-screen-xl mx-auto">
	<h2 class=" text-center mb-8 text-3xl font-bold uppercase">Experience online events</h2>
	<div class=" grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0  mb-10 md:mb-0 grid-cols-1">

		<!-- { $events.length === 0 ?
			<p>Failed to fetch events </p> : ""
		} -->

		{#each $events as event (event.id)}
			<div class=" mx-auto lg:mb-10 mb-10">
				<a href={`/${event.id / event.slug}`}>
					<div class="bg-white shadow-md rounded md:max-w-sm w-full ">
						<img class="	rounded-t-md w-full " src={`${event.coverImage}`} alt="event" />

						<div class="px-5 pb-5">
							<h3 class="text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight">
								{event.name}
							</h3>
							<p class=" text-gray-600 text-sm mb-3 capitalize">{event.venue}</p>
							<div class=" flex flex-row justify-between">
								<p class=" text-gray-700 text-xs font-bold">
									<Time timestamp={event.date} format="ddd, D MMM YYYY" />
								</p>
								<p class=" text-gray-700 text-xs font-bold">
									<Time timestamp={event.eventTime} format="h:mm A" />
								</p>
							</div>
						</div>
					</div>
				</a>
			</div>
		{/each}
	</div>

	<div class="flex flex-row justify-center">
		<Button link="/events" />
	</div>
</div>

<div class=" flex flex-col		 max-w-screen-xl mx-auto">
	<h2 class=" text-center mb-8 text-3xl font-bold uppercase">Experience online events</h2>
	<div class=" grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0  mb-10 md:mb-0 grid-cols-1">

		<!-- { $events.length === 0 ?
			<p>Failed to fetch events </p> : ""
		} -->

		{#each $events as event (event.id)}
			<div class=" mx-auto lg:mb-10 mb-10">
				<a href={`/${event.id / event.slug}`}>
					<div class="bg-white shadow-md rounded md:max-w-sm w-full ">
						<img class="	rounded-t-md w-full " src={`${event.coverImage}`} alt="event" />

						<div class="px-5 pb-5">
							<h3 class="text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight">
								{event.name}
							</h3>
							<p class=" text-gray-600 text-sm mb-3 capitalize">{event.venue}</p>
							<div class=" flex flex-row justify-between">
								<p class=" text-gray-700 text-xs font-bold">
									<Time timestamp={event.date} format="ddd, D MMM YYYY" />
								</p>
								<p class=" text-gray-700 text-xs font-bold">
									<Time timestamp={event.eventTime} format="h:mm A" />
								</p>
							</div>
						</div>
					</div>
				</a>
			</div>
		{/each}
	</div>

	<div class="flex flex-row justify-center">
		<Button link="/events" />
	</div>
</div>

<div class=" flex flex-col max-w-screen-xl mx-auto">
	<h2 class=" text-center mb-8 text-3xl font-bold uppercase">Experience online events</h2>
	<div class=" grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0  mb-10 md:mb-0 grid-cols-1">

		<!-- { $events.length === 0 ?
			<p>Failed to fetch events </p> : ""
		} -->

		{#each $events as event (event.id)}
			<div class=" mx-auto lg:mb-10 mb-10">
				<a href={`/${event.id / event.slug}`}>
					<div class="bg-white shadow-md rounded md:max-w-sm w-full ">
						<img class="	rounded-t-md w-full " src={`${event.coverImage}`} alt="event" />

						<div class="px-5 pb-5">
							<h3 class="text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight">
								{event.name}
							</h3>
							<p class=" text-gray-600 text-sm mb-3 capitalize">{event.venue}</p>
							<div class=" flex flex-row justify-between">
								<p class=" text-gray-700 text-xs font-bold">
									<Time timestamp={event.date} format="ddd, D MMM YYYY" />
								</p>
								<p class=" text-gray-700 text-xs font-bold">
									<Time timestamp={event.eventTime} format="h:mm A" />
								</p>
							</div>
						</div>
					</div>
				</a>
			</div>
		{/each}
	</div>

	<div class="flex flex-row justify-center">
		<Button link="/events" />
	</div>
</div>


<Footer />
