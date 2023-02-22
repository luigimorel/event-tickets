import {
	c as create_ssr_component,
	a as subscribe,
	v as validate_component,
	e as escape,
	f as each
} from '../../../../chunks/index.js';
import { p as page } from '../../../../chunks/stores.js';
import { s as singleEvent, F as Footer } from '../../../../chunks/store.js';
import { H as Header } from '../../../../chunks/styles.js';
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	let $page, $$unsubscribe_page;
	$$unsubscribe_page = subscribe(page, (value) => ($page = value));
	$page.params?.id;
	$$unsubscribe_page();
	return `${validate_component(Header, 'Header').$$render($$result, {}, {}, {})}

<div class="${'pt-40'}"><p>${escape(console.log(singleEvent))}</p>
	${each(singleEvent, (singoEvent) => {
		return `<p>${escape(singoEvent.name)}</p>`;
	})}</div>
${validate_component(Footer, 'Footer').$$render($$result, {}, {}, {})}`;
});
export { Page as default };
