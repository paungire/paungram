<script setup>
const props = defineProps(["passive"]);
</script>

<template>
	<div class="stick" :class="{ passive: props.passive }"></div>
</template>

<style scoped>
.stick {
	position: absolute;
	background-color: transparent;
	width: calc(100% - var(--spacing) * 6);
	height: 1px;
	bottom: -1px;
	left: calc(var(--spacing) * 3);
	transition: background-color 0.2s ease;
}
*:not(:last-of-type):not(:hover):not(:has(+ *:hover)) .stick {
	background-color: var(--border);
}
.stick.passive {
	background-color: var(--border);
}
</style>
