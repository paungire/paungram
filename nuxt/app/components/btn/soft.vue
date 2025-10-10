<script setup>
defineProps({
	icon: String,
	text: String,
});
</script>

<template>
	<UTooltip
		:text="text"
		:delay-duration="0"
		:content="{
			align: 'center',
			side: 'top',
		}"
	>
		<UButton
			class="rounded-full p-2 cursor-pointer text-muted dark:bg-[var(--ui-bg)] dark:hover:bg-[var(--ui-bg)] dark:hover:text-white light:bg-[var(--ui-bg)] light:hover:bg-[var(--ui-bg)] light:hover:text-black"
			:icon="icon"
			color="neutral"
			variant="ghost"
		/>
	</UTooltip>
</template>
