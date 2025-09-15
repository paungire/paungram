export default defineNuxtRouteMiddleware((to, from) => {
	// TODO
	const userIsAuthenticated = () => {
		return false;
	};
	if (!userIsAuthenticated()) {
		return navigateTo("/auth");
	}
});
