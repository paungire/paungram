export default defineNuxtConfig({
	compatibilityDate: "2025-09-11",
	css: ["~/assets/css/main.css"],
	modules: [
		"@nuxt/ui",
		"@nuxt/icon",
		"@pinia/nuxt",
		"pinia-plugin-persistedstate/nuxt",
		"nuxt-proxy-request",
	],
	proxy: {
		options: [
			{
				target: "http://localhost:8080",
				pathFilter: ["/api/v1/**"],
				pathRewrite: {
					"^/api/v1": "/api",
				},
				logLevel: "debug",
				changeOrigin: true,
			},
		],
	},
	ssr: true,
	// devServer: {
	// 	https: {
	// 		key: "./192.168.1.114+1-key.pem",
	// 		cert: "./192.168.1.114+1.pem",
	// 	},
	// },
});
