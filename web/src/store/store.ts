import { derived, writable } from 'svelte/store';

export const apiData = writable([]);

const devApiLink = 'https://events-21r9.onrender.com/api/v1/';
const productionApiLink = 'https://dronejobs.co/';

export let apiLink: string;

// Change the operator to turn the flag off or on
const flag = true;

if (flag === true) {
	apiLink = devApiLink;
} else {
	apiLink = productionApiLink;
}

export const events = derived(apiData, ($apiData) => {
	if ($apiData) {
		return $apiData;
	}
	return [];
});

export const singleEvent = derived(apiData, ($apiData) => {
	if ($apiData) {
		return $apiData;
	}
	return [];
});
