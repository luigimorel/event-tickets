<script lang="ts">
	function toggleMenu(flag: boolean) {
		let value = document.getElementById('menu');
		if (flag) {
			value?.classList.remove('hidden');
		} else {
			value?.classList.add('hidden');
		}
	}
</script>

<header>
	<nav class="bg-white px-2 sm:px-4 py-2.5 fixed w-full z-20 top-0 left-0 border-b border-gray-200">
		<div class="container mt-2 flex flex-wrap justify-between items-center mx-auto">
			<a href="/" class="flex items-center text-primary font-extrabold text-2xl"> Tickets </a>
			<div class="flex md:order-2 justify-center items-center">
				<a href="/create-account">
					<button
						type="button"
						class="text-white md:block hidden bg-primary hover:bg-secondary
					font-medium rounded-lg text-base px-5 py-4 text-center mr-5 md:mr-0">
						Create Account
					</button>
				</a>

				<button
					type="button"
					class="text-primary md:block hidden bg-white font-extrabold rounded-lg text-base px-5 py-4 text-center mr-3 md:mr-0">
					<a href="/login">Login</a>
				</button>

				<button
					on:click|preventDefault={() => toggleMenu(true)}
					data-collapse-toggle="navbar-sticky"
					type="button"
					class="inline-flex items-center p-2 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200"
					aria-controls="navbar-sticky"
					aria-expanded="false">
					<span class="sr-only">Open main menu</span>
					<svg
						class="w-6 h-6"
						aria-hidden="true"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
							clip-rule="evenodd" />
					</svg>
				</button>
			</div>
			<div
				class="hidden justify-between items-center w-full md:flex md:w-auto md:order-1"
				id="navbar-sticky">
				<ul
					class="flex flex-col p-4 mt-4 bg-gray-50 rounded-lg border border-gray-100 md:flex-row md:space-x-8
					md:mt- text-base md:font-medium md:border-0 md:bg-white">
					<li>
						<a
							href="/events"
							class="block py-2 pr-4 pl-3 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:p-0">
							Browse Events
						</a>
					</li>
				</ul>
			</div>
		</div>
	</nav>
</header>
