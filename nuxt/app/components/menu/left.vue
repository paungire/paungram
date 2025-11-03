<script setup>
const route = useRoute();
const authStore = useAuthStore();

const servers = ref([
	{
		label: "сервер",
		photo: "https://i.pravatar.cc/80?img=10",
		to: "/servers/10",
	},
	{
		label: "сервер",
		photo: "https://i.pravatar.cc/80?img=11",
		to: "/servers/11",
	},
	{
		label: "сервер",
		photo: "https://i.pravatar.cc/80?img=12",
		to: "/servers/12",
	},
]);
const items = computed(() => {
	// route.path = route.path || "";
	if (authStore.userId && route.path) {
		return [
			{
				label: "Личные сообщения",
				icon: "mingcute:message-3-fill",
				to: "/id/" + authStore.userId + "/friends",
				active: route.path.startsWith("/"),
				color: "primary",
			},
			...servers.value.map((server) => ({
				...server,
				active: route.path.startsWith(server.to),
			})),
			{
				label: "Найти сервер",
				icon: "mingcute:search-line",
				to: "/",
				active: route.path.startsWith("/"),
				color: "primary",
			},
			{
				label: "Создать сервер",
				icon: "mingcute:add-fill",
				to: "/",
				active: route.path.startsWith("/"),
				color: "primary",
			},
		];
	}
});
</script>

<template>
	<div class="w-max flex flex-col gap-0.25">
		<router-link
			v-for="(item, k) of items"
			:key="k"
			:to="item.to"
			class="relative"
		>
			<btn-soft-glass
				:icon="item.icon"
				:photo="item.photo"
				size="32"
				:color="item.color"
				:text="item.label"
				side="right"
			></btn-soft-glass>
			<stick :passive="true" v-if="k == 0" />
		</router-link>
	</div>
</template>
