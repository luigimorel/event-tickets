<script>
	import Card from '../lib/components/Card.svelte';
	import Footer from "../lib/components/Footer.svelte";
	import Header from '../lib/components/Header.svelte';
	import Hero from '../lib/components/Hero.svelte';
	import '../styles.css';
</script>

<Header />
<Hero />
<div class=" flex flex-col pt-32 max-w-screen-xl mx-auto">
	<h2 class=" text-center mb-8 text-3xl font-bold uppercase">Experience online events</h2>
	<div class=" grid lg:grid-cols-4 md:grid-cols-3 space-x-7 grid-cols-1">
		<Card />
		<Card />
		<Card />
		<Card />
	</div>

	<button
		class=" lg:mb-20 mb-4 mx-auto bg-white border-primary text-primary border font-bold text-xl py-3 px-6 rounded text-center items-center">
		See More
	</button>
</div>

<div class=" flex flex-col pt-4 max-w-screen-xl mx-auto">
	<h2 class=" text-center mb-8 text-3xl font-bold uppercase">Experience online events</h2>
	<div class=" grid lg:grid-cols-4 md:grid-cols-3 space-x-7 grid-cols-1">
		<Card />
		<Card />
		<Card />
		<Card />
	</div>
	<button
		class=" lg:mb-20 mb-4 mx-auto bg-white border-primary text-primary border font-bold text-xl py-3 px-6 rounded text-center items-center">
		See More
	</button>

</div>


<div class=" flex flex-col pt-4 max-w-screen-xl mx-auto">
	<h2 class=" text-center mb-8 text-3xl font-bold uppercase">WORTHWHILE FUNDRAISERS</h2>
	<div class=" grid lg:grid-cols-4 md:grid-cols-3 space-x-7 grid-cols-1">
		<Card />
		<Card />
		<Card />
		<Card />
	</div>
	<button
		class=" lg:mb-20 mb-4 mx-auto bg-white border-primary text-primary border font-bold text-xl py-3 px-6 rounded text-center items-center">
		See More
	</button>

</div>

<Footer/>
