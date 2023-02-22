import {
	c as create_ssr_component,
	d as add_attribute,
	a as subscribe,
	v as validate_component,
	e as escape,
	f as each
} from '../../../chunks/index.js';
import { T as Time } from '../../../chunks/Time.js';
import { e as events, F as Footer } from '../../../chunks/store.js';
import { H as Header } from '../../../chunks/styles.js';
const Sidebar = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	return `<div class="${'lg:w-3/12 w-full lg:pr-20'}"><div class="${'flex justify-start items-center relative mb-6 lg:mb-12'}"><input class="${'leading-none text-left text-gray-600 text-base placeholder:font-bold placeholder:text-gray-500 px-4 py-5 w-full border rounded border-primary outline-none'}" type="${'text'}" placeholder="${'Search for events'}">
		<svg class="${'absolute right-3 z-10 cursor-pointer'}"${add_attribute(
		'width',
		24,
		0
	)}${add_attribute(
		'height',
		24,
		0
	)} viewBox="${'0 0 24 24'}" fill="${'none'}" xmlns="${'http://www.w3.org/2000/svg'}"><path d="${'M10 17C13.866 17 17 13.866 17 10C17 6.13401 13.866 3 10 3C6.13401 3 3 6.13401 3 10C3 13.866 6.13401 17 10 17Z'}" stroke="${'#439A97'}" stroke-width="${'1.66667'}" stroke-linecap="${'round'}" stroke-linejoin="${'round'}"></path><path d="${'M21 21L15 15'}" stroke="${'#439A97'}" stroke-width="${'1.66667'}" stroke-linecap="${'round'}" stroke-linejoin="${'round'}"></path></svg></div>

	<div class="${'lg:block hidden'}">
		<div class="${'mt-5'}"><h3 class="${'uppercase text-xl font-bold mb-2.5'}">Upcoming events</h3>

			<div class="${'flex flex-col'}"><div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 cursor-pointer'}" type="${'radio'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Today
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 cursor-pointer'}" type="${'radio'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">This Week
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 cursor-pointer'}" type="${'radio'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Next 7 days
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 cursor-pointer'}" type="${'radio'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Next 30 days
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 cursor-pointer'}" type="${'radio'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">All Dates
					</label></div>

				<div class="${'flex flex-row items-center'}"><p>Select custom dates</p></div></div></div>

		

		<div class="${'mt-5'}"><h3 class="${'uppercase text-xl font-bold mb-2.5'}">Type</h3>

			<div class="${'flex flex-col'}"><div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Fundraiser
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Physical event
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Online event
					</label></div></div></div>

		

		<div class="${'mt-5'}"><h3 class="${'uppercase text-xl font-bold mb-2.5'}">City</h3>

			<div class="${'flex flex-col'}"><div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Kampala
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Lira
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Gulu
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Fort Portal
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Arua
					</label></div></div></div>

		

		<div class="${'mt-5'}"><h3 class="${'uppercase text-xl font-bold mb-2.5'}">Category</h3>

			<div class="${'flex flex-col'}"><div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Music
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Charity &amp; Causes
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Health &amp; Awareness
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Food &amp; Drinks
					</label></div>

				<div class="${'flex flex-row items-center mb-1'}"><input class="${'checked:bg-primary checked:border-primary mr-2 rounded-sm cursor-pointer'}" type="${'checkbox'}" name="${'radio'}" id="${'radio'}">

					<label class="${'form-check-label mt-1 inline-block text-gray-800'}" for="${'radio'}">Community
					</label></div></div></div></div></div>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	let $events, $$unsubscribe_events;
	$$unsubscribe_events = subscribe(events, (value) => ($events = value));
	$$unsubscribe_events();
	return `${validate_component(Header, 'Header').$$render($$result, {}, {}, {})}

<div class="${'lg:pt-40 pt-32 mb-10 container mx-auto'}"><div class="${'flex lg:flex-row flex-col px-4'}">${validate_component(
		Sidebar,
		'Sidebar'
	).$$render($$result, {}, {}, {})}

		<div class="${'lg:w-9/12 w-full'}"><h3 class="${'text-2xl px-4 mb-8'}">${escape(
		$events.length
	)} results found</h3>
			<div class="${'grid lg:grid-cols-4 md:grid-cols-3 px-4 md:gap-4 md:px-0 mb-10 md:mb-0 grid-cols-1'}">${each(
		$events,
		(event) => {
			return `<div class="${'lg:mb-10 mb-10'}"><a${add_attribute(
				'href',
				`/events${event.slug}/${event.id}`,
				0
			)}><div class="${'bg-white shadow-md rounded md:max-w-sm w-full '}"><img class="${'rounded-t-md w-full'}"${add_attribute(
				'src',
				`${event.coverImage}`,
				0
			)} alt="${'event'}">

								<div class="${'px-5 pb-5'}"><h3 class="${'text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight'}">${escape(
				event.name
			)}</h3>
									<p class="${'text-gray-600 capitalize text-sm mb-3'}">${escape(event.venue)}</p>
									<div class="${'flex flex-row justify-between'}"><p class="${'text-gray-700 text-xs font-extrabold'}">${validate_component(
				Time,
				'Time'
			).$$render(
				$$result,
				{
					timestamp: event.date,
					format: 'ddd, D MMM YYYY'
				},
				{},
				{}
			)}</p>
										<p class="${'text-gray-700 text-xs font-extrabold'}">${validate_component(
				Time,
				'Time'
			).$$render(
				$$result,
				{
					timestamp: event.eventTime,
					format: 'h:mm A'
				},
				{},
				{}
			)}</p>
									</div></div>
							</div></a>
					</div>`;
		}
	)}</div>
			<div class="${'flex justify-center my-8-center'}"><div class="${'flex items-center justify-center mr-8 gap-2'}"><svg class="${'w-5 h-5'}" fill="${'currentColor'}" viewBox="${'0 0 20 20'}" xmlns="${'http://www.w3.org/2000/svg'}"><path fill-rule="${'evenodd'}" d="${'M7.707 14.707a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l2.293 2.293a1 1 0 010 1.414z'}" clip-rule="${'evenodd'}"></path></svg>
					Prev
				  </div>
				  <div class="${'flex items-center gap-2'}">Next
					<svg class="${'w-5 h-5'}" fill="${'currentColor'}" viewBox="${'0 0 20 20'}" xmlns="${'http://www.w3.org/2000/svg'}"><path fill-rule="${'evenodd'}" d="${'M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z'}" clip-rule="${'evenodd'}"></path></svg></div></div></div></div></div>

${validate_component(Footer, 'Footer').$$render($$result, {}, {}, {})}`;
});
export { Page as default };
