import { derived, writable } from 'svelte/store';

export const apiData = writable([]);

const devApiLink = 'http://localhost:8080/api/v1/';
const productionApiLink = 'https://dronejobs.co/';

export let apiLink: string;

// Change the operator to turn the flag off or on
const flag = false;

if (flag === false) {
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
