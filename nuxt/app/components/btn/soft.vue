<script setup>
const props = defineProps(["icon", "text", "size", "color", "side", "photo"]);
</script>

<template>
	<div>
		<slot name="content">
			<btn-pie-icon
				:icon="icon"
				:size="size"
				:color="color"
				:side="side"
				:photo="photo"
			></btn-pie-icon>
		</slot>
	</div>
</template>
