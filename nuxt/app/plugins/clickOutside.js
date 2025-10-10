// plugins/clickOutside.client.js
import { defineNuxtPlugin } from "#app";

export default defineNuxtPlugin((nuxtApp) => {
	nuxtApp.vueApp.directive("click-outside", {
		beforeMount(el, binding) {
			el.__clickOutsideHandler__ = (event) => {
				console.log(el.contains(event.target));
				if (!(el === event.target || el.contains(event.target))) {
					binding.value && binding.value(event);
				}
			};
			document.addEventListener("click", el.__clickOutsideHandler__);
		},
		unmounted(el) {
			document.removeEventListener("click", el.__clickOutsideHandler__);
			delete el.__clickOutsideHandler__;
		},
	});
});
