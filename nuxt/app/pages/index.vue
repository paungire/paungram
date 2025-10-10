<script setup>
import { ref } from "vue";
import { useAuthStore } from "../stores/auth";
const { $apiFetch } = useNuxtApp();

const links = ref([
	{
		label: "погнали",
		to: "/auth/",
		icon: "pepicons-print:play",
	},
]);
</script>

<template>
	<div>
		<UPageHero
			title="паунграм"
			description="чатик для моих любимых друзьяшек"
			:links="links"
		/>
	</div>
</template>
