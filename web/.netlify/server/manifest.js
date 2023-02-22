export const manifest = {
	appDir: '_app',
	appPath: '_app',
	assets: new Set(['favicon.png', 'loginkey.svg']),
	mimeTypes: { '.png': 'image/png', '.svg': 'image/svg+xml' },
	_: {
		entry: {
			file: '_app/immutable/start-702200fe.js',
			imports: [
				'_app/immutable/start-702200fe.js',
				'_app/immutable/chunks/index-506043c6.js',
				'_app/immutable/chunks/singletons-f26365eb.js',
				'_app/immutable/chunks/index-196d5796.js'
			],
			stylesheets: [],
			fonts: []
		},
		nodes: [
			() => import('./nodes/0.js'),
			() => import('./nodes/1.js'),
			() => import('./nodes/2.js'),
			() => import('./nodes/3.js'),
			() => import('./nodes/4.js'),
			() => import('./nodes/5.js'),
			() => import('./nodes/6.js')
		],
		routes: [
			{
				id: '/',
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0], errors: [1], leaf: 2 },
				endpoint: null
			},
			{
				id: '/create-account',
				pattern: /^\/create-account\/?$/,
				params: [],
				page: { layouts: [0], errors: [1], leaf: 3 },
				endpoint: null
			},
			{
				id: '/events',
				pattern: /^\/events\/?$/,
				params: [],
				page: { layouts: [0], errors: [1], leaf: 4 },
				endpoint: null
			},
			{
				id: '/events/[id]',
				pattern: /^\/events\/([^/]+?)\/?$/,
				params: [{ name: 'id', optional: false, rest: false, chained: false }],
				page: { layouts: [0], errors: [1], leaf: 5 },
				endpoint: null
			},
			{
				id: '/login',
				pattern: /^\/login\/?$/,
				params: [],
				page: { layouts: [0], errors: [1], leaf: 6 },
				endpoint: null
			}
		],
		matchers: async () => {
			return {};
		}
	}
};
