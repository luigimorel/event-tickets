import { BROWSER } from 'esm-env';
let version = '';
function set_building(value) {}
function set_version(value) {
	version = value;
}
const browser = BROWSER;
let base = '';
let assets = '';
function set_paths(paths) {
	base = paths.base;
	assets = paths.assets || base;
}
export {
	set_building as a,
	browser as b,
	base as c,
	assets as d,
	set_version as e,
	set_paths as s,
	version as v
};
