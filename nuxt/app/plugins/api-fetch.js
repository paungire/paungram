// plugins/apiFetch.client.ts
export default defineNuxtPlugin(() => {
	const auth = useAuthStore();

	const apiFetch = async (url, options = {}) => {
		try {
			const result = await $fetch(url, {
				...options,
				credentials: "include",
				headers: {
					...(options.headers || {}),
					Authorization: auth.getAccessToken()
						? `Bearer ${auth.getAccessToken()}`
						: "",
				},
			});
			return result;
		} catch (error) {
			// Если 401 — пробуем рефреш
			if (error?.response?.status === 401) {
				const newToken = await auth.refresh();
				if (!newToken) {
					auth.clearToken();
					return navigateTo("/auth");
				}
				// 🔥 Повторяем исходный запрос
				return await $fetch(url, {
					...options,
					credentials: "include",
					headers: {
						...(options.headers || {}),
						Authorization: `Bearer ${newToken}`,
					},
				});
			}

			throw error;
		}
	};

	return {
		provide: {
			apiFetch,
		},
	};
});
