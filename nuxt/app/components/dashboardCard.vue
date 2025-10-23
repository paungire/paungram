<script setup>
const props = defineProps(["title", "link"]);
</script>
<template>
	<div
		class="bg-default rounded-3xl mb-8 break-inside-avoid p-4 sm:p-6 drop-shadow-xl"
	>
		<div>
			<router-link :to="props.link">{{ props.title }}</router-link>
			<slot></slot>
		</div>
	</div>
</template>
