<script setup>
const props = defineProps(["icon", "photo", "size", "color"]);

const colorClass = computed(() => {
	if (props.color) {
		return `text-${props.color}`;
	} else {
		return "text-[var(--ui-primary)]";
	}
});
</script>

<template>
	<div>
		<div v-if="icon" class="p-4">
			<Icon
				:class="colorClass"
				class="cursor-pointer block"
				:name="icon"
				:size="size ?? 20"
			></Icon>
		</div>
		<img v-else :src="photo" alt="" class="w-16 rounded-full p-2" />
	</div>
</template>
