<script lang="ts">
	export let label: string;
	export let type: string;
	export let name: string;
	export let id: string;
	export let placeholder: string;
	export let pattern: string | undefined = undefined;
</script>

<div class="flex flex-col mb-5">
	<label class="mb-2 font-bold lg:text-gray-800 text-white" for={id}>
		{label}
	</label>

	<input
		{type}
		{name}
		{id}
		{pattern}
		{placeholder}
		class="block w-full py-3 px-2 text-base rounded bg-white text-gray-800 border border-gray-800" />
</div>
