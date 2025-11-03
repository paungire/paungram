<script setup>
const props = defineProps(["icon", "text", "size", "color", "side", "photo"]);
</script>

<template>
	<btn-soft
		:icon="icon"
		:text="text"
		:size="size"
		:color="color"
		:side="side"
		:photo="photo"
	>
		<template #content>
			<div>
				<btn-glass>
					<btn-pie-icon
						:icon="icon"
						:size="size"
						:color="color"
						:side="side"
						:photo="photo"
					></btn-pie-icon>
				</btn-glass>
			</div>
		</template>
	</btn-soft>
</template>
