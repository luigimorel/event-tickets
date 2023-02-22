import {
	c as create_ssr_component,
	d as add_attribute,
	a as subscribe,
	v as validate_component,
	f as each,
	e as escape
} from '../../chunks/index.js';
import { T as Time } from '../../chunks/Time.js';
import { e as events, F as Footer } from '../../chunks/store.js';
import { H as Header } from '../../chunks/styles.js';
const Button = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	let { link } = $$props;
	if ($$props.link === void 0 && $$bindings.link && link !== void 0)
		$$bindings.link(link);
	return `<div class="${'flex flex-row justify-center mb-10'}"><a${add_attribute(
		'href',
		`${link}`,
		0
	)}><button class="${'flex justify-center lg:mb-20 mb-4 w-52 mx-auto sm:mt-14 mt-4 bg-white border-primary border-2 text-primary hover:bg-primary hover:text-white hover:delay-200 hover:transition-all font-bold text-xl md:py-3 py-2.5 md:px-6 px-4 rounded text-center items-center'}"><span class="${'md:pt-2 sm:pt-0'}">See More</span></button></a></div>`;
});
const Hero = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	return `<section class="${'relative bg-[url(https://images.unsplash.com/photo-1519671482749-fd09be7ccebf?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80)] bg-cover bg-center bg-no-repeat'}"><div class="${'absolute inset-0 sm:bg-transparent sm:bg-gradient-to-r sm:from-white/95 sm:to-white/25'}"></div>

	<div class="${'relative mx-auto max-w-screen-xl px-4 py-32 sm:px-6 lg:flex lg:h-screen lg:items-center lg:px-8'}"><div class="${'text-center sm:text-left'}"><h1 class="${'text-3xl text-white md:text-gray-900 font-extrabold sm:text-5xl'}">Let us help you find an
				<strong class="${'block font-extrabold text-primary '}">Event Of Your Liking. </strong></h1>

			<p class="${'mt-4 md:text-gray-900 text-white max-w-lg font-bold sm:text-xl sm:leading-relaxed'}">Tune into online events, webinars, lessons and more.
				<br>
				From wherever you are.
			</p>

			<div class="${'mt-8 flex flex-wrap gap-4 text-center'}"><a href="${'/events'}" class="${'block w-full rounded bg-primary px-12 py-5 text-base font-medium text-white hover:font-bold shadow hover:bg-secondary focus:outline-none focus:ring active:bg-secondary sm:w-auto'}">Browse Events
				</a>

				<a href="${'/create-account'}" class="${'block w-full rounded bg-white px-12 py-5 border-primary border-2 text-base font-medium text-primary shadow hover:text-white hover:font-bold hover:bg-secondary focus:outline-none focus:ring active:bg-secondary sm:w-auto'}">Create Account
				</a></div></div></div></section>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	let $events, $$unsubscribe_events;
	$$unsubscribe_events = subscribe(events, (value) => ($events = value));
	$$unsubscribe_events();
	return `${validate_component(Header, 'Header').$$render($$result, {}, {}, {})}

${validate_component(Hero, 'Hero').$$render($$result, {}, {}, {})}

<div class="${'flex flex-col pt-32 max-w-screen-xl mx-auto'}"><h2 class="${'text-center mb-8 text-3xl font-bold uppercase'}">Experience online events
	</h2>
	<div class="${'grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0 mb-10 md:mb-0 grid-cols-1'}">

		${each($events, (event) => {
			return `<div class="${'mx-auto lg:mb-10 mb-10'}"><a${add_attribute(
				'href',
				`/${event.id / event.slug}`,
				0
			)}><div class="${'bg-white shadow-md rounded md:max-w-sm w-full '}"><img class="${'rounded-t-md w-full '}"${add_attribute(
				'src',
				`${event.coverImage}`,
				0
			)} alt="${'event'}">

						<div class="${'px-5 pb-5'}"><h3 class="${'text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight'}">${escape(
				event.name
			)}</h3>
							<p class="${'text-gray-600 text-sm mb-3 capitalize'}">${escape(event.venue)}</p>
							<div class="${'flex flex-row justify-between'}"><p class="${'text-gray-700 text-xs font-bold'}">${validate_component(
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
								<p class="${'text-gray-700 text-xs font-bold'}">${validate_component(
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
		})}</div>

	<div class="${'flex flex-row justify-center'}">${validate_component(
		Button,
		'Button'
	).$$render($$result, { link: '/events' }, {}, {})}</div></div>

<div class="${'flex flex-col max-w-screen-xl mx-auto'}"><h2 class="${'text-center mb-8 text-3xl font-bold uppercase'}">Experience fundraisings
	</h2>

	<div class="${'grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0 mb-10 md:mb-0 grid-cols-1'}">

		${each($events, (event) => {
			return `<div class="${'mx-auto lg:mb-10 mb-10'}"><a${add_attribute(
				'href',
				`/${event.id / event.slug}`,
				0
			)}><div class="${'bg-white shadow-md rounded md:max-w-sm w-full '}"><img class="${'rounded-t-md w-full '}"${add_attribute(
				'src',
				`${event.coverImage}`,
				0
			)} alt="${'event'}">

						<div class="${'px-5 pb-5'}"><h3 class="${'text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight'}">${escape(
				event.name
			)}</h3>
							<p class="${'text-gray-600 text-sm mb-3 capitalize'}">${escape(event.venue)}</p>
							<div class="${'flex flex-row justify-between'}"><p class="${'text-gray-700 text-xs font-bold'}">${validate_component(
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
								<p class="${'text-gray-700 text-xs font-bold'}">${validate_component(
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
		})}</div>

	<div class="${'flex flex-row justify-center'}">${validate_component(
		Button,
		'Button'
	).$$render($$result, { link: '/events' }, {}, {})}</div></div>

<div class="${'flex flex-col max-w-screen-xl mx-auto'}"><h2 class="${'text-center mb-8 text-3xl font-bold uppercase'}">Experience concerts
	</h2>
	<div class="${'grid lg:grid-cols-4 md:grid-cols-3 gap-8 px-4 md:px-0 mb-10 md:mb-0 grid-cols-1'}">

		${each($events, (event) => {
			return `<div class="${'mx-auto lg:mb-10 mb-10'}"><a${add_attribute(
				'href',
				`/${event.id / event.slug}`,
				0
			)}><div class="${'bg-white shadow-md rounded md:max-w-sm w-full '}"><img class="${'rounded-t-md w-full '}"${add_attribute(
				'src',
				`${event.coverImage}`,
				0
			)} alt="${'event'}">

						<div class="${'px-5 pb-5'}"><h3 class="${'text-gray-900 font-semibold text-lg capitalize mt-4 tracking-tight'}">${escape(
				event.name
			)}</h3>
							<p class="${'text-gray-600 text-sm mb-3 capitalize'}">${escape(event.venue)}</p>
							<div class="${'flex flex-row justify-between'}"><p class="${'text-gray-700 text-xs font-bold'}">${validate_component(
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
								<p class="${'text-gray-700 text-xs font-bold'}">${validate_component(
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
		})}</div>

	<div class="${'flex flex-row justify-center'}">${validate_component(
		Button,
		'Button'
	).$$render($$result, { link: '/events' }, {}, {})}</div></div>

${validate_component(Footer, 'Footer').$$render($$result, {}, {}, {})}`;
});
export { Page as default };
