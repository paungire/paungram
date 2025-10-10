import { useAuthStore } from "../stores/auth";

export default defineNuxtPlugin(() => {
	const auth = useAuthStore();
	// просто вызов сторы — этого хватит, чтобы persist сработал до мидлвеаров
});
