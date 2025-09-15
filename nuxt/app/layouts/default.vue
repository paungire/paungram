<script setup>
const route = useRoute();

const items = ref([
	{
		label: "Политика",
		to: "/policy/",
		active: route.path.startsWith("/policy/"),
	},
	{
		label: "Профиль",
		to: "/profile/",
		active: route.path.startsWith("/profile/"),
	},
]);
</script>

<template>
	<div>
		<UHeader title="паунграм" to="/" mode="drawer">
			<template #right>
				<UColorModeButton />
			</template>

			<UNavigationMenu :items="items" />
		</UHeader>
		<main>
			<slot />
		</main>
	</div>
</template>
