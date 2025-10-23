<script setup>
const itemsFooter = ref([
	{
		label: "Политика конфиденциальности",
		to: "/policy/",
		active: useRoute().path.startsWith("/policy/"),
	},
]);
</script>

<template>
	<UApp>
		<UMain
			:style="{
				minHeight: `calc(100vh - var(--ui-header-height) - var(--ui-footer-height, 80px))`,
			}"
		>
			<slot />
		</UMain>
	</UApp>
</template>
