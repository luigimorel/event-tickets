<script lang="ts">
	export let link: string;
</script>

<div class="flex flex-row justify-center mb-10">
	<a href={`${link}`}>
		<button
			class=" flex justify-center lg:mb-20 mb-4 w-52 mx-auto sm:mt-14 mt-4 bg-white border-primary border-2 text-primary hover:bg-primary
			 hover:text-white hover:delay-200 hover:transition-all font-bold text-xl md:py-3 py-2.5 md:px-6 px-4 rounded text-center items-center">
			<span class="md:pt-2 sm:pt-0">See More</span>
		</button>
	</a>
</div>
