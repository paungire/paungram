<script setup>
const route = useRoute();
const authStore = useAuthStore();

const items = computed(() => {
	if (authStore.userId) {
		return [
			{
				label: "Профиль",
				to: "/id/" + authStore.userId,
				active: route.path.startsWith("/id/"),
			},
			{
				label: "Друзья",
				to: "/id/" + authStore.userId + "/friends/",
				active: route.path.startsWith("/friends/"),
			},
		];
	}
});
</script>
