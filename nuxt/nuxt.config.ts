export default defineNuxtConfig({
	compatibilityDate: "2025-09-11",
	css: ["~/assets/css/main.css"],
	modules: ["@nuxt/ui", "@nuxt/icon"],
	ssr: true,
	devServer: {
		https: {
			key: "./192.168.1.114+1-key.pem",
			cert: "./192.168.1.114+1.pem",
		},
	},
});
