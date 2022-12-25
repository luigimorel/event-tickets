<div class="lg:w-3/12 w-full lg:pr-20">
	<div class=" flex justify-start items-center relative mb-6 lg:mb-12">
		<input
			class="leading-none text-left text-gray-600 text-base placeholder:font-bold placeholder:text-gray-500 px-4 py-5 w-full border rounded border-primary  outline-none"
			type="text"
			placeholder="Search for events"
		/>
		<svg
			class="absolute right-3 z-10 cursor-pointer"
			width={24}
			height={24}
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg">
			<path
				d="M10 17C13.866 17 17 13.866 17 10C17 6.13401 13.866 3 10 3C6.13401 3 3 6.13401 3 10C3 13.866 6.13401 17 10 17Z"
				stroke="#439A97"
				stroke-width="1.66667"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
			<path
				d="M21 21L15 15"
				stroke="#439A97"
				stroke-width="1.66667"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
		</svg>
	</div>

	<div class="lg:block hidden">
		<!-- Sort by date -->
		<div class=" mt-5">
			<h3 class="uppercase text-xl font-bold mb-2.5">Upcoming events</h3>

			<div class=" flex flex-col">
				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 cursor-pointer"
						type="radio"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Today
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 cursor-pointer"
						type="radio"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						This Week
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 cursor-pointer"
						type="radio"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Next 7 days
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 cursor-pointer"
						type="radio"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Next 30 days
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 cursor-pointer"
						type="radio"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						All Dates
					</label>
				</div>

				<div class="flex flex-row items-center">
					<p>Select custom dates</p>
				</div>
			</div>
		</div>

		<!-- Sort by event type -->

		<div class=" mt-5">
			<h3 class="uppercase text-xl font-bold mb-2.5">Type</h3>

			<div class=" flex flex-col">
				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Fundraiser
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Physical event
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Online event
					</label>
				</div>
			</div>
		</div>

		<!-- Sort by city -->

		<div class=" mt-5">
			<h3 class="uppercase text-xl font-bold mb-2.5">City</h3>

			<div class=" flex flex-col">
				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Kampala
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Lira
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Gulu
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Fort Portal
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Arua
					</label>
				</div>
			</div>
		</div>

		<!-- Sort by category -->

		<div class=" mt-5">
			<h3 class="uppercase text-xl font-bold mb-2.5">Category</h3>

			<div class=" flex flex-col">
				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Music
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Charity & Causes
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Health & Awareness
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Food & Drinks
					</label>
				</div>

				<div class="flex flex-row items-center mb-1">
					<input
						class=" checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer"
						type="checkbox"
						name="radio"
						id="radio"
					/>

					<label
						class="form-check-label mt-1 inline-block text-gray-800"
						for="radio">
						Community
					</label>
				</div>
			</div>
		</div>
	</div>
</div>
