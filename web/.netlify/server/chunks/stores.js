import { g as getContext } from './index.js';
import { d as assets, v as version, b as browser } from './paths.js';
import { w as writable } from './index2.js';
import { DEV, BROWSER } from 'esm-env';
function notifiable_store(value) {
	const store = writable(value);
	let ready = true;
	function notify() {
		ready = true;
		store.update((val) => val);
	}
	function set(new_value) {
		ready = false;
		store.set(new_value);
	}
	function subscribe(run) {
		let old_value;
		return store.subscribe((new_value) => {
			if (old_value === void 0 || (ready && new_value !== old_value)) {
				run((old_value = new_value));
			}
		});
	}
	return { notify, set, subscribe };
}
function create_updated_store() {
	const { set, subscribe } = writable(false);
	let timeout;
	async function check() {
		if (DEV || !BROWSER) return false;
		clearTimeout(timeout);
		const res = await fetch(`${assets}/${'_app/version.json'}`, {
			headers: {
				pragma: 'no-cache',
				'cache-control': 'no-cache'
			}
		});
		if (res.ok) {
			const data = await res.json();
			const updated = data.version !== version;
			if (updated) {
				set(true);
				clearTimeout(timeout);
			}
			return updated;
		} else {
			throw new Error(`Version check failed: ${res.status}`);
		}
	}
	return {
		subscribe,
		check
	};
}
const stores = {
	url: notifiable_store({}),
	page: notifiable_store({}),
	navigating: writable(null),
	updated: create_updated_store()
};
const getStores = () => {
	const stores$1 = browser ? stores : getContext('__svelte__');
	const readonly_stores = {
		page: {
			subscribe: stores$1.page.subscribe
		},
		navigating: {
			subscribe: stores$1.navigating.subscribe
		},
		updated: stores$1.updated
	};
	Object.defineProperties(readonly_stores, {
		preloading: {
			get() {
				console.error(
					'stores.preloading is deprecated; use stores.navigating instead'
				);
				return {
					subscribe: stores$1.navigating.subscribe
				};
			},
			enumerable: false
		},
		session: {
			get() {
				removed_session();
				return {};
			},
			enumerable: false
		}
	});
	return readonly_stores;
};
const page = {
	subscribe(fn) {
		const store = getStores().page;
		return store.subscribe(fn);
	}
};
function removed_session() {
	throw new Error(
		'stores.session is no longer available. See https://github.com/sveltejs/kit/discussions/5883'
	);
}
export { page as p };
