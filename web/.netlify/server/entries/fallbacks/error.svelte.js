import {
	c as create_ssr_component,
	a as subscribe,
	e as escape
} from '../../chunks/index.js';
import { p as page } from '../../chunks/stores.js';
const Error = create_ssr_component(($$result, $$props, $$bindings, slots) => {
	let $page, $$unsubscribe_page;
	$$unsubscribe_page = subscribe(page, (value) => ($page = value));
	$$unsubscribe_page();
	return `<h1>${escape($page.status)}</h1>

<pre>${escape($page.error.message)}</pre>



${$page.error.frame ? `<pre>${escape($page.error.frame)}</pre>` : ``}
${$page.error.stack ? `<pre>${escape($page.error.stack)}</pre>` : ``}`;
});
export { Error as default };
