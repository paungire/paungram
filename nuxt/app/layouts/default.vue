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
		<UHeader title="паунграм" to="/" mode="drawer">
			<template #right>
				<UColorModeButton />
			</template>
		</UHeader>

		<UMain
			:style="{
				minHeight: `calc(100vh - var(--ui-header-height) - var(--ui-footer-height, 80px))`,
			}"
		>
			<slot />
		</UMain>
		<UFooter>
			<template #left>
				<p class="text-muted text-sm">
					Copyright © {{ new Date().getFullYear() }}
				</p>
			</template>

			<UNavigationMenu :items="itemsFooter" variant="link" />

			<template #right>
				<UButton
					icon="i-simple-icons-telegram"
					color="neutral"
					variant="ghost"
					to="https://go.nuxt.com/discord"
					target="_blank"
					aria-label="Telegram"
				/>
			</template>
		</UFooter>
	</UApp>
</template>
