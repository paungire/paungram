<!-- components/GlassButton.vue -->
<script setup>
const props = defineProps({
	icon: String,
	text: String,
	size: { type: Number, default: 20 },
	photo: String,
	color: { type: String, default: "var(--ui-primary)" },
	glow: { type: Boolean, default: true },
});
</script>

<template>
	<button
		class="glass-btn group relative flex items-center justify-center gap-2 px-4 py-2 rounded-xl overflow-hidden transition-all duration-300"
	>
		<!-- Фон подсветки -->
		<div
			v-if="glow"
			class="absolute inset-0 rounded-xl bg-gradient-to-b from-white/15 to-white/5 opacity-0 group-hover:opacity-100 blur-md transition-opacity duration-300"
		></div>

		<!-- Контент -->
		<div class="relative z-10 flex items-center gap-2 select-none">
			<Icon
				v-if="icon"
				:name="icon"
				:size="size"
				class="transition-transform duration-200 text-[var(--ui-text-color)] group-hover:scale-105 group-hover:text-white"
			/>
			<img
				v-else-if="photo"
				:src="photo"
				alt=""
				class="w-8 h-8 rounded-full object-cover transition-transform duration-200 group-hover:scale-105"
			/>
			<span
				v-if="text"
				class="text-sm font-medium text-[var(--ui-text-color)] group-hover:text-white transition-colors duration-200"
			>
				{{ text }}
			</span>
		</div>

		<!-- Свет по краям -->
		<div
			class="absolute inset-0 rounded-xl pointer-events-none border border-white/15 shadow-[inset_0_1px_1px_rgba(255,255,255,0.4),0_4px_10px_rgba(0,0,0,0.15)]"
		></div>
	</button>
</template>

<style scoped>
.glass-btn {
	background: rgba(255, 255, 255, 0.05);
	backdrop-filter: blur(12px);
	-webkit-backdrop-filter: blur(12px);
	transform: translateY(0);
	transition: all 0.25s ease;
}

.glass-btn:hover {
	background: rgba(255, 255, 255, 0.1);
	transform: translateY(-1px);
	box-shadow: 0 0 12px rgba(255, 255, 255, 0.1),
		inset 0 0 6px rgba(255, 255, 255, 0.15);
}

.glass-btn:active {
	transform: translateY(1px);
	background: rgba(255, 255, 255, 0.08);
	box-shadow: inset 0 2px 6px rgba(0, 0, 0, 0.3),
		0 0 4px rgba(255, 255, 255, 0.05);
}
</style>
