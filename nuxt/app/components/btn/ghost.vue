<script setup>
const props = defineProps(["icon", "text", "size", "color", "side", "photo"]);
</script>

<template>
	<UTooltip
		:text="text"
		:delay-duration="0.1"
		:content="{
			align: 'center',
			side: side ?? 'top',
		}"
	>
		<slot name="content">
			<btn-pie-icon v-bind="props"></btn-pie-icon>
		</slot>
	</UTooltip>
</template>
